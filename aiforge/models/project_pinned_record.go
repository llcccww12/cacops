package models

import (
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
)

// RepositoryPinnedRecord 用户与项目置顶
// 可能存在置顶时间存在，但是用户与项目解除的关系，若用户重新加入，置顶关系仍然存在
type RepositoryPinnedRecord struct {
	ID           int64              `xorm:"pk autoincr" json:"id"`
	UserID       int64              `xorm:"INDEX NOT NULL" json:"userId"`         // 用户id
	RepositoryID int64              `xorm:"INDEX NOT NULL" json:"repository_id"`  // 仓库的id
	PinnedUnix   timeutil.TimeStamp `xorm:"INDEX NOT NULL" json:"pinned_unix"`    // 置顶时间 0-取消置顶
	CreatedUnix  timeutil.TimeStamp `xorm:"INDEX created" json:"created_unix"`    // 创建时间
	CancelUnix   timeutil.TimeStamp `xorm:"INDEX cancel_unix" json:"cancel_unix"` // 取消时间
}

func (p *RepositoryPinnedRecord) TableName() string {
	return "repository_pinned_record"
}

// GetPinnedRecordByUserId 查询用户置顶记录
func GetPinnedRecordByUserId(userId int64) (map[int64]*RepositoryPinnedRecord, error) {
	var (
		records   = make([]*RepositoryPinnedRecord, 0) // 注意：这里不是指针)
		recordMap = make(map[int64]*RepositoryPinnedRecord)
	)

	if err := x.Table((&RepositoryPinnedRecord{}).TableName()).
		Where("user_id = ? and pinned_unix > 0", userId).
		Find(&records); err != nil {
		log.Error("GetPinnedRecordByUserId query error[%v]", err)
		return nil, err
	}

	for _, record := range records {
		recordMap[record.RepositoryID] = record
	}

	return recordMap, nil // 返回指针
}

// GetPinnedRecord 查询置顶记录
func GetPinnedRecord(userId, repositoryId int64) (*RepositoryPinnedRecord, error) {
	var record RepositoryPinnedRecord // 注意：这里不是指针

	has, err := x.Table(record.TableName()).
		Where("user_id = ? AND repository_id = ?", userId, repositoryId).
		Get(&record) // 直接传 &record，因为 record 现在是值类型

	if err != nil {
		log.Error("GetPinnedRecord query error[%v]", err)
		return nil, err
	}

	if !has {
		return nil, nil // 记录不存在
	}

	return &record, nil // 返回指针
}

// CreatePinnedRecord 创建置顶记录
func CreatePinnedRecord(record *RepositoryPinnedRecord) error {
	_, err := x.Table(record.TableName()).Insert(record)
	if err != nil {
		log.Info("CreatePinnedRecord insert error[%v]", err.Error())
		return err
	}

	return nil
}

// UpdatePinnedRecord 更新置顶时间
func UpdatePinnedRecord(record *RepositoryPinnedRecord) error {
	_, err := x.Table(record.TableName()).Where("user_id = ? and repository_id= ?", record.UserID, record.RepositoryID).Update(map[string]interface{}{
		"pinned_unix": record.PinnedUnix,
	})
	if err != nil {
		log.Info("UpdatePinnedRecord error[%v]", err.Error())
		return err
	}

	return nil
}

// CancelPinnedRecord 删除置顶时间
func CancelPinnedRecord(record *RepositoryPinnedRecord) error {
	_, err := x.Table(record.TableName()).Where("user_id = ? and repository_id= ?", record.UserID, record.RepositoryID).Update(map[string]interface{}{
		"pinned_unix": 0,
		"cancel_unix": timeutil.TimeStampNow(),
	})
	if err != nil {
		log.Info("UpdatePinnedRecord error[%v]", err.Error())
		return err
	}

	return nil
}
