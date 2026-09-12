package grampus

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/modules/notification"
	"code.gitea.io/gitea/modules/otel"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"github.com/go-resty/resty/v2"
)

var (
	restyClient *resty.Client
	HOST        string
	TOKEN       string
)

const (
	urlOpenApiV1 = "/openapi/v1/"

	urlGetToken             = urlOpenApiV1 + "token"
	urlTrainJob             = urlOpenApiV1 + "trainjob"
	urlGetResourceSpecs     = urlOpenApiV1 + "resourcespec"
	urlGetResourceSpecPools = urlOpenApiV1 + "resourceSpecPool"
	urlGetAiCenter          = urlOpenApiV1 + "sharescreen/aicenter"
	urlGetImages            = urlOpenApiV1 + "image"
	urlDelImages            = urlOpenApiV1 + "delimage"
	urlNotebookJob          = urlOpenApiV1 + "notebook"
	urlInferenceJob         = urlOpenApiV1 + "inference"

	errorIllegalToken        = 1005
	GrampusImageCommitMade   = 3
	GrampusImageCommitFailed = 4
	ImageExist               = 13003
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
	HOST = setting.Grampus.Host

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

func createNotebookJob(req models.CreateGrampusNotebookRequest) (*models.GrampusNotebookResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GrampusNotebookResponse

	retry := 0

sendjob:
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetBody(req).
		SetResult(&result).
		Post(HOST + urlNotebookJob)

	if err != nil {
		return nil, fmt.Errorf("resty CreateNotebookJob: %s", err)
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

	return &result, nil
}

func createJob(req models.CreateGrampusJobRequest) (*models.CreateGrampusJobResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.CreateGrampusJobResponse

	retry := 0

sendjob:
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
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

func GetNotebookJob(jobID string) (*models.GrampusNotebookResponse, error) {
	if jobID == "" {
		return nil, fmt.Errorf("jobID is emmpty")
	}

	checkSetting()
	client := getRestyClient()
	var result models.GrampusNotebookResponse

	retry := 0

sendjob:
	_, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + urlNotebookJob + "/" + jobID)

	if err != nil {
		return nil, fmt.Errorf("resty GetNotebookJob: %v", err)
	}

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

func GetJob(jobID string) (*models.GetGrampusJobResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetGrampusJobResponse

	retry := 0

sendjob:
	_, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + urlTrainJob + "/" + jobID)
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

func GetResourceSpecs(processorType string) (*models.GetGrampusResourceSpecsResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetGrampusResourceSpecsResult

	retry := 0

sendjob:
	res, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + urlGetResourceSpecs + "?processorType=" + processorType)
	log.Info("%+v", res)
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
	jsonStr, _ := json.Marshal(result)
	log.Info("GetResourceSpecs result=%s", string(jsonStr))
	return &result, nil
}

func GetResourceSpecPools(processorType string) (*models.GetGrampusResourceSpecPoolsResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetGrampusResourceSpecPoolsResult

	retry := 0

sendjob:
	res, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + urlGetResourceSpecPools + "?processorType=" + processorType)
	log.Info("%+v", res)
	if err != nil {
		return nil, fmt.Errorf("resty GetGrampusResourceSpecPools: %v", err)
	}

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("GetGrampusResourceSpecPools failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("GetGrampusResourceSpecPools failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}
	jsonStr, _ := json.Marshal(result)
	log.Info("GetGrampusResourceSpecPools result=%s", string(jsonStr))
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

	if result.ErrorCode != 0 {
		log.Error("GetImages failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("GetImages failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func GetAllBaseImages() (*models.GetGrampusImagesResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetGrampusImagesResult
	retry := 0
sendjob:
	_, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + urlGetImages)

	if err != nil {
		return nil, fmt.Errorf("resty GetAllBaseImages: %v", err)
	}
	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("GetAllBaseImages failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("GetImages failed(%d): %s", result.ErrorCode, result.ErrorMsg)
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
	url := HOST + urlTrainJob + "/" + jobID + "/task/0/replica/0/metrics"
	if len(nodeId) > 0 {
		url = HOST + urlTrainJob + "/" + jobID + "/task/0/replica/0/metrics/node/" + strconv.Itoa(nodeId[0])
	}
	if startTime > 0 {
		var step int64 = 60

		size := int64(math.Ceil(float64(endTime-startTime)/float64(step))) + 1

		url = url + "?startTime=" + strconv.FormatInt(startTime, 10) + "&step=" + strconv.FormatInt(step, 10) + "&size=" + strconv.FormatInt(size, 10)
	}
	res, err := client.R().
		SetAuthToken(TOKEN).
		Get(url)

	if err != nil {
		return result, fmt.Errorf("resty GetTrainJobLog: %v", err)
	}
	if err = json.Unmarshal([]byte(res.String()), &result); err != nil {
		log.Error("GetGrampusMetrics json.Unmarshal failed(%s): %v", res.String(), err.Error())
		return result, fmt.Errorf("json.Unmarshal failed(%s): %v", res.String(), err.Error())
	}
	if res.StatusCode() != http.StatusOK {
		return result, fmt.Errorf("Call GrampusMetrics failed(%d)", res.StatusCode())
	}
	return result, nil
}

func StopJob(jobID string, jobType ...string) (*models.GrampusStopJobResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GrampusStopJobResponse

	retry := 0

	url := urlTrainJob
	if len(jobType) > 0 {
		if jobType[0] == string(models.JobTypeDebug) {
			url = urlNotebookJob
		}
	}

sendjob:
	_, err := client.R().
		//SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetResult(&result).
		Post(HOST + url + "/" + jobID + "/stop")

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
		return &result, fmt.Errorf("GetJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func PostModelMigrate(jobID string) (*models.GrampusModelMigrateInfoResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GrampusModelMigrateInfoResponse

	retry := 0

sendjob:
	res, err := client.R().
		//SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetResult(&result).
		Post(HOST + urlTrainJob + "/" + jobID + "/modelMigrate")

	if err != nil {
		return &result, fmt.Errorf("resty ModelMigrate: %v", err)
	}
	log.Info("call modelMigrate res=%+v", res)
	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("ModelMigrate failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("GetJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}

	return &result, nil
}

func ModelMigrateInfo(jobID string) (*models.GrampusModelMigrateInfoResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GrampusModelMigrateInfoResponse

	retry := 0

sendjob:
	res, err := client.R().
		//SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + urlTrainJob + "/" + jobID + "/modelMigrateInfo")

	if err != nil {
		return &result, fmt.Errorf("resty ModelMigrateInfo: %v", err)
	}
	log.Info("call modelMigrateInfo res=%+v", res)
	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}

	if result.ErrorCode != 0 {
		log.Error("ModelMigrateInfo failed(%d): %s", result.ErrorCode, result.ErrorMsg)
		return &result, fmt.Errorf("GetJob failed(%d): %s", result.ErrorCode, result.ErrorMsg)
	}

	//针对modelarts调试任务且用户没有输出文件的情况，结果回传会失败，需要特殊处理成迁移成功
	if result.Status == int(models.GrampusMigrateResponseFailed) && strings.Contains(result.FailedReason, "bad response status 404 Not Found") {
		result.Status = int(models.GrampusMigrateResponseNoNeedMigrate)
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

func GetResourceAICenter() ([]models.GrampusAiCenter, error) {
	page := 1
	pageSize := 100
	result := make([]models.GrampusAiCenter, 0)
	for {
		res, err := GetAiCenters(page, pageSize)
		if err != nil {
			return nil, err
		}
		result = append(result, res.Infos...)
		if len(res.Infos) < pageSize {
			break
		}
		if page > 1000 {
			//防止接口异常导致无限循环
			break
		}
		page++
	}

	return result, nil
}

func GetResourceQueue() ([]models.GrampusResourceQueue, error) {
	res, err := GetResourceSpecs("")
	if err != nil {
		return nil, err
	}
	queueList := make([]models.GrampusResourceQueue, 0)
	queueMap := make(map[string]string, 0)
	for _, spec := range res.Infos {
		for _, c := range spec.Centers {
			centerId := c.ID
			computeResource := models.ParseComputeResourceFormGrampus(spec.SpecInfo.AccDeviceKind)
			if centerId == "" || computeResource == "" {
				continue
			}
			for _, queue := range c.ResourceSpec {
				queueCode := queue.ID
				accCardType := strings.ToUpper(spec.SpecInfo.AccDeviceModel)
				key := centerId + "_" + computeResource + "_" + accCardType + "_" + queueCode
				if _, has := queueMap[key]; has {
					continue
				}
				var hasInternet = int(models.NoInternet)
				if queue.IsNetAccess == models.GrampusNetAccess {
					hasInternet = int(models.HasInternet)
				}
				var isSupportVisual = false
				if queue.IsSupportVisual == models.GrampusTrue {
					isSupportVisual = true
				}
				var queueType = models.QueueTypePublic
				if queue.PoolType == models.GrampusPoolTypeExclusive {
					queueType = models.QueueTypeExclusive
				}
				queueMap[key] = ""
				queueList = append(queueList, models.GrampusResourceQueue{
					QueueCode:           queueCode,
					QueueName:           queue.Name,
					QueueType:           queueType,
					AiCenterCode:        centerId,
					AiCenterName:        c.Name,
					ComputeResource:     computeResource,
					AccCardType:         accCardType,
					HasInternet:         hasInternet,
					EnableVisualization: isSupportVisual,
				})
			}

		}
	}
	return queueList, nil
}

func GetNewResourceQueue() ([]models.GrampusResourceQueue, error) {
	res, err := GetResourceSpecPools("")
	if err != nil {
		return nil, err
	}
	queueList := make([]models.GrampusResourceQueue, 0)

	for _, pool := range res.Infos {
		log.Info("pool=%+v", pool)
		var queueType = models.QueueTypePublic
		if pool.PoolType == models.GrampusPoolTypeExclusive {
			queueType = models.QueueTypeExclusive
		}
		computeResource := models.GetComputeSourceStandardFormat(pool.ProcessorType)
		if computeResource == "" {
			log.Error("GetResourceQueue: computeResource is empty, pool=%+v", pool)
			continue
		}
		accCardType := strings.ToUpper(pool.CardType)
		var hasInternet = int(models.NoInternet)
		if pool.IsNetAccess == models.GrampusNetAccess {
			hasInternet = int(models.HasInternet)
		}
		var isSupportVisual = false
		if pool.IsSupportVisual == models.GrampusTrue {
			isSupportVisual = true
		}
		tmpQueue := models.GrampusResourceQueue{
			QueueCode:           pool.ID,
			QueueName:           pool.Name,
			QueueType:           queueType,
			AiCenterCode:        pool.CenterId,
			AiCenterName:        pool.CenterName,
			ComputeResource:     computeResource,
			AccCardType:         accCardType,
			HasInternet:         hasInternet,
			EnableVisualization: isSupportVisual,
			CardsTotalNum:       pool.CardNum,
		}
		log.Info("queue=%+v", tmpQueue)
		queueList = append(queueList, tmpQueue)

	}
	return queueList, nil
}

func GetDebugJobEvents(jobID string) (*models.GetGrampusDebugJobEventsResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetGrampusDebugJobEventsResponse

	retry := 0

sendjob:
	_, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + urlNotebookJob + "/" + jobID + "/events")

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

func GetTrainJobEvents(jobID string) (*models.GetGrampusJobEventsResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetGrampusJobEventsResponse

	retry := 0

sendjob:
	_, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + urlTrainJob + "/" + jobID + "/events")
	if err != nil {
		return nil, fmt.Errorf("resty GetTrainJobEvents: %v", err)
	}

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

func RestartNotebookJob(jobID string) (*models.GrampusNotebookRestartResponse, error) {
	checkSetting()
	client := getRestyClient()
	var restartResponse *models.GrampusNotebookRestartResponse
	retry := 0

sendjob:
	res, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&restartResponse).
		Post(HOST + urlNotebookJob + "/" + jobID + "/start")

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

	return restartResponse, nil
}

func CommitImage(task *models.Cloudbrain, computeResource, aiCenterId string, params models.CommitGrampusImageParams, doer *models.User) error {
	imageTag := strings.TrimSpace(params.ImageVersion)
	var err error
	dbImage, err := models.GetImageByTag(imageTag)

	if err != nil && !models.IsErrImageNotExist(err) {
		return fmt.Errorf("resty CommitImage: %v", err)
	}
	if dbImage != nil {
		return models.ErrorImageTagExist{
			Tag: imageTag,
		}

	}
	log.Info(fmt.Sprintf("Grmpus-CommitImage jobId[%s],computeResource[%s],aiCenterId[%s],params[%+v],doer[%+v]", task.JobID, computeResource, aiCenterId, params, doer))

	checkSetting()
	client := getRestyClient()
	var result models.CommitGrampusImageResult

	retry := 0
	newCtx, childCtx, span := otel.StartTrace(&entity.TraceInfo{TaskId: task.ID, SpanName: "CommitImage"}, false)

	defer func() {
		if err != nil {
			otel.RemoveTraceCache(task.ID, otel.GenerateCacheValue(span))
		}
		otel.FinalizeSpan(childCtx, err)
		otel.FinalizeSpan(newCtx, err)
	}()

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("traceparent", otel.GenerateHeader(span)).
		SetAuthToken(TOKEN).
		SetBody(params.CommitImageGrampusParams).
		SetResult(&result).
		Post(HOST + urlNotebookJob + "/" + task.JobID + "/save")

	if err != nil {
		return fmt.Errorf("resty CommitImage: %v", err)
	}

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}
	if res.StatusCode() != http.StatusOK {

		return fmt.Errorf("CommitImage err: %s", res.String())
	}

	if result.ErrorCode == ImageExist {
		return models.ErrorImageTagExist{
			Tag: imageTag,
		}
	}

	if result.Id == "" {
		return fmt.Errorf("CommitImage err: imageid is empty. %s", result.ErrorMsg)
	}

	aiCenterList := make(models.AiCenterImages, 0)
	ai := models.AiCenterImage{
		AiCenterId: aiCenterId,
	}
	aiCenterList = append(aiCenterList, ai)
	trainType, trainNum := models.ConvertTrainType(params.TrainType)

	image := models.Image{
		Type:                   models.NORMAL_TYPE,
		CloudbrainType:         params.CloudBrainType,
		UID:                    params.UID,
		IsPrivate:              params.IsPrivate,
		Tag:                    imageTag,
		Description:            params.Description,
		ImageID:                result.Id,
		Status:                 models.IMAGE_STATUS_COMMIT,
		ApplyStatus:            models.NoneApply,
		Framework:              params.Framework,
		FrameworkVersion:       params.FrameworkVersion,
		CudaVersion:            params.CudaVersion,
		CannVersion:            params.CannVersion,
		PythonVersion:          params.PythonVersion,
		OperationSystem:        params.OperationSystem,
		OperationSystemVersion: params.OperationSystemVersion,
		ThirdPackages:          params.ThirdPackages,
		ComputeResource:        params.ComputeResource,
		GrampusBaseImage:       0,
		AiCenterImages:         aiCenterList,
		DTKVersion:             params.DTKVersion,
		TrainType:              trainType,
		TrainTypeNum:           trainNum,
	}

	err = models.WithTx(func(ctx models.DBContext) error {
		models.UpdateAutoIncrementIndex()
		// if dbImage != nil {
		// 	dbImage.IsPrivate = params.IsPrivate
		// 	dbImage.Description = params.Description
		// 	dbImage.Status = models.IMAGE_STATUS_COMMIT
		// 	dbImage.Framework = params.Framework
		// 	dbImage.FrameworkVersion = params.FrameworkVersion
		// 	dbImage.CudaVersion = params.CudaVersion
		// 	dbImage.PythonVersion = params.PythonVersion
		// 	dbImage.OperationSystem = params.OperationSystem
		// 	dbImage.OperationSystemVersion = params.OperationSystemVersion
		// 	dbImage.ThirdPackages = params.ThirdPackages
		// 	dbImage.ComputeResource = params.ComputeResource
		// 	image = *dbImage
		// 	if err := models.UpdateLocalImage(dbImage); err != nil {
		// 		log.Error("Failed to update image record.", err)
		// 		return fmt.Errorf("CommitImage err: %s", res.String())
		// 	}

		// } else {
		if err := models.CreateLocalImage(&image); err != nil {
			log.Error("Failed to insert image record.", err)
			return fmt.Errorf("CommitImage err: %s", res.String())
		}
		//}
		if err := models.SaveImageTopics(image.ID, params.Topics...); err != nil {
			log.Error("Failed to insert image record.", err)
			return fmt.Errorf("CommitImage err: %s", res.String())
		}
		return nil
	})
	if err == nil {
		notification.NotifyCreateImage(doer, image)
	}
	return err
}

func CommitInferJobImage(task *models.Cloudbrain, computeResource, aiCenterId string, params models.CommitGrampusImageParams, doer *models.User) error {
	imageTag := strings.TrimSpace(params.ImageVersion)
	var err error
	dbImage, err := models.GetImageByTag(imageTag)

	if err != nil && !models.IsErrImageNotExist(err) {
		return fmt.Errorf("resty CommitImage: %v", err)
	}
	if dbImage != nil {
		return models.ErrorImageTagExist{
			Tag: imageTag,
		}

	}

	checkSetting()
	client := getRestyClient()
	var result models.CommitGrampusImageResult

	retry := 0

	newCtx, childCtx, span := otel.StartTrace(&entity.TraceInfo{TaskId: task.ID, SpanName: "CommitImage"}, false)

	defer func() {
		if err != nil {
			otel.RemoveTraceCache(task.ID, otel.GenerateCacheValue(span))
		}
		otel.FinalizeSpan(childCtx, err)
		otel.FinalizeSpan(newCtx, err)
	}()

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("traceparent", otel.GenerateHeader(span)).
		SetAuthToken(TOKEN).
		SetBody(params.CommitImageGrampusParams).
		SetResult(&result).
		Post(HOST + urlInferenceJob + "/" + task.JobID + "/save")

	if err != nil {
		return fmt.Errorf("resty CommitImage: %v", err)
	}

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}
	if res.StatusCode() != http.StatusOK {

		return fmt.Errorf("CommitImage err: %s", res.String())
	}

	if result.ErrorCode == ImageExist {
		return models.ErrorImageTagExist{
			Tag: imageTag,
		}
	}

	if result.Id == "" {
		return fmt.Errorf("CommitImage err: imageid is empty. %s", result.ErrorMsg)
	}

	aiCenterList := make(models.AiCenterImages, 0)
	ai := models.AiCenterImage{
		AiCenterId: aiCenterId,
	}
	aiCenterList = append(aiCenterList, ai)

	image := models.Image{
		Type:                   models.NORMAL_TYPE,
		CloudbrainType:         params.CloudBrainType,
		UID:                    params.UID,
		IsPrivate:              params.IsPrivate,
		Tag:                    imageTag,
		Description:            params.Description,
		ImageID:                result.Id,
		Status:                 models.IMAGE_STATUS_COMMIT,
		ApplyStatus:            models.NoneApply,
		Framework:              params.Framework,
		FrameworkVersion:       params.FrameworkVersion,
		CudaVersion:            params.CudaVersion,
		CannVersion:            params.CannVersion,
		PythonVersion:          params.PythonVersion,
		OperationSystem:        params.OperationSystem,
		OperationSystemVersion: params.OperationSystemVersion,
		ThirdPackages:          params.ThirdPackages,
		ComputeResource:        params.ComputeResource,
		GrampusBaseImage:       0,
		AiCenterImages:         aiCenterList,
		DTKVersion:             params.DTKVersion,
	}

	err = models.WithTx(func(ctx models.DBContext) error {
		models.UpdateAutoIncrementIndex()
		// if dbImage != nil {
		// 	dbImage.IsPrivate = params.IsPrivate
		// 	dbImage.Description = params.Description
		// 	dbImage.Status = models.IMAGE_STATUS_COMMIT
		// 	dbImage.Framework = params.Framework
		// 	dbImage.FrameworkVersion = params.FrameworkVersion
		// 	dbImage.CudaVersion = params.CudaVersion
		// 	dbImage.PythonVersion = params.PythonVersion
		// 	dbImage.OperationSystem = params.OperationSystem
		// 	dbImage.OperationSystemVersion = params.OperationSystemVersion
		// 	dbImage.ThirdPackages = params.ThirdPackages
		// 	dbImage.ComputeResource = params.ComputeResource
		// 	image = *dbImage
		// 	if err := models.UpdateLocalImage(dbImage); err != nil {
		// 		log.Error("Failed to update image record.", err)
		// 		return fmt.Errorf("CommitImage err: %s", res.String())
		// 	}

		// } else {
		if err := models.CreateLocalImage(&image); err != nil {
			log.Error("Failed to insert image record.", err)
			return fmt.Errorf("CommitImage err: %s", res.String())
		}
		//}
		if err := models.SaveImageTopics(image.ID, params.Topics...); err != nil {
			log.Error("Failed to insert image record.", err)
			return fmt.Errorf("CommitImage err: %s", res.String())
		}
		return nil
	})
	if err == nil {
		notification.NotifyCreateImage(doer, image)
	}
	return err
}

func GetImageStatus(image *models.Image) (*models.GrampusImageStatusResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GrampusImageStatusResult

	retry := 0

sendjob:
	res, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Get(HOST + urlGetImages + "/" + image.ImageID)

	if err != nil {
		return nil, fmt.Errorf("resty CommitImage: %v", err)
	}

	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}
	if res.StatusCode() != http.StatusOK {

		return nil, fmt.Errorf("CommitImage err: %s", res.String())
	}

	return &result, nil

}

func DeleteImage(image *models.Image) error {
	checkSetting()
	client := getRestyClient()
	log.Info("start to delete remote image=" + image.ImageID)
	var result models.GrampusResult
	retry := 0
sendjob:
	res, err := client.R().
		SetAuthToken(TOKEN).
		SetResult(&result).
		Delete(HOST + urlDelImages + "/" + image.ImageID)

	if err != nil {
		log.Info("DeleteImage error=" + err.Error())
		return fmt.Errorf("resty DeleteImage: %v", err)
	}
	if result.ErrorCode == errorIllegalToken && retry < 1 {
		retry++
		log.Info("retry get token")
		_ = getToken()
		goto sendjob
	}
	if res.StatusCode() != http.StatusOK {
		log.Info("res status code=" + fmt.Sprint(res.StatusCode()))
		return fmt.Errorf("DeleteImage err: %s", res.String())
	}
	return nil
}
