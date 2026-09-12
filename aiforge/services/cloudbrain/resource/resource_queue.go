package resource

import (
	"encoding/json"
	"fmt"
	"strings"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/grampus"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
)

func AddResourceQueue(req models.ResourceQueueReq) error {
	if _, err := models.InsertResourceQueue(req.ToDTO()); err != nil {
		return err
	}
	return nil
}

func UpdateResourceQueue(queueId int64, req models.ResourceQueueReq) error {
	if _, err := models.UpdateResourceCardsTotalNumAndInternetStatus(queueId, models.ResourceQueue{
		CardsTotalNum: req.CardsTotalNum,
		QueueType:     req.QueueType,
		QueueName:     req.QueueName,
		Remark:        req.Remark,
		HasInternet:   req.HasInternet,
	}, req.IsAvailable); err != nil {
		return err
	}
	return nil
}

func GetResourceQueueList(opts models.SearchResourceQueueOptions) (*models.ResourceQueueListRes, error) {
	n, r, err := models.SearchResourceQueue(opts)
	if err != nil {
		return nil, err
	}

	return models.NewResourceQueueListRes(n, r), nil
}

func GetResourceQueueCodes(opts models.GetQueueCodesOptions) ([]*models.ResourceQueueCodesRes, error) {
	r, err := models.GetResourceQueueCodes(opts)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func GetResourceAiCenters() ([]models.ResourceAiCenterRes, error) {
	r, err := models.GetResourceAiCenters()
	if err != nil {
		return nil, err
	}

	return r, nil
}

func filterGrampusImage(all []models.GrampusImage) []models.GrampusImage {
	result := make([]models.GrampusImage, 0)
	for _, tmp := range all {
		computeResource := getComputeResourceByProcessType(tmp.ProcessorType)
		if computeResource == "GCU" || computeResource == "MLU" ||
			computeResource == "ILUVATAR-GPGPU" || computeResource == "METAX-GPGPU" ||
			computeResource == "GPU" {
			result = append(result, tmp)
		}
	}
	return result
}

func getComputeResourceByProcessType(processType string) string {
	tail := strings.LastIndex(processType, "/")
	if tail > 0 {
		return strings.ToUpper(processType[tail+1:])
	}
	return strings.ToUpper(processType)
}

// SyncGrampusImage 这个函数主要是通过hu_jing_id来更新数据
func SyncGrampusImage(doerId int64) error {
	allImage, err := grampus.GetAllBaseImages()
	if err == nil {
		log.Info("Query image from grampus, length=" + fmt.Sprint(len(allImage.Infos)))
		dbAllImage, err := models.GetGrampusAllBaseImage()
		if err == nil {
			log.Info("start to deal db image update.")
			grampusImageList := allImage.Infos
			log.Info("sync grampus image length=." + fmt.Sprint(len(grampusImageList)))
			grampusMap := make(map[string]models.GrampusImage, 0)
			for _, image := range grampusImageList {
				grampusMap[image.ID] = image
			}
			dbMap := make(map[string]*models.Image, 0)
			for _, image := range dbAllImage {
				dbMap[image.HuJingId] = image
			}
			insertList := make([]models.GrampusImage, 0)
			updateList := make([]models.GrampusImage, 0)
			updateDbList := make([]*models.Image, 0)
			for _, image := range grampusImageList {
				if dbImage, ok := dbMap[image.ID]; !ok {
					insertList = append(insertList, image)
				} else {
					updateList = append(updateList, image)
					updateDbList = append(updateDbList, dbImage)
				}
			}
			deleteList := make([]*models.Image, 0)
			for _, image := range dbAllImage {
				if _, ok := grampusMap[image.HuJingId]; !ok {
					deleteList = append(deleteList, image)
				}
			}
			models.SyncGrampusAllBaseImageToDb(insertList, updateList, updateDbList, deleteList, doerId)
		} else {
			log.Error("failed query image db.error=" + err.Error())
		}
	} else {
		log.Error("failed query image grampus.error=" + err.Error())
	}
	return err
}

func SyncGrampusQueue(doerId int64) error {
	var r []models.GrampusResourceQueue
	var err error
	if setting.UseNewSyncQueueAPI {
		r, err = grampus.GetNewResourceQueue()
		jsostr, _ := json.Marshal(r)
		log.Info("SyncGrampusQueueNew result , length=%d, data=%s", len(r), string(jsostr))
		r2, _ := grampus.GetResourceQueue()
		jsostr2, _ := json.Marshal(r2)
		log.Info("SyncGrampusQueue result , length=%d, data=%s", len(r2), string(jsostr2))
	} else {
		r, err = grampus.GetResourceQueue()
		jsostr, _ := json.Marshal(r)
		log.Info("SyncGrampusQueue result , length=%d, data=%s", len(r), string(jsostr))

	}

	if err != nil {
		log.Error("Get grampus resource queue failed.err=%v", err)
		return err
	}

	queueUpdateList := make([]models.ResourceQueue, 0)
	queueInsertList := make([]models.ResourceQueue, 0)
	existIds := make([]int64, 0)

	for _, queue := range r {
		computeResource := queue.ComputeResource
		accCardType := queue.AccCardType
		oldQueue, err := models.GetResourceQueue(&models.ResourceQueue{
			Cluster:         models.C2NetCluster,
			AiCenterCode:    queue.AiCenterCode,
			ComputeResource: computeResource,
			AccCardType:     accCardType,
			QueueCode:       queue.QueueCode,
		})
		if err != nil {
			return err
		}

		hasInternet := queue.HasInternet
		if oldQueue == nil {
			queueInsertList = append(queueInsertList, models.ResourceQueue{
				Cluster:             models.C2NetCluster,
				AiCenterCode:        queue.AiCenterCode,
				AiCenterName:        queue.AiCenterName,
				ComputeResource:     computeResource,
				AccCardType:         accCardType,
				IsAutomaticSync:     true,
				HasInternet:         hasInternet,
				CreatedBy:           doerId,
				UpdatedBy:           doerId,
				QueueCode:           queue.QueueCode,
				QueueName:           queue.QueueName,
				QueueType:           queue.QueueType,
				IsAvailable:         true,
				EnableVisualization: queue.EnableVisualization,
				CardsTotalNum:       queue.CardsTotalNum,
			})
		} else {
			existIds = append(existIds, oldQueue.ID)
			queueUpdateList = append(queueUpdateList, models.ResourceQueue{
				ID:                  oldQueue.ID,
				ComputeResource:     computeResource,
				AiCenterName:        queue.AiCenterName,
				AccCardType:         accCardType,
				UpdatedBy:           doerId,
				HasInternet:         hasInternet,
				QueueName:           queue.QueueName,
				QueueType:           queue.QueueType,
				IsAvailable:         true,
				EnableVisualization: queue.EnableVisualization,
				CardsTotalNum:       queue.CardsTotalNum,
			})
		}

	}
	return models.SyncGrampusQueues(queueUpdateList, queueInsertList, existIds)
}

func SyncGrampusQueueAndSpecs() {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:", combinedErr)
		}
	}()
	log.Info("start to sync grampus centers queues and specs")
	SyncGrampusQueue(0)
	SyncGrampusSpecs(0)
	SyncGrampusAICenter()
	log.Info("sync grampus centers queues and specs finished")
	SyncGrampusSpecPools()
	log.Info("sync grampus queues cardNums finished")
}
