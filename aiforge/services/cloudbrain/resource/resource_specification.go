package resource

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/util"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/convert"
	"code.gitea.io/gitea/modules/grampus"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/modelarts"
	"code.gitea.io/gitea/modules/setting"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/admin/operate_log"
)

func AddResourceSpecification(doerId int64, req models.ResourceSpecificationReq) error {
	if req.Status == 0 {
		req.Status = models.SpecNotVerified
	}
	spec := req.ToDTO()
	_, err := models.InsertResourceSpecification(spec)
	if err != nil {
		return err
	}
	return nil
}

func UpdateSpecUnitPrice(doerId int64, specId int64, unitPrice float64, comment string) *response.BizError {
	oldSpec, err := models.GetResourceSpecification(&models.ResourceSpecification{ID: specId})
	if err != nil {
		return response.NewBizError(err)
	}
	if oldSpec == nil {
		return response.SPECIFICATION_NOT_EXIST
	}
	err = models.UpdateSpecUnitPriceById(specId, unitPrice)
	if err != nil {
		return response.NewBizError(err)
	}

	if oldSpec.UnitPrice != unitPrice {
		AddSpecOperateLog(doerId, "edit", operate_log.NewLogValues().Add("unitPrice", unitPrice), operate_log.NewLogValues().Add("unitPrice", oldSpec.UnitPrice), specId, comment)
	}
	return nil
}

func SyncGrampusSpecs(doerId int64) error {
	r, err := grampus.GetResourceSpecs("")
	if err != nil {
		return err
	}
	log.Info("SyncGrampusSpecs result = %+v", r)
	specUpdateList := make([]models.ResourceSpecification, 0)
	specInsertList := make([]models.ResourceSpecification, 0)
	existIds := make([]int64, 0)
	for _, info := range r.Infos {
		for _, c := range info.Centers {
			for _, resourceSpec := range c.ResourceSpec {
				computeResource := models.ParseComputeResourceFormGrampus(info.SpecInfo.AccDeviceKind)
				if computeResource == "" {
					continue
				}
				accCardType := strings.ToUpper(info.SpecInfo.AccDeviceModel)
				memGiB, err := models.ParseMemSizeFromGrampus(info.SpecInfo.MemorySize)
				gpuMemGiB, err := models.ParseMemSizeFromGrampus(info.SpecInfo.AccDeviceMemory)
				if err != nil {
					log.Error("ParseMemSizeFromGrampus error. MemorySize=%s AccDeviceMemory=%s", info.SpecInfo.MemorySize, info.SpecInfo.AccDeviceMemory)
				}
				// get resource queue.if queue not exist,skip it
				r, err := models.GetResourceQueue(&models.ResourceQueue{
					Cluster:         models.C2NetCluster,
					AiCenterCode:    c.ID,
					ComputeResource: computeResource,
					AccCardType:     accCardType,
					QueueCode:       resourceSpec.ID,
				})
				if err != nil || r == nil {
					continue
				}

				//Determine if this specification already exists.if exist,update params
				//if not exist,insert a new record and status is SpecNotVerified
				oldSpec, err := models.GetResourceSpecification(&models.ResourceSpecification{
					QueueId:      r.ID,
					SourceSpecId: info.ID,
				})
				if err != nil {
					return err
				}

				if oldSpec == nil {
					specInsertList = append(specInsertList, models.ResourceSpecification{
						QueueId:         r.ID,
						SourceSpecId:    info.ID,
						AccCardsNum:     info.SpecInfo.AccDeviceNum,
						CpuCores:        info.SpecInfo.CpuCoreNum,
						MemGiB:          memGiB,
						GPUMemGiB:       gpuMemGiB,
						Status:          models.SpecNotVerified,
						IsAutomaticSync: true,
						IsAvailable:     true,
						CreatedBy:       doerId,
						UpdatedBy:       doerId,
					})
				} else {
					existIds = append(existIds, oldSpec.ID)
					specUpdateList = append(specUpdateList, models.ResourceSpecification{
						ID:          oldSpec.ID,
						AccCardsNum: info.SpecInfo.AccDeviceNum,
						CpuCores:    info.SpecInfo.CpuCoreNum,
						MemGiB:      memGiB,
						GPUMemGiB:   gpuMemGiB,
						IsAvailable: true,
						UpdatedBy:   doerId,
					})
				}
			}

		}
	}
	return models.SyncGrampusSpecs(specUpdateList, specInsertList, existIds)
}

func SyncGrampusSpecPools() error {
	r, err := grampus.GetResourceSpecPools("")
	if err != nil {
		return err
	}
	log.Info("SyncGrampusSpecPools result = %+v", r)

	for _, info := range r.Infos {
		queue_id, err := strconv.ParseInt(info.ID, 10, 64)
		if err != nil {
			fmt.Printf("strconv.ParseInt error: %v\n", err)
		}
		resourceQueue := models.ResourceQueue{
			AiCenterCode:  info.CenterId,
			QueueCode:     info.ID,
			AccCardType:   info.CardType,
			CardsTotalNum: info.CardNum,
		}

		err = models.UpdateResourceQueueCardNum(resourceQueue)
		if err != nil {
			fmt.Printf("UpdateResourceQueueById error: %v, queueId is: %v\n", err, queue_id)
		}
	}
	return nil
}

// GetResourceSpecificationList returns specification and queue
func GetResourceSpecificationList(opts models.SearchResourceSpecificationOptions) (*models.ResourceSpecAndQueueListRes, error) {
	n, r, err := models.SearchResourceSpecification(opts)
	if err != nil {
		return nil, err
	}
	return models.NewResourceSpecAndQueueListRes(n, r), nil
}

// GetAllDistinctResourceSpecification returns specification and queue after distinct
// totalSize is always 0 here
func GetAllDistinctResourceSpecification(opts models.SearchResourceSpecificationOptions) (*models.ResourceSpecAndQueueListRes, error) {
	opts.Page = 0
	opts.PageSize = 1000
	opts.OrderBy = models.SearchSpecOrder4Standard
	_, r, err := models.SearchResourceSpecification(opts)
	if err != nil {
		return nil, err
	}
	nr := distinctResourceSpecAndQueue(r)
	return models.NewResourceSpecAndQueueListRes(0, nr), nil
}

func GetAllResourceSpecification(opts models.SearchResourceSpecificationOptions) ([]*models.ResourceSpecInfo, error) {
	opts.Page = 0
	opts.PageSize = 1000
	opts.OrderBy = models.SearchSpecOrder4Standard
	_, r, err := models.SearchResourceSpecification(opts)
	if err != nil {
		return nil, err
	}

	exclusiveMap := models.FindQueuesExclusiveMap()
	res := make([]*models.ResourceSpecInfo, len(r))
	for i := 0; i < len(r); i++ {
		res[i] = r[i].ConvertToResourceSpecInfo()
		if _, exists := exclusiveMap[res[i].QueueId]; exists {
			res[i].IsQueueExclusive = true
		}
	}
	return res, nil
}

func distinctResourceSpecAndQueue(r []models.ResourceSpecAndQueue) []models.ResourceSpecAndQueue {
	specs := make([]models.ResourceSpecAndQueue, 0, len(r))
	sourceSpecIdMap := make(map[string]models.ResourceSpecAndQueue, 0)
	for i := 0; i < len(r); i++ {
		spec := r[i]
		if spec.SourceSpecId == "" {
			specs = append(specs, spec)
			continue
		}
		if _, has := sourceSpecIdMap[spec.SourceSpecId]; has {
			//prefer to use on-shelf spec
			if sourceSpecIdMap[spec.SourceSpecId].Status != spec.Status && spec.Status == models.SpecOnShelf {
				for k, v := range specs {
					if v.ResourceSpecification.ID == sourceSpecIdMap[spec.SourceSpecId].ResourceSpecification.ID {
						specs[k] = spec
					}
				}
			}
			continue
		}
		specs = append(specs, spec)
		sourceSpecIdMap[spec.SourceSpecId] = spec
	}
	return specs
}

func GetResourceSpecificationScenes(specId int64) ([]models.ResourceSceneBriefRes, error) {
	r, err := models.GetSpecScenes(specId)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func ResourceSpecOnShelf(doerId int64, id int64, unitPrice float64, comment string) *response.BizError {
	spec, err := models.GetResourceSpecification(&models.ResourceSpecification{ID: id})
	if err != nil {
		return response.NewBizError(err)
	}
	if spec == nil {
		return response.SPECIFICATION_NOT_EXIST
	}
	if q, err := models.GetResourceQueue(&models.ResourceQueue{ID: spec.QueueId}); err != nil || q == nil || !q.IsAvailable {
		return response.RESOURCE_QUEUE_NOT_AVAILABLE
	}
	if !spec.IsAvailable {
		return response.SPECIFICATION_NOT_AVAILABLE
	}
	err = models.ResourceSpecOnShelf(id, unitPrice)
	if err != nil {
		return response.NewBizError(err)
	}
	if spec.UnitPrice != unitPrice {
		AddSpecOperateLog(doerId, "on-shelf", operate_log.NewLogValues().Add("UnitPrice", unitPrice), operate_log.NewLogValues().Add("UnitPrice", spec.UnitPrice), id, comment)
	} else {
		AddSpecOperateLog(doerId, "on-shelf", nil, nil, id, comment)
	}
	return nil

}
func ResourceSpecOffShelf(doerId int64, id int64, comment string) *response.BizError {
	_, err := models.ResourceSpecOffShelf(id)
	if err != nil {
		return response.NewBizError(err)
	}
	AddSpecOperateLog(doerId, "off-shelf", nil, nil, id, comment)
	return nil

}

func AddSpecOperateLog(doerId int64, operateType string, newValue, oldValue *models.LogValues, specId int64, comment string) {
	var newString = ""
	var oldString = ""
	if newValue != nil {
		newString = newValue.JsonString()
	}
	if oldValue != nil {
		oldString = oldValue.JsonString()
	}

	spec, err := models.GetResourceSpecification(&models.ResourceSpecification{ID: specId})
	if err != nil {
		log.Error("GetResourceSpecification failed: %v", err)
	}

	ResourceQueue, err := models.GetResourceQueueByID(spec.QueueId)
	if err != nil {
		log.Error("GetResourceQueue error.%v", err)
	}

	operate_log.Log(models.AdminOperateLog{
		BizType:         "SpecOperate",
		OperateType:     operateType,
		OldValue:        oldString,
		NewValue:        newString,
		RelatedId:       fmt.Sprint(specId),
		CreatedBy:       doerId,
		Comment:         comment,
		ComputeResource: ResourceQueue.ComputeResource,
		AccCardType:     ResourceQueue.AccCardType,
		AccCardNum:      spec.AccCardsNum,
		AiCenterCode:    ResourceQueue.AiCenterCode,
		AiCenterName:    ResourceQueue.AiCenterName,
		ResourceQueueID: ResourceQueue.ID,
	})
}

func FindAvailableSpecsForNewRightWithRandomQueue(userId int64, opts models.FindSpecsOptions) ([]*models.Specification, error) {
	log.Info("opts.JobType=" + string(opts.JobType))
	specs, err := models.QueryDistinctResourceWithRandomQueue(models.FindAvailSpecsOptions{
		UserId:            userId,
		ComputeResource:   opts.ComputeResource,
		Cluster:           opts.Cluster,
		JobType:           opts.JobType,
		SpecId:            opts.SpecId,
		VisualizeRequired: opts.VisualizeRequired,
	})
	if err != nil {
		log.Error("FindAvailableSpecs error.%v", err)
		return nil, err
	}
	return specs, err
}

func FindAvailableSpecsForNewRight(userId int64, opts models.FindSpecsOptions) ([]*models.AggregateSpecification, error) {
	log.Info("opts.JobType=" + string(opts.JobType))
	specs, err := models.QueryAggregatedResource(models.FindAvailSpecsOptions{
		UserId:            userId,
		ComputeResource:   opts.ComputeResource,
		Cluster:           opts.Cluster,
		JobType:           opts.JobType,
		SpecId:            opts.SpecId,
		VisualizeRequired: opts.VisualizeRequired,
	})
	if err != nil {
		log.Error("FindAvailableSpecs error.%v", err)
		return nil, err
	}
	return specs, err
}

func FindAvailableSpecs(userId int64, opts models.FindSpecsOptions) ([]*models.Specification, error) {
	opts.SpecStatus = models.SpecOnShelf

	isUserSpecial := models.IsUserInExclusivePool(userId)
	if isUserSpecial {
		opts.SceneType = models.SceneTypeExclusive
	} else {
		opts.SceneType = models.SceneTypePublic
	}

	r, err := models.FindSpecs(opts)
	if err != nil {
		log.Error("FindAvailableSpecs error.%v", err)
		return nil, err
	}
	//filter exclusive specs
	specs := models.FilterExclusiveSpecs(r, userId)

	//filter special queue
	specs = models.HandleSpecialQueues(specs, userId, opts)

	//distinct by sourceSpecId
	specs = models.DistinctSpecs(specs)

	return specs, err
}

func FindAvailableSpecs4Show(userId int64, opts models.FindSpecsOptions) ([]*api.SpecificationShow, error) {
	//specs, err := FindAvailableSpecs(userId, opts)
	specs, err := FindAvailableSpecsForNewRightWithRandomQueue(userId, opts)
	if err != nil {
		return nil, err
	}
	result := make([]*api.SpecificationShow, len(specs))
	for i, v := range specs {
		result[i] = convert.ToSpecification(v)
	}
	return result, nil
}

func GetAndCheckSpec(userId int64, specId int64, opts models.FindSpecsOptions) (*models.Specification, error) {
	log.Info("FindAvailableSpecs try to get spec. opts=%+v specId=%d ", opts, specId)
	if specId == 0 {
		return nil, nil
	}
	opts.SpecId = specId
	opts.HasInternet = models.QueryAllSpecs
	//r, err := FindAvailableSpecs(userId, opts)
	r, err := FindAvailableSpecsForNewRightWithRandomQueue(userId, opts)
	log.Info("FindAvailableSpecsForNewRight result=%+v", r)
	if err != nil {
		return nil, err
	}
	if r == nil || len(r) == 0 {
		return nil, nil
	}
	log.Info("FindAvailableSpecs final result r[0]=%+v", r[0])
	return r[0], nil
}

func InsertCloudbrainSpec(cloudbrainId int64, s *models.Specification) error {
	c := models.CloudbrainSpec{
		CloudbrainID:    cloudbrainId,
		SpecId:          s.ID,
		SourceSpecId:    s.SourceSpecId,
		AccCardsNum:     s.AccCardsNum,
		AccCardType:     s.AccCardType,
		CpuCores:        s.CpuCores,
		MemGiB:          s.MemGiB,
		GPUMemGiB:       s.GPUMemGiB,
		ShareMemGiB:     s.ShareMemGiB,
		ComputeResource: s.ComputeResource,
		UnitPrice:       s.UnitPrice,
		QueueId:         s.QueueId,
		QueueCode:       s.QueueCode,
		Cluster:         s.Cluster,
		AiCenterCode:    s.AiCenterCode,
		AiCenterName:    s.AiCenterName,
		IsExclusive:     s.IsSpecExclusive(),
		ExclusiveOrg:    s.ExclusiveOrg,
	}
	_, err := models.InsertCloudbrainSpec(c)
	if err != nil {
		log.Error("InsertCloudbrainSpec error.CloudbrainSpec=%v. err=%v", c, err)
		return err
	}
	return nil
}

func GetCloudbrainSpec(cloudbrainId int64) (*models.Specification, error) {
	c, err := models.GetCloudbrainSpecByID(cloudbrainId)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, nil
	}
	return c.ConvertToSpecification(), nil
}

func RefreshHistorySpec(scopeAll bool, ids []int64) (int64, int64, error) {
	var success int64
	var total int64

	if !scopeAll {
		if ids == nil || len(ids) == 0 {
			return 0, 0, nil
		}
		total = int64(len(ids))
		tasks, err := models.GetCloudbrainWithDeletedByIDs(ids)
		if err != nil {
			return total, 0, err
		}
		for _, task := range tasks {
			err = RefreshOneHistorySpec(task)
			if err != nil {
				log.Error("RefreshOneHistorySpec error.%v", err)
				continue
			}
			success++
		}

	} else {
		page := 1
		pageSize := 100
		n, err := models.CountNoSpecHistoricTask()
		if err != nil {
			log.Error("FindNoSpecHistoricTask CountNoSpecHistoricTask error. e=%v", err)
			return 0, 0, err
		}
		total = n
		for i := 0; i < 500; i++ {
			list, err := models.FindCloudbrainTask(page, pageSize)
			page++
			if err != nil {
				log.Error("FindCloudbrainTask error.page=%d pageSize=%d  e=%v", page, pageSize, err)
				return total, success, err
			}
			if len(list) == 0 {
				log.Info("RefreshHistorySpec. list is empty")
				break
			}
			for _, task := range list {
				s, err := GetCloudbrainSpec(task.ID)
				if err != nil {
					log.Error("RefreshHistorySpec GetCloudbrainSpec error.%v", err)
					continue
				}
				if s != nil {
					continue
				}
				err = RefreshOneHistorySpec(task)
				if err != nil {
					log.Error("RefreshOneHistorySpec error.%v", err)
					continue
				}
				success++
			}
			if len(list) < pageSize {
				log.Info("RefreshHistorySpec. list < pageSize")
				break
			}
		}
	}
	return total, success, nil

}

func RefreshOneHistorySpec(task *models.Cloudbrain) error {
	var spec *models.Specification
	var err error
	switch task.Type {
	case models.TypeCloudBrainOne:
		spec, err = getCloudbrainOneSpec(task)
	case models.TypeCloudBrainTwo:
		spec, err = getCloudbrainTwoSpec(task)
	case models.TypeC2Net:
		spec, err = getGrampusSpec(task)
	}
	if err != nil {
		log.Error("find spec error,task.ID=%d err=%v", task.ID, err)
		return err
	}
	if spec == nil {
		log.Error("find spec failed,task.ID=%d", task.ID)
		return errors.New("find spec failed")
	}
	return InsertCloudbrainSpec(task.ID, spec)
}

func getCloudbrainOneSpec(task *models.Cloudbrain) (*models.Specification, error) {
	if task.GpuQueue == "" {
		log.Info("gpu queue is empty.task.ID = %d", task.ID)
		return nil, nil
	}
	//find from config
	spec, err := findCloudbrainOneSpecFromConfig(task)
	if err != nil {
		log.Error("getCloudbrainOneSpec findCloudbrainOneSpecFromConfig error.%v", err)
		return nil, err
	}
	if spec != nil {
		return spec, nil
	}
	//find from remote
	return findCloudbrainOneSpecFromRemote(task)

}

func findCloudbrainOneSpecFromRemote(task *models.Cloudbrain) (*models.Specification, error) {
	time.Sleep(200 * time.Millisecond)
	log.Info("start findCloudbrainOneSpecFromRemote")
	result, err := cloudbrain.GetJob(task.JobID)
	if err != nil {
		log.Error("getCloudbrainOneSpec error. %v", err)
		return nil, err
	}

	if result == nil {
		log.Info("findCloudbrainOneSpecFromRemote failed,result is empty.task.ID=%d", task.ID)
		return nil, nil
	}
	jobRes, _ := models.ConvertToJobResultPayload(result.Payload)
	memSize, _ := models.ParseMemSizeFromGrampus(jobRes.Resource.Memory)
	if task.ComputeResource == "CPU/GPU" {
		task.ComputeResource = models.GPU
	}
	var shmMB float32
	if jobRes.Config.TaskRoles != nil && len(jobRes.Config.TaskRoles) > 0 {
		shmMB = float32(jobRes.Config.TaskRoles[0].ShmMB) / 1024
		if jobRes.Config.TaskRoles[0].ShmMB == 103600 {
			shmMB = 100
		} else if jobRes.Config.TaskRoles[0].ShmMB == 51800 {
			shmMB = 50
		}
	}
	opt := models.FindSpecsOptions{
		ComputeResource: task.ComputeResource,
		Cluster:         models.OpenICluster,
		AiCenterCode:    models.AICenterOfCloudBrainOne,
		QueueCode:       task.GpuQueue,
		AccCardsNum:     jobRes.Resource.NvidiaComGpu,
		UseAccCardsNum:  true,
		CpuCores:        jobRes.Resource.CPU,
		UseCpuCores:     true,
		MemGiB:          memSize,
		UseMemGiB:       memSize > 0,
		ShareMemGiB:     shmMB,
		UseShareMemGiB:  shmMB > 0,
		RequestAll:      true,
	}
	specs, err := models.FindSpecs(opt)
	if err != nil {
		log.Error("getCloudbrainOneSpec from remote error,%v", err)
		return nil, err
	}
	if len(specs) == 1 {
		return specs[0], nil
	}
	if len(specs) == 0 {
		s, err := InitQueueAndSpec(opt, "云脑一", "处理历史云脑任务时自动添加")
		if err != nil {
			log.Error("getCloudbrainOneSpec InitQueueAndSpec error.err=%v", err)
			return nil, nil
		}
		return s, nil
	}
	log.Error("Too many results matched.size=%d  opt=%+v", len(specs), opt)
	return nil, nil
}

func findCloudbrainOneSpecFromConfig(task *models.Cloudbrain) (*models.Specification, error) {
	//find from config
	var specConfig *models.ResourceSpec
	hasSpec := false
	if task.JobType == string(models.JobTypeTrain) {
		if cloudbrain.TrainResourceSpecs == nil {
			json.Unmarshal([]byte(setting.TrainResourceSpecs), &cloudbrain.TrainResourceSpecs)
		}
		for _, tmp := range cloudbrain.TrainResourceSpecs.ResourceSpec {
			if tmp.Id == task.ResourceSpecId {
				hasSpec = true
				specConfig = tmp
				break
			}
		}
	} else if task.JobType == string(models.JobTypeInference) {
		if cloudbrain.InferenceResourceSpecs == nil {
			json.Unmarshal([]byte(setting.InferenceResourceSpecs), &cloudbrain.InferenceResourceSpecs)
		}
		for _, tmp := range cloudbrain.InferenceResourceSpecs.ResourceSpec {
			if tmp.Id == task.ResourceSpecId {
				hasSpec = true
				specConfig = tmp
				break
			}
		}
	} else {
		if cloudbrain.ResourceSpecs == nil {
			json.Unmarshal([]byte(setting.ResourceSpecs), &cloudbrain.ResourceSpecs)
		}
		for _, tmp := range cloudbrain.ResourceSpecs.ResourceSpec {
			if tmp.Id == task.ResourceSpecId {
				hasSpec = true
				specConfig = tmp
				break

			}
		}
	}
	if !hasSpec && cloudbrain.SpecialPools != nil {

		for _, specialPool := range cloudbrain.SpecialPools.Pools {

			if specialPool.ResourceSpec != nil {

				for _, spec := range specialPool.ResourceSpec {
					if task.ResourceSpecId == spec.Id {
						hasSpec = true
						specConfig = spec
						break
					}
				}
			}
		}
	}
	if specConfig == nil {
		log.Error("getCloudbrainOneSpec from config failed,task.ResourceSpecId=%d", task.ResourceSpecId)
		return nil, nil
	}
	if task.ComputeResource == "CPU/GPU" {
		task.ComputeResource = models.GPU
	}

	shareMemMiB := float32(specConfig.ShareMemMiB) / 1024
	if specConfig.ShareMemMiB == 103600 {
		shareMemMiB = 100
	} else if specConfig.ShareMemMiB == 51800 {
		shareMemMiB = 50
	}
	opt := models.FindSpecsOptions{
		JobType:         models.JobType(task.JobType),
		ComputeResource: task.ComputeResource,
		Cluster:         models.OpenICluster,
		AiCenterCode:    models.AICenterOfCloudBrainOne,
		QueueCode:       task.GpuQueue,
		AccCardsNum:     specConfig.GpuNum,
		UseAccCardsNum:  true,
		CpuCores:        specConfig.CpuNum,
		UseCpuCores:     true,
		MemGiB:          float32(specConfig.MemMiB) / 1024,
		UseMemGiB:       true,
		ShareMemGiB:     shareMemMiB,
		UseShareMemGiB:  true,
		RequestAll:      true,
	}
	specs, err := models.FindSpecs(opt)
	if err != nil {
		log.Error("getCloudbrainOneSpec from config error,%v", err)
		return nil, err
	}
	if len(specs) > 1 {
		log.Error("Too many results matched.size=%d  opt=%+v", len(specs), opt)
		return nil, nil
	}
	if len(specs) == 0 {
		s, err := InitQueueAndSpec(opt, "云脑一", "处理历史云脑任务时自动添加")
		if err != nil {
			log.Error("getCloudbrainOneSpec InitQueueAndSpec error.err=%v", err)
			return nil, nil
		}
		return s, nil
	}
	return specs[0], nil
}

func getCloudbrainTwoSpec(task *models.Cloudbrain) (*models.Specification, error) {
	specMap, err := models.GetCloudbrainTwoSpecs()
	if err != nil {
		log.Error("InitCloudbrainTwoSpecs err.%v", err)
		return nil, err
	}
	if task.FlavorCode != "" {
		return specMap[task.FlavorCode], nil
	}
	time.Sleep(200 * time.Millisecond)
	log.Info("start getCloudbrainTwoSpec FromRemote")
	if task.JobType == string(models.JobTypeDebug) {
		result, err := modelarts.GetNotebook2(task.JobID)
		if err != nil {
			log.Error("getCloudbrainTwoSpec GetNotebook2 error.%v", err)
			return nil, err
		}
		if result != nil {
			return specMap[result.Flavor], nil
		}
	} else if task.JobType == string(models.JobTypeTrain) || task.JobType == string(models.JobTypeInference) {
		result, err := modelarts.GetTrainJob(task.JobID, strconv.FormatInt(task.VersionID, 10))
		if err != nil {
			log.Error("getCloudbrainTwoSpec GetTrainJob error:%v", task.JobName, err)
			return nil, err
		}
		if result != nil {
			return specMap[result.Flavor.Code], nil
		}
	}
	return nil, nil
}

func getGrampusSpec(task *models.Cloudbrain) (*models.Specification, error) {
	specMap, err := models.GetGrampusSpecs()
	if err != nil {
		log.Error("GetGrampusSpecs err.%v", err)
		return nil, err
	}
	if task.AiCenter != "" {
		c := strings.Split(task.AiCenter, "+")
		spec := specMap[task.FlavorCode+"_"+c[0]]
		if spec != nil {
			return spec, nil
		}
	}
	return specMap[task.FlavorCode], nil
}

func InitQueueAndSpec(opt models.FindSpecsOptions, aiCenterName string, remark string) (*models.Specification, error) {
	return models.InitQueueAndSpec(models.ResourceQueue{
		QueueCode:       opt.QueueCode,
		Cluster:         opt.Cluster,
		AiCenterCode:    opt.AiCenterCode,
		AiCenterName:    aiCenterName,
		ComputeResource: opt.ComputeResource,
		AccCardType:     models.GetCloudbrainOneAccCardType(opt.QueueCode),
		Remark:          remark,
	}, models.ResourceSpecification{
		AccCardsNum: opt.AccCardsNum,
		CpuCores:    opt.CpuCores,
		MemGiB:      opt.MemGiB,
		GPUMemGiB:   opt.GPUMemGiB,
		ShareMemGiB: opt.ShareMemGiB,
		Status:      models.SpecOffShelf,
		IsAvailable: true,
	})
}

func GetResourceListPaging(opts models.GetResourceListOpts) ([]*models.ResourceInfo4CardRequest, int64, error) {
	return models.GetResourceListPaging(opts)
}

// GetResourceListPagingV2 处理资源的list
func GetResourceListPagingV2(opts models.GetResourceListOpts) ([]*models.ResourceInfo4CardRequest, int64, error) {
	// 获取所有的资源ids
	specIds, err := models.GetAllSpecIds()
	if err != nil {
		return []*models.ResourceInfo4CardRequest{}, 0, err
	}

	var (
		rsp   = make([]*models.ResourceInfo4CardRequest, 0)
		total int64
	)

	// 分批函数处理，防止in sql过多，size = 【200】
	for _, chunkSpecIds := range util.ChunkInt64(specIds, util.PUBULIC_CHUNK_SIZE) {
		rtr, gTotal, mErr := models.GetResourceListPagingV2(opts, chunkSpecIds)
		if mErr != nil {
			return []*models.ResourceInfo4CardRequest{}, 0, mErr
		}

		// 总数叠加处理
		total += gTotal

		rsp = append(rsp, rtr...)

	}
	return rsp, total, nil
}

func GetAccCardList() ([]models.AccCardInfo, error) {
	return models.GetAccCardList()
}

// GetAccCardList 根据角色查询
func GetAccCardListV2(specIds []int64) ([]models.AccCardInfo, error) {
	return models.GetAccCardListV2(specIds)
}

func GetAvailableAICenter() ([]*models.ResourceAiCenterRes, error) {
	return models.GetAvailableResourceAiCenters()
}
func GetAvailableAICenterV2(specIds []int64) ([]*models.ResourceAiCenterRes, error) {
	return models.GetAvailableResourceAiCentersV2(specIds)
}
