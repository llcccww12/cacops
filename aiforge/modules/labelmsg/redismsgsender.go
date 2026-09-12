package labelmsg

import (
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
)

// 方法名

func SendAddAttachToLabelSys(attach string) error {

	redisclient := Get()

	//记得销毁本次链连接
	defer redisclient.Close()

	_, err := redisclient.Do("Publish", setting.LabelTaskName, attach)
	if err != nil {
		log.Critical("redis Publish failed.")
	}

	log.Info("LabelRedisQueue(%s) success", attach)

	return nil
}

func SendDeleteAttachToLabelSys(attach string) error {

	redisclient := Get()

	//记得销毁本次链连接
	defer redisclient.Close()

	_, err := redisclient.Do("Publish", setting.LabelDatasetDeleteQueue, attach)
	if err != nil {
		log.Critical("redis Publish failed.")
	}

	log.Info("LabelDatasetDeleteQueue(%s) success", attach)

	return nil
}

func SendDecompressAttachToLabelOBS(attach string) error {

	redisclient := Get()
	//记得销毁本次链连接
	defer redisclient.Close()

	_, err := redisclient.Do("Publish", setting.DecompressOBSTaskName, attach)
	if err != nil {
		log.Critical("redis Publish failed.")
		log.Info("send unzip err = " + err.Error())
		return err
	}

	log.Info("LabelDecompressOBSQueue(%s) success", attach)
	return nil
}
