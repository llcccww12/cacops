package repo

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"path"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
)

type DirRedisCache struct {
	DirList []storage.FileInfo
	Marker  string
}

func GetDirSomeFiles(ctx *context.Context) {
	uuid := ctx.Query("uuid")
	marker := ctx.Query("marker")
	pageSize := ctx.QueryInt("pageSize")
	prefix := ctx.Query("prefix")
	version := ctx.Query("version")
	cacheKey := uuid + prefix + marker + " " + fmt.Sprint(pageSize)
	//delimiter := ctx.Query("delimiter")
	var attach *models.Attachment
	var err error
	if version == "v2" {
		attach, err = getAttachmentV2(uuid)
	} else {
		attach, err = models.GetAttachmentByUUID(uuid)
	}

	if err != nil {
		log.Error("getAttachmentV2 error. %v", err)
		ctx.JSON(200, map[string]interface{}{
			"result_code": "-1",
			"msg":         err.Error(),
			"data":        "",
		})
		return
	}

	cacheResult, err := getAttachmentDirFromCache(cacheKey)
	if err == nil {
		if cacheResult.DirList != nil && len(cacheResult.DirList) > 0 {
			log.Info("load from cache.")
			//update redis status
			setAttachmentDirToCache(cacheKey, *cacheResult)
			ctx.JSON(200, map[string]interface{}{
				"result_code": "0",
				"data":        cacheResult.DirList,
				"marker":      cacheResult.Marker,
			})
			return
		}
	}
	if attach.Type == models.TypeCloudBrainOne {
		re, marker, err := storage.GetDirsSomeFileMinio(setting.Attachment.Minio.Bucket,
			attach.UnzipPath[0:len(attach.UnzipPath)-1], prefix, marker, pageSize)
		if err == nil {
			setAttachmentDirToCache(cacheKey, DirRedisCache{
				DirList: re,
				Marker:  marker,
			})
			ctx.JSON(200, map[string]interface{}{
				"result_code": "0",
				"data":        re,
				"marker":      marker,
			})
			return
		}
	} else if attach.Type == models.TypeCloudBrainTwo {
		re, marker, err := storage.GetDirsSomeFile(setting.Bucket,
			attach.UnzipPath[0:len(attach.UnzipPath)-1], prefix, marker, pageSize)
		if err == nil {
			setAttachmentDirToCache(cacheKey, DirRedisCache{
				DirList: re,
				Marker:  marker,
			})
			ctx.JSON(200, map[string]interface{}{
				"result_code": "0",
				"data":        re,
				"marker":      marker,
			})
			return
		}
	}
}

func getAttachmentV2(uuid string) (*models.Attachment, error) {
	dataset, err := models.GetDatasetRegistryByID(uuid)
	if err != nil {
		log.Error("GetDatasetV2ByID error. %v", err.Error())
		return nil, err
	}
	var cloudbrainType int
	if dataset.StorageType == string(entity.MINIO) {
		cloudbrainType = models.TypeCloudBrainOne
	} else if dataset.StorageType == string(entity.OBS) {
		cloudbrainType = models.TypeCloudBrainTwo
	} else {
		log.Error("GetDatasetV2ByID error. %v", "dataset.StorageType not support.")
		return nil, errors.New("dataset.StorageType not support.")
	}

	attach := &models.Attachment{
		UUID:      dataset.ID,
		UnzipPath: dataset.Path,
		Type:      cloudbrainType,
	}
	return attach, nil
}

func GetImageContent(ctx *context.Context) {
	uuid := ctx.Query("uuid")
	fileName := ctx.Query("filePath")
	storageType := ctx.QueryInt("type")
	if storageType == 0 {
		getMinioImageContent(uuid, fileName, ctx)
	} else {
		if storageType == 1 {
			getObsImageContent(uuid, fileName, ctx)
		}
	}
}

func GetTxtContent(ctx *context.Context) {
	uuid := ctx.Query("uuid")
	fileName := ctx.Query("filePath")
	storageType := ctx.QueryInt("type")
	if storageType == 0 {
		getMinioTxtContent(uuid, fileName, ctx)
	} else {
		if storageType == 1 {
			getObsTxtContent(uuid, fileName, ctx)
		}
	}
}

func getObsImageContent(uuid string, fileName string, ctx *context.Context) {
	isModel := ctx.QueryBool("isModel")
	version := ctx.Query("version")
	var attach *models.Attachment
	var err error
	if version == "v2" {
		attach, err = getAttachmentV2(uuid)
	} else {
		attach, err = models.GetAttachmentByUUID(uuid)
	}
	objectName := ""
	if err == nil {
		objectName = path.Join(attach.UnzipPath, fileName)
	} else {
		log.Info("download error. not found objectname.uuid=" + uuid)
	}
	//objectName := strings.TrimPrefix(path.Join(setting.BasePath+models.AttachmentRelativePath(uuid)+uuid, fileName), "/")
	if isModel {
		objectName = strings.TrimPrefix(path.Join(Model_prefix, path.Join(uuid[0:1], uuid[1:2], uuid, fileName)), "/")
	}
	log.Info("obs objectName=" + objectName)
	body, err := storage.ObsDownloadAFile(setting.Bucket, objectName)
	if err != nil {
		log.Info("upload error.")
	} else {
		defer body.Close()
		ctx.Resp.Header().Set("Content-Type", "image/jpg;charset=utf-8")
		p := make([]byte, 1024)
		var readErr error
		var readCount int
		// 读取对象内容
		for {
			readCount, readErr = body.Read(p)
			if readCount > 0 {
				ctx.Resp.Write(p[:readCount])
			}
			if readErr != nil {
				break
			}
		}
	}
}

func getObsTxtContent(uuid string, fileName string, ctx *context.Context) {
	isModel := ctx.QueryBool("isModel")
	version := ctx.Query("version")
	var attach *models.Attachment
	var err error
	if version == "v2" {
		attach, err = getAttachmentV2(uuid)
	} else {
		attach, err = models.GetAttachmentByUUID(uuid)
	}
	objectName := ""
	if err == nil {
		objectName = path.Join(attach.UnzipPath, fileName)
	} else {
		log.Info("download error. not found objectname.uuid=" + uuid)
	}
	//objectName := strings.TrimPrefix(path.Join(setting.BasePath+models.AttachmentRelativePath(uuid)+uuid, fileName), "/")
	if isModel {
		objectName = strings.TrimPrefix(path.Join(Model_prefix, path.Join(uuid[0:1], uuid[1:2], uuid, fileName)), "/")
	}
	log.Info("obs objectName=" + objectName)
	body, err := storage.ObsDownloadAFile(setting.Bucket, objectName)
	if err != nil {
		log.Info("download error.")
		ctx.JSON(200, map[string]interface{}{
			"result_code": "-1",
			"msg":         "Read file error.",
			"data":        "",
		})
		return
	} else {
		readS3Body(body, ctx)
	}
}

func readS3Body(body io.ReadCloser, ctx *context.Context) {
	defer body.Close()
	result := make([]string, 0)
	r := bufio.NewReader(body)
	size := 0
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			result = append(result, toUTF8(line))
			size += len(line)
			if !toDealLine(result, line, size, ctx) {
				return
			}
			break
		}
		if err != nil {
			log.Info("read error,err=" + err.Error())
			break
		}
		result = append(result, toUTF8(line))
		size += len(line)
		if !toDealLine(result, line, size, ctx) {
			return
		}
	}
	ctx.JSON(200, map[string]interface{}{
		"result_code": "0",
		"msg":         "",
		"data":        result,
	})
}

func toDealLine(result []string, line string, size int, ctx *context.Context) bool {
	log.Info("line count=" + fmt.Sprint(len(result)) + " size=" + fmt.Sprint(size))
	if len(result) > 2000 || size > 2*1024*1024 {
		ctx.JSON(200, map[string]interface{}{
			"result_code": "-1",
			"msg":         ctx.Tr("repo.attachmentfilesizetobig"),
			"data":        "",
		})
		return false
	}
	return true
}

func toUTF8(line string) string {
	lineBytes := []byte(line)
	if isUtf8(lineBytes) {
		return line
	}
	if isGBK(lineBytes) {
		log.Info("this is gbk")
		gbkBytes, err := simplifiedchinese.GBK.NewDecoder().Bytes(lineBytes) //gbk 转 utf-8
		if err == nil {
			return string(gbkBytes)
		}
	}
	return line
}

func toGBK(line string) string {
	lineBytes := []byte(line)
	if !isGBK(lineBytes) {
		gbkBytes, err := simplifiedchinese.GBK.NewEncoder().Bytes(lineBytes) //utf-8 转 gbk
		if err == nil {
			return string(gbkBytes)
		}
	}
	return line
}

func isGBK(data []byte) bool {
	length := len(data)
	var i int = 0
	for i < length {
		//fmt.Printf("for %x\n", data[i])
		if data[i] <= 0xff {
			//编码小于等于127,只有一个字节的编码，兼容ASCII吗
			i++
			continue
		} else {
			//大于127的使用双字节编码
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func getMinioImageContent(uuid string, fileName string, ctx *context.Context) {
	version := ctx.Query("version")
	var attach *models.Attachment
	var err error
	if version == "v2" {
		attach, err = getAttachmentV2(uuid)
	} else {
		attach, err = models.GetAttachmentByUUID(uuid)
	}
	objectName := ""
	if err == nil {
		objectName = path.Join(attach.UnzipPath, fileName)
	} else {
		log.Info("download error. not found objectname.uuid=" + uuid)
	}
	//objectName := strings.TrimPrefix(path.Join(setting.Attachment.Minio.BasePath+models.AttachmentRelativePath(uuid)+uuid, fileName), "/")
	log.Info("minio objectName=" + objectName)
	body, err := storage.Attachments.DownloadAFile(setting.Attachment.Minio.Bucket, objectName)
	if err != nil {
		log.Info("download error.")
	} else {
		defer body.Close()
		ctx.Resp.Header().Set("Content-Type", "image/jpg;charset=utf-8")
		p := make([]byte, 1024)
		var readErr error
		var readCount int
		// 读取对象内容
		for {
			readCount, readErr = body.Read(p)
			if readCount > 0 {
				ctx.Resp.Write(p[:readCount])
			}
			if readErr != nil {
				break
			}
		}
	}
}

func getMinioTxtContent(uuid string, fileName string, ctx *context.Context) {
	version := ctx.Query("version")
	var attach *models.Attachment
	var err error
	if version == "v2" {
		attach, err = getAttachmentV2(uuid)
	} else {
		attach, err = models.GetAttachmentByUUID(uuid)
	}
	objectName := ""
	if err == nil {
		objectName = path.Join(attach.UnzipPath, fileName)
	} else {
		log.Info("download error. not found objectname.uuid=" + uuid)
	}
	//objectName := strings.TrimPrefix(path.Join(setting.Attachment.Minio.BasePath+models.AttachmentRelativePath(uuid)+uuid, fileName), "/")
	log.Info("minio objectName=" + objectName)
	body, err := storage.Attachments.DownloadAFile(setting.Attachment.Minio.Bucket, objectName)
	if err != nil {
		log.Info("download error.")
		ctx.JSON(200, map[string]interface{}{
			"result_code": "-1",
			"msg":         "Read file error.",
			"data":        "",
		})
		return
	} else {
		readS3Body(body, ctx)
	}
}

func preNUm(data byte) int {
	str := fmt.Sprintf("%b", data)
	var i int = 0
	for i < len(str) {
		if str[i] != '1' {
			break
		}
		i++
	}
	return i
}

func isUtf8(data []byte) bool {
	for i := 0; i < len(data); {
		if data[i]&0x80 == 0x00 {
			// 0XXX_XXXX
			i++
			continue
		} else if num := preNUm(data[i]); num > 2 {
			// 110X_XXXX 10XX_XXXX
			// 1110_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_0XXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_10XX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_110X 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// preNUm() 返回首个字节的8个bits中首个0bit前面1bit的个数，该数量也是该字符所使用的字节数
			i++
			for j := 0; j < num-1; j++ {
				//判断后面的 num - 1 个字节是不是都是10开头
				if data[i]&0xc0 != 0x80 {
					return false
				}
				i++
			}
		} else {
			//其他情况说明不是utf-8
			return false
		}
	}
	return true
}

func setAttachmentDirToCache(msgKey string, dirList DirRedisCache) {
	msgMapJson, _ := json.Marshal(dirList)
	redisValue := string(msgMapJson)
	log.Info("set redis key=" + msgKey + " value=" + redisValue)
	re, err := redis_client.Setex(msgKey, redisValue, 30*24*3600*time.Second)
	if err == nil {
		log.Info("re =" + fmt.Sprint(re))
	} else {
		log.Info("set redis error:" + err.Error())
	}
}

func getAttachmentDirFromCache(msgKey string) (*DirRedisCache, error) {
	valueStr, err := redis_client.Get(msgKey)
	//msgMap := make(map[string]string, 0)
	fileInfos := &DirRedisCache{}
	if err == nil {
		if valueStr != "" {
			err1 := json.Unmarshal([]byte(valueStr), fileInfos)
			if err1 != nil {
				log.Info("unmarshal json failed. " + err1.Error())
				return nil, err1
			}
		} else {
			return nil, errors.New("cache is empty.")
		}
	} else {
		log.Info("Failed to load from reids. " + err.Error())
		return nil, err
	}
	return fileInfos, nil
}
