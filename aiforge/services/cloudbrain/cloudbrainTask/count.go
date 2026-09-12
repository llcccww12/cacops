package cloudbrainTask

import (
	"strconv"

	"code.gitea.io/gitea/models"
)

type StatusInfo struct {
	CloudBrainTypes  []int
	JobType          []models.JobType
	NotFinalStatuses []string
	ComputeResource  string
}

var CloudbrainOneNotFinalStatuses = []string{string(models.JobWaiting), string(models.JobRunning)}
var CloudbrainTwoNotFinalStatuses = []string{string(models.ModelArtsTrainJobInit), string(models.ModelArtsTrainJobImageCreating), string(models.ModelArtsTrainJobSubmitTrying), string(models.ModelArtsTrainJobWaiting), string(models.ModelArtsTrainJobRunning), string(models.ModelArtsTrainJobScaling), string(models.ModelArtsTrainJobCheckInit), string(models.ModelArtsTrainJobCheckRunning), string(models.ModelArtsTrainJobCheckRunningCompleted)}
var GrampusNotFinalStatuses = []string{models.GrampusStatusWaiting, models.GrampusStatusRunning, models.GrampusStatusPending}
var StatusInfoDict = map[string]StatusInfo{string(models.JobTypeDebug) + "-" + strconv.Itoa(models.TypeCloudBrainOne): {
	CloudBrainTypes:  []int{models.TypeCloudBrainOne},
	JobType:          []models.JobType{models.JobTypeDebug},
	NotFinalStatuses: CloudbrainOneNotFinalStatuses,
	ComputeResource:  models.GPUResource,
}, string(models.JobTypeTrain) + "-" + strconv.Itoa(models.TypeCloudBrainOne): {
	CloudBrainTypes:  []int{models.TypeCloudBrainOne},
	JobType:          []models.JobType{models.JobTypeTrain},
	NotFinalStatuses: CloudbrainOneNotFinalStatuses,
	ComputeResource:  models.GPUResource,
}, string(models.JobTypeInference) + "-" + strconv.Itoa(models.TypeCloudBrainOne): {
	CloudBrainTypes:  []int{models.TypeCloudBrainOne},
	JobType:          []models.JobType{models.JobTypeInference},
	NotFinalStatuses: CloudbrainOneNotFinalStatuses,
	ComputeResource:  models.GPUResource,
}, string(models.JobTypeBenchmark) + "-" + strconv.Itoa(models.TypeCloudBrainOne): {
	CloudBrainTypes:  []int{models.TypeCloudBrainOne},
	JobType:          []models.JobType{models.JobTypeBenchmark, models.JobTypeBrainScore, models.JobTypeSnn4imagenet, models.JobTypeSnn4Ecoset, models.JobTypeSim2BrainSNN},
	NotFinalStatuses: CloudbrainOneNotFinalStatuses,
	ComputeResource:  models.GPUResource,
}, string(models.JobTypeDebug) + "-" + strconv.Itoa(models.TypeCloudBrainTwo): {
	CloudBrainTypes:  []int{models.TypeCloudBrainTwo, models.TypeCDCenter},
	JobType:          []models.JobType{models.JobTypeDebug},
	NotFinalStatuses: []string{string(models.ModelArtsCreateQueue), string(models.ModelArtsCreating), string(models.ModelArtsStarting), string(models.ModelArtsReadyToStart), string(models.ModelArtsResizing), string(models.ModelArtsStartQueuing), string(models.ModelArtsRunning), string(models.ModelArtsRestarting)},
	ComputeResource:  models.NPUResource,
}, string(models.JobTypeTrain) + "-" + strconv.Itoa(models.TypeCloudBrainTwo): {
	CloudBrainTypes:  []int{models.TypeCloudBrainTwo},
	JobType:          []models.JobType{models.JobTypeTrain},
	NotFinalStatuses: CloudbrainTwoNotFinalStatuses,
	ComputeResource:  models.NPUResource,
}, string(models.JobTypeInference) + "-" + strconv.Itoa(models.TypeCloudBrainTwo): {
	CloudBrainTypes:  []int{models.TypeCloudBrainTwo},
	JobType:          []models.JobType{models.JobTypeInference},
	NotFinalStatuses: CloudbrainTwoNotFinalStatuses,
	ComputeResource:  models.NPUResource,
}, string(models.JobTypeTrain) + "-" + strconv.Itoa(models.TypeC2Net) + "-" + models.GPUResource: {
	CloudBrainTypes:  []int{models.TypeC2Net},
	JobType:          []models.JobType{models.JobTypeTrain},
	NotFinalStatuses: GrampusNotFinalStatuses,
	ComputeResource:  models.GPUResource,
}, string(models.JobTypeTrain) + "-" + strconv.Itoa(models.TypeC2Net) + "-" + models.NPUResource: {
	CloudBrainTypes:  []int{models.TypeC2Net},
	JobType:          []models.JobType{models.JobTypeTrain},
	NotFinalStatuses: GrampusNotFinalStatuses,
	ComputeResource:  models.NPUResource,
}, string(models.JobTypeDebug) + "-" + strconv.Itoa(models.TypeC2Net) + "-" + models.GPUResource: {
	CloudBrainTypes:  []int{models.TypeC2Net},
	JobType:          []models.JobType{models.JobTypeDebug},
	NotFinalStatuses: GrampusNotFinalStatuses,
	ComputeResource:  models.GPUResource,
}, string(models.JobTypeDebug) + "-" + strconv.Itoa(models.TypeC2Net) + "-" + models.NPUResource: {
	CloudBrainTypes:  []int{models.TypeC2Net},
	JobType:          []models.JobType{models.JobTypeDebug},
	NotFinalStatuses: GrampusNotFinalStatuses,
	ComputeResource:  models.NPUResource,
}, string(models.JobTypeDebug) + "-" + strconv.Itoa(models.TypeC2Net) + "-" + models.GCUResource: {
	CloudBrainTypes:  []int{models.TypeC2Net},
	JobType:          []models.JobType{models.JobTypeDebug},
	NotFinalStatuses: GrampusNotFinalStatuses,
	ComputeResource:  models.GCUResource,
}}

var AlleNotFinalStatuses = []string{string(models.JobWaiting), string(models.JobRunning), string(models.ModelArtsTrainJobInit), string(models.ModelArtsTrainJobImageCreating), string(models.ModelArtsTrainJobSubmitTrying), string(models.ModelArtsTrainJobWaiting), string(models.ModelArtsTrainJobRunning), string(models.ModelArtsTrainJobScaling), string(models.ModelArtsTrainJobCheckInit), string(models.ModelArtsTrainJobCheckRunning), string(models.ModelArtsTrainJobCheckRunningCompleted), models.GrampusStatusWaiting, models.GrampusStatusRunning, models.GrampusStatusPending,
	string(models.ModelArtsCreateQueue),
	string(models.ModelArtsCreating),
	string(models.ModelArtsStartQueuing),
	string(models.ModelArtsReadyToStart),
	string(models.ModelArtsStarting),
	string(models.ModelArtsRestarting),
	string(models.ModelArtsRunning),
	models.LocalStatusPreparing,
	models.LocalStatusCreating,
}

func GetNotFinalStatusTaskCount(uid int64, jobType string) (int, error) {
	jobNewType := []string{jobType}
	if models.IsBenchMarkJobType(jobType) {
		jobNewType = models.AllBenchMarkJobType()
	}

	return models.GetNotFinalStatusTaskCount(uid, AlleNotFinalStatuses, jobNewType)

}
