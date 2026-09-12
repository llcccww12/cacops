package schedule

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"path"
	"strings"
	"time"

	"code.gitea.io/gitea/entity"

	"code.gitea.io/gitea/modules/obs"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/grampus"
	"code.gitea.io/gitea/modules/labelmsg"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/redis/redis_lock"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
	"code.gitea.io/gitea/modules/util"
	"github.com/minio/minio-go"
)

const NPUModelDefaultName = "models.zip"

func GetModelScheduleStatus(jobId string) (models.ModelMigrateStatus, error) {
	job, err := models.GetCloudbrainByJobID(jobId)
	if err != nil {
		log.Error("GetModelScheduleStatus GetCloudbrainByJobID err.jobId=%s err=%v", jobId, err)
		return 0, errors.New("jobId not correct")
	}
	if !job.IsTerminal() {
		log.Info("GetModelScheduleStatus job is not terminal.jobId=%s", jobId)
		return models.ModelMigrateWaiting, nil
	}

	record, err := models.GetModelMigrateRecordByCloudbrainId(job.ID)
	if err != nil {
		log.Error("GetModelScheduleStatus GetModelMigrateRecordByCloudbrainId err.jobId=%s err=%v", jobId, err)
		if models.IsErrRecordNotExist(err) {
			return models.ModelMigrateSuccess, nil
		}
		return models.ModelMigrateFailed, err
	}

	if !record.IsFinished() {
		go HandleUnfinishedMigrateRecord(record)
	}

	return record.Status, nil
}

func RetryModelMigrate(job *models.Cloudbrain) error {
	if !job.IsTerminal() && job.Status != string(models.ModelSafetyTesting) {
		log.Info("RetryModelMigrate job is not terminal.id=%s", job.ID)
		return errors.New("task is not terminal")
	}
	if job.Cleared {
		log.Info("RetryModelMigrate job is not cleared.id=%s", job.ID)
		return errors.New("ai_task.task_cleared_cannot_get_result")
	}

	//避免并发问题，先尝试获取锁，获取锁以后再查最新的记录
	lock := redis_lock.NewDistributeLock(redis_key.RecordHandleLock(job.JobID))
	success, err := lock.LockWithWait(10*time.Second, 10*time.Second)
	if err != nil {
		log.Error("HandleUnfinishedMigrateRecord lock err.id=%d %v", job.ID, err)
		return err
	}
	if !success {
		log.Error("HandleUnfinishedMigrateRecord lock failed.ID=%d ", job.ID)
		return nil
	}
	defer lock.UnLock()

	record, err := models.GetModelMigrateRecordByCloudbrainId(job.ID)
	if err != nil {
		log.Error("RetryModelMigrate GetModelMigrateRecordByCloudbrainId err.id=%s err=%v", job.ID, err)
		if models.IsErrRecordNotExist(err) {
			return nil
		}
		return err
	}

	//只有两种情况可以再次调度，一是虎鲸调度失败 二是本地移桶失败
	if record.CurrentStep == models.GrampusMigrateFailed {
		log.Info("retry PostModelMigrate. record.id = %d", record.ID)
		_, err := grampus.PostModelMigrate(job.JobID)
		if err != nil {
			log.Error("PostModelMigrate err.%v", err)
			return err
		}
		models.IncreaseModelMigrateRetryCount(record.ID)
		if err := models.RollBackMigrateStatus(record, models.GrampusMigrating); err != nil {
			log.Error("UpdateModelMigrateStatusByStep err.%v", err)
			return err
		}
		return nil
	}

	if record.CurrentStep == models.BucketMoveFailed {
		log.Info("retry BucketMove. record.id = %d", record.ID)
		if err := models.RollBackMigrateStatus(record, models.GrampusMigrateSuccess); err != nil {
			log.Error("UpdateModelMigrateStatusByStep err.%v", err)
			return err
		}
		models.IncreaseModelMigrateRetryCount(record.ID)
		return nil
	}

	return errors.New("No need to retry,the model migration has been successful or is in the process.")

}

func HandleUnfinishedMigrateRecords() {
	page := 1
	pageSize := 100
	count := 0
	for {
		list, err := models.GetUnfinishedModelMigrateRecordsPaging(pageSize, page)
		if err != nil {
			log.Error("GetUnfinishedModelMigrateRecords err=%v", err)
			return
		}
		for _, r := range list {
			HandleUnfinishedMigrateRecord(r)
		}
		count += len(list)
		if len(list) < pageSize {
			break
		}
		if count > 100000 {
			break
		}
	}

}

func HandleUnfinishedMigrateRecord(r *models.ModelMigrateRecord) error {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:", combinedErr)
		}
	}()

	cloudbrain, err := models.GetCloudbrainByID(fmt.Sprint(r.CloudbrainID))
	if err != nil {
		log.Error("GetCloudbrainByID err. %v", err)
		return err
	}

	lock := redis_lock.NewDistributeLock(redis_key.RecordHandleLock(cloudbrain.JobID))
	success, err := lock.LockWithWait(10*time.Second, 10*time.Second)
	if err != nil {
		log.Error("HandleUnfinishedMigrateRecord lock err.ID=%d %v", r.ID, err)
		return err
	}
	if !success {
		log.Error("HandleUnfinishedMigrateRecord lock failed.ID=%d ", r.ID)
		return nil
	}
	defer lock.UnLock()
	//拿到锁以后重新查询一次
	r, err = models.GetModelMigrateRecordById(r.ID)
	if err != nil {
		log.Error("RetryModelMigrate GetModelMigrateRecordById err.Id=%s err=%v", r.ID, err)
		if models.IsErrRecordNotExist(err) {
			return nil
		}
		return err
	}

	if r.CurrentStep == models.GrampusMigrateInit || r.CurrentStep == models.GrampusMigrating {
		if err := UpdateModelMigrateStatusFromGrampus(r, cloudbrain.JobID); err != nil {
			log.Error("UpdateModelMigrateStatusFromGrampus err. %v", err)
			return err
		}
	}

	if r.CurrentStep == models.GrampusMigrateSuccess {
		if err := LocalMigrateOperate(cloudbrain.JobName, cloudbrain.ComputeResource, r); err != nil {
			log.Error("LocalMigrateOperate err. %v", err)
			return err
		}
	}

	return nil
}

func UpdateModelMigrateStatusFromGrampus(r *models.ModelMigrateRecord, jobId string) error {
	res, err := grampus.ModelMigrateInfo(jobId)
	if err != nil {
		log.Error("ModelMigrateInfo err. r.ID=%d  %v", r.ID, err)
		return err
	}
	if res.Status == int(models.GrampusMigrateResponseNoNeedMigrate) {
		if r.RetryCount < 2 {
			log.Warn("ModelMigrateInfo 404 err retry. r.ID=%d", r.ID)
			grampus.PostModelMigrate(jobId)
			models.IncreaseModelMigrateRetryCount(r.ID)
			return nil
		}

	}
	log.Info("grampus ModelMigrateInfo r.ID=%d res=%+v", r.ID, res)
	newStep := models.GrampusMigrateResponse(res.Status).ConvertToModelMigrateStep()
	if newStep == r.CurrentStep {
		log.Info("The status has not changed. r.ID=%d status=%d", r.ID, res.Status)
		return nil
	}
	err = updateModelMigrateFromRes(r, res)
	if err != nil {
		log.Error("updateModelMigrateFromRes err. r.ID=%d  %v", r.ID, err)
		return err
	}
	return nil
}

func LocalMigrateOperate(jobName, computeSource string, r *models.ModelMigrateRecord) error {
	log.Info("Grampus model migrate succeed,objectKey = %s  computeSource= %s", r.DestObjectKey, computeSource)
	err := models.UpdateModelMigrateStatusByStep(r, models.BucketMoving)
	if err != nil {
		log.Error("LocalMigrateOperate UpdateModelMigrateStatusByStep err. r.ID=%d step=%d err=%v", r.ID, models.BucketMoveFailed, err)
		return err
	}
	cloudbrain, err := models.GetCloudbrainByCloudbrainID(r.CloudbrainID)
	if err != nil {
		log.Error("LocalMigrateOperate GetCloudbrainByCloudbrainID err. r.ID=%d step=%d err=%v", r.ID, models.BucketMoveFailed, err)
		return err
	}
	aiConfig := cloudbrain.GetCloudbrainConfig()

	//判断来源是用的什么存储
	//再判断目标是什么存储
	sourceStorageType := entity.URCHIN_V2
	if setting.FORCE_OLD_OUTPUT_MIGRATE || r.MigrateService == models.OutputMigrateSystemUrchin1 {
		if aiConfig == nil || aiConfig.OutputStorageType == "" {
			if cloudbrain.ComputeResource == models.NPUResource {
				sourceStorageType = entity.OBS
			}
		} else {
			sourceStorageType = entity.StorageType(aiConfig.OutputStorageType)
		}
	}

	switch sourceStorageType {
	case entity.MINIO:
		err = OutputMigrate4Minio(aiConfig, r)
	case entity.OBS:
		err = OutputMigrate4OBS(aiConfig, r)
	case entity.URCHIN_V2:
		err = OutputMigrate4UrchinV2(aiConfig, r)
	default:
		log.Error("LocalMigrateOperate output storage type invalid")
		return fmt.Errorf(" output storage type invalid")
	}

	if err != nil {
		log.Error("LocalMigrateOperate output migrate  error.  r.ID=%d step=%d err=%v", r.ID, models.MigrateFinished, err)
		return err
	}

	if err := models.UpdateModelMigrateStatusByStep(r, models.MigrateFinished); err != nil {
		log.Error("LocalMigrateOperate UpdateModelMigrateStatusByStep  error.  r.ID=%d step=%d err=%v", r.ID, models.MigrateFinished, err)
	}
	return nil

}

func OutputMigrate4OBS(aiConfig *models.CloudbrainConfig, r *models.ModelMigrateRecord) error {
	targetObjectKeyPrex := strings.TrimSuffix(aiConfig.OutputObjectPrefix, "/") + "/"
	sourceObjectKey := strings.TrimSuffix(r.DestObjectKey, "/") + "/"
	if err := MoveBucketInOpenIOBS(sourceObjectKey, targetObjectKeyPrex, r.DestBucket, aiConfig.OutputBucket); err != nil {
		log.Error("MoveBucketInOpenIObs err.%v", err)
		if tmpErr := models.UpdateModelMigrateStatusByStep(r, models.BucketMoveFailed); tmpErr != nil {
			log.Error("UpdateModelMigrateStatusByStep  error.  r.ID=%d step=%d err=%v", r.ID, models.BucketMoveFailed, tmpErr)
		}
		return err
	}
	return nil
}

func OutputMigrate4Minio(aiConfig *models.CloudbrainConfig, r *models.ModelMigrateRecord) error {
	targetObjectKeyPrex := strings.TrimSuffix(aiConfig.OutputObjectPrefix, "/") + "/"
	sourceObjectKey := strings.TrimSuffix(r.DestObjectKey, "/") + "/"
	if err := MoveBucketInOpenIMinio(sourceObjectKey, targetObjectKeyPrex, r.DestBucket, aiConfig.OutputBucket); err != nil {
		log.Error("MoveBucketInOpenIMinio err.%v", err)
		if tmpErr := models.UpdateModelMigrateStatusByStep(r, models.BucketMoveFailed); tmpErr != nil {
			log.Error("UpdateModelMigrateStatusByStep  error.  r.ID=%d step=%d err=%v", r.ID, models.BucketMoveFailed, tmpErr)
		}
		return err
	}
	return nil
}

func OutputMigrate4UrchinV2(aiConfig *models.CloudbrainConfig, r *models.ModelMigrateRecord) error {
	aiConfig.OutputStorageType = string(entity.URCHIN_V2)
	aiConfig.OutputObjectPrefix = r.DataID
	_, err := models.UpdateCloudbrainConfigByTaskID(aiConfig.CloudbrainID, aiConfig)
	return err
}

func obsMkdir(dir string) error {
	input := &obs.PutObjectInput{}
	input.Bucket = setting.Bucket
	input.Key = dir
	_, err := storage.ObsCli.PutObject(input)
	if err != nil {
		log.Error("PutObject(%s) failed: %s", input.Key, err.Error())
		return err
	}

	return nil
}

func TryToUpdateNPUMoveBucketResult(record *models.ModelMigrateRecord, jobName, versionName string) error {
	if IsNPUModelDirHasFile(jobName, versionName) {
		if err := models.UpdateModelMigrateStatusByStep(record, models.MigrateFinished); err != nil {
			log.Error("UpdateModelMigrateStatusByStep  error.  r.ID=%d step=%d err=%v", record.ID, models.MigrateFinished, err)
			return err
		}
	}
	return nil
}

func updateModelMigrateFromRes(r *models.ModelMigrateRecord, res *models.GrampusModelMigrateInfoResponse) error {
	step := models.GrampusMigrateResponse(res.Status).ConvertToModelMigrateStep()
	err := models.UpdateModelMigrateStatusByStep(r, step)
	if err != nil {
		log.Error("UpdateModelMigrateStatusByStep err,ID=%d err=%v", r.ID, err)
		return err
	}
	r.DestBucket = res.DestBucket
	r.DestEndpoint = res.DestEndpoint
	r.DestObjectKey = res.DestObjectKey
	r.DestProxy = res.DestProxy
	r.Remark = strings.TrimPrefix(r.Remark+";"+util.TruncateString(res.FailedReason, 200), ";")
	r.SrcBucket = res.SrcBucket
	r.SrcEndpoint = res.SrcEndpoint
	r.SrcObjectKey = res.SrcObjectKey
	r.DataID = res.DataId
	r.GrampusMigrateTaskId = res.MigrateTaskId
	r.MigrateService = res.MigrateService
	err = models.UpdateModelMigrateRecordByStep(r)
	if err != nil {
		log.Error("updateModelMigrateFromRes UpdateModelMigrateRecord error.id=%d.err=%v", r.ID, err)
		return err
	}
	return nil
}

func MoveBucketInOpenIMinio(objectKeyPrefix, targetObjectPrefix, oldBucket, newBucket string) error {
	//删除目标桶原有内容
	err := DeleteDir4Minio(newBucket, targetObjectPrefix)
	if err != nil {
		log.Error("MoveBucketInOpenIMinio DeleteDir4Minio err.objectKeyPrefix=%s err=%v", objectKeyPrefix, err)
	}
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
	//删除源桶
	err = DeleteDir4Minio(oldBucket, objectKeyPrefix)
	if err != nil {
		log.Error("MoveBucketInOpenIMinio DeleteDir4Minio in old bucket err.objectKeyPrefix=%s err=%v", objectKeyPrefix, err)
	}
	log.Info("MoveBucketInOpenIMinio finished.objectKeyPrefix=%s ,total=%d", objectKeyPrefix, count)
	return nil
}

func DeleteDir4Minio(bucket, dirObjectKey string) error {
	var core = storage.MinioCore
	objectInfo := core.Client.ListObjects(bucket, dirObjectKey, true, nil)
	log.Info("DeleteDir4Minio start.objectKeyPrefix=%s", dirObjectKey)
	count := 0
	for object := range objectInfo {
		count++
		if object.Err != nil {
			log.Error("DeleteDir4Minio object.Err=%v", object.Err)
			return object.Err
		}
		log.Debug("DeleteDir4Minio object.Key=%s", object.Key)
		err := core.RemoveObject(bucket, object.Key)
		if err != nil {
			log.Error("DeleteDir4Minio RemoveObject err objectKey=%s .%v", object.Key, err)
		}
	}
	log.Info("DeleteDir4Minio finished.objectKeyPrefix=%s ,total=%d", dirObjectKey, count)
	return nil
}

func MoveBucketInOpenIOBS(objectKeyPrefix, targetObjectPrefix, oldBucket, newBucket string) error {
	var client = storage.ObsCli

	input := &obs.ListObjectsInput{}
	input.Bucket = oldBucket
	input.MaxKeys = 1000
	input.Prefix = objectKeyPrefix
	output, err := client.ListObjects(input)
	if err != nil {
		log.Error("MoveBucketInOpenIOBS list objects error.err=%v", err)
		return err
	}
	//删除目标桶原有内容
	err = DeleteDir4OBS(newBucket, targetObjectPrefix)
	if err != nil {
		log.Error("MoveBucketInOpenIOBS DeleteDir4OBS  err.objectKeyPrefix=%s err=%v", objectKeyPrefix, err)
	}
	log.Info("MoveBucketInOpenIOBS start.objectKeyPrefix=%s", objectKeyPrefix)
	count := 0
	for _, object := range output.Contents {
		if object.Key == input.Prefix {
			continue
		}
		if strings.HasSuffix(object.Key, "/") {
			continue
		}
		log.Debug("MoveBucketInOpenIOBS object.Key=%s", object.Key)
		newObjectKey := strings.Replace(object.Key, objectKeyPrefix, targetObjectPrefix, 1)
		err := MoveOBSFileBucket(client, object.Key, newObjectKey, oldBucket, newBucket)
		if err != nil {
			log.Error("MoveBucketInOpenIOBS MoveOBSFileBucket object.Key=%s Err=%v", object.Key, err)
			continue
		}
	}
	//删除源桶
	err = DeleteDir4OBS(oldBucket, objectKeyPrefix)
	if err != nil {
		log.Error("MoveBucketInOpenIOBS DeleteDir4OBS in old bucket err.objectKeyPrefix=%s err=%v", objectKeyPrefix, err)
	}
	log.Info("MoveBucketInOpenIOBS finished.objectKeyPrefix=%s ,total=%d", objectKeyPrefix, count)
	return nil
}

func DeleteDir4OBS(bucket, dirObjectKey string) error {
	var client = storage.ObsCli

	input := &obs.ListObjectsInput{}
	input.Bucket = bucket
	input.MaxKeys = 1000
	input.Prefix = dirObjectKey
	output, err := client.ListObjects(input)
	if err != nil {
		log.Error("DeleteDir4OBS list objects error.err=%v", err)
		return err
	}
	count := 0
	for _, object := range output.Contents {
		delObj := &obs.DeleteObjectInput{}
		delObj.Bucket = bucket
		delObj.Key = object.Key
		_, err = client.DeleteObject(delObj)
		if err != nil {
			log.Error("DeleteDir4OBS DeleteObject err oldObjectKey=%s .%v", object.Key, err)
		}
	}

	log.Info("DeleteDir4OBS finished.objectKeyPrefix=%s ,total=%d", dirObjectKey, count)
	return nil
}

func MoveBucketJust4LocalMinio(objectKeyPrefix, targetObjectPrefix, oldBucket, newBucket string) error {
	oldPath := path.Join(setting.Attachment.Minio.RealPath, oldBucket, objectKeyPrefix)
	newPath := path.Join(setting.Attachment.Minio.RealPath, newBucket, targetObjectPrefix)
	log.Info("MoveBucketJust4LocalMinio start.oldPath=%s  newPath=%s", oldPath, newPath)
	//重命名原有文件夹,防止已有该文件
	err, errStr := sudoMv(newPath, fmt.Sprintf("%s_%d", newPath, time.Now().Unix()))
	if err != nil {
		log.Error("MoveBucketJust4LocalMinio sudoMv error.oldPath=%s  newPath=%s Err=%v errStr=%s ", oldPath, newPath, err, errStr)
	}
	//移动（重命名）文件夹
	err, errStr = sudoMv(oldPath, newPath)
	if err != nil {
		log.Error("MoveBucketJust4LocalMinio sudoMv error.oldPath=%s  newPath=%s Err=%v errStr=%s ", oldPath, newPath, err, errStr)
		return err
	}
	log.Info("MoveBucketInOpenIMinio finished.oldPath=%s  newPath=%s  ", oldPath, newPath)
	return nil
}

func sudoMv(oldPath, newPath string) (error, string) {
	c := fmt.Sprintf("sudo mv %s %s", oldPath, newPath)
	log.Info("start to sudoMv,oldPath=%s  newPath=%s", oldPath, newPath)
	cmd := exec.Command("/bin/sh", "-c", c)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	log.Debug("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		log.Error("cmd.Run() failed,oldPath=%s  newPath=%s   err=%v\n", oldPath, newPath, err)
		return err, errStr
	}
	return nil, errStr
}

func MoveMinioFileBucket(core *minio.Core, oldObjectKey, newObjectKey, oldBucket, newBucket string) error {
	_, err := core.CopyObject(oldBucket, oldObjectKey, newBucket, newObjectKey, map[string]string{})

	if err != nil {
		log.Error("MoveBucketInOpenIMinio CopyObject  err oldObjectKey=%s .%v", oldObjectKey, err)
		//return err
	}

	//err = core.RemoveObject(oldBucket, oldObjectKey)
	//if err != nil {
	//	log.Error("MoveBucketInOpenIMinio RemoveObject err oldObjectKey=%s .%v", oldObjectKey, err)
	//}
	return err
}

func MoveOBSFileBucket(client *obs.ObsClient, oldObjectKey, newObjectKey, oldBucket, newBucket string) error {
	input := &obs.CopyObjectInput{}
	input.Bucket = newBucket
	input.Key = newObjectKey
	input.CopySourceBucket = oldBucket
	input.CopySourceKey = oldObjectKey
	_, err := client.CopyObject(input)

	if err != nil {
		log.Error("MoveOBSFileBucket CopyObject  err oldObjectKey=%s .%v", oldObjectKey, err)
		return err
	}
	//delObj := &obs.DeleteObjectInput{}
	//delObj.Bucket = oldBucket
	//delObj.Key = oldObjectKey
	//_, err = client.DeleteObject(delObj)
	//if err != nil {
	//	log.Error("MoveOBSFileBucket DeleteObject err oldObjectKey=%s .%v", oldObjectKey, err)
	//}
	return err
}

type DecompressReq struct {
	SourceFile string `json:"source_file"`
	DestPath   string `json:"dest_path"`
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

func IsNPUModelDirHasFile(jobName string, versionName string) bool {
	prefix := strings.TrimPrefix(path.Join(setting.TrainJobModelPath, jobName, setting.OutPutPath, versionName), "/")
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}
	fileInfos, err := storage.GetOneLevelAllObjectUnderDir(setting.Bucket, prefix, "")
	if err != nil {
		log.Info("IsNPUModelDirHasFile.get TrainJobListModel failed:", err)
		return false
	}

	if len(fileInfos) > 0 {
		return true
	}
	return len(fileInfos) > 0
}
