package models

import (
	"time"

	"code.gitea.io/gitea/modules/timeutil"
)

type ModelartsDeploy struct {
	JobID             string `xorm:"pk 'job_id'"`
	JobName           string
	DisplayJobName    string
	UserID            int64
	Status            string
	ModelID           string
	ModelType         string
	ModelName         string
	ModelStatus       string
	ServiceID         string
	ServiceName       string
	ServiceStatus     string
	CreateUnix        timeutil.TimeStamp
	UpdateUnix        timeutil.TimeStamp
	CompleteUnix      timeutil.TimeStamp
	InferAddr         string
	DeployUrl         string
	Finetune          bool
	FinetuneModelType int
	FinetuneCategory  int
	DeletedAt         time.Time `xorm:"deleted"`
}

type CreateDeployModelParams struct {
	ModelName      string   `json:"model_name"`
	ModelVersion   string   `json:"model_version"`
	ModelType      string   `json:"model_type"`
	SourceLocation string   `json:"source_location"`
	Runtime        string   `json:"runtime"`
	InstallType    []string `json:"install_type"`
	Prebuild       bool     `json:"prebuild"`
}

type CreateDeployModelResult struct {
	ModelID string `json:"model_id"`
}

type GetDeployModelResult struct {
	SourceLocation string `json:"source_location"`
	ModelName      string `json:"model_name"`
	ModelID        string `json:"model_id"`
	ModelStatus    string `json:"model_status"`
}

type DelDeployModelFail struct {
	ModelID   string `json:"model_id"`
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_message"`
}

type CreateDeployServiceParams struct {
	InferType   string           `json:"infer_type"`
	ServiceName string           `json:"service_name"`
	ClusterID   string           `json:"cluster_id"`
	Config      []ServiceConfig  `json:"config"`
	Schedule    []DeploySchedule `json:"schedule"`
}

type ServiceConfig struct {
	Specification string `json:"specification"`
	ModelID       string `json:"model_id"`
	InstanceCount int    `json:"instance_count"`
	//CustomSpec    ServiceCustomSpec `json:"custom_spec"`
}

type ServiceCustomSpec struct {
	CPU        int        `json:"cpu"`
	Memory     int        `json:"memory"`
	Ascend     int        `json:"ascend_a310"`
	CPUInfo    CPUInfo    `json:"cpu_info"`
	MemoryInfo MemoryInfo `json:"memory_info"`
	NPUInfo    NPUInfo    `json:"npu_info"`
}

type CPUInfo struct {
	CPU  int    `json:"cpu"`
	Arch string `json:"arch"`
}

type MemoryInfo struct {
	Memory int    `json:"memory"`
	Unit   string `json:"unit"`
}

type NPUInfo struct {
	NPU     int    `json:"npu"`
	Brand   string `json:"brand"`
	Version string `json:"version"`
	Unit    string `json:"unit"`
	Memory  int    `json:"memory"`
}

type DeploySchedule struct {
	Duration int    `json:"duration"`
	TimeUnit string `json:"time_unit"`
	Type     string `json:"type"`
}

type CreateDeployServiceResult struct {
	ServiceID   string   `json:"service_id"`
	ResourceIDs []string `json:"resource_ids"`
}

type GetDeployServiceResult struct {
	ServiceID   string `json:"service_id"`
	ServiceName string `json:"service_name"`
	Status      string `json:"status"`
	InferAddr   string `json:"access_address"`
}

type UpdateDeployServiceParams struct {
	Status string `json:"status"`
}

type PanguInferParams struct {
	Text string `json:"text"`
}
type PanguInferResult struct {
	GenerateResult string `json:"generate_result"`
}

type PanguInferError struct {
	ErrorCode string `json:"erno"`
	ErrorMsg  string `json:"msg"`
}

func CreateModelartsDeploy(deploy *ModelartsDeploy) (err error) {

	sess := x.NewSession()
	defer sess.Close()

	if err := sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.Insert(deploy); err != nil {
		return err
	}
	return sess.Commit()
}

// write a function to get the deployment by job id
func GetModelartsDeployByJobID(jobID string) (deploy *ModelartsDeploy, err error) {
	deploy = new(ModelartsDeploy)
	has, err := x.Where("job_id = ?", jobID).Get(deploy)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrModelartsDeployNotExist{jobID}
	}
	return deploy, nil
}

func UpdateDeploy(deploy *ModelartsDeploy) error {
	return updateDeploy(x, deploy)
}

func updateDeploy(e Engine, deploy *ModelartsDeploy) error {
	_, err := e.ID(deploy.JobID).AllCols().Update(deploy)
	return err
}

// write a function return all deployments
func GetAllModelartsDeploys() ([]*ModelartsDeploy, error) {
	return getAllModelartsDeploys(x)
}

func getAllModelartsDeploys(e Engine) ([]*ModelartsDeploy, error) {
	deploys := make([]*ModelartsDeploy, 0)
	return deploys, e.Find(&deploys)
}

// get deployment status by jobid
func GetModelartsDeployStatusByJobID(jobID string) (status string, err error) {
	deploy, err := GetModelartsDeployByJobID(jobID)
	if err != nil {
		return "", err
	}
	return deploy.Status, nil
}

func GetRunningServiceByUser(userID int64) ([]*ModelartsDeploy, error) {
	return getRunningServiceByUser(x, userID)
}
func getRunningServiceByUser(e Engine, userID int64) ([]*ModelartsDeploy, error) {
	deploys := make([]*ModelartsDeploy, 0, 10)
	return deploys, e.Where("user_id = ? AND status IN ('BUILDING','SUCCEEDED', 'DEPLOYING', 'WAITING')", userID).Find(&deploys)
}

func GetAllRunningService() ([]*ModelartsDeploy, error) {
	return getAllRunningService(x)
}
func getAllRunningService(e Engine) ([]*ModelartsDeploy, error) {
	deploys := make([]*ModelartsDeploy, 0, 10)
	return deploys, e.Where("service_status IN ('deploying', 'running')").Find(&deploys)
}

// todo delete deployment in database
func DeleteModelartsDeploy(jobID string) error {
	return deleteModelartsDeploy(x, jobID)
}
func deleteModelartsDeploy(e Engine, jobID string) error {
	_, err := e.Where("job_id = ?", jobID).Delete(&ModelartsDeploy{})
	return err
}

func DeployStatusConvert(status string) string {
	//status, err := GetModelartsDeployStatusByJobID(jobID)
	//if err != nil || status == "" {
	if status == "" {
		return status
	} else {
		var statusConvert string
		switch status {
		case "running":
			statusConvert = "SUCCEEDED"
		case "deploying":
			statusConvert = "DEPLOYING"
		case "publishing":
			statusConvert = "BUILDING"
		case "building":
			statusConvert = "BUILDING"
		case "published":
			statusConvert = "WAITING"
		case "stopped":
			statusConvert = "STOP"
		case "stopping":
			statusConvert = "STOP"
		default:
			statusConvert = "FAILED"
		}
		return statusConvert
	}
}

func GetModelartsDeployFinishTimebyJobID(jobID string) (finishTime timeutil.TimeStamp, err error) {
	finishTime = timeutil.TimeStamp(0)
	deploy, err := GetModelartsDeployByJobID(jobID)
	if err != nil || deploy.CompleteUnix == timeutil.TimeStamp(0) {
		return finishTime, err
	} else {
		finishTime = deploy.CompleteUnix.Add(int64(30 * 60))
		return finishTime, nil
	}
}
