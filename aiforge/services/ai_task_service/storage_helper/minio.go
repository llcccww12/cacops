package storage_helper

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/url"
	"path"
	"sort"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/minio_ext"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
	"github.com/minio/minio-go"
)

type MinioHelper struct {
	Bucket string
}

func (m *MinioHelper) AllocateDatasetNamespace(name, prefix string) (string, error) {
	return prefix, nil
}

func (m *MinioHelper) UploadDir(codePath, objectKeyPrefix string) error {
	objectKeyPrefix = m.TrimBucketPrefix(objectKeyPrefix)
	return UploadDirToMinio(codePath, objectKeyPrefix, "")
}

func (m *MinioHelper) UploadFile(objectKey string, r io.Reader) error {
	objectKey = m.TrimBucketPrefix(objectKey)
	_, err := storage.Attachments.UploadContent(m.GetBucket(), objectKey, r)
	if err != nil {
		return err
	}
	return nil
}

func (m *MinioHelper) GetJobDefaultObjectKeyPrefix(jobName string) string {
	return path.Join(setting.CBCodePathPrefix, jobName)
}
func (m *MinioHelper) GetRealPath(objectKey string) string {
	objectKey = m.TrimBucketPrefix(objectKey)
	return setting.Attachment.Minio.RealPath + setting.Attachment.Minio.Bucket + "/" + strings.TrimPrefix(objectKey, "/")
}

func (m *MinioHelper) GetBucket() string {
	if m.Bucket != "" {
		return m.Bucket
	}
	return setting.Attachment.Minio.Bucket
}

func (m *MinioHelper) GetEndpoint() string {
	if setting.Attachment.Minio.InternalEndpoint != "" {
		return setting.Attachment.Minio.InternalEndpoint
	}
	return setting.Attachment.Minio.Endpoint
}

func (m *MinioHelper) GetDataId(path string) string {
	return ""
}

const README = "README"

func (m *MinioHelper) MKDIR(path string, description ...string) error {
	path = m.TrimBucketPrefix(path)
	//无法直接创建空文件夹，上传一个readme文件模拟
	path = strings.TrimSuffix(path, "/") + "/" + README
	val := "read me."
	if description != nil && len(description) > 0 {
		val = description[0]
	}
	_, err := storage.Attachments.UploadContentWithSize(m.GetBucket(), path, int64(len(val)), bytes.NewReader([]byte(val)))
	return err
}

func (m *MinioHelper) OpenFile(objectKey string) (io.ReadCloser, error) {
	objectKey = m.TrimBucketPrefix(objectKey)

	reader, _, err := storage.MinioCore.GetObject(m.GetBucket(), objectKey, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func (m *MinioHelper) GetObject(path string) (*ObjectResponse, error) {
	path = m.TrimBucketPrefix(path)
	reader, obj, err := storage.MinioCore.GetObject(m.GetBucket(), path, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return &ObjectResponse{
		ContentLength: obj.Size,
		ContentType:   obj.ContentType,
		Body:          reader,
	}, nil
}

func (m *MinioHelper) GetObjectMeta(path string) (*ObjectMeta, error) {
	path = m.TrimBucketPrefix(path)
	output, err := storage.MinioCore.StatObject(m.GetBucket(), path, minio.StatObjectOptions{})
	if err != nil {
		log.Error("OpenFile err. path=%s err=%v", path, err)
		return nil, err
	}
	return &ObjectMeta{
		ContentLength: output.Size,
		ContentType:   output.ContentType,
		LastModified:  output.LastModified,
	}, nil
}

func (m *MinioHelper) GetOneLevelObjectsUnderDir(dirPath string, maxKeyArray ...int) ([]storage.FileInfo, error) {
	maxKey := setting.OUTPUT_SHOW_MAX_KEY
	if len(maxKeyArray) > 0 {
		maxKey = maxKeyArray[0]
	}
	res, err := m.GetOneLevelObjectsUnderDirWithMarker(dirPath, "", maxKey)
	if err != nil {
		log.Error("GetOneLevelObjectsUnderDirWithMarker error: %v", err)
		return nil, err
	}
	return res.Objects, nil
}

func (m *MinioHelper) GetOneLevelObjectsUnderDirWithMarker(dirPath string, marker string, maxKey int) (*ObjectListResponse, error) {
	dirPath = m.TrimBucketPrefix(dirPath)

	if !strings.HasSuffix(dirPath, "/") {
		dirPath += "/"
	}
	if maxKey <= 0 {
		maxKey = setting.OUTPUT_SHOW_MAX_KEY
	}
	r, err := storage.MinioCore.ListObjectsV2(m.GetBucket(), dirPath, marker, false, "/", maxKey, "")
	if err != nil {
		return nil, err
	}
	list := r.Contents

	fileInfos := make([]storage.FileInfo, 0)
	prefixLen := len(dirPath)
	for _, val := range list {
		var fileName string
		if val.Key == dirPath {
			continue
		}
		fileName = val.Key[prefixLen:]
		fileInfo := storage.FileInfo{
			ModTime:      val.LastModified.Local().Format("2006-01-02 15:04:05"),
			FileName:     fileName,
			Size:         val.Size,
			IsDir:        false,
			RelativePath: dirPath + "/" + fileName,
			FullPath:     val.Key,
		}
		fileInfos = append(fileInfos, fileInfo)
	}
	for _, val := range r.CommonPrefixes {
		fileName := strings.TrimSuffix(strings.TrimPrefix(val.Prefix, dirPath), "/")
		fileInfo := storage.FileInfo{
			FileName: fileName,
			IsDir:    true,
		}
		fileInfos = append(fileInfos, fileInfo)
	}
	return &ObjectListResponse{
		Objects:     fileInfos,
		IsTruncated: r.IsTruncated,
		NextMarker:  r.NextContinuationToken,
	}, nil

}

func (m *MinioHelper) GetAllObjectsUnderDir(prefix string, maxKeyArray ...int) ([]storage.FileInfo, error) {
	prefix = m.TrimBucketPrefix(prefix)

	prefix = strings.TrimSuffix(prefix, "/") + "/"
	prefixLen := len(prefix)
	delimiter := ""
	marker := ""
	index := 1
	fileInfoList := storage.FileInfoList{}
	maxKey := setting.OUTPUT_DOWNLOAD_MAX_KEY
	if len(maxKeyArray) > 0 {
		maxKey = maxKeyArray[0]
	}
	for {
		output, err := storage.MinioCore.ListObjects(m.GetBucket(), prefix, marker, delimiter, maxKey)
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
				marker = output.NextMarker
			} else {
				break
			}
		} else {
			log.Info("list error." + err.Error())
			return nil, err
		}
	}
	sort.Sort(fileInfoList)
	return fileInfoList, nil
}

func (m *MinioHelper) GetAllObjectsUnderDirWithMarker(prefix string, marker string, maxKeyArray ...int) (*ObjectListResponse, error) {
	prefix = m.TrimBucketPrefix(prefix)

	prefix = strings.TrimSuffix(prefix, "/") + "/"
	prefixLen := len(prefix)
	delimiter := ""
	fileInfoList := storage.FileInfoList{}
	maxKey := setting.OUTPUT_DOWNLOAD_MAX_KEY
	if len(maxKeyArray) > 0 {
		maxKey = maxKeyArray[0]
	}

	output, err := storage.MinioCore.ListObjects(m.GetBucket(), prefix, marker, delimiter, maxKey)
	if err != nil {
		log.Error("GetAllObjectsUnderDirWithMarker ListObjects error.%v", err)
		return nil, err
	}

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

	sort.Sort(fileInfoList)
	return &ObjectListResponse{
		NextMarker:  output.NextMarker,
		Objects:     fileInfoList,
		IsTruncated: output.IsTruncated,
	}, nil
}

func (m *MinioHelper) GetSignedDownloadUrl(key string) (string, error) {
	key = m.TrimBucketPrefix(key)

	fileName := key[strings.LastIndex(key, "/"):]
	fileName = strings.TrimPrefix(fileName, "/")
	if fileName == "" {
		fileName = fmt.Sprint(time.Now().Unix())
	}
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+fileName+"\"")

	var preURL *url.URL
	preURL, err := storage.MinioCore.PresignedGetObject(m.GetBucket(), key, storage.PresignedGetUrlExpireTime, reqParams)
	if err != nil {
		return "", err
	}

	return preURL.String(), nil
}

func (m *MinioHelper) GetS3DownloadUrl(key string) string {
	return ""
}

func (m *MinioHelper) CountDirSize(objectKey string) (int64, error) {

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

func (m *MinioHelper) CopyDir(sourcePath, targetPath string, filterSuffix []string) error {
	log.Info("CopyByPath sourcePath=%s,targetPath=%s", sourcePath, targetPath)
	allFiles, _ := m.GetAllObjectsUnderDir(sourcePath)
	var fileNames []string
	for _, file := range allFiles {
		if isMatchSuffix(file.FileName, filterSuffix) {
			continue
		}
		fileNames = append(fileNames, file.FileName)
	}
	log.Info("Previous task all files", fileNames)
	if len(fileNames) == 0 {
		return nil
	}
	for _, file := range fileNames {
		srcObjectName := path.Join(sourcePath, file)
		destObjectName := path.Join(targetPath, file)
		_, err := storage.MinioCore.Client.StatObject(m.GetBucket(), srcObjectName, minio.StatObjectOptions{})
		if err != nil {
			log.Info("Get file error:" + err.Error())
		}
		_, err = storage.MinioCore.CopyObject(m.GetBucket(), srcObjectName, m.GetBucket(), destObjectName, map[string]string{})
		if err != nil {
			log.Error("CopyByPath MinioCopyFiles error. sourcePath=%s targetPath=%s err=%v", sourcePath, targetPath, err)
			return err
		}
	}
	return nil
}

func (m *MinioHelper) CopyFile(sourcePath, targetPath string) error {
	sourcePath = m.TrimBucketPrefix(sourcePath)
	targetPath = m.TrimBucketPrefix(targetPath)

	_, err := storage.MinioCore.CopyObject(m.GetBucket(), sourcePath, m.GetBucket(), targetPath, map[string]string{})
	if err != nil {
		log.Error("CopyFile MinioCore.CopyObject error. sourcePath=%s targetPath=%s err=%v", sourcePath, targetPath, err)
		return err
	}
	return nil
}

func (m *MinioHelper) DeleteDir(dirPath string) (remainingFiles bool, err error) {
	dirPath = m.TrimBucketPrefix(dirPath)

	log.Info("DeleteDir dirPath=%s", dirPath)
	err = storage.Attachments.DeleteDir(dirPath)
	return false, err
}

func (m *MinioHelper) DeleteCollection(dirPath string) (remainingFiles bool, err error) {
	return m.DeleteDir(dirPath)
}

func (m *MinioHelper) DeleteFile(filePath string) error {
	filePath = m.TrimBucketPrefix(filePath)

	log.Info("DeleteFile filePath=%s", filePath)
	err := storage.Attachments.Delete(filePath)
	return err
}

func (m *MinioHelper) HasObject(path string) (bool, error) {
	path = m.TrimBucketPrefix(path)

	return storage.Attachments.HasObject(path)
}

func (m *MinioHelper) GetPartInfos(objectName string, uploadID string) ([]PartInfo, error) {
	objectName = m.TrimBucketPrefix(objectName)
	result := make([]PartInfo, 0)
	parts, err := storage.GetPartDetailInfos(objectName, m.GetBucket(), uploadID)
	if err != nil {
		return nil, err
	}
	for _, part := range parts {
		result = append(result, PartInfo{
			PartNumber:   part.PartNumber,
			ETag:         part.ETag,
			Size:         part.Size,
			LastModified: part.LastModified,
		})
	}
	return result, nil
}

func (m *MinioHelper) NewMultiPartUpload(objectName string) (string, error) {
	objectName = m.TrimBucketPrefix(objectName)
	return storage.NewMultiPartUpload(objectName)
}

func (m *MinioHelper) GenMultiPartSignedUrl(objectName string, uploadId string, partNumber int, partSize int64) (string, error) {
	objectName = m.TrimBucketPrefix(objectName)
	if partSize > minio_ext.MinPartSize {
		return "", errors.New(fmt.Sprintf("chunk size(%d) is too big", partSize))
	}
	return storage.GenMultiPartSignedUrl(objectName, uploadId, partNumber, partSize)
}

func (m *MinioHelper) CompleteMultiPartUpload(objectName string, uploadID string, totalChunks int) (string, error) {
	objectName = m.TrimBucketPrefix(objectName)
	return storage.CompleteMultiPartUpload(objectName, uploadID, totalChunks)
}

func (m *MinioHelper) GenSignedUrl(objectKey string) (string, error) {
	objectKey = m.TrimBucketPrefix(objectKey)
	return storage.GenSignedUrl(m.GetBucket(), objectKey)
}

func (m *MinioHelper) TrimBucketPrefix(path string) string {
	return strings.TrimPrefix(path, m.GetBucket()+"/")
}

func (m *MinioHelper) GetAllObjectByBucketAndPrefix(bucket string, prefix string) ([]storage.FileInfo, error) {

	return storage.GetAllObjectByBucketAndPrefixMinio(bucket, prefix)

}

func (m *MinioHelper) GetFilesSize(Files []string) int64 {
	return storage.MinioGetFilesSize(m.GetBucket(), Files)

}
func (m *MinioHelper) RemoveObject(path string) error {
	return storage.Attachments.DeleteDir(path)

}
func (m *MinioHelper) PutString(objectKey string, content string) error {
	_, err := storage.Attachments.UploadContent(m.GetBucket(), objectKey, strings.NewReader(content))
	return err

}

func (m *MinioHelper) GetOutputObjectKeyPrefix(jobName string, computeResource string, versionName string) string {
	return m.GetJobDefaultObjectKeyPrefix(jobName) + "/model/"
}

func (m *MinioHelper) ReportSizeChanged(objectKey string) error {
	return nil
}
