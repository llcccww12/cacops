package urchin

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/labelmsg"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
	"encoding/json"
	"fmt"
	"github.com/minio/minio-go"
	"strings"
)

type DecompressReq struct {
	SourceFile string `json:"source_file"`
	DestPath   string `json:"dest_path"`
}

var urfsClient Urchinfs

func getUrfsClient() {
	if urfsClient != nil {
		return
	}

	urfsClient = New()
}

func tryScheduleDir(endpoint, bucket, objectKey, dstPeer string) {
	println("new request dstPeer: ", dstPeer)
	urfs := New()

	scheduleResult, err := urfs.ScheduleDirToPeerByKey(endpoint, bucket, objectKey, dstPeer)
	if err != nil {
		print(err.Error())
		return
	}

	fmt.Printf("ScheduleDataToPeerByKey StatusCode:%v %v %v %v\n", scheduleResult.StatusCode, scheduleResult.DataEndpoint, scheduleResult.DataRoot, scheduleResult.DataPath)

	scheduleResult, err = urfs.CheckScheduleDirTaskStatusByKey(endpoint, bucket, objectKey, dstPeer)
	if err != nil {
		print(err.Error())
		return
	}
	fmt.Printf("CheckScheduleTaskStatusByKey StatusCode:%v %v %v %v\n", scheduleResult.StatusCode, scheduleResult.DataEndpoint, scheduleResult.DataRoot, scheduleResult.DataPath)
}

func MoveBucketInOpenIMinio(objectKeyPrefix, targetObjectPrefix, oldBucket, newBucket string) error {
	var core = storage.MinioCore
	objectInfo := core.Client.ListObjects(oldBucket, objectKeyPrefix, true, nil)
	log.Info("MoveBucketInOpenIMinio start.objectKeyPrefix=%s", objectKeyPrefix)
	count := 0
	for object := range objectInfo {
		count++
		if object.Err != nil {
			log.Error("MoveBucketInOpenIMinio object.Err=%v", object.Err)
			return object.Err
		}
		log.Debug("MoveBucketInOpenIMinio object.Key=%s", object.Key)
		newObjectKey := strings.Replace(object.Key, objectKeyPrefix, targetObjectPrefix, 1)
		err := MoveMinioFileBucket(core, object.Key, newObjectKey, oldBucket, newBucket)
		if err != nil {
			log.Error("MoveBucketInOpenIMinio MoveMinioFileBucket object.Key=%s Err=%v", object.Key, err)
			continue
		}
	}
	log.Info("MoveBucketInOpenIMinio finished.objectKeyPrefix=%s ,total=%d", objectKeyPrefix, count)
	return nil
}

func MoveMinioFileBucket(core *minio.Core, oldObjectKey, newObjectKey, oldBucket, newBucket string) error {
	_, err := core.CopyObject(oldBucket, oldObjectKey, newBucket, newObjectKey, map[string]string{})

	if err != nil {
		log.Error("MoveBucketInOpenIMinio CopyObject  err oldObjectKey=%s .%v", oldObjectKey, err)
		return err
	}

	err = core.RemoveObject(oldBucket, oldObjectKey)
	if err != nil {
		log.Error("MoveBucketInOpenIMinio RemoveObject err oldObjectKey=%s .%v", oldObjectKey, err)
	}
	return err
}

func HandleScheduleRecords() error {
	getUrfsClient()
	records, err := models.GetSchedulingRecord()
	if err != nil {
		log.Error("GetSchedulingRecord failed:%v", err)
		return err
	}

	for _, record := range records {
		var res *PeerResult
		var err error
		log.Info("try to check schedule status.record.EndPoint=%s record.Bucket=%s record.ObjectKey=%s record.ProxyServer=%s", record.EndPoint, record.Bucket, record.ObjectKey, record.ProxyServer)
		if record.IsDir {
			res, err = urfsClient.CheckScheduleDirTaskStatusByKey(record.EndPoint, record.Bucket, record.ObjectKey, record.ProxyServer)
		} else {
			res, err = urfsClient.CheckScheduleTaskStatusByKey(record.EndPoint, record.Bucket, record.ObjectKey, record.ProxyServer)
		}
		if err != nil {
			log.Error("CheckScheduleTaskStatusByKey(%d) failed:%v", record.ID, err)
			continue
		}
		log.Info("CheckScheduleTaskStatusByKey record.EndPoint=%s record.Bucket=%s record.ObjectKey=%s record.ProxyServer=%s res=%v", record.EndPoint, record.Bucket, record.ObjectKey, record.ProxyServer, res)
		record.Status = res.StatusCode
		models.UpdateScheduleCols(record, "status")

		err = handleScheduleResult(record, res)
		if err != nil {
			log.Error("HandleScheduleRecords handleScheduleResult err.%v", err)
			return err
		}

	}

	return nil
}

func handleScheduleResult(record *models.ScheduleRecord, res *PeerResult) error {
	var err error
	switch res.StatusCode {
	case models.StorageUrchinScheduleSucceed:
		log.Info("ScheduleDataToPeerByKey(%s) succeed", record.ObjectKey)
		models.UpdateScheduleLocalOperateStatus(record, models.MoveBucketOperating)
		if record.ComputeSource == models.GPUResource || record.ComputeSource == models.GCUResource {
			err = MoveBucketInOpenIMinio(res.DataPath, record.TargetObjectKey, res.DataRoot, setting.Attachment.Minio.Bucket)
			if err != nil {
				models.UpdateScheduleLocalOperateStatus(record, models.MoveBucketFailed)
				log.Error("GetBackModel MoveBucketInOpenIMinio err.%v", err)
				return err
			}
			models.UpdateScheduleLocalOperateStatus(record, models.MoveBucketSucceed)
		} else {
			decompress(res.DataRoot+"/"+res.DataPath, setting.Bucket+"/"+strings.TrimSuffix(record.ObjectKey, models.ModelSuffix))
		}

	case models.StorageUrchinScheduleProcessing:
		log.Info("ScheduleDataToPeerByKey(%s) processing", record.ObjectKey)
	case models.StorageUrchinScheduleFailed:
		log.Error("ScheduleDataToPeerByKey(%s) failed:%s", record.ObjectKey, res.StatusMsg)

	default:
		log.Info("ScheduleDataToPeerByKey(%s) failed, unknown StatusCode:%d", record.ObjectKey, res.StatusCode)
	}
	return nil
}

func decompress(sourceFile, destPath string) {
	req, _ := json.Marshal(DecompressReq{
		SourceFile: sourceFile,
		DestPath:   destPath,
	})
	err := labelmsg.SendDecompressAttachToLabelOBS(string(req))
	if err != nil {
		log.Error("SendDecompressTask to labelsystem (%s) failed:%s", sourceFile, err.Error())
	}
}

func interceptErrorMessages(err error) string {
	if err == nil {
		return ""
	}
	if len(err.Error()) > 200 {
		return err.Error()[:200]
	}
	return err.Error()
}
