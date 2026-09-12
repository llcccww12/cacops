package cloudbrain

import (
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"time"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/manager/client/grampus"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/eval"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/modelarts"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
)

func ClearCloudbrainResultSpace() {
	defer func() {
		if err := recover(); err != nil {
			log.Error("panic occurred:", err)
		}
	}()
	log.Info("clear cloudbrain one result space begin.")
	if !setting.ClearStrategy.Enabled {
		return
	}

	tasks, err := models.GetGPUStoppedNotDebugJobDaysAgo(setting.ClearStrategy.ResultSaveDays, setting.ClearStrategy.BatchSize)
	if err != nil {
		log.Warn("Failed to get cloudbrain, clear result failed.", err)
		return
	}
	debugTasks, err := models.GetGPUStoppedDebugJobDaysAgo(setting.ClearStrategy.ResultSaveDays, setting.ClearStrategy.DebugJobSize)
	if err != nil {
		log.Warn("Failed to get gpu debug cloudbrain.", err)
		return
	}

	tasks = append(tasks, debugTasks...)

	debugTasksNotDaysAgo, err := models.GetGPUStoppedDebugJobNotDaysAgo(setting.ClearStrategy.ResultSaveDays, setting.ClearStrategy.BatchSize)
	if err != nil {
		log.Warn("Failed to get gpu debug cloudbrain.", err)
		return
	}
	for _, task := range debugTasksNotDaysAgo {
		if task.Duration > 0 {
			continue
		}
		sameNameTasks, _ := models.GetCloudbrainByJobNameDesc(task.JobName)
		if len(sameNameTasks) > 0 && time.Now().Unix()-int64(sameNameTasks[len(sameNameTasks)-1].UpdatedUnix) > int64(setting.ClearStrategy.ResultSaveDays*24*60*60) {
			var runtime int64
			for _, sameNameTask := range sameNameTasks {
				if task.ID != sameNameTask.ID {
					if time.Now().Unix()-int64(sameNameTask.UpdatedUnix) <= int64(setting.ClearStrategy.ResultSaveDays*24*60*60) {
						runtime += sameNameTask.Duration
						if runtime > 0 {
							break
						}
					} else {
						break
					}
				}

			}
			if runtime == 0 {
				tasks = append(tasks, task)
			}

		}
	}

	var ids []int64
	for _, task := range tasks {
		if task.Type == models.TypeC2Net {
			DeleteLocalJobStorage(task.JobName)
			deleteC2NetTask(task)
			if task.JobType == string(models.JobTypeEval) {
				eval.ClearEvalTaskResult(task.ID)
			}
			log.Info("clear TypeC2Net,name=" + task.JobName)
			ids = append(ids, task.ID)
		}
	}
	err = models.UpdateCloudBrainRecordsCleared(ids)
	if err != nil {
		log.Warn("Failed to set cloudbrain cleared status", err)
	}
	//如果云脑表处理完了，通过遍历minio对象处理历史垃圾数据，如果存在的话
	if len(tasks) < setting.ClearStrategy.BatchSize+setting.ClearStrategy.DebugJobSize {
		clearLocalHistoryTrashFile()

	}
	log.Info("clear cloudbrain one result space end.")

}

func ClearCloudbrain(task *models.Cloudbrain) {
	DeleteLocalJobStorage(task.JobName)
	deleteC2NetTask(task)
	if task.JobType == string(models.JobTypeEval) {
		eval.ClearEvalTaskResult(task.ID)
	}
	models.UpdateCloudBrainRecordsCleared([]int64{task.ID})
}

func clearLocalHistoryTrashFile() {
	files, err := ioutil.ReadDir(setting.JobPath)
	processCount := 0
	if err != nil {
		log.Warn("Can not browser local job path.")
	} else {
		SortModTimeAscend(files)
		for _, file := range files {
			//清理n天前的历史垃圾数据，清理job目录
			if file.Name() != "" && file.ModTime().Before(time.Now().AddDate(0, 0, -setting.ClearStrategy.TrashSaveDays)) {
				has, err := models.IsCloudbrainExistByJobName(file.Name())
				if err == nil && !has {
					os.RemoveAll(setting.JobPath + file.Name())
					log.Info("clear job in local trash:" + file.Name())
					processCount++
				}
				if processCount == setting.ClearStrategy.BatchSize {
					break
				}
			} else {
				break
			}

		}

	}

}

func SortModTimeAscend(files []os.FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().Before(files[j].ModTime())
	})
}

func DeleteLocalJobStorage(jobName string) error {

	if jobName == "" {
		return nil
	}
	//delete local
	localJobPath := setting.JobPath + jobName
	err := os.RemoveAll(localJobPath)
	if err != nil {
		log.Error("RemoveAll(%s) failed:%v", localJobPath, err)
	}

	return err
}

func deleteC2NetTask(taskInfo *models.Cloudbrain) {
	grampus.DeleteJob(&entity.JobIdAndVersionId{JobID: taskInfo.JobID, JobType: taskInfo.JobType, TaskID: taskInfo.ID})
}

func DeleteNPUJobStorage(taskInfo *models.Cloudbrain) int64 {
	if taskInfo.Type == models.TypeCloudBrainTwo {
		if taskInfo.JobType == string(models.JobTypeTrain) || taskInfo.JobType == string(models.JobTypeInference) ||
			taskInfo.JobType == string(models.JobTypeModelSafety) {
			modelarts.DelTrainJobVersion(taskInfo.JobID, strconv.FormatInt(taskInfo.VersionID, 10))
		}
		if taskInfo.JobType == string(models.JobTypeDebug) {
			modelarts.DelNotebook2(taskInfo.JobID)
		}

		latestTask, err := models.GetCloudbrainByJobIDAndIsLatestVersion(taskInfo.JobID, modelarts.IsLatestVersion)
		if err == nil {
			if latestTask.ID == taskInfo.ID {
				deleteModelArtsStorage(taskInfo.JobName, "")
			} else {
				deleteModelArtsStorage(taskInfo.JobName, taskInfo.VersionName)
			}
		} else {
			deleteModelArtsStorage(taskInfo.JobName, taskInfo.VersionName)
		}
		return taskInfo.ID
	}

	return 0
}

func deleteModelArtsStorage(jobName string, versionName string) {
	if jobName == "" {
		return
	}
	localJobPath := setting.JobPath + jobName
	log.Info("delete local path=" + localJobPath)
	err := os.RemoveAll(localJobPath)
	if err != nil {
		log.Error("RemoveAll(%s) failed:%v", localJobPath, err)
	}
	dirPath := setting.CodePathPrefix + jobName
	if versionName != "" {
		dirPath = setting.CodePathPrefix + jobName + "/output/" + versionName
	}
	//delete oss
	log.Info("delete obs path=" + dirPath)
	err = storage.ObsRemoveObject(setting.Bucket, dirPath)
	if err != nil {
		log.Error("ObsRemoveObject(%s) failed:%v", localJobPath, err)
	}
}
