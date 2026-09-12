package task

import (
	"encoding/json"
	"fmt"
	"path"
	"regexp"
	"strings"

	"code.gitea.io/gitea/services/role"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/otel"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/modules/util"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/context"
	"code.gitea.io/gitea/services/cloudbrain/cloudbrainTask"
	"code.gitea.io/gitea/services/cloudbrain/modelmanage"
	"code.gitea.io/gitea/services/cloudbrain/resource"
	"code.gitea.io/gitea/services/reward/point/account"
)

type CreationHandler interface {
	BuildRequest4Restart(ctx *context.CreationContext) *response.BizError
	CheckParamFormat(ctx *context.CreationContext) *response.BizError
	CheckPrivilege4Continue(ctx *context.CreationContext) *response.BizError
	HandleReqParameters(ctx *context.CreationContext) *response.BizError
	CheckMultiRequest(ctx *context.CreationContext) *response.BizError
	CheckDisplayJobName(ctx *context.CreationContext) *response.BizError
	LoadSpec(ctx *context.CreationContext) *response.BizError
	CheckPointBalance(ctx *context.CreationContext) *response.BizError
	CheckDatasets(ctx *context.CreationContext) *response.BizError
	CheckBranchExists(ctx *context.CreationContext) *response.BizError
	CheckBootFile(ctx *context.CreationContext) *response.BizError
	CheckSourceTaskIsCleared(ctx *context.CreationContext) *response.BizError
	BuildContainerData(ctx *context.CreationContext) *response.BizError
	InsertCloudbrainRecord4Async(ctx *context.CreationContext) *response.BizError
	CallCreationAPI(ctx *context.CreationContext) *response.BizError
	AfterCallCreationAPI4Async(ctx *context.CreationContext) *response.BizError
	AfterCallCreationAPI4Sync(ctx *context.CreationContext) *response.BizError
	CreateCloudbrainRecord4Restart(ctx *context.CreationContext) *response.BizError
	CallRestartAPI(ctx *context.CreationContext) *response.BizError
	NotifyCreation(ctx *context.CreationContext) *response.BizError
	HandleErr4Async(ctx *context.CreationContext) *response.BizError
	CheckNotebookCount(ctx *context.CreationContext) *response.BizError
	GetAvailableQueues(ctx *context.CreationContext) *response.BizError
	CheckImageAvailable(ctx *context.CreationContext) *response.BizError
}

// DefaultCreationHandler  CreationHandler的默认实现，公共逻辑可以在此结构体中实现
type DefaultCreationHandler struct {
}

func (g DefaultCreationHandler) BuildContainerData(ctx *context.CreationContext) *response.BizError {
	return nil
}

func (DefaultCreationHandler) BuildRequest4Restart(ctx *context.CreationContext) *response.BizError {
	task := ctx.SourceCloudbrain
	if task == nil {
		log.Error("ctx.Cloudbrain not exists")
		return response.AI_TASK_NOT_EXISTS
	}
	if !task.IsTerminal() {
		log.Error("ctx.Cloudbrain not terminal")
		return response.AI_TASK_NOT_FINISHED
	}
	//再次调试前尝试修正一下资源规格（主要是为了处理历史数据）
	correctAITaskSpec(task)
	oldSpec, err := resource.GetCloudbrainSpec(task.ID)
	if err != nil {
		log.Error("GetCloudbrainSpec err. %v", err)
		return response.SPEC_NOT_AVAILABLE
	}
	imageUrl := task.Image
	imageName := task.Image
	imageId := task.ImageID

	if task.Image == "" {
		imageUrl = task.EngineName
		imageName = task.EngineName
	}
	if imageId == "" && task.EngineID > 0 {
		imageId = fmt.Sprint(task.EngineID)
	}
	ctx.Request = &entity.CreateReq{
		JobType:               models.JobType(task.JobType),
		DisplayJobName:        task.DisplayJobName,
		JobName:               task.JobName,
		SpecId:                oldSpec.ID,
		ComputeSourceStr:      task.GetStandardComputeSource(),
		Cluster:               entity.GetClusterTypeFromCloudbrainType(task.Type),
		WorkServerNumber:      task.WorkServerNumber,
		BranchName:            task.BranchName,
		ImageUrl:              imageUrl,
		ImageID:               imageId,
		ImageName:             imageName,
		PretrainModelName:     task.ModelName,
		Description:           task.Description,
		LabelName:             task.LabelName,
		DatasetUUIDStr:        task.Uuid,
		Params:                task.Parameters,
		BootFile:              task.BootFile,
		PretrainModelId:       task.ModelId,
		ReqCommitID:           task.CommitID,
		IsFileNoteBookRequest: task.BootFile != "",
		IsRestartRequest:      true,
		DatasetNames:          task.DatasetName,
		DatasetAlias:          task.DatasetAlias,
		HasInternet:           models.SpecInternetQuery(task.HasInternet),
		TimeLimit:             task.TimeLimit,
		EndPoint:              task.EndPoint,
		Port:                  task.Port,
	}
	log.Info("BuildRequest4Restart success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return nil
}

func (g DefaultCreationHandler) CheckDatasets(ctx *context.CreationContext) *response.BizError {
	log.Info("Start to CheckDatasets.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	datasetUUIDStr := ctx.Request.DatasetUUIDStr
	if datasetUUIDStr == "" {
		return nil
	}
	//check datasets num
	uuids := strings.Split(datasetUUIDStr, ";")
	if ctx.Config.DatasetsMaxNum > 0 && len(uuids) > ctx.Config.DatasetsMaxNum {
		log.Error("the dataset count(%d) exceed the limit", len(uuids))
		return response.DATASET_NUMBER_OVER_LIMIT
	}
	datasetList, err := models.GetDatasetRegistryListByIDs(uuids)
	if err != nil {
		log.Error("GetDatasetInfo failed: %v", err)
		return response.SYSTEM_ERROR
	}

	if len(datasetList) < len(uuids) {
		log.Info("CheckDataset hasDatasetDeleted.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
		return response.DATASET_NOT_EXISTS
	}

	for _, uuid := range uuids {
		if !canReadSubject(ctx.User, uuid, models.DatasetSubject) {
			log.Warn("CheckDataset can not read.uuid=%s ", uuid)
			return response.DATASET_CANNOT_READ
		}
	}

	//check datasets size
	var attachSize int64
	for _, infos := range datasetList {
		attachSize += infos.Size
	}
	limitSizeGB := ctx.Config.DatasetsLimitSizeGB
	if limitSizeGB > 0 && attachSize > int64(limitSizeGB*1024*1024*1024) {
		log.Error("The DatasetSize exceeds the limit (%dGB)", limitSizeGB) // GB
		return response.DATASET_SIZE_OVER_LIMIT.WithParams(limitSizeGB)
	}

	var datasetNames, datasetAlias string
	for i := 0; i < len(uuids); i++ {
		for _, dataset := range datasetList {
			if dataset.ID == uuids[i] {
				datasetNames += dataset.Name + ";"
				datasetAlias += dataset.DisplayName() + ";"
			}
		}
	}
	ctx.Request.DatasetNames = strings.TrimSuffix(datasetNames, ";")
	ctx.Request.DatasetAlias = strings.TrimSuffix(datasetAlias, ";")
	log.Info("CheckDatasets success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return nil
}

func canReadSubject(user *models.User, subjectID string, subjectType models.SubjectType) bool {
	subjectContext, err := models.GetSubjectContext(subjectID, subjectType)
	if err != nil {
		return false
	}
	if subjectContext.OwnerID == user.ID {
		subjectContext.Owner = user
	} else {
		err = subjectContext.GetOwner()
		if err != nil {
			return false
		}
	}

	permission, err := models.GetUserSubjectPermission(subjectContext, user)
	if err != nil {
		return false
	}
	return permission.AccessMode >= models.AccessModeRead

}

func (g DefaultCreationHandler) CheckModels(ctx *context.CreationContext) *response.BizError {
	log.Info("Start to CheckModels.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	modelIdStr := ctx.Request.PretrainModelId
	if modelIdStr == "" {
		return nil
	}
	//check model num
	uuids := strings.Split(modelIdStr, ";")
	if ctx.Config.ModelMaxNum > 0 && len(uuids) > ctx.Config.ModelMaxNum {
		log.Error("the dataset count(%d) exceed the limit", len(uuids))
		return response.MODEL_NUMBER_OVER_LIMIT
	}

	modelInfoMaps, err := models.QueryModelMapsByIds(uuids)
	if err != nil {
		log.Error("QueryModelsByIds failed: %v", err)
		return response.SYSTEM_ERROR
	}

	if len(modelInfoMaps) < len(uuids) {
		log.Info("CheckModels has model deleted.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
		return response.MODEL_NOT_EXISTS
	}

	for _, uuid := range uuids {
		if !canReadSubject(ctx.User, uuid, models.AimodelSubject) {
			log.Warn("CheckModel can not read.uuid=%s ", uuid)
			return response.MODEL_CANNOT_READ
		}
	}

	//check datasets size
	var attachSize int64
	for _, infos := range modelInfoMaps {
		attachSize += infos.Size
	}
	limitSizeGB := ctx.Config.ModelLimitSizeGB
	if limitSizeGB > 0 && attachSize > int64(limitSizeGB*1024*1024*1024) {
		log.Error("The model size exceeds the limit (%dGB)", limitSizeGB) // GB
		return response.MODEL_SIZE_OVER_LIMIT.WithParams(limitSizeGB)
	}

	var modelNames string
	for i := 0; i < len(uuids); i++ {
		m := modelInfoMaps[uuids[i]]
		modelNames += m.Name + ";"
	}
	ctx.Request.ModelNames = strings.TrimSuffix(modelNames, ";")
	log.Info("CheckModels success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return nil
}

func (DefaultCreationHandler) CheckBranchExists(ctx *context.CreationContext) *response.BizError {
	log.Info("Start to CheckBranchExists.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	if ctx.GitRepo == nil || ctx.Request.BranchName == "" {
		return nil
	}
	log.Info("ctx.GitRepo=" + ctx.GitRepo.Path + " ctx.Request.BranchName=" + ctx.Request.BranchName)
	if !ctx.GitRepo.IsBranchExist(ctx.Request.BranchName) {
		return response.BRANCH_NOT_EXISTS
	}
	log.Info("CheckBranchExists success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return nil
}

func (DefaultCreationHandler) CheckSourceTaskIsCleared(ctx *context.CreationContext) *response.BizError {
	log.Info("Start to CheckSourceTaskIsCleared.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	task := ctx.SourceCloudbrain
	if task == nil {
		return nil
	}
	if task.Cleared {
		return response.RESULT_CLEARD
	}
	log.Info("CheckSourceTaskIsCleared success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return nil
}

func hasModelFileDeleted(modelId, pretrainModelCkptName string) bool {
	if modelId == "" {
		return false
	}
	CkptNames := strings.Split(pretrainModelCkptName, ";")
	for _, ckptName := range CkptNames {
		if !modelmanage.HasModelFileByModelId(modelId, ckptName) {
			return true
		}
	}
	return false
}

func hasModelNumOverLimit(pretrainModelCkptName string) bool {
	CkptNames := strings.Split(pretrainModelCkptName, ";")
	if len(CkptNames) > 30 {
		return true
	}
	return false
}

var jobNamePattern = regexp.MustCompile(`^[a-z0-9][a-z0-9-_]{1,34}[a-z0-9-]$`)
var dockerImageRegex = regexp.MustCompile(`^(?:[a-zA-Z0-9.-]+(?::[0-9]+)?/)?[a-zA-Z0-9._-]+(?:/[a-zA-Z0-9._-]+)*(?::[a-zA-Z0-9._-]+)?(?:@sha256:[a-fA-F0-9]{64})?$`)

func (DefaultCreationHandler) CheckParamFormat(ctx *context.CreationContext) *response.BizError {
	req := ctx.Request
	log.Info("Start to CheckParam.displayJobName=%s jobType=%s cluster=%s", req.DisplayJobName, req.JobType, req.Cluster)

	c := models.GetComputeSourceInstance(req.ComputeSourceStr)
	if c == nil {
		log.Error("ComputeSourceStr invalid")
		return response.PARAM_ERROR
	}
	ctx.Request.ComputeSource = c

	if !jobNamePattern.MatchString(req.DisplayJobName) {
		log.Error("DisplayJobName invalid")
		return response.PARAM_ERROR
	}
	if req.ImageUrl != "" {
		if !dockerImageRegex.MatchString(req.ImageUrl) {
			log.Error("ImageUrl invalid")
			return response.PARAM_ERROR
		}
	}
	if len(req.Description) > 1000 {
		log.Error("Description too long,displayJobName=%s", req.DisplayJobName)
		return response.PARAM_ERROR
	}
	if req.Description != "" {
		req.Description = util.StripHTML(req.Description)
	}
	if req.ImageName != "" {
		req.ImageName = util.StripHTML(req.ImageName)
	}

	ctx.Request.BootFile = strings.TrimSpace(ctx.Request.BootFile)

	if ctx.Request.SourceCloudbrainId > 0 {
		sourceTask, _ := models.GetCloudbrainByCloudbrainID(ctx.Request.SourceCloudbrainId)
		ctx.SourceCloudbrain = sourceTask
	}

	log.Info("CheckParam success.displayJobName=%s jobType=%s cluster=%s", req.DisplayJobName, req.JobType, req.Cluster)
	return nil
}

func (DefaultCreationHandler) CheckPrivilege4Continue(ctx *context.CreationContext) *response.BizError {
	req := ctx.Request
	//继续训练或者创建新版本时需要校验对旧云脑任务的权限
	if !ctx.Request.IsContinueRequest {
		return nil
	}
	log.Info("Start to CheckPrivilege4Continue.displayJobName=%s jobType=%s cluster=%s", req.DisplayJobName, req.JobType, req.Cluster)
	oldCloudbrainId := req.SourceCloudbrainId
	if oldCloudbrainId <= 0 {
		return response.PARAM_ERROR
	}
	oldCloudbrain, err := models.GetCloudbrainByCloudbrainID(oldCloudbrainId)
	if err != nil {
		log.Error("CheckPrivilege4NewVersion get old cloudbrain task error.oldCloudbrainId=%d err=%v", oldCloudbrainId, err)
		if models.IsErrRecordNotExist(err) {
			return response.PARAM_ERROR
		}
		return response.SYSTEM_ERROR
	}
	ctx.SourceCloudbrain = oldCloudbrain

	if oldCloudbrain.UserID != ctx.User.ID && !ctx.User.IsAdmin {
		return response.INSUFFICIENT_PERMISSION
	}

	log.Info("CheckPrivilege4Continue success.displayJobName=%s jobType=%s cluster=%s", req.DisplayJobName, req.JobType, req.Cluster)
	return nil
}

func (DefaultCreationHandler) HandleReqParameters(ctx *context.CreationContext) *response.BizError {
	req := ctx.Request
	var parameters models.Parameters
	if req.Params != "" {
		err := json.Unmarshal([]byte(req.Params), &parameters)
		if err != nil {
			log.Error("Failed to Unmarshal params: %s (%v)", req.Params, err)
			return response.PARAM_ERROR
		}
		// label去掉所有的空格，value去掉首位的空格
		for i := 0; i < len(parameters.Parameter); i++ {
			parameters.Parameter[i].Label = strings.ReplaceAll(parameters.Parameter[i].Label, " ", "")
			parameters.Parameter[i].Value = strings.TrimSpace(parameters.Parameter[i].Value)
		}
		ctx.Request.ParamArray = parameters
		p, err := json.Marshal(parameters)
		if err == nil {
			ctx.Request.Params = string(p)
		}
	}
	return nil
}

func (DefaultCreationHandler) CheckBootFile(ctx *context.CreationContext) *response.BizError {
	req := ctx.Request
	branch := req.BranchName
	if req.BootFile == "" {
		return response.PARAM_ERROR
	}
	if !strings.HasSuffix(strings.TrimSpace(req.BootFile), ".py") {
		log.Error("the boot file(%s) must be a python file", strings.TrimSpace(req.BootFile))
		return response.BOOT_FILE_MUST_BE_PYTHON
	}
	if branch == "" {
		branch = ctx.Repository.DefaultBranch
	}
	commit, err := ctx.GitRepo.GetBranchCommit(branch)
	if err != nil {
		log.Error("CheckBootFile GetBranchCommit error,repoId:=%d err=%v", ctx.Repository.ID, err)
		return response.BOOT_FILE_NOT_EXIST
	}
	if _, err := commit.GetTreeEntryByPath(req.BootFile); err != nil {
		log.Error("CheckBootFile GetTreeEntryByPath error,repoId:=%d BootFile=%s err=%v", ctx.Repository.ID, req.BootFile, err)
		return response.BOOT_FILE_NOT_EXIST
	}
	return nil
}

func (DefaultCreationHandler) CheckMultiRequest(ctx *context.CreationContext) *response.BizError {
	jobType := string(ctx.Request.JobType)
	log.Info("Start to CheckMulti success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	count, err := cloudbrainTask.GetNotFinalStatusTaskCount(ctx.User.ID, jobType)
	if err != nil {
		log.Error("GetGrampusCountByUserID failed:%v", err)
		return response.SYSTEM_ERROR
	}

	limitNum := GetUserMultiLimitNum(ctx.User.ID, ctx.Request.JobType)
	if count >= limitNum {
		log.Error("the user already has running or waiting task.")
		return response.MULTI_TASK.WithParams(count)
	}
	log.Info("CheckMulti success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	return nil
}

func (DefaultCreationHandler) CheckDisplayJobName(ctx *context.CreationContext) *response.BizError {
	repo := ctx.Repository
	displayJobName := ctx.Request.DisplayJobName
	jobType := string(ctx.Request.JobType)
	log.Info("Start to CheckDisplayJobName.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	//check whether the task name in the project is duplicated
	tasks, err := models.GetCloudbrainsByDisplayJobName(repo.ID, jobType, displayJobName)
	if err == nil {
		if len(tasks) != 0 {
			log.Error("the job name did already exist.displayJobName=%s", displayJobName)
			return response.JOB_NAME_ALREADY_USED
		}
	} else {
		if !models.IsErrJobNotExist(err) {
			log.Error("system error, %v", err)
			return response.SYSTEM_ERROR
		}
	}
	log.Info("CheckDisplayJobName success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return nil
}

func (DefaultCreationHandler) LoadSpec(ctx *context.CreationContext) *response.BizError {
	log.Info("Start to LoadSpec.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	//check specification
	spec, err := resource.GetAndCheckSpec(ctx.User.ID, ctx.Request.SpecId, models.FindSpecsOptions{
		JobType:         ctx.Request.JobType,
		ComputeResource: ctx.Request.ComputeSource.Name,
		Cluster:         ctx.Request.Cluster.GetParentCluster(),
		HasInternet:     ctx.Request.HasInternet,
	})
	if err != nil || spec == nil {
		return response.SPEC_NOT_AVAILABLE
	}
	ctx.Spec = spec
	log.Info("LoadSpec success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	return nil
}

func (DefaultCreationHandler) InsertCloudbrainRecord4Async(ctx *context.CreationContext) *response.BizError {
	log.Info("Start to InsertCloudbrainRecord4Async.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	req := ctx.Request

	imageUrl := req.ImageUrl
	if req.ImageUrl == "" && req.ImageName != "" {
		imageUrl = req.ImageName
	}
	taskType := req.Cluster.GetCloudbrainType()
	if taskType == models.TypeCloudBrainTwo && setting.ModelartsCD.Enabled {
		taskType = models.TypeCDCenter
	}
	branchName := req.BranchName
	//在线运行notebook的请求分支放到了另一个字段
	if req.IsFileNoteBookRequest {
		branchName = req.FileBranchName
	}
	//if a normal user set TimeLimit, it does not work
	//if !role.UserHasRole(ctx.User.ID, models.Subscriber) {
	if !role.UserHasOper(ctx.User.ID, role.ROLE_OPER_DEBUG_TIME) {
		req.TimeLimit = 0
	}
	var repoId int64
	if ctx.Repository != nil {
		repoId = ctx.Repository.ID
	}
	var sourceCloudbrainId int64
	if ctx.SourceCloudbrain != nil {
		sourceCloudbrainId = ctx.SourceCloudbrain.ID
	}
	c := &models.Cloudbrain{
		Status:            models.LocalStatusPreparing,
		UserID:            ctx.User.ID,
		RepoID:            repoId,
		JobName:           req.JobName,
		DisplayJobName:    req.DisplayJobName,
		JobType:           string(req.JobType),
		Type:              taskType,
		Uuid:              req.DatasetUUIDStr,
		DatasetName:       req.DatasetNames,
		DatasetAlias:      req.DatasetAlias,
		CommitID:          ctx.CommitID,
		IsLatestVersion:   "1",
		VersionCount:      1,
		ComputeResource:   req.ComputeSource.GetCloudbrainFormat(),
		ImageID:           req.ImageID,
		Image:             imageUrl,
		BranchName:        branchName,
		Parameters:        req.Params,
		BootFile:          req.BootFile,
		Description:       req.Description,
		WorkServerNumber:  req.WorkServerNumber,
		EngineName:        imageUrl,
		Spec:              ctx.Spec,
		ModelName:         req.ModelNames,
		LabelName:         req.LabelName,
		ModelId:           req.PretrainModelId,
		SubTaskName:       models.SubTaskName,
		CreatedUnix:       timeutil.TimeStampNow(),
		UpdatedUnix:       timeutil.TimeStampNow(),
		GpuQueue:          ctx.Spec.QueueCode,
		AppName:           req.AppName,
		HasInternet:       int(req.HasInternet),
		TimeLimit:         req.TimeLimit,
		Port:              req.Port,
		EndPoint:          req.EndPoint,
		VisualizeRequired: req.VisualizeRequired,
		AimRequired:       setting.AimConfig.Enabled && role.UserHasOper(ctx.User.ID, role.ROLE_OPER_MonitorAdmin),
		SourceID:          sourceCloudbrainId,
		TemplateID:        req.TemplateID,
	}

	err := models.CreateCloudbrain(c)

	if err != nil {
		log.Error("DefaultHandler InsertCloudbrainRecord(%s) failed:%v", req.DisplayJobName, err)
		return response.NewBizError(err)
	}
	ctx.NewCloudbrain = c
	log.Info("InsertCloudbrainRecord4Async success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	return nil
}

func (DefaultCreationHandler) AfterCallCreationAPI4Sync(ctx *context.CreationContext) *response.BizError {
	log.Info("Start to AfterCallCreationAPI4Sync.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	req := ctx.Request
	res := ctx.Response
	if res == nil || res.Error != nil {
		//如果是同步调用，此时API直接返回错误，无需创建云脑记录
		return response.NewBizError(res.Error)
	}

	imageUrl := req.ImageUrl
	if req.ImageUrl == "" && req.ImageName != "" {
		imageUrl = req.ImageName
	}
	taskType := req.Cluster.GetCloudbrainType()
	if taskType == models.TypeCloudBrainTwo && setting.ModelartsCD.Enabled {
		taskType = models.TypeCDCenter
	}
	c := &models.Cloudbrain{
		UserID:           ctx.User.ID,
		RepoID:           ctx.Repository.ID,
		JobName:          req.JobName,
		DisplayJobName:   req.DisplayJobName,
		JobType:          string(req.JobType),
		Type:             taskType,
		Uuid:             req.DatasetUUIDStr,
		DatasetName:      req.DatasetNames,
		DatasetAlias:     req.DatasetAlias,
		CommitID:         ctx.CommitID,
		IsLatestVersion:  "1",
		ComputeResource:  req.ComputeSource.GetCloudbrainFormat(),
		ImageID:          req.ImageID,
		Image:            imageUrl,
		BranchName:       req.BranchName,
		Parameters:       req.Params,
		BootFile:         req.BootFile,
		Description:      req.Description,
		WorkServerNumber: req.WorkServerNumber,
		EngineName:       imageUrl,
		Spec:             ctx.Spec,
		ModelName:        req.ModelNames,
		LabelName:        req.LabelName,
		ModelId:          req.PretrainModelId,
		SubTaskName:      models.SubTaskName,
		JobID:            res.JobID,
		Status:           TransAITaskStatus(res.Status),
		CreatedUnix:      res.CreateTime,
		UpdatedUnix:      res.CreateTime,
		GpuQueue:         ctx.Spec.QueueCode,
		Config:           ctx.BuildCloudbrainConfig(),
	}

	config := ctx.BuildCloudbrainConfig()
	if config != nil {
		c.TrainUrl = path.Join("/", config.OutputBucket, config.OutputObjectPrefix)
		c.LogUrl = path.Join("/", config.LogBucket, config.LogObjectPrefix)
	}
	err := models.CreateCloudbrain(c)

	if err != nil {
		log.Error("DefaultHandler AfterCallCreationAPI4Sync(%s) failed:%v", req.DisplayJobName, err)
		return response.NewBizError(err)
	}
	ctx.NewCloudbrain = c
	log.Info("AfterCallCreationAPI4Sync success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return nil
}

func (DefaultCreationHandler) AfterCallCreationAPI4Async(ctx *context.CreationContext) *response.BizError {
	log.Info("Start to AfterCallCreationAPI4Async.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	c := ctx.NewCloudbrain
	if c == nil {
		log.Error("cloudbrain not exists.")
		return response.AI_TASK_NOT_EXISTS
	}
	res := ctx.Response
	if res == nil || res.Status == "" {
		log.Error("Response not exists.")
		return response.SYSTEM_ERROR
	}
	//更新commitId
	c.CommitID = ctx.CommitID
	c.JobID = res.JobID
	c.Status = TransAITaskStatus(res.Status)
	c.CreatedUnix = res.CreateTime
	c.UpdatedUnix = res.CreateTime
	c.DatasetName = ctx.Request.DatasetNames
	c.DatasetAlias = ctx.Request.DatasetAlias
	c.VersionName = res.VersionName
	c.VersionID = res.VersionID

	config := ctx.BuildCloudbrainConfig()
	if config != nil {
		c.TrainUrl = path.Join("/", config.OutputBucket, config.OutputObjectPrefix)
		c.LogUrl = path.Join("/", config.LogBucket, config.LogObjectPrefix)
	}
	err := models.UpdateJob(c)
	if err != nil {
		log.Error("AfterCallCreationAPI4Async UpdateJob err.displayJobName=%s jobType=%s cluster=%s err=%v", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster, err)
		return response.NewBizError(err)
	}
	log.Info("AfterCallCreationAPI4Async success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	//更新cloudbrain_config表
	config.CloudbrainID = c.ID
	_, err = models.InsertCloudbrainConfig(config)
	if err != nil {
		log.Error("InsertCloudbrainConfig error,config=%+v err=%v", config, err)
	}
	return nil

}

func (DefaultCreationHandler) CreateCloudbrainRecord4Restart(ctx *context.CreationContext) *response.BizError {
	log.Info("Start to CreateCloudbrainRecord4Restart.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	res := ctx.Response
	if res == nil {
		return nil
	}
	if res.Error != nil {
		return response.RESTART_FAILED
	}
	req := ctx.Request
	c := &models.Cloudbrain{
		Status:           TransAITaskStatus(res.Status),
		UserID:           ctx.SourceCloudbrain.UserID,
		RepoID:           ctx.SourceCloudbrain.RepoID,
		JobID:            res.JobID,
		JobName:          req.JobName,
		DisplayJobName:   req.DisplayJobName,
		JobType:          string(req.JobType),
		Type:             ctx.SourceCloudbrain.Type,
		Uuid:             req.DatasetUUIDStr,
		DatasetName:      req.DatasetNames,
		DatasetAlias:     req.DatasetAlias,
		CommitID:         ctx.SourceCloudbrain.CommitID,
		IsLatestVersion:  "1",
		ComputeResource:  req.ComputeSource.GetCloudbrainFormat(),
		ImageID:          req.ImageID,
		Image:            req.ImageUrl,
		BranchName:       req.BranchName,
		Parameters:       req.Params,
		BootFile:         req.BootFile,
		Description:      req.Description,
		WorkServerNumber: req.WorkServerNumber,
		EngineName:       req.ImageUrl,
		CreatedUnix:      res.CreateTime,
		UpdatedUnix:      res.CreateTime,
		Spec:             ctx.Spec,
		ModelName:        req.ModelNames,
		LabelName:        req.LabelName,
		SubTaskName:      models.SubTaskName,
		ModelId:          req.PretrainModelId,
		GpuQueue:         ctx.Spec.QueueCode,
		HasInternet:      int(req.HasInternet),
		TimeLimit:        ctx.SourceCloudbrain.TimeLimit,
		EndPoint:         req.EndPoint,
		Port:             req.Port,
	}
	err := models.RestartCloudbrain(ctx.SourceCloudbrain, c)

	if err != nil {
		log.Error("DefaultHandler CreateCloudbrain(%s) failed:%v", req.DisplayJobName, err)
		return response.SYSTEM_ERROR
	}
	ctx.NewCloudbrain = c
	otel.UpdateTraceCache(ctx.SourceCloudbrain.ID, c.ID)
	log.Info("CreateCloudbrainRecord4Restart success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	return nil
}

func TransAITaskStatus(oldStatus string) string {
	switch oldStatus {
	case models.GrampusStatusPending:
		return models.GrampusStatusWaiting
	}
	return strings.ToUpper(oldStatus)
}

func (DefaultCreationHandler) CheckPointBalance(ctx *context.CreationContext) *response.BizError {
	log.Info("Start to CheckPointBalance.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	if ctx.Spec == nil {
		return response.SPEC_NOT_AVAILABLE
	}
	if !account.IsPointBalanceEnough(ctx.User.ID, models.PointDeductCondition{SpecUnitPrice: ctx.Spec.UnitPrice, WorkServerNumber: ctx.Request.WorkServerNumber}) {
		return response.INSUFFICIENT_POINT_BALANCE
	}
	log.Info("CheckPointBalance success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	return nil
}

func (DefaultCreationHandler) GetAvailableQueues(ctx *context.CreationContext) *response.BizError {
	log.Info("Start to GetAvailableQueues.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	ctx.Queues = ctx.Spec.GetAvailableQueuesForNewRight(models.GetAvailableCenterIdOpts{
		UserId:            ctx.User.ID,
		JobType:           ctx.Request.JobType,
		HasInternet:       ctx.Request.HasInternet,
		VisualizeRequired: ctx.Request.VisualizeRequired,
	})
	log.Info("GetAvailableQueues %+v ", ctx.Queues)
	log.Info("GetAvailableQueues success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return nil
}

func (DefaultCreationHandler) CallCreationAPI(ctx *context.CreationContext) *response.BizError {
	log.Error("CallCreationAPI not implements")
	return response.SYSTEM_ERROR
}
func (DefaultCreationHandler) CallRestartAPI(ctx *context.CreationContext) *response.BizError {
	log.Error("CallRestartAPI not implements")
	return response.SYSTEM_ERROR
}

func (DefaultCreationHandler) HandleErr4Async(ctx *context.CreationContext) *response.BizError {
	log.Info("Start to HandleErr4Async.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)

	c := ctx.NewCloudbrain
	if c == nil {
		log.Error("HandleErr4Async cloudbrain not exists.")
		return nil
	}
	if ctx.Response == nil || ctx.Response.Error == nil {
		log.Error("HandleErr4Async response err is not exists.")
		return nil
	}
	responseErr := ctx.Response.Error

	cloudbrain, err := models.GetCloudbrainByCloudbrainID(c.ID)
	if cloudbrain == nil {
		log.Error("HandleErr4Async GetCloudbrainByCloudbrainID err.id=%d err=%v", c.ID, err)
		return nil
	}
	//只处理处于PREPARING状态的任务
	if !cloudbrain.IsPreparing() {
		return nil
	}
	//处理调用集群接口返回错误的情况
	if models.IsNetworkError(responseErr) {
		//如果是网络错误，创建是否成功未知，不修改状态，等待定时任务处理
		cloudbrain.Status = models.LocalStatusCreating
	} else {
		//非网络错误则认为创建失败
		cloudbrain.Status = models.LocalStatusFailed
		cloudbrain.FailedReason = responseErr.Error()
	}
	cloudbrain.CommitID = ctx.CommitID
	cloudbrain.DatasetName = ctx.Request.DatasetNames
	cloudbrain.DatasetAlias = ctx.Request.DatasetAlias

	err = models.UpdateJob(cloudbrain)
	if err != nil {
		log.Error("HandleErr4Async UpdateJob err.displayJobName=%s jobType=%s cluster=%s err=%v", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster, err)
		return response.NewBizError(err)
	}
	log.Info("HandleErr4Async success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return nil
}

func (g DefaultCreationHandler) NotifyCreation(ctx *context.CreationContext) *response.BizError {
	return nil
}

func (DefaultCreationHandler) CheckNotebookCount(ctx *context.CreationContext) *response.BizError {

	if ctx.Request.JobType == models.JobTypeDebug {
		count, err := models.GetNotebooksCountByUser(ctx.User.ID)
		if err != nil {
			log.Warn("can not get user notebook count", err)
		}
		if count >= int64(setting.NotebookStrategy.MaxNumberPerUser) {
			return response.NOTEBOOK_EXCEED_MAX_NUM.WithParams(setting.NotebookStrategy.MaxNumberPerUser)
		}
	}
	return nil
}

func (DefaultCreationHandler) CheckImageAvailable(ctx *context.CreationContext) *response.BizError {
	if ctx.Request.SourceCloudbrainId <= 0 {
		return nil
	}
	log.Info("start to gcu CheckImageAvailable")
	if ctx.Request.ComputeSource.Name == models.GCU {
		image_name := ctx.Request.ImageUrl
		log.Info("image_name=" + image_name)
		re, err := models.FindImageByImageUrl(image_name)
		if err != nil {
			return response.IMAGE_NOT_AVAILABLE
		}
		if len(re) == 0 {
			return response.IMAGE_NOT_AVAILABLE
		}
	}
	return nil
}
