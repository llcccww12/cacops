package storage_helper

import (
	"encoding/json"
	"io"
	"net/url"
	"path"
	"sort"
	"strings"

	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/obs"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
)

type OBSHelper struct {
	Bucket string
}

func (m *OBSHelper) AllocateDatasetNamespace(name, prefix string) (string, error) {
	return prefix, nil
}

func (m *OBSHelper) UploadDir(codePath, objectKeyPrefix string) error {
	objectKeyPrefix = m.TrimBucketPrefix(objectKeyPrefix)

	return UploadDirToObs(codePath, objectKeyPrefix, "")
}

func (m *OBSHelper) UploadFile(objectKey string, r io.Reader) error {
	objectKey = m.TrimBucketPrefix(objectKey)
	input := &obs.PutObjectInput{}
	input.Bucket = m.GetBucket()
	input.Key = objectKey
	input.Body = r
	_, err := storage.ObsCli.PutObject(input)
	if err != nil {
		return err
	}
	return nil
}

func (m *OBSHelper) GetJobDefaultObjectKeyPrefix(jobName string) string {
	return path.Join(setting.CodePathPrefix, jobName)
}

func (m *OBSHelper) GetRealPath(objectKey string) string {
	return ""
}

func (m *OBSHelper) GetBucket() string {
	if m.Bucket != "" {
		return m.Bucket
	}
	return setting.Bucket
}

func (m *OBSHelper) GetDataId(path string) string {
	return ""
}

func (m *OBSHelper) MKDIR(path string, description ...string) error {
	path = m.TrimBucketPrefix(path)
	path = strings.TrimSuffix(path, "/") + "/"
	input := &obs.PutObjectInput{}
	input.Bucket = setting.Bucket
	input.Key = path
	input.Metadata = map[string]string{
		"type": "dir",
	}
	_, err := storage.ObsCli.PutObject(input)
	if err != nil {
		log.Error("PutObject(%s) failed: %s", input.Key, err.Error())
		return err
	}

	return nil
}

func (m *OBSHelper) GetEndpoint() string {
	index := strings.Index(setting.Endpoint, "//")
	endpoint := setting.Endpoint[index+2:]
	return endpoint
}

func (m *OBSHelper) GetOneLevelObjectsUnderDir(dirPath string, maxKeyArray ...int) ([]storage.FileInfo, error) {
	maxKey := setting.OUTPUT_SHOW_MAX_KEY + 1
	if len(maxKeyArray) > 0 {
		maxKey = maxKeyArray[0]
	}
	resp, err := m.GetOneLevelObjectsUnderDirWithMarker(dirPath, "", maxKey)
	if err != nil {
		return nil, err
	}
	return resp.Objects, nil
}

func (m *OBSHelper) GetOneLevelObjectsUnderDirWithMarker(dirPath string, marker string, maxKey int) (*ObjectListResponse, error) {
	dirPath = m.TrimBucketPrefix(dirPath)
	input := &obs.ListObjectsInput{}
	input.Bucket = m.GetBucket()
	input.Prefix = dirPath
	input.Delimiter = "/"
	input.Marker = marker
	if maxKey <= 0 {
		maxKey = setting.OUTPUT_SHOW_MAX_KEY
	}
	input.MaxKeys = maxKey
	if !strings.HasSuffix(input.Prefix, "/") {
		input.Prefix += "/"
	}
	fileInfos := make([]storage.FileInfo, 0)
	prefixLen := len(input.Prefix)
	index := 1
	output, err := storage.ObsCli.ListObjects(input)
	if err != nil {
		if obsError, ok := err.(obs.ObsError); ok {
			log.Error("Code:%s, Message:%s", obsError.Code, obsError.Message)
		}
		return nil, err
	}
	log.Info("Page:%d\n", index)
	index++
	for _, val := range output.Contents {
		var fileName string
		if val.Key == input.Prefix {
			continue
		}
		fileName = val.Key[prefixLen:]
		fileInfo := storage.FileInfo{
			ModTime:      val.LastModified.Local().Format("2006-01-02 15:04:05"),
			FileName:     fileName,
			Size:         val.Size,
			IsDir:        false,
			RelativePath: strings.TrimSuffix(dirPath, "/") + "/" + fileName,
			FullPath:     val.Key,
		}
		fileInfos = append(fileInfos, fileInfo)
	}
	for _, val := range output.CommonPrefixes {
		fileName := strings.TrimSuffix(strings.TrimPrefix(val, input.Prefix), "/")
		fileInfo := storage.FileInfo{
			FileName: fileName,
			IsDir:    true,
		}
		fileInfos = append(fileInfos, fileInfo)
	}
	return &ObjectListResponse{
		NextMarker:  output.NextMarker,
		IsTruncated: output.IsTruncated,
		Objects:     fileInfos,
	}, nil
}

func (m *OBSHelper) GetAllObjectsUnderDir(prefix string, maxKeyArray ...int) ([]storage.FileInfo, error) {
	//某些情况下，prefix可能是以bucket开头的，所以需要去掉bucket前缀
	prefix = m.TrimBucketPrefix(prefix)
	prefix = strings.TrimSuffix(prefix, "/") + "/"
	bucket := m.GetBucket()
	input := &obs.ListObjectsInput{}
	input.Bucket = bucket
	input.MaxKeys = 1000
	input.Prefix = prefix
	maxKey := setting.OUTPUT_DOWNLOAD_MAX_KEY
	if len(maxKeyArray) > 0 {
		maxKey = maxKeyArray[0]
	}
	input.MaxKeys = maxKey

	index := 1
	fileInfoList := storage.FileInfoList{}

	prefixLen := len(prefix)
	log.Info("full obs path:", input.Bucket+input.Prefix)
	log.Info("prefix=" + input.Prefix)
	for {
		output, err := storage.ObsCli.ListObjects(input)
		if err == nil {
			log.Info("Page:%d\n", index)
			index++
			for _, val := range output.Contents {
				var isDir bool
				if prefixLen == len(val.Key) {
					continue
				}
				if strings.HasSuffix(val.Key, "/") {
					isDir = true
				} else {
					isDir = false
				}
				if isDir {
					continue
				}
				fileInfo := storage.FileInfo{
					ModTime:      val.LastModified.Format("2006-01-02 15:04:05"),
					FileName:     strings.TrimPrefix(val.Key[prefixLen:], "/"),
					Size:         val.Size,
					IsDir:        isDir,
					ParenDir:     "",
					RelativePath: val.Key,
					FullPath:     val.Key,
				}
				fileInfoList = append(fileInfoList, fileInfo)
			}
			if output.IsTruncated {
				input.Marker = output.NextMarker
			} else {
				break
			}
		} else {
			if obsError, ok := err.(obs.ObsError); ok {
				log.Info("Code:%s\n", obsError.Code)
				log.Info("Message:%s\n", obsError.Message)
			}
			return nil, err
		}
	}
	sort.Sort(fileInfoList)
	return fileInfoList, nil
}

func (m *OBSHelper) GetAllObjectsUnderDirWithMarker(prefix string, marker string, maxKeyArray ...int) (*ObjectListResponse, error) {
	//某些情况下，prefix可能是以bucket开头的，所以需要去掉bucket前缀
	prefix = m.TrimBucketPrefix(prefix)
	prefix = strings.TrimSuffix(prefix, "/") + "/"
	bucket := m.GetBucket()
	input := &obs.ListObjectsInput{}
	input.Bucket = bucket
	input.MaxKeys = 1000
	input.Prefix = prefix
	input.Marker = marker
	maxKey := setting.OUTPUT_DOWNLOAD_MAX_KEY
	if len(maxKeyArray) > 0 {
		maxKey = maxKeyArray[0]
	}
	input.MaxKeys = maxKey

	index := 1
	fileInfoList := storage.FileInfoList{}

	prefixLen := len(prefix)
	output, err := storage.ObsCli.ListObjects(input)
	if err == nil {
		index++
		for _, val := range output.Contents {
			var isDir bool
			if prefixLen == len(val.Key) {
				continue
			}
			if strings.HasSuffix(val.Key, "/") {
				isDir = true
			} else {
				isDir = false
			}
			if isDir {
				continue
			}
			fileInfo := storage.FileInfo{
				ModTime:      val.LastModified.Format("2006-01-02 15:04:05"),
				FileName:     strings.TrimPrefix(val.Key[prefixLen:], "/"),
				Size:         val.Size,
				IsDir:        isDir,
				ParenDir:     "",
				RelativePath: val.Key,
				FullPath:     val.Key,
			}
			fileInfoList = append(fileInfoList, fileInfo)
		}
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			log.Error("OBS error %v", obsError)
		}
		return nil, err
	}
	return &ObjectListResponse{
		NextMarker:  output.NextMarker,
		IsTruncated: output.IsTruncated,
		Objects:     fileInfoList,
	}, nil
}

func (m *OBSHelper) TrimBucketPrefix(path string) string {
	return strings.TrimPrefix(path, m.GetBucket()+"/")
}

func (m *OBSHelper) OpenFile(path string) (io.ReadCloser, error) {
	path = m.TrimBucketPrefix(path)
	input := &obs.GetObjectInput{}
	input.Bucket = m.GetBucket()
	input.Key = path
	output, err := storage.ObsCli.GetObject(input)
	if err != nil {
		log.Error("OpenFile err. path=%s err=%v", path, err)
		return nil, err
	}
	return output.Body, nil
}

func (m *OBSHelper) GetObject(path string) (*ObjectResponse, error) {
	path = m.TrimBucketPrefix(path)
	input := &obs.GetObjectInput{}
	input.Bucket = m.GetBucket()
	input.Key = path
	output, err := storage.ObsCli.GetObject(input)
	if err != nil {
		log.Error("OpenFile err. path=%s err=%v", path, err)
		return nil, err
	}
	return &ObjectResponse{
		ContentLength: output.ContentLength,
		ContentType:   output.ContentType,
		Body:          output.Body,
	}, nil
}

func (m *OBSHelper) GetObjectMeta(path string) (*ObjectMeta, error) {
	path = m.TrimBucketPrefix(path)
	output, err := storage.ObsCli.GetObjectMetadata(&obs.GetObjectMetadataInput{
		Bucket: m.GetBucket(),
		Key:    path,
	})
	if err != nil {
		log.Error("OpenFile err. path=%s err=%v", path, err)
		return nil, err
	}
	return &ObjectMeta{
		ContentLength: output.ContentLength,
		ContentType:   output.ContentType,
		LastModified:  output.LastModified,
	}, nil
}

func (m *OBSHelper) GetSignedDownloadUrl(key string) (string, error) {
	key = m.TrimBucketPrefix(key)
	input := &obs.CreateSignedUrlInput{}
	input.Bucket = m.GetBucket()
	input.Key = key

	input.Expires = 60 * 60
	input.Method = obs.HttpMethodGet
	comma := strings.LastIndex(key, "/")
	filename := key
	if comma != -1 {
		filename = key[comma+1:]
	}
	reqParams := make(map[string]string)
	filename = url.PathEscape(filename)
	reqParams["response-content-disposition"] = "attachment; filename=\"" + filename + "\""
	input.QueryParams = reqParams
	output, err := storage.ObsCli.CreateSignedUrl(input)
	if err != nil {
		log.Error("CreateSignedUrl failed:", err.Error())
		return "", err
	}

	return output.SignedUrl, nil
}

func (m *OBSHelper) GetS3DownloadUrl(key string) string {
	key = m.TrimBucketPrefix(key)
	return "s3://" + setting.Bucket + "/" + strings.TrimPrefix(key, "/")
}

func (m *OBSHelper) CopyDir(sourcePath, targetPath string, filterSuffix []string) error {
	log.Info("CopyDir sourcePath=%s,targetPath=%s", sourcePath, targetPath)
	allFiles, _ := m.GetAllObjectsUnderDir(sourcePath)
	var fileNames []string
	for _, file := range allFiles {
		if isMatchSuffix(file.FileName, filterSuffix) {
			continue
		}
		fileNames = append(fileNames, file.FileName)
	}
	log.Info("Previous task all files", fileNames)
	_, err := storage.ObsCopyManyFile(m.GetBucket(), sourcePath, m.GetBucket(), targetPath, fileNames)
	if err != nil {
		log.Error("CopyDir ObsCopyManyFile error. sourcePath=%s targetPath=%s err=%v", sourcePath, targetPath, err)
		return err
	}
	return nil
}

func (m *OBSHelper) CopyFile(sourcePath, targetPath string) error {
	sourcePath = m.TrimBucketPrefix(sourcePath)
	targetPath = m.TrimBucketPrefix(targetPath)

	err := storage.ObsCopyFile(m.GetBucket(), sourcePath, m.GetBucket(), targetPath)
	if err != nil {
		log.Error("CopyFile ObsCopyFile error. sourcePath=%s targetPath=%s err=%v", sourcePath, targetPath, err)
		return err
	}
	return nil
}

func (m *OBSHelper) CountDirSize(objectKey string) (int64, error) {
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

func (m *OBSHelper) DeleteDir(dirPath string) (remainingFiles bool, err error) {
	dirPath = m.TrimBucketPrefix(dirPath)
	log.Info("DeleteDir filePath=%s", dirPath)
	marker := ""
	for {
		res, err := m.GetAllObjectsUnderDirWithMarker(dirPath, marker)
		if err != nil {
			return true, err
		}
		if len(res.Objects) == 0 {
			break
		}
		marker = res.NextMarker
		objectsTodelete := make([]obs.ObjectToDelete, 0)
		for _, f := range res.Objects {
			objectsTodelete = append(objectsTodelete, obs.ObjectToDelete{
				Key: f.RelativePath,
			})
		}
		input := &obs.DeleteObjectsInput{}
		input.Bucket = m.GetBucket()
		input.Quiet = true
		input.Objects = objectsTodelete
		output, err := storage.ObsCli.DeleteObjects(input)
		if err != nil {
			remainingFiles = true
			log.Error("DeleteDir DeleteObjects error.input=%+v err=%v", input, err)
			if !res.IsTruncated {
				break
			}
			continue
		}
		if len(output.Errors) > 0 {
			remainingFiles = true
			jsonData, _ := json.Marshal(output.Errors)
			log.Error("Part of objects delete failed.%s", string(jsonData))
		}
		if !res.IsTruncated {
			break
		}
	}

	return remainingFiles, nil
}

func (m *OBSHelper) DeleteCollection(dirPath string) (remainingFiles bool, err error) {
	return m.DeleteDir(dirPath)
}

func (m *OBSHelper) DeleteFile(filePath string) error {
	filePath = m.TrimBucketPrefix(filePath)

	log.Info("DeleteFile filePath=%s", filePath)
	input := &obs.DeleteObjectInput{}
	input.Bucket = m.GetBucket()
	input.Key = filePath
	_, err := storage.ObsCli.DeleteObject(input)
	return err
}

func (m *OBSHelper) HasObject(path string) (bool, error) {
	path = m.TrimBucketPrefix(path)
	return storage.ObsHasObject(path)
}

const MAX_LIST_PARTS = 1000

func (m *OBSHelper) GetPartInfos(objectName string, uploadID string) ([]PartInfo, error) {
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

func (m *OBSHelper) NewMultiPartUpload(objectName string) (string, error) {
	objectName = m.TrimBucketPrefix(objectName)
	return storage.NewObsMultiPartUpload(objectName)
}

func (m *OBSHelper) GenMultiPartSignedUrl(objectName string, uploadId string, partNumber int, partSize int64) (string, error) {
	objectName = m.TrimBucketPrefix(objectName)
	return storage.ObsGenMultiPartSignedUrl(objectName, uploadId, partNumber)
}

func (m *OBSHelper) GenSignedUrl(objectKey string) (string, error) {
	objectKey = m.TrimBucketPrefix(objectKey)
	return storage.ObsGenSignedUrl(m.GetBucket(), objectKey)
}

func (m *OBSHelper) CompleteMultiPartUpload(objectName string, uploadID string, totalChunks int) (string, error) {
	objectName = m.TrimBucketPrefix(objectName)
	return "", storage.CompleteObsMultiPartUpload(objectName, uploadID, totalChunks)
}

func (m *OBSHelper) GetAllObjectByBucketAndPrefix(bucket string, prefix string) ([]storage.FileInfo, error) {

	return storage.GetAllObjectByBucketAndPrefix(bucket, prefix)

}

func (m *OBSHelper) GetFilesSize(Files []string) int64 {
	return storage.ObsGetFilesSize(m.GetBucket(), Files)

}
func (m *OBSHelper) RemoveObject(path string) error {
	return storage.ObsRemoveObject(m.GetBucket(), path)

}

func (m *OBSHelper) PutString(objectKey string, content string) error {

	return storage.PutStringToObs(m.GetBucket(), objectKey, content)

}

func (m *OBSHelper) GetOutputObjectKeyPrefix(jobName string, computeResource string, versionName string) string {
	objectkey := path.Join(m.GetJobDefaultObjectKeyPrefix(jobName), setting.OutPutPath, versionName) + "/"
	if computeResource != "NPU" {
		objectkey = path.Join(m.GetJobDefaultObjectKeyPrefix(jobName), cloudbrain.ModelMountPath, versionName) + "/"
	}
	return objectkey
}

func (m *OBSHelper) ReportSizeChanged(objectKey string) error {
	return nil
}
