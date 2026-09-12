package storage_helper

import (
	"fmt"
	"io"
	"net/url"
	"path"
	"strconv"
	"strings"

	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/obs"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
)

type UrchinV2Helper struct {
	Bucket string
}

// objectKey = collectionID + "/" + relativePath
func (c *UrchinV2Helper) splitPath(p string) (collectionID, relativePath string) {
	p = strings.Trim(p, "/")
	idx := strings.Index(p, "/")
	if idx == -1 {
		return p, ""
	}
	return p[:idx], p[idx+1:]
}

func (m *UrchinV2Helper) AllocateDatasetNamespace(name, prefix string) (string, error) {
	return storage.UrchinClient.CreateCollection(name)
}

func (m *UrchinV2Helper) UploadDir(localPath, objectKeyPrefix string) error {
	collectionId, relativePath := m.splitPath(objectKeyPrefix)
	return storage.UrchinClient.UploadLocalDir(collectionId, localPath, relativePath)
}

func (m *UrchinV2Helper) UploadFile(objectKey string, r io.Reader) error {
	collectionId, relativePath := m.splitPath(objectKey)
	return storage.UrchinClient.UploadFile(collectionId, relativePath, r)
}

func (m *UrchinV2Helper) GetJobDefaultObjectKeyPrefix(jobName string) string {
	return path.Join(setting.CodePathPrefix, jobName)
}

func (m *UrchinV2Helper) GetRealPath(objectKey string) string {
	return ""
}

func (m *UrchinV2Helper) GetBucket() string {
	return ""
}
func (m *UrchinV2Helper) MKDIR(path string, description ...string) error {
	collectionId, relativePath := m.splitPath(path)
	return storage.UrchinClient.MKDIR(collectionId, relativePath)
}

func (m *UrchinV2Helper) GetEndpoint() string {
	return ""
}

func (m *UrchinV2Helper) GetDataId(path string) string {
	return path
}

func (m *UrchinV2Helper) GetOneLevelObjectsUnderDir(dirPath string, maxKeyArray ...int) ([]storage.FileInfo, error) {
	var maxKey int
	if len(maxKeyArray) <= 0 {
		maxKey = setting.OUTPUT_SHOW_MAX_KEY + 1
	}

	index := 1
	fileInfoList := make([]storage.FileInfo, 0)
	marker := ""
	for {
		output, err := m.GetOneLevelObjectsUnderDirWithMarker(dirPath, marker, maxKey)
		if err != nil {
			return nil, err
		}
		fileInfoList = append(fileInfoList, output.Objects...)

		log.Info("GetOneLevelObjectsUnderDir Page:%d\n", index)
		index++

		if output.IsTruncated {
			marker = output.NextMarker
		} else {
			break
		}
	}

	return fileInfoList, nil
}

func (m *UrchinV2Helper) GetOneLevelObjectsUnderDirWithMarker(dirPath string, marker string, maxKey int) (*ObjectListResponse, error) {
	collectionId, relativePath := m.splitPath(dirPath)

	if maxKey <= 0 {
		maxKey = setting.OUTPUT_SHOW_MAX_KEY
	}
	if relativePath != "" && !strings.HasSuffix(relativePath, "/") {
		relativePath = relativePath + "/"
	}
	res, err := storage.UrchinClient.ListPrefixObjectsWithMarkerAndDelimeter(collectionId, relativePath, marker, int32(maxKey))
	if err != nil {
		return nil, err
	}
	if res == nil {
		return &ObjectListResponse{}, nil
	}
	files := make([]storage.FileInfo, 0, len(res.Data.List))
	for _, val := range res.Data.CommonPrefixes {
		fileName := strings.TrimSuffix(strings.TrimPrefix(val, relativePath), "/")
		// fileName := strings.TrimSuffix(val, "/")
		file := storage.FileInfo{
			FileName: fileName,
			IsDir:    true,
		}
		files = append(files, file)
	}
	for i := 0; i < len(res.Data.List); i++ {
		val := res.Data.List[i]
		if val.Key == relativePath {
			continue
		}
		sizeStr := val.Size
		var size int64 = 0
		if sizeStr != "" {
			sizeTmp, err := strconv.ParseInt(sizeStr, 10, 64)
			if err == nil {
				size = sizeTmp
			}
		}
		file := storage.FileInfo{
			ModTime:      val.LastModified.Local().Format("2006-01-02 15:04:05"),
			FileName:     strings.TrimPrefix(val.Key[len(relativePath):], "/"),
			Size:         size,
			IsDir:        false,
			ParenDir:     "",
			RelativePath: path.Join(collectionId, strings.TrimPrefix(val.Key, collectionId)),
			FullPath:     path.Join(collectionId, strings.TrimPrefix(val.Key, collectionId)),
		}
		files = append(files, file)
	}

	return &ObjectListResponse{
		NextMarker:  res.Data.NextMarker,
		IsTruncated: res.Data.IsTruncated,
		Objects:     files,
	}, nil

}

func (m *UrchinV2Helper) GetAllObjectsUnderDir(prefix string, maxKeyArray ...int) ([]storage.FileInfo, error) {
	index := 1
	fileInfoList := make([]storage.FileInfo, 0)
	marker := ""
	for {
		output, err := m.GetAllObjectsUnderDirWithMarker(prefix, marker, maxKeyArray...)
		if err != nil {
			return nil, err
		}
		fileInfoList = append(fileInfoList, output.Objects...)

		log.Info("GetAllObjectsUnderDir Page:%d\n", index)
		index++

		if output.IsTruncated {
			marker = output.NextMarker
		} else {
			break
		}
	}

	return fileInfoList, nil
}

func (m *UrchinV2Helper) GetAllObjectsUnderDirWithMarker(prefix string, marker string, maxKeyArray ...int) (*ObjectListResponse, error) {
	collectionId, relativePath := m.splitPath(prefix)

	var maxKey int
	if len(maxKeyArray) <= 0 {
		maxKey = 1000
	} else {
		maxKey = maxKeyArray[0]
	}
	res, err := storage.UrchinClient.ListPrefixObjectsWithMarker(collectionId, marker, relativePath, int32(maxKey))
	if err != nil {
		return nil, err
	}
	if res == nil || len(res.Data.List) == 0 {
		return &ObjectListResponse{}, nil
	}
	files := make([]storage.FileInfo, 0, len(res.Data.List))

	for i := 0; i < len(res.Data.List); i++ {
		val := res.Data.List[i]
		if val.Key == relativePath {
			continue
		}
		sizeStr := val.Size
		var size int64 = 0
		if sizeStr != "" {
			sizeTmp, err := strconv.ParseInt(sizeStr, 10, 64)
			if err == nil {
				size = sizeTmp
			}
		}
		var isDir bool
		if strings.HasSuffix(val.Key, "/") {
			isDir = true
		} else {
			isDir = false
		}
		if isDir {
			continue
		}
		file := storage.FileInfo{
			ModTime:      val.LastModified.Local().Format("2006-01-02 15:04:05"),
			FileName:     strings.TrimPrefix(val.Key[len(relativePath):], "/"),
			Size:         size,
			IsDir:        isDir,
			ParenDir:     "",
			RelativePath: path.Join(collectionId, strings.TrimPrefix(val.Key, collectionId)),
			FullPath:     path.Join(collectionId, strings.TrimPrefix(val.Key, collectionId)),
		}
		files = append(files, file)
	}

	return &ObjectListResponse{
		NextMarker:  res.Data.NextMarker,
		IsTruncated: res.Data.IsTruncated,
		Objects:     files,
	}, nil
}

func (m *UrchinV2Helper) TrimBucketPrefix(path string) string {
	return path
}

func (m *UrchinV2Helper) OpenFile(path string) (io.ReadCloser, error) {
	collectionId, relativePath := m.splitPath(path)
	res, err := storage.UrchinClient.GetObject(collectionId, relativePath)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (m *UrchinV2Helper) GetObject(path string) (*ObjectResponse, error) {
	collectionId, relativePath := m.splitPath(path)

	res, err := storage.UrchinClient.GetObject(collectionId, relativePath)
	if err != nil {
		return nil, err
	}
	return &ObjectResponse{
		ContentLength: res.ContentLength,
		ContentType:   res.ContentType,
		Body:          res.Body,
	}, nil
}

func (m *UrchinV2Helper) GetObjectMeta(path string) (*ObjectMeta, error) {
	collectionId, relativePath := m.splitPath(path)

	res, err := storage.UrchinClient.GetObjectMeta(collectionId, relativePath)
	if err != nil {
		return nil, err
	}
	return &ObjectMeta{
		ContentLength: res.ContentLength,
		ContentType:   res.ContentType,
		LastModified:  res.LastModified,
	}, nil
}

func (m *UrchinV2Helper) GetSignedDownloadUrl(sourcePath string) (string, error) {
	collectionId, relativePath := m.splitPath(sourcePath)

	reqParams := make(map[string]string)

	name := path.Base(relativePath)
	name = url.PathEscape(name)

	reqParams["response-content-disposition"] =
		"attachment; filename=\"" + name + "\""
	url, err := storage.UrchinClient.GetObjectSignedUrl(collectionId, relativePath, reqParams)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (m *UrchinV2Helper) GetS3DownloadUrl(path string) string {
	return ""
}

func (m *UrchinV2Helper) CopyDir(sourcePath, targetPath string, filterSuffix []string) error {
	sourceCollectionId, sourceRelativePath := m.splitPath(sourcePath)
	targetCollectionId, targetRelativePath := m.splitPath(targetPath)

	index := 1
	marker := ""
	for {
		output, err := m.GetAllObjectsUnderDirWithMarker(sourcePath, marker, 1000)
		if err != nil || output == nil {
			log.Error("CopyDir GetAllObjectsUnderDirWithMarker error.ourcePath=%s targetPath=%s err=%v", sourcePath, targetPath, err)
			return err
		}
		for i := 0; i < len(output.Objects); i++ {
			file := output.Objects[i]
			if isMatchSuffix(file.FileName, filterSuffix) {
				continue
			}
			sourceFilePath := path.Join(sourceRelativePath, file.RelativePath)
			targetFilePath := path.Join(targetRelativePath, file.RelativePath)
			err := storage.UrchinClient.CopyObject(sourceCollectionId, sourceFilePath, targetCollectionId, targetFilePath)
			if err != nil {
				log.Error("CopyDir CopyObject error.ourcePath=%s targetPath=%s err=%v", sourcePath, targetPath, err)
				return err
			}
		}

		log.Info("CopyDir GetAllObjectsUnderDirWithMarker Page:%d\n", index)
		index++

		if output.IsTruncated {
			marker = output.NextMarker
		} else {
			break
		}
	}
	return nil
}

func (m *UrchinV2Helper) CopyFile(sourcePath, targetPath string) error {
	sourceCollectionId, sourceRelativePath := m.splitPath(sourcePath)

	targetCollectionId, targetRelativePath := m.splitPath(targetPath)

	return storage.UrchinClient.CopyObject(sourceCollectionId, sourceRelativePath, targetCollectionId, targetRelativePath)
}

func (m *UrchinV2Helper) DeleteDir(dirPath string) (remainingFiles bool, err error) {
	collectionId, relativePath := m.splitPath(dirPath)

	index := 1
	marker := ""
	for {
		output, err := m.GetAllObjectsUnderDirWithMarker(dirPath, marker, 1000)
		if err != nil {
			log.Error("DeleteDir GetAllObjectsUnderDirWithMarker error.dirPath=%s err=%v", dirPath, err)
			return true, err
		}
		for i := 0; i < len(output.Objects); i++ {
			file := output.Objects[i]
			filePath := path.Join(relativePath, file.RelativePath)
			err := storage.UrchinClient.DeleteFile(collectionId, filePath)
			if err != nil {
				log.Error("DeleteDir: delete file error.collectionId=%s filePath=%s err=%v", collectionId, filePath, err)
				remainingFiles = true
				continue
			}
		}

		log.Info("DeleteDir GetAllObjectsUnderDirWithMarker Page:%d\n", index)
		index++

		if output.IsTruncated {
			marker = output.NextMarker
		} else {
			break
		}
	}
	return false, nil
}

func (m *UrchinV2Helper) DeleteCollection(dirPath string) (remainingFiles bool, err error) {
	collectionId, _ := m.splitPath(dirPath)

	err = storage.UrchinClient.DeleteObject(collectionId)
	if err != nil {
		log.Error("DeleteCollection: delete object error.collectionId=%s err=%v", collectionId, err)
		return true, err
	}
	return false, nil
}

func (m *UrchinV2Helper) DeleteFile(path string) error {
	collectionId, relativePath := m.splitPath(path)

	return storage.UrchinClient.DeleteFile(collectionId, relativePath)
}

func (m *UrchinV2Helper) HasObject(path string) (bool, error) {
	collectionId, relativePath := m.splitPath(path)

	res, err := storage.UrchinClient.GetObjectMeta(collectionId, relativePath)
	if err != nil || res == nil {
		return false, nil
	}
	return true, nil
}

func (m *UrchinV2Helper) GetPartInfos(objectName string, uploadID string) ([]PartInfo, error) {
	result := make([]PartInfo, 0)

	output := &obs.ListPartsOutput{}
	partNumberMarker := 0
	for {
		temp, err := storage.ObsCli.ListParts(&obs.ListPartsInput{
			Bucket:           m.GetBucket(),
			Key:              objectName,
			UploadId:         uploadID,
			MaxParts:         MAX_LIST_PARTS,
			PartNumberMarker: partNumberMarker,
		})
		if err != nil {
			log.Error("ListParts failed:", err.Error())
			return nil, err
		}

		partNumberMarker = temp.NextPartNumberMarker
		log.Info("uuid:%s, MaxParts:%d, PartNumberMarker:%d, NextPartNumberMarker:%d, len:%d", objectName, temp.MaxParts, temp.PartNumberMarker, temp.NextPartNumberMarker, len(temp.Parts))

		for _, partInfo := range temp.Parts {
			output.Parts = append(output.Parts, obs.Part{
				PartNumber: partInfo.PartNumber,
				ETag:       partInfo.ETag,
			})
		}

		if !temp.IsTruncated {
			break
		} else {
			continue
		}
	}
	for _, part := range output.Parts {
		result = append(result, PartInfo{
			PartNumber:   part.PartNumber,
			ETag:         part.ETag,
			Size:         part.Size,
			LastModified: part.LastModified,
		})
	}
	return result, nil
}

func (m *UrchinV2Helper) NewMultiPartUpload(objectName string) (string, error) {
	collectionId, relativePath := m.splitPath(objectName)
	return storage.UrchinClient.InitiateMultipartUpload(collectionId, relativePath)
}

func (m *UrchinV2Helper) GenMultiPartSignedUrl(objectName string, uploadId string, partNumber int, partSize int64) (string, error) {
	collectionId, relativePath := m.splitPath(objectName)
	return storage.UrchinClient.CreateMultipartUploadSignedUrl(collectionId, uploadId, relativePath, partNumber)
}

func (m *UrchinV2Helper) GenSignedUrl(objectKey string) (string, error) {
	collectionId, relativePath := m.splitPath(objectKey)
	return storage.UrchinClient.CreatePutObjectSignedUrl(collectionId, relativePath)

}

func (m *UrchinV2Helper) CompleteMultiPartUpload(objectName string, uploadID string, totalChunks int) (string, error) {
	collectionId, relativePath := m.splitPath(objectName)
	return "", storage.UrchinClient.CompleteMultiPartUpload(collectionId, uploadID, relativePath, totalChunks)
}

func (m *UrchinV2Helper) GetAllObjectByBucketAndPrefix(bucket string, prefix string) ([]storage.FileInfo, error) {
	return nil, nil

}

func (m *UrchinV2Helper) GetFilesSize(Files []string) int64 {
	var fileTotalSize int64
	for _, file := range Files {
		collectionId, relativePath := m.splitPath(file)
		out, err := storage.UrchinClient.GetObjectMeta(collectionId, relativePath)
		if err != nil {
			log.Info("Get File error, error=" + err.Error())
			continue
		}
		fileTotalSize += out.ContentLength
	}
	return fileTotalSize

}
func (m *UrchinV2Helper) RemoveObject(path string) error {
	remainingFiles, err := m.DeleteDir(path)
	if err != nil {
		return err
	}
	if remainingFiles {
		return fmt.Errorf("Some files failed to be deleted.")
	}
	return nil
}

func (m *UrchinV2Helper) PutString(objectKey string, content string) error {
	return m.UploadFile(objectKey, strings.NewReader(content))

}

func (m *UrchinV2Helper) GetOutputObjectKeyPrefix(jobName string, computeResource string, versionName string) string {
	objectkey := path.Join(m.GetJobDefaultObjectKeyPrefix(jobName), setting.OutPutPath, versionName) + "/"
	if computeResource != "NPU" {
		objectkey = path.Join(m.GetJobDefaultObjectKeyPrefix(jobName), cloudbrain.ModelMountPath, versionName) + "/"
	}
	return objectkey
}

func (m *UrchinV2Helper) CountDirSize(objectKey string) (int64, error) {
	var fileTotalSize int64
	index := 1
	marker := ""
	for {
		output, err := m.GetAllObjectsUnderDirWithMarker(objectKey, marker, 1000)
		if err != nil {
			log.Error("CountDirSize: GetAllObjectsUnderDirWithMarker error.dirPath=%s err=%v", objectKey, err)
			return 0, err
		}
		for i := 0; i < len(output.Objects); i++ {
			fileTotalSize += output.Objects[i].Size
		}

		log.Info("CountD:irSize GetAllObjectsUnderDirWithMarker Page:%d\n", index)
		index++

		if output.IsTruncated {
			marker = output.NextMarker
		} else {
			break
		}
	}
	return fileTotalSize, nil
}

func (m *UrchinV2Helper) ReportSizeChanged(objectKey string) error {
	collectionId, _ := m.splitPath(objectKey)
	err := storage.UrchinClient.ReportSizeChanged(collectionId)
	if err != nil {
		log.Error("ReportSizeChanged err. objectKey=%s err=%v", err)
	}
	err2 := storage.UrchinClient.UpdateObjectLastModifyTime(collectionId)
	if err2 != nil {
		log.Error("UpdateObjectLastModifyTime err. objectKey=%s err=%v", err2)
	}
	return err
}
