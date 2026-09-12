package dync_parser

import (
	"errors"
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
)

func init() {
	RegisterParser(entity.DynConfigTypeDataset, &DatasetParser{})
}

type DatasetParser struct {
}

func (d *DatasetParser) ParseObject(config entity.DynConfig) (interface{}, error) {
	if config.Value == nil {
		return nil, nil
	}
	// 类型断言
	m, ok := config.Value.(map[string]interface{})
	if !ok {
		return nil, errors.New("value is not a map")
	}

	val, exists := m["name"]
	if !exists {
		return "", errors.New("dataset config format error")
	}
	str, ok := val.(string)
	if !ok {
		return "", errors.New("dataset name is not a string")
	}

	return d.parseSingleDataset(str)
}

func (d *DatasetParser) ParseList(config entity.DynConfig) (interface{}, error) {
	var datasetConfigs []*entity.DatasetConfig4Show
	list, ok := config.Value.([]interface{})
	if !ok {
		return nil, errors.New("value is not a list")
	}

	var newList []map[string]interface{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		if !ok {
			return nil, errors.New("list item is not a map")
		}
		newList = append(newList, m)
	}

	for i := 0; i < len(newList); i++ {
		val, exists := newList[i]["name"]
		if !exists {
			continue
		}
		str, ok := val.(string)
		if !ok {
			continue
		}
		datasetConfig, err := d.parseSingleDataset(str)
		if err != nil {
			log.Error("Parse dataset config error. err=%v", err)
			continue
		}
		datasetConfigs = append(datasetConfigs, datasetConfig)
	}
	return datasetConfigs, nil
}

func (d *DatasetParser) parseSingleDataset(name string) (*entity.DatasetConfig4Show, error) {
	trimmed := strings.Trim(name, "/")

	parts := strings.Split(trimmed, "/")
	if len(parts) != 2 {
		log.Error("Dataset name format error. name=%s", name)
		return nil, errors.New("dataset name format error")
	}
	dataset, err := models.GetDatasetRegistryByOwnerNameAndDatasetName(parts[0], parts[1])
	if err != nil {
		log.Error("Get dataset registry error.name=%s, err=%v", name, err)
		return nil, errors.New("dataset name error")
	}
	// 组装数据集配置
	datasetConfig := &entity.DatasetConfig4Show{
		Name:      dataset.Name,
		ID:        dataset.ID,
		Tags:      dataset.Tags,
		Tasks:     dataset.Tasks,
		License:   dataset.License,
		OwnerName: parts[0],
		Recommend: dataset.Recommend,
		Alias:     dataset.Alias,
	}
	return datasetConfig, nil
}
