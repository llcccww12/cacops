package ai_task

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/eval"
	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/otel"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/modules/util"
	"code.gitea.io/gitea/routers/common"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/schedule"
	"code.gitea.io/gitea/services/ai_task_service/task"
	"code.gitea.io/gitea/services/cloudbrain/resource"
	"code.gitea.io/gitea/services/role"
)

func CreateAITask(ctx *context.Context, form entity.CreateReq) {
	log.Info("start to here.....")
	handCreateReq(&form)
	if ctx.Repo == nil && form.JobType != models.JobTypeDebug && form.JobType != models.JobTypeGeneral && form.JobType != models.JobTypeSuperCompute {
		ctx.JSON(http.StatusOK, response.PARAM_ERROR)
	}
	var gitRepo *git.Repository
	var repo *models.Repository
	if ctx.Repo != nil {
		gitRepo = ctx.Repo.GitRepo
		repo = ctx.Repo.Repository
	}
	var traceErr error

	traceContext, _ := otel.StartTraceParent(&entity.TraceInfo{TaskName: form.DisplayJobName, JobType: string(form.JobType), SpanName: "CreateAITask"}, "POST")

	defer func() {

		otel.FinalizeSpan(traceContext, traceErr)
	}()
	res, err := task.CreateAITask(form, gitRepo, repo, ctx.User, &traceContext)
	if err != nil {
		traceErr = err.ToError()
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(res))
}
func BatchDelAITask(ctx *context.Context, from entity.DeleteIDs) {
	log.Info("BatchDelAITask begin")
	failed_ids := make([]int64, 0)
	for _, id := range from.ID {
		job, err1 := models.GetCloudbrainByCloudbrainID(id)
		if err1 != nil {
			log.Error("GetCloudbrainByID failed:%v", err1.Error())
			failed_ids = append(failed_ids, id)
			continue
		}
		if !isAdminOrOwnerOrJobCreator(ctx, job) {

			log.Error("!isAdminOrJobCreater error:%v")

			ctx.Error(http.StatusForbidden)
			return
		}
		t, _ := task.GetAITaskTemplateByCloudbrainId(id)
		if t == nil {
			log.Error("can not delete task param error。 %d", id)
			failed_ids = append(failed_ids, id)
			continue
		}
		err := t.Delete(id)
		if err != nil {
			log.Error("Delete error.%v   id:%d", err, id)
			failed_ids = append(failed_ids, id)
			continue
		}

	}

	if len(failed_ids) == 0 {
		ctx.JSON(http.StatusOK, response.OuterSuccess())
	} else {
		ctx.JSON(http.StatusOK, response.OuterErrorWithData(response.RESPONSE_CODE_ERROR_DEFAULT, "Delete Failed", failed_ids))
	}
}

func isAdminOrOwnerOrJobCreator(ctx *context.Context, job *models.Cloudbrain) bool {
	if !ctx.IsSigned {
		return false
	}
	repo, _ := models.GetRepositoryByID(job.RepoID)
	if repo == nil {
		return ctx.IsUserSiteAdmin() || ctx.User.ID == job.UserID
	} else {
		return repo.OwnerID == ctx.User.ID || ctx.IsUserSiteAdmin() || ctx.User.ID == job.UserID
	}

}

func DelAITask(ctx *context.Context) {
	id := ctx.QueryInt64("id")
	t, _ := task.GetAITaskTemplateByCloudbrainId(id)
	if t == nil {
		log.Error("param error")
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.PARAM_ERROR, ctx))
		return
	}
	err := t.Delete(id)
	if err != nil {
		log.Error("Delete error.%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccess())
}
func StopAITask(ctx *context.Context) {
	id := ctx.QueryInt64("id")
	t, err := task.GetAITaskTemplateByCloudbrainId(id)
	if err != nil {
		log.Error("param error")
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.PARAM_ERROR, ctx))
		return
	}
	res, err := t.Stop(id)
	if err != nil {
		log.Error("Stop error.%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	res.Tr(ctx.Language())
	res.ClearNonPublicFields()
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(res))
}
func RestartAITask(ctx *context.Context) {
	id := ctx.QueryInt64("id")
	res, bizErr := task.RestartAITask(id, ctx.Repo.GitRepo, ctx.Repo.Repository, ctx.User)
	if bizErr != nil {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(bizErr, ctx))
		return
	}

	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(res))
}

func GetEvalResult(ctx *context.Context) {
	id := ctx.QueryInt64("id")
	task, err := models.GetCloudbrainByCloudbrainID(id)
	if err != nil {
		ctx.NotFound("not found", err)
		return
	}
	if task.JobType != string(models.JobTypeEval) {
		ctx.NotFound("not found", nil)
		return
	}
	if task.UserID != ctx.User.ID && !ctx.IsUserSiteAdmin() {
		ctx.Context.Error(http.StatusUnauthorized)
		return
	}
	result, err := eval.GetEvalTaskResult(id)
	if err != nil {
		ctx.JSON(http.StatusOK, structs.EvalResultResponse{Code: 1, Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func GetEvalDetailResult(ctx *context.Context, detailBody structs.TaskDetailRequest) {
	task, err := models.GetCloudbrainByCloudbrainID(detailBody.TaskId)
	if err != nil {
		ctx.NotFound("not found", err)
		return
	}
	if task.JobType != string(models.JobTypeEval) {
		ctx.NotFound("not found", nil)
		return
	}
	if task.UserID != ctx.User.ID && !ctx.IsUserSiteAdmin() {
		ctx.Context.Error(http.StatusUnauthorized)
		return
	}

	result, err := eval.GetEvalTaskDetailResult(detailBody)
	if err != nil {
		ctx.JSON(http.StatusOK, structs.EvalDetailResultResponse{Code: 1, Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func GetAITaskLoss(ctx *context.Context) {
	id := ctx.QueryInt64("id")

	logFileName := ctx.Query("file_name")
	job, bizErr := models.GetCloudbrainByCloudbrainID(id)
	if bizErr != nil {
		log.Error("GetAITaskLog GetCloudbrainByCloudbrainID err.%v", bizErr)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.AI_TASK_NOT_EXISTS, ctx))
		return
	}
	t, err := task.GetAITaskTemplateFromCloudbrain(job)
	if err != nil {
		log.Error("param error")
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	res, err := t.GetLog(entity.QueryLogOpts{
		CloudbrainId: id,
		BaseLine:     0,
		Lines:        0,
		Order:        entity.UP,
		NodeId:       0,
		LogFileName:  logFileName,
	})
	if err != nil {
		log.Error("GetAITaskLog error.%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	losses := make([]entity.FinetuneLoss, 0)
	if res.Content != "" {
		scanner := bufio.NewScanner(strings.NewReader(res.Content))
		for scanner.Scan() {
			text := scanner.Text()
			text = strings.TrimSpace(text)
			if strings.Contains(text, "{'loss'") && strings.Contains(text, "grad_norm") && strings.Contains(text, "learning_rate") && strings.Contains(text, "epoch") {
				loss := entity.FinetuneLoss{}

				begin := strings.Index(text, "{'loss'")
				end := strings.LastIndex(text, "}")
				if end > begin {
					loss_str := text[begin : end+1]
					loss_str = strings.ReplaceAll(loss_str, "'", "\"")
					err := json.Unmarshal([]byte(loss_str), &loss)
					if err != nil {
						log.Warn("text can not parse error.%v %s", err, text)
					} else {
						losses = append(losses, loss)
					}
				}

			}

		}

	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(losses))
}

func GetAITaskLog(ctx *context.Context) {
	id := ctx.QueryInt64("id")
	baseLine := ctx.QueryInt64("base_line")
	lines := ctx.QueryInt64("lines")
	order := ctx.Query("order")
	nodeId := ctx.QueryInt("node_id")
	logFileName := ctx.Query("file_name")
	job, bizErr := models.GetCloudbrainByCloudbrainID(id)
	if bizErr != nil {
		log.Error("GetAITaskLog GetCloudbrainByCloudbrainID err.%v", bizErr)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.AI_TASK_NOT_EXISTS, ctx))
		return
	}
	t, err := task.GetAITaskTemplateFromCloudbrain(job)
	if err != nil {
		log.Error("param error")
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	res, err := t.GetLog(entity.QueryLogOpts{
		CloudbrainId: id,
		BaseLine:     baseLine,
		Lines:        lines,
		Order:        entity.Direction(order),
		NodeId:       nodeId,
		LogFileName:  logFileName,
	})
	if err != nil {
		log.Error("GetAITaskLog error.%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	if res.Content != "" && cloudbrain.IsAdminOrOwnerOrJobCreater(ctx, job, nil) {
		res.CanLogDownload = true
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(res))
}

func DownloadAITaskLog(ctx *context.Context) {
	id := ctx.QueryInt64("id")
	nodeId := ctx.QueryInt("node_id")
	logFileName := ctx.Query("file_name")
	cloudbrain, bizErr := models.GetCloudbrainByCloudbrainID(id)
	if bizErr != nil {
		log.Error("DownloadAITaskLog GetCloudbrainByCloudbrainID err.%v", bizErr)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.AI_TASK_NOT_EXISTS, ctx))
		return
	}
	t, err := task.GetAITaskTemplateFromCloudbrain(cloudbrain)
	if err != nil {
		log.Error("param error")
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	res, err := t.GetLogDownloadInfo(entity.GetLogDownloadInfoReq{
		CloudbrainId: id,
		NodeId:       nodeId,
		LogFileName:  logFileName,
	})
	if err != nil {
		log.Error("DownloadAITaskLog error.%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}

	if res == nil || res.IsEmpty() {
		log.Error("DownloadAITaskLog error.%v", err)
		ctx.JSON(http.StatusNotFound, "")
		return
	}

	tmpErr := common.WriteDownloadContent2Resp(ctx, res)
	if tmpErr != nil {
		log.Error("DownloadAITaskLog error.%v", tmpErr)
		ctx.JSON(http.StatusOK, response.OuterResponseError(tmpErr))
		return
	}

}

func DownloadOutputFile(ctx *context.Context) {
	id := ctx.QueryInt64("id")
	fileName := ctx.Query("file_name")
	parentDir := ctx.Query("parent_dir")
	cloudbrain, bizErr := models.GetCloudbrainByCloudbrainID(id)
	if bizErr != nil {
		log.Error("DownloadOutputFile GetCloudbrainByCloudbrainID err.%v", bizErr)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.AI_TASK_NOT_EXISTS, ctx))
		return
	}
	t, err := task.GetAITaskTemplateFromCloudbrain(cloudbrain)
	if err != nil {
		log.Error("param error")
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	res, err := t.GetSingleOutputDownloadInfo(entity.GetSingleDownloadInfoReq{
		CloudbrainId: id,
		FileName:     fileName,
		ParentDir:    parentDir,
	})
	if err != nil {
		log.Error("DownloadOutputFile error.%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}

	if res == nil || res.IsEmpty() {
		log.Error("DownloadOutputFile error.%v", err)
		ctx.JSON(http.StatusNotFound, "")
		return
	}

	tmpErr := common.WriteDownloadContent2Resp(ctx, res)
	if tmpErr != nil {
		log.Error("DownloadAITaskLog error.%v", tmpErr)
		ctx.JSON(http.StatusOK, response.OuterResponseError(tmpErr))
		return
	}

}

func DownloadAllOutputFile(ctx *context.Context) {
	id := ctx.QueryInt64("id")
	cloudbrain, bizErr := models.GetCloudbrainByCloudbrainID(id)
	if bizErr != nil {
		log.Error("DownloadAllOutputFile GetCloudbrainByCloudbrainID err.%v", bizErr)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.AI_TASK_NOT_EXISTS, ctx))
		return
	}
	t, err := task.GetAITaskTemplateFromCloudbrain(cloudbrain)
	if err != nil {
		log.Error("param error")
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	resultFileName := cloudbrain.JobName + ".zip"
	ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+url.QueryEscape(resultFileName))
	ctx.Resp.Header().Set("Content-Type", "application/octet-stream")
	zipWriter := zip.NewWriter(ctx.Resp)
	defer zipWriter.Close()
	err = t.DownloadAllOutput(entity.DownloadAllFileReq{
		CloudbrainId: id,
		ZIPWriter:    zipWriter,
	})
	if err != nil {
		log.Error("DownloadAllOutput error.%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
}

func GetAITaskInfo(ctx *context.Context) {
	id := ctx.QueryInt64("id")
	job, bizErr := models.GetCloudbrainByCloudbrainID(id)
	if bizErr != nil {
		log.Error("GetAITaskInfo GetCloudbrainByCloudbrainID err.%v", bizErr)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.AI_TASK_NOT_EXISTS, ctx))
		return
	}
	t, err := task.GetAITaskTemplateFromCloudbrain(job)
	if err != nil {
		log.Error("param error")
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	resultTask, err := t.Query(id)
	if err != nil {
		log.Error("Query error.id=%d err=%v", id, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	//加载关联版本
	earlyVersionList, bizErr := task.QueryTaskEarlyVersionList(id)
	if bizErr != nil {
		log.Error("QueryTaskEarlyVersionList err.id=%d  err=%v", id, err)
		ctx.JSON(http.StatusOK, response.OuterResponseError(bizErr))
		return
	}
	isSubscriber := false
	if ctx.User != nil {
		isSubscriber = role.UserHasOper(ctx.User.ID, role.ROLE_OPER_ONLINE_INFER_PATH)
		if ctx.User.IsAdmin {
			isSubscriber = true
		}
	}

	res := &entity.QueryAITaskRes{
		Task:                 resultTask,
		CanDownload:          cloudbrain.CanDownloadJob(ctx, job),
		EarlyVersionList:     earlyVersionList,
		CanCreateVersion:     job.CanUserModify(ctx.User),
		CanModify:            job.CanUserModify(ctx.User),
		CanDelete:            job.CanUserDelete(ctx.User, ctx.Repo.IsOwner()),
		CanCreateTemplate:    job.CanCreateTemplate(ctx.User),
		CanExperience:        job.CanExperience(ctx.User),
		CanFintuneExperience: job.CanFintuneExperience(ctx.User),
		IsSubscriber:         isSubscriber,
	}
	//根据权限去掉数据集和模型信息、EndPoint等信息
	res.TryToRemoveNoNeedInfo(ctx.User)
	//国际化
	res.Tr(ctx.Language())
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(res))
}

func GetAITaskBriefInfo(ctx *context.Context) {
	id := ctx.QueryInt64("id")
	t, err := task.GetAITaskTemplateByCloudbrainId(id)
	if err != nil {
		log.Error("param error")
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	res, err := t.BriefQuery(id)
	if err != nil {
		log.Error("BriefQuery error.%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	res.Tr(ctx.Language())
	res.ClearNonPublicFields()
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(res))
}

func GetAITaskOutput(ctx *context.Context) {
	id := ctx.QueryInt64("id")
	parentDir := ctx.Query("parent_dir")
	cloudbrainTask, bizErr := models.GetCloudbrainByCloudbrainID(id)
	if bizErr != nil {
		log.Error("GetAITaskOutput GetCloudbrainByCloudbrainID err.%v", bizErr)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.AI_TASK_NOT_EXISTS, ctx))
		return
	}
	t, err := task.GetAITaskTemplateFromCloudbrain(cloudbrainTask)
	if err != nil {
		log.Error("param error")
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	res, err := t.GetOutput(id, parentDir)
	if err != nil {
		log.Error("GetOutput error.%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	res.CanReschedule = cloudbrain.CanDeleteJob(ctx, cloudbrainTask)
	res.CanDownload = cloudbrain.CanDownloadJob(ctx, cloudbrainTask)
	// 获取 output 存储限制
	if ctx.User != nil {
		_, res.OutputSizeLimit = role.GetUserContainerStorageLimits(ctx.User.ID)
	}

	m := map[string]interface{}{"output": res}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

func GetAllAITaskOutput(ctx *context.Context) {
	id := ctx.QueryInt64("id")
	suffixStr := ctx.Query("suffix")
	var suffix []string
	if suffixStr != "" {
		suffixList := strings.Split(suffixStr, "|")
		for i := 0; i < len(suffixList); i++ {
			if suffixList[i] != "" {
				suffix = append(suffix, suffixList[i])
			}
		}
	}
	cloudbrainTask, bizErr := models.GetCloudbrainByCloudbrainID(id)
	if bizErr != nil {
		log.Error("GetAITaskOutput GetCloudbrainByCloudbrainID err.%v", bizErr)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.AI_TASK_NOT_EXISTS, ctx))
		return
	}
	t, err := task.GetAITaskTemplateFromCloudbrain(cloudbrainTask)
	if err != nil {
		log.Error("param error")
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	res, err := t.GetAllOutput(entity.GetAllOutputReq{
		CloudbrainId: cloudbrainTask.ID,
		Suffix:       suffix,
	})
	if err != nil {
		log.Error("GetAllOutput error.%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}

	m := map[string]interface{}{"output": res}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

func GetNoteBookUrlById(id int64) (string, *response.BizError) {
	t, err := task.GetAITaskTemplateByCloudbrainId(id)
	if err != nil {
		log.Error("param error")
		return "", err
	}
	url, err := t.GetDebugUrl(id, "")
	return url, err
}

func GetSelfEndPointUrlById(id int64) (string, *response.BizError) {
	t, err := task.GetAITaskTemplateByCloudbrainId(id)
	if err != nil {
		log.Error("param error")
		return "", err
	}
	fileName := ""
	url, err := t.GetSelfEndPointUrl(id, fileName)
	return url, err
}

func GetSelfEndPointUrl(ctx *context.Context) {
	id := ctx.QueryInt64("id")
	t, err := task.GetAITaskTemplateByCloudbrainId(id)
	if err != nil {
		log.Error("param error")
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	fileName := ctx.QueryTrim("file")
	url, err := t.GetSelfEndPointUrl(id, fileName)
	if err != nil {
		log.Error("GetSelfEndPointUrl error.%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	m := map[string]interface{}{"url": url}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

func GetNotebookUrl(ctx *context.Context) {
	id := ctx.QueryInt64("id")
	t, err := task.GetAITaskTemplateByCloudbrainId(id)
	if err != nil {
		log.Error("param error")
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	fileName := ctx.QueryTrim("file")
	url, err := t.GetDebugUrl(id, fileName)
	if err != nil {
		log.Error("GetNotebookUrl error.%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	// 如果url中包含ptoken=，split url and ptoken
	if strings.Contains(url, "&ptoken=") {
		urlArr := strings.Split(url, "&ptoken=")
		m := map[string]interface{}{"url": urlArr[0], "ptoken": urlArr[1]}
		ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
		return
	}
	m := map[string]interface{}{"url": url}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

func GetVisualizeUrl(ctx *context.Context) {
	id := ctx.QueryInt64("id")
	t, err := task.GetAITaskTemplateByCloudbrainId(id)
	if err != nil {
		log.Error("param error")
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	url, err := t.GetVisualizeUrl(id)
	if err != nil {
		log.Error("GetVisualizeUrl error.%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	m := map[string]interface{}{"url": url}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

func GetNodeInfo(ctx *context.Context) {
	id := ctx.QueryInt64("id")
	t, err := task.GetAITaskTemplateByCloudbrainId(id)
	if err != nil {
		log.Error("param error")
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	res, err := t.GetNodeInfo(id)
	if err != nil {
		log.Error("GetNodeInfo error.%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}

	m := map[string]interface{}{"nodes": res}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(m))
}

func GetImageInfoBySelectedSpec(ctx *context.Context) {
	jobType := ctx.Query("job_type")

	if models.JobType(jobType) == (models.JobTypeOnlineInference) {
		jobType = string(models.JobTypeDebug)
	}
	log.Info("required jobType=" + jobType)
	computeSourceName := ctx.Query("compute_source")
	clusterType := ctx.Query("cluster_type")

	computeSource := models.GetComputeSourceInstance(computeSourceName)
	specId := ctx.QueryInt64("spec_id")
	hasInternet := ctx.QueryInt("has_internet")
	visualizeRequired := ctx.QueryBool("visualize_required")

	spec, err := resource.GetAndCheckSpec(ctx.User.ID, specId, models.FindSpecsOptions{
		JobType:           models.JobType(jobType),
		ComputeResource:   computeSourceName,
		Cluster:           clusterType,
		HasInternet:       models.SpecInternetQuery(hasInternet),
		VisualizeRequired: visualizeRequired,
	})

	if err != nil || spec == nil {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.SPEC_NOT_AVAILABLE, ctx))
		return
	}

	result, bizerr := task.GetAvailableImageInfoBySpec(entity.GetAITaskCreationImageInfoReq{
		ClusterType:       entity.ClusterType(clusterType),
		ComputeSource:     computeSource,
		Spec:              spec,
		JobType:           models.JobType(jobType),
		UserID:            ctx.User.ID,
		VisualizeRequired: visualizeRequired,
		HasInternet:       models.SpecInternetQuery(hasInternet),
	})
	if bizerr != nil {
		log.Error("GetAITaskImageCreationInfo error,err=%v", bizerr)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(bizerr, ctx))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(result))
}

func GetCreationRequiredInfo(ctx *context.Context) {
	jobType := ctx.Query("job_type")
	var isOnlineType bool
	log.Info("required jobType=" + jobType)
	computeSourceName := ctx.Query("compute_source")
	clusterType := ctx.Query("cluster_type")
	VisualizeRequired := ctx.QueryBool("visualize_required")
	computeSource := models.GetComputeSourceInstance(computeSourceName)
	var gitRepo *git.Repository
	var repo *models.Repository
	if ctx.Repo != nil {
		gitRepo = ctx.Repo.GitRepo
		repo = ctx.Repo.Repository
	}

	result, err := task.GetAITaskCreationInfo(entity.GetAITaskCreationInfoReq{
		User:              ctx.User,
		JobType:           models.JobType(jobType),
		ClusterType:       entity.ClusterType(clusterType),
		Repo:              repo,
		GitRepo:           gitRepo,
		ComputeSource:     computeSource,
		IsOnlineType:      isOnlineType,
		VisualizeRequired: VisualizeRequired,
	})
	if err != nil {
		log.Error("GetAITaskCreationInfo error,err=%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(result))
}

func GetCreationRequiredRepoInfo(ctx *context.Context) {
	result, err := task.GetAITaskCreationRepoInfo(entity.GetAITaskCreationInfoReq{
		GitRepo: ctx.Repo.GitRepo,
		Repo:    ctx.Repo.Repository,
	})
	if err != nil {
		log.Error("GetAITaskCreationInfo error,err=%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(result))
}

func GetAITaskList(ctx *context.Context) {
	jobType := ctx.Query("job_type")
	status := ctx.Query("job_status")
	aiCenter := ctx.Query("ai_center")
	cluster := ctx.Query("cluster")
	excludeJobTypes := ctx.QueryStrings("exclude_job_types")
	computeSourceName := ctx.Query("compute_source")
	isAimRequired := ctx.QueryBool("aim")
	page := ctx.QueryInt("page")
	computeSource := models.GetComputeSourceInstance(computeSourceName)
	if page <= 0 {
		page = 1
	}
	jobTypes := make([]string, 0)
	if jobType != "" {
		jobTypes = append(jobTypes, jobType)
	}
	excludeStatus := []string{}
	if status == "other" {
		statusStr := ctx.Query("exclude_status")
		if statusStr == "" {
			ctx.JSON(http.StatusOK, response.PARAM_ERROR)
		}
		excludeStatus = strings.Split(statusStr, ",")
	}
	result, err := task.GetRepoAITaskList(entity.GetTaskListReq{
		ListOptions: models.ListOptions{
			PageSize: setting.UI.IssuePagingNum,
			Page:     page,
		},
		ComputeSource:   computeSource,
		JobTypes:        jobTypes,
		RepoID:          ctx.Repo.Repository.ID,
		Operator:        ctx.User,
		IsRepoOwner:     ctx.Repo.IsOwner(),
		AICenter:        aiCenter,
		JobStatus:       status,
		ExcludeStatus:   excludeStatus,
		Cluster:         cluster,
		ExcludeJobTypes: excludeJobTypes,
		IsAimRequired:   isAimRequired,
	})
	if err != nil {
		log.Error("GetAITaskList error,err=%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	result.CanCreateTask = cloudbrain.CanCreateOrDebugJob(ctx)
	result.IsRepoEmpty = ctx.Repo.Repository.IsEmpty
	result.IsSubscriber = false
	result.IsAdmin = false
	if ctx.User != nil {
		//result.IsSubscriber = role.UserHasRole(ctx.User.ID, models.Subscriber)
		result.IsSubscriber = role.UserHasOper(ctx.User.ID, role.ROLE_OPER_ONLINE_INFER_PATH)
		if ctx.User.IsAdmin {
			result.IsSubscriber = true
			result.IsAdmin = true
		}
	}

	for _, r := range result.Tasks {
		r.Task.Tr(ctx.Language())
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(result))
}

func GetAITaskOperationProfile(ctx *context.Context) {
	id := ctx.QueryInt64("id")
	t, err := task.GetAITaskTemplateByCloudbrainId(id)
	if err != nil {
		log.Error("param error")
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	r, err := t.GetOperationProfile(id)
	if err != nil {
		log.Error("GetOperationProfile error.%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(r))
}

func GetAITaskResourceUsage(ctx *context.Context) {
	id := ctx.QueryInt64("id")
	nodeId := ctx.QueryInt("node_id")
	logFileName := ctx.Query("file_name")
	t, err := task.GetAITaskTemplateByCloudbrainId(id)
	if err != nil {
		log.Error("param error")
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	r, err := t.GetResourceUsage(entity.GetResourceUsageOpts{
		CloudbrainId: id,
		NodeId:       nodeId,
		LogFileName:  logFileName,
	})
	if err != nil {
		log.Error("GetOperationProfile error.%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(r))
}
func GetAllMonitorAITask(ctx *context.Context) {

	fileName := "cloudbrain.csv"
	status := ctx.QueryTrim("status")
	minDuration := ctx.QueryInt("duration")
	days := ctx.QueryInt("days")

	ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+url.QueryEscape(fileName))
	ctx.Resp.Header().Set("Content-Type", "application/octet-stream")
	err := task.MonitorTaskFile(status, minDuration, days, ctx)
	if err != nil {
		log.Error("get monitor tasks err", err)
	}

}

func RetryModelSchedule(ctx *context.APIContext) {
	id := ctx.QueryInt64("id")
	if id <= 0 {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.PARAM_ERROR, ctx))
		return
	}
	job, err := models.GetCloudbrainByCloudbrainID(id)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.PARAM_ERROR, ctx))
		return
	}
	err = schedule.RetryModelMigrate(job)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NewBizError(err), ctx))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func handCreateReq(req *entity.CreateReq) {
	req.JobName = util.ConvertDisplayJobNameToJobName(req.DisplayJobName)
	if req.WorkServerNumber == 0 {
		req.WorkServerNumber = 1
	}

}

func GenerateSDKCode(ctx *context.Context) {
	datasetNames := ctx.QueryStrings("dataset_name")
	pretrainModelNames := ctx.QueryStrings("pretrain_model_name")
	parameterKeys := ctx.QueryStrings("param_key")
	jobType := ctx.Query("job_type")
	visualizeRequired := ctx.QueryBool("visualize_required")
	code := task.GenerateSDKCode(entity.SDKCodeOpts{
		DatasetNames:       datasetNames,
		PretrainModelNames: pretrainModelNames,
		ParameterKeys:      parameterKeys,
		JobType:            models.JobType(jobType),
		VisualizeRequired:  visualizeRequired,
	})
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(map[string]string{"code": code}))
}

func GetMyAITaskList(ctx *context.Context) {
	jobType := ctx.Query("job_type")
	jobStatus := ctx.Query("job_status")
	aiCenter := ctx.Query("ai_center")
	cluster := ctx.Query("cluster")
	appName := ctx.Query("app_name")
	OrderBy := ctx.Query("order_by")
	computeSourceName := ctx.Query("compute_source")
	excludeStatus := []string{}
	if jobStatus == "other" {
		statusStr := ctx.Query("exclude_status")
		if statusStr == "" {
			ctx.JSON(http.StatusOK, response.PARAM_ERROR)
		}
		excludeStatus = strings.Split(statusStr, ",")
	}
	keyword := strings.Trim(ctx.Query("q"), " ")
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := ctx.QueryInt("pageSize")
	if pageSize <= 0 || pageSize > 100 {
		pageSize = setting.UI.IssuePagingNum
	}
	computeSource := models.GetComputeSourceInstance(computeSourceName)
	result, err := task.GetMyAITaskList(entity.GetMyTaskListReq{
		ListOptions: models.ListOptions{
			PageSize: pageSize,
			Page:     page,
		},
		ComputeSource: computeSource,
		JobType:       jobType,
		User:          ctx.User,
		IsRepoOwner:   ctx.Repo.IsOwner(),
		JobStatus:     jobStatus,
		Keyword:       keyword,
		AICenter:      aiCenter,
		Cluster:       cluster,
		ExcludeStatus: excludeStatus,
		AppName:       appName,
		OrderBy:       OrderBy,
	})
	if err != nil {
		log.Error("GetMyAITaskList error,err=%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	result.CanCreateTask = cloudbrain.CanCreateOrDebugJob(ctx)
	result.IsSubscriber = false
	result.IsAdmin = false
	if ctx.User != nil {
		//result.IsSubscriber = role.UserHasRole(ctx.User.ID, models.Subscriber)
		result.IsSubscriber = role.UserHasOper(ctx.User.ID, role.ROLE_OPER_ONLINE_INFER_PATH)
		if ctx.User.IsAdmin {
			result.IsSubscriber = true
			result.IsAdmin = true
		}
	}

	for _, r := range result.Tasks {
		r.Task.Tr(ctx.Language())
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(result))
}

func GetAITaskList4Admin(ctx *context.Context) {
	BeginTime, EndTime := getTaskBeginAndEndTime(ctx)
	jobType := ctx.Query("job_type")
	jobStatus := ctx.Query("job_status")
	aiCenter := ctx.Query("ai_center")
	cluster := ctx.Query("cluster")
	orderBy := ctx.Query("order_by")
	computeSourceName := ctx.Query("compute_source")
	excludeStatus := []string{}
	if jobStatus == "other" {
		statusStr := ctx.Query("exclude_status")
		if statusStr == "" {
			ctx.JSON(http.StatusOK, response.PARAM_ERROR)
		}
		excludeStatus = strings.Split(statusStr, ",")
	}

	keyword := strings.Trim(ctx.Query("q"), " ")
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := ctx.QueryInt("pageSize")
	if pageSize <= 0 || pageSize > 100 {
		pageSize = setting.UI.IssuePagingNum
	}
	computeSource := models.GetComputeSourceInstance(computeSourceName)
	result, err := task.GetAITaskList4Admin(entity.GetMyTaskListReq{
		ListOptions: models.ListOptions{
			PageSize: pageSize,
			Page:     page,
		},
		ComputeSource: computeSource,
		JobType:       jobType,
		User:          ctx.User,
		JobStatus:     jobStatus,
		Keyword:       keyword,
		AICenter:      aiCenter,
		Cluster:       cluster,
		ExcludeStatus: excludeStatus,
		BeginTimeUnix: BeginTime.Unix(),
		EndTimeUnix:   EndTime.Unix(),
		OrderBy:       orderBy,
	})
	if err != nil {
		log.Error("GetMyAITaskList error,err=%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	result.CanCreateTask = cloudbrain.CanCreateOrDebugJob(ctx)
	result.IsSubscriber = false
	result.IsAdmin = false
	if ctx.User != nil {
		//result.IsSubscriber = role.UserHasRole(ctx.User.ID, models.Subscriber)
		result.IsSubscriber = role.UserHasOper(ctx.User.ID, role.ROLE_OPER_ONLINE_INFER_PATH)
		if ctx.User.IsAdmin {
			result.IsSubscriber = true
			result.IsAdmin = true
		}
	}
	for _, r := range result.Tasks {
		r.Task.Tr(ctx.Language())
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(result))
}

func getTaskBeginAndEndTime(ctx *context.Context) (time.Time, time.Time) {
	queryType := ctx.QueryTrim("type")
	now := time.Now()
	beginTimeStr := ctx.QueryTrim("beginTime")
	endTimeStr := ctx.QueryTrim("endTime")

	var beginTime time.Time
	var endTime time.Time
	if queryType != "" {
		if queryType == "all" {
			recordCloudbrain, err := models.GetRecordBeginTime()
			if err != nil {
				log.Error("Can not get recordCloudbrain", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("repo.record_begintime_get_err"))
			}
			brainRecordBeginTime := recordCloudbrain[0].Cloudbrain.CreatedUnix.AsTime()
			beginTime = brainRecordBeginTime
			endTime = now
		} else if queryType == "today" {
			beginTime = now.AddDate(0, 0, 0)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = now

		} else if queryType == "yesterday" {
			beginTime = now.AddDate(0, 0, -1)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		} else if queryType == "last_7day" {
			beginTime = now.AddDate(0, 0, -7)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = now
		} else if queryType == "last_30day" {
			beginTime = now.AddDate(0, 0, -30)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = now
		} else if queryType == "current_month" {
			endTime = now
			beginTime = time.Date(endTime.Year(), endTime.Month(), 1, 0, 0, 0, 0, now.Location())

		} else if queryType == "current_year" {
			endTime = now
			beginTime = time.Date(endTime.Year(), 1, 1, 0, 0, 0, 0, now.Location())
		} else if queryType == "last_month" {
			lastMonthTime := now.AddDate(0, -1, 0)
			beginTime = time.Date(lastMonthTime.Year(), lastMonthTime.Month(), 1, 0, 0, 0, 0, now.Location())
			endTime = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		} else if queryType == "last_year" {
			lastYear := now.Year() - 1                                           // 获取去年的年份
			beginTime = time.Date(lastYear, 1, 1, 0, 0, 0, 0, now.Location())    // 去年1月1日 00:00:00
			endTime = time.Date(lastYear, 12, 31, 23, 59, 59, 0, now.Location()) // 去年12月31日 23:59:59
		}

	} else {
		if beginTimeStr == "" || endTimeStr == "" {
			//如果查询类型和开始时间结束时间都未设置，按queryType=last_30day处理
			beginTime = now.AddDate(0, 0, -30)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = now
		} else {
			beginTime = now.AddDate(0, 0, -30)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = now
		}

	}
	return beginTime, endTime
}
