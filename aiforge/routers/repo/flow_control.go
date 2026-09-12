package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/setting"
)

const (
	REDIS_FLOW_ATTACHMENT_KEY       = "flow_attachment_key"
	REDIS_FLOW_MODEL_ATTACHMENT_KEY = "flow_model_attachment_key"
)

var mutex *sync.RWMutex = new(sync.RWMutex)
var modelMutex *sync.Mutex = new(sync.Mutex)

func CheckFlowForDataset(ctx *context.Context) error {
	if ctx.User == nil {
		return errors.New("User not login.")
	}
	log.Info("start to check flow for upload dataset file.")
	fileName := ctx.Query("file_name")
	currentTimeNow := time.Now()
	currentLongTime := currentTimeNow.Unix()
	last24Hour := currentTimeNow.AddDate(0, 0, -1).Unix()
	filechunks, err := models.GetFileChunksByUserId(ctx.User.ID, last24Hour, true)
	if err == nil {
		if len(filechunks) > setting.FLOW_CONTROL.ATTACHEMENT_NUM_A_USER_LAST24HOUR {
			log.Info("A single user cannot upload more than " + fmt.Sprint(setting.FLOW_CONTROL.ATTACHEMENT_NUM_A_USER_LAST24HOUR) + " files within the last 24 hours. so " + fileName + " is rejected. user id=" + fmt.Sprint(ctx.User.ID))
			return errors.New("A single user cannot upload more than " + fmt.Sprint(setting.FLOW_CONTROL.ATTACHEMENT_NUM_A_USER_LAST24HOUR) + " files within the last 24 hours.")
		}
		var totalSize int64
		totalSize += ctx.QueryInt64("size")
		concurrentUpload := 0
		for _, file := range filechunks {
			totalSize += file.Size
			if (currentLongTime - int64(file.CreatedUnix)) < 10*60 {
				log.Info("the file " + file.Md5 + " in 10min upload." + file.CreatedUnix.Format("2006-01-02 15:04:05"))
				concurrentUpload += 1
			} else {
				log.Info("the file " + file.Md5 + " not in 10min upload." + file.CreatedUnix.Format("2006-01-02 15:04:05"))
			}
		}
		log.Info("The concurrentUpload is " + fmt.Sprint(concurrentUpload) + " to checked " + fileName + ". user id=" + fmt.Sprint(ctx.User.ID))
		if concurrentUpload >= setting.FLOW_CONTROL.ATTACHEMENT_NUM_A_USER_LAST10M {
			log.Info("A single user cannot upload more than " + fmt.Sprint(setting.FLOW_CONTROL.ATTACHEMENT_NUM_A_USER_LAST10M) + " files within the past 10 minutes. so " + fileName + " is rejected. user id=" + fmt.Sprint(ctx.User.ID))
			return errors.New("A single user cannot upload more than " + fmt.Sprint(setting.FLOW_CONTROL.ATTACHEMENT_NUM_A_USER_LAST10M) + " files within the past 10 minutes.")
		}
		if totalSize >= setting.FLOW_CONTROL.ATTACHEMENT_SIZE_A_USER*1024*1024*1024 {
			log.Info("The total file size uploaded by a single user within the past 24 hours cannot exceed " + fmt.Sprint(setting.FLOW_CONTROL.ATTACHEMENT_SIZE_A_USER) + "G. so " + fileName + " is rejected. user id=" + fmt.Sprint(ctx.User.ID))
			return errors.New("The total file size uploaded by a single user within the past 24 hours cannot exceed " + fmt.Sprint(setting.FLOW_CONTROL.ATTACHEMENT_SIZE_A_USER) + "G.")
		}
	}
	return nil
}

func AddFileNameToCache(datasetId int64, fileName string, userId int64) {
	mutex.Lock()
	defer mutex.Unlock()
	cacheMap := getSDKUploadFileMap(REDIS_FLOW_ATTACHMENT_KEY)
	expireTimeKeys := make([]string, 0)
	currentTime := time.Now().Unix()
	for tmpKey, tmpValue := range cacheMap {
		time, err := strconv.ParseInt(tmpValue, 10, 64)
		if err == nil {
			if currentTime-time > 24*3600 {
				expireTimeKeys = append(expireTimeKeys, tmpKey)
				continue
			}
		}
	}
	for _, delKey := range expireTimeKeys {
		delete(cacheMap, delKey)
	}
	key := fmt.Sprint(datasetId) + "_" + fileName + "_" + fmt.Sprint(userId)
	value := fmt.Sprint(time.Now().Unix())
	cacheMap[key] = value
	log.Info("set key=" + key + " value=" + value + " to cache.")
	setSDKUploadFileCache(REDIS_FLOW_ATTACHMENT_KEY, cacheMap)
}

func AddModelFileNameToCache(modelId string, fileName string, userId int64) {
	modelMutex.Lock()
	defer modelMutex.Unlock()
	cacheMap := getSDKUploadFileMap(REDIS_FLOW_MODEL_ATTACHMENT_KEY)
	expireTimeKeys := make([]string, 0)
	currentTime := time.Now().Unix()
	for tmpKey, tmpValue := range cacheMap {
		time, err := strconv.ParseInt(tmpValue, 10, 64)
		if err == nil {
			if currentTime-time > 24*3600 {
				expireTimeKeys = append(expireTimeKeys, tmpKey)
				continue
			}
		}
	}
	for _, delKey := range expireTimeKeys {
		delete(cacheMap, delKey)
	}
	key := modelId + "_" + fileName + "_" + fmt.Sprint(userId)
	value := fmt.Sprint(time.Now().Unix())
	cacheMap[key] = value
	log.Info("set key=" + key + " value=" + value + " to cache.")
	setSDKUploadFileCache(REDIS_FLOW_MODEL_ATTACHMENT_KEY, cacheMap)
}

func RemoveFileFromCache(datasetId int64, fileName string, userId int64) {
	mutex.Lock()
	defer mutex.Unlock()
	key := fmt.Sprint(datasetId) + "_" + fileName + "_" + fmt.Sprint(userId)
	cacheMap := getSDKUploadFileMap(REDIS_FLOW_ATTACHMENT_KEY)
	delete(cacheMap, key)
	log.Info("remove key=" + key + " from cache.")
	setSDKUploadFileCache(REDIS_FLOW_ATTACHMENT_KEY, cacheMap)
}

func RemoveModelFileFromCache(modelId string, fileName string, userId int64) {
	modelMutex.Lock()
	defer modelMutex.Unlock()
	key := modelId + "_" + fileName + "_" + fmt.Sprint(userId)
	cacheMap := getSDKUploadFileMap(REDIS_FLOW_MODEL_ATTACHMENT_KEY)
	delete(cacheMap, key)
	log.Info("remove key=" + key + " from cache.")
	setSDKUploadFileCache(REDIS_FLOW_MODEL_ATTACHMENT_KEY, cacheMap)
}

func getSDKUploadFileMap(msgKey string) map[string]string {
	valueStr, err := redis_client.Get(msgKey)
	msgMap := make(map[string]string, 0)
	if err == nil {
		if valueStr != "" {
			err1 := json.Unmarshal([]byte(valueStr), &msgMap)
			if err1 != nil {
				log.Info("unmarshal json failed. " + err1.Error())
			}
		}
	} else {
		log.Info("Failed to load from reids. " + err.Error())
	}
	return msgMap
}

func setSDKUploadFileCache(msgKey string, msgMap map[string]string) {
	msgMapJson, _ := json.Marshal(msgMap)
	redisValue := string(msgMapJson)
	log.Info("set redis key=" + msgKey + " value=" + redisValue)
	re, err := redis_client.Setex(msgKey, redisValue, 24*3600*time.Second)
	if err == nil {
		log.Info("re =" + fmt.Sprint(re))
	} else {
		log.Info("set redis error:" + err.Error())
	}
}

func CheckFlowForDatasetSDK() error {
	cacheMap := getSDKUploadFileMap(REDIS_FLOW_ATTACHMENT_KEY)
	currentTime := time.Now().Unix()
	count := 0
	for _, tmpValue := range cacheMap {
		time, err := strconv.ParseInt(tmpValue, 10, 64)
		if err == nil {
			if currentTime-time > 24*3600 {
				continue
			}
		}
		count += 1
	}
	log.Info("total find " + fmt.Sprint(count) + " uploading files.")
	if count >= setting.FLOW_CONTROL.ALL_ATTACHEMENT_NUM_SDK {
		log.Info("The number of datasets uploaded using the SDK simultaneously cannot exceed " + fmt.Sprint(setting.FLOW_CONTROL.ALL_ATTACHEMENT_NUM_SDK))
		return errors.New("The number of datasets uploaded using the SDK simultaneously cannot exceed " + fmt.Sprint(setting.FLOW_CONTROL.ALL_ATTACHEMENT_NUM_SDK))
	}
	return nil
}

func CheckFlowForModelSDK() error {
	cacheMap := getSDKUploadFileMap(REDIS_FLOW_MODEL_ATTACHMENT_KEY)
	currentTime := time.Now().Unix()
	count := 0
	for _, tmpValue := range cacheMap {
		time, err := strconv.ParseInt(tmpValue, 10, 64)
		if err == nil {
			if currentTime-time > 24*3600 {
				continue
			}
		}
		count += 1
	}
	log.Info("total find " + fmt.Sprint(count) + " uploading files.")
	if count >= setting.FLOW_CONTROL.ALL_ATTACHEMENT_NUM_SDK {
		log.Info("The number of model files uploaded using the SDK simultaneously cannot exceed " + fmt.Sprint(setting.FLOW_CONTROL.ALL_ATTACHEMENT_NUM_SDK))
		return errors.New("The number of model files uploaded using the SDK simultaneously cannot exceed " + fmt.Sprint(setting.FLOW_CONTROL.ALL_ATTACHEMENT_NUM_SDK))
	}
	return nil
}

func CheckFlowForModel(ctx *context.Context) error {
	if ctx.User == nil {
		return errors.New("User not login.")
	}
	log.Info("start to check flow for upload model file.")
	fileName := ctx.Query("file_name")
	currentTimeNow := time.Now()
	currentLongTime := currentTimeNow.Unix()
	last24Hour := currentTimeNow.AddDate(0, 0, -1).Unix()
	filechunks, err := models.GetModelFileChunksByUserId(ctx.User.ID, last24Hour, true)
	if err == nil {
		if len(filechunks) >= setting.FLOW_CONTROL.ATTACHEMENT_NUM_A_USER_LAST24HOUR {
			log.Info("A single user cannot upload more than " + fmt.Sprint(setting.FLOW_CONTROL.ATTACHEMENT_NUM_A_USER_LAST24HOUR) + " files within the last 24 hours. so " + fileName + " is rejected. user id=" + fmt.Sprint(ctx.User.ID))
			return errors.New("A single user cannot upload more than " + fmt.Sprint(setting.FLOW_CONTROL.ATTACHEMENT_NUM_A_USER_LAST24HOUR) + " files within the last 24 hours.")
		}
		var totalSize int64
		totalSize += ctx.QueryInt64("size")
		concurrentUpload := 0
		for _, file := range filechunks {
			totalSize += file.Size
			if (currentLongTime - int64(file.CreatedUnix)) < 10*60 {
				log.Info("the file " + file.Md5 + " in 10min upload." + file.CreatedUnix.Format("2006-01-02 15:04:05"))
				concurrentUpload += 1
			} else {
				log.Info("the file " + file.Md5 + " not in 10min upload." + file.CreatedUnix.Format("2006-01-02 15:04:05"))
			}
		}
		log.Info("The concurrentUpload is " + fmt.Sprint(concurrentUpload) + " to checked " + fileName + ". user id=" + fmt.Sprint(ctx.User.ID))
		if concurrentUpload >= setting.FLOW_CONTROL.ATTACHEMENT_NUM_A_USER_LAST10M {
			log.Info("A single user cannot upload more than " + fmt.Sprint(setting.FLOW_CONTROL.ATTACHEMENT_NUM_A_USER_LAST10M) + " files within the past 10 minutes. so " + fileName + " is rejected. user id=" + fmt.Sprint(ctx.User.ID))
			return errors.New("A single user cannot upload more than " + fmt.Sprint(setting.FLOW_CONTROL.ATTACHEMENT_NUM_A_USER_LAST10M) + " files within the past 10 minutes.")
		}
		if totalSize >= setting.FLOW_CONTROL.ATTACHEMENT_SIZE_A_USER*1024*1024*1024 {
			log.Info("The total file size uploaded by a single user within the past 24 hours cannot exceed " + fmt.Sprint(setting.FLOW_CONTROL.ATTACHEMENT_SIZE_A_USER) + "G. so " + fileName + " is rejected. user id=" + fmt.Sprint(ctx.User.ID))
			return errors.New("The total file size uploaded by a single user within the past 24 hours cannot exceed " + fmt.Sprint(setting.FLOW_CONTROL.ATTACHEMENT_SIZE_A_USER) + "G.")
		}
	}
	return nil
}
