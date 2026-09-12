package models

import (
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
)

type ResourceExclusivePool struct {
	ID              int64 `xorm:"pk autoincr"`
	SceneId         int64
	OrgName         string
	JobType         string
	Cluster         string
	QueueId         int64
	ComputeResource string
	CreatedTime     timeutil.TimeStamp `xorm:"created"`
	CreatedBy       int64
	UpdatedTime     timeutil.TimeStamp `xorm:"updated"`
	UpdatedBy       int64
}

func FindExclusivePools() ([]*ResourceExclusivePool, error) {
	sq := make([]*ResourceExclusivePool, 0)

	err := x.Find(&sq)
	if err != nil {
		return nil, err
	}
	return sq, nil
}

func InsertExclusivePools(queue []ResourceExclusivePool) (int64, error) {
	return x.Insert(&queue)
}

func IsQueueInExclusivePool(queueId int64) bool {
	n, _ := x.In("queue_id", queueId).Count(&ResourceExclusivePool{})
	return n > 0
}

func FindExclusiveQueueIds() []int64 {
	existsIds := make([]int64, 0)
	err := x.Table("resource_exclusive_pool").Distinct("queue_id").Find(&existsIds)
	if err != nil {
		log.Error("FindQueuesExclusiveMap err.%v", err)
		return existsIds
	}

	return existsIds
}

func FindQueuesExclusiveMap() map[int64]string {
	resultMap := make(map[int64]string, 0)

	existsIds := FindExclusiveQueueIds()

	for _, id := range existsIds {
		resultMap[id] = ""
	}
	return resultMap
}
