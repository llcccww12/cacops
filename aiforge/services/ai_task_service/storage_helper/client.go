package storage_helper

import (
	"errors"
	"io"
	"path"
	"strings"
	"time"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/storage"
)

type UploaderConfig struct {
	Bucket   string
	Endpoint string
}

type ObjectResponse struct {
	ContentLength int64
	ContentType   string
	Body          io.ReadCloser
}

type ObjectMeta struct {
	ContentLength int64
	ContentType   string
	LastModified  time.Time
}

type ObjectListResponse struct {
	NextMarker  string
	Objects     []storage.FileInfo
	IsTruncated bool
}

type PartInfo struct {
	PartNumber   int
	ETag         string
	Size         int64
	LastModified time.Time
}

type UploadDirReq struct {
	LocalPath          string
	TargetObjectPrefix string
	ReqId              string
}

type MKDIRReq struct {
	ReqId       string
	ObjectKey   string
	Description string
}

// dirPath string, maxKeyArray ...int
type ListObjectsReq struct {
	Prefix    string
	MaxKey    int
	Marker    string
	ReqId     string
	Delimiter string
}

type GetObjectReq struct {
	ReqId string
	path  string
}

type StorageHelper interface {
	AllocateDatasetNamespace(name, prefix string) (string, error)
	UploadDir(codePath, objectKeyPrefix string) error
	UploadFile(objectKey string, r io.Reader) error
	GetRealPath(objectKey string) string
	GetBucket() string
	GetEndpoint() string
	GetDataId(path string) string
	GetJobDefaultObjectKeyPrefix(jobName string) string
	MKDIR(path string, description ...string) error
	GetOneLevelObjectsUnderDir(dirPath string, maxKeyArray ...int) ([]storage.FileInfo, error)
	GetOneLevelObjectsUnderDirWithMarker(dirPath string, marker string, maxKey int) (*ObjectListResponse, error)
	GetAllObjectsUnderDir(prefix string, maxKeyArray ...int) ([]storage.FileInfo, error)
	GetAllObjectsUnderDirWithMarker(prefix string, marker string, maxKeyArray ...int) (*ObjectListResponse, error)
	OpenFile(path string) (io.ReadCloser, error)
	GetObject(path string) (*ObjectResponse, error)
	GetObjectMeta(path string) (*ObjectMeta, error)
	GetSignedDownloadUrl(key string) (string, error)
	GetS3DownloadUrl(key string) string
	CopyDir(sourcePath, targetPath string, filterSuffix []string) error
	CopyFile(sourcePath, targetPath string) error
	DeleteDir(dirPath string) (remainingFiles bool, err error)
	DeleteFile(filePath string) error
	DeleteCollection(path string) (remainingFiles bool, err error)
	GetAllObjectByBucketAndPrefix(bucket string, prefix string) ([]storage.FileInfo, error)
	GetFilesSize(Files []string) int64
	RemoveObject(path string) error
	PutString(objectKey string, content string) error
	GetOutputObjectKeyPrefix(jobName string, computeResource string, versionName string) string
	HasObject(path string) (bool, error)
	GetPartInfos(objectName string, uploadID string) ([]PartInfo, error)
	NewMultiPartUpload(objectName string) (string, error)
	GenMultiPartSignedUrl(objectName string, uploadId string, partNumber int, partSize int64) (string, error)
	CompleteMultiPartUpload(objectName string, uploadID string, totalChunks int) (string, error)
	GenSignedUrl(objectKey string) (string, error)
	TrimBucketPrefix(path string) string
	CountDirSize(objectKey string) (int64, error)
	ReportSizeChanged(objectKey string) error
}

func SelectStorageHelperFromStorageType(storageType entity.StorageType, bucket ...string) StorageHelper {
	var bucketName string
	if len(bucket) > 0 {
		bucketName = bucket[0]
	}
	switch storageType {
	case entity.OBS:
		return &OBSHelper{
			Bucket: bucketName,
		}
	case entity.MINIO:
		return &MinioHelper{
			Bucket: bucketName,
		}
	case entity.URCHIN_V2:
		return &UrchinV2Helper{}
	}

	return nil
}

func SelectStorageHelperFromStorageIntType(storageType int) StorageHelper {
	switch storageType {
	case 1:
		return &OBSHelper{}
	case 0:
		return &MinioHelper{}
	}
	return nil
}
func GetStorageIntTypeFromStorageType(storageType entity.StorageType) int {
	switch storageType {
	case entity.OBS:
		return 1
	case entity.MINIO:
		return 0
	}
	return -1
}

func GetStorageTypeFromIntType(storageType int) entity.StorageType {
	switch storageType {
	case 1:
		return entity.OBS
	case 0:
		return entity.MINIO
	}
	return ""
}

func isMatchSuffix(fileName string, filterSuffix []string) bool {
	for _, s := range filterSuffix {
		if strings.HasSuffix(fileName, s) {
			return true
		}
	}
	return false

}

func CopyFileBetweenStorage(old, new StorageHelper, oldFilePath, newFilePath string) error {
	body, err := old.OpenFile(oldFilePath)
	if err != nil {
		return err
	}
	defer body.Close()
	return new.UploadFile(newFilePath, body)
}

func CopyDirBetweenStorage(old, new StorageHelper, oldDirPath, newDirPath string) error {
	files, err := old.GetAllObjectsUnderDir(oldDirPath)
	for _, file := range files {
		newFilePath := path.Join(newDirPath, file.FileName)
		err = CopyFileBetweenStorage(old, new, file.RelativePath, newFilePath)
		if err != nil {
			log.Error("transfer file between storage error. file=%+v err=%v", file, err)
			return err
		}
	}
	return nil
}

func Copy(old, new entity.StorageType, oldObjectKey, newObjectKey string) error {
	isDir := strings.HasSuffix(oldObjectKey, "/")
	isCrossStorage := !(old == new)

	oldHelper := SelectStorageHelperFromStorageType(old)
	newHelper := SelectStorageHelperFromStorageType(new)
	if oldHelper == nil || newHelper == nil {
		return errors.New("Transfer failed ")
	}
	if isDir && isCrossStorage {
		return CopyDirBetweenStorage(oldHelper, newHelper, oldObjectKey, newObjectKey)
	}
	if isDir && !isCrossStorage {
		return oldHelper.CopyDir(oldObjectKey, newObjectKey, []string{})
	}
	if !isDir && isCrossStorage {
		return CopyFileBetweenStorage(oldHelper, newHelper, oldObjectKey, newObjectKey)
	}
	if !isDir && !isCrossStorage {
		return oldHelper.CopyFile(oldObjectKey, newObjectKey)
	}
	return nil
}
