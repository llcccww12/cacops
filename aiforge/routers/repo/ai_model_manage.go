package repo

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"

	"code.gitea.io/gitea/entity"

	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_model"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
	"code.gitea.io/gitea/services/cloudbrain/modelmanage"
	"code.gitea.io/gitea/services/cloudbrain/resource"
	"code.gitea.io/gitea/services/storage_limit"

	"code.gitea.io/gitea/services/repository"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/notification"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
	uuid "github.com/satori/go.uuid"
)

const (
	Attachment_model                 = "model"
	Model_prefix                     = "aimodels/"
	tplModelManageIndex              = "repo/modelmanage/index"
	tplModelManageDownload           = "repo/modelmanage/download"
	tplModelInfo                     = "repo/modelmanage/showinfo"
	tplCreateLocalModelInfo          = "repo/modelmanage/create_local"
	tplCreateLocalForUploadModelInfo = "repo/modelmanage/fileupload"
	tplCreateOnlineModelInfo         = "repo/modelmanage/create_online"

	MODEL_LATEST      = 1
	MODEL_NOT_LATEST  = 0
	MODEL_MAX_SIZE    = 1024 * 1024 * 1024
	STATUS_COPY_MODEL = 1
	STATUS_FINISHED   = 0
	STATUS_ERROR      = 2
)

func saveModelByParameters(aiTask *models.Cloudbrain, name string, version string, label string, description string, engine int, ctx *context.Context) (string, error) {
	uuid := uuid.NewV4()
	id := uuid.String()
	modelPath := id
	var lastNewModelId string
	var modelSize int64

	log.Info("find task name:" + aiTask.JobName)
	aimodels := models.QueryModelByName(name, ctx.Repo.Repository.ID)
	if len(aimodels) > 0 {
		for _, model := range aimodels {
			if model.Version == version {
				return "", errors.New(ctx.Tr("repo.model.manage.create_error"))
			}
			if model.New == MODEL_LATEST {
				lastNewModelId = model.ID
			}
		}
	}

	modelSelectedFile := ctx.Query("modelSelectedFile")
	//download model zip   //train type

	storageType := models.GetDefaultStorageType()

	spec, err := resource.GetCloudbrainSpec(aiTask.ID)
	if err == nil {
		specJson, _ := json.Marshal(spec)
		aiTask.FlavorName = string(specJson)
	}
	accuracy := make(map[string]string)
	accuracy["F1"] = ""
	accuracy["Recall"] = ""
	accuracy["Accuracy"] = ""
	accuracy["Precision"] = ""
	accuracyJson, _ := json.Marshal(accuracy)
	log.Info("accuracyJson=" + string(accuracyJson))
	aiTask.ContainerIp = ""
	aiTaskJson, _ := json.Marshal(aiTask)
	isPrivate := ctx.QueryBool("isPrivate")
	license := ctx.Query("license")
	model := &models.AiModelManage{
		ID:              id,
		Version:         version,
		VersionCount:    len(aimodels) + 1,
		Label:           label,
		Name:            name,
		Description:     description,
		New:             MODEL_LATEST,
		Type:            storageType,
		Path:            modelPath,
		Size:            modelSize,
		AttachmentId:    aiTask.Uuid,
		RepoId:          ctx.Repo.Repository.ID,
		UserId:          ctx.User.ID,
		CodeBranch:      aiTask.BranchName,
		CodeCommitID:    aiTask.CommitID,
		Engine:          int64(engine),
		TrainTaskInfo:   string(aiTaskJson),
		Accuracy:        string(accuracyJson),
		Status:          STATUS_COPY_MODEL,
		IsPrivate:       isPrivate,
		ComputeResource: aiTask.ComputeResource,
		License:         license,
	}

	err = models.SaveModelToDb(model)
	if err != nil {
		return "", err
	}
	if len(lastNewModelId) > 0 {
		//udpate status and version count
		models.ModifyModelNewProperty(lastNewModelId, MODEL_NOT_LATEST, 0)
	}
	var units []models.RepoUnit
	var deleteUnitTypes []models.UnitType
	units = append(units, models.RepoUnit{
		RepoID: ctx.Repo.Repository.ID,
		Type:   models.UnitTypeModelManage,
		Config: &models.ModelManageConfig{
			EnableModelManage: true,
		},
	})
	deleteUnitTypes = append(deleteUnitTypes, models.UnitTypeModelManage)

	models.UpdateRepositoryUnits(ctx.Repo.Repository, units, deleteUnitTypes)

	go asyncToCopyModel(aiTask, id, modelSelectedFile, ctx)

	log.Info("save model end.")
	if !model.IsPrivate {
		notification.NotifyOtherTask(ctx.User, ctx.Repo.Repository, id, name, models.ActionCreateNewModelTask)
	}
	return id, nil
}

func asyncToCopyModel(aiTask *models.Cloudbrain, id string, modelSelectedFile string, ctx *context.Context) {
	destKeyNamePrefix := Model_prefix + models.AttachmentRelativePath(id) + "/"
	aiConfig := aiTask.GetCloudbrainConfig()
	if aiConfig == nil {
		log.Error("GetCloudbrainConfig empty, cloudbrain=%s", aiTask.JobName)
		return
	}
	if aiConfig.OutputStorageType == string(entity.OBS) {
		//destKeyNamePrefix := Model_prefix + models.AttachmentRelativePath(modelUUID) + "/"
		modelPath, modelSize, err := downloadModelFromObsStorage(aiConfig.OutputObjectPrefix, "", aiConfig.OutputBucket, modelSelectedFile, destKeyNamePrefix, ctx)
		if err != nil {
			UpdateStatus(id, 0, STATUS_ERROR, modelPath, err.Error())
			log.Info("download model from CloudBrainTwo faild." + err.Error())
		} else {
			UpdateStatus(id, modelSize, STATUS_FINISHED, modelPath, "")
			insertModelFile(id)
		}
	} else {
		modelPath, modelSize, err := downloadModelFromMinioStorage(aiConfig.OutputObjectPrefix, aiConfig.OutputBucket, modelSelectedFile, destKeyNamePrefix, ctx)

		if err != nil {
			UpdateStatus(id, 0, STATUS_ERROR, modelPath, err.Error())
			log.Info("download model from CloudBrainOne faild." + err.Error())
		} else {
			UpdateStatus(id, modelSize, STATUS_FINISHED, modelPath, "")
			insertModelFile(id)
		}
	}
	if aiTask.ModelId != "" {
		modelids := strings.Split(aiTask.ModelId, ";")
		for _, v := range modelids {
			log.Info("model:" + v + "  DerivativeCount +1")
			models.ModifyModelDerivativeCount(v)
		}
	}
	ai_model.UpdateModelMeta(id)
}

func insertModelFile(id string) {
	model, _ := models.QueryModelById(id)
	helper := storage_helper.SelectStorageHelperFromStorageIntType(model.Type)
	bucket := helper.GetBucket()
	if len(model.Path) > len(bucket)+1 {
		files, err := helper.GetAllObjectByBucketAndPrefix(bucket, model.Path[len(bucket)+1:])
		if err != nil {
			log.Info("Failed to query model size from obs. id=" + id)
		}
		for _, file := range files {
			modelFile := &models.AiModelFile{
				ModelID: id,
				Name:    file.FileName,
				Size:    file.Size,
			}
			models.SaveModelFile(modelFile)
		}
	}
}

func UpdateStatus(id string, modelSize int64, status int, modelPath string, statusDesc string) {
	if len(statusDesc) > 400 {
		statusDesc = statusDesc[0:400]
	}
	m, _ := models.QueryModelById(id)
	err := models.ModifyModelStatus(id, modelSize, status, modelPath, statusDesc)
	if err != nil {
		log.Info("update status error." + err.Error())
	}
	if m != nil {
		if modelSize > 0 && m.Size == 0 {
			go repository.ResetRepoModelNum(m.RepoId)
		}
	}

}

func SaveNewNameModel(ctx *context.Context) {
	if !ctx.Repo.CanWrite(models.UnitTypeModelManage) {
		ctx.Error(403, ctx.Tr("repo.model_noright"))
		return
	}
	name := ctx.Query("name")
	if name == "" {
		ctx.Error(500, fmt.Sprintf("name or version is null."))
		return
	}

	aimodels := models.QueryModelByName(name, ctx.Repo.Repository.ID)
	if len(aimodels) > 0 {
		ctx.Error(500, ctx.Tr("repo.model_rename"))
		return
	}
	SaveModel(ctx)
	//ctx.Status(200)
	log.Info("save model end.")
}

func SaveLocalModel(ctx *context.Context) {
	if !ctx.Repo.CanWrite(models.UnitTypeModelManage) {
		ctx.Error(403, ctx.Tr("repo.model_noright"))
		return
	}
	re := map[string]string{
		"code": "-1",
	}
	log.Info("save SaveLocalModel start.")
	uuid := uuid.NewV4()
	id := uuid.String()
	name := ctx.Query("name")
	version := ctx.Query("version")
	if version == "" {
		version = "0.0.1"
	}
	label := ctx.Query("label")
	description := ctx.Query("description")
	engine := ctx.QueryInt("engine")
	storageType := models.GetDefaultStorageType()
	isPrivate := ctx.QueryBool("isPrivate")
	if ctx.Repo.Repository.IsPrivate {
		if !isPrivate {
			re["msg"] = "Private repo cannot create public model."
			ctx.JSON(200, re)
			return
		}
	}
	modelActualPath := ""

	if storageType == models.StorageTypeMinio {
		destKeyNamePrefix := Model_prefix + models.AttachmentRelativePath(id) + "/"
		modelActualPath = setting.Attachment.Minio.Bucket + "/" + destKeyNamePrefix

	} else if storageType == models.StorageTypeObs {
		destKeyNamePrefix := Model_prefix + models.AttachmentRelativePath(id) + "/"
		modelActualPath = setting.Bucket + "/" + destKeyNamePrefix
	}
	var lastNewModelId string
	repoId := ctx.Repo.Repository.ID
	aimodels := models.QueryModelByName(name, repoId)
	if len(aimodels) > 0 {
		for _, model := range aimodels {
			if model.Version == version {
				re["msg"] = ctx.Tr("repo.model.manage.create_error")
				ctx.JSON(200, re)
				return
			}
			if model.New == MODEL_LATEST {
				lastNewModelId = model.ID
			}
		}
	}
	license := ctx.Query("license")
	model := &models.AiModelManage{
		ID:            id,
		Version:       version,
		ModelType:     models.MODEL_LOCAL_TYPE,
		VersionCount:  len(aimodels) + 1,
		Label:         label,
		Name:          name,
		Description:   description,
		New:           MODEL_LATEST,
		Type:          storageType,
		Path:          modelActualPath,
		Size:          0,
		AttachmentId:  "",
		RepoId:        repoId,
		UserId:        ctx.User.ID,
		Engine:        int64(engine),
		TrainTaskInfo: "",
		Accuracy:      "",
		Status:        STATUS_FINISHED,
		IsPrivate:     isPrivate,
		License:       license,
	}

	err := models.SaveModelToDb(model)
	if err != nil {
		re["msg"] = err.Error()
		ctx.JSON(200, re)
		return
	}
	if len(lastNewModelId) > 0 {
		//udpate status and version count
		models.ModifyModelNewProperty(lastNewModelId, MODEL_NOT_LATEST, 0)
	}
	var units []models.RepoUnit
	var deleteUnitTypes []models.UnitType
	units = append(units, models.RepoUnit{
		RepoID: ctx.Repo.Repository.ID,
		Type:   models.UnitTypeModelManage,
		Config: &models.ModelManageConfig{
			EnableModelManage: true,
		},
	})
	deleteUnitTypes = append(deleteUnitTypes, models.UnitTypeModelManage)

	models.UpdateRepositoryUnits(ctx.Repo.Repository, units, deleteUnitTypes)

	log.Info("save model end.")
	if !model.IsPrivate {
		notification.NotifyOtherTask(ctx.User, ctx.Repo.Repository, id, name, models.ActionCreateNewModelTask)
	}
	re["code"] = "0"
	re["id"] = id
	ctx.JSON(200, re)
}

func getSize(files []storage.FileInfo) int64 {
	var size int64
	for _, file := range files {
		size += file.Size
	}
	return size
}

func UpdateModelSize(modeluuid string, objectName string) {
	model, err := models.QueryModelById(modeluuid)
	if err == nil {
		var size int64
		helper := storage_helper.SelectStorageHelperFromStorageIntType(model.Type)
		bucket := helper.GetBucket()
		if strings.HasPrefix(model.Path, bucket+"/"+Model_prefix) {
			files, err := helper.GetAllObjectByBucketAndPrefix(bucket, model.Path[len(bucket)+1:])
			if err != nil {
				log.Info("Failed to query model size from obs. id=" + modeluuid)
			}
			size = getSize(files)
			models.ModifyModelSize(modeluuid, size)
			modelFileName := objectName
			index := strings.LastIndex(objectName, "/")
			if index > 0 {
				modelFileName = objectName[index+1:]
			}
			log.Info("modelFileName=" + modelFileName)
			for _, file := range files {
				log.Info("fileName=" + file.FileName)
				if file.FileName == modelFileName {
					modelFile := &models.AiModelFile{
						ModelID: modeluuid,
						Name:    file.FileName,
						Size:    file.Size,
					}
					models.SaveModelFile(modelFile)
				}
			}
		}
		if model.Size == 0 && size > 0 {
			go repository.ResetRepoModelNum(model.RepoId)
		}
	} else {
		log.Info("not found model,uuid=" + modeluuid)
	}
}

func SaveModel(ctx *context.Context) {
	if !ctx.Repo.CanWrite(models.UnitTypeModelManage) {
		ctx.Error(403, ctx.Tr("repo.model_noright"))
		return
	}
	log.Info("save model start.")
	JobId := ctx.Query("jobId")
	VersionName := ctx.Query("versionName")
	name := ctx.Query("name")
	version := ctx.Query("version")
	label := ctx.Query("label")
	description := ctx.Query("description")
	engine := ctx.QueryInt("engine")
	modelSelectedFile := ctx.Query("modelSelectedFile")
	log.Info("engine=" + fmt.Sprint(engine) + " modelSelectedFile=" + modelSelectedFile)
	re := map[string]string{
		"code": "-1",
	}

	var aiTask *models.Cloudbrain
	var err error
	//云脑重构:适配用id的方式请求
	cloudbrainId := ctx.QueryInt64("cloudbrain_id")
	if cloudbrainId > 0 {
		aiTask, err = models.GetCloudbrainByCloudbrainID(cloudbrainId)
	} else {
		aiTask, err = models.GetCloudbrainByJobIDAndVersionName(JobId, VersionName)
		if err != nil {
			aiTask, err = models.GetRepoCloudBrainByJobID(ctx.Repo.Repository.ID, JobId)
		}
	}
	if err != nil {
		log.Error("save model error." + err.Error())
		re["msg"] = err.Error()
		return
	}

	isPrivate := ctx.QueryBool("isPrivate")
	if ctx.Repo.Repository.IsPrivate {
		if !isPrivate {
			re["msg"] = "Private repo cannot create public model."
			ctx.JSON(200, re)
			return
		}
	}
	if modelSelectedFile == "" {
		re["msg"] = "Not selected model file."
		ctx.JSON(200, re)
		return
	}
	if name == "" || version == "" {
		re["msg"] = "name or version is null."
		ctx.JSON(200, re)
		return
	}
	id, err := saveModelByParameters(aiTask, name, version, label, description, engine, ctx)
	if err != nil {
		log.Info("save model error." + err.Error())
		re["msg"] = err.Error()
	} else {
		re["code"] = "0"
		re["id"] = id
	}
	ctx.JSON(200, re)
	log.Info("save model end.")
}

func downloadModelFromObsStorage(objectKeyPrefix string, parentDir string, bucket string, modelSelectedFile string, destKeyNamePrefix string, ctx *context.Context) (string, int64, error) {
	objectKey := strings.TrimPrefix(path.Join(objectKeyPrefix, parentDir), "/")
	//if trainUrl != "" {
	//	objectkey = strings.Trim(trainUrl[len(setting.Bucket)+1:], "/")
	//}
	prefix := objectKey + "/"
	filterFiles := strings.Split(modelSelectedFile, ";")
	Files := make([]string, 0)
	for _, shortFile := range filterFiles {
		Files = append(Files, prefix+shortFile)
	}
	totalSize := storage.ObsGetFilesSize(setting.Bucket, Files)
	if storage_limit.IsFileUploadOverLimit(ctx.User.ID, totalSize) {
		return "", 0, errors.New(ctx.Tr("common_error.total_file_size_storage_limit"))
	}
	if float64(totalSize) > setting.MaxModelSizeForWhole*MODEL_MAX_SIZE {
		return "", 0, errors.New("Cannot create model, as model is exceed " + fmt.Sprint(setting.MaxModelSize) + "G.")
	}

	modelDbResult, err := storage.GetOneLevelAllObjectUnderDir(bucket, objectKey, "")
	log.Info("bucket=" + setting.Bucket + "  objectkey=" + objectKey)
	if err != nil {
		log.Info("get TrainJobListModel failed:", err)
		return "", 0, err
	}
	if len(modelDbResult) == 0 {
		return "", 0, errors.New("Cannot create model, as model is empty.")
	}

	//destKeyNamePrefix := Model_prefix + models.AttachmentRelativePath(modelUUID) + "/"
	size, err := storage.ObsCopyManyFile(setting.Bucket, prefix, setting.Bucket, destKeyNamePrefix, filterFiles)

	dataActualPath := setting.Bucket + "/" + destKeyNamePrefix
	return dataActualPath, size, err
}

func downloadModelFromMinioStorage(objectKeyPrefix string, bucket string, modelSelectedFile string, destKeyNamePrefix string, ctx *context.Context) (string, int64, error) {
	modelSrcPrefix := strings.TrimSuffix(objectKeyPrefix, "/") + "/"
	log.Info("destKeyNamePrefix=" + destKeyNamePrefix + "   modelSrcPrefix=" + modelSrcPrefix + "  bucket=" + bucket)
	filterFiles := strings.Split(modelSelectedFile, ";")
	Files := make([]string, 0)
	for _, shortFile := range filterFiles {
		Files = append(Files, modelSrcPrefix+shortFile)
	}
	totalSize := storage.MinioGetFilesSize(bucket, Files)
	if storage_limit.IsFileUploadOverLimit(ctx.User.ID, totalSize) {
		return "", 0, errors.New(ctx.Tr("common_error.total_file_size_storage_limit"))
	}
	if float64(totalSize) > setting.MaxModelSizeForWhole*MODEL_MAX_SIZE {
		return "", 0, errors.New("Cannot create model, as model is exceed " + fmt.Sprint(setting.MaxModelSize) + "G.")
	}

	size, err := storage.MinioCopyFiles(bucket, modelSrcPrefix, destKeyNamePrefix, filterFiles)

	dataActualPath := bucket + "/" + destKeyNamePrefix
	return dataActualPath, size, err

}

func DeleteModelFileInternal(id string, fileName string) error {
	model, err := models.QueryModelById(id)
	if err == nil {
		var totalSize int64
		if model.ModelType == models.MODEL_LOCAL_TYPE || model.ModelType == models.MODEL_HF_TYPE {
			helper := storage_helper.SelectStorageHelperFromStorageIntType(model.Type)
			bucketName := helper.GetBucket()
			if len(model.Path) > (len(bucketName) + 1) {
				objectName := model.Path[len(bucketName)+1:] + fileName
				log.Info("delete bucket=" + bucketName + " path=" + objectName)
				if strings.HasPrefix(model.Path, bucketName+"/"+Model_prefix) {
					totalSize = helper.GetFilesSize([]string{objectName})
					err = helper.RemoveObject(objectName)
					if err != nil {
						log.Info("Failed to delete model. id=" + id)
						return err
					} else {
						log.Info("delete obs file size is:" + fmt.Sprint(totalSize))
						newSize := model.Size - totalSize
						if newSize < 0 {
							newSize = 0
						}
						models.ModifyModelSize(id, newSize)
						modelFile := &models.AiModelFile{
							Name:    fileName,
							ModelID: id,
						}
						models.DeleteModelFile(modelFile)
						ai_model.UpdateModelMeta(id)
					}
				}
			}
		}
		if (model.Size - totalSize) <= 0 {
			go repository.ResetRepoModelNum(model.RepoId)
		}
	}
	return nil
}

func DeleteModelFile(ctx *context.Context) {
	log.Info("delete model start.")
	id := ctx.Query("id")
	fileName := ctx.Query("fileName")
	err := DeleteModelFileInternal(id, fileName)
	if err != nil {
		ctx.JSON(200, map[string]string{
			"code": "-1",
			"msg":  err.Error(),
		})
	}

	ctx.JSON(200, map[string]string{
		"code": "0",
	})
}

func DeleteModel(ctx *context.Context) {
	log.Info("delete model start.")
	id := ctx.Query("id")
	err := deleteModelByID(ctx, id)
	if err != nil {
		re := map[string]string{
			"code": "-1",
		}
		re["msg"] = err.Error()
		ctx.JSON(200, re)
	} else {
		ctx.JSON(200, map[string]string{
			"code": "0",
		})
	}
}

func BatchDeleteModel(ctx *context.Context) {
	log.Info("delete model start.")
	ids := ctx.Query("id")
	idList := strings.Split(ids, ";")
	succeedIdList := make([]string, 0)
	for _, id := range idList {
		err := deleteModelByID(ctx, id)
		if err != nil {
			log.Warn("delete model (%s) failed:%s", id, err.Error())
			continue
		} else {
			succeedIdList = append(succeedIdList, id)
		}
	}
	ctx.JSON(200, map[string][]string{
		"id": succeedIdList,
	})
}

func UpdateAllModelMeta(ctx *context.Context) {
	log.Info("Start to update all model meta")
	ids := ctx.QueryStrings("model_id")
	updateAll := ctx.QueryBool("update_all")
	if !updateAll && len(ids) == 0 {
		ctx.JSON(http.StatusOK, response.OuterSuccessWithData(map[string]int{"count": 0}))
		return
	}
	count := 0
	if updateAll {
		pageSize := 100
		pageNum := 1
		for {
			ids, err := models.QueryModelIdsByPaging(pageSize, pageNum, "created_unix")
			if err != nil {
				log.Error("UpdateAllModelMeta QueryModelIdsByPaging err.%v", err)
				ctx.JSON(http.StatusOK, response.OuterResponseError(err))
				return
			}
			for _, id := range ids {
				ai_model.UpdateModelMeta(id)
				count++
			}
			if len(ids) < pageSize {
				break
			}
			pageNum++
		}

	} else {
		for _, id := range ids {
			ai_model.UpdateModelMeta(id)
			count++
		}
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(map[string]int{"count": count}))
}

func QueryModelMetaById(ctx *context.Context) {
	log.Info("Start to query model meta")
	id := ctx.Query("model_id")
	data, err := ai_model.QueryModelMeta(id)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterResponseError(err))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(map[string]interface{}{"meta_data": data}))
}

func deleteModelByID(ctx *context.Context, id string) error {
	log.Info("delete model start. id=" + id)
	model, err := models.QueryModelById(id)
	if !isCanDelete(ctx, model.UserId) {
		return errors.New(ctx.Tr("repo.model_noright"))
	}
	if err == nil {
		helper := storage_helper.SelectStorageHelperFromStorageIntType(model.Type)
		bucketName := helper.GetBucket()

		log.Info("bucket=" + bucketName + " path=" + model.Path)
		if strings.HasPrefix(model.Path, bucketName+"/"+Model_prefix) {

			err = helper.RemoveObject(model.Path[len(bucketName)+1:])
			if err != nil {
				log.Info("Failed to delete model. id=" + id)
				return err
			}
		}

		err = models.DeleteModelById(id)
		if err == nil { //find a model to change new
			aimodels := models.QueryModelByName(model.Name, model.RepoId)
			if model.New == MODEL_LATEST {
				if len(aimodels) > 0 {
					//udpate status and version count
					models.ModifyModelNewProperty(aimodels[0].ID, MODEL_LATEST, len(aimodels))
				}
			} else {
				for _, tmpModel := range aimodels {
					if tmpModel.New == MODEL_LATEST {
						models.ModifyModelNewProperty(tmpModel.ID, MODEL_LATEST, len(aimodels))
						break
					}
				}
			}
			if model.Size > 0 {
				go repository.ResetRepoModelNum(model.RepoId)
			}
		}

		// delete hf transfer file record db
		if model.ModelType == models.MODEL_HF_TYPE {
			err := models.DeleteHfFilesByModelId(model.ID)
			if err != nil {
				log.Info("delete model file error." + err.Error())
				return err
			}
		}

	}
	return err
}

func DownloadMultiModelFile(ctx *context.Context) {
	log.Info("DownloadMultiModelFile start.")
	id := ctx.Query("id")
	log.Info("id=" + id)
	task, err := models.QueryModelById(id)
	if err != nil {
		log.Error("no such model!", err.Error())
		ctx.ServerError("no such model:", err)
		return
	}
	if !isCanDownload(ctx, task) {
		ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
		return
	}

	path := Model_prefix + models.AttachmentRelativePath(id) + "/"
	if task.Type == models.TypeCloudBrainTwo {
		downloadFromCloudBrainTwo(path, task, ctx, id)
	} else if task.Type == models.TypeCloudBrainOne {
		downloadFromCloudBrainOne(path, task, ctx, id)
	}
}

func MinioDownloadManyFile(path string, ctx *context.Context, returnFileName string, allFile []storage.FileInfo) {
	ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+url.QueryEscape(returnFileName))
	ctx.Resp.Header().Set("Content-Type", "application/octet-stream")
	w := zip.NewWriter(ctx.Resp)
	defer w.Close()
	for _, oneFile := range allFile {
		if oneFile.IsDir {
			log.Info("zip dir name:" + oneFile.FileName)
		} else {
			log.Info("zip file name:" + oneFile.FileName)
			fDest, err := w.Create(oneFile.FileName)
			if err != nil {
				log.Info("create zip entry error, download file failed: %s\n", err.Error())
				ctx.ServerError("download file failed:", err)
				return
			}
			log.Info("minio file path=" + (path + oneFile.FileName))
			body, err := storage.Attachments.DownloadAFile(setting.Attachment.Minio.Bucket, path+oneFile.FileName)
			if err != nil {
				log.Info("download file failed: %s\n", err.Error())
				ctx.ServerError("download file failed:", err)
				return
			} else {
				// 读取对象内容
				bodyReader(body, fDest)
			}
		}
	}

}

func bodyReader(body io.ReadCloser, fDest io.Writer) {
	defer body.Close()
	p := make([]byte, 1024)
	var readErr error
	var readCount int

	for {
		readCount, readErr = body.Read(p)
		if readCount > 0 {
			fDest.Write(p[:readCount])
		}
		if readErr != nil {
			break
		}
	}
}

func downloadFromCloudBrainOne(path string, task *models.AiModelManage, ctx *context.Context, id string) {
	allFile, err := storage.GetAllObjectByBucketAndPrefixMinio(setting.Attachment.Minio.Bucket, path)
	if err == nil {
		//count++
		models.ModifyModelDownloadCount(id)
		returnFileName := task.Name + "_" + task.Version + ".zip"
		MinioDownloadManyFile(path, ctx, returnFileName, allFile)
	} else {
		log.Info("error,msg=" + err.Error())
		ctx.ServerError("no file to download.", err)
	}
}

func ObsDownloadManyFile(path string, ctx *context.Context, returnFileName string, allFile []storage.FileInfo) {
	ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+url.QueryEscape(returnFileName))
	ctx.Resp.Header().Set("Content-Type", "application/octet-stream")
	w := zip.NewWriter(ctx.Resp)
	defer w.Close()
	for _, oneFile := range allFile {
		if oneFile.IsDir {
			log.Info("zip dir name:" + oneFile.FileName)
		} else {
			log.Info("zip file name:" + oneFile.FileName)
			fDest, err := w.Create(oneFile.FileName)
			if err != nil {
				log.Info("create zip entry error, download file failed: %s\n", err.Error())
				ctx.ServerError("download file failed:", err)
				return
			}
			body, err := storage.ObsDownloadAFile(setting.Bucket, path+oneFile.FileName)
			if err != nil {
				log.Info("download file failed: %s\n", err.Error())
				ctx.ServerError("download file failed:", err)
				return
			} else {
				// 读取对象内容
				obsBodyReader(body, fDest)
			}
		}
	}
}

func obsBodyReader(body io.ReadCloser, fDest io.Writer) {
	defer body.Close()
	p := make([]byte, 1024)
	var readErr error
	var readCount int

	for {
		readCount, readErr = body.Read(p)
		if readCount > 0 {
			fDest.Write(p[:readCount])
		}
		if readErr != nil {
			break
		}
	}
}

func downloadFromCloudBrainTwo(path string, task *models.AiModelManage, ctx *context.Context, id string) {
	allFile, err := storage.GetAllObjectByBucketAndPrefix(setting.Bucket, path)
	if err == nil {
		//count++
		models.ModifyModelDownloadCount(id)
		returnFileName := task.Name + "_" + task.Version + ".zip"
		ObsDownloadManyFile(path, ctx, returnFileName, allFile)
	} else {
		log.Info("error,msg=" + err.Error())
		ctx.ServerError("no file to download.", err)
	}
}

func QueryTrainJobVersionList(ctx *context.Context) {

	VersionListTasks, err := QueryTrainJobVersionListApi(ctx)
	if err != nil {
		ctx.JSON(200, nil)
	} else {
		ctx.JSON(200, VersionListTasks)
	}
}

func QueryTrainJobVersionListApi(ctx *context.Context) ([]*models.Cloudbrain, error) {
	log.Info("query train job version list. start.")
	JobID := ctx.Query("jobId")
	if JobID == "" {
		JobID = ctx.Query("JobId")
	}
	VersionListTasks, count, err := models.QueryModelTrainJobVersionList(JobID)
	log.Info("query return count=" + fmt.Sprint(count))

	return VersionListTasks, err
}

func QueryTrainJobList(ctx *context.Context) {
	VersionListTasks, err := QueryTrainJobListApi(ctx)
	if err != nil {
		ctx.JSON(200, nil)
	} else {
		ctx.JSON(200, VersionListTasks)
	}
}

func QueryTrainJobListApi(ctx *context.Context) ([]*models.Cloudbrain, error) {
	repoId := ctx.QueryInt64("repoId")
	VersionListTasks, count, err := models.QueryModelTrainJobList(repoId)
	log.Info("query return count=" + fmt.Sprint(count))

	return VersionListTasks, err
}

func QueryTrainModelFileById(ctx *context.Context) ([]storage.FileInfo, error) {
	JobID := ctx.Query("jobId")
	VersionListTasks, count, err := models.QueryModelTrainJobVersionList(JobID)
	if err == nil {
		if count == 1 {
			task := VersionListTasks[0]
			jobName := task.JobName
			VersionName := task.VersionName
			taskConfig := task.GetCloudbrainConfig()
			storageType := storage_helper.GetStorageIntTypeFromStorageType(entity.StorageType(taskConfig.OutputStorageType))

			modelDbResult, err := getModelFromObjectSave(jobName, task.ComputeResource, storageType, VersionName)
			return modelDbResult, err
		}
	}
	log.Info("get TypeCloudBrainTwo TrainJobListModel failed:", err)
	return nil, errors.New("Not found task.")
}

func getModelFromObjectSave(jobName string, computeResource string, storageType int, VersionName string) ([]storage.FileInfo, error) {
	helper := storage_helper.SelectStorageHelperFromStorageIntType(storageType)
	if helper == nil {
		return nil, errors.New("Not support.")

	}
	objectkey := helper.GetOutputObjectKeyPrefix(jobName, computeResource, VersionName)
	bucket := helper.GetBucket()

	modelDbResult, err := helper.GetAllObjectByBucketAndPrefix(setting.Bucket, objectkey)
	log.Info("bucket=" + bucket + "  objectkey=" + objectkey)
	if err != nil {
		log.Info("get TypeCloudBrainTwo TrainJobListModel failed:", err)
		return nil, err
	} else {
		return modelDbResult, nil
	}

}

func QueryTrainModelList(ctx *context.Context) {
	log.Info("query train job list. start.")
	jobName := ctx.Query("jobName")
	VersionName := ctx.Query("versionName")
	if VersionName == "" {
		VersionName = ctx.Query("VersionName")
	}
	task, err := models.GetCloudbrainByName(jobName)
	if err != nil {
		log.Info("get task failed:", err)
		ctx.JSON(200, "")
		return
	}
	taskConfig := task.GetCloudbrainConfig()
	storageType := storage_helper.GetStorageIntTypeFromStorageType(entity.StorageType(taskConfig.OutputStorageType))
	modelDbResult, err := getModelFromObjectSave(jobName, task.ComputeResource, storageType, VersionName)
	if err != nil {
		log.Info("get TypeCloudBrainTwo TrainJobListModel failed:", err)
		ctx.JSON(200, "")
	} else {
		ctx.JSON(200, modelDbResult)
		return
	}
}

func DownloadSingleModelFile(ctx *context.Context) {
	log.Info("DownloadSingleModelFile start.")
	id := ctx.Params(":ID")
	parentDir := ctx.Query("parentDir")
	fileName := ctx.Query("fileName")
	path := Model_prefix + models.AttachmentRelativePath(id) + "/" + parentDir + fileName
	task, err := models.QueryModelById(id)
	if err != nil {
		log.Error("no such model!", err.Error())
		ctx.ServerError("no such model:", err)
		return
	}
	if !isCanDownload(ctx, task) {
		ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
		return
	}
	if ctx.User != nil {
		log.Info("download  user is  %v, model is %v", ctx.User.ID, id)
	}

	if task.Type == models.StorageTypeObs {

		url, err := storage.GetObsCreateSignedUrlByBucketAndKey(setting.Bucket, path)
		if err != nil {
			log.Error("GetObsCreateSignedUrl failed: %v", err.Error(), ctx.Data["msgID"])
			ctx.ServerError("GetObsCreateSignedUrl", err)
			return
		}
		//count++
		models.ModifyModelDownloadCount(id)
		http.Redirect(ctx.Resp, ctx.Req.Request, url, http.StatusMovedPermanently)

	} else if task.Type == models.StorageTypeMinio {
		log.Info("start to down load minio file.")
		url, err := storage.Attachments.PresignedGetURL(path, fileName)
		if err != nil {
			log.Error("Get minio get SignedUrl failed: %v", err.Error(), ctx.Data["msgID"])
			ctx.ServerError("Get minio get SignedUrl failed", err)
			return
		}
		models.ModifyModelDownloadCount(id)
		http.Redirect(ctx.Resp, ctx.Req.Request, url, http.StatusMovedPermanently)
	}

}

func ShowModelInfo(ctx *context.Context) {
	ctx.Data["ID"] = ctx.Query("id")
	ctx.Data["name"] = ctx.Query("name")
	ctx.Data["isModelManage"] = true
	ctx.Data["ModelManageAccess"] = ctx.Repo.CanWrite(models.UnitTypeModelManage)

	ctx.HTML(200, tplModelInfo)
}

func QueryModelObjById(ctx *context.Context) *models.AiModelManage {
	id := ctx.Query("id")
	model, err := models.QueryModelById(id)
	if err == nil {
		model.IsCanOper = isOperModifyOrDelete(ctx, model.UserId)
		model.IsCanDelete = isCanDelete(ctx, model.UserId)
		model.IsCanDownload = isCanDownload(ctx, model)
		removeIpInfo(model)
		return model
	} else {
		return nil
	}
}

func QueryModelById(ctx *context.Context) {
	model := QueryModelObjById(ctx)
	if model != nil {
		ctx.JSON(http.StatusOK, model)
	} else {
		ctx.JSON(http.StatusNotFound, nil)
	}
}

func QueryModelObjByName(ctx *context.Context) []*models.AiModelManage {
	name := ctx.Query("name")
	log.Info("Show single ModelInfo start.name=" + name)
	modelArrays := models.QueryModelByName(name, ctx.Repo.Repository.ID)
	modelResult := make([]*models.AiModelManage, 0)
	isCanReadPrivateModel := isQueryPrivateModel(ctx)
	userIds := make([]int64, len(modelArrays))
	for i, model := range modelArrays {
		model.IsCanOper = isOperModifyOrDelete(ctx, model.UserId)
		model.IsCanDownload = isCanDownload(ctx, model)
		model.IsCanDelete = isCanDelete(ctx, model.UserId)
		model.RepoName = ctx.Repo.Repository.Name
		model.RepoOwnerName = ctx.Repo.Repository.OwnerName
		model.RepoDisplayName = ctx.Repo.Repository.DisplayName()

		userIds[i] = model.UserId
		if ctx.User != nil {
			re := models.QueryModelCollectByUserId(model.ID, ctx.User.ID)
			if re != nil && len(re) > 0 {
				model.IsCollected = true
			}
		}
		if model.IsPrivate {
			if !isCanReadPrivateModel {
				continue
			}
		}
		setModelRelateInfo(model)
		modelResult = append(modelResult, model)
	}
	userNameMap := models.QueryUserName(userIds)

	for _, model := range modelResult {
		removeIpInfo(model)
		value := userNameMap[model.UserId]
		if value != nil {
			model.UserName = value.Name
			model.UserRelAvatarLink = value.RelAvatarLink()
		}
	}
	return modelResult
}

func ShowSingleModel(ctx *context.Context) {
	ctx.JSON(http.StatusOK, QueryModelObjByName(ctx))
}

func removeIpInfo(model *models.AiModelManage) {
	reg, _ := regexp.Compile(`[[:digit:]]{1,3}\.[[:digit:]]{1,3}\.[[:digit:]]{1,3}\.[[:digit:]]{1,3}`)
	taskInfo := model.TrainTaskInfo
	taskInfo = reg.ReplaceAllString(taskInfo, "")
	model.TrainTaskInfo = taskInfo
}

func ShowOneVersionOtherModel(ctx *context.Context) {
	repoId := ctx.Repo.Repository.ID
	name := ctx.Query("name")
	aimodels := models.QueryModelByName(name, repoId)

	userIds := make([]int64, len(aimodels))
	for i, model := range aimodels {
		model.IsCanOper = isOperModifyOrDelete(ctx, model.UserId)
		model.IsCanDownload = isCanDownload(ctx, model)
		model.IsCanDelete = isCanDelete(ctx, model.UserId)
		userIds[i] = model.UserId
	}
	userNameMap := models.QueryUserName(userIds)

	for _, model := range aimodels {
		removeIpInfo(model)
		value := userNameMap[model.UserId]
		if value != nil {
			model.UserName = value.Name
			model.UserRelAvatarLink = value.RelAvatarLink()
		}
	}

	if len(aimodels) > 0 {
		ctx.JSON(200, aimodels[1:])
	} else {
		ctx.JSON(200, aimodels)
	}
}

func SetModelCount(ctx *context.Context) {
	isQueryPrivate := isQueryPrivateModel(ctx)
	repoId := ctx.Repo.Repository.ID
	Type := -1
	_, count, _ := models.QueryModel(&models.AiModelQueryOptions{
		ListOptions: models.ListOptions{
			Page:     1,
			PageSize: 2,
		},
		RepoID:         repoId,
		Type:           Type,
		New:            MODEL_LATEST,
		IsOnlyThisRepo: true,
		Status:         -1,
		FrameFilter:    -1,
		IsQueryPrivate: isQueryPrivate,
	})
	ctx.Data["MODEL_COUNT"] = count
}

func ShowModelTemplate(ctx *context.Context) {
	ctx.Data["isModelManage"] = true
	repoId := ctx.Repo.Repository.ID
	SetModelCount(ctx)
	ctx.Data["ModelManageAccess"] = ctx.Repo.CanWrite(models.UnitTypeModelManage)
	_, trainCount, _ := models.QueryModelTrainJobList(repoId)
	log.Info("query train count=" + fmt.Sprint(trainCount))
	ctx.Data["TRAIN_COUNT"] = trainCount
	ctx.Data["max_model_size"] = setting.MaxModelSize * MODEL_MAX_SIZE
	ctx.HTML(200, tplModelManageIndex)
}

func isQueryRight(ctx *context.Context) bool {
	if ctx.Repo.Repository.IsPrivate {
		if ctx.Repo.CanRead(models.UnitTypeModelManage) || ctx.User.IsAdmin || ctx.Repo.IsAdmin() || ctx.Repo.IsOwner() {
			return true
		}
		return false
	} else {
		return true
	}
}

func isCanDownload(ctx *context.Context, task *models.AiModelManage) bool {
	if !task.IsPrivate {
		return true
	}
	if ctx.User == nil {
		return false
	}
	isCollaborator, err := ctx.Repo.Repository.IsCollaborator(ctx.User.ID)
	if err != nil {
		log.Info("query error.")
	}
	isTeamMember, err := ctx.Repo.Repository.IsInRepoTeam(ctx.User.ID)
	if err != nil {
		log.Info("query IsInRepoTeam error." + err.Error())
	}
	if ctx.User.IsAdmin || ctx.User.ID == task.UserId || isCollaborator || isTeamMember {
		return true
	}
	if ctx.Repo.IsOwner() {
		return true
	}
	if !task.IsPrivate {
		return true
	}
	return false
}

func isModelCanDownload(user *models.User, repo *models.Repository, task *models.AiModelManage) bool {
	if user == nil {
		return false
	}
	//log.Info("enter here1... ")
	if repo == nil {
		log.Info("repo is nil here1... ")
	}
	isCollaborator, err := repo.IsCollaborator(user.ID)
	if err != nil {
		log.Info("query error.")
	}
	//log.Info("enter here2...")
	isTeamMember, err := repo.IsInRepoTeam(user.ID)
	if err != nil {
		log.Info("query IsInRepoTeam error." + err.Error())
	}
	//log.Info("enter here3...")
	if user.IsAdmin || user.ID == task.UserId || isCollaborator || isTeamMember {
		return true
	}
	if repo.OwnerID == user.ID {
		return true
	}
	if !task.IsPrivate {
		return true
	}
	return false
}

func isQueryPrivateModel(ctx *context.Context) bool {
	if ctx.User == nil {
		return false
	}
	isCollaborator, err := ctx.Repo.Repository.IsCollaborator(ctx.User.ID)
	if err != nil {
		log.Info("query IsCollaborator error." + err.Error())
	}
	isTeamMember, err := ctx.Repo.Repository.IsInRepoTeam(ctx.User.ID)
	if err != nil {
		log.Info("query IsInRepoTeam error." + err.Error())
	}
	if ctx.User.IsAdmin || isCollaborator || isTeamMember {
		return true
	}
	if ctx.Repo.IsOwner() {
		return true
	}
	return false
}

func isCanDelete(ctx *context.Context, modelUserId int64) bool {
	if ctx.User == nil {
		return false
	}
	if ctx.User.ID == modelUserId {
		return true
	}
	return isAdminRight(ctx)
}

func isAdminRight(ctx *context.Context) bool {
	if ctx.User.IsAdmin {
		return true
	}
	if ctx.Repo.IsOwner() {
		return true
	}
	permission, err := models.GetUserRepoPermission(ctx.Repo.Repository, ctx.User)
	if err != nil {
		log.Error("GetUserRepoPermission failed:%v", err.Error())
		return false
	} else {
		log.Info("permission.AccessMode=" + string(permission.AccessMode))
	}
	if permission.AccessMode >= models.AccessModeAdmin {
		return true
	}
	return false
}

func isCanReadPrivateModel(ctx *context.Context, modelUserId int64) bool {
	if ctx.User == nil {
		log.Info("user is nil")
		return false
	}
	if ctx.User.IsAdmin || ctx.User.ID == modelUserId {
		return true
	}
	return isAdminRight(ctx)
}

func isOperModifyOrDelete(ctx *context.Context, modelUserId int64) bool {
	if ctx.User == nil {
		log.Info("user is nil")
		return false
	}
	if ctx.User.IsAdmin || ctx.User.ID == modelUserId {
		return true
	}
	return isAdminRight(ctx)
}

func IsOperModifyOrDelete(ctx *context.Context, modelUserId int64) bool {
	return isOperModifyOrDelete(ctx, modelUserId)
}

func ShowModelPageInfo(ctx *context.Context) {
	log.Info("ShowModelInfo start.")
	if !isQueryRight(ctx) {
		ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
		return
	}
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := ctx.QueryInt("pageSize")
	if pageSize <= 0 {
		pageSize = setting.UI.IssuePagingNum
	}
	isQueryPrivate := isQueryPrivateModel(ctx)
	repoId := ctx.Repo.Repository.ID
	Type := -1
	modelResult, count, err := models.QueryModel(&models.AiModelQueryOptions{
		ListOptions: models.ListOptions{
			Page:     page,
			PageSize: pageSize,
		},
		RepoID:         repoId,
		Type:           Type,
		New:            MODEL_LATEST,
		IsOnlyThisRepo: true,
		Status:         -1,
		IsQueryPrivate: isQueryPrivate,
		FrameFilter:    -1,
	})
	if err != nil {
		ctx.ServerError("Cloudbrain", err)
		return
	}

	userIds := make([]int64, len(modelResult))
	for i, model := range modelResult {
		model.IsCanOper = isOperModifyOrDelete(ctx, model.UserId)
		model.IsCanDelete = isCanDelete(ctx, model.UserId)
		model.IsCanDownload = isCanDownload(ctx, model)
		userIds[i] = model.UserId
	}

	userNameMap := models.QueryUserName(userIds)

	for _, model := range modelResult {
		removeIpInfo(model)
		value := userNameMap[model.UserId]
		if value != nil {
			model.UserName = value.Name
			model.UserRelAvatarLink = value.RelAvatarLink()
		}
	}

	mapInterface := make(map[string]interface{})
	mapInterface["data"] = modelResult
	mapInterface["count"] = count
	ctx.JSON(http.StatusOK, mapInterface)
}

func ModifyModelPrivate(ctx *context.Context) {
	id := ctx.Query("id")
	isPrivate := ctx.QueryBool("isPrivate")
	re := map[string]string{
		"code": "-1",
	}
	task, err := models.QueryModelById(id)
	if err != nil || task == nil {
		re["msg"] = err.Error()
		log.Error("no such model!", err.Error())
		ctx.JSON(200, re)
		return
	}
	if !isOperModifyOrDelete(ctx, task.UserId) {
		re["msg"] = "No right to operation."
		ctx.JSON(200, re)
		return
	}
	err = models.ModifyModelPrivate(id, isPrivate)
	if err == nil {
		re["code"] = "0"
		ctx.JSON(200, re)
		log.Info("modify success.")
	} else {
		re["msg"] = err.Error()
		ctx.JSON(200, re)
		log.Info("Failed to modify.id=" + id + "  isprivate=" + fmt.Sprint(isPrivate) + " error:" + err.Error())
	}

}

func ModifyModelInfo(ctx *context.Context) {
	log.Info("modify model start.")
	id := ctx.Query("id")
	re := map[string]string{
		"code": "-1",
	}
	task, err := models.QueryModelById(id)
	if err != nil {
		re["msg"] = err.Error()
		log.Error("no such model!", err.Error())
		ctx.JSON(200, re)
		return
	}
	if !isOperModifyOrDelete(ctx, task.UserId) {
		re["msg"] = "No right to operation."
		ctx.JSON(200, re)
		return
	}

	name := ctx.Query("name")
	label := ctx.Query("label")
	description := ctx.Query("description")
	engine := ctx.QueryInt("engine")
	isPrivate := ctx.QueryBool("isPrivate")
	license := ctx.Query("license")
	aimodels := models.QueryModelByName(name, task.RepoId)
	if aimodels != nil && len(aimodels) > 0 {
		if len(aimodels) == 1 {
			if aimodels[0].ID != task.ID {
				re["msg"] = ctx.Tr("repo.model.manage.create_error")
				ctx.JSON(200, re)
				return
			}
		} else {
			re["msg"] = ctx.Tr("repo.model.manage.create_error")
			ctx.JSON(200, re)
			return
		}
	}
	err = models.ModifyLocalModel(id, name, label, description, engine, isPrivate, license)
	if task.Name != name {
		aimodels = models.QueryModelByName(task.Name, task.RepoId)
		if aimodels != nil && len(aimodels) > 0 {
			for _, model := range aimodels {
				models.ModifyLocalModel(model.ID, name, model.Label, model.Description, int(model.Engine), model.IsPrivate, model.License)
			}
		}
	}
	if license != task.License {
		mdContent := getMDContent(task)
		if mdContent != "" {
			if task.License != "" {
				log.Info("start to remove license.")
				mdContent = removeLicenseFromMD(mdContent, task.License)
			}
			mdContent = addLicenseMDToReadme(mdContent, license)
			updateReadMeMd(task, mdContent)
		}
	}
	if err != nil {
		re["msg"] = err.Error()
		ctx.JSON(200, re)
		return
	} else {
		re["code"] = "0"
		ctx.JSON(200, re)
	}
}

func QueryModelListForPredict(ctx *context.Context) {
	repoId := ctx.Repo.Repository.ID
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = -1
	}
	pageSize := ctx.QueryInt("pageSize")
	if pageSize <= 0 {
		pageSize = -1
	}
	isQueryPrivate := isQueryPrivateModel(ctx)
	//IsOnlyThisRepo := ctx.QueryBool("isOnlyThisRepo")
	modelResult, count, err := models.QueryModel(&models.AiModelQueryOptions{
		ListOptions: models.ListOptions{
			Page:     page,
			PageSize: pageSize,
		},
		RepoID:         repoId,
		Type:           -1,
		New:            -1,
		Status:         0,
		IsOnlyThisRepo: true,
		IsQueryPrivate: isQueryPrivate,
		FrameFilter:    -1,
	})
	if err != nil {
		ctx.ServerError("Cloudbrain", err)
		return
	}
	log.Info("query return count=" + fmt.Sprint(count))

	nameList := make([]string, 0)

	nameMap := make(map[string][]*models.AiModelManage)
	for _, model := range modelResult {
		model.TrainTaskInfo = ""
		model.Accuracy = ""
		//removeIpInfo(model)
		if _, value := nameMap[model.Name]; !value {
			models := make([]*models.AiModelManage, 0)
			models = append(models, model)
			nameMap[model.Name] = models
			nameList = append(nameList, model.Name)
		} else {
			nameMap[model.Name] = append(nameMap[model.Name], model)
		}
	}

	mapInterface := make(map[string]interface{})
	mapInterface["nameList"] = nameList
	mapInterface["nameMap"] = nameMap
	ctx.JSON(http.StatusOK, mapInterface)
}

func QueryModelFileForPredict(ctx *context.Context) {
	id := ctx.Query("id")
	if id == "" {
		id = ctx.Query("ID")
	}
	ctx.JSON(http.StatusOK, QueryModelFileByID(id))
}

func QueryModelFileByID(id string) []storage.FileInfo {
	model, err := models.QueryModelById(id)
	if err != nil {
		log.Error("no such model!", err.Error())
		return nil
	}
	return modelmanage.QueryModelFileByModel(model)
}

func QueryOneLevelModelFile(ctx *context.Context) {
	id := ctx.Query("id")
	if id == "" {
		id = ctx.Query("ID")
	}
	parentDir := ctx.Query("parentDir")
	model, err := models.QueryModelById(id)
	if err != nil {
		log.Error("no such model!", err.Error())
		ctx.JSON(http.StatusOK, nil)
		return
	}
	if model.IsPrivate && !isCanReadPrivateModel(ctx, model.UserId) {
		ctx.JSON(401, map[string]string{"msg": "No right to read."})
		return
	}
	ctx.JSON(http.StatusOK, queryOneLevelModelFile(model, parentDir))
}

func queryOneLevelModelFile(model *models.AiModelManage, parentDir string) []storage.FileInfo {
	fileinfos := make([]storage.FileInfo, 0)
	if model.Type == models.StorageTypeObs {
		log.Info("obs list model file.model.Path=" + model.Path)
		if len(model.Path) > len(setting.Bucket)+1 {
			prefix := model.Path[len(setting.Bucket)+1:] + parentDir
			fileinfos, _ = storage.GetOneLevelAllObjectUnderDir(setting.Bucket, prefix, "")
		}
	} else if model.Type == models.StorageTypeMinio {
		log.Info("minio list model file.model.Path=" + model.Path)
		if len(model.Path) > len(setting.Attachment.Minio.Bucket)+1 {
			prefix := model.Path[len(setting.Attachment.Minio.Bucket)+1:] + parentDir
			fileinfos, _ = storage.GetOneLevelAllObjectUnderDirMinio(setting.Attachment.Minio.Bucket, prefix, "")
		}
	}
	if fileinfos == nil {
		fileinfos = make([]storage.FileInfo, 0)
	}
	return fileinfos
}

func CreateLocalModel(ctx *context.Context) {
	ctx.Data["isModelManage"] = true
	ctx.Data["ModelManageAccess"] = ctx.Repo.CanWrite(models.UnitTypeModelManage)

	ctx.HTML(200, tplCreateLocalModelInfo)
}

func CreateLocalModelForUpload(ctx *context.Context) {
	ctx.Data["uuid"] = ctx.Query("uuid")
	ctx.Data["isModelManage"] = true
	ctx.Data["ModelManageAccess"] = ctx.Repo.CanWrite(models.UnitTypeModelManage)
	ctx.Data["max_model_size"] = setting.MaxModelSize * MODEL_MAX_SIZE
	ctx.HTML(200, tplCreateLocalForUploadModelInfo)
}

func CreateOnlineModel(ctx *context.Context) {
	ctx.Data["isModelManage"] = true
	ctx.Data["ModelManageAccess"] = ctx.Repo.CanWrite(models.UnitTypeModelManage)

	ctx.HTML(200, tplCreateOnlineModelInfo)
}

func QueryModelCollectNum(ctx *context.Context) {
	id := ctx.Query("id")
	record := models.QueryModelCollectNum(id)
	ctx.JSON(200, record)
}

func ModelCollect(ctx *context.Context) {
	id := ctx.Query("id")
	isCollected := ctx.QueryBool("collected")
	re := map[string]string{
		"code": "-1",
	}
	task, err := models.QueryModelById(id)
	if err != nil || task == nil {
		re["msg"] = err.Error()
		log.Error("no such model!", err.Error())
		ctx.JSON(200, re)
		return
	}
	if ctx.User == nil {
		re["msg"] = "user not login."
		re["code"] = "401"
		ctx.JSON(200, re)
		return
	}
	record := models.QueryModelCollectByUserId(id, ctx.User.ID)
	if isCollected {
		if record == nil || len(record) == 0 {
			log.Info("user collect the model.user id=" + fmt.Sprint(ctx.User.ID) + " model id=" + id)
			err := models.SaveModelCollect(&models.AiModelCollect{
				ModelID: id,
				UserId:  ctx.User.ID,
			})
			if err == nil {
				re["code"] = "0"
			} else {
				re["msg"] = err.Error()
			}
		}
	} else {
		if record != nil && len(record) > 0 {
			log.Info("user delete collect the model.user id=" + fmt.Sprint(ctx.User.ID) + " model id=" + id)
			err := models.DeleteModelCollect(&models.AiModelCollect{
				ID: record[0].ID,
			})
			if err == nil {
				re["code"] = "0"
			} else {
				re["msg"] = err.Error()
			}
		}
	}
	num := models.QueryModelCollectNum(id)
	models.ModifyModelCollectedNum(id, num)
	ctx.JSON(200, re)
}

func SaveHfModel(ctx *context.Context, modelName string, label string, description string, engine int, isPrivate bool) (*models.AiModelManage, error) {
	log.Info("[hf_model]SaveHfModel: save SaveHfModel start.")
	uuid := uuid.NewV4()
	id := uuid.String()
	name := modelName
	version := "0.0.1"
	taskType := 1
	license := ""
	if ctx.Repo.Repository.IsPrivate {
		msg := "SaveHfModel: Private repo cannot create public model"
		log.Error(msg)
		return nil, errors.New(msg)
	}
	modelActualPath := ""
	computeResource := ""
	destKeyNamePrefix := Model_prefix + models.AttachmentRelativePath(id) + "/"
	modelActualPath = setting.Bucket + "/" + destKeyNamePrefix
	computeResource = models.NPUResource
	var lastNewModelId string
	repoId := ctx.Repo.Repository.ID
	aimodels := models.QueryModelByName(name, repoId)
	if len(aimodels) > 0 {
		for _, model := range aimodels {
			if model.Version == version {
				msg := ctx.Tr("repo.model.manage.create_error")
				log.Error(msg)
				return nil, errors.New(msg)
			}
			if model.New == MODEL_LATEST {
				lastNewModelId = model.ID
			}
		}
	}
	model := &models.AiModelManage{
		ID:              id,
		Version:         version,
		ModelType:       models.MODEL_HF_TYPE,
		VersionCount:    len(aimodels) + 1,
		Label:           label,
		Name:            name,
		Description:     description,
		New:             MODEL_LATEST,
		Type:            taskType,
		Path:            modelActualPath,
		Size:            0,
		AttachmentId:    "",
		RepoId:          repoId,
		UserId:          ctx.User.ID,
		Engine:          int64(engine),
		TrainTaskInfo:   "",
		Accuracy:        "",
		Status:          STATUS_COPY_MODEL,
		IsPrivate:       isPrivate,
		ComputeResource: computeResource,
		License:         license,
	}

	err := models.SaveModelToDb(model)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	if len(lastNewModelId) > 0 {
		models.ModifyModelNewProperty(lastNewModelId, MODEL_NOT_LATEST, 0)
	}
	var units []models.RepoUnit
	var deleteUnitTypes []models.UnitType
	units = append(units, models.RepoUnit{
		RepoID: ctx.Repo.Repository.ID,
		Type:   models.UnitTypeModelManage,
		Config: &models.ModelManageConfig{
			EnableModelManage: true,
		},
	})
	deleteUnitTypes = append(deleteUnitTypes, models.UnitTypeModelManage)

	models.UpdateRepositoryUnits(ctx.Repo.Repository, units, deleteUnitTypes)

	log.Info("[hf_model] SaveHfModel: save hf model end.")
	if !model.IsPrivate {
		notification.NotifyOtherTask(ctx.User, ctx.Repo.Repository, id, name, models.ActionCreateNewModelTask)
	}
	return model, nil
}
