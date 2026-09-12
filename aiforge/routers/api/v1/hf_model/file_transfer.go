package hf_model

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
	"time"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
	"code.gitea.io/gitea/services/subject_service"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
)

func StartTransferFilesAPI(ctx *context.APIContext) {
	numFiles := ctx.QueryInt("num_files")
	hfRepoId := ctx.Query("hf_repo_id")
	heartbeatSec := ctx.QueryInt("heartbeat_sec")
	resp := initResponseJson()

	// 获取模型所有者
	modelOwner, err := GetModelOwner()
	if err != nil {
		resp.Msg = ctx.Tr("hf_model.internal_error", err.Error())
		ctx.JSON(http.StatusOK, resp)
		return
	}

	// 从数据库获取等待传输的文件
	waitlist, err := models.GetWaitlist(modelOwner, numFiles, hfRepoId)
	if err != nil {
		resp.Msg = ctx.Tr("hf_model.internal_error", "cannot get waitlist")
		ctx.JSON(http.StatusOK, resp)
		return
	}

	modelIds := make(map[string]bool)
	// 更新文件状态并记录心跳时间
	for _, file := range waitlist {
		models.UpdateHfFileStatusAndHeartbeat(file.ID, models.HfTransferStatusOngoing, time.Now(), heartbeatSec)
		if _, exists := modelIds[file.ModelId]; !exists {
			modelIds[file.ModelId] = true
		}
	}

	go setAimodelsOngoing(modelIds)
	// 返回响应
	resp.Code = 1
	resp.Msg = "success"
	resp.Data = map[string]interface{}{
		"count": len(waitlist),
		"files": waitlist,
	}
	ctx.JSON(http.StatusOK, resp)
}

func UpdateFileStatusAPI(ctx *context.APIContext) {
	fileID := ctx.QueryInt64("file_id")
	fileStatus := ctx.QueryInt("status") // 1: success, 0: ongoing; 2: failed
	resp := initResponseJson()

	fileRecord, _ := models.QueryFileById(fileID)
	err := models.UpdateHfFileStatus(fileID, fileStatus)
	if err != nil {
		log.Error("[hf_model] failed to update: %v", err)
		resp.Msg = err.Error()
		ctx.JSON(http.StatusOK, resp)
		return
	}

	_, err = CompleteAiModel(fileRecord.ModelId)
	if err != nil {
		log.Error("[hf_model] failed to complete ai model: %v", err)
		resp.Msg = err.Error()
		ctx.JSON(http.StatusOK, resp)
		return
	}

	resp.Code = 1
	resp.Msg = "success"
	ctx.JSON(http.StatusOK, resp)
}

// added but no use in V20260319
func BatchUpdateFileStatusAPI(ctx *context.APIContext) {
	fileIdStrs := ctx.Req.URL.Query()["file_id"]
	fileStatus := ctx.QueryInt("status") // 1: success, 0: ongoing; 2: failed
	resp := initResponseJson()

	if len(fileIdStrs) == 0 {
		resp.Msg = "file_id is required"
		ctx.JSON(http.StatusOK, resp)
		return
	}

	// Parse all file IDs
	fileIDs := make([]int64, 0, len(fileIdStrs))
	for _, idStr := range fileIdStrs {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			log.Error("[hf_model] invalid file_id: %s", idStr)
			continue
		}
		fileIDs = append(fileIDs, id)
	}

	if len(fileIDs) == 0 {
		resp.Msg = "no valid file_id provided"
		ctx.JSON(http.StatusOK, resp)
		return
	}

	// Heartbeat only (status=0): batch update last_heartbeat timestamp
	if fileStatus == models.HfTransferStatusOngoing {
		err := models.BatchUpdateHfFileHeartbeat(fileIDs)
		if err != nil {
			log.Error("[hf_model] batch heartbeat update failed: %v", err)
			resp.Msg = err.Error()
			ctx.JSON(http.StatusOK, resp)
			return
		}
		resp.Code = 1
		resp.Msg = "success"
		ctx.JSON(http.StatusOK, resp)
		return
	}

	// Completion/failure (status=1 or 2): batch update status
	err := models.BatchUpdateHfFileStatus(fileIDs, fileStatus)
	if err != nil {
		log.Error("[hf_model] batch status update failed: %v", err)
		resp.Msg = err.Error()
		ctx.JSON(http.StatusOK, resp)
		return
	}

	// Check model completion for affected models
	modelIds := make(map[string]bool)
	for _, fid := range fileIDs {
		fileRecord, _ := models.QueryFileById(fid)
		if fileRecord != nil {
			modelIds[fileRecord.ModelId] = true
		}
	}
	for modelId := range modelIds {
		_, err := CompleteAiModel(modelId)
		if err != nil {
			log.Error("[hf_model] failed to complete ai model %s: %v", modelId, err)
		}
	}

	resp.Code = 1
	resp.Msg = "success"
	ctx.JSON(http.StatusOK, resp)
}

func CheckHfHeartbeats() {
	CheckExpiredFiles()
	CheckIncompleteModels()
}

func CheckExpiredFiles() {
	files, err := models.GetRunningFiles()
	if err != nil {
		log.Error("[hf_model] failed to get running files: %v", err)
		return
	}
	for _, record := range files {
		lastHeartbeat := record.LastHeartbeat
		heartbeatSec := record.HeartbeatSec
		log.Info(fmt.Sprintf("[hf_model] Running file: %v,%s", record.ID, record.Filename))

		timeMin := time.Duration(heartbeatSec+30) * time.Second
		log.Info(fmt.Sprintf("[hf_model] %v timeMin lastHeartbeat: %v", record.ID, timeMin))
		timeSince := time.Since(lastHeartbeat.AsTime())
		log.Info(fmt.Sprintf("[hf_model] %v timeSince lastHeartbeat: %v", record.ID, timeSince))

		if timeSince > timeMin {
			aimodel, _ := models.GetAimodelByID(record.ModelId)
			helper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(aimodel.StorageType))
			fileObjectKey := path.Join(aimodel.Path, record.Filename)
			hasFile, err := helper.HasObject(fileObjectKey)
			if err != nil {
				log.Error("[hf_model]:Error checking file:" + err.Error())
			}
			log.Info(fmt.Sprintf("[hf_model] hasFile: %v", hasFile))

			statusNew := models.HfTransferStatusFailure
			if hasFile {
				statusNew = models.HfTransferStatusSuccess
			}
			log.Info("[hf_model]: file fileStatusNew after checking: " + fmt.Sprint(statusNew))
			err = models.UpdateHfFileStatus(record.ID, statusNew)
			if err != nil {
				log.Error("[hf_model]:Error updating transfer record on file: " + record.Filename + err.Error())
			}
		}
	}
}

func CheckIncompleteModels() {
	hfModels, err := models.QueryHfModelByStatus(models.HfTransferStatusOngoing)
	if err != nil {
		log.Error("[hf_model] failed to query running models: %v", err)
		return
	}
	log.Info(fmt.Sprintf("[hf_model] hfModels Running hfModels: %v", hfModels))
	for _, model := range hfModels {
		completed, err := CompleteAiModel(model.ID)
		if err != nil {
			log.Error("[hf_model] failed to complete ai model: %v", err)
		}
		log.Info(fmt.Sprintf("[hf_model]  %v completed=%v", model.Name, completed))
	}
}

func CompleteAiModel(modelId string) (bool, error) {
	isComplete, _ := models.IsTransferComplete(modelId)
	if isComplete {
		allSuccess, _ := models.IsTransferSuccess(modelId)
		if allSuccess {
			subject_service.DoAfterAimodelFileChanged(modelId)
			err := models.UpdateAimodelByStatus(modelId, models.HfTransferStatusSuccess, "")
			if err != nil {
				log.Error("UpdateAimodelByStatus faild." + err.Error())
			}
		} else {
			err := models.UpdateAimodelByStatus(modelId, models.HfTransferStatusFailure, "")
			if err != nil {
				log.Error("UpdateAimodelByStatus faild." + err.Error())
			}
		}
	}
	return isComplete, nil
}

func setAimodelsOngoing(modelIdsMap map[string]bool) {
	for id := range modelIdsMap {
		err := models.UpdateAimodelByStatus(id, models.HfTransferStatusOngoing, "")
		if err != nil {
			log.Error("SetAimodelsOngoing failed: " + err.Error())
		}
	}
}
