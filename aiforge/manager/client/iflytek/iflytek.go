package iflytek

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"net/url"
	"strings"
	"time"
)

var (
	restyClient         *resty.Client
	DATA_SERVICE_HOST   string
	TRAIN_SERVICE_HOST  string
	TOKEN               string
	ONLINE_SERVICE_HOST string
	MODEL_SERVICE_HOST  string
)

func checkSetting() {
	DATA_SERVICE_HOST = setting.IFLYTekConfig.DataServiceHost
	TRAIN_SERVICE_HOST = setting.IFLYTekConfig.TrainServiceHost
	TOKEN = setting.IFLYTekConfig.TOKEN
	ONLINE_SERVICE_HOST = setting.IFLYTekConfig.OnlineServiceHost
	MODEL_SERVICE_HOST = setting.IFLYTekConfig.ModelServiceHost
}

const (
	APIPrefix = "/api/v1"

	UploadFileUrl          = APIPrefix + "/file/upload"
	FlintDatasetGroupsUrl  = APIPrefix + "/flintDatasetGroups"
	FlintDatasetVersionUrl = APIPrefix + "/flintDataset"
	DeleteDatasetUrl       = APIPrefix + "/flintDatasetGroups/%s"

	FlintTrainTaskUrl         = APIPrefix + "/flintTrainTasks"
	FlintTrainTaskInfoUrl     = APIPrefix + "/flintTrainTaskVersions/unPage"
	StopFlintTrainTaskUrl     = APIPrefix + "/flintTrainTaskVersions/stop/%d"
	FlintTrainTaskOperateUrl  = APIPrefix + "/flintTrainTasks/%d"
	FlintTrainTaskLogUrl      = APIPrefix + "/flintTrainTaskVersions/getLogs/%d"
	FlintTrainTaskMetaDataUrl = APIPrefix + "/metadata/train-params/SFT"
	FlintTrainTaskProfileUrl  = APIPrefix + "/metadata/train-profile/SFT"

	PublishModelUrl     = APIPrefix + "/flintModelRepos"
	QueryModelDetailUrl = APIPrefix + "/flintModelRepos/%s"

	OnlineServicePrefix          = "/api/v1/aiservice"
	CreateOnlineServiceUrl       = OnlineServicePrefix + "/deploy"
	QueryOnlineServiceStatusUrl  = OnlineServicePrefix + "/status"
	UpdateOnlineServiceStatusUrl = OnlineServicePrefix + "/update"
	DeleteOnlineServiceStatusUrl = OnlineServicePrefix + "/deploy"
)

const (
	CODE_DUPLICATE_DATASET_NAME = 3005
)

func getRestyClient() *resty.Client {
	if restyClient == nil {
		restyClient = resty.New()
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return restyClient
}

// FlintDatasetGroup 用于定义请求的 JSON 数据结构
type FlintDatasetGroup struct {
	Description string `json:"description"`
	LabelId     string `json:"labelId"`
	Name        string `json:"name"`
}

type FlintDatasetGroupResponse struct {
	IFLYTekCommonResponse
	ID          string `json:"id"`
	TenantID    string `json:"tenantId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	LabelID     string `json:"labelId"`
	LabelName   string `json:"labelName"`
	Deleted     int    `json:"deleted"`
}

// FlintDatasetVersionReq 用于定义请求的 JSON 数据结构
type FlintDatasetVersionReq struct {
	Version        string `json:"version"`
	Description    string `json:"description"`
	Source         int    `json:"source"`
	DataType       int    `json:"dataType"`
	OriDataPath    string `json:"oriDataPath"`
	MarkStatus     int    `json:"markStatus"`
	PremarkStatus  int    `json:"premarkStatus"`
	MarkToolId     int    `json:"markToolId"`
	Prompt         string `json:"prompt"`
	Users          []User `json:"users"`
	DatasetGroupId string `json:"datasetGroupId"`
}

type FlintDatasetVersionResponse struct {
	IFLYTekCommonResponse
	ID             string `json:"id"`
	TenantID       string `json:"tenantId"`
	Name           string `json:"name"`
	Version        string `json:"version"`
	Source         int    `json:"source"`
	SourceCount    int    `json:"sourceCount"` // 使用指针来处理可能的null值
	ImportStatus   int    `json:"importStatus"`
	MarkStatus     int    `json:"markStatus"`
	PublishStatus  int    `json:"publishStatus"` // 使用指针来处理可能的null值
	DatasetGroupId string `json:"datasetGroupId"`
	LabelId        string `json:"labelId"`
	LabelName      string `json:"labelName"`
	Count          int    `json:"count"`       // 使用指针来处理可能的null值
	FailCount      int    `json:"failCount"`   // 使用指针来处理可能的null值
	ImportCount    int    `json:"importCount"` // 使用指针来处理可能的null值
	Size           int64  `json:"size"`        // 使用指针来处理可能的null值
	DataType       int    `json:"dataType"`
	QuestionType   string `json:"questionType"` // 使用指针来处理可能的null值
	MarkToolId     int    `json:"markToolId"`
	PremarkStatus  int    `json:"premarkStatus"`
	Prompt         string `json:"prompt"`
	DataPath       string `json:"dataPath"` // 使用指针来处理可能的null值
	OriDataPath    string `json:"oriDataPath"`
	Description    string `json:"description"`
	MarkTaskId     string `json:"markTaskId"` // 使用指针来处理可能的null值
	IsMarked       int    `json:"isMarked"`
	Deleted        int    `json:"deleted"`
}

type FlintTrainTaskReq struct {
	Description string      `json:"description"`
	Name        string      `json:"name"`
	TrainType   int         `json:"trainType"`
	InitVersion InitVersion `json:"initVersion"`
}

type InitVersion struct {
	IncrementalTrain bool           `json:"incrementalTrain"`
	AssetQueue       string         `json:"assetQueue"`
	TrainDataSets    []TrainDataSet `json:"trainDataSets"`
	BaseModelCode    string         `json:"baseModelCode"`
	MaxRunTime       int            `json:"maxRunTime"`
	TestDataSets     []TrainDataSet `json:"testDataSets"`
	RunNow           bool           `json:"runNow"`
	ExtendInfo       string         `json:"extendInfo"`
}

type AssetQueue struct {
	Name             string     `json:"name"`
	Remark           string     `json:"remark"`
	PoolId           string     `json:"poolId"`
	PoolName         string     `json:"poolName"`
	PoolType         string     `json:"poolType"`
	PoolStatus       string     `json:"poolStatus"`
	PoolBizStatus    int        `json:"poolBizStatus"`
	CreatedName      string     `json:"createdName"`
	Users            []User     `json:"users"`
	SpecUsed         []SpecUsed `json:"specUsed"`
	Labels           Labels     `json:"labels"`
	TaskCount        int        `json:"taskCount"`
	ID               string     `json:"id"`
	CreatedBy        string     `json:"createdBy"`
	CreatedDate      time.Time  `json:"createdDate"`
	LastModifiedBy   string     `json:"lastModifiedBy"`
	LastModifiedDate time.Time  `json:"lastModifiedDate"`
	TenantId         string     `json:"tenantId"`
	TenantName       string     `json:"tenantName"`
	RelationId       string     `json:"relationId"`
	RelationName     string     `json:"relationName"`
	Deleted          int        `json:"deleted"`
	QueueStatus      int        `json:"queueStatus"`
	ConfigType       int        `json:"configType"`
}

type User struct {
	ID          string `json:"id"`
	QueueId     string `json:"queueId"`
	UserId      string `json:"userId"`
	Account     string `json:"account"`
	UserName    string `json:"userName"`
	Status      string `json:"status"`
	CreatedName string `json:"createdName"`
}

type SpecUsed struct {
	Uuid     string `json:"uuid"`
	Used     int    `json:"used"`
	Spec     Spec   `json:"spec"`
	Quantity int    `json:"quantity"`
}

type Spec struct {
	Cpu             int             `json:"cpu"`
	Memory          int             `json:"memory"`
	AcceleratorCard AcceleratorCard `json:"acceleratorCard"`
}

type AcceleratorCard struct {
	Manufacturer string `json:"manufacturer"`
	Type         string `json:"type"`
	Quantity     int    `json:"quantity"`
}

type Labels struct {
	TenantId string `json:"tenantId"`
	CreateBy string `json:"createBy"`
}

type TrainDataSet struct {
	ID            string `json:"id"`
	TenantId      string `json:"tenantId"`
	Name          string `json:"name"`
	Source        int    `json:"source"`
	ImportStatus  int    `json:"importStatus"`
	MarkStatus    int    `json:"markStatus"`
	LabelName     string `json:"labelName"`
	LabelId       string `json:"labelId"`
	DataType      int    `json:"dataType"`
	MarkToolId    int    `json:"markToolId"`
	IsMarked      int    `json:"isMarked"`
	PremarkStatus int    `json:"premarkStatus"`
	Prompt        string `json:"prompt"`
	Deleted       int    `json:"deleted"`
	Version       string `json:"version"`
	OriDataPath   string `json:"oriDataPath"`
	DataPath      string `json:"dataPath"`
}

type CreateFlintTrainTaskResponse struct {
	IFLYTekCommonResponse
	BaseModelName       string `json:"baseModelName"`
	CreatedDate         string `json:"createdDate"`
	Deleted             int    `json:"deleted"`
	Description         string `json:"description"`
	ID                  int64  `json:"id"`
	LatestVersionId     int    `json:"latestVersionId"`
	LatestVersionStatus string `json:"latestVersionStatus"`
	ModelId             string `json:"modelId"`
	Name                string `json:"name"`
	TenantId            string `json:"tenantId"`
	TrainType           int    `json:"trainType"`
	VersionCount        int    `json:"versionCount"`
}

type FlintTrainTaskInfo struct {
	AssetQueueID     string      `json:"assetQueueId"`
	CreatedDate      string      `json:"createdDate"`
	Deleted          int         `json:"deleted"`
	Duration         string      `json:"duration"`
	ID               int64       `json:"id"`
	IncrementalTrain bool        `json:"incrementalTrain"`
	MaxRunTime       int         `json:"maxRunTime"`
	Name             string      `json:"name"`
	RunNow           bool        `json:"runNow"`
	TaskId           int64       `json:"taskId"`
	TenantId         string      `json:"tenantId"`
	TrainEndTime     interface{} `json:"trainEndTime"`
	TrainStartTime   interface{} `json:"trainStartTime"`
	TrainStatus      int         `json:"trainStatus"`
	ModelPath        string      `json:"modelPath"`
}

type IFLYTekCommonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type IFLYTekServiceResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    bool   `json:"data"`
}

type ContentItem struct {
	ID          string `json:"id"`
	TenantId    string `json:"tenantId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	LabelID     string `json:"labelId"`
	LabelName   string `json:"labelName"`
}

type QueryDatasetResponse struct {
	Content          []ContentItem `json:"content"`
	Last             bool          `json:"last"`
	TotalElements    int           `json:"totalElements"`
	TotalPages       int           `json:"totalPages"`
	First            bool          `json:"first"`
	NumberOfElements int           `json:"numberOfElements"`
	Size             int           `json:"size"`
	Number           int           `json:"number"`
	Empty            bool          `json:"empty"`
}

type GetMetaDataOpts struct {
	BaseModel       string
	TrainingMethod  string
	SystemStructure string
	AcceleratorCard string
}

type OriginParam struct {
	NodeTag          string   `json:"nodeTag"`
	Tag              string   `json:"tag"`
	CnName           string   `json:"cnName"`
	ParamType        string   `json:"paramType"`
	Description      string   `json:"description"`
	Default          string   `json:"default"`
	Suggest          []string `json:"suggest"`
	Display          bool     `json:"display"`
	Form             string   `json:"form"`
	Require          bool     `json:"require"`
	ReadOnly         bool     `json:"readOnly"`
	Type             string   `json:"type"`
	Env              string   `json:"env"`
	Min              string   `json:"min"`
	Max              string   `json:"max"`
	DisplayCondition string   `json:"displayCondition"`
}

type PublishAsBetterOption struct {
	IsPublishAsBetterOption bool   `json:"IsPublishAsBetterOption"`
	ParamLevel              string `json:"ParamLevel"`
}

type NodeParam struct {
	OriginParam           OriginParam           `json:"originParam"`
	ParamName             string                `json:"paramName"`
	PublishAsBetterOption PublishAsBetterOption `json:"publishAsBetterOption"`
	Value                 interface{}           `json:"value"` // 使用 interface{} 以支持不同类型的值
}

type ParamsMap map[string][]NodeParam

type TemplateData struct {
	Params     ParamsMap  `json:"params"`
	TemplateId string     `json:"templateId"`
	Profile    []Category `json:"profile"`
}

type Condition struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Status      string `json:"status"`
}

type Category struct {
	Condition []Condition `json:"condition"`
	Name      string      `json:"name"`
	QueryName string      `json:"queryName"`
	Value     string      `json:"value"`
}
type GetLogResponse struct {
	Data LogData `json:"data"`
}

type LogData struct {
	Nodes []Node `json:"Nodes"`
	Sid   string `json:"Sid"`
}

type Node struct {
	Log    []LogEntry `json:"Log"` // 使用指针，因为Log字段可以是null
	NodeId string     `json:"NodeId"`
}

type LogEntry struct {
	Log     string `json:"Log"`
	PodName string `json:"PodName"`
}

type CreateOnlineServiceReq struct {
	ServiceId       string `json:"service_id"`
	PretrainedModel string `json:"pretrained_model"`
	LoraPath        string `json:"lora_path"`
}

// UploadFile 方法用于上传文件
func UploadFile(fileStream io.Reader, fileName string) (string, error) {
	checkSetting()
	client := getRestyClient()
	log.Info("UploadFile fileName:%s", fileName)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("SKYBOX_TOKEN_USER_KEY", TOKEN).
		SetFormData(map[string]string{
			"dataType": "1",
		}).
		SetFileReader("file", fileName, fileStream).
		Post(DATA_SERVICE_HOST + UploadFileUrl)

	if err != nil {
		log.Error("resty UploadFile: %v", err)
		return "", err
	}

	log.Info("UploadFile success.resp= %+v", resp.String())
	return resp.String(), nil
}

//AddFlintDatasetGroups 创建数据集
func AddFlintDatasetGroups(req FlintDatasetGroup) (*FlintDatasetGroupResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result FlintDatasetGroupResponse
	log.Info("AddFlintDatasetGroups REQ:%+v", req)
sendjob:
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("SKYBOX_TOKEN_USER_KEY", TOKEN).
		SetBody(req).
		SetResult(&result).
		Post(DATA_SERVICE_HOST + FlintDatasetGroupsUrl)

	if err != nil {
		log.Error("resty AddFlintDatasetGroups: %v", err)
		return nil, err
	}
	if res.IsError() {
		log.Error("AddFlintDatasetGroups error.req= %+v res=%+v", req, res)
		commonRes := IFLYTekCommonResponse{}
		err = json.Unmarshal([]byte(res.String()), &commonRes)
		if err != nil {
			return nil, errors.New(res.String())
		}
		if commonRes.Code == CODE_DUPLICATE_DATASET_NAME {
			err = DeleteDatasetByName(req.LabelId, req.Name)
			if err != nil {
				log.Error("AddFlintDatasetGroups failed.DeleteDatasetByName err.req= %+v err=%v", req, err)
				return nil, err
			}
			goto sendjob
		}
	}
	log.Info("AddFlintDatasetGroups success.req= %+v res=%+v", req, res)
	return &result, nil

}

//AddFlintDatasetVersion 创建数据集版本
func AddFlintDatasetVersion(req FlintDatasetVersionReq) (*TrainDataSet, error) {
	checkSetting()
	client := getRestyClient()
	var result TrainDataSet
	log.Info("AddFlintDatasetVersion REQ:%+v", req)

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("SKYBOX_TOKEN_USER_KEY", TOKEN).
		SetBody(req).
		SetResult(&result).
		Post(DATA_SERVICE_HOST + FlintDatasetVersionUrl)

	if err != nil {
		log.Error("resty AddFlintDatasetVersion: %v", err)
		return nil, err
	}
	if res.IsError() {
		log.Error("AddFlintDatasetVersion error.req= %+v res=%+v", req, res)
		return nil, errors.New(res.String())
	}
	log.Info("AddFlintDatasetVersion success. req= %+v res=%+v", req, res)
	result.DataPath = result.OriDataPath
	return &result, nil
}

//CreateFlintTrainTasks 创建SFT任务
func CreateFlintTrainTask(req FlintTrainTaskReq) (*CreateFlintTrainTaskResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result CreateFlintTrainTaskResponse
	log.Info("CreateFlintTrainTask REQ:%+v", req)

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("SKYBOX_TOKEN_USER_KEY", TOKEN).
		SetBody(req).
		SetResult(&result).
		Post(TRAIN_SERVICE_HOST + FlintTrainTaskUrl)

	if err != nil {
		log.Error("resty CreateFlintTrainTask: %v", err)
		return nil, err
	}
	if res.IsError() {
		log.Error("CreateFlintTrainTask error.req= %+v res=%+v", req, res)
		return nil, errors.New(res.String())
	}
	log.Info("CreateFlintTrainTask success.res=%+v", res)
	return &result, nil
}

//GetFlintTrainTaskMetaData
func GetFlintTrainTaskMetaData(opts GetMetaDataOpts) (*TemplateData, error) {
	checkSetting()
	client := getRestyClient()
	var result TemplateData
	log.Info("GetFlintTrainTaskMetaData opts:%+v", opts)

	res, err := client.R().
		SetHeader("SKYBOX_TOKEN_USER_KEY", TOKEN).
		SetResult(&result).
		Get(TRAIN_SERVICE_HOST + FlintTrainTaskMetaDataUrl + "?baseModel=" + opts.BaseModel + "&trainingMethod=" + opts.TrainingMethod + "&systemStructure=" + opts.SystemStructure + "&acceleratorCard=" + url.QueryEscape(opts.AcceleratorCard))

	if err != nil {
		log.Error("resty GetFlintTrainTaskMetaData: %v", err)
		return nil, err
	}
	if res.IsError() {
		log.Error("GetFlintTrainTaskMetaData error.opts:%+v res=%+v", opts, res)
		return nil, errors.New(res.String())
	}

	log.Info("GetFlintTrainTaskMetaData success.opts:%+v res=%+v", opts, res)
	return &result, nil
}

//GetFlintTrainTaskProfile
func GetFlintTrainTaskProfile() ([]Category, error) {
	checkSetting()
	client := getRestyClient()
	var result []Category
	log.Info("Start to GetFlintTrainTaskProfile")

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("SKYBOX_TOKEN_USER_KEY", TOKEN).
		SetResult(&result).
		Get(TRAIN_SERVICE_HOST + FlintTrainTaskProfileUrl)

	if err != nil {
		log.Error("resty GetFlintTrainTaskProfile: %v", err)
		return nil, err
	}
	if res.IsError() {
		log.Error("GetFlintTrainTaskProfile error. err=%v", err)
		return nil, errors.New(res.String())
	}

	log.Info("GetFlintTrainTaskProfile success.res=%+v", res)
	return result, nil
}

//QueryFlintTrainTask 查询SFT任务
func QueryFlintTrainTask(taskId int64) (*FlintTrainTaskInfo, error) {
	checkSetting()
	client := getRestyClient()
	var resultArray []FlintTrainTaskInfo
	var result *FlintTrainTaskInfo
	log.Info("QueryFlintTrainTask taskId:%s", taskId)

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("SKYBOX_TOKEN_USER_KEY", TOKEN).
		SetResult(&resultArray).
		Get(TRAIN_SERVICE_HOST + FlintTrainTaskInfoUrl + "?taskId.equals=" + fmt.Sprint(taskId))

	if err != nil {
		log.Error("resty QueryFlintTrainTask: %v", err)
		return nil, err
	}
	if res.IsError() {
		log.Error("QueryFlintTrainTask error.req=%s res=%+v", taskId, res)
		return nil, errors.New(res.String())
	}
	if len(resultArray) == 0 {
		result = &FlintTrainTaskInfo{
			TaskId:      taskId,
			TrainStatus: 6,
		}
	}
	if len(resultArray) > 0 {
		result = &resultArray[0]
	}

	log.Info("QueryFlintTrainTask success.req=%s res=%+v", taskId, res)
	return result, nil
}

//StopFlintTrainTask 停止SFT任务
func StopFlintTrainTask(taskId int64) (*IFLYTekCommonResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result IFLYTekCommonResponse
	log.Info("StopFlintTrainTask taskId:%s", taskId)

	res, err := client.R().
		SetHeader("SKYBOX_TOKEN_USER_KEY", TOKEN).
		SetResult(&result).
		Post(TRAIN_SERVICE_HOST + fmt.Sprintf(StopFlintTrainTaskUrl, taskId))

	if err != nil {
		log.Error("resty StopFlintTrainTask: %v", err)
		return nil, err
	}
	if res.IsError() {
		log.Error("StopFlintTrainTask error.req= %s res=%+v", taskId, res)
		return nil, errors.New(res.String())
	}
	log.Info("StopFlintTrainTask success.req= %s res=%+v", taskId, res)
	return &result, nil
}

//DeleteFlintTrainTask 删除SFT任务
func DeleteFlintTrainTask(taskId int64) (*IFLYTekCommonResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result IFLYTekCommonResponse
	log.Info("DeleteFlintTrainTask taskId:%s", taskId)

	res, err := client.R().
		SetHeader("SKYBOX_TOKEN_USER_KEY", TOKEN).
		SetResult(&result).
		Delete(TRAIN_SERVICE_HOST + fmt.Sprintf(FlintTrainTaskOperateUrl, taskId))

	if err != nil {
		log.Error("resty DeleteFlintTrainTask: %v", err)
		return nil, err
	}
	if res.IsError() {
		log.Error("DeleteFlintTrainTask error.req= %s res=%+v", taskId, res)
		return nil, errors.New(res.String())
	}
	log.Info("DeleteFlintTrainTask success.req= %s res=%+v", taskId, res)
	return &result, nil
}

//GetFlintTrainTaskLog 获取SFT任务日志
func GetFlintTrainTaskLog(taskId int64) (*GetLogResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result GetLogResponse
	log.Info("GetFlintTrainTaskLog taskId:%s", taskId)

	res, err := client.R().
		SetHeader("SKYBOX_TOKEN_USER_KEY", TOKEN).
		SetResult(&result).
		Get(TRAIN_SERVICE_HOST + fmt.Sprintf(FlintTrainTaskLogUrl, taskId))

	if err != nil {
		log.Error("resty GetFlintTrainTaskLog: %v", err)
		return nil, err
	}
	if res.IsError() {
		log.Error("GetFlintTrainTaskLog error.req= %s res=%+v", taskId, res)
		return nil, errors.New(res.String())
	}
	log.Info("GetFlintTrainTaskLog success.req= %s res=%+v", taskId, res)
	return &result, nil
}

//DeleteDataset 删除数据集
func DeleteDataset(datasetId string) (*TrainDataSet, error) {
	checkSetting()
	client := getRestyClient()
	var result TrainDataSet
	log.Info("DeleteDataset datasetId:%s", datasetId)

	res, err := client.R().
		SetHeader("SKYBOX_TOKEN_USER_KEY", TOKEN).
		SetResult(&result).
		Delete(DATA_SERVICE_HOST + fmt.Sprintf(DeleteDatasetUrl, datasetId))

	if err != nil {
		log.Error("resty DeleteDataset: %v", err)
		return nil, err
	}
	if res.IsError() {
		log.Error("DeleteDataset error.req= %+v res=%+v", datasetId, res)
		return nil, errors.New(res.String())
	}
	log.Info("DeleteDataset success. req= %+v res=%+v", datasetId, res)
	return &result, nil
}

func DeleteDatasetByName(labelId, datasetName string) error {
	//根据名称查询你数据集id
	datasets, err := QueryDataset(labelId, datasetName)
	if err != nil {
		log.Error("DeleteDatasetByName failed,labelId=%s datasetName=%s,err=%v ", labelId, datasetName, err)
		return err
	}
	if datasets == nil {
		log.Info("DeleteDatasetByName failed,labelId=%s datasetName=%s,err=result is empty ", labelId, datasetName)
		return nil
	}
	var datasetId string
	for _, dataset := range datasets.Content {
		if dataset.Name == datasetName {
			datasetId = dataset.ID
			break
		}
	}
	//根据id删除数据集
	if datasetId == "" {
		log.Info("DeleteDatasetByName failed,labelId=%s datasetName=%s,err=result is empty ", labelId, datasetName)
		return nil
	}

	_, err = DeleteDataset(datasetId)
	return err
}

///flintDatasetGroups?labelId=670ce12428c3ca0cade97489&name=test&size=10&page=1&sort=lastModifiedDate,desc
//QueryDataset 查询数据集
func QueryDataset(labelId, datasetName string) (*QueryDatasetResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result QueryDatasetResponse
	log.Info("QueryDataset labelId=%s,datasetName=%s", labelId, datasetName)

	res, err := client.R().
		SetHeader("SKYBOX_TOKEN_USER_KEY", TOKEN).
		SetResult(&result).
		Get(DATA_SERVICE_HOST + FlintDatasetGroupsUrl + "?labelId=" + labelId + "&name=" + datasetName + "&size=100&page=1&sort=lastModifiedDate,desc")

	if err != nil {
		log.Error("resty QueryDataset,labelId=%s,datasetName=%s", labelId, datasetName)
		return nil, err
	}
	if res.IsError() {
		log.Error("QueryDataset error.labelId=%s,datasetName=%s res=%+v", labelId, datasetName, res)
		return nil, errors.New(res.String())
	}
	log.Info("QueryDataset success.labelId=%s,datasetName=%s  res=%+v", labelId, datasetName, res)
	return &result, nil
}

//CreateOnlineService 创建在线服务
func CreateOnlineService(req CreateOnlineServiceReq) (*IFLYTekCommonResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result IFLYTekCommonResponse
	log.Info("CreateOnlineService req:%+v", req)

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("SKYBOX_TOKEN_USER_KEY", TOKEN).
		SetBody(req).
		SetResult(&result).
		Post(ONLINE_SERVICE_HOST + CreateOnlineServiceUrl)

	if err != nil {
		log.Error("resty CreateOnlineService: %v", err)
		return nil, err
	}
	if res.IsError() {
		log.Error("CreateOnlineService error.req= %+v res=%+v", req, res)
		return nil, errors.New(res.String())
	}
	if result.Code != 0 {
		log.Error("CreateOnlineService error.req= %+v res=%+v", req, res)
		return nil, errors.New(res.String())
	}
	log.Info("CreateOnlineService success.req= %+v res=%+v", req, res)
	return &result, nil
}

//QueryOnlineServiceStatus
func QueryOnlineServiceStatus(serviceId string) (*IFLYTekServiceResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result IFLYTekServiceResponse
	log.Info("QueryOnlineServiceStatus serviceId=%s", serviceId)
	req := map[string]string{"service_id": serviceId}

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("SKYBOX_TOKEN_USER_KEY", TOKEN).
		SetBody(req).
		SetResult(&result).
		Post(ONLINE_SERVICE_HOST + QueryOnlineServiceStatusUrl)

	if err != nil {
		log.Error("resty QueryOnlineServiceStatus, serviceId=%s", serviceId)
		return nil, err
	}
	if res.IsError() {
		log.Error("QueryOnlineServiceStatus error.serviceId=%s, res=%+v", serviceId, res)
		return nil, errors.New(res.String())
	}
	log.Info("QueryOnlineServiceStatus success.serviceId=%s, res=%+v", serviceId, res)
	return &result, nil
}

//StopOnlineService
func StopOnlineService(serviceId string) (*IFLYTekCommonResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result IFLYTekCommonResponse
	log.Info("StopOnlineService serviceId=%s", serviceId)
	req := map[string]string{"service_id": serviceId}

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(req).
		SetResult(&result).
		Delete(ONLINE_SERVICE_HOST + CreateOnlineServiceUrl)

	if err != nil {
		log.Error("resty StopOnlineService, serviceId=%s", serviceId)
		return nil, err
	}
	if res.IsError() {
		log.Error("StopOnlineService error.serviceId=%s, res=%+v", serviceId, res)
		return nil, errors.New(res.String())
	}
	log.Info("StopOnlineService success.serviceId=%s, res=%+v", serviceId, res)
	return &result, nil
}

// AuthParams 包含鉴权所需的参数
type AuthParams struct {
	APIKey    string
	APISecret string
	Params    url.Values
}

type PublishModelReq struct {
	Name            string `json:"name"`
	BaseModel       string `json:"baseModel"`
	TrainingMethod  string `json:"trainingMethod"`
	SystemStructure string `json:"systemStructure"`
	AcceleratorCard string `json:"acceleratorCard"`
	ModelPath       string `json:"modelPath"`
	ModelType       string `json:"modelType"`
	TaskVersionId   int64  `json:"taskVersionId"`
	Source          string `json:"source"`
	Description     string `json:"description"`
	CreatedBy       string `json:"createdBy"`
	TrainType       int    `json:"trainType"`
}

type PublishModelResponse struct {
	AcceleratorCard string  `json:"acceleratorCard"`
	BaseModel       string  `json:"baseModel"`
	ClusterId       *string `json:"clusterId"` // 使用指针类型来表示可能为null的字段
	Description     string  `json:"description"`
	ID              string  `json:"id"`
	ModelPath       string  `json:"modelPath"`
	ModelType       string  `json:"modelType"`
	Name            string  `json:"name"`
	Params          *string `json:"params"` // 使用指针类型来表示可能为null的字段
	Size            *int    `json:"size"`   // 使用指针类型来表示可能为null的字段
	Source          string  `json:"source"`
	Status          int     `json:"status"`
	SystemStructure string  `json:"systemStructure"`
	TaskVersionId   int     `json:"taskVersionId"`
	TaskVersionName *string `json:"taskVersionName"` // 使用指针类型来表示可能为null的字段
	TemplateId      *string `json:"templateId"`      // 使用指针类型来表示可能为null的字段
	TenantId        string  `json:"tenantId"`
	TrainType       int     `json:"trainType"`
	Version         string  `json:"version"`
}

//PublishModel
func PublishModel(req PublishModelReq) (*PublishModelResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result PublishModelResponse
	log.Info("CreateOnlineService req:%+v", req)

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("SKYBOX_TOKEN_USER_KEY", TOKEN).
		SetBody(req).
		SetResult(&result).
		Post(MODEL_SERVICE_HOST + PublishModelUrl)

	if err != nil {
		log.Error("resty CreateOnlineService: %v", err)
		return nil, err
	}
	if res.IsError() {
		if strings.Contains(res.String(), "\"code\":4000") {
			return nil, models.ErrModelPublished{}
		}
		log.Error("CreateOnlineService error.req= %+v res=%+v", req, res)
		return nil, errors.New(res.String())
	}
	if strings.Contains(res.String(), "\"code\":4000") {
		return nil, models.ErrModelPublished{}
	}
	log.Info("CreateOnlineService success.req= %+v res=%+v", req, res)
	return &result, nil
}

type QueryModelDetailResponse struct {
	AcceleratorCard string          `json:"acceleratorCard"`
	BaseModel       string          `json:"baseModel"`
	FlintItemVoList FlintItemVoList `json:"flintItemVoList"`
	ID              string          `json:"id"`
	ModelFrame      *string         `json:"modelFrame"`
	ModelManagerId  *string         `json:"modelManagerId"`
	ModelMark       string          `json:"modelMark"`
	ModelPath       string          `json:"modelPath"`
	ModelType       string          `json:"modelType"`
	Name            string          `json:"name"`
	Params          *string         `json:"params"`
	Size            *int            `json:"size"`
	Source          string          `json:"source"`
	Status          int             `json:"status"`
	SystemStructure string          `json:"systemStructure"`
	TrainType       int             `json:"trainType"`
	Version         string          `json:"version"`
}

// FlintItemVoList represents the structure of the flintItemVoList field in the JSON data
type FlintItemVoList struct {
	Content          []FlintItem `json:"content"`
	Empty            bool        `json:"empty"`
	First            bool        `json:"first"`
	Last             bool        `json:"last"`
	Number           int         `json:"number"`
	NumberOfElements int         `json:"numberOfElements"`
}

// FlintItem represents the structure of each item in the content field of FlintItemVoList
type FlintItem struct {
	BucketName string  `json:"bucketName"`
	Category   *string `json:"category"`
	Count      *int    `json:"count"`
	DataPath   string  `json:"dataPath"`
	FileSize   string  `json:"fileSize"`
	Format     string  `json:"format"`
	IsDir      bool    `json:"isDir"`
	Name       string  `json:"name"`
}

func QueryModelDetail(modelId string) (*QueryModelDetailResponse, error) {
	checkSetting()
	client := getRestyClient()
	var result QueryModelDetailResponse
	log.Info("QueryModelDetail modelId:%s", modelId)

	res, err := client.R().
		SetHeader("SKYBOX_TOKEN_USER_KEY", TOKEN).
		SetResult(&result).
		Get(MODEL_SERVICE_HOST + fmt.Sprintf(QueryModelDetailUrl, modelId))

	if err != nil {
		log.Error("resty QueryModelDetail: %v", err)
		return nil, err
	}
	if res.IsError() {
		log.Error("QueryModelDetail error.modelId:%s res=%+v", modelId, res)
		return nil, errors.New(res.String())
	}
	log.Info("QueryModelDetail success.modelId:%s res=%+v", modelId, res)
	return &result, nil
}
