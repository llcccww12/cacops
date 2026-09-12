package cloudbrainTask

import (
	"fmt"
	"os"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
)

func DeleteCloudbrainJobStorage(jobName string, cloudbrainType int) error {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:", combinedErr)
		}
	}()

	//delete local
	if jobName == "" {
		log.Info("This is error.jobName is empty.")
		return nil
	}
	localJobPath := setting.JobPath + jobName
	err := os.RemoveAll(localJobPath)
	if err != nil {
		log.Error("RemoveAll(%s) failed:%v", localJobPath, err)
	}

	deleteS3Storage(jobName)
	return nil
}

func deleteS3Storage(jobName string) {
	//delete oss
	dirPath := setting.CodePathPrefix + jobName + "/"
	err := storage.ObsRemoveObject(setting.Bucket, dirPath)
	if err != nil {
		log.Error("ObsRemoveObject(%s) failed:%v", dirPath, err)
	}

	dirPath = setting.CBCodePathPrefix + jobName + "/"
	err = storage.Attachments.DeleteDir(dirPath)
	if err != nil {
		log.Error("DeleteDir(%s) failed:%v", dirPath, err)
	}

}
