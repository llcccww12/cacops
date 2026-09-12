package models

import "code.gitea.io/gitea/modules/timeutil"

type DatasetStar struct {
	ID          int64              `xorm:"pk autoincr"`
	UID         int64              `xorm:"UNIQUE(s)"`
	DatasetID   int64              `xorm:"UNIQUE(s)"`
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
}

// StarRepo or unstar repository.
func StarDataset(userID, datasetID int64, star bool) error {
	sess := x.NewSession()
	defer sess.Close()

	if err := sess.Begin(); err != nil {
		return err
	}

	if star {
		if isDatasetStaring(sess, userID, datasetID) {
			return nil
		}

		if _, err := sess.Insert(&DatasetStar{UID: userID, DatasetID: datasetID}); err != nil {
			return err
		}
		if _, err := sess.Exec("UPDATE `dataset` SET num_stars = num_stars + 1 WHERE id = ?", datasetID); err != nil {
			return err
		}
		if _, err := sess.Exec("UPDATE `user` SET num_dataset_stars = num_dataset_stars + 1 WHERE id = ?", userID); err != nil {
			return err
		}
	} else {
		if !isDatasetStaring(sess, userID, datasetID) {
			return nil
		}

		if _, err := sess.Delete(&DatasetStar{0, userID, datasetID, 0}); err != nil {
			return err
		}
		if _, err := sess.Exec("UPDATE `dataset` SET num_stars = num_stars - 1 WHERE id = ?", datasetID); err != nil {
			return err
		}
		if _, err := sess.Exec("UPDATE `user` SET num_dataset_stars = num_dataset_stars - 1 WHERE id = ?", userID); err != nil {
			return err
		}
	}

	return sess.Commit()
}

func IsDatasetStaringByRepoId(userID, repoID int64) bool {
	dataset, _ := GetDatasetByRepo(&Repository{ID: repoID})
	if dataset == nil {
		return false
	}
	return isDatasetStaring(x, userID, dataset.ID)
}

func IsDatasetStaring(userID, datasetID int64) bool {
	return isDatasetStaring(x, userID, datasetID)

}

func isDatasetStaring(e Engine, userID, datasetID int64) bool {
	has, _ := e.Get(&DatasetStar{0, userID, datasetID, 0})
	return has
}

func GetDatasetIdsStarByUser(userID int64) []int64 {
	var datasets []int64
	_ = x.Table("dataset_star").Where("uid=?", userID).
		Cols("dataset_star.dataset_id").Find(&datasets)
	return datasets
}

func GetStaredUsersByDatasetID(datasetID int64) []int64 {
	var datasets []int64
	_ = x.Table("dataset_star").Where("dataset_id=?", datasetID).
		Cols("dataset_star.uid").Find(&datasets)
	return datasets
}

func HandleOldDatasetCollections(oldDatasetID int64, newDatasetId string, userId int64) error {
	_, err := x.Exec("DELETE from  dataset_collection WHERE dataset_id = ?", newDatasetId)
	if err != nil {
		return err
	}
	_, err = x.Exec("INSERT INTO dataset_collection (user_id, dataset_id, created_unix) SELECT uid, ?, created_unix FROM dataset_star WHERE dataset_id = ?", newDatasetId, oldDatasetID)
	if err != nil {
		return err
	}
	_, err = x.Exec("UPDATE dataset_registry SET num_collections = (SELECT COUNT(*) FROM dataset_collection WHERE dataset_id = ?) WHERE id = ? ", newDatasetId, newDatasetId)
	if err != nil {
		return err
	}

	var ids = make([]int64, 0)
	err = x.Table("dataset_collection").Cols("user_id").Where("dataset_id = ?", newDatasetId).Find(&ids)
	if err != nil {
		return err
	}
	for _, id := range ids {
		_, err = x.Exec("UPDATE public.user SET num_dataset_stars = (SELECT COUNT(*) FROM dataset_collection WHERE user_id = ?) WHERE id = ? ", id, id)
		if err != nil {
			return err
		}
	}

	return nil

}
