package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"code.gitea.io/gitea/services/cloudbrain/modelmanage"

	"code.gitea.io/gitea/services/lock"

	"code.gitea.io/gitea/routers/response"

	"code.gitea.io/gitea/services/cloudbrain/cloudbrainTask"

	"code.gitea.io/gitea/modules/dataset"

	"code.gitea.io/gitea/services/cloudbrain/resource"

	"code.gitea.io/gitea/services/reward/point/account"

	"code.gitea.io/gitea/modules/auth"
	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/grampus"
	"code.gitea.io/gitea/modules/modelarts"

	// "code.gitea.io/gitea/modules/notification"
	// "code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/modules/util"
	"github.com/unknwon/com"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	ai_task "code.gitea.io/gitea/services/ai_task_service/task"
	cloudbrainService "code.gitea.io/gitea/services/cloudbrain"
)

const (
	tplGrampusTrainJobShow base.TplName = "repo/grampus/trainjob/show"
	tplGrampusNotebookShow base.TplName = "repo/grampus/notebook/show"

	//GPU "repo/grampus/notebook/gpu/image/submit.tmpl"
	tplGrampusNotebookGPUNew       base.TplName = "repo/grampus/notebook/gpu/new"
	tplGrampusNotebookGPUImageShow base.TplName = "repo/cloudbrain/image/submit"
	tplGrampusTrainJobGPUNew       base.TplName = "repo/grampus/trainjob/gpu/new"

	//NPU
	tplGrampusNotebookNPUNew base.TplName = "repo/grampus/notebook/npu/new"
	tplGrampusTrainJobNPUNew base.TplName = "repo/grampus/trainjob/npu/new"
	//GCU
	tplGrampusNotebookGCUNew base.TplName = "repo/grampus/notebook/gcu/new"
	tplGrampusTrainJobGCUNew base.TplName = "repo/grampus/trainjob/gcu/new"

	//MLU
	tplGrampusNotebookMLUNew base.TplName = "repo/grampus/notebook/mlu/new"
	tplGrampusTrainJobMLUNew base.TplName = "repo/grampus/trainjob/mlu/new"

	//IluvatarGPGPU
	tplGrampusTrainJobIluvatarGPGPUNew base.TplName = "repo/grampus/trainjob/iluvatar-gpgpu/new"

	//IluvatarGPGPU
	tplGrampusTrainJobDCUNew base.TplName = "repo/grampus/trainjob/dcu/new"

	//BIREN-GPU
	tplGrampusTrainJobBIRENGPUNew base.TplName = "repo/grampus/trainjob/biren-gpu/new"

	//C2NET notebook
	tplGrampusNotebookNew base.TplName = "repo/grampus/notebook/new"

	// Inference job
	tplGrampusInferenceNew  base.TplName = "repo/grampus/inference/new"
	tplGrampusInferenceShow base.TplName = "repo/grampus/inference/show"
)

func GrampusInferenceNew(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(http.StatusOK, tplGrampusInferenceNew)
}

func GrampusInferenceShow(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(http.StatusOK, tplGrampusInferenceShow)
}

func GrampusNotebookNew(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(http.StatusOK, tplGrampusNotebookNew)
}

func GrampusTrainJobGPUNew(ctx *context.Context) {
	ctx.Data["IsCreate"] = true
	ctx.Data["PageIsCloudBrain"] = true

	ctx.HTML(http.StatusOK, tplGrampusTrainJobGPUNew)
}

func GrampusTrainJobNPUNew(ctx *context.Context) {
	ctx.Data["IsCreate"] = true
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(200, tplGrampusTrainJobNPUNew)
}

func GrampusTrainJobGCUNew(ctx *context.Context) {
	ctx.Data["IsCreate"] = true
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(http.StatusOK, tplGrampusTrainJobGCUNew)
}

func GrampusTrainJobIluvatarGPGPUNew(ctx *context.Context) {
	ctx.Data["IsCreate"] = true
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(http.StatusOK, tplGrampusTrainJobIluvatarGPGPUNew)
}

func GrampusTrainJobBIRENGPUNew(ctx *context.Context) {
	ctx.Data["IsCreate"] = true
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(http.StatusOK, tplGrampusTrainJobBIRENGPUNew)
}

func GrampusTrainJobDCUNew(ctx *context.Context) {
	ctx.Data["IsCreate"] = true
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(http.StatusOK, tplGrampusTrainJobDCUNew)
}

func grampusTrainJobNewDataPrepare(ctx *context.Context, processType string) error {
	ctx.Data["PageIsCloudBrain"] = true

	var displayJobName = cloudbrainService.GetDisplayJobName(ctx.User.Name)
	ctx.Data["display_job_name"] = displayJobName

	//get valid images
	if processType == grampus.ProcessorTypeNPU || processType == grampus.ProcessorTypeGCU {
		images, err := grampus.GetImages(processType, string(models.JobTypeTrain))
		if err != nil {
			log.Error("GetImages failed:", err.Error())
		} else {
			ctx.Data["images"] = images.Infos
		}
	}

	//prepare available specs
	if processType == grampus.ProcessorTypeNPU {
		prepareGrampusSpecs(ctx, models.NPU)
	} else if processType == grampus.ProcessorTypeGPU {
		prepareGrampusSpecs(ctx, models.GPU)
	} else if processType == grampus.ProcessorTypeGCU {
		prepareGrampusSpecs(ctx, models.GCU)
	}

	//get branches
	branches, _, err := ctx.Repo.GitRepo.GetBranches(0, 0)
	if err != nil {
		log.Error("GetBranches error:", err.Error())
	} else {
		ctx.Data["branches"] = branches
	}

	ctx.Data["branchName"] = ctx.Repo.BranchName

	if processType == grampus.ProcessorTypeGPU {
		ctx.Data["datasetType"] = models.TypeCloudBrainOne
		waitCount := cloudbrain.GetWaitingCloudbrainCount(models.TypeC2Net, models.GPUResource, models.JobTypeTrain)
		ctx.Data["WaitCount"] = waitCount
		NotStopTaskCount, _ := cloudbrainTask.GetNotFinalStatusTaskCount(ctx.User.ID, string(models.JobTypeTrain))
		ctx.Data["NotStopTaskCount"] = NotStopTaskCount
	} else if processType == grampus.ProcessorTypeNPU {
		ctx.Data["datasetType"] = models.TypeCloudBrainTwo
		waitCount := cloudbrain.GetWaitingCloudbrainCount(models.TypeC2Net, models.NPUResource, models.JobTypeTrain)
		ctx.Data["WaitCount"] = waitCount
		NotStopTaskCount, _ := cloudbrainTask.GetNotFinalStatusTaskCount(ctx.User.ID, string(models.JobTypeTrain))
		ctx.Data["NotStopTaskCount"] = NotStopTaskCount
		setGrampusMultiNodeIfConfigureMatch(ctx)
	} else if processType == grampus.ProcessorTypeGCU {
		ctx.Data["datasetType"] = models.TypeCloudBrainAll
		waitCount := cloudbrain.GetWaitingCloudbrainCount(models.TypeC2Net, models.GCUResource, models.JobTypeTrain)
		ctx.Data["WaitCount"] = waitCount
		NotStopTaskCount, _ := cloudbrainTask.GetNotFinalStatusTaskCount(ctx.User.ID, string(models.JobTypeTrain))
		ctx.Data["NotStopTaskCount"] = NotStopTaskCount
	}

	if ctx.Cloudbrain != nil {
		uuids, datasetNames := dataset.GetFilterDeletedAttachments(ctx.Cloudbrain.Uuid)
		ctx.Data["attachment"] = uuids
		ctx.Data["boot_file"] = ctx.Cloudbrain.BootFile
		ctx.Data["image_id"] = ctx.Cloudbrain.ImageID
		ctx.Data["run_para_list"] = ctx.Cloudbrain.Parameters
		ctx.Data["description"] = ctx.Cloudbrain.Description
		ctx.Data["branch_name"] = ctx.Cloudbrain.BranchName
		ctx.Data["engine_name"] = ctx.Cloudbrain.EngineName
		ctx.Data["work_server_number"] = ctx.Cloudbrain.WorkServerNumber
		if ctx.Cloudbrain.Image != "" {
			ctx.Data["image"] = ctx.Cloudbrain.Image
		} else {
			ctx.Data["image"] = ctx.Cloudbrain.EngineName
		}
		ctx.Data["dataset_name"] = datasetNames
		ctx.Data["model_name"] = ctx.Cloudbrain.ModelName

		ctx.Data["model_version"] = ctx.Cloudbrain.ModelVersion
		ctx.Data["ckpt_name"] = ctx.Cloudbrain.CkptName
		ctx.Data["model_id"] = ctx.Cloudbrain.ModelId
		ctx.Data["label_names"] = ctx.Cloudbrain.LabelName
		ctx.Data["pre_train_model_url"] = ctx.Cloudbrain.PreTrainModelUrl
		spec, _ := resource.GetCloudbrainSpec(ctx.Cloudbrain.ID)
		if spec != nil {
			ctx.Data["spec_id"] = spec.ID
		}

	}
	return nil
}

func prepareGrampusSpecs(ctx *context.Context, computeResource string, jobType ...models.JobType) {
	tempJobType := models.JobTypeTrain
	if len(jobType) > 0 {
		tempJobType = jobType[0]
	}
	noteBookSpecs, _ := resource.FindAvailableSpecsForNewRightWithRandomQueue(ctx.User.ID, models.FindSpecsOptions{
		JobType:         tempJobType,
		ComputeResource: computeResource,
		Cluster:         models.C2NetCluster,
	})
	ctx.Data["Specs"] = noteBookSpecs
}

func grampusParamCheckCreateTrainJob(form auth.CreateGrampusTrainJobForm) error {
	if !strings.HasSuffix(strings.TrimSpace(form.BootFile), ".py") {
		log.Error("the boot file(%s) must be a python file", form.BootFile)
		return errors.New("启动文件必须是python文件")
	}

	if form.BranchName == "" {
		log.Error("the branch must not be null!", form.BranchName)
		return errors.New("代码分支不能为空！")
	}

	return nil
}

func GrampusTrainJobGpuCreate(ctx *context.Context, form auth.CreateGrampusTrainJobForm) {
	ctx.Data["IsCreate"] = true
	grampusTrainJobGpuCreate(ctx, form)
}

func grampusTrainJobGpuCreate(ctx *context.Context, form auth.CreateGrampusTrainJobForm) {

	displayJobName := form.DisplayJobName
	jobName := util.ConvertDisplayJobNameToJobName(displayJobName)
	uuid := form.Attachment
	description := form.Description
	bootFile := strings.TrimSpace(form.BootFile)
	params := form.Params
	repo := ctx.Repo.Repository
	codeLocalPath := setting.JobPath + jobName + cloudbrain.CodeMountPath + "/"
	//codeMinioPath := setting.CBCodePathPrefix + jobName + cloudbrain.CodeMountPath + "/"
	branchName := form.BranchName
	image := strings.TrimSpace(form.Image)
	tpl := tplGrampusTrainJobGPUNew

	if !jobNamePattern.MatchString(displayJobName) {
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGPU)
		ctx.RenderWithErr(ctx.Tr("repo.cloudbrain_jobname_err"), tpl, &form)
		return
	}

	//check specification
	spec, err := resource.GetAndCheckSpec(ctx.User.ID, form.SpecId, models.FindSpecsOptions{
		JobType:         models.JobTypeTrain,
		ComputeResource: models.GPU,
		Cluster:         models.C2NetCluster,
	})
	if err != nil || spec == nil {
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGPU)
		ctx.RenderWithErr("Resource specification not available", tpl, &form)
		return
	}

	if !account.IsPointBalanceEnough(ctx.User.ID, models.PointDeductCondition{SpecUnitPrice: spec.UnitPrice}) {
		log.Error("point balance is not enough,userId=%d specId=%d", ctx.User.ID, spec.ID)
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGPU)
		ctx.RenderWithErr(ctx.Tr("points.insufficient_points_balance"), tplGrampusTrainJobGPUNew, &form)
		return
	}

	lockOperator, errMsg := cloudbrainService.Lock4CloudbrainCreation(&lock.LockContext{Task: &models.Cloudbrain{DisplayJobName: displayJobName, JobType: string(models.JobTypeTrain)}, User: ctx.User})
	defer func() {
		if lockOperator != nil {
			lockOperator.Unlock()
		}
	}()

	if errMsg != "" {
		log.Error("lock processed failed:%s", errMsg, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGPU)
		ctx.RenderWithErr(ctx.Tr(errMsg), tpl, &form)
		return
	}

	bootFileExist, err := ctx.Repo.FileExists(bootFile, branchName)
	if err != nil || !bootFileExist {
		log.Error("Get bootfile error:", err, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGPU)
		ctx.RenderWithErr(ctx.Tr("repo.cloudbrain_bootfile_err"), tpl, &form)
		return
	}

	//check count limit
	count, err := cloudbrainTask.GetNotFinalStatusTaskCount(ctx.User.ID, string(models.JobTypeTrain))
	if err != nil {
		log.Error("GetGrampusCountByUserID failed:%v", err, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGPU)
		ctx.RenderWithErr("system error", tpl, &form)
		return
	} else {
		if count >= 1 {
			log.Error("the user already has running or waiting task", ctx.Data["MsgID"])
			grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGPU)
			ctx.RenderWithErr("you have already a running or waiting task, can not create more", tpl, &form)
			return
		}
	}

	//check param
	if err := grampusParamCheckCreateTrainJob(form); err != nil {
		log.Error("paramCheckCreateTrainJob failed:(%v)", err, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGPU)
		ctx.RenderWithErr(err.Error(), tpl, &form)
		return
	}

	//check whether the task name in the project is duplicated
	tasks, err := models.GetCloudbrainsByDisplayJobName(repo.ID, string(models.JobTypeTrain), displayJobName)
	if err == nil {
		if len(tasks) != 0 {
			log.Error("the job name did already exist", ctx.Data["MsgID"])
			grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGPU)
			ctx.RenderWithErr("the job name did already exist", tpl, &form)
			return
		}
	} else {
		if !models.IsErrJobNotExist(err) {
			log.Error("system error, %v", err, ctx.Data["MsgID"])
			grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGPU)
			ctx.RenderWithErr("system error", tpl, &form)
			return
		}
	}

	//check dataset

	datasetInfos, datasetNames, err := models.GetDatasetInfo(uuid, models.GPU)
	if err != nil {
		log.Error("GetDatasetInfo failed: %v", err, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGPU)
		ctx.RenderWithErr(ctx.Tr("cloudbrain.error.dataset_select"), tpl, &form)
		return
	}

	//prepare code and out path
	_, err = ioutil.ReadDir(codeLocalPath)
	if err == nil {
		os.RemoveAll(codeLocalPath)
	}

	if err := downloadZipCode(ctx, codeLocalPath, branchName); err != nil {
		log.Error("downloadZipCode failed, server timed out: %s (%v)", repo.FullName(), err, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGPU)
		ctx.RenderWithErr(ctx.Tr("cloudbrain.load_code_failed"), tpl, &form)
		return
	}

	//todo: upload code (send to file_server todo this work?)
	//upload code
	if err := uploadCodeToMinio(codeLocalPath+"/", jobName, cloudbrain.CodeMountPath+"/"); err != nil {
		log.Error("Failed to uploadCodeToMinio: %s (%v)", repo.FullName(), err, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGPU)
		ctx.RenderWithErr(ctx.Tr("cloudbrain.load_code_failed"), tpl, &form)
		return
	}

	modelPath := setting.JobPath + jobName + cloudbrain.ModelMountPath + "/"
	if err := mkModelPath(modelPath); err != nil {
		log.Error("Failed to mkModelPath: %s (%v)", repo.FullName(), err, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGPU)
		ctx.RenderWithErr(ctx.Tr("cloudbrain.load_code_failed"), tpl, &form)
		return
	}

	/*if form.IsContinue { // zhisuan GPU 继续训练，将旧任务输出文件拷贝至新任务输出路径
		srcPath := setting.JobPath + form.PreJobName + cloudbrain.ModelMountPath + "/"
		destPath := setting.JobPath + jobName + cloudbrain.ModelMountPath + "/"
		err := MinioCopyResults(srcPath, destPath)
		if err != nil {
			log.Error("Grampus GPU: Copy Prev Task Result File failed:", err)
			grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGPU)
			ctx.RenderWithErr("Failed to copy output files from previous train job", tpl, &form)
			return
		}
	}*/

	//init model readme
	if err := uploadCodeToMinio(modelPath, jobName, cloudbrain.ModelMountPath+"/"); err != nil {
		log.Error("Failed to uploadCodeToMinio: %s (%v)", repo.FullName(), err, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGPU)
		ctx.RenderWithErr(ctx.Tr("cloudbrain.load_code_failed"), tpl, &form)
		return
	}

	var datasetRemotePath, allFileName string
	for _, datasetInfo := range datasetInfos {
		if datasetRemotePath == "" {
			datasetRemotePath = datasetInfo.DataLocalPath
			allFileName = datasetInfo.FullName
		} else {
			datasetRemotePath = datasetRemotePath + ";" + datasetInfo.DataLocalPath
			allFileName = allFileName + ";" + datasetInfo.FullName
		}

	}

	//prepare command
	preTrainModelPath := getPreTrainModelPath(form.PreTrainModelUrl, form.CkptName)

	command, err := generateCommand(repo.Name, grampus.ProcessorTypeGPU, bootFile, params, setting.CBCodePathPrefix+jobName+cloudbrain.ModelMountPath+"/", datasetNames, form.CkptName, "")
	if err != nil {
		log.Error("Failed to generateCommand: %s (%v)", displayJobName, err, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGPU)
		ctx.RenderWithErr("Create task failed, internal error", tpl, &form)
		return
	}

	commitID, _ := ctx.Repo.GitRepo.GetBranchCommitID(branchName)

	req := &grampus.GenerateTrainJobReq{
		JobName:         jobName,
		DisplayJobName:  displayJobName,
		ComputeResource: models.GPUResource,
		ProcessType:     grampus.ProcessorTypeGPU,
		Command:         command,
		ImageUrl:        image,
		Description:     description,
		BootFile:        bootFile,
		Uuid:            uuid,
		CommitID:        commitID,
		BranchName:      branchName,
		Params:          form.Params,
		EngineName:      image,
		DatasetNames:    datasetNames,
		DatasetInfos:    datasetInfos,

		IsLatestVersion:   modelarts.IsLatestVersion,
		VersionCount:      modelarts.VersionCountOne,
		WorkServerNumber:  1,
		Spec:              spec,
		PreTrainModelPath: preTrainModelPath,
	}

	if form.ModelName != "" { //使用预训练模型训练
		req.ModelName = form.ModelName
		req.LabelName = form.LabelName
		req.CkptName = form.CkptName
		req.ModelId = form.ModelId
		req.ModelVersion = form.ModelVersion
		req.PreTrainModelUrl = form.PreTrainModelUrl
		if !modelmanage.HasModelFileByModelId(req.ModelId, req.CkptName) { //使用预训练模型训练
			grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGPU)
			ctx.RenderWithErr(ctx.Tr("repo.train.manage.model_not_exist"), tpl, &form)
			return
		}
	}

	_, err = grampus.GenerateTrainJob(ctx, req)
	if err != nil {
		log.Error("GenerateTrainJob failed:%v", err.Error(), ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGPU)
		ctx.RenderWithErr(err.Error(), tpl, &form)
		return
	}
	ctx.Redirect(setting.AppSubURL + ctx.Repo.RepoLink + "/modelarts/train-job")
}

func grampusTrainJobGcuCreate(ctx *context.Context, form auth.CreateGrampusTrainJobForm) {

	displayJobName := form.DisplayJobName
	jobName := util.ConvertDisplayJobNameToJobName(displayJobName)
	uuid := form.Attachment
	description := form.Description
	bootFile := strings.TrimSpace(form.BootFile)
	params := form.Params
	repo := ctx.Repo.Repository
	codeLocalPath := setting.JobPath + jobName + cloudbrain.CodeMountPath + "/"
	//codeMinioPath := setting.CBCodePathPrefix + jobName + cloudbrain.CodeMountPath + "/"
	branchName := form.BranchName
	image := strings.TrimSpace(form.Image)
	imageId := strings.TrimSpace(form.ImageID)
	tpl := tplGrampusTrainJobGCUNew

	if !jobNamePattern.MatchString(displayJobName) {
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGCU)
		ctx.RenderWithErr(ctx.Tr("repo.cloudbrain_jobname_err"), tpl, &form)
		return
	}

	//check specification
	spec, err := resource.GetAndCheckSpec(ctx.User.ID, form.SpecId, models.FindSpecsOptions{
		JobType:         models.JobTypeTrain,
		ComputeResource: models.GCU,
		Cluster:         models.C2NetCluster,
	})
	if err != nil || spec == nil {
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGCU)
		ctx.RenderWithErr("Resource specification not available", tpl, &form)
		return
	}

	if !account.IsPointBalanceEnough(ctx.User.ID, models.PointDeductCondition{SpecUnitPrice: spec.UnitPrice}) {
		log.Error("point balance is not enough,userId=%d specId=%d", ctx.User.ID, spec.ID)
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGCU)
		ctx.RenderWithErr(ctx.Tr("points.insufficient_points_balance"), tpl, &form)
		return
	}

	lockOperator, errMsg := cloudbrainService.Lock4CloudbrainCreation(&lock.LockContext{Task: &models.Cloudbrain{DisplayJobName: displayJobName, JobType: string(models.JobTypeTrain)}, User: ctx.User})
	defer func() {
		if lockOperator != nil {
			lockOperator.Unlock()
		}
	}()

	if errMsg != "" {
		log.Error("lock processed failed:%s", errMsg, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGCU)
		ctx.RenderWithErr(ctx.Tr(errMsg), tpl, &form)
		return
	}

	bootFileExist, err := ctx.Repo.FileExists(bootFile, branchName)
	if err != nil || !bootFileExist {
		log.Error("Get bootfile error:", err, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGCU)
		ctx.RenderWithErr(ctx.Tr("repo.cloudbrain_bootfile_err"), tpl, &form)
		return
	}

	//check count limit
	count, err := cloudbrainTask.GetNotFinalStatusTaskCount(ctx.User.ID, string(models.JobTypeTrain))
	if err != nil {
		log.Error("GetGrampusCountByUserID failed:%v", err, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGCU)
		ctx.RenderWithErr("system error", tpl, &form)
		return
	} else {
		if count >= 1 {
			log.Error("the user already has running or waiting task", ctx.Data["MsgID"])
			grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGCU)
			ctx.RenderWithErr("you have already a running or waiting task, can not create more", tpl, &form)
			return
		}
	}

	//check param
	if err := grampusParamCheckCreateTrainJob(form); err != nil {
		log.Error("paramCheckCreateTrainJob failed:(%v)", err, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGCU)
		ctx.RenderWithErr(err.Error(), tpl, &form)
		return
	}

	//check whether the task name in the project is duplicated
	tasks, err := models.GetCloudbrainsByDisplayJobName(repo.ID, string(models.JobTypeTrain), displayJobName)
	if err == nil {
		if len(tasks) != 0 {
			log.Error("the job name did already exist", ctx.Data["MsgID"])
			grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGCU)
			ctx.RenderWithErr("the job name did already exist", tpl, &form)
			return
		}
	} else {
		if !models.IsErrJobNotExist(err) {
			log.Error("system error, %v", err, ctx.Data["MsgID"])
			grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGCU)
			ctx.RenderWithErr("system error", tpl, &form)
			return
		}
	}

	//check dataset

	datasetInfos, datasetNames, err := models.GetDatasetInfo(uuid, models.GCU)
	if err != nil {
		log.Error("GetDatasetInfo failed: %v", err, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGCU)
		ctx.RenderWithErr(ctx.Tr("cloudbrain.error.dataset_select"), tpl, &form)
		return
	}

	//prepare code and out path
	_, err = ioutil.ReadDir(codeLocalPath)
	if err == nil {
		os.RemoveAll(codeLocalPath)
	}

	if err := downloadZipCode(ctx, codeLocalPath, branchName); err != nil {
		log.Error("downloadZipCode failed, server timed out: %s (%v)", repo.FullName(), err, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGCU)
		ctx.RenderWithErr(ctx.Tr("cloudbrain.load_code_failed"), tpl, &form)
		return
	}

	//todo: upload code (send to file_server todo this work?)
	//upload code
	if err := uploadCodeToMinio(codeLocalPath+"/", jobName, cloudbrain.CodeMountPath+"/"); err != nil {
		log.Error("Failed to uploadCodeToMinio: %s (%v)", repo.FullName(), err, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGCU)
		ctx.RenderWithErr(ctx.Tr("cloudbrain.load_code_failed"), tpl, &form)
		return
	}

	modelPath := setting.JobPath + jobName + cloudbrain.ModelMountPath + "/"
	if err := mkModelPath(modelPath); err != nil {
		log.Error("Failed to mkModelPath: %s (%v)", repo.FullName(), err, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGCU)
		ctx.RenderWithErr(ctx.Tr("cloudbrain.load_code_failed"), tpl, &form)
		return
	}

	//init model readme
	if err := uploadCodeToMinio(modelPath, jobName, cloudbrain.ModelMountPath+"/"); err != nil {
		log.Error("Failed to uploadCodeToMinio: %s (%v)", repo.FullName(), err, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGCU)
		ctx.RenderWithErr(ctx.Tr("cloudbrain.load_code_failed"), tpl, &form)
		return
	}

	var datasetRemotePath, allFileName string
	for _, datasetInfo := range datasetInfos {
		if datasetRemotePath == "" {
			datasetRemotePath = datasetInfo.DataLocalPath
			allFileName = datasetInfo.FullName
		} else {
			datasetRemotePath = datasetRemotePath + ";" + datasetInfo.DataLocalPath
			allFileName = allFileName + ";" + datasetInfo.FullName
		}

	}

	//prepare command
	preTrainModelPath := getPreTrainModelPath(form.PreTrainModelUrl, form.CkptName)

	command, err := generateCommand(repo.Name, grampus.ProcessorTypeGCU, bootFile, params, setting.CBCodePathPrefix+jobName+cloudbrain.ModelMountPath+"/", datasetNames, form.CkptName, "")
	if err != nil {
		log.Error("Failed to generateCommand: %s (%v)", displayJobName, err, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGCU)
		ctx.RenderWithErr("Create task failed, internal error", tpl, &form)
		return
	}

	commitID, _ := ctx.Repo.GitRepo.GetBranchCommitID(branchName)

	req := &grampus.GenerateTrainJobReq{
		JobName:         jobName,
		DisplayJobName:  displayJobName,
		ComputeResource: models.GCUResource,
		ProcessType:     grampus.ProcessorTypeGCU,
		Command:         command,
		// ImageUrl:        image,
		ImageId:      imageId,
		Description:  description,
		BootFile:     bootFile,
		Uuid:         uuid,
		CommitID:     commitID,
		BranchName:   branchName,
		Params:       form.Params,
		EngineName:   image,
		DatasetNames: datasetNames,
		DatasetInfos: datasetInfos,

		IsLatestVersion:   modelarts.IsLatestVersion,
		VersionCount:      modelarts.VersionCountOne,
		WorkServerNumber:  1,
		Spec:              spec,
		PreTrainModelPath: preTrainModelPath,
	}

	if form.ModelName != "" { //使用预训练模型训练
		req.ModelName = form.ModelName
		req.LabelName = form.LabelName
		req.CkptName = form.CkptName
		req.ModelVersion = form.ModelVersion
		req.PreTrainModelUrl = form.PreTrainModelUrl
		req.ModelId = form.ModelId
		if !modelmanage.HasModelFileByModelId(req.ModelId, req.CkptName) { //使用预训练模型训练
			grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGCU)
			ctx.RenderWithErr(ctx.Tr("repo.train.manage.model_not_exist"), tpl, &form)
			return
		}
	}

	_, err = grampus.GenerateTrainJob(ctx, req)
	if err != nil {
		log.Error("GenerateTrainJob failed:%v", err.Error(), ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeGCU)
		ctx.RenderWithErr(err.Error(), tpl, &form)
		return
	}
	ctx.Redirect(setting.AppSubURL + ctx.Repo.RepoLink + "/modelarts/train-job")
}

func getPreTrainModelPath(pretrainModelDir string, fileName string) string {
	index := strings.Index(pretrainModelDir, "/")
	if index > 0 {
		filterBucket := pretrainModelDir[index+1:]
		return filterBucket + fileName
	} else {
		return ""
	}
}

func getPreTrainModelPaths(pretrainModelDir string, fileName []string) []string {
	var paths []string
	index := strings.Index(pretrainModelDir, "/")
	if index > 0 {
		filterBucket := pretrainModelDir[index+1:]
		for _, name := range fileName {
			paths = append(paths, filterBucket+name)
		}
	}
	return paths
}

func GrampusTrainJobVersionCreate(ctx *context.Context, form auth.CreateGrampusTrainJobForm) {
	ctx.Data["IsCreate"] = false

	if form.IsContinue {
		form.PreJobName = ctx.Cloudbrain.JobName
	}

	computeResource := ctx.Query("compute_resource")
	if computeResource == "" {
		computeResource = models.GCUResource
	}
	if computeResource == models.GPUResource {
		grampusTrainJobGpuCreate(ctx, form)
	} else if computeResource == models.NPUResource {
		grampusTrainJobNpuCreate(ctx, form)
	} else if computeResource == models.GCUResource {
		grampusTrainJobGcuCreate(ctx, form)
	} else {
		ctx.ServerError("resource error", errors.New("compute resource is not support"))
		return
	}

}

func GrampusTrainJobNpuCreate(ctx *context.Context, form auth.CreateGrampusTrainJobForm) {
	ctx.Data["IsCreate"] = true
	grampusTrainJobNpuCreate(ctx, form)
}

func grampusTrainJobNpuCreate(ctx *context.Context, form auth.CreateGrampusTrainJobForm) {

	displayJobName := form.DisplayJobName
	jobName := util.ConvertDisplayJobNameToJobName(displayJobName)
	uuid := form.Attachment
	description := form.Description
	bootFile := strings.TrimSpace(form.BootFile)
	params := form.Params
	repo := ctx.Repo.Repository
	codeLocalPath := setting.JobPath + jobName + modelarts.CodePath
	codeObsPath := grampus.JobPath + jobName + modelarts.CodePath
	branchName := form.BranchName
	isLatestVersion := modelarts.IsLatestVersion
	versionCount := modelarts.VersionCountOne
	engineName := form.EngineName
	tpl := tplGrampusTrainJobNPUNew

	if !jobNamePattern.MatchString(displayJobName) {
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeNPU)
		ctx.RenderWithErr(ctx.Tr("repo.cloudbrain_jobname_err"), tpl, &form)
		return
	}

	//check specification
	spec, err := resource.GetAndCheckSpec(ctx.User.ID, form.SpecId, models.FindSpecsOptions{
		JobType:         models.JobTypeTrain,
		ComputeResource: models.NPU,
		Cluster:         models.C2NetCluster,
	})
	if err != nil || spec == nil {
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeNPU)
		ctx.RenderWithErr("Resource specification not available", tpl, &form)
		return
	}
	if !account.IsPointBalanceEnough(ctx.User.ID, models.PointDeductCondition{SpecUnitPrice: spec.UnitPrice}) {
		log.Error("point balance is not enough,userId=%d specId=%d", ctx.User.ID, spec.ID)
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeNPU)
		ctx.RenderWithErr(ctx.Tr("points.insufficient_points_balance"), tplGrampusTrainJobNPUNew, &form)
		return
	}

	lockOperator, errMsg := cloudbrainService.Lock4CloudbrainCreation(&lock.LockContext{Task: &models.Cloudbrain{DisplayJobName: displayJobName, JobType: string(models.JobTypeTrain)}, User: ctx.User})
	defer func() {
		if lockOperator != nil {
			lockOperator.Unlock()
		}
	}()

	if errMsg != "" {
		log.Error("lock processed failed:%s", errMsg, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeNPU)
		ctx.RenderWithErr(ctx.Tr(errMsg), tpl, &form)
		return
	}

	errMsg = cloudbrainTask.CheckGrampusNPUMultiNode(ctx.User.ID, form.WorkServerNumber)
	if errMsg != "" {
		log.Error("multi node match failed:%s", errMsg, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeNPU)
		ctx.RenderWithErr(ctx.Tr(errMsg), tpl, &form)
		return
	}

	bootFileExist, err := ctx.Repo.FileExists(bootFile, branchName)
	if err != nil || !bootFileExist {
		log.Error("Get bootfile error:", err, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeNPU)
		ctx.RenderWithErr(ctx.Tr("repo.cloudbrain_bootfile_err"), tpl, &form)
		return
	}

	//check count limit
	count, err := cloudbrainTask.GetNotFinalStatusTaskCount(ctx.User.ID, string(models.JobTypeTrain))
	if err != nil {
		log.Error("GetGrampusCountByUserID failed:%v", err, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeNPU)
		ctx.RenderWithErr("system error", tpl, &form)
		return
	} else {
		if count >= 1 {
			log.Error("the user already has running or waiting task", ctx.Data["MsgID"])
			grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeNPU)
			ctx.RenderWithErr("you have already a running or waiting task, can not create more", tpl, &form)
			return
		}
	}

	//check param
	if err := grampusParamCheckCreateTrainJob(form); err != nil {
		log.Error("paramCheckCreateTrainJob failed:(%v)", err)
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeNPU)
		ctx.RenderWithErr(err.Error(), tpl, &form)
		return
	}

	//check whether the task name in the project is duplicated
	tasks, err := models.GetCloudbrainsByDisplayJobName(repo.ID, string(models.JobTypeTrain), displayJobName)
	if err == nil {
		if len(tasks) != 0 {
			log.Error("the job name did already exist", ctx.Data["MsgID"])
			grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeNPU)
			ctx.RenderWithErr("the job name did already exist", tpl, &form)
			return
		}
	} else {
		if !models.IsErrJobNotExist(err) {
			log.Error("system error, %v", err, ctx.Data["MsgID"])
			grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeNPU)
			ctx.RenderWithErr("system error", tpl, &form)
			return
		}
	}

	//check dataset
	datasetInfos, datasetNames, err := models.GetDatasetInfo(uuid, models.NPU)
	if err != nil {
		log.Error("GetDatasetInfo failed: %v", err, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeNPU)
		ctx.RenderWithErr(ctx.Tr("cloudbrain.error.dataset_select"), tpl, &form)
		return
	}

	//prepare code and out path
	_, err = ioutil.ReadDir(codeLocalPath)
	if err == nil {
		os.RemoveAll(codeLocalPath)
	}

	if err := downloadZipCode(ctx, codeLocalPath, branchName); err != nil {
		log.Error("downloadZipCode failed, server timed out: %s (%v)", repo.FullName(), err)
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeNPU)
		ctx.RenderWithErr(ctx.Tr("cloudbrain.load_code_failed"), tpl, &form)
		return
	}

	//todo: upload code (send to file_server todo this work?)
	/**if err := obsMkdir(setting.CodePathPrefix + jobName + modelarts.OutputPath); err != nil {
		log.Error("Failed to obsMkdir_output: %s (%v)", repo.FullName(), err)
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeNPU)
		ctx.RenderWithErr(ctx.Tr("cloudbrain.load_code_failed"), tpl, &form)
		return
	}*/

	if err := uploadCodeToObs(codeLocalPath, jobName, ""); err != nil {
		log.Error("Failed to uploadCodeToObs: %s (%v)", repo.FullName(), err)
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeNPU)
		ctx.RenderWithErr(ctx.Tr("cloudbrain.load_code_failed"), tpl, &form)
		return
	}

	/*if form.IsContinue { // qizhi NPU 继续训练，将旧任务输出文件拷贝至新任务输出路径
		srcPath := path.Join(setting.CodePathPrefix + form.PreJobName + modelarts.OutputPath)
		destPath := path.Join(setting.CodePathPrefix + jobName + modelarts.OutputPath)
		err := ObsCopyResults(srcPath, destPath)
		if err != nil {
			log.Error("Copy Prev Task Result File failed:", err)
			grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeNPU)
			ctx.RenderWithErr("Failed to copy output files from previous train job", tplModelArtsTrainJobVersionNew, &form)
			return
		}
	}*/

	var datasetRemotePath, allFileName string
	for _, datasetInfo := range datasetInfos {
		if datasetRemotePath == "" {
			datasetRemotePath = datasetInfo.DataLocalPath + "'" + datasetInfo.FullName + "'"
			allFileName = datasetInfo.FullName
		} else {
			datasetRemotePath = datasetRemotePath + ";" + datasetInfo.DataLocalPath + "'" + datasetInfo.FullName + "'"
			allFileName = allFileName + ";" + datasetInfo.FullName
		}

	}

	ckptNames := strings.Split(form.CkptName, ";")

	//prepare command
	preTrainModelPaths := getPreTrainModelPaths(form.PreTrainModelUrl, ckptNames)
	command, err := generateCommand(repo.Name, grampus.ProcessorTypeNPU, bootFile, params, setting.CodePathPrefix+jobName+modelarts.OutputPath, datasetNames, form.CkptName, grampus.GetNpuModelRemoteObsUrl(jobName))
	if err != nil {
		log.Error("Failed to generateCommand: %s (%v)", displayJobName, err, ctx.Data["MsgID"])
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeNPU)
		ctx.RenderWithErr("Create task failed, internal error", tpl, &form)
		return
	}

	commitID, _ := ctx.Repo.GitRepo.GetBranchCommitID(branchName)

	req := &grampus.GenerateTrainJobReq{
		JobName:           jobName,
		DisplayJobName:    displayJobName,
		ComputeResource:   models.NPUResource,
		ProcessType:       grampus.ProcessorTypeNPU,
		Command:           command,
		ImageId:           form.ImageID,
		Description:       description,
		CodeObsPath:       codeObsPath,
		BootFileUrl:       codeObsPath + bootFile,
		BootFile:          bootFile,
		WorkServerNumber:  form.WorkServerNumber,
		Uuid:              uuid,
		CommitID:          commitID,
		IsLatestVersion:   isLatestVersion,
		BranchName:        branchName,
		Params:            form.Params,
		EngineName:        engineName,
		VersionCount:      versionCount,
		TotalVersionCount: modelarts.TotalVersionCount,
		DatasetNames:      datasetNames,
		DatasetInfos:      datasetInfos,
		Spec:              spec,
		CodeName:          strings.ToLower(repo.Name),
	}
	if form.ModelName != "" { //使用预训练模型训练
		req.ModelName = form.ModelName
		req.LabelName = form.LabelName
		req.CkptName = form.CkptName
		req.ModelId = form.ModelId
		req.ModelVersion = form.ModelVersion
		req.PreTrainModelUrl = form.PreTrainModelUrl
		req.PreTrainModelPaths = preTrainModelPaths
		req.CkptNames = ckptNames
		for _, ckptName := range req.CkptNames {
			if !modelmanage.HasModelFileByModelId(req.ModelId, ckptName) { //使用预训练模型训练
				grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeNPU)
				ctx.RenderWithErr(ctx.Tr("repo.train.manage.model_not_exist"), tpl, &form)
				return
			}
		}
	}

	_, err = grampus.GenerateTrainJob(ctx, req)
	if err != nil {
		log.Error("GenerateTrainJob failed:%v", err.Error())
		grampusTrainJobNewDataPrepare(ctx, grampus.ProcessorTypeNPU)
		ctx.RenderWithErr(err.Error(), tpl, &form)
		return
	}
	ctx.Redirect(setting.AppSubURL + ctx.Repo.RepoLink + "/modelarts/train-job")
}

func setGrampusMultiNodeIfConfigureMatch(ctx *context.Context) {
	grampus.InitMultiNode()
	if grampus.MultiNodeConfig != nil {
		for _, info := range grampus.MultiNodeConfig.Info {
			if isInOrg, _ := models.IsOrganizationMemberByOrgName(info.Org, ctx.User.ID); isInOrg {
				ctx.Data["WorkNode"] = info.Node
				break
			}
		}
	}
}

func GetGrampusNotebook(ctx *context.APIContext) {
	var (
		err error
	)

	ID := ctx.Params(":id")
	job, err := models.GetCloudbrainByID(ID)
	if err != nil {
		ctx.NotFound("", err)
		log.Error("GetCloudbrainByID failed:", err)
		return
	}

	var jobAfter *models.Cloudbrain
	if job.IsNewAITask() {
		jobAfter, _ = ai_task.UpdateCloudbrain(job)
	} else {
		jobAfter, err = cloudbrainTask.SyncGrampusNotebookStatus(job)
	}

	aiCenterName := cloudbrainService.GetAiCenterShow(jobAfter.AiCenter, ctx.Context)

	if err != nil {
		ctx.NotFound(err)
		log.Error("Sync cloud brain one status failed:", err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"ID":             ID,
		"JobName":        jobAfter.JobName,
		"JobStatus":      jobAfter.Status,
		"DetailedStatus": jobAfter.DetailedStatus,
		"AiCenter":       aiCenterName,
		"CreatedTime":    jobAfter.CreatedUnix.Format("2006-01-02 15:04:05"),
		"CompletedTime":  jobAfter.UpdatedUnix.Format("2006-01-02 15:04:05"),
		"JobDuration":    jobAfter.TrainJobDuration,
	})
}

func GrampusStopJob(ctx *context.Context) {
	if res, isHandled, err := ai_task.HandleNewAITaskStop(ctx.Cloudbrain.ID); isHandled {
		if err != nil {
			log.Error("StopJob(%s) failed:%v", ctx.Cloudbrain.JobName, err, ctx.Data["msgID"])
			ctx.JSON(200, map[string]interface{}{
				"result_code": "-1",
				"error_msg":   ctx.Tr("cloudbrain.Stopped_failed"),
				"status":      "",
				"id":          ctx.Params(":id"),
				"StatusOK":    0,
			})
			return
		}
		ctx.JSON(200, map[string]interface{}{
			"result_code": "0",
			"error_msg":   "",
			"status":      res.Status,
			"id":          ctx.Params(":id"),
			"StatusOK":    0,
		})
		return
	}
	cloudbrainTask.GrampusStopJob(ctx)
}

func GrampusNotebookDel(ctx *context.Context) {
	var listType = ctx.Query("listType")

	if isHandled, err := ai_task.HandleNewAITaskDelete(ctx.Cloudbrain.ID); isHandled {
		if err != nil {
			log.Error("DeleteJob(%s) failed:%v", ctx.Cloudbrain.JobName, err, ctx.Data["msgID"])
			ctx.ServerError(err.Error(), err)
			return
		}
	} else {
		if err := cloudbrainTask.DeleteGrampusJob(ctx); err != nil {
			log.Error("deleteGrampusJob failed: %v", err, ctx.Data["msgID"])
			ctx.ServerError(err.Error(), err)
			return
		}
	}

	var isAdminPage = ctx.Query("isadminpage")
	var isHomePage = ctx.Query("ishomepage")
	if ctx.IsUserSiteAdmin() && isAdminPage == "true" {
		ctx.Redirect(setting.AppSubURL + "/admin" + "/cloudbrains")
	} else if isHomePage == "true" {
		ctx.Redirect(setting.AppSubURL + "/cloudbrains")
	} else {
		ctx.Redirect(setting.AppSubURL + ctx.Repo.RepoLink + "/debugjob?debugListType=" + listType)
	}
}

func GrampusTrainJobDel(ctx *context.Context) {
	var listType = ctx.Query("listType")
	if isHandled, err := ai_task.HandleNewAITaskDelete(ctx.Cloudbrain.ID); isHandled {
		if err != nil {
			log.Error("DeleteJob(%s) failed:%v", ctx.Cloudbrain.JobName, err, ctx.Data["msgID"])
			ctx.ServerError(err.Error(), err)
			return
		}
	} else if err := cloudbrainTask.DeleteGrampusJob(ctx); err != nil {
		log.Error("deleteGrampusJob failed: %v", err, ctx.Data["msgID"])
		ctx.ServerError(err.Error(), err)
		return
	}

	var isAdminPage = ctx.Query("isadminpage")
	var isHomePage = ctx.Query("ishomepage")
	if ctx.IsUserSiteAdmin() && isAdminPage == "true" {
		ctx.Redirect(setting.AppSubURL + "/admin" + "/cloudbrains")
	} else if isHomePage == "true" {
		ctx.Redirect(setting.AppSubURL + "/cloudbrains")
	} else {
		ctx.Redirect(setting.AppSubURL + ctx.Repo.RepoLink + "/modelarts/train-job?listType=" + listType)
	}
}

type NotebookDataset struct {
	DatasetUrl string `json:"dataset_url"`
}

func GrampusNotebookShow(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(http.StatusOK, tplGrampusNotebookShow)
	return
}

func GrampusTrainJobShow(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.HTML(http.StatusOK, tplGrampusTrainJobShow)
	/*
		var task *models.Cloudbrain
		task, err := models.GetCloudbrainByJobIDWithDeleted(ctx.Params(":jobid"))
		if err != nil {
			log.Error("GetCloudbrainByJobID failed:" + err.Error())
			ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
			return
		}
		task.ContainerIp = ""
		task.User, _ = models.GetUserByID(task.UserID)
		if task.DeletedAt.IsZero() { //normal record
			result, err := grampus.GetJob(task.JobID)
			if err != nil {
				log.Error("GetJob failed:" + err.Error())
				ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
				return
			}

			if result != nil {
				if len(result.JobInfo.Tasks[0].CenterID) == 1 && len(result.JobInfo.Tasks[0].CenterName) == 1 {
					task.AiCenter = result.JobInfo.Tasks[0].CenterID[0] + "+" + result.JobInfo.Tasks[0].CenterName[0]
				}
				oldStatus := task.Status
				task.Status = grampus.TransTrainJobStatus(result.JobInfo.Status)
				if task.Status != oldStatus || task.Status == models.GrampusStatusRunning {
					task.Duration = result.JobInfo.RunSec
					if task.Duration < 0 {
						task.Duration = 0
					}
					task.TrainJobDuration = models.ConvertDurationToStr(task.Duration)

					if task.StartTime == 0 && result.JobInfo.StartedAt > 0 {
						task.StartTime = timeutil.TimeStamp(result.JobInfo.StartedAt)
					}
					if task.EndTime == 0 && models.IsTrainJobTerminal(task.Status) && task.StartTime > 0 {
						task.EndTime = task.StartTime.Add(task.Duration)
					}
					task.CorrectCreateUnix()
					if oldStatus != task.Status {
						notification.NotifyChangeCloudbrainStatus(task, oldStatus)
					}
				}
				err = models.UpdateJob(task)
				if err != nil {
					log.Error("UpdateJob failed:" + err.Error())
				}
			}
		}

		if len(task.Parameters) > 0 {
			var parameters models.Parameters
			err := json.Unmarshal([]byte(task.Parameters), &parameters)
			if err != nil {
				log.Error("Failed to Unmarshal Parameters: %s (%v)", task.Parameters, err)
				ctx.ServerError("system error", err)
				return
			}

			if len(parameters.Parameter) > 0 {
				paramTemp := ""
				for _, Parameter := range parameters.Parameter {
					param := Parameter.Label + " = " + Parameter.Value + "; "
					paramTemp = paramTemp + param
				}
				task.Parameters = paramTemp[:len(paramTemp)-2]
			} else {
				task.Parameters = ""
			}
		}

		taskList := make([]*models.Cloudbrain, 0)
		taskList = append(taskList, task)
		prepareSpec4Show(ctx, task)

		ctx.Data["version_list_task"] = taskList
		ctx.Data["datasetDownload"] = GetCloudBrainDataSetInfo(task.Uuid, task.DatasetName, false)
		ctx.Data["canDownload"] = cloudbrain.CanDownloadJob(ctx, task)
		ctx.Data["displayJobName"] = task.DisplayJobName
		ctx.Data["canReschedule"] = cloudbrain.CanDeleteJob(ctx, task)

		ctx.Data["ai_center"] = cloudbrainService.GetAiCenterShow(task.AiCenter, ctx)

		ctx.HTML(http.StatusOK, tplGrampusTrainJobShow)
	*/
}

func GrampusDownloadLog(ctx *context.Context) {
	jobID := ctx.Params(":jobid")
	job, err := models.GetCloudbrainByJobID(jobID)
	if err != nil {
		log.Error("GetCloudbrainByJobID failed: %v", err, ctx.Data["MsgID"])
		ctx.ServerError(err.Error(), err)
		return
	}
	fileName := job.JobName + "-log.txt"

	nodeIdStr := ctx.Params(":nodeId")
	var content string
	if nodeIdStr != "" {
		nodeId, _ := strconv.Atoi(nodeIdStr)
		fileName = job.JobName + "-" + strconv.Itoa(nodeId+1) + "-log.txt"
		if job.WorkServerNumber < 1 || nodeId > job.WorkServerNumber-1 {
			ctx.NotFound("query parameter is wrong", nil)
			return
		}
		content, err = grampus.GetTrainJobLog(job.JobID, nodeId)
	} else {
		content, err = grampus.GetTrainJobLog(job.JobID)
	}
	if err != nil {
		log.Error("GetTrainJobLog failed: %v", err, ctx.Data["MsgID"])
		content = ""
	}

	ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	ctx.Resp.Header().Set("Content-Type", "application/octet-stream")
	var b []byte = []byte(content)
	ctx.Resp.Write(b)
}

func GrampusGetLog(ctx *context.Context) {
	jobID := ctx.Params(":jobid")
	job, err := models.GetCloudbrainByJobID(jobID)
	if err != nil {
		log.Error("GetCloudbrainByJobID failed: %v", err, ctx.Data["MsgID"])
		ctx.ServerError(err.Error(), err)
		return
	}

	result, err := grampus.GetJob(jobID)
	if err != nil {
		log.Error("GetJob(%s) failed:%v", job.JobName, err)
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"JobName":        job.JobName,
			"Content":        "",
			"CanLogDownload": false,
		})
		return
	}
	exitDiagnostics := ""
	if result != nil {
		exitDiagnostics = result.ExitDiagnostics
	}

	nodeIdStr := ctx.Params(":nodeId")
	var content string
	if nodeIdStr != "" {
		nodeId, _ := strconv.Atoi(nodeIdStr)
		if job.WorkServerNumber < 1 || nodeId > job.WorkServerNumber-1 {
			ctx.NotFound("query parameter is wrong", nil)
			return
		}
		content, err = grampus.GetTrainJobLog(job.JobID, nodeId)
	} else {
		content, err = grampus.GetTrainJobLog(job.JobID)
	}

	if err != nil {
		log.Error("GetTrainJobLog failed: %v", err, ctx.Data["MsgID"])
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"JobName":        job.JobName,
			"Content":        exitDiagnostics,
			"CanLogDownload": false,
		})
		return
	}

	if result != nil {
		job.Status = grampus.TransTrainJobStatus(result.JobInfo.Status)
		if job.Status == models.GrampusStatusFailed {
			content = content + "\n" + exitDiagnostics
		}
	}

	canLogDownload := err == nil && job.IsUserHasRight(ctx.User)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"JobName":        job.JobName,
		"Content":        content,
		"CanLogDownload": canLogDownload,
	})

	return
}

func GrampusMetrics(ctx *context.Context) {
	jobID := ctx.Params(":jobid")
	job, err := models.GetCloudbrainByJobID(jobID)
	if err != nil {
		log.Error("GetCloudbrainByJobID failed: %v", err, ctx.Data["MsgID"])
		ctx.ServerError(err.Error(), err)
		return
	}
	var result models.NewModelArtsMetricStatisticResult
	if job.IsNPUTask() {
		nodeIdStr := ctx.Params(":nodeId")
		if nodeIdStr != "" {
			nodeId, _ := strconv.Atoi(nodeIdStr)
			if job.WorkServerNumber < 1 || nodeId > job.WorkServerNumber-1 {
				ctx.NotFound("query parameter is wrong", nil)
				return
			}
			result, err = grampus.GetGrampusMetrics(job.JobID, 0, 0, nodeId)
		} else {
			result, err = grampus.GetGrampusMetrics(job.JobID, 0, 0)
		}
	} else if job.IsGPUTask() {
		startTime := int64(job.StartTime)
		if startTime == 0 {
			startTime = time.Now().Unix() - 30*60
		}
		endTime := int64(job.EndTime)
		if endTime == 0 {
			endTime = time.Now().Unix()
		}

		result, err = grampus.GetGrampusMetrics(job.JobID, startTime, endTime)
	}

	if err != nil {
		log.Error("GetTrainJobLog failed: %v", err, ctx.Data["MsgID"])
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"JobID":       jobID,
		"Interval":    0,
		"MetricsInfo": result.MetricsInfo,
	})

	return
}
func GrampusDebugJobEvents(ctx *context.Context) {
	ID := ctx.Params(":id")
	job, err := models.GetCloudbrainByID(ID)
	if err != nil {
		log.Error("GetCloudbrainByID failed: %v", err, ctx.Data["MsgID"])
		ctx.ServerError(err.Error(), err)
		return
	}

	var grampusJobEvent models.GrampusJobEvents
	getJobResult, err := grampus.GetJob(job.JobID)
	if err != nil {
		log.Error("GetJob(%s) failed:%v", job.JobName, err)
	}
	if getJobResult != nil {
		grampusJobEvent.Reason = getJobResult.ExitDiagnostics
	}
	jobExitEvent := models.GetGrampusDebugJobEventsResponse{
		NotebookEvents: []models.GrampusJobEvents{
			grampusJobEvent,
		},
	}

	result, err := grampus.GetDebugJobEvents(job.JobID)
	if err != nil {
		log.Error("GetDebugJobEvents failed: %v", err, ctx.Data["MsgID"])
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"JobID":     ID,
			"JobEvents": jobExitEvent,
		})
		return
	}
	result.NotebookEvents = append(result.NotebookEvents, grampusJobEvent)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"JobID":     ID,
		"JobEvents": result.NotebookEvents,
	})
	return
}

func GrampusTrainJobEvents(ctx *context.Context) {
	jobID := ctx.Params(":jobid")
	job, err := models.GetCloudbrainByJobID(jobID)
	if err != nil {
		log.Error("GetCloudbrainByJobID failed: %v", err, ctx.Data["MsgID"])
		ctx.ServerError(err.Error(), err)
		return
	}

	result, err := grampus.GetTrainJobEvents(job.JobID)
	if err != nil {
		log.Error("GetJobEvents failed: %v", err, ctx.Data["MsgID"])
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"JobID":     jobID,
		"JobEvents": result.JobEvents,
	})

	return
}

func generateCommand(repoName, processorType, bootFile, paramSrc, outputRemotePath, datasetName, pretrainModelFileName, modelRemoteObsUrl string) (string, error) {
	var command string

	//prepare
	workDir := grampus.NpuWorkDir
	if processorType == grampus.ProcessorTypeNPU {
		command += "pwd;cd " + workDir + grampus.CommandPrepareScriptNpu
	} else if processorType == grampus.ProcessorTypeGPU {
		workDir = grampus.GpuWorkDir
		command += "pwd; cd " + workDir + fmt.Sprintf(grampus.CommandPrepareScriptGpu)
	} else if processorType == grampus.ProcessorTypeGCU {
		workDir = grampus.GcuWorkDir
		command += "pwd; cd " + workDir + fmt.Sprintf(grampus.CommandPrepareScriptGpu)
	}
	//unzip code & dataset
	if processorType == grampus.ProcessorTypeNPU {
		//no need to process
	} else if processorType == grampus.ProcessorTypeGPU {
		unZipDatasetCommand := cloudbrainTask.GenerateDatasetUnzipCommand(datasetName)
		commandUnzip := "cd " + workDir + "code;echo \"start unzip code\";unzip -q master.zip;echo \"start to unzip dataset\";cd " + workDir + "dataset;" + unZipDatasetCommand
		command += commandUnzip
	} else if processorType == grampus.ProcessorTypeGCU {
		unZipDatasetCommand := cloudbrainTask.GenerateDatasetUnzipCommand(datasetName)
		commandUnzip := "cd " + workDir + "code;echo \"start unzip code\";unzip -q master.zip;echo \"start to unzip dataset\";cd " + workDir + "dataset;" + unZipDatasetCommand
		command += commandUnzip
	}

	command += "echo \"unzip finished;start to exec code;\";"

	// set export
	var commandExport string
	if processorType == grampus.ProcessorTypeNPU {
		commandExport = "export bucket=" + setting.Bucket + " && export remote_path=" + outputRemotePath + ";"
	} else if processorType == grampus.ProcessorTypeGPU {
		commandExport = "export env=" + setting.Grampus.Env + " && export remote_path=" + outputRemotePath + ";"
	} else if processorType == grampus.ProcessorTypeGCU {
		commandExport = "export env=" + setting.Grampus.Env + " && export remote_path=" + outputRemotePath + ";"
	}

	command += commandExport

	//exec code
	var parameters models.Parameters
	var paramCode string

	if len(paramSrc) != 0 {
		err := json.Unmarshal([]byte(paramSrc), &parameters)
		if err != nil {
			log.Error("Failed to Unmarshal params: %s (%v)", paramSrc, err)
			return command, err
		}

		for _, parameter := range parameters.Parameter {
			paramCode += " --" + parameter.Label + "=" + parameter.Value
		}
	}

	var commandCode string
	if processorType == grampus.ProcessorTypeNPU {
		paramCode += " --model_url=" + modelRemoteObsUrl
		commandCode = "source /home/ma-user/.bashrc;python /home/ma-user/davinci/train/davincirun.py python /home/ma-user/openi.py " + paramCode + ";"
	} else if processorType == grampus.ProcessorTypeGPU {
		if pretrainModelFileName != "" {
			paramCode += " --ckpt_url" + "='" + workDir + "pretrainmodel/" + pretrainModelFileName + "'"
		}
		commandCode = "cd " + workDir + "code/" + strings.ToLower(repoName) + ";python " + bootFile + paramCode + ";"
	} else if processorType == grampus.ProcessorTypeGCU {
		if pretrainModelFileName != "" {
			paramCode += " --ckpt_url" + "='" + workDir + "pretrainmodel/" + pretrainModelFileName + "'"
		}
		commandCode = "cd " + workDir + "code/" + strings.ToLower(repoName) + ";python3 " + bootFile + paramCode + ";"
	}

	command += commandCode

	//get exec result
	commandGetRes := "result=$?;"
	command += commandGetRes

	////upload models
	//if processorType == grampus.ProcessorTypeNPU {
	//	// no need to upload
	//} else if processorType == grampus.ProcessorTypeGPU {
	//	commandUpload := "cd " + workDir + setting.Grampus.SyncScriptProject + "/;./uploader_for_gpu " + setting.Grampus.Env + " " + outputRemotePath + " " + workDir + "output/;"
	//	command += commandUpload
	//}

	//check exec result
	commandCheckRes := "bash -c \"[[ $result -eq 0 ]] && exit 0 || exit -1\""
	command += commandCheckRes

	return command, nil
}

func processPretrainModelParameter(pretrainModelPath string, pretrainModelFileName string, commandDownload string) string {
	commandDownloadTemp := commandDownload
	if pretrainModelPath != "" {
		commandDownloadTemp += " '" + pretrainModelPath + "' '" + pretrainModelFileName + "'"
	}
	commandDownloadTemp += ";"
	return commandDownloadTemp
}

func downloadZipCode(ctx *context.Context, codePath, branchName string) error {
	archiveType := git.ZIP
	archivePath := codePath

	if !com.IsDir(archivePath) {
		if err := os.MkdirAll(archivePath, os.ModePerm); err != nil {
			log.Error("MkdirAll failed:" + err.Error())
			return err
		}
	}

	// Get corresponding commit.
	var (
		commit *git.Commit
		err    error
	)

	gitRepo := ctx.Repo.GitRepo
	if err != nil {
		log.Error("OpenRepository failed:" + err.Error())
		return err
	}

	if gitRepo.IsBranchExist(branchName) {
		commit, err = gitRepo.GetBranchCommit(branchName)
		if err != nil {
			log.Error("GetBranchCommit failed:" + err.Error())
			return err
		}
	} else {
		log.Error("the branch is not exist: " + branchName)
		return fmt.Errorf("The branch does not exist.")
	}

	archivePath = path.Join(archivePath, grampus.CodeArchiveName)
	if !com.IsFile(archivePath) {
		if err := commit.CreateArchive(archivePath, git.CreateArchiveOpts{
			Format: archiveType,
			Prefix: setting.Repository.PrefixArchiveFiles,
		}); err != nil {
			log.Error("CreateArchive failed:" + err.Error())
			return err
		}
	}

	return nil
}
func HandleTaskWithAiCenter(ctx *context.Context) {
	log.Info("HandleTaskWithAiCenter start")
	updateCounts := 0
	cloudBrains, err := models.GetC2NetWithAiCenterWrongJob()
	if err != nil {
		log.Error("GetC2NetWithAiCenterWrongJob failed:" + err.Error())
		return
	}
	if len(cloudBrains) == 0 {
		log.Info("HandleC2NetWithAiCenterWrongJob:no task need handle")
		return
	}
	cloudBrainCounts := len(cloudBrains)
	for _, task := range cloudBrains {
		result, err := grampus.GetJob(task.JobID)
		if err != nil {
			log.Error("GetJob failed:" + err.Error())
			continue
		}
		if len(result.JobInfo.Tasks) != 0 {
			if len(result.JobInfo.Tasks[0].CenterID) == 1 && len(result.JobInfo.Tasks[0].CenterName) == 1 {
				task.AiCenter = result.JobInfo.Tasks[0].CenterID[0] + "+" + result.JobInfo.Tasks[0].CenterName[0]
			}
			err = models.UpdateJob(task)
			if err != nil {
				log.Error("UpdateJob failed:" + err.Error())
			}
			updateCounts++
		}
	}
	r := make(map[string]interface{}, 0)
	r["cloudBrainCounts"] = cloudBrainCounts
	r["updateCounts"] = updateCounts
	ctx.JSON(http.StatusOK, response.SuccessWithData(r))
}

func GrampusNotebookDebug(ctx *context.Context) {

	err := cloudbrainTask.GrampusNotebookDebug(ctx)
	if err != nil {
		ctx.NotFound(err.Error(), nil)
	}

}

func getImageComputeResource(processType string) string {
	tail := strings.LastIndex(processType, "/")
	if tail > 0 {
		return strings.ToUpper(processType[tail+1:])
	}
	return strings.ToUpper(processType)
}
func GrampusCommitImageShow(ctx *context.Context) {
	ctx.Data["PageIsCloudBrain"] = true
	ctx.Data["Type"] = ctx.Cloudbrain.Type
	ctx.Data["AiCenter"] = ctx.Cloudbrain.AiCenter
	ctx.Data["ComputeResource"] = getImageComputeResource(ctx.Cloudbrain.ComputeResource)
	ctx.HTML(200, tplGrampusNotebookGPUImageShow)
}

func GrampusCommitImage(ctx *context.Context, form auth.CommitImageGrampusForm) {

	if !GrampusNamePattern.MatchString(form.Tag) {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("repo.images.name_format_err")))
		return
	}

	if utf8.RuneCountInString(form.Description) > 1000 {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("dataset.description_format_err", "1000")))
		return
	}

	validTopics, errMessage := checkTopics(form.Topics)
	if errMessage != "" {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr(errMessage)))
		return
	}

	imageInfo, _ := models.GetImageByImageId(form.ImageId)

	if ctx.Cloudbrain.JobType == string(models.JobTypeComfyuiExperience) || ctx.Cloudbrain.JobType == string(models.JobTypeOnlineInference) {
		err := grampus.CommitInferJobImage(ctx.Cloudbrain, ctx.Cloudbrain.ComputeResource, ctx.Cloudbrain.GetAiCenter(), models.CommitGrampusImageParams{
			CommitImageGrampusParams: models.CommitImageGrampusParams{
				ImageName:    setting.Grampus.GPUImageCommonName,
				Description:  form.Description,
				ImageVersion: form.Tag,
				TaskName:     "task0",
				TrainType:    imageInfo.TrainType,
			}, IsPrivate: form.IsPrivate,
			CloudBrainType:         form.Type,
			Topics:                 validTopics,
			UID:                    ctx.User.ID,
			Framework:              form.Framework,
			FrameworkVersion:       form.FrameworkVersion,
			PythonVersion:          form.PythonVersion,
			CudaVersion:            form.CudaVersion,
			CannVersion:            form.CannVersion,
			DTKVersion:             form.DTKVersion,
			OperationSystem:        form.OperationSystem,
			OperationSystemVersion: form.OperationSystemVersion,
			ThirdPackages:          form.ThirdPackages,
			ComputeResource:        form.ComputeResource,
		}, ctx.User)
		if err != nil {
			log.Error("CommitImage(%s) failed:%v", ctx.Cloudbrain.JobName, err.Error(), ctx.Data["msgID"])
			if models.IsErrImageTagExist(err) || strings.Contains(err.Error(), "Image already exists") || strings.Contains(err.Error(), "image exists") {
				ctx.JSON(200, models.BaseErrorMessage(ctx.Tr("repo.image_exist")))

			} else if models.IsErrorImageCommitting(err) {
				ctx.JSON(200, models.BaseErrorMessage(ctx.Tr("repo.image_committing")))
			} else if isOver20GError(err) {
				ctx.JSON(200, models.BaseErrorMessage(ctx.Tr("repo.image_over_20g")))
			} else {
				ctx.JSON(200, models.BaseErrorMessage(ctx.Tr("repo.image_commit_fail")))
			}

			return
		}
	} else {
		err := grampus.CommitImage(ctx.Cloudbrain, ctx.Cloudbrain.ComputeResource, ctx.Cloudbrain.GetAiCenter(), models.CommitGrampusImageParams{
			CommitImageGrampusParams: models.CommitImageGrampusParams{
				ImageName:    setting.Grampus.GPUImageCommonName,
				Description:  form.Description,
				ImageVersion: form.Tag,
				TaskName:     "task0",
				TrainType:    imageInfo.TrainType,
			}, IsPrivate: form.IsPrivate,
			CloudBrainType:         form.Type,
			Topics:                 validTopics,
			UID:                    ctx.User.ID,
			Framework:              form.Framework,
			FrameworkVersion:       form.FrameworkVersion,
			PythonVersion:          form.PythonVersion,
			CudaVersion:            form.CudaVersion,
			CannVersion:            form.CannVersion,
			DTKVersion:             form.DTKVersion,
			OperationSystem:        form.OperationSystem,
			OperationSystemVersion: form.OperationSystemVersion,
			ThirdPackages:          form.ThirdPackages,
			ComputeResource:        form.ComputeResource,
		}, ctx.User)
		if err != nil {
			log.Error("CommitImage(%s) failed:%v", ctx.Cloudbrain.JobName, err.Error(), ctx.Data["msgID"])
			if models.IsErrImageTagExist(err) || strings.Contains(err.Error(), "Image already exists") || strings.Contains(err.Error(), "image exists") {
				ctx.JSON(200, models.BaseErrorMessage(ctx.Tr("repo.image_exist")))

			} else if models.IsErrorImageCommitting(err) {
				ctx.JSON(200, models.BaseErrorMessage(ctx.Tr("repo.image_committing")))
			} else if isOver20GError(err) {
				ctx.JSON(200, models.BaseErrorMessage(ctx.Tr("repo.image_over_20g")))
			} else {
				ctx.JSON(200, models.BaseErrorMessage(ctx.Tr("repo.image_commit_fail")))
			}

			return
		}
	}
	ctx.JSON(200, models.BaseOKMessage)
}
