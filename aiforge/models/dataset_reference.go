package models

import (
	"strconv"
	"strings"

	"code.gitea.io/gitea/modules/timeutil"
)

type DatasetReference struct {
	ID          int64              `xorm:"pk autoincr"`
	RepoID      int64              `xorm:"INDEX unique"`
	DatasetID   string             `xorm:"TEXT"`
	CreatedUnix timeutil.TimeStamp `xorm:"INDEX created"`
}

func GetDatasetIdsByRepoID(repoID int64) []int64 {
	var datasets []int64
	var datasetIds []string
	_ = x.Table("dataset_reference").Where("repo_id=?", repoID).
		Cols("dataset_reference.dataset_id").Find(&datasetIds)
	if len(datasetIds) > 0 {
		for _, datasetIdStr := range strings.Split(datasetIds[0], ",") {
			datasetId, err := strconv.ParseInt(datasetIdStr, 10, 64)
			if err != nil {
				continue
			}
			datasets = append(datasets, datasetId)
		}
	}

	return datasets
}

func HasReferenceDataset(repoID int64) bool {

	var datasetIds []string
	_ = x.Table("dataset_reference").Where("repo_id=?", repoID).
		Cols("dataset_reference.dataset_id").Find(&datasetIds)
	return len(datasetIds) > 0
}

func getReferenceDatasetStr(repoID int64) string {

	var datasetIds []string
	_ = x.Table("dataset_reference").Where("repo_id=?", repoID).
		Cols("dataset_reference.dataset_id").Find(&datasetIds)
	if len(datasetIds) > 0 {
		return datasetIds[0]
	}
	return ""
}

func DeleteReferenceDatasetIdsByRepoID(repoID int64) error {

	_, err := x.Exec("delete from dataset_reference where repo_id=?", repoID)
	return err
}

func NewDatasetIdsByRepoID(repoID int64, datasetIds []int64) error {
	if len(datasetIds) == 0 { //关联数据集数组为空
		DeleteReferenceDatasetIdsByRepoID(repoID)
	}
	var datasetsStrArray []string
	for _, datasetId := range datasetIds {
		datasetsStrArray = append(datasetsStrArray, strconv.FormatInt(datasetId, 10))
	}

	newDatasetStr := strings.Join(datasetsStrArray, ",")
	oldDatasetStr := getReferenceDatasetStr(repoID)
	if newDatasetStr == oldDatasetStr { //关联数据集无变化，不需要处理
		return nil
	}
	if oldDatasetStr != "" { //已经存在关联数据集
		_, err := x.Exec("update dataset_reference set dataset_id=? where repo_id=?", newDatasetStr, repoID)

		return err
	} else {
		datasetReference := DatasetReference{
			DatasetID: newDatasetStr,
			RepoID:    repoID,
		}

		_, err := x.Insert(datasetReference)
		return err
	}

}
