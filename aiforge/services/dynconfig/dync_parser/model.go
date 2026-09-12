package dync_parser

import (
	"errors"
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
)

func init() {
	RegisterParser(entity.DynConfigTypeModel, &ModelParser{})
}

type ModelParser struct {
}

func (d *ModelParser) ParseObject(config entity.DynConfig) (interface{}, error) {
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
		return "", errors.New("model config format error")
	}
	str, ok := val.(string)
	if !ok {
		return "", errors.New("model name is not a string")
	}

	return d.parseSingleModel(str)
}

func (d *ModelParser) ParseList(config entity.DynConfig) (interface{}, error) {
	var modelConfigs []*entity.ModelConfig4Show
	list, ok := config.Value.([]interface{})
	if !ok {
		log.Error("Parse model config error. config=%v", config)
		return nil, errors.New("value is not a list")
	}

	var newList []map[string]interface{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		if !ok {
			log.Error("Parse model config error. item=%v", item)
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
		modelConfig, err := d.parseSingleModel(str)
		if err != nil {
			log.Error("Parse model config error. err=%v", err)
			continue
		}
		modelConfigs = append(modelConfigs, modelConfig)
	}
	return modelConfigs, nil
}

func (d *ModelParser) parseSingleModel(name string) (*entity.ModelConfig4Show, error) {
	trimmed := strings.Trim(name, "/")

	parts := strings.Split(trimmed, "/")
	if len(parts) != 2 {
		log.Error("Aimodel name format error. name=%s", name)
		return nil, errors.New("aimodel name format error")
	}

	owner, err := models.GetUserByName(parts[0])
	if err != nil {
		log.Error("Get aimodel owner error.name=%s, err=%v", name, err)
		return nil, errors.New("aimodel owner error")
	}

	aimodel, err := models.GetAiModelByByOwnerAndName(owner.ID, parts[1])
	if err != nil {
		log.Error("Get aimodel error.name=%s, err=%v", name, err)
		return nil, errors.New("aimodel name error")
	}

	isRecommend := false
	if aimodel.Recommend == 1 {
		isRecommend = true
	}

	// 组装数据集配置
	aimodelConfig := &entity.ModelConfig4Show{
		ID:           aimodel.ID,
		Name:         aimodel.Name,
		Alias:        aimodel.Alias,
		ExternalName: aimodel.ExternalName,
		Label:        aimodel.Label,
		OwnerName:    owner.Name,
		Avatar:       owner.AvatarLink(),
		Recommend:    isRecommend,
		ModelType:    aimodel.ModelType,
		Engine:       aimodel.Engine,
	}
	return aimodelConfig, nil
}
