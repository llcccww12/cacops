package subject_service

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/minio_ext"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/upload"
	"code.gitea.io/gitea/modules/util"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
	"code.gitea.io/gitea/services/storage_limit"
)

func GetFileChunks(req entity.UploadChunkRequest) (*entity.UploadChunkResponse, error) {
	fileMD5 := req.FileMD5
	fileName := req.FileName
	subjectId := req.SubjectID
	subjectType := req.SubjectType
	userId := req.UserId

	var chunks = make([]int, 0)
	var err error

	uploadHelper := GetUploadHelper(models.SubjectType(subjectType), subjectId)
	if uploadHelper == nil {
		log.Error("GetUploadHelper failed, dataType=%d", subjectType)
		return nil, errors.New("Data type error")
	}
	storageType := uploadHelper.GetStorageType()
	storageHelper := storage_helper.SelectStorageHelperFromStorageType(storageType)
	if storageHelper == nil {
		log.Error("storageHelper failed, datastorageTypeType=%d", storageType)
		return nil, errors.New("Storage type error")
	}

	var fileChunk *models.UploadChunk
	if fileName != "" {
		fileChunk, err = models.GetUploadChunkByFileNameAndUser(uploadHelper.GetFileStoragePath(fileName, subjectId), userId, string(storageType), subjectId)
		log.Info("GetUploadChunkByFileNameAndUser fileChunk=%+v err=%v", fileChunk, err)
		if err != nil {
			if models.IsErrFileChunkNotExist(err) {
				return &entity.UploadChunkResponse{}, nil
			}

			return nil, err
		}

		if fileChunk.Md5 != fileMD5 && fileChunk.IsUploaded == models.FileNotUploaded {
			//同名，且上一个没有传完，直接覆盖。否则走后面提示
			models.DeleteUploadChunk(fileChunk)
			return &entity.UploadChunkResponse{}, nil
		}

	} else {
		fileChunk, err = models.GetUploadChunkByMD5AndUser(fileMD5, userId, string(storageType), subjectId)
		if err != nil {
			if models.IsErrFileChunkNotExist(err) {
				return &entity.UploadChunkResponse{}, nil
			}
			return nil, err
		}
	}

	isExist, err := storageHelper.HasObject(fileChunk.ObjectName)
	if err != nil {
		log.Error("HasObject failed:objectName = %s %v", fileChunk.ObjectName, err.Error())
		return nil, err
	}

	if isExist {
		log.Info("The file is exist. has uploaded. path=%s", fileChunk.ObjectName)
		if fileChunk.IsUploaded == models.DataFileNotUploaded {
			log.Info("the file has been uploaded but not recorded")
			fileChunk.IsUploaded = models.DataFileUploaded
			if err = models.UpdateChunkUploadedStatusByUUID(fileChunk.IsUploaded, fileChunk.UUID); err != nil {
				log.Error("UpdateFileChunk failed:%v", err)
			}
		}
		return &entity.UploadChunkResponse{
			UUID:     fileChunk.UUID,
			Uploaded: fileChunk.IsUploaded,
			Chunks:   chunks,
			FileName: fileName,
		}, nil
	}
	if fileChunk.IsUploaded == models.DataFileNotUploaded {
		log.Info("the file has been recorded but not uploaded")
		fileChunk.IsUploaded = models.FileNotUploaded
		if err = models.UpdateChunkUploadedStatusByUUID(fileChunk.IsUploaded, fileChunk.UUID); err != nil {
			log.Error("UpdateFileChunk failed:%v", err)
		}
	}
	parts, err := storageHelper.GetPartInfos(fileChunk.ObjectName, fileChunk.UploadID)
	if err != nil {
		log.Error("GetPartInfos failed:%v", err)
		models.DeleteUploadChunk(fileChunk)
		return &entity.UploadChunkResponse{}, nil
	}

	for _, part := range parts {
		chunks = append(chunks, part.PartNumber)
	}

	return &entity.UploadChunkResponse{
		UUID:     fileChunk.UUID,
		Uploaded: fileChunk.IsUploaded,
		Chunks:   chunks,
		FileName: fileName,
	}, nil
}

var uploadMutex *sync.Mutex = new(sync.Mutex)

func NewUploadMultipart(req entity.NewMultipartRequest) (*entity.NewMuiltipartResponse, error) {
	if !setting.Attachment.Enabled {
		return nil, errors.New("attachment is not enabled")
	}
	if err := CheckFlow(req.User.ID, req.FileName, req.Size, req.SubjectId, req.SubjectType); err != nil {
		log.Error("CheckFlow failed: %v", err)
		return nil, err
	}

	uploadHelper := GetUploadHelperBySubjectContext(req.SubjectAccessContext)
	if uploadHelper == nil {
		log.Error("GetUploadHelper failed, dataType=%d", req.SubjectType)
		return nil, errors.New("Data type error")
	}
	storageType := uploadHelper.GetStorageType()
	uploadMutex.Lock()
	defer uploadMutex.Unlock()

	err := upload.VerifyFileType(req.FileType, strings.Split(setting.Attachment.AllowedTypes, ","))
	if err != nil {
		return nil, err
	}
	storageHelper := storage_helper.SelectStorageHelperFromStorageType(storageType)
	if storageHelper == nil {
		log.Error("SelectStorageHelperFromStorageType failed, StorageType=%d", storageType)
		return nil, errors.New("Storage type error")
	}

	if req.TotalChunkCounts > minio_ext.MaxPartsCount {
		return nil, errors.New(fmt.Sprintf("chunk counts(%d) is too much", req.TotalChunkCounts))

	}

	isLimited := storage_limit.IsSubjectFileUploadOverLimit(req.Size, req.SubjectAccessContext)
	if isLimited {
		return nil, errors.New("common_error.file_size_storage_limit")
	}

	if req.Size > minio_ext.MaxMultipartPutObjectSize {
		return nil, errors.New(fmt.Sprintf("file size(%d) is too big", req.Size))
	}

	uuid := util.UUID()
	objectName := uploadHelper.GetFileStoragePath(req.FileName, req.SubjectId)
	uploadID, err := storageHelper.NewMultiPartUpload(objectName)
	if err != nil {
		log.Error("NewMultiPartUpload err.req = %+v err=%v", req, err)
		return nil, err
	}
	_, err = models.InsertUploadChunk(&models.UploadChunk{
		UUID:        uuid,
		Md5:         req.MD5,
		SubjectID:   req.SubjectId,
		ObjectName:  objectName,
		UploadID:    uploadID,
		TotalChunks: req.TotalChunkCounts,
		Size:        req.Size,
		UserID:      req.User.ID,
		StorageType: string(storageType),
		SubjectType: req.SubjectType,
	})
	if err != nil {
		return nil, err

	}
	err = SetFlowCache(req.User.ID, req.SubjectType, req.SubjectId, req.FileName, req.Size)
	if err != nil {
		log.Error("NewUploadMultipart SetFlowCacheAfterUploaded err.req = %+v err=%v", req, err)
	}
	return &entity.NewMuiltipartResponse{
		UUID: uuid,
	}, nil

}

func GetMultipartUploadUrl(req entity.GetMultipartUrlRequest) (string, error) {
	uuid := req.UUID
	partNumber := req.PartNumber
	size := req.Size

	fileChunk, err := models.GetUploadChunkByUUID(uuid)
	if err != nil {
		return "", err
	}
	storageHelper := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(fileChunk.StorageType))
	if storageHelper == nil {
		log.Error("SelectStorageHelperFromStorageType failed, StorageType=%d", fileChunk.StorageType)
		return "", errors.New("Storage type error")
	}
	url, err := storageHelper.GenMultiPartSignedUrl(fileChunk.ObjectName, fileChunk.UploadID, partNumber, size)
	if err != nil {
		log.Error("GenMultiPartSignedUrl failed: %v", err)
		return "", err
	}

	return url, nil
}

func CompleteModelMultipart(req entity.CompleteMultipartRequest) error {
	uuid := req.UUID
	fileChunk, err := models.GetUploadChunkByUUID(uuid)
	if err != nil {
		return err
	}

	storageType := entity.StorageType(fileChunk.StorageType)
	storageHelper := storage_helper.SelectStorageHelperFromStorageType(storageType)
	if storageHelper == nil {
		log.Error("SelectStorageHelperFromStorageType failed, StorageType=%d", storageType)
		return errors.New("Storage type error")
	}
	uploadHelper := GetUploadHelper(models.SubjectType(fileChunk.SubjectType))
	if uploadHelper == nil {
		log.Error("GetUploadHelper failed, dataType=%d", fileChunk.SubjectType)
		return errors.New("Data type error")
	}
	_, err = storageHelper.CompleteMultiPartUpload(fileChunk.ObjectName, fileChunk.UploadID, fileChunk.TotalChunks)
	if err != nil {
		log.Error("CompleteMultiPartUpload failed: %v", err)
		return err
	}

	err = models.UpdateChunkUploadedStatusByUUID(models.FileUploaded, uuid)
	if err != nil {
		return err
	}
	uploadHelper.DoAfterUploadedSuccess(fileChunk.SubjectID)
	fileName := uploadHelper.GetFileNameByStoragePath(fileChunk.ObjectName, fileChunk.SubjectID)
	UpdateFlowCacheAfterUploaded(fileChunk.SubjectType, fileChunk.SubjectID, fileName)
	return nil
}

func GetUploadUrl(req entity.GetUploadUrlRequest) (string, error) {
	subjectCtx := req.SubjectContext
	if !setting.Attachment.Enabled {
		return "", errors.New("attachment is not enabled")
	}
	if err := CheckFlow(req.User.ID, req.FileName, req.Size, subjectCtx.SubjectID, int(subjectCtx.SubjectType)); err != nil {
		log.Error("CheckFlow failed: %v", err)
		return "", err
	}
	uploadHelper := GetUploadHelperBySubjectContext(subjectCtx)
	if uploadHelper == nil {
		log.Error("GetUploadHelper failed, dataType=%d", subjectCtx.SubjectType)
		return "", errors.New("Data type error")
	}
	storageType := uploadHelper.GetStorageType()
	storageHelper := storage_helper.SelectStorageHelperFromStorageType(storageType)
	if storageHelper == nil {
		log.Error("SelectStorageHelperFromStorageType failed, StorageType=%d", storageType)
		return "", errors.New("Storage type error")
	}
	err := upload.VerifyFileType(req.FileType, strings.Split(setting.Attachment.AllowedTypes, ","))
	if err != nil {
		return "", err
	}
	isLimited := storage_limit.IsSubjectFileUploadOverLimit(req.Size, subjectCtx)
	if isLimited {
		if subjectCtx.Owner.IsOrganization() {
			if subjectCtx.SubjectType == models.DatasetSubject {
				return "", errors.New("dataset_registry.org_not_allowed_to_upload_dataset_file")
			} else {
				return "", errors.New("aimodel.org_not_allowed_to_upload_aimodel_file")
			}
		}
		return "", errors.New("common_error.file_size_storage_limit")
	}
	path := uploadHelper.GetFileStoragePath(req.FileName, subjectCtx.SubjectID)

	url, err := storageHelper.GenSignedUrl(path)
	if err != nil {
		log.Error("GenSignedUrl failed: %v", err)
		return "", err
	}
	SetFlowCache(req.User.ID, int(subjectCtx.SubjectType), subjectCtx.SubjectID, req.FileName, req.Size)
	return url, nil
}

func CompleteUpload(req entity.CompleteUploadRequest) error {
	uploadHelper := GetUploadHelperBySubjectContext(req.SubjectContext)
	if uploadHelper == nil {
		log.Error("GetUploadHelper failed, dataType=%d", req.SubjectContext.SubjectType)
		return errors.New("Data type error")
	}
	go uploadHelper.DoAfterUploadedSuccess(req.SubjectContext.SubjectID)
	go UpdateFlowCacheAfterUploaded(int(req.SubjectContext.SubjectType), req.SubjectContext.SubjectID, req.FileNameList...)

	return nil
}
