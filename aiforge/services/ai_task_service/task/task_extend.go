package task

import (
	"encoding/json"
	"fmt"
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/storage"
)

func GetModelDownload(task *models.Cloudbrain) []*models.Model4Show {
	var pretrainModelList []*models.Model4Show
	var model *models.AiModelManage
	var err error
	if task.ModelId == "" {
		return []*models.Model4Show{}
	}
	modelIdArray := task.GetModelIdArray()
	modelNameArray := task.GetModelNameArray()
	for i := 0; i < len(modelIdArray); i++ {
		modelId := modelIdArray[i]
		model, err = models.QueryModelById(modelId)
		if err != nil || model == nil {
			oldModelName := ""
			if len(modelNameArray) > i {
				oldModelName = modelNameArray[i]
			}
			pretrainModelList = append(pretrainModelList, &models.Model4Show{
				IsDelete: true,
				Name:     oldModelName,
			})
			continue
		}

		var repositoryLink string
		if r, err := models.QueryModelRepoByModelID(modelId); err == nil {
			repositoryLink = r.Link()
		}

		var ownerName string
		alias := model.DisplayName()
		model.GetOwner()
		if model.Owner != nil {
			ownerName = model.Owner.Name
		}
		pretrainModelList = append(pretrainModelList, &models.Model4Show{
			ID:             modelId,
			IsDelete:       false,
			Name:           model.Name,
			RepositoryLink: repositoryLink,
			Size:           model.Size,
			Alias:          alias,
			OwnerName:      ownerName,
		})
	}
	return pretrainModelList
}

func getModelLocalLink(model *models.AiModelManage) string {
	index := strings.Index(model.Path, "/")
	key := model.Path[index+1:]
	url, _ := storage.GetObsCreateSignedUrlByBucketAndKey(setting.Bucket, key)
	return url
}

func GetCloudBrainDataSetInfo(task *models.Cloudbrain) []*models.DatasetDownload {
	datasetDownload := getCloudBrainDatasetInfo4Local(task.Uuid, task.DatasetName, true)

	//非虎鲸的任务返回本地地址
	if task.Type != models.TypeC2Net {
		return datasetDownload
	}
	//虎鲸的任务需要返回调度后的地址
	datasetObsUrlList := make([]entity.NotebookDataset, 0)
	_ = json.Unmarshal([]byte(task.DataUrl), &datasetObsUrlList)

	return datasetDownload
}

func getCloudBrainDatasetInfo4Local(uuid string, datasetname string, isNeedDown bool) []*models.DatasetDownload {
	datasetDownload := make([]*models.DatasetDownload, 0)
	if len(uuid) == 0 {
		return datasetDownload
	}
	uuidList := strings.Split(uuid, ";")
	datasetnameList := strings.Split(datasetname, ";")
	for i, uuidStr := range uuidList {
		name := ""
		link := ""
		url := ""
		isDelete := false
		var size int64
		attachment, err := models.GetAttachmentByUUID(uuidStr)
		if err != nil {
			log.Error("GetAttachmentByUUID failed:%v", err.Error())
			if len(datasetnameList) <= i || len(datasetname) == 0 {
				continue
			}
			name = datasetnameList[i]
			isDelete = true
		} else {
			name = attachment.Name
			size = attachment.Size
			dataset, err := models.GetDatasetByID(attachment.DatasetID)
			if err != nil {
				log.Error("GetDatasetByID failed:%v", err.Error())
			} else {
				repo, err := models.GetRepositoryByID(dataset.RepoID)
				if err != nil {
					log.Error("GetRepositoryByID failed:%v", err.Error())
				} else {
					link = repo.Link() + "/datasets"
				}
			}
			if isNeedDown {
				url = attachment.S3DownloadURL()
			}
		}

		datasetDownload = append(datasetDownload, &models.DatasetDownload{
			DatasetName:         name,
			DatasetDownloadLink: url,
			RepositoryLink:      link,
			IsDelete:            isDelete,
			UUID:                uuidStr,
			Size:                size,
		})
	}
	log.Info("dataset length=" + fmt.Sprint(len(datasetDownload)))
	return datasetDownload
}
func GetCloudBrainDataSetRegistryInfo(task *models.Cloudbrain) []*models.DatasetRegistryDownload {
	datasetDownload := getCloudBrainDatasetInfoRegistry4Local(task.Uuid, task.DatasetName, task.DatasetAlias)
	return datasetDownload
}

func getCloudBrainDatasetInfoRegistry4Local(uuid string, datasetname string, datasetAlias string) []*models.DatasetRegistryDownload {
	datasetDownload := make([]*models.DatasetRegistryDownload, 0)
	if len(uuid) == 0 {
		return datasetDownload
	}
	uuidList := strings.Split(uuid, ";")
	datasetnameList := strings.Split(datasetname, ";")
	datasetAliasList := strings.Split(datasetAlias, ";")
	for i, uuidStr := range uuidList {
		name := ""
		alias := ""
		ownerName := ""
		isDelete := false
		var size int64
		dataset, err := models.GetDatasetRegistryByID(uuidStr)
		// attachment, err := models.GetAttachmentByUUID(uuidStr)
		if err != nil {
			log.Error("GetDatasetRegistryByID failed:%v", err)
			if len(datasetAliasList) > i {
				alias = datasetAliasList[i]
			}
			if len(datasetnameList) > i {
				name = datasetnameList[i]
			}
			if alias == "" {
				alias = name
			}
			isDelete = true
		} else {
			alias = dataset.DisplayName()
			name = dataset.Name
			size = dataset.Size
			dataset.GetOwner()
			if dataset.Owner != nil {
				ownerName = dataset.Owner.Name
			}
		}

		datasetDownload = append(datasetDownload, &models.DatasetRegistryDownload{
			DatasetName:  name,
			IsDelete:     isDelete,
			UUID:         uuidStr,
			Size:         size,
			OwnerName:    ownerName,
			DatasetAlias: alias,
		})
	}
	log.Info("dataset length=" + fmt.Sprint(len(datasetDownload)))
	return datasetDownload
}

// 根据实际调度的智算中心修正规格
func correctAITaskSpec(task *models.Cloudbrain) {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:%v", combinedErr)
		}
	}()
	if task.AiCenter == "" {
		return
	}
	s := strings.Split(task.AiCenter, "+")
	if len(s) < 2 {
		return
	}

	realCenterCode := s[0]
	if realCenterCode == "" {
		return
	}

	oldSpec, err := models.GetCloudbrainSpecByID(task.ID)
	if err != nil {
		log.Error("correctAITaskSpec GetCloudbrainSpecByID err.taskId=%d err=%v", task.ID, err)
		return
	}
	if oldSpec == nil {
		log.Error("correctAITaskSpec GetCloudbrainSpecByID spec is empty.taskId=%d ", task.ID)
		return
	}
	if oldSpec.AiCenterCode == realCenterCode && oldSpec.QueueCode == task.QueueCode {
		return
	}
	//所属资源池队列不一样时才需要处理
	r, err := models.FindSpecs(models.FindSpecsOptions{
		SourceSpecId: oldSpec.SourceSpecId,
		AiCenterCode: realCenterCode,
		QueueCode:    task.QueueCode,
		RequestAll:   true,
	})
	if err != nil {
		log.Error("correctAITaskSpec FindSpecs err.taskId=%d err=%v", task.ID, err)
		return
	}
	if r == nil || len(r) == 0 {
		log.Error("correctAITaskSpec FindSpecs 0.taskId=%d ", task.ID)
		return
	}
	var newestQueue = r[0]
	if len(r) > 1 {
		for _, s := range r {
			if s.QueueId > newestQueue.QueueId {
				newestQueue = s
			}
		}
	}
	n, err := models.UpdateCloudbrainSpec(task.ID, newestQueue)
	if err == nil && n > 0 {
		log.Info("correctAITaskSpec success,taskId=%d oldCenter=%s realCenter=%s", task.ID, oldSpec.AiCenterCode, realCenterCode)
	}
}

func getModelContainerLink(dataUrl string, modelName string) string {
	if dataUrl == "" {
		return ""
	}
	datasetObsUrlList := make([]entity.NotebookDataset, 0)
	_ = json.Unmarshal([]byte(dataUrl), &datasetObsUrlList)
	for _, datasetObs := range datasetObsUrlList {
		if datasetObs.DatasetName == modelName {
			return datasetObs.DatasetUrl
		}
	}
	return ""
}
