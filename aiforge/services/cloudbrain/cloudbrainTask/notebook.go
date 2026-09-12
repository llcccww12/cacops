package cloudbrainTask

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"code.gitea.io/gitea/entity"

	"code.gitea.io/gitea/modules/grampus"
	"code.gitea.io/gitea/modules/notification"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/services/cloudbrain/modelmanage"

	"code.gitea.io/gitea/services/lock"

	"code.gitea.io/gitea/modules/notebook"

	"code.gitea.io/gitea/modules/modelarts"
	"code.gitea.io/gitea/modules/modelarts_cd"

	"code.gitea.io/gitea/modules/git"

	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/services/cloudbrain/resource"
	"code.gitea.io/gitea/services/reward/point/account"

	"code.gitea.io/gitea/modules/setting"
	cloudbrainService "code.gitea.io/gitea/services/cloudbrain"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	api "code.gitea.io/gitea/modules/structs"
)

const NoteBookExtension = ".ipynb"
const CPUType = 0
const GPUType = 1
const NPUType = 2
const CharacterLength = 2550

var notebookDirInvalidChar = []string{"<", ">", "'", "\"", ";", "\\", "`", "=", "#", "$", "%", "^", "(", ")"}

func FileNotebookCreate(ctx *context.Context, option api.CreateFileNotebookJobOption) {

	if ctx.Written() {
		return
	}

	if path.Ext(option.File) != NoteBookExtension {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.notebook_select_wrong")))
		return
	}
	if len(GetBootFile(option.File, option.OwnerName, option.ProjectName, option.BranchName)) > CharacterLength {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.notebook_path_too_long")))
		return
	}
	if len(option.BranchName) > CharacterLength {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.notebook_branch_name_too_long")))
		return
	}
	if branchNameContainsNotebookInvalidChar(option.BranchName) {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.notebook_branch_name_not_support")))
		return
	}

	var err error
	isNotebookFileExist, _ := isNoteBookFileExist(ctx, option)
	if !isNotebookFileExist {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.notebook_file_not_exist")))
		return
	}

	sourceRepo, err := models.GetRepositoryByOwnerAndName(option.OwnerName, option.ProjectName)
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.notebook_file_not_exist")))
		return
	}

	permission, err := models.GetUserRepoPermission(sourceRepo, ctx.User)
	if err != nil {
		log.Error("Get permission failed", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.notebook_file_no_right")))
		return
	}

	if !permission.CanRead(models.UnitTypeCode) {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.notebook_file_no_right")))
		return
	}

	noteBook, _ := models.GetWaitOrRunFileNotebookByRepo(sourceRepo.ID, ctx.User.ID, getComputeResource(option.Type), option.Image)
	if noteBook != nil {

		if isNotebookSpecMath(option, noteBook) && isImageMatch(option.Image, noteBook) {
			if !isRepoBranchMatch(option, noteBook) {
				err = downloadCode(sourceRepo, getCodePath(noteBook.JobName, sourceRepo, option.BranchName), option.BranchName)
				if err != nil {
					log.Error("download code failed", err)
					if !strings.Contains(err.Error(), "already exists and is not an empty directory") {
						ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("cloudbrain.load_code_failed")))
						return
					}
				}
			}
			if !isRepoFileMatch(option, noteBook) {
				if len(noteBook.BootFile)+len(GetBootFile(option.File, option.OwnerName, option.ProjectName, option.BranchName))+1 <= CharacterLength {
					noteBook.BootFile += ";" + GetBootFile(option.File, option.OwnerName, option.ProjectName, option.BranchName)
				} else {
					ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.notebook_path_too_long")))
					return
				}
				if len(noteBook.BranchName)+len(option.BranchName)+1 <= CharacterLength {
					noteBook.BranchName += ";" + option.BranchName
				} else {
					ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.notebook_branch_name_too_long")))
					return
				}

				if len(noteBook.Description)+len(getDescription(option))+1 <= CharacterLength {
					noteBook.Description += ";" + getDescription(option)
				}

				err := models.UpdateJob(noteBook)
				if err != nil {
					log.Error("GenerateNotebook2 failed, %v", err, ctx.Data["MsgID"])
					ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
					return
				}
			}

			ctx.JSON(http.StatusOK, models.BaseMessageApi{
				Code:    0,
				Message: fmt.Sprint(noteBook.ID),
			})
			return
		}

	}

	//if option.Type <= GPUType {
	//	cloudBrainFileNoteBookCreate(ctx, option, repo, sourceRepo)
	//} else {
	//	modelartsFileNoteBookCreate(ctx, option, repo, sourceRepo, imageIdNpu)
	//}

}

func branchNameContainsNotebookInvalidChar(branchName string) bool {

	for _, invalidChar := range notebookDirInvalidChar {
		if strings.Contains(branchName, invalidChar) {
			return true
		}
	}
	return false

}

func isImageMatch(npuImage string, book *models.Cloudbrain) bool {
	return (!book.IsNPUTask()) || book.Image == npuImage

}
func FileNotebookStatus(ctx *context.Context, option api.CreateFileNotebookJobOption) {
	if ctx.Written() {
		return
	}

	if path.Ext(option.File) != NoteBookExtension {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.notebook_select_wrong")))
		return
	}

	isNotebookFileExist, _ := isNoteBookFileExist(ctx, option)
	if !isNotebookFileExist {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("repo.notebook_file_not_exist")))
		return
	}
	var task *models.Cloudbrain
	var err error
	if option.ID > 0 {
		task, err = models.GetCloudbrainByCloudbrainID(option.ID)
	} else {
		task, err = models.GetCloudbrainByJobID(option.JobId)
	}
	if err != nil {
		log.Error("job not found:"+option.JobId, err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi("Job id may not be right. can not find job."))
		return
	}
	if task.BootFile == "" || task.Status != string(models.ModelArtsRunning) {
		log.Warn("Boot file is empty or status is running. ")
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi("Boot file is empty or status is running."))
		return
	}
	if !isRepoFileMatch(option, task) {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi("can not math repo file."))
		return
	}
	debugBaseUrl, token, err := getBaseUrlAndToken(task)
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
		return
	}

	if uploadNotebookFileIfCannotBroswer(debugBaseUrl, GetBootFile(option.File, option.OwnerName, option.ProjectName, option.BranchName), task, token) {
		ctx.JSON(http.StatusOK, models.BaseOKMessageApi)
	} else {
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi("upload failed."))

	}

}

func getBaseUrlAndToken(task *models.Cloudbrain) (string, string, error) {
	var debugBaseUrl string
	var token string
	if task.Type == models.TypeC2Net {
		log.Info("get grampus debug url begin: ", task.JobID)
		result, err := grampus.GetNotebookJob(task.JobID)
		if err != nil {
			return "", "", err
		}
		if result == nil {
			return "", "", fmt.Errorf("can not get job response.")
		}
		convertRes := entity.ConvertGrampusNotebookResponse(result.JobInfo)
		debugBaseUrl = convertRes.Url
		token = convertRes.Token
		if debugBaseUrl == "" {
			log.Error("notebook job not found:"+task.JobID, err)
			return "", "", fmt.Errorf("can not get job or job is invalid.")
		}

	} else {
		var result *models.GetNotebook2Result
		var err error
		if task.Type == models.TypeCloudBrainTwo {
			result, err = modelarts.GetNotebook2(task.JobID)
		} else if task.Type == models.TypeCDCenter {
			result, err = modelarts_cd.GetNotebook(task.JobID)
		}
		if err != nil || result == nil || result.Status != string(models.ModelArtsRunning) || result.Url == "" {
			log.Error("notebook job not found:"+task.JobID, err)
			return "", "", fmt.Errorf("can not get job or job is invalid.")
		}

		debugBaseUrl = result.Url
		token = result.Token

	}
	return debugBaseUrl, token, nil
}

func uploadNotebookFileIfCannotBroswer(debugBaseUrl string, bootFile string, task *models.Cloudbrain, token string) bool {
	c := &notebook.NotebookContent{
		Url:      debugBaseUrl,
		Path:     bootFile,
		PathType: "file",
		Token:    token,
	}
	if c.IsNotebookFileCanBrowser() {
		return true
	} else {
		c.SetCookiesAndCsrf()
		c.UploadNoteBookFile(task)
		return c.IsNotebookFileCanBrowser()
	}

}

func isNotebookSpecMath(option api.CreateFileNotebookJobOption, book *models.Cloudbrain) bool {
	if option.Type == NPUType || option.Type == CPUType {
		return true
	}
	spec, err := models.GetCloudbrainSpecByID(book.ID)
	if err != nil {
		log.Warn("can not get spec ", err)
		return false
	}
	return spec.AccCardsNum > 0
}

func isRepoBranchMatch(option api.CreateFileNotebookJobOption, book *models.Cloudbrain) bool {
	bootFiles := strings.Split(book.BootFile, ";")

	for _, bootFile := range bootFiles {
		splits := strings.Split(bootFile, "/")
		if len(splits) >= 4 {
			if splits[0] == option.OwnerName && splits[1] == option.ProjectName && splits[2] == option.BranchName {
				return true
			}
		}
	}
	return false

}

func isRepoFileMatch(option api.CreateFileNotebookJobOption, book *models.Cloudbrain) bool {
	bootFiles := strings.Split(book.BootFile, ";")
	branches := strings.Split(book.BranchName, ";")

	for i, bootFile := range bootFiles {
		if branches[i] == option.BranchName && GetBootFile(option.File, option.OwnerName, option.ProjectName, option.BranchName) == bootFile {
			return true
		}
	}

	return false

}
func UploadNotebookFiles(task *models.Cloudbrain) {
	if task.Status == string(models.JobRunning) && task.BootFile != "" {

		debugBaseUrl, token, err := getBaseUrlAndToken(task)
		if err != nil {
			log.Error("can not get base url:", err)
			return
		}
		bootFiles := strings.Split(task.BootFile, ";")

		for _, bootFile := range bootFiles {
			uploadNotebookFileIfCannotBroswer(debugBaseUrl, bootFile, task, token)
		}

	}
}

func getComputeResource(optionType int) string {
	if optionType == CPUType {
		return ""
	}
	if optionType == GPUType {
		return models.GPUResource
	} else {
		return models.NPUResource
	}

}

func getCodePath(jobName string, repo *models.Repository, branchName string) string {
	return setting.JobPath + jobName + cloudbrain.CodeMountPath + "/" + repo.OwnerName + "/" + repo.Name + "/" + branchName
}

func getDescription(option api.CreateFileNotebookJobOption) string {
	des := option.OwnerName + "/" + option.ProjectName + "/" + option.BranchName + "/" + option.File
	if len(des) <= CharacterLength {
		return des
	}
	return ""
}

func isNoteBookFileExist(ctx *context.Context, option api.CreateFileNotebookJobOption) (bool, error) {
	repoPathOfNoteBook := models.RepoPath(option.OwnerName, option.ProjectName)

	gitRepoOfNoteBook, err := git.OpenRepository(repoPathOfNoteBook)
	if err != nil {
		log.Error("RepoRef Invalid repo "+repoPathOfNoteBook, err.Error())
		return false, err
	}
	// We opened it, we should close it
	defer func() {
		// If it's been set to nil then assume someone else has closed it.
		if gitRepoOfNoteBook != nil {
			gitRepoOfNoteBook.Close()
		}
	}()
	fileExist, err := fileExists(gitRepoOfNoteBook, option.File, option.BranchName)
	if err != nil || !fileExist {
		log.Error("Get file error:", err, ctx.Data["MsgID"])

		return false, err
	}
	return true, nil
}

func GetBootFile(filePath string, ownerName string, projectName string, branchName string) string {
	return ownerName + "/" + projectName + "/" + branchName + "/" + filePath
}

func fileExists(gitRepo *git.Repository, path string, branch string) (bool, error) {

	commit, err := gitRepo.GetBranchCommit(branch)
	if err != nil {
		return false, err
	}
	if _, err := commit.GetTreeEntryByPath(path); err != nil {
		return false, err
	}
	return true, nil
}

func GrampusNotebookUrl(result *models.GrampusNotebookResponse) string {
	if len(result.JobInfo.Tasks) > 0 {
		return (result.JobInfo.Tasks[0].Url + "?token=" + result.JobInfo.Tasks[0].Token)
	}
	return ""
}

func GrampusNotebookDebug(ctx *context.Context) error {

	result, err := grampus.GetNotebookJob(ctx.Cloudbrain.JobID)

	if err != nil {
		return err
	}
	if len(result.JobInfo.Tasks) > 0 {
		ctx.Redirect(GrampusNotebookUrl(result))
		return nil
	}
	return fmt.Errorf("Can not find the job.")
}

func GrampusStopJob(ctx *context.Context) {
	var ID = ctx.Params(":id")
	var resultCode = "0"
	var errorMsg = ""
	var status = ""

	task := ctx.Cloudbrain
	for {
		if task.Status == models.GrampusStatusStopped || task.Status == models.GrampusStatusFailed || task.Status == models.GrampusStatusSucceeded {
			log.Error("the job(%s) has been stopped", task.JobName, ctx.Data["msgID"])
			resultCode = "-1"
			errorMsg = ctx.Tr("cloudbrain.Already_stopped")
			break
		}

		res, err := grampus.StopJob(task.JobID, task.JobType)
		if err != nil {
			log.Error("StopJob(%s) failed:%v", task.JobName, err, ctx.Data["msgID"])
			resultCode = strconv.Itoa(res.ErrorCode)
			errorMsg = ctx.Tr("cloudbrain.Stopped_failed")
			break
		}
		oldStatus := task.Status
		task.Status = GetStopJobResponseStatus(res)
		if task.EndTime == 0 {
			task.EndTime = timeutil.TimeStampNow()
		}
		task.ComputeAndSetDuration()
		if oldStatus != task.Status {
			notification.NotifyChangeCloudbrainStatus(task, oldStatus)
		}
		err = models.UpdateJob(task)
		if err != nil {
			log.Error("UpdateJob(%s) failed:%v", task.JobName, err, ctx.Data["msgID"])
			resultCode = "-1"
			errorMsg = "system error"
			break
		}

		status = task.Status
		break
	}

	ctx.JSON(200, map[string]interface{}{
		"result_code": resultCode,
		"error_msg":   errorMsg,
		"status":      status,
		"id":          ID,
		"StatusOK":    0,
	})
}
func GrampusNotebookRestart(ctx *context.Context) {
	var id = ctx.Params(":id")
	var resultCode = "-1"
	var errorMsg = ""
	var status = ""
	var spec *models.Specification

	task := ctx.Cloudbrain
	if ctx.Written() {
		return
	}

	lockOperator, errMsg := cloudbrainService.Lock4CloudbrainRestart(&lock.LockContext{Repo: ctx.Repo.Repository, Task: &models.Cloudbrain{JobType: task.JobType}, User: ctx.User})
	defer func() {
		if lockOperator != nil {
			lockOperator.Unlock()
		}
	}()

	if errMsg != "" {
		log.Error("lock processed failed:%s", errMsg, ctx.Data["MsgID"])
		errorMsg = ctx.Tr(errMsg)
	}

	for {
		if errorMsg != "" {
			break
		}

		if task.Status != models.GrampusStatusStopped && task.Status != models.GrampusStatusSucceeded && task.Status != models.GrampusStatusFailed {
			log.Error("the job(%s) is not stopped", task.JobName, ctx.Data["MsgID"])
			errorMsg = "the job is not stopped"
			break
		}

		count, err := GetNotFinalStatusTaskCount(ctx.User.ID, string(models.JobTypeDebug))

		if err != nil {
			log.Error("GetCloudbrainNotebookCountByUserID failed:%v", err, ctx.Data["MsgID"])
			errorMsg = "system error"
			break
		} else {
			if count >= 1 {
				log.Error("the user already has running or waiting task", ctx.Data["MsgID"])
				resultCode = "2"
				errorMsg = ctx.Tr("repo.cloudbrain.morethanonejob")
				break
			}
		}

		oldSpec, err := resource.GetCloudbrainSpec(task.ID)
		if err != nil || oldSpec == nil {
			log.Error("NotebookManage GetCloudbrainSpec error.%v", err)
			errorMsg = "Resource specification not available"
			break
		}

		computeSourceSimple := models.GPU
		action := models.ActionCreateGrampusGPUDebugTask
		if task.ComputeResource == models.NPUResource {
			computeSourceSimple = models.NPU
			action = models.ActionCreateGrampusNPUDebugTask
		} else if task.ComputeResource == models.GCUResource {
			computeSourceSimple = models.GCU
			action = models.ActionCreateGrampusGCUDebugTask
		} else if task.ComputeResource == models.DCU {
			computeSourceSimple = models.DCU
			action = models.ActionCreateGrampusDCUDebugTask
		}

		spec, err = resource.GetAndCheckSpec(ctx.User.ID, oldSpec.ID, models.FindSpecsOptions{
			JobType:         models.JobType(task.JobType),
			ComputeResource: computeSourceSimple,
			Cluster:         models.C2NetCluster,
		})
		if err != nil || spec == nil {
			log.Error("NotebookManage GetAndCheckSpec error.task.id = %d", task.ID)
			errorMsg = "Resource specification not support any more"
			break
		}
		if !account.IsPointBalanceEnough(ctx.User.ID, models.PointDeductCondition{SpecUnitPrice: spec.UnitPrice}) {
			log.Error("point balance is not enough,userId=%d specId=%d", ctx.User.ID, spec.ID)
			errorMsg = ctx.Tr("points.insufficient_points_balance")
			break
		}
		if task.IsGPUTask() || task.IsGCUTask() {
			if _, err := os.Stat(getOldJobPath(task)); err != nil {
				log.Error("Can not find job minio path", err)
				resultCode = "-1"
				errorMsg = ctx.Tr("cloudbrain.result_cleared")
				break
			}
		}

		if !modelmanage.HasModelFile(task) { //使用预训练模型训练
			errorMsg = ctx.Tr("repo.debug.manage.model_not_exist")
			break
		}
		if hasDatasetDeleted(task) {
			errorMsg = ctx.Tr("repo.debug.manage.dataset_not_exist")
			break
		}

		createTime := timeutil.TimeStampNow()

		res, err := grampus.RestartNotebookJob(task.JobID)
		if err != nil {
			log.Error("ManageNotebook2(%s) failed:%v", task.DisplayJobName, err.Error(), ctx.Data["MsgID"])
			errorMsg = ctx.Tr("repo.debug_again_fail")
			break
		}

		if res.GrampusResult.ErrorCode != 0 || res.NewId == "" {
			log.Error("ManageNotebook2 failed:" + res.GrampusResult.ErrorMsg)
			errorMsg = ctx.Tr("repo.debug_again_fail")
			if res.GrampusResult.ErrorCode == 5005 {
				errorMsg = ctx.Tr("repo.debug_again_fail_forever")
			}

			break
		}

		if res.Status == models.GrampusStatusPending {
			res.Status = models.GrampusStatusWaiting
		}

		newTask := &models.Cloudbrain{
			Status:           res.Status,
			UserID:           task.UserID,
			RepoID:           task.RepoID,
			JobID:            res.NewId,
			JobName:          task.JobName,
			DisplayJobName:   task.DisplayJobName,
			JobType:          task.JobType,
			Type:             task.Type,
			Uuid:             task.Uuid,
			Image:            task.Image,
			ImageID:          task.ImageID,
			EngineID:         task.EngineID,
			CommitID:         task.CommitID,
			EngineName:       task.EngineName,
			IsLatestVersion:  "1",
			BranchName:       task.BranchName,
			DatasetName:      task.DatasetName,
			ComputeResource:  task.ComputeResource,
			Description:      task.Description,
			CreatedUnix:      createTime,
			UpdatedUnix:      createTime,
			Spec:             spec,
			ModelName:        task.ModelName,
			ModelVersion:     task.ModelVersion,
			LabelName:        task.LabelName,
			PreTrainModelUrl: task.PreTrainModelUrl,
			CkptName:         task.CkptName,
			ModelId:          task.ModelId,
			WorkServerNumber: 1,
		}

		err = models.RestartCloudbrain(task, newTask)
		if err != nil {
			log.Error("RestartCloudbrain(%s) failed:%v", task.JobName, err.Error(), ctx.Data["MsgID"])
			errorMsg = "system error"
			break
		}

		id = strconv.FormatInt(newTask.ID, 10)

		status = res.Status
		resultCode = "0"

		notification.NotifyOtherTask(ctx.User, nil, id, newTask.DisplayJobName, action)

		break
	}

	ctx.JSON(200, map[string]string{
		"result_code": resultCode,
		"error_msg":   errorMsg,
		"status":      status,
		"id":          id,
	})
}

func getOldJobPath(task *models.Cloudbrain) string {
	return setting.Attachment.Minio.RealPath + setting.Attachment.Minio.Bucket + "/" + setting.CBCodePathPrefix + task.JobName
}

func hasDatasetDeleted(task *models.Cloudbrain) bool {
	if task.Uuid == "" {
		return false
	}
	uuids := strings.Split(task.Uuid, ";")
	attachs, _ := models.GetAttachmentsByUUIDs(uuids)
	return len(attachs) < len(uuids)
}
func GetStopJobResponseStatus(res *models.GrampusStopJobResponse) string {
	newStatus := models.GrampusStatusStopping
	if res != nil && res.Status != "" {
		newStatus = grampus.TransTrainJobStatus(res.Status)
	}
	return newStatus
}

func DeleteGrampusJob(ctx *context.Context) error {
	task := ctx.Cloudbrain
	if task.Status != models.GrampusStatusStopped && task.Status != models.GrampusStatusSucceeded && task.Status != models.GrampusStatusFailed {
		log.Error("the job(%s) has not been stopped", task.JobName, ctx.Data["msgID"])
		return errors.New(ctx.Tr("cloudbrain.Not_Stopped"))
	}

	err := models.DeleteJob(task)
	if err != nil {
		log.Error("DeleteJob failed: %v", err, ctx.Data["msgID"])
		return errors.New(ctx.Tr("cloudbrain.Delete_failed"))
	}

	storageType := models.TypeCloudBrainOne
	if task.ComputeResource == models.NPUResource {
		storageType = models.TypeCloudBrainTwo
	}
	DeleteCloudbrainJobStorage(task.JobName, storageType)

	return nil
}

func getNpuImageId(option api.CreateFileNotebookJobOption) (string, error) {
	if option.Type != NPUType {
		return "", fmt.Errorf("type is not npu.")
	}
	if option.Image == "" {
		return "", nil
	}
	for _, imageInfo := range setting.StImageInfos.ImageInfo {
		if imageInfo.Value == option.Image {
			return imageInfo.Id, nil
		}
	}
	return "", fmt.Errorf("invalid image parameter")
}
