package finetune

import (
	"net/http"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/notification"
	"code.gitea.io/gitea/modules/setting"
	api "code.gitea.io/gitea/modules/structs"

	// for deployment
	"code.gitea.io/gitea/modules/modelarts"
	"code.gitea.io/gitea/modules/storage"
	"code.gitea.io/gitea/modules/timeutil"
)

const (
	PanguModelType         = "MindSpore"
	PanguModelVersion      = "1.0.0"
	PanguDeployType        = "real-time"
	PanguDeployInstanceCnt = 1
	//PanguDeploySpec         = "custom"
	PanguDeployScheduleType = "stop"
)

func CreateAIModel(ctx *context.APIContext, jobID string) error {
	deploy, _ := models.GetModelartsDeployByJobID(jobID)

	jobName := deploy.JobName
	deployPath := modelarts.JobPath[1:] + jobName + "/model/"
	panguBucket := setting.FineTune.Pangu.Deploy.CodeBucket
	panguCodePath := setting.FineTune.Pangu.Deploy.CodeObsPath
	if _, err := storage.ObsCopyAllFile(panguBucket, panguCodePath, setting.Bucket, deployPath); err != nil {
		log.Error("panguService: jobID %s Modelart ModelartsDeploy: Failed to copy %s to /model (%v)", jobID, panguCodePath, err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("deployment.code_copy_failed")))
		return err
	}
	//log.Info("panguService: 拷贝推理代码文件成功, 源路径: %s, 目标路径: (%s)", panguBucket+"/"+panguCodePath, setting.Bucket+"/"+deployPath)

	trainModel := modelarts.JobPath[1:] + jobName + modelarts.OutputPath + "V0001/" + setting.FineTune.Pangu.Deploy.ModelName
	deployModel := deployPath + setting.FineTune.Pangu.Deploy.ModelName
	if err := storage.ObsCopyFile(setting.Bucket, trainModel, setting.Bucket, deployModel); err != nil {
		log.Error("panguService: jobID %s Modelart ModelartsDeploy: Failed to copy %s to /model (%v)", jobID, trainModel, err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("deployment.model_copy_failed")))
		return err
	}
	//log.Info("panguService: 拷贝模型成功, 源路径: %s, 目标路径: %s", setting.Bucket+"/"+trainModel, setting.Bucket+"/"+deployPath)
	deploy.DeployUrl = setting.Bucket + "/" + deployPath
	deploy.UpdateUnix = timeutil.TimeStampNow()
	models.UpdateDeploy(deploy)

	deployModelReq := &modelarts.GenerateDeployModelReq{
		JobID:          jobID,
		ModelName:      jobName + "_model",
		ModelVersion:   PanguModelVersion,
		ModelType:      PanguModelType,
		SourceLocation: setting.Endpoint[:8] + setting.Bucket + "." + setting.Endpoint[8:] + "/" + deployPath,
		UserID:         deploy.UserID,
		Runtime:        setting.FineTune.Pangu.Deploy.Runtime,
		InstallType:    []string{PanguDeployType},
	}

	modelID, err := modelarts.GenerateDeployModel(ctx.Context, deployModelReq)
	if err != nil {
		log.Error("panguService: 微调任务 [%s] AI应用 API调用失败:%v", jobID, err.Error())
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("deployment.builidng_fail")))
		return err
	}
	log.Info("panguService: 微调任务 [%s], AI应用API调用成功, 模型ID [%s]", jobID, modelID)
	return nil
}

func PanguServiceCreateQueue() {
	runningDeploys, _ := models.GetAllRunningService()
	var cntQueue int
	if runningDeploys == nil {
		cntQueue = setting.FineTune.Pangu.Deploy.MaxDeployNum
		//log.Info("panguService: 部署队列 当前无部署任务，可用运行资源 %v 卡。", cntQueue)
	} else {
		cntQueue = setting.FineTune.Pangu.Deploy.MaxDeployNum - len(runningDeploys)
		log.Info("panguService: 部署队列 当前正在运行 %v 部署任务，可用运行资源 %v 卡。", len(runningDeploys), cntQueue)
		if cntQueue == 0 {
			//log.Info("panguService: 资源已满，停止当前定时任务。")
			return
		}
	}

	deployQueues, _ := models.GetModelartsDeployQueue(cntQueue)
	if deployQueues == nil {
		//log.Error("panguService: 部署队列 空，停止当前定时任务。")
		return
	}

	for _, deployQueue := range deployQueues {
		if deployQueue.ServiceID != "" {
			//服务重启
			err := modelarts.UpdateDeployService(deployQueue.ServiceID, models.UpdateDeployServiceParams{
				Status: "running",
			})
			if err != nil {
				log.Error("panguService: 微调任务 [%s],部署服务重启失败, 服务ID %s", deployQueue.JobID, deployQueue.ServiceID)
				models.DeleteModelartsDeployQueueByJobID(deployQueue.JobID)
				continue
			}
			log.Error("panguService: 微调任务 [%s],部署服务重启成功", deployQueue.JobID)
		} else {
			//初次部署服务
			serviceReq := &modelarts.GenerateDeployServiceReq{
				JobID:         deployQueue.JobID,
				InferType:     PanguDeployType,
				ServiceName:   deployQueue.ModelName + "-service",
				Spec:          setting.FineTune.Pangu.Deploy.Specification,
				Duration:      setting.FineTune.Pangu.Deploy.Duration + setting.FineTune.Pangu.Deploy.WarmupDuration,
				TimeUnit:      setting.FineTune.Pangu.Deploy.TimeUnit,
				ScheduleType:  PanguDeployScheduleType,
				ModelID:       deployQueue.ModelID,
				InstanceCount: PanguDeployInstanceCnt,
			}
			serviceID, err := modelarts.GenerateDeployService(serviceReq)
			if err != nil {
				log.Error("panguService: 微调任务 [%s],部署服务创建失败, 服务ID %s", deployQueue.JobID, serviceID)
				models.DeleteModelartsDeployQueueByJobID(deployQueue.JobID)
				continue
			}
			log.Info("panguService: 微调任务 [%s] 部署服务创建成功, serviceID %s", deployQueue.JobID, serviceID)
		}
		models.DeleteModelartsDeployQueueByJobID(deployQueue.JobID)
	}
}

func SyncPanguDeployStatus() {
	deployments, _ := models.GetAllModelartsDeploys()
	for _, deploy := range deployments {
		var statusNew string
		if deployQueue, _ := models.GetModelartsDeployQueueByJobID(deploy.JobID); deployQueue != nil || deploy.Status == "FAILED" {
			log.Info("panguServic: 微调任务 [%s] 不更新状态, status %s", deploy.JobID, deploy.Status)
			continue
		} else if deploy.ServiceID == "" {
			// 查模型
			deployNew, err := modelarts.GetDeployModel(deploy.ModelID)
			if err != nil {
				log.Error("panguService: 微调任务 [%s] Get DeployModel API failed:%v", deploy.DisplayJobName, err.Error())
				return
			}
			statusNew = models.DeployStatusConvert(deployNew.ModelStatus)
			statusOld := deploy.Status
			deploy.Status = statusNew
			deploy.ModelStatus = deployNew.ModelStatus
			log.Info("panguService: status update 微调任务 [%s], OLD %s, NEW %s, model api %s", deploy.JobID, statusOld, statusNew, deployNew.ModelStatus)
			if statusOld == "BUILDING" && statusNew == "WAITING" {
				models.CreateModelartsDeployQueue(&models.ModelartsDeployQueue{
					JobID:          deploy.JobID,
					ModelID:        deploy.ModelID,
					ModelName:      deploy.ModelName,
					DisplayJobName: deploy.DisplayJobName,
					CreateUnix:     timeutil.TimeStampNow(),
				})
			}
		} else {
			// 查服务
			deployNew, err := modelarts.GetDeployService(deploy.ServiceID)
			if err != nil {
				log.Error("panguService: 微调任务 [%s] Get DeployService API failed:%v", deploy.DisplayJobName, err.Error())
				return
			}
			statusNew = models.DeployStatusConvert(deployNew.Status)
			deploy.ServiceStatus = deployNew.Status
			// 部署成功后，等待30分钟再允许用户打开推理界面
			if statusNew == "SUCCEEDED" {
				currentTime := timeutil.TimeStampNow()
				//log.Info("panguService：deploy.CompleteUnix.IsZero() %v, deploy.CompleteUnix == timeutil.TimeStamp(0) %v", deploy.CompleteUnix.IsZero(), deploy.CompleteUnix == timeutil.TimeStamp(0))
				if deploy.CompleteUnix == timeutil.TimeStamp(0) {
					warmupTime := int64(setting.FineTune.Pangu.Deploy.WarmupDuration * 60)
					deploy.CompleteUnix = currentTime.Add(warmupTime)
					deploy.UpdateUnix = currentTime
					models.UpdateDeploy(deploy)
					log.Info("panguService：%s 部署成功，预热开始，预计完成时间：%v", deploy.DisplayJobName, deploy.CompleteUnix)
					continue
				} else if deploy.CompleteUnix >= currentTime {
					log.Info("panguService：%s 预热中，预计完成时间：%v", deploy.DisplayJobName, deploy.CompleteUnix)
					continue
				}
			}
			statusOld := deploy.Status
			deploy.Status = statusNew
			deploy.InferAddr = deployNew.InferAddr
			log.Info("panguService: status update 微调任务 [%s], OLD %s, NEW %s, service api %s", deploy.JobID, statusOld, statusNew, deployNew.Status)
			if statusNew == "SUCCEEDED" && statusOld == "DEPLOYING" {
				notification.NotifyChangeFinetuneStatus(deploy)
				log.Info("panguService：%s 部署服务预热完成，发送微信通知", deploy.DisplayJobName)
			}
		}
		deploy.UpdateUnix = timeutil.TimeStampNow()
		models.UpdateDeploy(deploy)
	}
}

func GetPanguDeployStatus(ctx *context.APIContext) {
	var jobID = ctx.Params(":jobid")

	status, err := models.GetModelartsDeployStatusByJobID(jobID)
	if err != nil {
		log.Info("panguService: GetPanguDeployStatus, jobID %s, err %v", jobID, status, err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
		return
	}

	finishTime, err := models.GetModelartsDeployFinishTimebyJobID(jobID)
	if err != nil {
		log.Info("panguService: GetModelartsDeployFinishTimebyJobID, jobID %s, err %v", jobID, status, err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(err.Error()))
		return
	}

	log.Info("panguService: GetPanguDeployStatus, jobID %s, status %s, finishTime %s", jobID, status, finishTime)

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"fineTuneDeployStatus":     status,
		"fineTuneDeployFinishTime": finishTime,
	})
}

func ServiceInference(ctx *context.APIContext, option api.PanguInferenceOption) {
	deploy, err := models.GetModelartsDeployByJobID(option.JobID)
	if err != nil {
		log.Error("panguService: %s Get ModelartsDeploy from DB failed:%v", deploy.DisplayJobName, err.Error())
		return
	}

	//serviceID := deploy.ServiceID
	inferAddr := deploy.InferAddr
	inferResults, err := modelarts.SendInferenceDeploy(inferAddr, option.Text)
	if err != nil {
		log.Error("panguService: %s SendInferenceDeploy API failed:%v", deploy.DisplayJobName, err.Error())
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 1,
			"text": err.Error(),
		})
		return
	}
	log.Info("panguService: %s 推理结果返回成功, %s", deploy.DisplayJobName, inferResults.GenerateResult)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"text": inferResults.GenerateResult,
	})
}

func ServiceUpdate(ctx *context.APIContext, option api.UpdateFineTuneDeployOption) {

	jobID := option.JobID
	deploy, err := models.GetModelartsDeployByJobID(jobID)
	if err != nil {
		log.Error("panguService: Get ModelartsDeploy from DB failed:%v", err.Error())
		return
	}

	if option.Status == "running" {
		// get running service count
		deployments, _ := models.GetRunningServiceByUser(ctx.User.ID)
		if len(deployments) >= setting.FineTune.Pangu.Deploy.MaxDeployPerUser {
			log.Error("panguService: 每个用户最多只能同时部署%v个模型", ctx.User.ID, len(deployments))
			ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("deployment.deploy_max", setting.FineTune.Pangu.Deploy.MaxDeployPerUser)))
			return
		}
		//log.Info("panguService: 重启服务开始")
		deploy.Status = "WAITING"
		deploy.CompleteUnix = timeutil.TimeStamp(0)
		models.CreateModelartsDeployQueue(&models.ModelartsDeployQueue{
			JobID:      jobID,
			ModelID:    deploy.ModelID,
			ModelName:  deploy.ModelName,
			ServiceID:  deploy.ServiceID,
			CreateUnix: timeutil.TimeStampNow(),
		})
	} else {
		log.Info("panguService: 停止服务")
		deploy.Status = "STOP"
		err := modelarts.UpdateDeployService(deploy.ServiceID, models.UpdateDeployServiceParams{
			Status: option.Status,
		})
		if err != nil {
			log.Error("panguService: %s Update DeployService API failed:%v", deploy.DisplayJobName, err.Error())
			ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(ctx.Tr("deployment.stop_service_failed")))
		}
	}
	deploy.UpdateUnix = timeutil.TimeStampNow()
	models.UpdateDeploy(deploy)
	ctx.JSON(http.StatusOK, models.BaseMessageApi{
		Code:    0,
		Message: option.JobID,
	})
}
