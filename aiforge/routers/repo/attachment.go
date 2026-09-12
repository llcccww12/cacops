// Copyright 2017 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repo

import (
	contexExt "context"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"path"
	"strconv"
	"strings"

	"code.gitea.io/gitea/modules/auth"

	"code.gitea.io/gitea/modules/base"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/labelmsg"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/minio_ext"
	"code.gitea.io/gitea/modules/notification"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
	"code.gitea.io/gitea/modules/upload"
	"code.gitea.io/gitea/modules/worker"
	repo_service "code.gitea.io/gitea/services/repository"
	"code.gitea.io/gitea/services/storage_limit"
	gouuid "github.com/satori/go.uuid"
)

const (
	//result of decompress
	DecompressSuccess                = "0"
	DecompressFailed                 = "1"
	tplAttachmentUpload base.TplName = "repo/attachment/upload"
	tplAttachmentEdit   base.TplName = "repo/attachment/edit"
)

type CloudBrainDataset struct {
	UUID       string `json:"id"`
	Name       string `json:"name"`
	Path       string `json:"place"`
	UserName   string `json:"provider"`
	CreateTime string `json:"created_at"`
}

type UploadForm struct {
	UploadID   string         `form:"uploadId"`
	UuID       string         `form:"uuid"`
	PartSize   int64          `form:"size"`
	Offset     int64          `form:"offset"`
	PartNumber int            `form:"chunkNumber"`
	PartFile   multipart.File `form:"file"`
}

func RenderAttachmentSettings(ctx *context.Context) {
	renderAttachmentSettings(ctx)
}

func renderAttachmentSettings(ctx *context.Context) {
	ctx.Data["IsAttachmentEnabled"] = setting.Attachment.Enabled
	ctx.Data["AttachmentStoreType"] = setting.Attachment.StoreType
	ctx.Data["AttachmentAllowedTypes"] = setting.Attachment.AllowedTypes
	ctx.Data["AttachmentMaxSize"] = setting.Attachment.MaxSize
	ctx.Data["AttachmentMaxFiles"] = setting.Attachment.MaxFiles
}

func UploadAttachmentUI(ctx *context.Context) {
	ctx.Data["datasetId"] = ctx.Query("datasetId")
	ctx.Data["PageIsDataset"] = true
	ctx.Data["max_model_size"] = int64(setting.MaxModelSize * MODEL_MAX_SIZE)
	ctx.HTML(200, tplAttachmentUpload)

}

func EditAttachmentUI(ctx *context.Context) {

	id, _ := strconv.ParseInt(ctx.Params(":id"), 10, 64)
	ctx.Data["PageIsDataset"] = true
	attachment, _ := models.GetAttachmentByID(id)
	if attachment == nil {
		ctx.Error(404, "The attachment does not exits.")
	}
	ctx.Data["Attachment"] = attachment
	ctx.HTML(200, tplAttachmentEdit)

}

func EditAttachment(ctx *context.Context, form auth.EditAttachmentForm) {

	err := models.UpdateAttachmentDescription(&models.Attachment{
		ID:          form.ID,
		Description: form.Description,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("dataset.edit_attachment_fail")))
	}
	ctx.JSON(http.StatusOK, models.BaseOKMessage)

}

// UploadAttachment response for uploading issue's attachment
func UploadAttachment(ctx *context.Context) {
	if !setting.Attachment.Enabled {
		ctx.Error(404, "attachment is not enabled")
		return
	}

	file, header, err := ctx.Req.FormFile("file")
	if err != nil {
		ctx.Error(500, fmt.Sprintf("FormFile: %v", err))
		return
	}
	defer file.Close()

	buf := make([]byte, 1024)
	n, _ := file.Read(buf)
	if n > 0 {
		buf = buf[:n]
	}

	err = upload.VerifyAllowedContentType(buf, strings.Split(setting.Attachment.AllowedTypes, ","))
	if err != nil {
		ctx.Error(400, err.Error())
		return
	}

	datasetID, _ := strconv.ParseInt(ctx.Req.FormValue("dataset_id"), 10, 64)

	attach, err := models.NewAttachment(&models.Attachment{
		IsPrivate:  true,
		UploaderID: ctx.User.ID,
		Name:       header.Filename,
		DatasetID:  datasetID,
	}, buf, file)
	if err != nil {
		ctx.Error(500, fmt.Sprintf("NewAttachment: %v", err))
		return
	}

	log.Trace("New attachment uploaded: %s", attach.UUID)
	ctx.JSON(200, map[string]string{
		"uuid": attach.UUID,
	})
}

func UpdatePublicAttachment(ctx *context.Context) {
	file := ctx.Query("file")
	isPrivate, _ := strconv.ParseBool(ctx.Query("is_private"))
	attach, err := models.GetAttachmentByUUID(file)
	if err != nil {
		ctx.Error(404, err.Error())
		return
	}
	attach.IsPrivate = isPrivate
	if isPrivate {
		if err := models.UpdateAttachment(attach); err != nil {
			ctx.ServerError("UpdateAttachment", err)
			return
		}
	} else {

		dataset, err := models.GetDatasetByID(attach.DatasetID)
		if err != nil {
			ctx.Error(404, err.Error())
			return
		}
		repo, err := models.GetRepositoryByID(dataset.RepoID)
		if err != nil {
			ctx.Error(404, err.Error())
			return
		}
		if repo.IsPrivate {
			repo.IsPrivate = false
			if err := models.UpdateRepositoryAndAttachment(repo, attach); err != nil {
				ctx.ServerError("UpdateRepositoryAndAttachment", err)
				return
			}

		} else {
			if err := models.UpdateAttachment(attach); err != nil {
				ctx.ServerError("UpdateAttachment", err)
				return
			}
		}

	}

}

// DeleteAttachment response for deleting issue's attachment
func DeleteAttachment(ctx *context.Context) {
	file := ctx.Query("file")

	attach, err := models.GetAttachmentByUUID(file)
	if err != nil {
		ctx.Error(400, err.Error())
		return
	}

	//issue 214: mod del-dataset permission
	if !models.CanDelAttachment(ctx.IsSigned, ctx.User, attach) {
		ctx.Error(403)
		return
	}
	re := models.IsNeedToDeleteAttachPath(attach)
	go deleteAttachmentFile(attach, re)

	models.DeleteAttachmentByID(attach.ID)
	models.DeleteFileChunkById(attach.UUID)

	ctx.JSON(200, map[string]string{
		"uuid": attach.UUID,
	})
}

func BatchDeleteAttachment(ctx *context.Context) {
	uuids := ctx.Query("file")

	uuidList := strings.Split(uuids, ";")
	succeedUuidList := []string{}

	for _, uuid := range uuidList {
		attach, err := models.GetAttachmentByUUID(uuid)
		if err != nil {
			log.Warn("Get attachment (%s) failed:%s", uuid, err.Error())
			continue
		}

		//issue 214: mod del-dataset permission
		if !models.CanDelAttachment(ctx.IsSigned, ctx.User, attach) {
			log.Warn("Can not delete attachment (%s) for user:%s", uuid, ctx.User.Name)
			continue
		}
		re := models.IsNeedToDeleteAttachPath(attach)
		go deleteAttachmentFile(attach, re)

		models.DeleteAttachmentByID(attach.ID)
		models.DeleteFileChunkById(attach.UUID)
		succeedUuidList = append(succeedUuidList, uuid)

	}

	ctx.JSON(200, map[string][]string{
		"uuid": succeedUuidList,
	})

}

func deleteAttachmentFile(attach *models.Attachment, isNeedToDeleteFile bool) {
	repo_service.DecreaseRepoDatasetNum(attach.DatasetID)
	if isNeedToDeleteFile {
		models.DeleteSingleAttachment(attach)
		attachjson, _ := json.Marshal(attach)
		labelmsg.SendDeleteAttachToLabelSys(string(attachjson))
		DeleteAllUnzipFile(attach, "")
	}
}

func DownloadUserIsOrgOrCollaboration(ctx *context.Context, attach *models.Attachment) bool {
	dataset, err := models.GetDatasetByID(attach.DatasetID)
	if err != nil {
		log.Info("query dataset error")
	} else {
		repo, err := models.GetRepositoryByID(dataset.RepoID)
		if err != nil {
			log.Info("query repo error.")
		} else {
			repo.GetOwner()
			if ctx.User != nil {

				if repo.Owner.IsOrganization() {
					if repo.Owner.IsUserPartOfOrg(ctx.User.ID) {
						log.Info("org user may visit the attach.")
						return true
					}
				}
				isCollaborator, _ := repo.IsCollaborator(ctx.User.ID)
				if isCollaborator {
					log.Info("Collaborator user may visit the attach.")
					return true
				}
			}
		}
	}
	return false
}

// GetAttachment serve attachements
func GetAttachment(ctx *context.Context) {

	attach, err := models.GetAttachmentByUUID(ctx.Params(":uuid"))
	if err != nil {
		if models.IsErrAttachmentNotExist(err) {
			ctx.Error(404)
		} else {
			ctx.ServerError("GetAttachmentByUUID", err)
		}
		return
	}

	repository, unitType, err := attach.LinkedRepository()
	if err != nil {
		ctx.ServerError("LinkedRepository", err)
		return
	}
	dataSet, err := attach.LinkedDataSet()
	if err != nil {
		ctx.ServerError("LinkedDataSet", err)
		return
	}

	if repository == nil && dataSet != nil {
		repository, _ = models.GetRepositoryByID(dataSet.RepoID)
		unitType = models.UnitTypeDatasets
	}

	if repository == nil { //If not linked
		//if !(ctx.IsSigned && attach.UploaderID == ctx.User.ID) && attach.IsPrivate { //We block if not the uploader
		//log.Info("ctx.IsSigned =" + fmt.Sprintf("%v", ctx.IsSigned))
		if !(ctx.IsSigned && attach.UploaderID == ctx.User.ID) && attach.IsPrivate && !DownloadUserIsOrgOrCollaboration(ctx, attach) { //We block if not the uploader
			ctx.Error(http.StatusNotFound)
			return
		}

	} else { //If we have the repository we check access
		perm, errPermission := models.GetUserRepoPermission(repository, ctx.User)
		if errPermission != nil {
			ctx.Error(http.StatusInternalServerError, "GetUserRepoPermission", errPermission.Error())
			return
		}
		if !perm.CanRead(unitType) {
			ctx.Error(http.StatusNotFound)
			return
		}
	}

	if dataSet != nil {
		if !ctx.IsSigned && dataSet.IsPrivate() {
			ctx.SetCookie("redirect_to", setting.AppSubURL+ctx.Req.URL.RequestURI(), 0, setting.AppSubURL)
			ctx.Redirect(setting.AppSubURL + "/user/login")
			return
		} else {
			isPermit, err := models.GetUserDataSetPermission(dataSet, ctx.User)
			if err != nil {
				ctx.Error(http.StatusInternalServerError, "GetUserDataSetPermission", err.Error())
				return
			}
			if !isPermit {
				ctx.Error(http.StatusNotFound)
				return
			}
		}
	}
	if ctx.User != nil {
		log.Info("download  user is  %v, attachment is %v", ctx.User.ID, ctx.Params(":uuid"))
	}
	url := ""
	if attach.Type == models.StorageTypeObs {

		url, err = storage.ObsGetPreSignedUrl(attach.RealPath(), attach.Name)
		if err != nil {
			ctx.ServerError("ObsGetPreSignedUrl", err)
			return
		}
		if dataSet != nil {
			http.Redirect(ctx.Resp, ctx.Req.Request, url, http.StatusMovedPermanently)
			return
		}

	}

	//If we have matched and access to release or issue
	if setting.Attachment.StoreType == storage.MinioStorageType {

		if attach.Type == models.StorageTypeMinio {
			url, err = storage.Attachments.PresignedGetURL(attach.RealPath(), attach.Name)
			if err != nil {
				ctx.ServerError("PresignedGetURL", err)
				return
			}
		}

		if err = increaseDownloadCount(attach, dataSet); err != nil {
			ctx.ServerError("Update", err)
			return
		}
		if dataSet != nil {
			http.Redirect(ctx.Resp, ctx.Req.Request, url, http.StatusMovedPermanently)
		} else {
			fr, err := storage.Attachments.Open(attach.RelativePath())
			if err != nil {
				ctx.ServerError("Open", err)
				return
			}
			defer fr.Close()
			if err = ServeData(ctx, attach.Name, fr); err != nil {
				ctx.ServerError("ServeData", err)
				return
			}
		}

	} else {
		fr, err := storage.Attachments.Open(attach.RelativePath())
		if err != nil {
			ctx.ServerError("Open", err)
			return
		}
		defer fr.Close()
		if err = increaseDownloadCount(attach, dataSet); err != nil {
			ctx.ServerError("Update", err)
			return
		}
		if err = ServeData(ctx, attach.Name, fr); err != nil {
			ctx.ServerError("ServeData", err)
			return
		}
	}

}

func increaseDownloadCount(attach *models.Attachment, dataSet *models.Dataset) error {
	if err := attach.IncreaseDownloadCount(); err != nil {
		return err
	}

	if dataSet != nil {
		if err := models.IncreaseDownloadCount(dataSet.ID); err != nil {
			return err
		}
	}

	return nil
}

// Get a presigned url for put object
func GetPresignedPutObjectURL(ctx *context.Context) {
	if !setting.Attachment.Enabled {
		ctx.Error(404, "attachment is not enabled")
		return
	}

	err := upload.VerifyFileType(ctx.Params("file_type"), strings.Split(setting.Attachment.AllowedTypes, ","))
	if err != nil {
		ctx.Error(400, err.Error())
		return
	}

	if setting.Attachment.StoreType == storage.MinioStorageType {
		uuid := gouuid.NewV4().String()
		url, err := storage.Attachments.PresignedPutURL(models.AttachmentRelativePath(uuid))
		if err != nil {
			ctx.ServerError("PresignedPutURL", err)
			return
		}

		ctx.JSON(200, map[string]string{
			"uuid": uuid,
			"url":  url,
		})
	} else {
		ctx.Error(404, "storage type is not enabled")
		return
	}
}

// AddAttachment response for add attachment record
func AddAttachment(ctx *context.Context) {
	typeCloudBrain := ctx.QueryInt("type")
	fileName := ctx.Query("file_name")
	err := checkTypeCloudBrain(typeCloudBrain)
	if err != nil {
		ctx.ServerError("checkTypeCloudBrain failed", err)
		return
	}

	uuid := ctx.Query("uuid")
	has := false
	if typeCloudBrain == models.TypeCloudBrainOne {
		has, err = storage.Attachments.HasObject(setting.Attachment.Minio.BasePath + models.AttachmentRelativePath(uuid))
		if err != nil {
			ctx.ServerError("HasObject", err)
			return
		}
	} else {
		has, err = storage.ObsHasObject(setting.BasePath + models.AttachmentRelativePath(uuid) + "/" + fileName)
		if err != nil {
			ctx.ServerError("ObsHasObject", err)
			return
		}
	}

	if !has {
		ctx.Error(404, "attachment has not been uploaded")
		return
	}
	datasetId := ctx.QueryInt64("dataset_id")
	dataset, err := models.GetDatasetByID(datasetId)
	if err != nil {
		ctx.Error(404, "dataset does not exist.")
		return
	}

	attachment, err := models.InsertAttachment(&models.Attachment{
		UUID:       uuid,
		UploaderID: ctx.User.ID,
		IsPrivate:  dataset.IsPrivate(),
		Name:       fileName,
		Size:       ctx.QueryInt64("size"),
		DatasetID:  ctx.QueryInt64("dataset_id"),
		Type:       typeCloudBrain,
	})

	if err != nil {
		ctx.Error(500, fmt.Sprintf("InsertAttachment: %v", err))
		return
	}

	if attachment.DatasetID != 0 {
		if isCanDecompress(attachment.Name) {
			if typeCloudBrain == models.TypeCloudBrainOne {
				err = worker.SendDecompressTask(contexExt.Background(), uuid, attachment.Name)
				if err != nil {
					log.Error("SendDecompressTask(%s) failed:%s", uuid, err.Error())
				} else {
					attachment.DecompressState = models.DecompressStateIng
					err = models.UpdateAttachment(attachment)
					if err != nil {
						log.Error("UpdateAttachment state(%s) failed:%s", uuid, err.Error())
					}
				}
			}
			//todo:decompress type_two
		}
	}

	ctx.JSON(200, map[string]string{
		"result_code": "0",
	})
}

func isCanDecompress(name string) bool {
	if strings.HasSuffix(name, ".zip") || strings.HasSuffix(name, ".tar.gz") || strings.HasSuffix(name, ".tgz") {
		return true
	}
	return false
}

func UpdateAttachmentDecompressState(ctx *context.Context) {
	uuid := ctx.Query("uuid")
	result := ctx.Query("result")
	attach, err := models.GetAttachmentByUUID(uuid)
	if err != nil {
		log.Error("GetAttachmentByUUID(%s) failed:%s", uuid, err.Error())
		return
	}

	if result == DecompressSuccess {
		attach.DecompressState = models.DecompressStateDone
	} else if result == DecompressFailed {
		attach.DecompressState = models.DecompressStateFailed
	} else {
		log.Error("result is error:", result)
		return
	}

	err = models.UpdateAttachment(attach)
	if err != nil {
		log.Error("UpdateAttachment(%s) failed:%s", uuid, err.Error())
		return
	}
	log.Info("start to send msg to labelsystem ")

	dataset, _ := models.GetDatasetByID(attach.DatasetID)

	var labelMap map[string]string
	labelMap = make(map[string]string)
	labelMap["UUID"] = uuid
	labelMap["Type"] = fmt.Sprint(attach.Type)
	labelMap["UploaderID"] = fmt.Sprint(attach.UploaderID)
	labelMap["RepoID"] = fmt.Sprint(dataset.RepoID)
	labelMap["AttachName"] = attach.Name
	attachjson, _ := json.Marshal(labelMap)
	labelmsg.SendAddAttachToLabelSys(string(attachjson))

	log.Info("end to send msg to labelsystem ")

	ctx.JSON(200, map[string]string{
		"result_code": "0",
	})
}

func dealGlobalUniqueness(fileMD5 string, size int64, datasetId int64, storageType int, user *models.User, fileName string) map[string]string {
	fileChunks, err := models.GetFileChunkByMD5AndType(fileMD5, storageType, size)
	if err != nil {
		return nil
	} else {
		if fileChunks == nil || len(fileChunks) == 0 {
			return nil
		}
		for _, chunk := range fileChunks {
			if chunk.UserID == user.ID {
				log.Info("has find equal user id, so return nil.")
				return nil
			}
		}
		fileChunk := fileChunks[0]
		if fileChunk.IsUploaded == models.FileUploaded {
			attach, err := models.GetAttachmentByUUID(fileChunk.UUID)
			if err == nil {
				uuid := gouuid.NewV4().String()
				_, err = models.InsertFileChunk(&models.FileChunk{
					UUID:        uuid,
					UserID:      user.ID,
					UploadID:    fileChunk.UploadID,
					IsUploaded:  fileChunk.IsUploaded,
					Md5:         fileMD5,
					Size:        size,
					TotalChunks: fileChunk.TotalChunks,
					Type:        storageType,
				})
				if err != nil {
					log.Info("insert filechunk error." + err.Error())
				}
				dataset, _ := models.GetDatasetByID(datasetId)
				log.Warn("insert attachment to datasetId:" + strconv.FormatInt(dataset.ID, 10))
				state := attach.DecompressState
				if attach.DecompressState == 3 {
					log.Info("attach.UnzipPath=" + attach.UnzipPath)
					files, err := storage.GetOneLevelObjectsUnderDir(setting.Bucket, attach.UnzipPath, "")
					if err == nil {
						if len(files) > 0 {
							state = 1
						}
					}
				}
				attachment, err := models.InsertAttachment(&models.Attachment{
					UUID:            uuid,
					UploaderID:      user.ID,
					IsPrivate:       dataset.IsPrivate(),
					Name:            fileName,
					Size:            size,
					DatasetID:       datasetId,
					Description:     "",
					Type:            storageType,
					Path:            attach.Path,
					UnzipPath:       attach.UnzipPath,
					DecompressState: state,
				})
				if err != nil {
					log.Info("insert to attachement error", err)
					return nil
				}
				attachment.UpdateDatasetUpdateUnix()
				go repo_service.IncreaseRepoDatasetNum(dataset.ID)

				repository, _ := models.GetRepositoryByID(dataset.RepoID)
				notification.NotifyOtherTask(user, repository, fmt.Sprint(repository.IsPrivate, attachment.IsPrivate), attachment.Name, models.ActionUploadAttachment)

				return map[string]string{
					"uuid":        fileChunk.UUID,
					"uploaded":    strconv.Itoa(fileChunk.IsUploaded),
					"uploadID":    fileChunk.UploadID,
					"chunks":      strconv.Itoa(0),
					"attachID":    strconv.Itoa(int(attachment.ID)),
					"datasetID":   strconv.Itoa(int(datasetId)),
					"fileName":    "",
					"datasetName": "",
				}
			} else {
				log.Info("find sql is err:" + err.Error())
			}
		}
	}
	return nil
}

func GetSuccessChunks(ctx *context.Context) {
	fileMD5 := ctx.Query("md5")
	fileName := ctx.Query("file_name")
	size := ctx.QueryInt64("size")
	datasetId := ctx.QueryInt64("dataset_id")
	log.Info("size=" + fmt.Sprint(size))
	log.Info("datasetId=" + fmt.Sprint(datasetId))
	var chunks string
	if strings.Contains(fileName, " ") {
		ctx.JSON(http.StatusBadRequest, ctx.Tr("dataset.file_name_invalid"))
		return
	}

	if storage_limit.IsFileUploadOverLimit(ctx.User.ID, size) {
		ctx.JSON(200, map[string]string{
			"result_code": "-1",
			"msg":         ctx.Tr("common_error.file_size_storage_limit"),
		})
		return
	}
	storageType := models.GetDefaultStorageType()

	re := dealGlobalUniqueness(fileMD5, size, datasetId, storageType, ctx.User, fileName)
	if re != nil {
		ctx.JSON(200, re)
		return
	}

	fileChunk, err := models.GetFileChunkByMD5AndUser(fileMD5, ctx.User.ID, storageType)
	if err != nil {
		if models.IsErrFileChunkNotExist(err) {
			ctx.JSON(200, map[string]string{
				"uuid":     "",
				"uploaded": "0",
				"uploadID": "",
				"chunks":   "",
			})
		} else {
			ctx.ServerError("GetFileChunkByMD5", err)
		}
		return
	}

	isExist := false
	if storageType == models.StorageTypeMinio {
		oldAttachment, err := models.GetAttachmentByUUID(fileChunk.UUID)
		if err == nil {
			log.Info("start to check oldAttachment, path=" + oldAttachment.Path)
			isExist, err = storage.Attachments.HasObject(oldAttachment.Path)
			if err != nil {
				ctx.ServerError("ObsHasObject failed", err)
				return
			}
		} else {
			isExist, err = storage.Attachments.HasObject(setting.Attachment.Minio.BasePath + models.AttachmentRelativePath(fileChunk.UUID))
			if err != nil {
				ctx.ServerError("HasObject failed", err)
				return
			}
		}

	} else {
		oldFileName := fileName
		oldAttachment, err := models.GetAttachmentByUUID(fileChunk.UUID)
		if err == nil {
			log.Info("start to check oldAttachment, path=" + oldAttachment.Path)
			isExist, err = storage.ObsHasObject(oldAttachment.Path)
			if err != nil {
				ctx.ServerError("ObsHasObject failed", err)
				return
			}
		} else {
			isExist, err = storage.ObsHasObject(setting.BasePath + models.AttachmentRelativePath(fileChunk.UUID) + "/" + oldFileName)
			if err != nil {
				ctx.ServerError("ObsHasObject failed", err)
				return
			}
		}
	}

	if isExist {
		if fileChunk.IsUploaded == models.FileNotUploaded {
			log.Info("the file has been uploaded but not recorded")
			fileChunk.IsUploaded = models.FileUploaded
			if err = models.UpdateFileChunk(fileChunk); err != nil {
				log.Error("UpdateFileChunk failed:", err.Error())
			}
		}
	} else {
		if fileChunk.IsUploaded == models.FileUploaded {
			log.Info("the file has been recorded but not uploaded")
			fileChunk.IsUploaded = models.FileNotUploaded
			if err = models.UpdateFileChunk(fileChunk); err != nil {
				log.Error("UpdateFileChunk failed:", err.Error())
			}
		}

		if storageType == models.StorageTypeMinio {
			chunks, err = storage.GetPartInfos(strings.TrimPrefix(path.Join(setting.Attachment.Minio.BasePath, path.Join(fileChunk.UUID[0:1], fileChunk.UUID[1:2], fileChunk.UUID)), "/"), fileChunk.UploadID)
			if err != nil {
				log.Error("GetPartInfos failed:%v", err.Error())
			}
		} else {
			chunks, err = storage.GetObsPartInfos(strings.TrimPrefix(path.Join(setting.BasePath, path.Join(fileChunk.UUID[0:1], fileChunk.UUID[1:2], fileChunk.UUID, fileName)), "/"), fileChunk.UploadID)
			if err != nil {
				log.Error("GetObsPartInfos failed:%v", err.Error())
			}
		}

		if err != nil {
			models.DeleteFileChunk(fileChunk)
			ctx.JSON(200, map[string]string{
				"uuid":     "",
				"uploaded": "0",
				"uploadID": "",
				"chunks":   "",
			})
			return
		}
	}

	var attachID int64
	attach, err := models.GetAttachmentByUUID(fileChunk.UUID)
	if err != nil {
		if models.IsErrAttachmentNotExist(err) {
			attachID = 0
		} else {
			ctx.ServerError("GetAttachmentByUUID", err)
			return
		}
	} else {
		attachID = attach.ID
	}

	if attach == nil {
		ctx.JSON(200, map[string]string{
			"uuid":        fileChunk.UUID,
			"uploaded":    strconv.Itoa(fileChunk.IsUploaded),
			"uploadID":    fileChunk.UploadID,
			"chunks":      string(chunks),
			"attachID":    "0",
			"datasetID":   "0",
			"fileName":    "",
			"datasetName": "",
		})
		return
	}

	dataset, err := models.GetDatasetByID(attach.DatasetID)
	if err != nil {
		ctx.ServerError("GetDatasetByID", err)
		return
	}

	repo, err := models.GetRepositoryByID(dataset.RepoID)
	if err != nil {
		ctx.ServerError("GetRepositoryByID", err)
		return
	}

	ctx.JSON(200, map[string]string{
		"uuid":        fileChunk.UUID,
		"uploaded":    strconv.Itoa(fileChunk.IsUploaded),
		"uploadID":    fileChunk.UploadID,
		"chunks":      string(chunks),
		"attachID":    strconv.Itoa(int(attachID)),
		"datasetID":   strconv.Itoa(int(attach.DatasetID)),
		"fileName":    attach.Name,
		"datasetName": dataset.Title,
		"repoName":    repo.Name,
		"repoOwner":   repo.OwnerName,
	})

}

func NewMultipartForApi(ctx *context.Context, isFlowControl bool) (map[string]string, error) {
	if !setting.Attachment.Enabled {
		return nil, errors.New("attachment is not enabled")
	}
	storageType := models.GetDefaultStorageType()
	fileMD5 := ctx.Query("md5")
	fileChunk, err := models.GetFileChunkByMD5AndUser(fileMD5, ctx.User.ID, storageType)
	if err == nil {
		if fileChunk != nil {
			log.Info("cannot reupload,name" + ctx.Query("file_name"))
			return nil, errors.New("Cannot upload repeatedly,name is " + ctx.Query("file_name"))
		}
	}
	if isFlowControl {
		err = CheckFlowForDataset(ctx)
		if err != nil {
			log.Info("check error," + err.Error())
			return nil, err
		}
	}
	err = upload.VerifyFileType(ctx.Query("fileType"), strings.Split(setting.Attachment.AllowedTypes, ","))
	if err != nil {
		log.Info("VerifyFileType error," + err.Error())
		return nil, errors.New("Not support file type.")
	}

	fileName := ctx.Query("file_name")

	totalChunkCounts := ctx.QueryInt("totalChunkCounts")
	if totalChunkCounts > minio_ext.MaxPartsCount {
		log.Info(fmt.Sprintf("chunk counts(%d) is too much", totalChunkCounts))
		return nil, errors.New(fmt.Sprintf("chunk counts(%d) is too much", totalChunkCounts))

	}

	fileSize := ctx.QueryInt64("size")
	if storage_limit.IsFileUploadOverLimit(ctx.User.ID, fileSize) {
		return nil, errors.New(ctx.Tr("common_error.file_size_storage_limit"))

	}
	if fileSize > minio_ext.MaxMultipartPutObjectSize {
		log.Info(fmt.Sprintf("file size(%d) is too big", fileSize))
		return nil, errors.New(fmt.Sprintf("file size(%d) is too big", fileSize))
	}

	uuid := gouuid.NewV4().String()
	var uploadID string
	if storageType == models.StorageTypeMinio {
		uploadID, err = storage.NewMultiPartUpload(strings.TrimPrefix(path.Join(setting.Attachment.Minio.BasePath, path.Join(uuid[0:1], uuid[1:2], uuid)), "/"))
		if err != nil {
			log.Info("NewMultipart " + err.Error())
			return nil, err
		}
	} else {
		uploadID, err = storage.NewObsMultiPartUpload(strings.TrimPrefix(path.Join(setting.BasePath, path.Join(uuid[0:1], uuid[1:2], uuid, fileName)), "/"))
		if err != nil {
			log.Info("NewObsMultiPartUpload " + err.Error())
			return nil, err
		}
	}
	_, err = models.InsertFileChunk(&models.FileChunk{
		UUID:        uuid,
		UserID:      ctx.User.ID,
		UploadID:    uploadID,
		Md5:         ctx.Query("md5"),
		Size:        fileSize,
		TotalChunks: totalChunkCounts,
		Type:        storageType,
	})

	if err != nil {
		log.Info(fmt.Sprintf("InsertFileChunk: %v", err))
		return nil, err
	}
	return map[string]string{
		"uuid":     uuid,
		"uploadID": uploadID,
	}, nil

}

func NewMultipart(ctx *context.Context) {
	re, err := NewMultipartForApi(ctx, false)
	if err != nil {
		ctx.JSON(200, map[string]string{
			"result_code": "-1",
			"msg":         err.Error(),
		})
		//ctx.ServerError("NewMultipart failed", err)
		return
	}
	re["result_code"] = "0"
	ctx.JSON(200, re)
}

func PutOBSProxyUpload(ctx *context.Context) {
	uuid := ctx.Query("uuid")
	uploadID := ctx.Query("uploadId")
	partNumber := ctx.QueryInt("partNumber")
	fileName := ctx.Query("file_name")

	RequestBody := ctx.Req.Body()

	if RequestBody == nil {
		ctx.Error(500, fmt.Sprintf("FormFile: %v", RequestBody))
		return
	}
	objectName := strings.TrimPrefix(path.Join(setting.BasePath, path.Join(uuid[0:1], uuid[1:2], uuid, fileName)), "/")
	err := storage.ObsMultiPartUpload(objectName, uploadID, partNumber, fileName, RequestBody.ReadCloser())
	if err != nil {
		log.Info("upload error.")
	}
}

func GetOBSProxyDownload(ctx *context.Context) {
	uuid := ctx.Query("uuid")
	fileName := ctx.Query("file_name")
	objectName := strings.TrimPrefix(path.Join(setting.BasePath, path.Join(uuid[0:1], uuid[1:2], uuid, fileName)), "/")
	body, err := storage.ObsDownloadAFile(setting.Bucket, objectName)
	if err != nil {
		log.Info("upload error.")
	} else {
		defer body.Close()
		ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+fileName)
		ctx.Resp.Header().Set("Content-Type", "application/octet-stream")
		p := make([]byte, 1024)
		var readErr error
		var readCount int
		// 读取对象内容
		for {
			readCount, readErr = body.Read(p)
			if readCount > 0 {
				ctx.Resp.Write(p[:readCount])
				//fmt.Printf("%s", p[:readCount])
			}
			if readErr != nil {
				break
			}
		}
	}
}

func GetMultipartUploadUrl(ctx *context.Context) {
	uuid := ctx.Query("uuid")
	uploadID := ctx.Query("uploadID")
	partNumber := ctx.QueryInt("chunkNumber")
	size := ctx.QueryInt64("size")
	fileName := ctx.Query("file_name")

	storageType := models.GetDefaultStorageType()

	url := ""
	var err error
	if storageType == models.StorageTypeMinio {
		if size > minio_ext.MinPartSize {
			ctx.Error(400, fmt.Sprintf("chunk size(%d) is too big", size))
			return
		}

		url, err = storage.GenMultiPartSignedUrl(strings.TrimPrefix(path.Join(setting.Attachment.Minio.BasePath, path.Join(uuid[0:1], uuid[1:2], uuid)), "/"), uploadID, partNumber, size)
		if err != nil {
			ctx.Error(500, fmt.Sprintf("GenMultiPartSignedUrl failed: %v", err))
			return
		}
	} else {
		if setting.PROXYURL != "" {
			url = setting.PROXYURL + "/obs_proxy_multipart?uuid=" + uuid + "&uploadId=" + uploadID + "&partNumber=" + fmt.Sprint(partNumber) + "&file_name=" + fileName
			log.Info("return url=" + url)
		} else {
			url, err = storage.ObsGenMultiPartSignedUrl(strings.TrimPrefix(path.Join(setting.BasePath, path.Join(uuid[0:1], uuid[1:2], uuid, fileName)), "/"), uploadID, partNumber)
			if err != nil {
				ctx.Error(500, fmt.Sprintf("ObsGenMultiPartSignedUrl failed: %v", err))
				return
			}
			log.Info("url=" + url)
		}
	}
	ctx.JSON(200, map[string]string{
		"url": url,
	})
}

func CompleteMultipart(ctx *context.Context) {
	uuid := ctx.Query("uuid")
	uploadID := ctx.Query("uploadID")
	storageType := models.GetDefaultStorageType()
	fileName := ctx.Query("file_name")

	log.Warn("uuid:" + uuid)

	fileChunk, err := models.GetFileChunkByUUID(uuid)
	if err != nil {
		ctx.JSON(200, map[string]string{
			"result_code": "-1",
			"msg":         "The upload file not found.",
		})
		return
	}

	if storageType == models.StorageTypeMinio {
		_, err = storage.CompleteMultiPartUpload(strings.TrimPrefix(path.Join(setting.Attachment.Minio.BasePath, path.Join(fileChunk.UUID[0:1], fileChunk.UUID[1:2], fileChunk.UUID)), "/"), uploadID, fileChunk.TotalChunks)
		if err != nil {
			ctx.JSON(200, map[string]string{
				"result_code": "-1",
				"msg":         fmt.Sprintf("CompleteMultiPartUpload failed: %v", err),
			})
			//ctx.Error(500, fmt.Sprintf("CompleteMultiPartUpload failed: %v", err))
			return
		}
	} else {
		err = storage.CompleteObsMultiPartUpload(strings.TrimPrefix(path.Join(setting.BasePath, path.Join(fileChunk.UUID[0:1], fileChunk.UUID[1:2], fileChunk.UUID, fileName)), "/"), uploadID, fileChunk.TotalChunks)
		if err != nil {
			ctx.JSON(200, map[string]string{
				"result_code": "-1",
				"msg":         fmt.Sprintf("CompleteObsMultiPartUpload failed: %v", err),
			})
			//ctx.Error(500, fmt.Sprintf("CompleteObsMultiPartUpload failed: %v", err))
			return
		}
	}

	fileChunk.IsUploaded = models.FileUploaded

	err = models.UpdateFileChunk(fileChunk)
	if err != nil {
		ctx.JSON(200, map[string]string{
			"result_code": "-1",
			"msg":         fmt.Sprintf("UpdateFileChunk: %v", err),
		})
		//ctx.Error(500, fmt.Sprintf("UpdateFileChunk: %v", err))
		return
	}

	err = finishedUploadAttachment(ctx.QueryInt64("dataset_id"), ctx.QueryInt64("size"), uuid, fileName, ctx.Query("description"), storageType, ctx.User)
	if err == nil {
		ctx.JSON(200, map[string]string{
			"result_code": "0",
		})
	} else {
		ctx.JSON(200, map[string]string{
			"result_code": "-1",
			"msg":         err.Error(),
		})
	}

}

func finishedUploadAttachment(datasetId, size int64, uuid, fileName, description string, storageType int, User *models.User) error {
	dataset, _ := models.GetDatasetByID(datasetId)
	log.Warn("insert attachment to datasetId:" + strconv.FormatInt(dataset.ID, 10))
	objectKey := "attachment/" + uuid[0:1] + "/" + uuid[1:2] + "/" + uuid + "/" + fileName
	unzipPath := "attachment/" + uuid[0:1] + "/" + uuid[1:2] + "/" + uuid + uuid + "/"
	if storageType == models.StorageTypeMinio {
		objectKey = "attachments/" + uuid[0:1] + "/" + uuid[1:2] + "/" + uuid
		unzipPath = "attachments/" + uuid[0:1] + "/" + uuid[1:2] + "/" + uuid + uuid + "/"
	}
	attachment, err := models.InsertAttachment(&models.Attachment{
		UUID:        uuid,
		UploaderID:  User.ID,
		IsPrivate:   dataset.IsPrivate(),
		Name:        fileName,
		Size:        size,
		DatasetID:   datasetId,
		Description: description,
		Type:        storageType,
		Path:        objectKey,
		UnzipPath:   unzipPath,
	})
	if err != nil {
		return errors.New("InsertAttachment: " + err.Error())
	}
	attachment.UpdateDatasetUpdateUnix()
	go repo_service.IncreaseRepoDatasetNum(dataset.ID)
	if attachment.DatasetID != 0 {
		if isCanDecompress(attachment.Name) {
			if storageType == models.StorageTypeMinio {
				err = worker.SendDecompressTask(contexExt.Background(), uuid, attachment.Name)
				if err != nil {
					log.Error("SendDecompressTask(%s) failed:%s", uuid, err.Error())
				} else {
					updateAttachmentDecompressStateIng(attachment)
				}
			}
			if storageType == models.StorageTypeObs {
				log.Info("start to send msg to unzip =" + attachment.Name)
				attachjson, _ := json.Marshal(attachment)
				err = labelmsg.SendDecompressAttachToLabelOBS(string(attachjson))
				if err != nil {
					log.Error("SendDecompressTask to labelsystem (%s) failed:%s", attachment.UUID, err.Error())
				} else {
					updateAttachmentDecompressStateIng(attachment)
				}
			}
		} else {
			var labelMap map[string]string
			labelMap = make(map[string]string)
			labelMap["UUID"] = uuid
			labelMap["Type"] = fmt.Sprint(attachment.Type)
			labelMap["UploaderID"] = fmt.Sprint(attachment.UploaderID)
			labelMap["RepoID"] = fmt.Sprint(dataset.RepoID)
			labelMap["AttachName"] = attachment.Name
			attachjson, _ := json.Marshal(labelMap)
			labelmsg.SendAddAttachToLabelSys(string(attachjson))
		}
	}
	repository, _ := models.GetRepositoryByID(dataset.RepoID)
	notification.NotifyOtherTask(User, repository, fmt.Sprint(repository.IsPrivate, attachment.IsPrivate), attachment.Name, models.ActionUploadAttachment)
	return nil
}

func HandleUnDecompressAttachment() {
	attachs, err := models.GetUnDecompressAttachments()
	if err != nil {
		log.Error("GetUnDecompressAttachments failed:", err.Error())
		return
	}

	for _, attach := range attachs {
		if attach.Type == models.TypeCloudBrainOne {
			err = worker.SendDecompressTask(contexExt.Background(), attach.UUID, attach.Name)
			if err != nil {
				log.Error("SendDecompressTask(%s)  failed:%s", attach.UUID, err.Error())
			} else {
				updateAttachmentDecompressStateIng(attach)
			}
		} else if attach.Type == models.TypeCloudBrainTwo {
			attachjson, _ := json.Marshal(attach)
			err = labelmsg.SendDecompressAttachToLabelOBS(string(attachjson))
			if err != nil {
				log.Error("SendDecompressTask to labelsystem (%s) failed:%s", attach.UUID, err.Error())
			} else {
				updateAttachmentDecompressStateIng(attach)
			}
		}
	}
	return
}
func updateAttachmentDecompressStateIng(attach *models.Attachment) {
	attach.DecompressState = models.DecompressStateIng
	err := models.UpdateAttachment(attach)
	if err != nil {
		log.Error("UpdateAttachment state(%s) failed:%s", attach.UUID, err.Error())
	}
}

func QueryAllPublicDataset(ctx *context.Context) {
	attachs, err := models.GetAllPublicAttachments()
	if err != nil {
		ctx.JSON(200, map[string]string{
			"result_code": "-1",
			"error_msg":   err.Error(),
			"data":        "",
		})
		return
	}

	queryDatasets(ctx, attachs)
}

func QueryPrivateDataset(ctx *context.Context) {
	username := ctx.Params(":username")
	attachs, err := models.GetPrivateAttachments(username)
	if err != nil {
		ctx.JSON(200, map[string]string{
			"result_code": "-1",
			"error_msg":   err.Error(),
			"data":        "",
		})
		return
	}

	for _, attach := range attachs {
		attach.Name = username
	}

	queryDatasets(ctx, attachs)
}

func queryDatasets(ctx *context.Context, attachs []*models.AttachmentUsername) {
	var datasets []CloudBrainDataset
	if len(attachs) == 0 {
		log.Info("dataset is null")
		ctx.JSON(200, map[string]string{
			"result_code": "0",
			"error_msg":   "",
			"data":        "",
		})
		return
	}

	for _, attch := range attachs {
		has, err := storage.Attachments.HasObject(setting.Attachment.Minio.BasePath + models.AttachmentRelativePath(attch.UUID))
		if err != nil || !has {
			continue
		}

		datasets = append(datasets, CloudBrainDataset{strconv.FormatInt(attch.ID, 10),
			attch.Attachment.Name,
			setting.Attachment.Minio.RealPath +
				setting.Attachment.Minio.Bucket + "/" +
				setting.Attachment.Minio.BasePath +
				models.AttachmentRelativePath(attch.UUID) +
				attch.UUID,
			attch.Name,
			attch.CreatedUnix.Format("2006-01-02 03:04:05 PM")})
	}

	data, err := json.Marshal(datasets)
	if err != nil {
		log.Error("json.Marshal failed:", err.Error())
		ctx.JSON(200, map[string]string{
			"result_code": "-1",
			"error_msg":   err.Error(),
			"data":        "",
		})
		return
	}

	ctx.JSON(200, map[string]string{
		"result_code": "0",
		"error_msg":   "",
		"data":        string(data),
	})
	return
}

func checkTypeCloudBrain(typeCloudBrain int) error {
	if typeCloudBrain != models.TypeCloudBrainOne && typeCloudBrain != models.TypeCloudBrainTwo {
		log.Error("type error:", typeCloudBrain)
		return errors.New("type error")
	}
	return nil
}
