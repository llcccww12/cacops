// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package storage

import (
	"errors"
	"io"
	"net/url"
	"path"
	"sort"
	"strconv"
	"strings"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/obs"
	"code.gitea.io/gitea/modules/setting"

	"github.com/unknwon/com"
)

type FileInfo struct {
	FileName        string `json:"FileName"`
	ModTime         string `json:"ModTime"`
	IsDir           bool   `json:"IsDir"`
	Size            int64  `json:"Size"`
	ParenDir        string `json:"ParenDir"`
	UUID            string `json:"UUID"`
	RelativePath    string `json:"RelativePath"`
	FullPath        string `json:"FullPath"`
	IsSupportPrivew bool   `json:"IsSupportPreview"`
}
type FileInfoList []FileInfo

const MAX_LIST_PARTS = 1000

func (ulist FileInfoList) Swap(i, j int) { ulist[i], ulist[j] = ulist[j], ulist[i] }
func (ulist FileInfoList) Len() int      { return len(ulist) }
func (ulist FileInfoList) Less(i, j int) bool {
	return strings.Compare(ulist[i].FileName, ulist[j].FileName) > 0
}

//check if has the object
func ObsHasObject(path string) (bool, error) {
	hasObject := false

	input := &obs.GetObjectMetadataInput{}
	input.Bucket = setting.Bucket
	input.Key = path
	_, err := ObsCli.GetObjectMetadata(input)
	if err == nil {
		hasObject = true
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			log.Error("GetObjectMetadata failed(%d): %s", obsError.StatusCode, obsError.Message)
		} else {
			log.Error("%v", err.Error())
		}
	}

	return hasObject, nil
}

func listAllParts(uuid, uploadID, key string) (output *obs.ListPartsOutput, err error) {
	output = &obs.ListPartsOutput{}
	partNumberMarker := 0
	for {
		temp, err := ObsCli.ListParts(&obs.ListPartsInput{
			Bucket:           setting.Bucket,
			Key:              key,
			UploadId:         uploadID,
			MaxParts:         MAX_LIST_PARTS,
			PartNumberMarker: partNumberMarker,
		})
		if err != nil {
			log.Error("ListParts failed:", err.Error())
			return output, err
		}

		partNumberMarker = temp.NextPartNumberMarker
		log.Info("uuid:%s, MaxParts:%d, PartNumberMarker:%d, NextPartNumberMarker:%d, len:%d", uuid, temp.MaxParts, temp.PartNumberMarker, temp.NextPartNumberMarker, len(temp.Parts))

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

	return output, nil
}

func GetObsPartInfos(objectName, uploadID string) (string, error) {
	key := objectName

	allParts, err := listAllParts(objectName, uploadID, key)
	if err != nil {
		log.Error("listAllParts failed: %v", err)
		return "", err
	}

	var chunks string
	for _, partInfo := range allParts.Parts {
		chunks += strconv.Itoa(partInfo.PartNumber) + "-" + partInfo.ETag + ","
	}

	return chunks, nil
}

func NewObsMultiPartUpload(objectName string) (string, error) {
	input := &obs.InitiateMultipartUploadInput{}
	input.Bucket = setting.Bucket
	input.Key = objectName

	output, err := ObsCli.InitiateMultipartUpload(input)
	if err != nil {
		log.Error("InitiateMultipartUpload failed:", err.Error())
		return "", err
	}

	return output.UploadId, nil
}

func CompleteObsMultiPartUpload(objectName, uploadID string, totalChunks int) error {
	input := &obs.CompleteMultipartUploadInput{}
	input.Bucket = setting.Bucket
	input.Key = objectName
	input.UploadId = uploadID

	allParts, err := listAllParts(objectName, uploadID, input.Key)
	if err != nil {
		log.Error("listAllParts failed: %v", err)
		return err
	}

	if len(allParts.Parts) != totalChunks {
		log.Error("listAllParts number(%d) is not equal the set total chunk number(%d)", len(allParts.Parts), totalChunks)
		return errors.New("the parts is not complete")
	}

	input.Parts = allParts.Parts

	output, err := ObsCli.CompleteMultipartUpload(input)
	if err != nil {
		log.Error("CompleteMultipartUpload failed:", err.Error())
		return err
	}

	log.Info("uuid:%s, RequestId:%s", objectName, output.RequestId)

	return nil
}

func ObsMultiPartUpload(objectName string, uploadId string, partNumber int, fileName string, putBody io.ReadCloser) error {
	input := &obs.UploadPartInput{}
	input.Bucket = setting.Bucket
	input.Key = objectName
	input.UploadId = uploadId
	input.PartNumber = partNumber
	input.Body = putBody
	output, err := ObsCli.UploadPart(input)
	if err == nil {
		log.Info("RequestId:%s\n", output.RequestId)
		log.Info("ETag:%s\n", output.ETag)
		return nil
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			log.Info(obsError.Code)
			log.Info(obsError.Message)
			return obsError
		} else {
			log.Error("error:", err.Error())
			return err
		}
	}

}

//delete all file under the dir path
func ObsRemoveObject(bucket string, path string) error {
	log.Info("Bucket=" + bucket + "  path=" + path)
	if len(path) == 0 {
		return errors.New("path canot be null.")
	}
	input := &obs.ListObjectsInput{}
	input.Bucket = bucket
	// 设置每页100个对象
	input.MaxKeys = 100
	input.Prefix = path
	index := 1
	log.Info("prefix=" + input.Prefix)
	for {
		output, err := ObsCli.ListObjects(input)
		if err == nil {
			log.Info("Page:%d\n", index)
			index++
			for _, val := range output.Contents {
				log.Info("delete obs file:" + val.Key)
				delObj := &obs.DeleteObjectInput{}
				delObj.Bucket = setting.Bucket
				delObj.Key = val.Key
				ObsCli.DeleteObject(delObj)
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
			return err
		}
	}
	return nil
}

func ObsDownloadAFile(bucket string, key string) (io.ReadCloser, error) {
	input := &obs.GetObjectInput{}
	input.Bucket = bucket
	input.Key = key
	output, err := ObsCli.GetObject(input)
	if err == nil {
		log.Info("StorageClass:%s, ETag:%s, ContentType:%s, ContentLength:%d, LastModified:%s\n",
			output.StorageClass, output.ETag, output.ContentType, output.ContentLength, output.LastModified)
		return output.Body, nil
	} else if obsError, ok := err.(obs.ObsError); ok {
		log.Error("Code:%s, Message:%s", obsError.Code, obsError.Message)
		return nil, obsError
	} else {
		return nil, err
	}
}

func ObsModelDownload(JobName string, fileName string) (io.ReadCloser, error) {
	input := &obs.GetObjectInput{}
	input.Bucket = setting.Bucket
	input.Key = strings.TrimPrefix(path.Join(setting.TrainJobModelPath, JobName, setting.OutPutPath, fileName), "/")
	output, err := ObsCli.GetObject(input)
	if err == nil {
		log.Info("StorageClass:%s, ETag:%s, ContentType:%s, ContentLength:%d, LastModified:%s\n",
			output.StorageClass, output.ETag, output.ContentType, output.ContentLength, output.LastModified)
		return output.Body, nil
	} else if obsError, ok := err.(obs.ObsError); ok {
		log.Error("Code:%s, Message:%s", obsError.Code, obsError.Message)
		return nil, obsError
	} else {
		return nil, err
	}
}

func ObsGetFilesSize(srcBucket string, Files []string) int64 {
	var fileTotalSize int64
	for _, file := range Files {
		log.Info("file=" + file)
		out, err := ObsCli.GetObjectMetadata(&obs.GetObjectMetadataInput{
			Bucket: srcBucket,
			Key:    file,
		})
		if err != nil {
			log.Info("Get File error, error=" + err.Error())
			continue
		}
		fileTotalSize += out.ContentLength
	}
	return fileTotalSize
}

func ObsCheckAndGetFileSize(srcBucket string, key string) (bool, int64) {
	out, err := ObsCli.GetObjectMetadata(&obs.GetObjectMetadataInput{
		Bucket: srcBucket,
		Key:    key,
	})
	if err != nil {
		log.Info("ObsCheckAndGetFilesSize error, error=%v", err)
		return false, 0
	}
	return true, out.ContentLength
}

func ObsCopyManyFile(srcBucket string, srcPath string, destBucket string, destPath string, Files []string) (int64, error) {

	var fileTotalSize int64
	srcPath = strings.TrimSuffix(srcPath, "/") + "/"
	destPath = strings.TrimSuffix(destPath, "/") + "/"

	for _, file := range Files {
		srcKey := srcPath + strings.TrimPrefix(file, "/")
		destKey := destPath + strings.TrimPrefix(file, "/")
		log.Info("srcKey=" + srcKey + " destKey=" + destKey)
		out, err := ObsCli.GetObjectMetadata(&obs.GetObjectMetadataInput{
			Bucket: srcBucket,
			Key:    srcKey,
		})
		if err != nil {
			log.Info("Get File error, error=" + err.Error())
			continue
		}
		ObsCopyFile(srcBucket, srcKey, destBucket, destKey)
		fileTotalSize += out.ContentLength
	}

	return fileTotalSize, nil
}

func ObsCopyAllFile(srcBucket string, srcPath string, destBucket string, destPath string) (int64, error) {
	input := &obs.ListObjectsInput{}
	input.Bucket = srcBucket
	// 设置每页100个对象
	input.MaxKeys = 100
	input.Prefix = srcPath
	index := 1
	length := len(srcPath)
	var fileTotalSize int64
	log.Info("prefix=" + input.Prefix)
	for {
		output, err := ObsCli.ListObjects(input)
		if err == nil {
			log.Info("Page:%d\n", index)
			index++
			for _, val := range output.Contents {
				destKey := destPath + val.Key[length:]
				ObsCopyFile(srcBucket, val.Key, destBucket, destKey)
				fileTotalSize += val.Size
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
			return 0, err
		}
	}
	return fileTotalSize, nil
}

func ObsCopyFile(srcBucket string, srcKeyName string, destBucket string, destKeyName string) error {
	input := &obs.CopyObjectInput{}
	input.Bucket = destBucket
	input.Key = destKeyName
	input.CopySourceBucket = srcBucket
	input.CopySourceKey = srcKeyName
	_, err := ObsCli.CopyObject(input)
	if err == nil {
		log.Info("copy success,destBuckName:%s, destkeyname:%s", destBucket, destKeyName)
	} else {
		log.Info("copy failed,,destBuckName:%s, destkeyname:%s", destBucket, destKeyName)
		if obsError, ok := err.(obs.ObsError); ok {
			log.Info(obsError.Code)
			log.Info(obsError.Message)
		}
		return err
	}
	return nil
}

func GetOneLevelAllObjectUnderDir(bucket string, prefixRootPath string, relativePath string) ([]FileInfo, error) {
	input := &obs.ListObjectsInput{}
	input.Bucket = bucket
	input.Prefix = prefixRootPath + relativePath
	if !strings.HasSuffix(input.Prefix, "/") {
		input.Prefix += "/"
	}
	fileInfos := make([]FileInfo, 0)
	prefixLen := len(input.Prefix)
	fileMap := make(map[string]bool, 0)
	index := 1
	for {
		output, err := ObsCli.ListObjects(input)
		if err == nil {
			log.Info("Page:%d\n  input.Prefix=v%", index, input.Prefix)
			log.Info("input.Prefix=" + input.Prefix)
			index++
			for _, val := range output.Contents {
				var isDir bool
				var fileName string
				log.Info("val.key=" + val.Key)
				if val.Key == input.Prefix {
					continue
				}
				fileName = val.Key[prefixLen:]
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
				input.Marker = output.NextMarker
			} else {
				break
			}
		} else {
			if obsError, ok := err.(obs.ObsError); ok {
				log.Error("Code:%s, Message:%s", obsError.Code, obsError.Message)
			}
			return nil, err
		}
	}
	return fileInfos, nil
}

func GetOneLevelObjectsUnderDir(bucket string, prefixRootPath string, relativePath string) ([]FileInfo, error) {
	input := &obs.ListObjectsInput{}
	input.Bucket = bucket
	input.Prefix = prefixRootPath + relativePath
	input.Delimiter = "/"
	if !strings.HasSuffix(input.Prefix, "/") {
		input.Prefix += "/"
	}
	fileInfos := make([]FileInfo, 0)
	prefixLen := len(input.Prefix)
	index := 1
	output, err := ObsCli.ListObjects(input)
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
		fileInfo := FileInfo{
			ModTime:  val.LastModified.Local().Format("2006-01-02 15:04:05"),
			FileName: fileName,
			Size:     val.Size,
			IsDir:    false,
			ParenDir: relativePath,
		}
		fileInfos = append(fileInfos, fileInfo)
	}
	for _, val := range output.CommonPrefixes {
		fileName := strings.TrimSuffix(strings.TrimPrefix(val, input.Prefix), "/")
		fileInfo := FileInfo{
			FileName: fileName,
			IsDir:    true,
			ParenDir: strings.TrimPrefix(val, prefixRootPath),
		}
		fileInfos = append(fileInfos, fileInfo)
	}
	return fileInfos, nil
}

func GetAllObjectByBucketAndPrefix(bucket string, prefix string) ([]FileInfo, error) {
	input := &obs.ListObjectsInput{}
	input.Bucket = bucket
	// 设置每页100个对象
	input.MaxKeys = 100
	input.Prefix = prefix
	index := 1
	fileInfoList := FileInfoList{}

	prefixLen := len(prefix)
	log.Info("full obs path:", input.Bucket+input.Prefix)
	log.Info("prefix=" + input.Prefix)
	for {
		output, err := ObsCli.ListObjects(input)
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
					ModTime:      val.LastModified.Format("2006-01-02 15:04:05"),
					FileName:     strings.TrimPrefix(val.Key[prefixLen:], "/"),
					Size:         val.Size,
					IsDir:        isDir,
					ParenDir:     "",
					RelativePath: val.Key,
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

func GetObsListObject(jobName, outPutPath, parentDir, versionName string) ([]FileInfo, error) {
	input := &obs.ListObjectsInput{}
	input.Bucket = setting.Bucket
	input.Prefix = strings.TrimPrefix(path.Join(setting.TrainJobModelPath, jobName, outPutPath, versionName, parentDir), "/")
	log.Info("obs prefix", input.Prefix)
	if !strings.HasSuffix(input.Prefix, "/") {
		input.Prefix += "/"
	}
	output, err := ObsCli.ListObjects(input)
	fileInfos := make([]FileInfo, 0)
	prefixLen := len(input.Prefix)
	fileMap := make(map[string]bool, 0)
	if err == nil {
		for _, val := range output.Contents {
			log.Info("val key=" + val.Key)
			var isDir bool
			var fileName string
			if val.Key == input.Prefix {
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
			ParenDir := parentDir
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
		sort.Slice(fileInfos, func(i, j int) bool {
			return fileInfos[i].ModTime > fileInfos[j].ModTime
		})
		return fileInfos, err
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			log.Error("Code:%s, Message:%s", obsError.Code, obsError.Message)
		}
		return nil, err
	}
}

func ObsGenMultiPartSignedUrl(objectName string, uploadId string, partNumber int) (string, error) {

	input := &obs.CreateSignedUrlInput{}
	input.Bucket = setting.Bucket
	input.Key = objectName
	input.Expires = 7 * 24 * 60 * 60
	input.Method = obs.HttpMethodPut

	input.QueryParams = map[string]string{
		"partNumber": com.ToStr(partNumber, 10),
		"uploadId":   uploadId,
		//"partSize":	com.ToStr(partSize,10),
	}

	output, err := ObsCli.CreateSignedUrl(input)
	if err != nil {
		log.Error("CreateSignedUrl failed:", err.Error())
		return "", err
	}

	return output.SignedUrl, nil
}

func ObsGenSignedUrl(bucketName, objectName string) (string, error) {

	input := &obs.CreateSignedUrlInput{}
	input.Bucket = bucketName
	input.Key = objectName
	input.Expires = 7 * 24 * 60 * 60
	input.Method = obs.HttpMethodPut

	input.QueryParams = map[string]string{}

	output, err := ObsCli.CreateSignedUrl(input)
	if err != nil {
		log.Error("CreateSignedUrl failed:", err.Error())
		return "", err
	}

	return output.SignedUrl, nil
}

func GetObsCreateSignedUrlByBucketAndKey(bucket, key string) (string, error) {
	input := &obs.CreateSignedUrlInput{}
	input.Bucket = bucket
	input.Key = key

	input.Expires = 7 * 24 * 60 * 60
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
	output, err := ObsCli.CreateSignedUrl(input)
	if err != nil {
		log.Error("CreateSignedUrl failed:", err.Error())
		return "", err
	}

	return output.SignedUrl, nil
}

func GetObsCreateSignedUrl(jobName, parentDir, fileName string) (string, error) {
	return GetObsCreateSignedUrlByBucketAndKey(setting.Bucket, strings.TrimPrefix(path.Join(setting.TrainJobModelPath, jobName, setting.OutPutPath, parentDir, fileName), "/"))
}

func ObsGetPreSignedUrl(objectName, fileName string) (string, error) {
	input := &obs.CreateSignedUrlInput{}
	input.Method = obs.HttpMethodGet
	input.Key = objectName
	input.Bucket = setting.Bucket
	input.Expires = 7 * 24 * 60 * 60

	fileName = url.PathEscape(fileName)
	reqParams := make(map[string]string)
	reqParams["response-content-disposition"] = "attachment; filename=\"" + fileName + "\""
	input.QueryParams = reqParams
	output, err := ObsCli.CreateSignedUrl(input)
	if err != nil {
		log.Error("CreateSignedUrl failed:", err.Error())
		return "", err
	}

	return output.SignedUrl, nil
}

func ObsCreateObject(path string) error {
	input := &obs.PutObjectInput{}
	input.Bucket = setting.Bucket
	input.Key = path

	_, err := ObsCli.PutObject(input)
	if err != nil {
		log.Error("PutObject failed:", err.Error())
		return err
	}

	return nil
}

func GetObsLogFileName(prefix string) ([]FileInfo, error) {
	input := &obs.ListObjectsInput{}
	input.Bucket = setting.Bucket
	input.Prefix = prefix

	output, err := ObsCli.ListObjects(input)
	if err != nil {
		log.Error("PutObject failed:", err.Error())
		return nil, err
	}
	if output == nil || len(output.Contents) == 0 {
		return nil, errors.New("obs log files not exist")
	}
	fileInfos := make([]FileInfo, 0)
	for _, val := range output.Contents {
		//result[num] = c.Key
		if strings.HasSuffix(val.Key, ".log") {
			log.Info("log fileName=" + val.Key)
			fileInfo := FileInfo{
				ModTime:  val.LastModified.Local().Format("2006-01-02 15:04:05"),
				FileName: val.Key[len(prefix)-3:], //加上  job
				Size:     val.Size,
				IsDir:    false,
				ParenDir: prefix[0 : len(prefix)-3],
			}
			fileInfos = append(fileInfos, fileInfo)
		}

	}
	return fileInfos, nil
}

func IsObjectExist4Obs(bucket, key string) (bool, error) {

	_, err := ObsCli.GetObjectMetadata(&obs.GetObjectMetadataInput{
		Bucket: bucket,
		Key:    key,
	})
	if err != nil {
		log.Error("GetObjectMetadata error.%v", err)
		return false, err
	}
	return true, nil
}

func PutStringToObs(bucket, key string, fileContent string) error {
	log.Info("PutStringToObs bucket=" + bucket + " key=" + key)
	input := &obs.PutObjectInput{}
	input.Bucket = bucket
	input.Key = key
	input.Body = strings.NewReader(fileContent)
	_, err := ObsCli.PutObject(input)
	if err != nil {
		if obsError, ok := err.(obs.ObsError); ok {
			log.Info("Message:%s\n", obsError.Message)
		}
	}
	return err
}

func PutReaderToObs(bucket, key string, reader io.Reader) error {
	log.Info("PutReaderToObs bucket=" + bucket + " key=" + key)
	input := &obs.PutObjectInput{}
	input.Bucket = bucket
	input.Key = key
	input.Body = reader
	_, err := ObsCli.PutObject(input)
	if err != nil {
		if obsError, ok := err.(obs.ObsError); ok {
			log.Info("Message:%s\n", obsError.Message)
		}
	}
	return err
}

func GetDirsList(bucket string, prefixRootPath string, prefix string, dirLevel int) ([]string, error) {
	input := &obs.ListObjectsInput{}
	input.Bucket = bucket
	input.Prefix = prefixRootPath + prefix
	if !strings.HasSuffix(input.Prefix, "/") {
		input.Prefix += "/"
	}
	prefixLen := len(input.Prefix)
	fileMap := make(map[string]bool, 0)
	index := 1
	for {
		output, err := ObsCli.ListObjects(input)
		if err == nil {
			log.Info("Page:%d\n  input.Prefix=v%", index, input.Prefix)
			log.Info("input.Prefix=" + input.Prefix)
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
				input.Marker = output.NextMarker
			} else {
				break
			}
		} else {
			if obsError, ok := err.(obs.ObsError); ok {
				log.Error("Code:%s, Message:%s", obsError.Code, obsError.Message)
			}
			return nil, err
		}
	}
	result := make([]string, 0)
	for k, _ := range fileMap {
		result = append(result, k)
	}
	return result, nil
}

func getPathFromPathArrays(files []string, currentLevel int) string {
	from := 0
	re := ""
	for {
		re += files[from]
		if from >= currentLevel {
			break
		}
		from = from + 1
		re += "/"
	}
	return re
}

func GetDirsSomeFile(bucket string, prefixRootPath string, prefix string, marker string, pageSize int) ([]FileInfo, string, error) {
	input := &obs.ListObjectsInput{}
	input.Bucket = bucket
	input.Prefix = prefixRootPath + prefix
	if !strings.HasSuffix(input.Prefix, "/") {
		input.Prefix += "/"
	}
	input.Marker = marker
	fileInfos := make([]FileInfo, 0)
	//prefixLen := len(input.Prefix)
	fileMap := make(map[string]bool, 0)
	index := 1
	for {
		output, err := ObsCli.ListObjects(input)
		if err == nil {
			log.Info("Page:%d  input.Prefix=v%", index, input.Prefix)
			log.Info("input.Prefix=" + input.Prefix)
			index++
			isBreak := false
			for _, val := range output.Contents {
				var isDir bool
				var fileName string
				//log.Info("val.Key=" + val.Key)
				if val.Key == input.Prefix {
					continue
				}
				fileNameRune := []rune(val.Key)
				prefixRune := []rune(input.Prefix)

				fileName = string(fileNameRune[len(prefixRune):])
				//log.Info("fileName=" + fileName)
				//val.Key[prefixLen:]
				files := strings.Split(fileName, "/")
				if fileMap[files[0]] {
					continue
				} else {
					fileMap[files[0]] = true
				}
				ParenDir := prefix
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
					log.Info("marker 1= " + marker)
					isBreak = true
					break
				}
			}
			if isBreak {
				break
			}
			if output.IsTruncated {
				input.Marker = output.NextMarker
				marker = output.NextMarker
				log.Info("marker 2= " + marker)
			} else {
				marker = ""
				log.Info("marker 3= " + marker)
				break
			}
		} else {
			if obsError, ok := err.(obs.ObsError); ok {
				log.Error("Code:%s, Message:%s", obsError.Code, obsError.Message)
			}
			return nil, "", err
		}
	}
	return fileInfos, marker, nil
}

func SetObsObjectMetaData(bucket, key string, metaData map[string]string) error {
	log.Info("SetObsObjectMetaData bucket=%s  key=%s metaData=%v", bucket, key, metaData)
	input := &obs.SetObjectMetadataInput{}
	input.Bucket = bucket
	input.Key = key
	input.Metadata = metaData
	_, err := ObsCli.SetObjectMetadata(input)
	if err != nil {
		if obsError, ok := err.(obs.ObsError); ok {
			log.Info("Message:%s\n", obsError.Message)
			if obsError.StatusCode == 404 {
				putInput := &obs.PutObjectInput{}
				putInput.Bucket = bucket
				putInput.Key = key
				ObsCli.PutObject(putInput)
				_, err = ObsCli.SetObjectMetadata(input)
			}
		}
	}
	return err
}

func GetObsObjectMetaData(bucket, key string) (map[string]string, error) {
	log.Info("GetObsObjectMetaData bucket=%s  key=%s", bucket, key)
	input := &obs.GetObjectMetadataInput{}
	input.Bucket = bucket
	input.Key = key
	output, err := ObsCli.GetObjectMetadata(input)
	if err != nil {
		if obsError, ok := err.(obs.ObsError); ok {
			log.Info("Message:%s\n", obsError.Message)
		}
		return nil, err
	}
	if output == nil {
		return map[string]string{}, nil
	}
	return output.Metadata, err
}
