package grampus

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"code.gitea.io/gitea/modules/modelarts"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/notification"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
)

const (
	JobPath = "job/"

	ProcessorTypeNPU = "npu.huawei.com/NPU"
	ProcessorTypeGPU = "nvidia.com/gpu"
	ProcessorTypeGCU = "enflame-tech.com/gcu"
	ProcessorTypeMLU = "cambricon.com/mlu"

	GcuWorkDir              = "/tmp/"
	GpuWorkDir              = "/tmp/"
	NpuWorkDir              = "/cache/"
	NpuLocalLogUrl          = "/tmp/train.log"
	CommandPrepareScriptNpu = ";mkdir -p output;mkdir -p code;mkdir -p dataset;mkdir -p pretrainmodel;"

	CodeArchiveName = "master.zip"

	BucketRemote       = "grampus"
	RemoteModelPath    = "/output"
	autoStopDurationMs = 4 * 60 * 60 * 1000
	CommandGpuDebug    = "jupyter lab --ServerApp.shutdown_no_activity_timeout=%s --TerminalManager.cull_inactive_timeout=%s --TerminalManager.cull_interval=%s --MappingKernelManager.cull_idle_timeout=%s --MappingKernelManager.cull_interval=%s --MappingKernelManager.cull_connected=True --MappingKernelManager.cull_busy=True --no-browser --ip=0.0.0.0 --allow-root --notebook-dir='%s' --port=$OCTOPUS_NOTEBOOK_PORT --LabApp.token='' --LabApp.allow_origin='*' --LabApp.base_url=$OCTOPUS_NOTEBOOK_BASE_URL;"
)

var (
	poolInfos   *models.PoolInfos
	FlavorInfos *setting.StFlavorInfos
	ImageInfos  *setting.StImageInfosModelArts

	SpecialPools *models.SpecialPools

	CommandPrepareScriptGpu = ";mkdir -p output;mkdir -p code;mkdir -p dataset;mkdir -p pretrainmodel;"

	MultiNodeConfig *modelarts.MultiNodes
)

type GenerateTrainJobReq struct {
	JobName  string
	Command  string
	ImageUrl string //与image_id二选一，都有的情况下优先image_url
	ImageId  string

	DisplayJobName    string
	Uuid              string
	Description       string
	CodeObsPath       string
	BootFile          string
	BootFileUrl       string
	DataUrl           string
	TrainUrl          string
	WorkServerNumber  int
	EngineID          int64
	CommitID          string
	IsLatestVersion   string
	BranchName        string
	PreVersionId      int64
	PreVersionName    string
	VersionCount      int
	EngineName        string
	TotalVersionCount int
	ComputeResource   string
	ProcessType       string

	DatasetNames       string
	DatasetInfos       map[string]models.DatasetInfo
	Params             string
	ModelName          string
	LabelName          string
	CkptName           string
	ModelId            string
	ModelVersion       string
	PreTrainModelPath  string
	PreTrainModelUrl   string
	Spec               *models.Specification
	CodeName           string
	PreTrainModelPaths []string
	CkptNames          []string
}

type GenerateNotebookJobReq struct {
	JobName           string
	Command           string
	ImageUrl          string
	ImageId           string
	DisplayJobName    string
	Uuid              string
	Description       string
	CodeStoragePath   string
	CommitID          string
	BranchName        string
	ComputeResource   string
	ProcessType       string
	DatasetNames      string
	DatasetInfos      map[string]models.DatasetInfo
	ModelName         string
	LabelName         string
	CkptName          string
	ModelId           string
	ModelVersion      string
	PreTrainModelPath string
	PreTrainModelUrl  string
	Spec              *models.Specification
	CodeName          string
	ModelPath         string //参考启智GPU调试， 挂载/model目录用户的模型可以输出到这个目录
	ModelStorageType  int
}

func getEndPoint() string {
	index := strings.Index(setting.Endpoint, "//")
	endpoint := setting.Endpoint[index+2:]
	return endpoint
}

func getDatasetGrampus(datasetInfos map[string]models.DatasetInfo) []models.GrampusDataset {
	var datasetGrampus []models.GrampusDataset
	endPoint := getEndPoint()
	for _, datasetInfo := range datasetInfos {
		datasetGrampus = append(datasetGrampus, models.GrampusDataset{
			Name:          datasetInfo.FullName,
			Bucket:        setting.Bucket,
			EndPoint:      endPoint,
			ObjectKey:     datasetInfo.DataLocalPath + datasetInfo.FullName,
			ReadOnly:      true,
			ContainerPath: "/tmp/dataset/" + datasetInfo.FullName,
		})

	}
	return datasetGrampus
}
func getDatasetGPUGrampus(datasetInfos map[string]models.DatasetInfo, containerPath string) []models.GrampusDataset {
	var datasetGrampus []models.GrampusDataset
	for _, datasetInfo := range datasetInfos {
		datasetGrampus = append(datasetGrampus, models.GrampusDataset{
			Name:          datasetInfo.FullName,
			Bucket:        setting.Attachment.Minio.Bucket,
			EndPoint:      setting.Attachment.Minio.Endpoint,
			ObjectKey:     datasetInfo.DataLocalPath,
			ReadOnly:      true,
			ContainerPath: containerPath + "/" + datasetInfo.FullName,
		})

		//command += "cp " + containerPath + "/'" + datasetInfo.Name + "'/" + uuid + " " + cpPath + "/'" + datasetInfo.FullName + "';"

	}
	return datasetGrampus
}
func getDatasetGCUGrampus(datasetInfos map[string]models.DatasetInfo, containerPath string) []models.GrampusDataset {
	var datasetGrampus []models.GrampusDataset
	obsEndPoint := getEndPoint()
	for _, datasetInfo := range datasetInfos {
		if datasetInfo.Type == models.TypeCloudBrainOne {
			datasetGrampus = append(datasetGrampus, models.GrampusDataset{
				Name:          datasetInfo.FullName,
				Bucket:        setting.Attachment.Minio.Bucket,
				EndPoint:      setting.Attachment.Minio.Endpoint,
				ObjectKey:     datasetInfo.DataLocalPath,
				ReadOnly:      true,
				ContainerPath: containerPath + "/" + datasetInfo.FullName,
			})

		} else {
			datasetGrampus = append(datasetGrampus, models.GrampusDataset{
				Name:          datasetInfo.FullName,
				Bucket:        setting.Bucket,
				EndPoint:      obsEndPoint,
				ObjectKey:     datasetInfo.DataLocalPath + datasetInfo.FullName,
				ContainerPath: containerPath + "/" + datasetInfo.Name,
			})
		}

	}
	return datasetGrampus
}

func GenerateNotebookJob(ctx *context.Context, req *GenerateNotebookJobReq) (jobId string, err error) {
	createTime := timeutil.TimeStampNow()
	containerPrefix := ""
	if ProcessorTypeGCU == req.ProcessType {
		containerPrefix = "/tmp"
	}

	var datasetGrampus []models.GrampusDataset
	var codeGrampus models.GrampusDataset
	imageUrl := req.ImageUrl
	if ProcessorTypeNPU == req.ProcessType {
		datasetGrampus = getDatasetGrampus(req.DatasetInfos)
		if len(req.ModelName) != 0 {
			datasetGrampus = append(datasetGrampus, models.GrampusDataset{
				Name:      req.ModelName,
				Bucket:    setting.Bucket,
				EndPoint:  getEndPoint(),
				ReadOnly:  true,
				ObjectKey: req.PreTrainModelPath,
			})
		}

		codeGrampus = models.GrampusDataset{
			Name:      req.CodeName,
			Bucket:    setting.Bucket,
			EndPoint:  getEndPoint(),
			ObjectKey: req.CodeStoragePath + cloudbrain.DefaultBranchName + ".zip",
			ReadOnly:  false,
		}
		imageUrl = ""
		req.Command = ""
	} else {
		if ProcessorTypeGCU == req.ProcessType {
			datasetGrampus = getDatasetGCUGrampus(req.DatasetInfos, containerPrefix+"/dataset")
		} else {
			datasetGrampus = getDatasetGPUGrampus(req.DatasetInfos, "/dataset")
		}
		if len(req.ModelName) != 0 {
			if req.ModelStorageType == models.TypeCloudBrainOne {
				datasetGrampus = append(datasetGrampus, models.GrampusDataset{
					Name:          req.ModelName,
					Bucket:        setting.Attachment.Minio.Bucket,
					EndPoint:      setting.Attachment.Minio.Endpoint,
					ObjectKey:     req.PreTrainModelPath,
					ReadOnly:      true,
					ContainerPath: containerPrefix + cloudbrain.PretrainModelMountPath + "/" + req.CkptName,
				})
			} else {
				datasetGrampus = append(datasetGrampus, models.GrampusDataset{
					Name:          req.ModelName,
					Bucket:        setting.Bucket,
					EndPoint:      getEndPoint(),
					ReadOnly:      true,
					ObjectKey:     req.PreTrainModelPath,
					ContainerPath: containerPrefix + cloudbrain.PretrainModelMountPath + "/" + req.CkptName,
				})
			}

		}
		codeArchiveName := cloudbrain.DefaultBranchName + ".zip"
		codeGrampus = models.GrampusDataset{
			Name:          req.CodeName,
			Bucket:        setting.Attachment.Minio.Bucket,
			EndPoint:      setting.Attachment.Minio.Endpoint,
			ObjectKey:     req.CodeStoragePath + codeArchiveName,
			ReadOnly:      false,
			ContainerPath: containerPrefix + cloudbrain.CodeMountPath + "/" + codeArchiveName,
		}
		if ProcessorTypeGCU == req.ProcessType {
			imageUrl = ""
		}
		req.Command = fmt.Sprintf(CommandGpuDebug, setting.CullIdleTimeout, setting.CullIdleTimeout, setting.CullInterval, setting.CullIdleTimeout, setting.CullInterval, containerPrefix+cloudbrain.CodeMountPath)
		log.Info("debug command:" + req.Command)

	}
	datasetGrampusJson, _ := json.Marshal(datasetGrampus)
	log.Info("datasetGrampusJson=" + string(datasetGrampusJson))
	jobResult, err := createNotebookJob(models.CreateGrampusNotebookRequest{
		Name: req.JobName,
		Tasks: []models.GrampusNotebookTask{
			{
				Name:             req.JobName,
				ResourceSpecId:   req.Spec.SourceSpecId,
				ImageId:          req.ImageId,
				ImageUrl:         imageUrl,
				Datasets:         datasetGrampus,
				Code:             codeGrampus,
				AutoStopDuration: autoStopDurationMs,
				Capacity:         setting.Capacity,
				Command:          req.Command,
				CenterID: req.Spec.GetAvailableCenterIds(models.GetAvailableCenterIdOpts{
					UserId:  ctx.User.ID,
					JobType: models.JobTypeDebug,
				}),
			},
		},
	})
	if err != nil {
		log.Error("createNotebookJob failed: %v", err.Error())
		return "", err
	}

	jobID := jobResult.JobInfo.JobID
	err = models.CreateCloudbrain(&models.Cloudbrain{
		Status:           TransTrainJobStatus(jobResult.JobInfo.Status),
		UserID:           ctx.User.ID,
		RepoID:           ctx.Repo.Repository.ID,
		JobID:            jobID,
		JobName:          req.JobName,
		DisplayJobName:   req.DisplayJobName,
		JobType:          string(models.JobTypeDebug),
		Type:             models.TypeC2Net,
		Uuid:             req.Uuid,
		DatasetName:      req.DatasetNames,
		CommitID:         req.CommitID,
		IsLatestVersion:  "1",
		ComputeResource:  req.ComputeResource,
		ImageID:          req.ImageId,
		BranchName:       req.BranchName,
		Description:      req.Description,
		WorkServerNumber: 1,
		EngineName:       req.ImageUrl,
		CreatedUnix:      createTime,
		UpdatedUnix:      createTime,
		Spec:             req.Spec,
		ModelName:        req.ModelName,
		ModelVersion:     req.ModelVersion,
		LabelName:        req.LabelName,
		PreTrainModelUrl: req.PreTrainModelUrl,
		CkptName:         req.CkptName,
		ModelId:          req.ModelId,
	})

	if err != nil {
		log.Error("CreateCloudbrain(%s) failed:%v", req.DisplayJobName, err.Error())
		return "", err
	}

	var actionType models.ActionType
	if req.ComputeResource == models.NPUResource {
		actionType = models.ActionCreateGrampusNPUDebugTask
	} else if req.ComputeResource == models.GPUResource {
		actionType = models.ActionCreateGrampusGPUDebugTask
	} else if req.ComputeResource == models.GCUResource {
		actionType = models.ActionCreateGrampusGCUDebugTask
	}
	task, err := models.GetCloudbrainByJobID(jobID)
	if err != nil {
		log.Error("GetCloudbrainByJobID failed: %v", err.Error())
		return "", err
	}

	stringId := strconv.FormatInt(task.ID, 10)
	notification.NotifyOtherTask(ctx.User, ctx.Repo.Repository, stringId, req.DisplayJobName, actionType)

	return jobID, nil
}

func GenerateTrainJob(ctx *context.Context, req *GenerateTrainJobReq) (jobId string, err error) {
	createTime := timeutil.TimeStampNow()

	var datasetGrampus, modelGrampus []models.GrampusDataset
	var ckptGrampus, codeGrampus, outputGrampus models.GrampusDataset
	if ProcessorTypeNPU == req.ProcessType {
		datasetGrampus = getDatasetGrampus(req.DatasetInfos)
		for i, ckptName := range req.CkptNames {
			if len(req.CkptNames) != 0 {
				ckptGrampus = models.GrampusDataset{
					Name:          ckptName,
					Bucket:        setting.Bucket,
					EndPoint:      getEndPoint(),
					ObjectKey:     req.PreTrainModelPaths[i],
					ContainerPath: "/tmp/pretrainmodel/" + req.CkptName,
					ReadOnly:      true,
				}
			}
			modelGrampus = append(modelGrampus, ckptGrampus)
		}
		codeGrampus = models.GrampusDataset{
			Name:          req.CodeName,
			Bucket:        setting.Bucket,
			EndPoint:      getEndPoint(),
			ObjectKey:     req.CodeObsPath + cloudbrain.DefaultBranchName + ".zip",
			ReadOnly:      false,
			ContainerPath: "/tmp/code/" + cloudbrain.DefaultBranchName + ".zip",
		}
		outputGrampus = models.GrampusDataset{
			ContainerPath:   "/tmp/output",
			GetBackEndpoint: getEndPoint(),
		}
	} else if ProcessorTypeGPU == req.ProcessType {
		datasetGrampus = getDatasetGPUGrampus(req.DatasetInfos, "/tmp/dataset")
		if len(req.ModelName) != 0 {
			modelGrampus = []models.GrampusDataset{ //model save as obs
				{
					Name:          req.ModelName,
					Bucket:        setting.Bucket,
					EndPoint:      getEndPoint(),
					ReadOnly:      true,
					ObjectKey:     req.PreTrainModelPath,
					ContainerPath: "/tmp/pretrainmodel/" + req.CkptName,
				},
			}
		}
		codeGrampus = models.GrampusDataset{
			Name:          req.CodeName,
			Bucket:        setting.Attachment.Minio.Bucket,
			EndPoint:      setting.Attachment.Minio.Endpoint,
			ObjectKey:     setting.CBCodePathPrefix + req.JobName + cloudbrain.CodeMountPath + "/" + cloudbrain.DefaultBranchName + ".zip",
			ReadOnly:      false,
			ContainerPath: "/tmp/code/" + cloudbrain.DefaultBranchName + ".zip",
		}
		outputGrampus = models.GrampusDataset{
			ContainerPath:   "/tmp/output",
			GetBackEndpoint: setting.Attachment.Minio.Endpoint,
		}

	} else if ProcessorTypeGCU == req.ProcessType {
		datasetGrampus = getDatasetGCUGrampus(req.DatasetInfos, "/tmp/dataset")
		if len(req.ModelName) != 0 {
			modelGrampus = []models.GrampusDataset{ //model save as obs
				{
					Name:          req.ModelName,
					Bucket:        setting.Bucket,
					EndPoint:      getEndPoint(),
					ReadOnly:      true,
					ObjectKey:     req.PreTrainModelPath,
					ContainerPath: "/tmp/pretrainmodel/" + req.CkptName,
				},
			}
		}
		codeGrampus = models.GrampusDataset{
			Name:          req.CodeName,
			Bucket:        setting.Attachment.Minio.Bucket,
			EndPoint:      setting.Attachment.Minio.Endpoint,
			ObjectKey:     setting.CBCodePathPrefix + req.JobName + cloudbrain.CodeMountPath + "/" + cloudbrain.DefaultBranchName + ".zip",
			ReadOnly:      false,
			ContainerPath: "/tmp/code/" + cloudbrain.DefaultBranchName + ".zip",
		}
		outputGrampus = models.GrampusDataset{
			ContainerPath:   "/tmp/output",
			GetBackEndpoint: setting.Attachment.Minio.Endpoint,
		}
	}

	modelGrampusJson, _ := json.Marshal(modelGrampus)
	log.Info("train job modelGrampus=" + string(modelGrampusJson))

	jobResult, err := createJob(models.CreateGrampusJobRequest{
		Name: req.JobName,
		Tasks: []models.GrampusTasks{
			{
				Name:           req.JobName,
				Command:        req.Command,
				ResourceSpecId: req.Spec.SourceSpecId,
				ImageId:        req.ImageId,
				ImageUrl:       req.ImageUrl,
				CenterID: req.Spec.GetAvailableCenterIds(models.GetAvailableCenterIdOpts{
					UserId:  ctx.User.ID,
					JobType: models.JobTypeTrain,
				}),
				ReplicaNum:       1,
				Datasets:         datasetGrampus,
				Models:           modelGrampus,
				Code:             codeGrampus,
				BootFile:         req.BootFile,
				OutPut:           outputGrampus,
				WorkServerNumber: req.WorkServerNumber,
			},
		},
	})
	if err != nil {
		log.Error("createJob failed: %v", err.Error())
		return "", err
	}

	jobID := jobResult.JobInfo.JobID
	err = models.CreateCloudbrain(&models.Cloudbrain{
		Status:            TransTrainJobStatus(jobResult.JobInfo.Status),
		UserID:            ctx.User.ID,
		RepoID:            ctx.Repo.Repository.ID,
		JobID:             jobID,
		JobName:           req.JobName,
		DisplayJobName:    req.DisplayJobName,
		JobType:           string(models.JobTypeTrain),
		Type:              models.TypeC2Net,
		Uuid:              req.Uuid,
		DatasetName:       req.DatasetNames,
		CommitID:          req.CommitID,
		IsLatestVersion:   req.IsLatestVersion,
		ComputeResource:   req.ComputeResource,
		ImageID:           req.ImageId,
		TrainUrl:          req.TrainUrl,
		BranchName:        req.BranchName,
		Parameters:        req.Params,
		BootFile:          req.BootFile,
		DataUrl:           req.DataUrl,
		Description:       req.Description,
		WorkServerNumber:  req.WorkServerNumber,
		EngineName:        req.EngineName,
		VersionCount:      req.VersionCount,
		TotalVersionCount: req.TotalVersionCount,
		CreatedUnix:       createTime,
		UpdatedUnix:       createTime,
		Spec:              req.Spec,
		ModelName:         req.ModelName,
		ModelVersion:      req.ModelVersion,
		LabelName:         req.LabelName,
		PreTrainModelUrl:  req.PreTrainModelUrl,
		CkptName:          req.CkptName,
		ModelId:           req.ModelId,
	})

	if err != nil {
		log.Error("CreateCloudbrain(%s) failed:%v", req.DisplayJobName, err.Error())
		return "", err
	}

	var actionType models.ActionType
	if req.ComputeResource == models.NPUResource {
		actionType = models.ActionCreateGrampusNPUTrainTask
	} else if req.ComputeResource == models.GPUResource {
		actionType = models.ActionCreateGrampusGPUTrainTask
	} else if req.ComputeResource == models.GCUResource {
		actionType = models.ActionCreateGrampusGCUTrainTask
	} else if req.ComputeResource == models.ILUVATAR {
		actionType = models.ActionCreateGrampusILUVATARTrainTask
	} else if req.ComputeResource == models.DCU {
		actionType = models.ActionCreateGrampusDCUTrainTask
	} else if req.ComputeResource == models.BIREN {
		actionType = models.ActionCreateGrampusBIRENGPUTrainTask
	} else if req.ComputeResource == models.METAX {
		actionType = models.ActionCreateGrampusMETAXGPGPUTrainTask
	}

	notification.NotifyOtherTask(ctx.User, ctx.Repo.Repository, jobID, req.DisplayJobName, actionType)

	return jobID, nil
}

func getCentersParamter(ctx *context.Context, req *GenerateTrainJobReq) ([]string, []string) {
	var centerID []string
	var centerName []string

	includeCenters := make(map[string]string)
	excludeCenters := make(map[string]string)

	if SpecialPools != nil {
		for _, pool := range SpecialPools.Pools {
			if !pool.IsExclusive && strings.Contains(req.ComputeResource, pool.Type) {
				org, _ := models.GetOrgByName(pool.Org)
				if org != nil {
					isOrgMember, _ := models.IsOrganizationMember(org.ID, ctx.User.ID)
					if isOrgMember {
						for _, info := range pool.Pool {
							includeCenters[info.Queue] = info.Value
						}
					} else {
						for _, info := range pool.Pool {
							excludeCenters[info.Queue] = info.Value
						}
					}
				}
			}
		}

	}

	if len(includeCenters) > 0 {
		//如果有专属资源池，根据专属资源池指定智算中心
		for k, v := range includeCenters {
			centerID = append(centerID, k)
			centerName = append(centerName, v)
		}
	} else if len(excludeCenters) > 0 {
		//否则，有要排除的中心，先获取所有中心，删除其中的排除中心，得到指定的智算中心
		allCenters := make(map[string]string)
		specs, err := GetResourceSpecs(req.ProcessType)
		if err == nil {
			for _, info := range specs.Infos {
				for _, center := range info.Centers {
					allCenters[center.ID] = center.Name
				}

			}
		}

		for k, _ := range excludeCenters {
			delete(allCenters, k)
		}

		for k, v := range allCenters {
			centerID = append(centerID, k)
			centerName = append(centerName, v)
		}

	}
	return centerID, centerName
}

func TransTrainJobStatus(status string) string {
	if status == models.GrampusStatusPending {
		status = models.GrampusStatusWaiting
	}

	return strings.ToUpper(status)
}

func GetNpuModelRemoteObsUrl(jobName string) string {
	return "s3:///" + BucketRemote + "/" + GetNpuModelObjectKey(jobName)
}

func GetNpuModelObjectKey(jobName string) string {
	return setting.CodePathPrefix + jobName + RemoteModelPath + "/" + models.ModelSuffix
}
func GetGPUModelObjectKey4Grampus(jobName string) string {
	return setting.CodePathPrefix + jobName + "0" + RemoteModelPath
}
func GetGPUModelObjectKey(jobName string) string {
	return setting.CBCodePathPrefix + jobName + cloudbrain.ModelMountPath + "/"
}

func GetRemoteEndPoint(aiCenterID string) string {
	var endPoint string
	for _, info := range setting.CenterInfos.Info {
		if info.CenterID == aiCenterID {
			endPoint = info.Endpoint
			break
		}
	}

	return endPoint
}

func GetCenterProxy(aiCenterID string) string {
	var proxy string
	for _, info := range setting.CenterInfos.Info {
		if info.CenterID == aiCenterID {
			proxy = info.StorageProxyServer
			break
		}
	}

	return proxy
}

func InitMultiNode() {
	if MultiNodeConfig == nil && setting.Grampus.MultiNode != "" {
		json.Unmarshal([]byte(setting.Grampus.MultiNode), &MultiNodeConfig)
	}

}
