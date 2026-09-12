package repo

import (
	"errors"
	"fmt"
	"path"
	"strconv"
	"strings"

	"code.gitea.io/gitea/services/ai_model"
	"code.gitea.io/gitea/services/storage_limit"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/minio_ext"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
	"code.gitea.io/gitea/modules/upload"
	gouuid "github.com/satori/go.uuid"
)

func GetModelChunks(ctx *context.Context) {
	fileMD5 := ctx.Query("md5")
	storageType := models.GetDefaultStorageType()
	fileName := ctx.Query("file_name")
	modeluuid := ctx.Query("modeluuid")

	var chunks string

	var fileChunk *models.ModelFileChunk
	var err error
	if fileName != "" {
		fileChunk, err = models.GetModelFileChunkByFileNameAndUser(getObjectName(fileName, modeluuid), ctx.User.ID, storageType, modeluuid)
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
		} else {
			if fileChunk.Md5 != fileMD5 && fileChunk.IsUploaded == models.FileNotUploaded {
				//同名，且上一个没有传完，直接覆盖。否则走后面提示
				models.DeleteModelFileChunk(fileChunk)
				ctx.JSON(200, map[string]string{
					"uuid":     "",
					"uploaded": "0",
					"uploadID": "",
					"chunks":   "",
				})
				return
			}
		}
	} else {
		fileChunk, err = models.GetModelFileChunkByMD5AndUser(fileMD5, ctx.User.ID, storageType, modeluuid)
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
	}

	isExist := false
	if storageType == models.StorageTypeMinio {
		isExist, err = storage.Attachments.HasObject(fileChunk.ObjectName)
		if isExist {
			log.Info("The file is exist in minio. has uploaded.path=" + fileChunk.ObjectName)
		} else {
			log.Info("The file is not exist in minio..")
		}
		if err != nil {
			ctx.ServerError("HasObject failed", err)
			return
		}
	} else {
		isExist, err = storage.ObsHasObject(fileChunk.ObjectName)
		if isExist {
			log.Info("The file is exist in obs. has uploaded. path=" + fileChunk.ObjectName)
		} else {
			log.Info("The file is not exist in obs.")
		}
		if err != nil {
			ctx.ServerError("ObsHasObject failed", err)
			return
		}
	}

	if isExist {
		if fileChunk.IsUploaded == models.FileNotUploaded {
			log.Info("the file has been uploaded but not recorded")
			fileChunk.IsUploaded = models.FileUploaded
			if err = models.UpdateModelFileChunk(fileChunk); err != nil {
				log.Error("UpdateFileChunk failed:", err.Error())
			}
		}
		modelname := ""
		model, err := models.QueryModelById(modeluuid)
		if err == nil && model != nil {
			modelname = model.Name
		}
		fileNameUploaded := strings.Split(fileChunk.ObjectName, modeluuid+"/")[1]
		ctx.JSON(200, map[string]string{
			"uuid":      fileChunk.UUID,
			"uploaded":  strconv.Itoa(fileChunk.IsUploaded),
			"uploadID":  fileChunk.UploadID,
			"chunks":    string(chunks),
			"attachID":  "0",
			"modeluuid": modeluuid,
			"fileName":  fileNameUploaded,
			"modelName": modelname,
		})
	} else {
		if fileChunk.IsUploaded == models.FileUploaded {
			log.Info("the file has been recorded but not uploaded")
			fileChunk.IsUploaded = models.FileNotUploaded
			if err = models.UpdateModelFileChunk(fileChunk); err != nil {
				log.Error("UpdateFileChunk failed:", err.Error())
			}
		}

		if storageType == models.StorageTypeMinio {
			chunks, err = storage.GetPartInfos(fileChunk.ObjectName, fileChunk.UploadID)
			if err != nil {
				log.Error("GetPartInfos failed:%v", err.Error())
			}
		} else {
			chunks, err = storage.GetObsPartInfos(fileChunk.ObjectName, fileChunk.UploadID)
			if err != nil {
				log.Error("GetObsPartInfos failed:%v", err.Error())
			}
		}
		if err != nil {
			models.DeleteModelFileChunk(fileChunk)
			ctx.JSON(200, map[string]string{
				"uuid":     "",
				"uploaded": "0",
				"uploadID": "",
				"chunks":   "",
			})
		} else {
			ctx.JSON(200, map[string]string{
				"uuid":      fileChunk.UUID,
				"uploaded":  strconv.Itoa(fileChunk.IsUploaded),
				"uploadID":  fileChunk.UploadID,
				"chunks":    string(chunks),
				"attachID":  "0",
				"datasetID": "0",
				"fileName":  "",
				"modelName": "",
			})
		}
	}
}

func getObjectName(filename string, modeluuid string) string {
	return strings.TrimPrefix(path.Join(Model_prefix, path.Join(modeluuid[0:1], modeluuid[1:2], modeluuid, filename)), "/")
}

func NewModelMultipart(ctx *context.Context) {
	re, err := NewModelMultipartForApi(ctx, false)
	if err != nil {
		ctx.JSON(200, map[string]string{
			"result_code": "-1",
			"msg":         err.Error(),
		})

		return
	}
	re["result_code"] = "0"
	ctx.JSON(200, re)
}

func NewModelMultipartForApi(ctx *context.Context, isFlowControl bool) (map[string]string, error) {
	if !setting.Attachment.Enabled {
		return nil, errors.New("attachment is not enabled")
	}
	fileName := ctx.Query("file_name")
	modeluuid := ctx.Query("modeluuid")
	err := upload.VerifyFileType(ctx.Query("fileType"), strings.Split(setting.Attachment.AllowedTypes, ","))
	if err != nil {
		return nil, err
	}
	if isFlowControl {
		err = CheckFlowForModel(ctx)
		if err != nil {
			log.Info("check error," + err.Error())
			return nil, err
		}
	}
	storageType := models.GetDefaultStorageType()

	totalChunkCounts := ctx.QueryInt("totalChunkCounts")
	if totalChunkCounts > minio_ext.MaxPartsCount {
		return nil, errors.New(fmt.Sprintf("chunk counts(%d) is too much", totalChunkCounts))

	}

	fileSize := ctx.QueryInt64("size")
	if storage_limit.IsFileUploadOverLimit(ctx.User.ID, fileSize) {
		return nil, errors.New(ctx.Tr("common_error.file_size_storage_limit"))

	}
	if fileSize > minio_ext.MaxMultipartPutObjectSize {
		return nil, errors.New(fmt.Sprintf("file size(%d) is too big", fileSize))
	}

	uuid := gouuid.NewV4().String()
	var uploadID string
	var objectName string
	if storageType == models.StorageTypeMinio {
		objectName = strings.TrimPrefix(path.Join(Model_prefix, path.Join(modeluuid[0:1], modeluuid[1:2], modeluuid, fileName)), "/")
		uploadID, err = storage.NewMultiPartUpload(objectName)
		if err != nil {
			return nil, err
		}
	} else {

		objectName = strings.TrimPrefix(path.Join(Model_prefix, path.Join(modeluuid[0:1], modeluuid[1:2], modeluuid, fileName)), "/")
		uploadID, err = storage.NewObsMultiPartUpload(objectName)
		if err != nil {
			return nil, err
		}
	}

	_, err = models.InsertModelFileChunk(&models.ModelFileChunk{
		UUID:        uuid,
		UserID:      ctx.User.ID,
		UploadID:    uploadID,
		Md5:         ctx.Query("md5"),
		Size:        fileSize,
		ObjectName:  objectName,
		ModelUUID:   modeluuid,
		TotalChunks: totalChunkCounts,
		Type:        storageType,
	})

	if err != nil {
		return nil, err

	}
	return map[string]string{
		"uuid":     uuid,
		"uploadID": uploadID,
	}, nil

}

func GetModelMultipartUploadUrl(ctx *context.Context) {
	uuid := ctx.Query("uuid")
	uploadID := ctx.Query("uploadID")
	partNumber := ctx.QueryInt("chunkNumber")
	size := ctx.QueryInt64("size")
	storageType := models.GetDefaultStorageType()

	fileChunk, err := models.GetModelFileChunkByUUID(uuid)
	if err != nil {
		if models.IsErrFileChunkNotExist(err) {
			ctx.Error(404)
		} else {
			ctx.ServerError("GetFileChunkByUUID", err)
		}
		return
	}
	url := ""
	if storageType == models.StorageTypeMinio {
		if size > minio_ext.MinPartSize {
			ctx.Error(400, fmt.Sprintf("chunk size(%d) is too big", size))
			return
		}
		url, err = storage.GenMultiPartSignedUrl(fileChunk.ObjectName, uploadID, partNumber, size)
		if err != nil {
			ctx.Error(500, fmt.Sprintf("GenMultiPartSignedUrl failed: %v", err))
			return
		}
	} else {
		url, err = storage.ObsGenMultiPartSignedUrl(fileChunk.ObjectName, uploadID, partNumber)
		if err != nil {
			ctx.Error(500, fmt.Sprintf("ObsGenMultiPartSignedUrl failed: %v", err))
			return
		}
		log.Info("url=" + url)

	}

	ctx.JSON(200, map[string]string{
		"url": url,
	})
}

func CompleteModelMultipart(ctx *context.Context) {
	uuid := ctx.Query("uuid")
	uploadID := ctx.Query("uploadID")
	storageType := models.GetDefaultStorageType()
	modeluuid := ctx.Query("modeluuid")
	log.Warn("uuid:" + uuid)
	log.Warn("modeluuid:" + modeluuid)

	fileChunk, err := models.GetModelFileChunkByUUID(uuid)
	if err != nil {
		if models.IsErrFileChunkNotExist(err) {
			ctx.Error(404)
		} else {
			ctx.ServerError("GetFileChunkByUUID", err)
		}
		return
	}

	if storageType == models.StorageTypeMinio {
		_, err = storage.CompleteMultiPartUpload(fileChunk.ObjectName, uploadID, fileChunk.TotalChunks)
		if err != nil {
			ctx.Error(500, fmt.Sprintf("CompleteMultiPartUpload failed: %v", err))
			return
		}
	} else {
		err = storage.CompleteObsMultiPartUpload(fileChunk.ObjectName, uploadID, fileChunk.TotalChunks)
		if err != nil {
			ctx.Error(500, fmt.Sprintf("CompleteObsMultiPartUpload failed: %v", err))
			return
		}
	}

	fileChunk.IsUploaded = models.FileUploaded

	err = models.UpdateModelFileChunk(fileChunk)
	if err != nil {
		ctx.Error(500, fmt.Sprintf("UpdateFileChunk: %v", err))
		return
	}
	//更新模型大小信息
	UpdateModelSize(modeluuid, fileChunk.ObjectName)
	ai_model.UpdateModelMeta(modeluuid)

	ctx.JSON(200, map[string]string{
		"result_code": "0",
	})

}
