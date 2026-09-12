package cloudbrain

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/notification"

	"code.gitea.io/gitea/modules/log"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/setting"
	"github.com/go-resty/resty/v2"
)

var (
	restyClient  *resty.Client
	HOST         string
	TOKEN        string
	ImagesUrlMap = map[string]string{Public: "/rest-server/api/v1/image/public/list/", Custom: "/rest-server/api/v1/image/list/"}
)

const (
	JobHasBeenStopped   = "S410"
	Public              = "public"
	Custom              = "custom"
	LogPageSize         = 500
	errInvalidToken     = "S401"
	LogPageTokenExpired = "5m"
	pageSize            = 15
	QueuesDetailUrl     = "/rest-server/api/v2/queuesdetail"
)

func getRestyClient() *resty.Client {
	if restyClient == nil {
		restyClient = resty.New()
	}
	return restyClient
}

func checkSetting() {
	if len(HOST) != 0 && len(TOKEN) != 0 && restyClient != nil {
		return
	}
	_ = loginCloudbrain()
}

func loginCloudbrain() error {
	conf := setting.GetCloudbrainConfig()

	username := conf.Username
	password := conf.Password
	HOST = conf.Host
	var loginResult models.CloudBrainLoginResult

	client := getRestyClient()

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{"username": username, "password": password, "expiration": conf.Expiration}).
		SetResult(&loginResult).
		Post(HOST + "/rest-server/api/v1/token")
	if err != nil {
		return fmt.Errorf("resty loginCloudbrain: %s", err)
	}

	if loginResult.Code != Success {
		return fmt.Errorf("%s: %s", loginResult.Msg, res.String())
	}

	TOKEN = loginResult.Payload["token"].(string)
	return nil
}

func GetQueuesDetail() (*map[string]int, error) {
	checkSetting()
	client := getRestyClient()
	var jobResult models.QueueDetailResult

	var result = make(map[string]int, 0)
	retry := 0
sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetResult(&jobResult).
		Get(HOST + QueuesDetailUrl)

	if err != nil {
		return nil, fmt.Errorf("resty get queues detail failed: %s", err)
	}

	if (res.StatusCode() == http.StatusUnauthorized || jobResult.Code == errInvalidToken) && retry < 1 {
		retry++
		_ = loginCloudbrain()
		goto sendjob
	}

	if jobResult.Code != Success {
		return nil, fmt.Errorf("jobResult err: %s", res.String())
	}

	for k, v := range jobResult.Payload {

		result[k] = v.JobScheduleInfo.Pending

	}
	return &result, nil

}

func CreateJob(jobName string, createJobParams models.CreateJobParams) (*models.CreateJobResult, error) {
	checkSetting()
	client := getRestyClient()
	var jobResult models.CreateJobResult

	retry := 0
	reqPara, _ := json.Marshal(createJobParams)
	log.Warn("job req:", string(reqPara[:]))

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetBody(createJobParams).
		SetResult(&jobResult).
		Post(HOST + "/rest-server/api/v1/jobs/")

	if err != nil {
		log.Error("resty create job: %s", err)
		return nil, models.NetworkError{}
	}

	var response models.CloudBrainResult
	json.Unmarshal(res.Body(), &response)
	if response.Code == errInvalidToken && retry < 1 {
		retry++
		_ = loginCloudbrain()
		goto sendjob
	}

	if jobResult.Code != Success {
		return &jobResult, fmt.Errorf("jobResult err: %s", res.String())
	}
	return &jobResult, nil
}

func GetJob(jobID string) (*models.GetJobResult, error) {
	checkSetting()
	// http://192.168.204.24/rest-server/api/v1/jobs/90e26e500c4b3011ea0a251099a987938b96
	client := getRestyClient()
	var getJobResult models.GetJobResult

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetResult(&getJobResult).
		Get(HOST + "/rest-server/api/v1/jobs/" + jobID)

	if err != nil {
		return nil, fmt.Errorf("resty GetJob: %v", err)
	}

	var response models.CloudBrainResult
	json.Unmarshal(res.Body(), &response)
	if response.Code == errInvalidToken && retry < 1 {
		retry++
		_ = loginCloudbrain()
		goto sendjob
	}

	if getJobResult.Code != Success {
		return &getJobResult, fmt.Errorf("jobResult GetJob err: %s", res.String())
	}

	return &getJobResult, nil
}

func GetJobListByName(jobName string) (*models.GetJobListResult, error) {
	checkSetting()
	// http://192.168.204.24/rest-server/api/v1/jobs/90e26e500c4b3011ea0a251099a987938b96
	client := getRestyClient()
	var getJobResult models.GetJobListResult

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetResult(&getJobResult).
		Get(HOST + "/rest-server/api/v1/jobs?size=20&offset=0&username=" + setting.GetCloudbrainConfig().Username + "&jobName=" + jobName)

	if err != nil {
		return nil, fmt.Errorf("resty GetJob: %v", err)
	}

	var response models.CloudBrainResult
	json.Unmarshal(res.Body(), &response)
	if response.Code == errInvalidToken && retry < 1 {
		retry++
		_ = loginCloudbrain()
		goto sendjob
	}

	if getJobResult.Code != Success {
		return &getJobResult, fmt.Errorf("jobResult GetJob err: %s", res.String())
	}

	return &getJobResult, nil
}

func GetImagesPageable(page int, size int, imageType string, name string) (*models.GetImagesResult, error) {
	checkSetting()
	client := getRestyClient()
	var getImagesResult models.GetImagesResult

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetQueryString(getQueryString(page, size, name)).
		SetResult(&getImagesResult).
		Get(HOST + ImagesUrlMap[imageType])

	if err != nil {
		return nil, fmt.Errorf("resty GetImages: %v", err)
	}

	var response models.CloudBrainResult
	json.Unmarshal(res.Body(), &response)
	if response.Code == errInvalidToken && retry < 1 {
		retry++
		_ = loginCloudbrain()
		goto sendjob
	}

	if getImagesResult.Code != Success {
		return &getImagesResult, fmt.Errorf("getImagesResult err: %s", res.String())
	}

	getImagesResult.Payload.TotalPages = getTotalPages(getImagesResult, size)
	return &getImagesResult, nil
}

func getTotalPages(getImagesResult models.GetImagesResult, size int) int {
	totalCount := getImagesResult.Payload.Count
	var totalPages int
	if totalCount%size != 0 {
		totalPages = totalCount/size + 1
	} else {
		totalPages = totalCount / size
	}
	return totalPages
}

func getQueryString(page int, size int, name string) string {
	if strings.TrimSpace(name) == "" {
		return fmt.Sprintf("pageIndex=%d&pageSize=%d", page, size)
	}
	return fmt.Sprintf("pageIndex=%d&pageSize=%d&name=%s", page, size, name)
}

func CommitImage(jobID string, params models.CommitImageParams, doer *models.User) error {
	imageTag := strings.TrimSpace(params.ImageTag)

	dbImageInC2net, err := models.GetImageByTagAndCloudbrainType(imageTag, models.TypeC2Net)
	if err != nil && !models.IsErrImageNotExist(err) {
		return fmt.Errorf("resty CommitImage: %v", err)
	}
	if dbImageInC2net != nil {
		return models.ErrorImageTagExist{
			Tag: imageTag,
		}
	}

	dbImage, err := models.GetImageByTagAndCloudbrainType(imageTag, models.TypeCloudBrainOne)

	if err != nil && !models.IsErrImageNotExist(err) {
		return fmt.Errorf("resty CommitImage: %v", err)
	}
	var createTime time.Time
	var isSetCreatedUnix = false
	if dbImage != nil {
		if dbImage.UID != params.UID {
			return models.ErrorImageTagExist{
				Tag: imageTag,
			}
		} else {
			if dbImage.Status == models.IMAGE_STATUS_COMMIT {
				return models.ErrorImageCommitting{
					Tag: imageTag,
				}

			} else { //覆盖提交

				result, err := GetImagesPageable(1, pageSize, Custom, "")
				if err == nil && result.Code == "S000" {
					for _, v := range result.Payload.ImageInfo {
						if v.Place == dbImage.Place {
							isSetCreatedUnix = true
							createTime, _ = time.Parse(time.RFC3339, v.Createtime)
							break
						}
					}
				}

			}
		}
	}

	checkSetting()
	client := getRestyClient()
	var result models.CommitImageResult

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetBody(params.CommitImageCloudBrainParams).
		SetResult(&result).
		Post(HOST + "/rest-server/api/v1/jobs/" + jobID + "/commitImage")

	if err != nil {
		return fmt.Errorf("resty CommitImage: %v", err)
	}

	var response models.CloudBrainResult
	json.Unmarshal(res.Body(), &response)
	if response.Code == errInvalidToken && retry < 1 {
		retry++
		_ = loginCloudbrain()
		goto sendjob
	}

	if result.Code != Success {
		return fmt.Errorf("CommitImage err: %s", res.String())
	}

	image := models.Image{
		Type:                   models.NORMAL_TYPE,
		CloudbrainType:         params.CloudBrainType,
		UID:                    params.UID,
		IsPrivate:              params.IsPrivate,
		Tag:                    imageTag,
		Description:            params.ImageDescription,
		Place:                  setting.Cloudbrain.ImageURLPrefix + imageTag,
		Status:                 models.IMAGE_STATUS_COMMIT,
		ApplyStatus:            models.NoneApply,
		Framework:              params.Framework,
		FrameworkVersion:       params.FrameworkVersion,
		CudaVersion:            params.CudaVersion,
		PythonVersion:          params.PythonVersion,
		OperationSystem:        params.OperationSystem,
		OperationSystemVersion: params.OperationSystemVersion,
		ThirdPackages:          params.ThirdPackages,
	}

	err = models.WithTx(func(ctx models.DBContext) error {
		models.UpdateAutoIncrementIndex()
		if dbImage != nil {
			dbImage.IsPrivate = params.IsPrivate
			dbImage.Description = params.ImageDescription
			dbImage.Status = models.IMAGE_STATUS_COMMIT
			dbImage.Framework = params.Framework
			dbImage.FrameworkVersion = params.FrameworkVersion
			dbImage.CudaVersion = params.CudaVersion
			dbImage.PythonVersion = params.PythonVersion
			dbImage.OperationSystem = params.OperationSystem
			dbImage.OperationSystemVersion = params.OperationSystemVersion
			dbImage.ThirdPackages = params.ThirdPackages
			image = *dbImage
			if err := models.UpdateLocalImage(dbImage); err != nil {
				log.Error("Failed to update image record.", err)
				return fmt.Errorf("CommitImage err: %s", res.String())
			}

		} else {
			if err := models.CreateLocalImage(&image); err != nil {
				log.Error("Failed to insert image record.", err)
				return fmt.Errorf("CommitImage err: %s", res.String())
			}
		}
		if err := models.SaveImageTopics(image.ID, params.Topics...); err != nil {
			log.Error("Failed to insert image record.", err)
			return fmt.Errorf("CommitImage err: %s", res.String())
		}
		return nil
	})
	if err == nil {
		go updateImageStatus(image, isSetCreatedUnix, createTime)
		notification.NotifyCreateImage(doer, image)
	}
	return err
}

func CommitAdminImage(params models.CommitImageParams, doer *models.User) error {
	imageTag := strings.TrimSpace(params.ImageTag)
	exist, err := models.IsImageExist(imageTag)

	if err != nil {
		return fmt.Errorf("resty CommitImage: %v", err)
	}
	if exist {
		return models.ErrorImageTagExist{
			Tag: imageTag,
		}
	}

	trainType, trainNum := models.ConvertTrainType(params.TrainType)

	image := models.Image{
		CloudbrainType:         params.CloudBrainType,
		UID:                    params.UID,
		IsPrivate:              params.IsPrivate,
		Tag:                    imageTag,
		Description:            params.ImageDescription,
		Place:                  params.Place,
		Status:                 models.IMAGE_STATUS_SUCCESS,
		Type:                   params.Type,
		ApplyStatus:            models.NoneApply,
		Framework:              params.Framework,
		FrameworkVersion:       params.FrameworkVersion,
		CudaVersion:            params.CudaVersion,
		PythonVersion:          params.PythonVersion,
		DTKVersion:             params.DTKVersion,
		TrainType:              trainType,
		TrainTypeNum:           trainNum,
		OperationSystem:        params.OperationSystem,
		OperationSystemVersion: params.OperationSystemVersion,
		ThirdPackages:          params.ThirdPackages,
		ComputeResource:        params.ComputeResource,
	}

	err = models.WithTx(func(ctx models.DBContext) error {

		if err := models.CreateLocalImage(&image); err != nil {
			log.Error("Failed to insert image record.", err)
			return fmt.Errorf("resty CommitImage: %v", err)
		}

		if err := models.SaveImageTopics(image.ID, params.Topics...); err != nil {
			log.Error("Failed to insert image record.", err)
			return fmt.Errorf("resty CommitImage: %v", err)
		}
		return nil
	})
	if err == nil {
		notification.NotifyCreateImage(doer, image)
	}
	return err
}

func updateImageStatus(image models.Image, isSetCreatedUnix bool, createTime time.Time) {
	attemps := 60
	commitSuccess := false

	for i := 0; i < attemps; i++ {
		time.Sleep(20 * time.Second)
		log.Info("the " + strconv.Itoa(i) + " times query cloudbrain images.Imagetag：" + image.Tag + "isSetCreate:" + strconv.FormatBool(isSetCreatedUnix))
		result, err := GetImagesPageable(1, pageSize, Custom, "")
		if err == nil && result.Code == "S000" {

			log.Info("images count:" + strconv.Itoa(result.Payload.Count))
			for _, v := range result.Payload.ImageInfo {
				// TODO 这里是否要删除place的判断
				if v.Place == image.Place && (!isSetCreatedUnix || (isSetCreatedUnix && createTimeUpdated(v, createTime))) {
					image.Status = models.IMAGE_STATUS_SUCCESS
					models.UpdateLocalImageStatus(&image)
					commitSuccess = true
					break
				}

			}

		}

		if commitSuccess {
			break
		}

	}
	if !commitSuccess {
		image.Status = models.IMAGE_STATUS_Failed
		models.UpdateLocalImageStatus(&image)
	}

}

func createTimeUpdated(v *models.ImageInfo, createTime time.Time) bool {
	newTime, err := time.Parse(time.RFC3339, v.Createtime)
	if err != nil {
		return false
	}
	return newTime.After(createTime)
}

func StopJob(jobID string) error {
	checkSetting()
	client := getRestyClient()
	var result models.CloudBrainResult

	retry := 0

sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetResult(&result).
		Delete(HOST + "/rest-server/api/v1/jobs/" + jobID)

	if err != nil {
		return fmt.Errorf("resty StopJob: %v", err)
	}

	var response models.CloudBrainResult
	json.Unmarshal(res.Body(), &response)
	if response.Code == errInvalidToken && retry < 1 {
		retry++
		_ = loginCloudbrain()
		goto sendjob
	}

	if result.Code != Success {
		if result.Code == JobHasBeenStopped {
			log.Info("StopJob(%s) failed:%s", jobID, result.Msg)
		} else {
			return fmt.Errorf("StopJob err: %s", res.String())
		}
	}

	return nil
}

func GetJobLog(jobID string) (*models.GetJobLogResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetJobLogResult
	req := models.GetJobLogParams{
		Size: strconv.Itoa(LogPageSize),
		Sort: "log.offset",
		QueryInfo: models.QueryInfo{
			MatchInfo: models.MatchInfo{
				PodName: jobID + "-task1-0",
			},
		},
	}

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetBody(req).
		SetResult(&result).
		Post(HOST + "es/_search?_source=message&scroll=" + LogPageTokenExpired)

	if err != nil {
		log.Error("GetJobLog failed: %v", err)
		return &result, fmt.Errorf("resty GetJobLog: %v, %s", err, res.String())
	}

	if !strings.Contains(res.Status(), strconv.Itoa(http.StatusOK)) {
		log.Error("res.Status(): %s, response: %s", res.Status(), res.String())
		return &result, errors.New(res.String())
	}

	return &result, nil
}

func GetJobAllLog(scrollID string) (*models.GetJobLogResult, error) {
	checkSetting()
	client := getRestyClient()
	var result models.GetJobLogResult
	req := models.GetAllJobLogParams{
		Scroll:   LogPageTokenExpired,
		ScrollID: scrollID,
	}

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetBody(req).
		SetResult(&result).
		Post(HOST + "es/_search/scroll")

	if err != nil {
		log.Error("GetJobAllLog failed: %v", err)
		return &result, fmt.Errorf("resty GetJobAllLog: %v, %s", err, res.String())
	}

	if !strings.Contains(res.Status(), strconv.Itoa(http.StatusOK)) {
		log.Error("res.Status(): %s, response: %s", res.Status(), res.String())
		return &result, errors.New(res.String())
	}

	return &result, nil
}

func DeleteJobLogToken(scrollID string) error {
	checkSetting()
	client := getRestyClient()
	var result models.DeleteJobLogTokenResult
	req := models.DeleteJobLogTokenParams{
		ScrollID: scrollID,
	}

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(TOKEN).
		SetBody(req).
		SetResult(&result).
		Delete(HOST + "es/_search/scroll")

	if err != nil {
		log.Error("DeleteJobLogToken failed: %v", err)
		return fmt.Errorf("resty DeleteJobLogToken: %v, %s", err, res.String())
	}

	if !strings.Contains(res.Status(), strconv.Itoa(http.StatusOK)) {
		log.Error("res.Status(): %s, response: %s", res.Status(), res.String())
		return errors.New(res.String())
	}

	if !result.Succeeded {
		log.Error("DeleteJobLogToken failed")
		return errors.New("DeleteJobLogToken failed")
	}

	return nil
}
