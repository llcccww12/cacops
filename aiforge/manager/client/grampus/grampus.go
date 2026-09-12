package grampus

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/otel"
	"code.gitea.io/gitea/modules/setting"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
)

var (
	restyClient *resty.Client
	HOST        string
	TOKEN       string
)

const (
	urlOpenApiV1 = "/openapi/v1/"

	urlGetToken         = urlOpenApiV1 + "token"
	urlTrainJob         = urlOpenApiV1 + "trainjob"
	urlGetResourceSpecs = urlOpenApiV1 + "resourcespec"
	urlGetAiCenter      = urlOpenApiV1 + "sharescreen/aicenter"
	urlGetImages        = urlOpenApiV1 + "image"
	urlListUserImages   = urlOpenApiV1 + "listUserImage"
	urlNotebookJob      = urlOpenApiV1 + "notebook"
	urlInferenceJob     = urlOpenApiV1 + "inference"
	urlModelService     = urlOpenApiV1 + "serviceModel"
	urlModelAppService  = urlOpenApiV1 + "service"

	errorIllegalToken             = 1005
	errorCannotStopCreatingJob    = 5008
	errorCannotStopSavingImageJob = 5009
	MAX_MATRICS_SIZE              = 1000
	errorNotSupportResouceUsage   = 6003
)

type GetTokenParams struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type GetTokenResult struct {
	Token      string `json:"token"`
	Expiration int64  `json:"expiration"`
}

func getRestyClient() *resty.Client {
	if restyClient == nil {
		restyClient = resty.New()
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return restyClient
}

func checkSetting() {
	if len(HOST) != 0 && len(TOKEN) != 0 && restyClient != nil {
		return
	}

	err := getToken()
	if err != nil {
		log.Error("getToken failed:%v", err)
	}
}

func getToken() error {
	HOST = strings.TrimSuffix(setting.Grampus.Host, "/")

	client := getRestyClient()
	params := GetTokenParams{
		UserName: setting.Grampus.UserName,
		Password: setting.Grampus.Password,
	}

	var result GetTokenResult
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(params).
		SetResult(&result).
		Post(HOST + urlGetToken)
	if err != nil {
		return fmt.Errorf("resty getToken: %v", err)
	}

	if res.StatusCode() != http.StatusOK {
		return fmt.Errorf("getToken failed:%s", res.String())
	}

	TOKEN = result.Token

	return nil
}

func GetToken() error {
	return getToken()
}

func CreateInferenceJob(req models.CreateGrampusInferenceRequest, traceInfo *entity.TraceInfo) (*models.GrampusNotebookResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GrampusNotebookResponse
	reqJson, _ := json.Marshal(req)
	log.Info("Online infer REQ:" + string(reqJson))
	retry := 0
	var err error
	childCtx, span := otel.StartTraceChild(traceInfo, "POST")

	defer func() {
		if err != nil {
			otel.RemoveTraceCache(traceInfo.TaskId, otel.GenerateCacheValue(span))
		}
		otel.FinalizeSpan(childCtx, err)
	}()

sendjob:
	_, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("traceparent", otel.GenerateHeader(span)).
		SetAuthToken(TOKEN).
		SetBody(req).
		SetResult(&result).
		Post(HOST + urlInferenceJob)

	if err != nil {
		log.Error("resty CreateInferenceJob: %v", err)
		return nil, models.NetworkError{}
	}

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("CreateInferenceJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("CreateNotebookJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}
	log.Info("CreateInferenceJob success.req.JobName = %s ,result=%+v", req.Name, result)
	return &result, nil
}

func GetInferenceJob(opts *entity.JobIdAndVersionId) (*models.GrampusNotebookResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GrampusNotebookResponse

	retry := 0
	var err error
	newCtx, childCtx, span := otel.StartTrace(&entity.TraceInfo{TaskId: opts.TaskID, SpanName: "QueryTask"}, false)

	defer func() {
		if err != nil {
			otel.RemoveTraceCache(opts.TaskID, otel.GenerateCacheValue(span))
		}
		otel.FinalizeSpan(childCtx, err)
		otel.FinalizeSpan(newCtx, err)
	}()

sendjob:
	body, err := client.R().
		SetHeader("traceparent", otel.GenerateHeader(span)).
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + urlInferenceJob + "/" + opts.JobID)

	if err != nil {
		return nil, fmt.Errorf("resty GetNotebookJob: %v", err)
	}
	log.Info("%+v", body)
	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("GetNotebookJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return nil, fmt.Errorf("GetNotebookJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}
func CreateNotebookJob(req models.CreateGrampusNotebookRequest, traceInfo *entity.TraceInfo) (*models.GrampusNotebookResponse, error) {

	checkSetting()
	client := getRestyClient()
	var result models.GrampusNotebookResponse
	reqJson, _ := json.Marshal(req)
	log.Info("grampus notebook req:" + string(reqJson))
	retry := 0
	var err error

	childCtx, span := otel.StartTraceChild(traceInfo, "POST")

	defer func() {
		if err != nil {
			otel.RemoveTraceCache(traceInfo.TaskId, otel.GenerateCacheValue(span))
		}
		otel.FinalizeSpan(childCtx, err)
	}()

sendjob:
	_, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("traceparent", otel.GenerateHeader(span)).
		SetAuthToken(TOKEN).
		SetBody(req).
		SetResult(&result).
		Post(HOST + urlNotebookJob)

	if err != nil {
		log.Error("resty CreateNotebookJob: %v", err)
		return nil, models.NetworkError{}
	}

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("CreateNotebookJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("CreateNotebookJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}
	log.Info("CreateNotebookJob success.req.JobName = %s ,result=%+v", req.Name, result)

	return &result, nil
}

func CreateJob(req models.CreateGrampusJobRequest, traceInfo *entity.TraceInfo) (*models.CreateGrampusJobResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.CreateGrampusJobResponse

	retry := 0
	log.Info("grampus job req:%+v", req)
	var err error

	childCtx, span := otel.StartTraceChild(traceInfo, "POST")

	defer func() {
		if err != nil {
			otel.RemoveTraceCache(traceInfo.TaskId, otel.GenerateCacheValue(span))
		}
		otel.FinalizeSpan(childCtx, err)
	}()
sendjob:
	_, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("traceparent", otel.GenerateHeader(span)).
		SetAuthToken(TOKEN).
		SetBody(req).
		SetResult(&result).
		Post(HOST + urlTrainJob)

	if err != nil {
		return nil, fmt.Errorf("resty CreateJob: %s", err)
	}

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("CreateJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("CreateJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func GetNotebookJob(opts *entity.JobIdAndVersionId) (*models.GrampusNotebookResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GrampusNotebookResponse

	retry := 0
	var err error
	newCtx, childCtx, span := otel.StartTrace(&entity.TraceInfo{TaskId: opts.TaskID, SpanName: "QueryTask"}, false)

	defer func() {
		if err != nil {
			otel.RemoveTraceCache(opts.TaskID, otel.GenerateCacheValue(span))
		}
		otel.FinalizeSpan(childCtx, err)
		otel.FinalizeSpan(newCtx, err)
	}()

sendjob:
	body, err := client.R().
		SetHeader("traceparent", otel.GenerateHeader(span)).
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + urlNotebookJob + "/" + opts.JobID)

	if err != nil {
		return nil, fmt.Errorf("resty GetNotebookJob: %v", err)
	}
	log.Info("%+v", body)
	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("GetNotebookJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return nil, fmt.Errorf("GetNotebookJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func GetJobs(jobIds []string) (*models.GetGrampusJobListResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetGrampusJobListResponse

	retry := 0
	var err error
	newCtx, childCtx, span := otel.StartTraceSimple("QueryTaskList")

	defer func() {

		otel.FinalizeSpan(childCtx, err)
		otel.FinalizeSpan(newCtx, err)
	}()

sendjob:
	r, err := client.R().
		SetAuthToken(TOKEN).
		SetHeader("traceparent", otel.GenerateHeader(span)).
		SetQueryParam("ids", strings.Join(jobIds, ",")).
		SetResult(&result).
		Get(HOST + urlTrainJob)
	if err != nil {
		return nil, fmt.Errorf("resty GetJob: %v", err)
	}

	log.Info("%+v", r)

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("GetJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return nil, fmt.Errorf("GetJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func GetJob(opts *entity.JobIdAndVersionId) (*models.GetGrampusJobResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetGrampusJobResponse

	retry := 0
	var err error
	newCtx, childCtx, span := otel.StartTrace(&entity.TraceInfo{TaskId: opts.TaskID, SpanName: "QueryTask"}, false)

	defer func() {
		if err != nil {
			otel.RemoveTraceCache(opts.TaskID, otel.GenerateCacheValue(span))
		}
		otel.FinalizeSpan(childCtx, err)
		otel.FinalizeSpan(newCtx, err)
	}()

sendjob:

	_, err = client.R().
		SetHeader("traceparent", otel.GenerateHeader(span)).
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + urlTrainJob + "/" + opts.JobID)
	if err != nil {
		return nil, fmt.Errorf("resty GetJob: %v", err)
	}

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("GetJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return nil, fmt.Errorf("GetJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func GetJobListByJobName(jobName string) (*models.GetGrampusJobListResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetGrampusJobListResponse

	retry := 0

sendjob:
	_, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + urlTrainJob + "?pageIndex=1&pageSize=20&searchKey=" + jobName)
	if err != nil {
		return nil, fmt.Errorf("resty GetJobListByJobName: %v", err)
	}

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("GetJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return nil, fmt.Errorf("GetJobListByJobName failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func GetResourceSpecs(processorType string) (*models.GetGrampusResourceSpecsResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetGrampusResourceSpecsResult

	retry := 0

sendjob:
	_, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + urlGetResourceSpecs + "?processorType=" + processorType)

	if err != nil {
		return nil, fmt.Errorf("resty GetResourceSpecs: %v", err)
	}

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("GetResourceSpecs failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("GetResourceSpecs failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func GetImages(processorType string, jobType string) (*models.GetGrampusImagesResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetGrampusImagesResult

	retry := 0

	queryType := "TrainJob"
	if jobType == string(models.JobTypeDebug) {
		queryType = "Notebook"
	}

sendjob:
	_, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + urlGetImages + "?processorType=" + processorType + "&trainType=" + queryType)

	if err != nil {
		return nil, fmt.Errorf("resty GetImages: %v", err)
	}

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}
	jsonstr, _ := json.Marshal(result)
	log.Info("GetImages resp:", string(jsonstr))

	if result.ErrorCode != 0 {
		log.Error("GetImages failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("GetImages failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func GetUserImages(processorType string, jobType string) (*models.GetGrampusImagesResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetGrampusImagesResult

	retry := 0

	queryType := "TrainJob"
	if jobType == string(models.JobTypeDebug) {
		queryType = "Notebook"
	}

sendjob:
	_, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + urlListUserImages + "?processorType=" + processorType + "&trainType=" + queryType)

	if err != nil {
		return nil, fmt.Errorf("resty GetUserImages: %v", err)
	}

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}
	jsonstr, _ := json.Marshal(result)
	log.Info("GetUserImages resp:", string(jsonstr))

	if result.ErrorCode != 0 {
		log.Error("GetUserImages failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("GetUserImages failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func GetTrainJobLog(jobID string, nodeId ...int) (string, error) {
	checkSetting()
	client := getRestyClient()
	var logContent string

	url := HOST + urlTrainJob + "/" + jobID + "/task/0/replica/0/log"
	if len(nodeId) > 0 {
		url = HOST + urlTrainJob + "/" + jobID + "/task/0/replica/0/log/node/" + strconv.Itoa(nodeId[0])
	}

	res, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&logContent).
		Get(url)

	if err != nil {
		return logContent, fmt.Errorf("resty GetTrainJobLog: %v", err)
	}

	if res.StatusCode() != http.StatusOK {
		var temp models.GrampusResult
		if err = json.Unmarshal([]byte(res.String()), &temp); err != nil {
			log.Error("json.Unmarshal failed(%s): %v", res.String(), err.Error())
			return logContent, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
		}
		log.Error("GetTrainJobLog failed(%d):%s(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
		return logContent, fmt.Errorf("GetTrainJobLog failed(%d):%d(%s)", res.StatusCode(), temp.ErrorCode, temp.ErrorMsg)
	}

	logContent = res.String()

	return logContent, nil
}

func GetGrampusMetrics(jobID string, startTime int64, endTime int64, nodeId ...int) (models.NewModelArtsMetricStatisticResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.NewModelArtsMetricStatisticResult
	var grampusResult models.GrampusMetricStatisticResult
	url := HOST + urlTrainJob + "/" + jobID + "/task/0/replica/0/metrics"
	if len(nodeId) > 0 {
		url = HOST + urlTrainJob + "/" + jobID + "/task/0/replica/0/metrics/node/" + strconv.Itoa(nodeId[0])
	}
	var step int64 = 60
	if startTime > 0 {

		size := int64(math.Ceil(float64(endTime-startTime)/float64(step))) + 1

		if size > MAX_MATRICS_SIZE {
			step = int64(math.Ceil(float64(size*step) / float64(MAX_MATRICS_SIZE)))
			tempStep := step / 60
			if step%60 != 0 {
				step = (tempStep + 1) * 60
			} else {
				step = tempStep * 60
			}
			size = int64(math.Ceil(float64(endTime-startTime)/float64(step))) + 1

		}

		url = url + "?startTime=" + strconv.FormatInt(startTime, 10) + "&step=" + strconv.FormatInt(step, 10) + "&size=" + strconv.FormatInt(size, 10)
	}
	res, err := client.R().
		SetAuthToken(TOKEN).
		Get(url)

	if err != nil {
		return result, fmt.Errorf("resty GetTrainJobLog: %v", err)
	}

	if err = json.Unmarshal([]byte(res.String()), &grampusResult); err != nil {
		log.Error("GetGrampusMetrics json.Unmarshal failed(%s): %v", res.String(), err.Error())
		return result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
	}
	if res.StatusCode() != http.StatusOK {
		return result, fmt.Errorf("Call GrampusMetrics failed(%d)", res.StatusCode())
	}
	if grampusResult.ErrorCode == errorNotSupportResouceUsage {

		return result, fmt.Errorf(strconv.Itoa(errorNotSupportResouceUsage))
	}
	result = models.NewModelArtsMetricStatisticResult{
		MetricsInfo: grampusResult.MetricsInfo,
		Step:        step,
	}
	return result, nil
}

func StopJob(opts *entity.JobIdAndVersionId) (*models.GrampusStopJobResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GrampusStopJobResponse

	retry := 0

	url := urlTrainJob
	if opts.JobType == string(models.JobTypeDebug) {
		url = urlNotebookJob
	}

	var err error

	newCtx, childCtx, span := otel.StartTrace(&entity.TraceInfo{TaskId: opts.TaskID, SpanName: "StopTask"}, false, "POST")

	defer func() {
		if err != nil {
			otel.RemoveTraceCache(opts.TaskID, otel.GenerateCacheValue(span))
		}
		otel.FinalizeSpan(childCtx, err)
		otel.FinalizeSpan(newCtx, err)
	}()

sendjob:
	_, err = client.R().
		SetHeader("traceparent", otel.GenerateHeader(span)).
		SetAuthToken(TOKEN).
		SetResult(&result).
		Post(HOST + url + "/" + opts.JobID + "/stop")

	if err != nil {
		return &result, fmt.Errorf("resty StopTrainJob: %v", err)
	}

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("GetJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		if result.ErrorCode == errorCannotStopCreatingJob {
			return &result, models.ErrCannotStopCreatingGrampusJob{}
		}
		if result.ErrorCode == errorCannotStopSavingImageJob {
			return &result, models.ErrCannotStopSavingImageJob{}
		}
		return &result, fmt.Errorf("GetJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func GetAiCenters(pageIndex, pageSize int) (*models.GetGrampusAiCentersResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetGrampusAiCentersResult

	retry := 0

sendjob:
	_, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + urlGetAiCenter + "?pageIndex=" + fmt.Sprint(pageIndex) + "&pageSize=" + fmt.Sprint(pageSize))

	if err != nil {
		return nil, fmt.Errorf("resty GetAiCenters: %v", err)
	}

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("GetAiCenters failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("GetAiCenters failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func RestartNotebookJob(jobID string, autoStopDuration int64, traceInfo *entity.TraceInfo) (*models.GrampusNotebookRestartResponse, error) {
	checkSetting()
	client := getRestyClient()
	var restartResponse *models.GrampusNotebookRestartResponse
	retry := 0

	var err error
	newCtx, childCtx, span := otel.StartTrace(traceInfo, true, "POST")

	defer func() {
		if err != nil {
			otel.RemoveTraceCache(traceInfo.TaskId, otel.GenerateCacheValue(span))
		}
		otel.FinalizeSpan(childCtx, err)
		otel.FinalizeSpan(newCtx, err)
	}()

sendjob:
	res, err := client.R().
		SetHeader("traceparent", otel.GenerateHeader(span)).
		SetAuthToken(TOKEN).
		SetResult(&restartResponse).
		Post(HOST + urlNotebookJob + "/" + jobID + "/start?autoStopDuration=" + strconv.FormatInt(autoStopDuration, 10))

	if err != nil {
		return nil, fmt.Errorf("resty grampus restart note book job: %v", err)
	}
	if restartResponse.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		log.Error("resty grampus restart note book job failed(%s): %v", res.String(), err.Error())
		return nil, fmt.Errorf("resty grampus restart note book job failed: %v", err)
	}
	log.Info("RestartNotebookJob success.jobId = %s ,result=%+v", jobID, restartResponse)

	return restartResponse, nil
}

func GetDebugJobEvents(opts *entity.JobIdAndVersionId) (*models.GetGrampusDebugJobEventsResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetGrampusDebugJobEventsResponse

	retry := 0
	var err error
	newCtx, childCtx, span := otel.StartTrace(&entity.TraceInfo{TaskId: opts.TaskID, SpanName: "Operational Overview"}, false)

	defer func() {
		if err != nil {
			otel.RemoveTraceCache(opts.TaskID, otel.GenerateCacheValue(span))
		}
		otel.FinalizeSpan(childCtx, err)
		otel.FinalizeSpan(newCtx, err)
	}()

sendjob:
	res, err := client.R().
		SetAuthToken(TOKEN).
		SetHeader("traceparent", otel.GenerateHeader(span)).
		SetResult(&result).
		Get(HOST + urlNotebookJob + "/" + opts.JobID + "/events")
	log.Info("res=%v", res)
	if err != nil {
		return nil, fmt.Errorf("resty GetDebugJobEvents: %v", err)
	}

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("GetDebugJobEvents failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return nil, fmt.Errorf("GetDebugJobEvents failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func GetTrainJobEvents(opts *entity.JobIdAndVersionId) (*models.GetGrampusJobEventsResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetGrampusJobEventsResponse

	retry := 0
	var err error
	newCtx, childCtx, span := otel.StartTrace(&entity.TraceInfo{TaskId: opts.TaskID, SpanName: "Operational Overview"}, false)

	defer func() {
		if err != nil {
			otel.RemoveTraceCache(opts.TaskID, otel.GenerateCacheValue(span))
		}
		otel.FinalizeSpan(childCtx, err)
		otel.FinalizeSpan(newCtx, err)
	}()

sendjob:
	res, err := client.R().
		SetAuthToken(TOKEN).
		SetHeader("traceparent", otel.GenerateHeader(span)).
		SetResult(&result).
		Get(HOST + urlTrainJob + "/" + opts.JobID + "/events")
	if err != nil {
		return nil, fmt.Errorf("resty GetTrainJobEvents: %v", err)
	}
	log.Info("res=%+v", res)

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("GetTrainJobEvents failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return nil, fmt.Errorf("GetTrainJobEvents failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func DeleteJob(opts *entity.JobIdAndVersionId) (*models.GrampusDeleteJobResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GrampusDeleteJobResponse

	retry := 0

	url := urlTrainJob
	if opts.JobType != "" {

		if opts.JobType == string(models.JobTypeDebug) {
			url = urlNotebookJob
		} else if opts.JobType == string(models.JobTypeInference) {
			url = urlInferenceJob
		} else if opts.JobType == string(models.JobTypeModelExperience) {
			url = urlInferenceJob
		}
	}

	var err error
	newCtx, childCtx, span := otel.StartTrace(&entity.TraceInfo{TaskId: opts.TaskID, SpanName: "DeleteTask"}, false, "DELETE")

	defer func() {
		if err != nil {
			otel.RemoveTraceCache(opts.TaskID, otel.GenerateCacheValue(span))
		}
		otel.FinalizeSpan(childCtx, err)
		otel.FinalizeSpan(newCtx, err)
	}()

sendjob:
	_, err = client.R().
		SetHeader("traceparent", otel.GenerateHeader(span)).
		SetAuthToken(TOKEN).
		SetResult(&result).
		Delete(HOST + url + "/" + opts.JobID)

	if err != nil {
		return &result, fmt.Errorf("resty StopTrainJob: %v", err)
	}

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("Delete Job failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("Delete Job failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}
	log.Info("delete grampus task, re=" + result.Info + "  grampus jobId=" + opts.JobID)
	return &result, nil
}

func GetAvailableModelServices() (*models.GrampusServiceModels, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GrampusServiceModels

	retry := 0

sendjob:
	res, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + urlModelService)
	log.Info("res=%v", res)
	if err != nil {
		return nil, fmt.Errorf("resty GetModelService: %v", err)
	}

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("GetModelService failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return nil, fmt.Errorf("GetModelService failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}
	if res.StatusCode() != http.StatusOK {
		return &result, fmt.Errorf("Call GetModelService failed(%d)", res.StatusCode())
	}

	return &result, nil

}

func CreateModelAppServiceJob(req models.CreateGrampusModelAppRequest, traceInfo *entity.TraceInfo) (*models.GrampusNotebookResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GrampusNotebookResponse
	reqJson, _ := json.Marshal(req)
	log.Info("Online model service REQ:" + string(reqJson))
	retry := 0
	var err error
	childCtx, span := otel.StartTraceChild(traceInfo, "POST")

	defer func() {
		if err != nil {
			otel.RemoveTraceCache(traceInfo.TaskId, otel.GenerateCacheValue(span))
		}
		otel.FinalizeSpan(childCtx, err)
	}()

sendjob:
	res, err := client.R().
		SetHeader("traceparent", otel.GenerateHeader(span)).
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetBody(req).
		SetResult(&result).
		Post(HOST + urlModelAppService)

	if err != nil {
		log.Error("resty CreateModelAppJob: %v", err)
		return nil, models.NetworkError{}
	}

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("CreateModelAppJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("CreateModelAppJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}
	if res.StatusCode() != http.StatusOK {
		return &result, fmt.Errorf("CreateModelAppJob failed(%d)", res.StatusCode())
	}
	log.Info("CreateModelAppJob success.req.JobName = %s ,result=%+v", req.Name, result)
	return &result, nil
}

func GetModelAppServiceJob(opts *entity.JobIdAndVersionId) (*models.GrampusNotebookResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GrampusNotebookResponse

	retry := 0
	var err error
	newCtx, childCtx, span := otel.StartTrace(&entity.TraceInfo{TaskId: opts.TaskID, SpanName: "QueryTask"}, false)

	defer func() {
		if err != nil {
			otel.RemoveTraceCache(opts.TaskID, otel.GenerateCacheValue(span))
		}
		otel.FinalizeSpan(childCtx, err)
		otel.FinalizeSpan(newCtx, err)
	}()

sendjob:
	_, err = client.R().
		SetHeader("traceparent", otel.GenerateHeader(span)).
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + urlModelAppService + "/" + opts.JobID)
	if err != nil {
		return nil, fmt.Errorf("resty GetModelAppServiceJob: %v", err)
	}

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("GetModelAppServiceJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return nil, fmt.Errorf("GetModelAppServiceJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}
func StopModelAppServiceJob(opts *entity.JobIdAndVersionId) (*models.GrampusStopJobResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GrampusStopJobResponse

	retry := 0

	url := urlModelAppService
	var err error
	newCtx, childCtx, span := otel.StartTrace(&entity.TraceInfo{TaskId: opts.TaskID, SpanName: "StopTask"}, false, "POST")

	defer func() {
		if err != nil {
			otel.RemoveTraceCache(opts.TaskID, otel.GenerateCacheValue(span))
		}
		otel.FinalizeSpan(childCtx, err)
		otel.FinalizeSpan(newCtx, err)
	}()

sendjob:
	_, err = client.R().
		SetHeader("traceparent", otel.GenerateHeader(span)).
		SetAuthToken(TOKEN).
		SetResult(&result).
		Post(HOST + url + "/" + opts.JobID + "/stop")

	if err != nil {
		return &result, fmt.Errorf("resty StopModelAppServiceJob: %v", err)
	}

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("StopModelAppServiceJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		if result.ErrorCode == errorCannotStopCreatingJob {
			return &result, &models.ErrCannotStopCreatingGrampusJob{}
		}
		return &result, fmt.Errorf("StopModelAppServiceJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}
func DeleteModelAppServiceJob(opts *entity.JobIdAndVersionId) (*models.GrampusDeleteJobResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GrampusDeleteJobResponse

	retry := 0
	var err error
	newCtx, childCtx, span := otel.StartTrace(&entity.TraceInfo{TaskId: opts.TaskID, SpanName: "DeleteTask"}, false, "DELETE")

	defer func() {
		if err != nil {
			otel.RemoveTraceCache(opts.TaskID, otel.GenerateCacheValue(span))
		}
		otel.FinalizeSpan(childCtx, err)
		otel.FinalizeSpan(newCtx, err)
	}()
sendjob:
	_, err = client.R().
		SetHeader("traceparent", otel.GenerateHeader(span)).
		SetAuthToken(TOKEN).
		SetResult(&result).
		Delete(HOST + urlModelAppService + "/" + opts.JobID)

	if err != nil {
		return &result, fmt.Errorf("resty Delete ModelApp job: %v", err)
	}

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("Delete Job failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("Delete Job failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}
	log.Info("delete grampus task, re=" + result.Info + "  grampus jobId=" + opts.JobID)
	return &result, nil

}

type ModelAppApiSetKey struct {
	AK       string `json:"ak"`
	SK       string `json:"sk"`
	EndPoint string `json:"ep"`
}
type ModelAppCopyLoadModelReq struct {
	OBSDir   string `json:"obs_dir"`
	Template string `json:"template"`
}

type ModelAppCLoadModelStatusRes struct {
	models.BaseMessageApi
	models.GrampusResult
}

type ChatParameter struct {
	Messages          []models.PayloadText `json:"messages" binding:"Required"`
	MaxTokens         int                  `json:"max_tokens,omitempty"`
	Temperature       float64              `json:"temperature,omitempty"`
	TopP              float64              `json:"top_p,omitempty"`
	RepetitionPenalty float64              `json:"repetition_penalty,omitempty"`
}

func ModelAppSetKey(jobId string) error {
	getToken()
	checkSetting()
	client := getRestyClient()

	var result models.BaseMessageApi
	setkeyBody := ModelAppApiSetKey{
		AK:       setting.AccessKeyID,
		SK:       setting.SecretAccessKey,
		EndPoint: setting.Endpoint,
	}

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetBody(setkeyBody).
		SetResult(&result).
		Post(HOST + urlModelAppService + "/" + jobId + "/forward/setkey")

	if err != nil {
		return fmt.Errorf("resty  ModelApp setkey: %v", err)
	}
	log.Info("ModelAppSetKeyRes:%v", result)
	log.Info("ModelAppSetKeyRes:%v", res)

	if res.StatusCode() != http.StatusOK {
		return fmt.Errorf("ModelApp setkey failed(%d)", res.StatusCode())
	}
	return nil

}
func ModelAppCopyLoadModel(jobId string, modelId string, sourceCloudBrainId int64, template string, modelResultRelativePath string) error {
	if modelId == "" && sourceCloudBrainId == 0 {
		return fmt.Errorf("ModelAppCopyLoadModel failed. modelId is empty,sourceCloudBrainId=0")
	}
	tempModelPath := ""
	if modelId != "" {
		aiModel, err := models.QueryModelById(modelId)
		if err != nil {
			return err
		}

		tempModelPath = aiModel.Path

	} else {
		sourceCloudbrain, err := models.GetCloudbrainByCloudbrainID(sourceCloudBrainId)
		if err != nil {
			return err
		}

		if sourceCloudbrain.Cleared {
			return fmt.Errorf("ModelAppCopyLoadModel failed. the task is cleared.")
		}
		c := sourceCloudbrain.GetCloudbrainConfig()
		if c != nil {
			tempModelPath = path.Join(setting.Bucket, c.OutputObjectPrefix)
		}
		if modelResultRelativePath != "" {
			tempModelPath = path.Join(tempModelPath, modelResultRelativePath)
		}
	}
	tempModelPath = strings.TrimSuffix(tempModelPath, "/")

	checkSetting()
	client := getRestyClient()

	body := ModelAppCopyLoadModelReq{
		OBSDir:   "obs://" + tempModelPath,
		Template: template,
	}
	log.Info("ModelAppCopyLoadModelReq:%v", body)
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetBody(body).
		Post(HOST + urlModelAppService + "/" + jobId + "/forward/deployingservices")

	if err != nil {
		return fmt.Errorf("resty  ModelApp Loadmodel: %v", err)
	}

	log.Info("ModelAppCopyModelRes:%v", res)
	//ignore 502,504 error,because  user annother api check load status
	if res.StatusCode() != http.StatusOK && res.StatusCode() != http.StatusGatewayTimeout && res.StatusCode() != http.StatusBadGateway {
		return fmt.Errorf("ModelApp Loadmodel failed(%d)", res.StatusCode())
	}
	return nil

}

func CheckModelLoadStatus(jobId string) error {
	checkSetting()
	var result ModelAppCLoadModelStatusRes
	client := getRestyClient()
	retry := 0

sendjob:
	res, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + urlModelAppService + "/" + jobId + "/forward/querystatus")
	if err != nil {
		return fmt.Errorf("resty  ModelApp setkey: %v", err)
	}

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}

	if res.StatusCode() != http.StatusOK {
		return fmt.Errorf("ModelApp setkey failed(%d)", res.StatusCode())
	}

	if result.Code != 0 {
		log.Error("CheckModelLoadStatus failed: %s", result.Code, result.Message)
		return fmt.Errorf("CheckModelLoadStatus failed: %s", result.Message)
	}
	return nil

}

func ModelAppServiceCopyAndLoadModel(jobID string, modelId string, sourceCloudBrainId int64, template string, modelResultRelativePath string) error {
	err := ModelAppSetKey(jobID)
	if err != nil {
		return err
	}
	err = ModelAppCopyLoadModel(jobID, modelId, sourceCloudBrainId, template, modelResultRelativePath)
	if err != nil {
		return err
	}
	//15 minutes timeout
	attemps := 90
	for i := 0; i < attemps; i++ {
		time.Sleep(10 * time.Second)
		err = CheckModelLoadStatus(jobID)
		if err == nil {
			return nil
		}
	}

	return fmt.Errorf("CheckModelLoadStatus failed: Timeout")

}

func ModelAppInference(jobId string, chatParameter ChatParameter) (*resty.Response, error) {
	getToken()
	client := getRestyClient()
	return client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetBody(chatParameter).
		SetDoNotParseResponse(true).
		Post(HOST + urlModelAppService + "/" + jobId + "/forward/inference")

}

func ClickOnce(url string) {
	client := getRestyClient()
	log.Info("click url=" + url)
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		Get(url)
	if err != nil {
		log.Info("error=" + err.Error())
	} else {
		if res != nil {
			jsonstr, _ := json.Marshal(res)
			log.Info("job resp:", jsonstr)
		}
	}
}

func SendMsgToWebsocket(wsurl string, msg string) {
	//wsurl := strings.Replace(tmpurl, "http", "ws", 1)
	log.Info("wsurl =" + wsurl)
	conn, _, err := websocket.DefaultDialer.Dial(wsurl, nil)
	if err == nil {
		defer conn.Close()
		err1 := conn.WriteJSON(msg)
		//err1 := conn.WriteJSON("{\"msg\":\"send_hash\"}")
		if err1 != nil {
			log.Info("websocket send msg error=" + err1.Error())
		}
	} else {
		log.Info("websocket error=" + err.Error())
	}
}
