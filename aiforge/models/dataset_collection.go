package models

import (
	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
)

type DatasetCollection struct {
	ID          int64              `xorm:"pk autoincr"`
	UserID      int64              `xorm:"UNIQUE(s)"`
	DatasetID   string             `xorm:"uuid UNIQUE(s)"`
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
}

// StarRepo or unstar repository.
func CollectionDataset(userID int64, datasetID string, collect bool) error {
	sess := x.NewSession()
	defer sess.Close()

	if err := sess.Begin(); err != nil {
		return err
	}

	if collect {
		if isDatasetCollecting(sess, userID, datasetID) {
			return nil
		}

		if _, err := sess.Insert(&DatasetCollection{UserID: userID, DatasetID: datasetID}); err != nil {
			return err
		}
		if _, err := sess.Exec("UPDATE `dataset_registry` SET num_collections = num_collections + 1 WHERE id = ?", datasetID); err != nil {
			return err
		}
		if _, err := sess.Exec("UPDATE `user` SET num_dataset_stars = num_dataset_stars + 1 WHERE id = ?", userID); err != nil {
			return err
		}
	} else {
		if !isDatasetCollecting(sess, userID, datasetID) {
			return nil
		}

		if _, err := sess.Delete(&DatasetCollection{0, userID, datasetID, 0}); err != nil {
			return err
		}
		if _, err := sess.Exec("UPDATE `dataset_registry` SET num_collections = num_collections - 1 WHERE id = ?", datasetID); err != nil {
			return err
		}
		if _, err := sess.Exec("UPDATE `user` SET num_dataset_stars = num_dataset_stars - 1 WHERE id = ?", userID); err != nil {
			return err
		}
	}

	return sess.Commit()
}

func isDatasetCollecting(e Engine, userID int64, datasetID string) bool {
	has, _ := e.Get(&DatasetCollection{0, userID, datasetID, 0})
	return has
}

func IsDatasetCollecting(userID int64, datasetID string) bool {
	return isDatasetCollecting(x, userID, datasetID)
}

func GetDatasetCollectionsMap(userId int64, datasetIds []string) (map[string]bool, error) {
	if len(datasetIds) == 0 {
		return nil, nil
	}
	var collections []DatasetCollection
	err := x.Where(
		builder.NewCond().
			And(builder.Eq{"user_id": userId}).
			And(builder.In("dataset_id", datasetIds))).
		Find(&collections)
	if err != nil {
		return nil, err
	}
	collectionMap := make(map[string]bool)
	for _, datasetId := range datasetIds {
		for _, collection := range collections {
			if collection.DatasetID == datasetId {
				collectionMap[datasetId] = true
				break
			}
		}
		if _, exists := collectionMap[datasetId]; !exists {
			collectionMap[datasetId] = false
		}
	}
	return collectionMap, nil
}

func GetCollectionUserIdList(datasetId string) ([]int64, error) {
	var ids = make([]int64, 0)
	err := x.Table("dataset_collection").Cols("user_id").Where("dataset_id = ?", datasetId).Find(&ids)
	if err != nil {
		return nil, err
	}
	return ids, nil
}
