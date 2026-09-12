package task

import (
	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/cloudbrain/cloudbrainTask"
	"code.gitea.io/gitea/services/reward/point/account"
	"code.gitea.io/gitea/services/role"
)

func GetAITaskCreationRepoInfo(req entity.GetAITaskCreationInfoReq) (*entity.CreationRequiredRepoInfo, *response.BizError) {
	result := &entity.CreationRequiredRepoInfo{}
	if req.GitRepo != nil {
		if branches, _, err := req.GitRepo.GetBranches(0, 0); err == nil {
			result.Branches = branches
			if len(req.Repo.DefaultBranch) > 0 && req.GitRepo.IsBranchExist(req.Repo.DefaultBranch) {
				result.DefaultBranch = req.Repo.DefaultBranch
			} else if len(branches) > 0 {
				result.DefaultBranch = branches[0]
			}
		}
	}
	return result, nil

}

func GetAITaskCreationInfo(req entity.GetAITaskCreationInfoReq) (*entity.CreationRequiredInfo, *response.BizError) {
	result := &entity.CreationRequiredInfo{}

	//查询排队信息
	waitCount := cloudbrain.GetWaitingCloudbrainCount(req.ClusterType.GetCloudbrainType(), req.ComputeSource.GetCloudbrainFormat(), req.JobType)
	result.WaitCount = waitCount
	//查询是否有正在运行的任务
	notStopTaskCount := 0
	if req.IsOnlineType {
		notStopTaskCount, _ = cloudbrainTask.GetNotFinalStatusTaskCount(req.User.ID, string(models.JobTypeOnlineInference))
	} else {
		notStopTaskCount, _ = cloudbrainTask.GetNotFinalStatusTaskCount(req.User.ID, string(req.JobType))
	}
	result.NotStopTaskCount = notStopTaskCount
	limitNum := GetUserMultiLimitNum(req.User.ID, req.JobType)
	result.CanCreateMore = notStopTaskCount < limitNum
	result.LimitCount = limitNum
	result.NoteBookCanCreateMore = true
	if req.JobType == models.JobTypeDebug {
		count, err := models.GetNotebooksCountByUser(req.User.ID)
		if err != nil {
			log.Warn("can not get user notebook count", err)
		}
		if count >= int64(setting.NotebookStrategy.MaxNumberPerUser) {
			result.NoteBookCanCreateMore = false
		}
	}
	result.NotebookLimitCount = setting.NotebookStrategy.MaxNumberPerUser
	//获取代码分支
	if req.GitRepo != nil {
		if branches, _, err := req.GitRepo.GetBranches(0, 0); err == nil {
			result.Branches = branches
			if len(req.Repo.DefaultBranch) > 0 && req.GitRepo.IsBranchExist(req.Repo.DefaultBranch) {
				result.DefaultBranch = req.Repo.DefaultBranch
			} else if len(branches) > 0 {
				result.DefaultBranch = branches[0]
			}
		}
	}

	//查询积分余额
	if a, err := account.GetAccount(req.User.ID); err == nil {
		result.PointAccount = entity.ParsePointAccountInfo(a)
	}
	//积分开关
	result.PaySwitch = setting.CloudBrainPaySwitch

	t, err := GetAITaskTemplate(req.JobType, req.ClusterType)

	if err != nil {
		log.Error("param error")
		return nil, err
	}

	//生成任务名称
	result.DisplayJobName = t.GetDisplayJobName(req.User.Name)
	specsMap := make(map[string][]*structs.SpecificationShow, 0)
	log.Info("DisplayJobName=" + result.DisplayJobName)
	//查询所有资源规格
	if specs, err := t.GetSpecs(entity.GetSpecOpts{
		UserId:            req.User.ID,
		JobType:           req.JobType,
		ComputeSource:     *req.ComputeSource,
		VisualizeRequired: req.VisualizeRequired,
	}); err == nil {
		log.Info("find spec success.")
		specsMap["all"] = specs
		specsMap["has_internet"] = filterHasInternetSpecs(specs)
		specsMap["no_internet"] = filterNoInternetSpecs(specs)
	}

	result.Specs = specsMap
	// 查询镜像列表
	if images, canUseAll, err := t.GetImages(*req.ComputeSource, ""); err == nil {
		result.Images = images
		result.CanUseAllImages = canUseAll
	}

	c := t.GetConfig(entity.AITaskConfigKey{ComputeSource: req.ComputeSource.GetCloudbrainFormat()})
	result.Config = entity.AITaskCreationConfig{
		//DatasetMaxSize: setting.DebugAttachSize * 1000 * 1000 * 1000,
		DatasetMaxSize: c.DatasetsLimitSizeGB * 1024 * 1024 * 1024,
		DatasetsMaxNum: c.DatasetsMaxNum,
		ModelMaxSize:   c.ModelLimitSizeGB * 1024 * 1024 * 1024,
		ModelMaxNum:    c.ModelMaxNum,
	}
	//查询可用节点数
	if workerNums, err := t.GetAllowedWorkerNum(req.User.ID, req.ComputeSource); err == nil {
		result.AllowedWorkerNum = workerNums
	} else {
		result.AllowedWorkerNum = []int{1}
	}
	//result.IsSubscriber = role.UserHasRole(req.User.ID, models.Subscriber)
	if req.JobType == models.JobTypeDebug {
		result.IsSubscriber = role.UserHasOper(req.User.ID, role.ROLE_OPER_DEBUG_TIME)
	} else if req.JobType == models.JobTypeOnlineInference {
		result.IsSubscriber = role.UserHasOper(req.User.ID, role.ROLE_OPER_ONLINE_INFER_PATH)
	} else {
		result.IsSubscriber = false
	}

	// 查询容器存储配额
	result.CodeSizeLimit, result.OutputSizeLimit = role.GetUserContainerStorageLimits(req.User.ID)

	return result, nil
}

func filterHasInternetSpecs(allSpecs []*structs.SpecificationShow) []*structs.SpecificationShow {
	hasInternetSpecs := make([]*structs.SpecificationShow, 0)
	for i := 0; i < len(allSpecs); i++ {
		if allSpecs[i].NetworkCapableQueuesExist {
			hasInternetSpecs = append(hasInternetSpecs, allSpecs[i])
		}
	}
	return hasInternetSpecs
}

func filterNoInternetSpecs(allSpecs []*structs.SpecificationShow) []*structs.SpecificationShow {
	noInternetSpecs := make([]*structs.SpecificationShow, 0)
	for i := 0; i < len(allSpecs); i++ {
		if allSpecs[i].NetworkIncapableQueuesExist {
			noInternetSpecs = append(noInternetSpecs, allSpecs[i])
		}
	}
	return noInternetSpecs
}

func GetAvailableImageInfoBySpec(req entity.GetAITaskCreationImageInfoReq) (*entity.ImageRequiredInfo, *response.BizError) {
	result := &entity.ImageRequiredInfo{}
	t, err := GetAITaskTemplate(req.JobType, req.ClusterType)

	if err != nil {
		log.Error("param error")
		return nil, err
	}
	queues := req.Spec.GetAvailableQueues(models.GetAvailableCenterIdOpts{
		UserId:            req.UserID,
		JobType:           req.JobType,
		VisualizeRequired: req.VisualizeRequired,
		HasInternet:       req.HasInternet,
	})
	//load from db
	reImages := make([]entity.ClusterImage, 0, 10)
	result.Images = reImages
	if images, canUseAll, err := t.GetImages(*req.ComputeSource, req.Spec.AccCardType, queues...); err == nil {
		result.Images = append(result.Images, images...)
		result.CanUseAllImages = canUseAll
	}
	return result, nil

}
