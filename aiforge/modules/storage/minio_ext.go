package storage

import (
	"encoding/xml"
	"errors"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/minio_ext"
	"code.gitea.io/gitea/modules/setting"

	miniov6 "github.com/minio/minio-go/v6"
)

const (
	PresignedUploadPartUrlExpireTime = time.Hour * 24 * 7
)

type ComplPart struct {
	PartNumber int    `json:"partNumber"`
	ETag       string `json:"eTag"`
}

type CompleteParts struct {
	Data []ComplPart `json:"completedParts"`
}

// completedParts is a collection of parts sortable by their part numbers.
// used for sorting the uploaded parts before completing the multipart request.
type completedParts []miniov6.CompletePart

func (a completedParts) Len() int           { return len(a) }
func (a completedParts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a completedParts) Less(i, j int) bool { return a[i].PartNumber < a[j].PartNumber }

// completeMultipartUpload container for completing multipart upload.
type completeMultipartUpload struct {
	XMLName xml.Name               `xml:"http://s3.amazonaws.com/doc/2006-03-01/ CompleteMultipartUpload" json:"-"`
	Parts   []miniov6.CompletePart `xml:"Part"`
}

var (
	adminClient *minio_ext.Client = nil
	coreClient  *miniov6.Core     = nil
)

var mutex *sync.Mutex

func init() {
	mutex = new(sync.Mutex)
}

func getClients() (*minio_ext.Client, *miniov6.Core, error) {
	var client *minio_ext.Client
	var core *miniov6.Core
	mutex.Lock()

	defer mutex.Unlock()

	if nil != adminClient && nil != coreClient {
		client = adminClient
		core = coreClient
		return client, core, nil
	}

	var err error
	minio := setting.Attachment.Minio
	if nil == adminClient {
		adminClient, err = minio_ext.New(
			minio.Endpoint,
			minio.AccessKeyID,
			minio.SecretAccessKey,
			minio.UseSSL,
		)
		if nil != err {
			return nil, nil, err
		}
	}

	client = adminClient

	if nil == coreClient {
		coreClient, err = miniov6.NewCore(
			minio.Endpoint,
			minio.AccessKeyID,
			minio.SecretAccessKey,
			minio.UseSSL,
		)
		if nil != err {
			return nil, nil, err
		}
	}

	core = coreClient

	return client, core, nil
}

func GenMultiPartSignedUrl(objectName string, uploadId string, partNumber int, partSize int64) (string, error) {
	minioClient, _, err := getClients()
	if err != nil {
		log.Error("getClients failed:", err.Error())
		return "", err
	}

	minio := setting.Attachment.Minio
	bucketName := minio.Bucket

	return minioClient.GenUploadPartSignedUrl(uploadId, bucketName, objectName, partNumber, partSize, PresignedUploadPartUrlExpireTime, setting.Attachment.Minio.Location)
}

func GenSignedUrl(bucketName, objectName string) (string, error) {
	minioClient, _, err := getClients()
	if err != nil {
		log.Error("getClients failed:", err.Error())
		return "", err
	}

	return minioClient.GenUploadSignedUrl(bucketName, objectName, PresignedUploadPartUrlExpireTime, setting.Attachment.Minio.Location)
}

func GetAllObjectByBucketAndPrefixMinio(bucket string, prefix string) ([]FileInfo, error) {
	_, core, err := getClients()
	if err != nil {
		log.Error("getClients failed:", err.Error())
		return nil, err
	}
	prefixLen := len(prefix)
	delimiter := ""
	marker := ""
	index := 1
	fileInfoList := FileInfoList{}
	for {
		output, err := core.ListObjects(bucket, prefix, marker, delimiter, 1000)
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
				fileInfo := FileInfo{
					ModTime:  val.LastModified.Format("2006-01-02 15:04:05"),
					FileName: val.Key[prefixLen:],
					Size:     val.Size,
					IsDir:    isDir,
					ParenDir: "",
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

func GetOneLevelAllObjectUnderDirMinio(bucket string, prefixRootPath string, relativePath string) ([]FileInfo, error) {
	_, core, err := getClients()
	if err != nil {
		log.Error("getClients failed:", err.Error())
		return nil, err
	}

	Prefix := prefixRootPath + relativePath
	if !strings.HasSuffix(Prefix, "/") {
		Prefix += "/"
	}
	log.Info("bucket=" + bucket + " Prefix=" + Prefix)
	marker := ""
	fileInfos := make([]FileInfo, 0)
	prefixLen := len(Prefix)
	fileMap := make(map[string]bool, 0)
	index := 1
	for {
		output, err := core.ListObjects(bucket, Prefix, marker, "", 1000)
		if err == nil {
			log.Info("Page:%d\n", index)
			index++
			for _, val := range output.Contents {
				log.Info("val key=" + val.Key)
				var isDir bool
				var fileName string
				if val.Key == Prefix {
					continue
				}
				fileName = val.Key[prefixLen:]
				log.Info("fileName =" + fileName)
				files := strings.Split(fileName, "/")
				if fileMap[files[0]] {
					continue
				} else {
					fileMap[files[0]] = true
				}
				ParenDir := relativePath
				fileName = files[0]
				if len(files) > 1 {
					isDir = true
					ParenDir += fileName + "/"
				} else {
					isDir = false
				}
				fileInfo := FileInfo{
					ModTime:  val.LastModified.Local().Format("2006-01-02 15:04:05"),
					FileName: fileName,
					Size:     val.Size,
					IsDir:    isDir,
					ParenDir: ParenDir,
				}
				fileInfos = append(fileInfos, fileInfo)
			}
			if output.IsTruncated {
				marker = output.NextMarker
			} else {
				break
			}

		} else {

			log.Error("Message:%s", err.Error())

			return nil, err
		}
	}
	return fileInfos, err
}

func MinioGetFilesSize(bucketName string, Files []string) int64 {
	_, core, err := getClients()
	var fileTotalSize int64
	fileTotalSize = 0
	if err != nil {
		log.Error("getClients failed:", err.Error())
		return fileTotalSize
	}
	for _, file := range Files {
		log.Info("file=" + file)
		meta, err := core.StatObject(bucketName, file, miniov6.StatObjectOptions{})
		if err != nil {
			log.Info("Get file error:" + err.Error())
		}
		fileTotalSize += meta.Size
	}
	return fileTotalSize
}

func MinioCopyFiles(bucketName string, srcPath string, destPath string, Files []string) (int64, error) {
	_, core, err := getClients()
	var fileTotalSize int64
	fileTotalSize = 0
	if err != nil {
		log.Error("getClients failed:", err.Error())
		return fileTotalSize, err
	}

	for _, file := range Files {
		srcObjectName := srcPath + file
		destObjectName := destPath + file
		log.Info("srcObjectName=" + srcObjectName + " destObjectName=" + destObjectName)
		meta, err := core.StatObject(bucketName, srcObjectName, miniov6.StatObjectOptions{})
		if err != nil {
			log.Info("Get file error:" + err.Error())
		}
		core.CopyObject(bucketName, srcObjectName, bucketName, destObjectName, meta.UserMetadata)
		fileTotalSize += meta.Size
	}

	return fileTotalSize, nil
}

func MinioCopyAFile(srcBucketName, srcObjectName, destBucketName, destObjectName string) (int64, error) {
	_, core, err := getClients()
	var fileTotalSize int64
	fileTotalSize = 0
	if err != nil {
		log.Error("getClients failed:", err.Error())
		return fileTotalSize, err
	}
	meta, err := core.StatObject(srcBucketName, srcObjectName, miniov6.StatObjectOptions{})
	if err != nil {
		log.Info("Get file error:" + err.Error())
	}
	core.CopyObject(srcBucketName, srcObjectName, destBucketName, destObjectName, meta.UserMetadata)
	fileTotalSize = meta.Size
	return fileTotalSize, nil
}

func MinioPathCopy(bucketName string, srcPath string, destPath string) (int64, error) {
	_, core, err := getClients()
	var fileTotalSize int64
	fileTotalSize = 0
	if err != nil {
		log.Error("getClients failed:", err.Error())
		return fileTotalSize, err
	}
	delimiter := ""
	marker := ""
	for {
		output, err := core.ListObjects(bucketName, srcPath, marker, delimiter, 1000)
		if err == nil {
			for _, val := range output.Contents {
				srcObjectName := val.Key
				destObjectName := destPath + srcObjectName[len(srcPath):]
				log.Info("srcObjectName=" + srcObjectName + " destObjectName=" + destObjectName)
				core.CopyObject(bucketName, srcObjectName, bucketName, destObjectName, val.UserMetadata)
				fileTotalSize += val.Size
			}
			if output.IsTruncated {
				marker = output.NextMarker
			} else {
				break
			}
		} else {
			log.Info("list error." + err.Error())
			return 0, err
		}
	}
	return fileTotalSize, nil
}

func NewMultiPartUpload(objectName string) (string, error) {
	_, core, err := getClients()
	if err != nil {
		log.Error("getClients failed:", err.Error())
		return "", err
	}

	minio := setting.Attachment.Minio
	bucketName := minio.Bucket

	return core.NewMultipartUpload(bucketName, objectName, miniov6.PutObjectOptions{})
}

func CompleteMultiPartUpload(objectName string, uploadID string, totalChunks int) (string, error) {
	client, core, err := getClients()
	if err != nil {
		log.Error("getClients failed:", err.Error())
		return "", err
	}

	minio := setting.Attachment.Minio
	bucketName := minio.Bucket

	log.Info("bucketName=" + bucketName + " objectName=" + objectName + " uploadID=" + uploadID)
	partInfos, err := client.ListObjectParts(bucketName, objectName, uploadID)
	if err != nil {
		log.Error("ListObjectParts failed:", err.Error())
		return "", err
	}

	if len(partInfos) != totalChunks {
		log.Error("ListObjectParts number(%d) is not equal the set total chunk number(%d)", len(partInfos), totalChunks)
		return "", errors.New("the parts is not complete")
	}

	var complMultipartUpload completeMultipartUpload
	for _, partInfo := range partInfos {
		complMultipartUpload.Parts = append(complMultipartUpload.Parts, miniov6.CompletePart{
			PartNumber: partInfo.PartNumber,
			ETag:       partInfo.ETag,
		})
	}

	// Sort all completed parts.
	sort.Sort(completedParts(complMultipartUpload.Parts))

	return core.CompleteMultipartUpload(bucketName, objectName, uploadID, complMultipartUpload.Parts)
}

func GetPartInfos(objectName string, uploadID string) (string, error) {
	minioClient, _, err := getClients()
	if err != nil {
		log.Error("getClients failed:", err.Error())
		return "", err
	}

	minio := setting.Attachment.Minio
	bucketName := minio.Bucket

	partInfos, err := minioClient.ListObjectParts(bucketName, objectName, uploadID)
	if err != nil {
		log.Error("ListObjectParts failed:", err.Error())
		return "", err
	}

	var chunks string
	for _, partInfo := range partInfos {
		chunks += strconv.Itoa(partInfo.PartNumber) + "-" + partInfo.ETag + ","
	}

	return chunks, nil
}

func GetPartDetailInfos(objectName, bucketName, uploadID string) (map[int]minio_ext.ObjectPart, error) {
	minioClient, _, err := getClients()
	if err != nil {
		log.Error("getClients failed:", err.Error())
		return nil, err
	}

	partInfos, err := minioClient.ListObjectParts(bucketName, objectName, uploadID)
	if err != nil {
		log.Error("ListObjectParts failed:", err.Error())
		return nil, err
	}

	return partInfos, nil
}

func IsObjectExist4Minio(bucket, objectName string) (bool, error) {
	_, core, err := getClients()
	if err != nil {
		log.Error("getClients failed:", err.Error())
		return false, err
	}

	_, err = core.StatObject(bucket, objectName, miniov6.StatObjectOptions{})
	if err != nil {
		log.Error("GetObjectMetadata error.%v", err)
		return false, err
	}

	return true, nil
}

func MinioCheckAndGetFileSize(srcBucket string, key string) (bool, int64) {
	_, core, err := getClients()
	if err != nil {
		log.Error("getClients failed:", err.Error())
		return false, 0
	}

	meta, err := core.StatObject(srcBucket, key, miniov6.StatObjectOptions{})
	if err != nil {
		log.Info("MinioCheckAndGetFileSize error, error=%v", err)
		return false, 0
	}
	return true, meta.Size
}

func GetDirsListMinio(bucket string, prefixRootPath string, prefix string, dirLevel int) ([]string, error) {
	_, core, err := getClients()
	if err != nil {
		log.Error("getClients failed:", err.Error())
		return nil, err
	}
	minioPrefix := prefixRootPath + prefix
	if !strings.HasSuffix(minioPrefix, "/") {
		minioPrefix += "/"
	}
	prefixLen := len(minioPrefix)
	delimiter := ""
	marker := ""
	index := 1
	fileMap := make(map[string]bool, 0)
	log.Info("buckent=" + bucket + " minioPrefix=" + minioPrefix)
	for {
		output, err := core.ListObjects(bucket, minioPrefix, marker, delimiter, 1000)
		if err == nil {
			log.Info("Page:%d\n", index)
			index++
			for _, val := range output.Contents {
				fileName := val.Key[prefixLen:]
				files := strings.Split(fileName, "/")
				currentLevel := 0
				for {
					if currentLevel >= dirLevel || currentLevel >= (len(files)-1) {
						break
					}
					path := getPathFromPathArrays(files, currentLevel)
					if !fileMap[path] {
						fileMap[path] = true
					}
					currentLevel = currentLevel + 1
				}
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
	result := make([]string, 0)
	for k, _ := range fileMap {
		result = append(result, k)
	}
	return result, nil
}

func GetDirsSomeFileMinio(bucket string, prefixRootPath string, relativePath string, marker string, pageSize int) ([]FileInfo, string, error) {
	_, core, err := getClients()
	if err != nil {
		log.Error("getClients failed:", err.Error())
		return nil, "", err
	}

	Prefix := prefixRootPath + relativePath
	if !strings.HasSuffix(Prefix, "/") {
		Prefix += "/"
	}
	log.Info("bucket=" + bucket + " Prefix=" + Prefix)
	fileInfos := make([]FileInfo, 0)

	fileMap := make(map[string]bool, 0)
	index := 1
	for {
		output, err := core.ListObjects(bucket, Prefix, marker, "", 1000)
		if err == nil {
			log.Info("Page:%d\n", index)
			index++
			isBreak := false
			for _, val := range output.Contents {
				//log.Info("val key=" + val.Key)
				var isDir bool
				var fileName string
				if val.Key == Prefix {
					continue
				}
				fileNameRune := []rune(val.Key)
				prefixRune := []rune(Prefix)

				fileName = string(fileNameRune[len(prefixRune):])
				files := strings.Split(fileName, "/")
				if fileMap[files[0]] {
					continue
				} else {
					fileMap[files[0]] = true
				}
				ParenDir := relativePath
				fileName = files[0]
				if len(files) > 1 {
					isDir = true
					ParenDir += fileName + "/"
				} else {
					isDir = false
				}
				fileInfo := FileInfo{
					ModTime:  val.LastModified.Local().Format("2006-01-02 15:04:05"),
					FileName: fileName,
					Size:     val.Size,
					IsDir:    isDir,
					ParenDir: ParenDir,
				}
				fileInfos = append(fileInfos, fileInfo)
				if len(fileInfos) >= pageSize {
					marker = val.Key
					isBreak = true
					break
				}
			}
			if isBreak {
				break
			}
			if output.IsTruncated {
				marker = output.NextMarker
			} else {
				marker = ""
				break
			}
		} else {
			log.Error("Message:%s", err.Error())
			return nil, "", err
		}
	}
	return fileInfos, marker, err
}
