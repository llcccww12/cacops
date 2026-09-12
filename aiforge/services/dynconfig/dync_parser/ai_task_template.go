package dync_parser

import (
	"errors"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
)

func init() {
	RegisterParser(entity.DynConfigTypeAITaskTemplate, &AITaskTemplateParser{})
}

type AITaskTemplateParser struct {
}

func (d *AITaskTemplateParser) ParseObject(config entity.DynConfig) (interface{}, error) {
	if config.Value == nil {
		return nil, nil
	}
	// 类型断言
	m, ok := config.Value.(map[string]interface{})
	if !ok {
		return nil, errors.New("value is not a map")
	}

	val, exists := m["id"]
	if !exists {
		return "", errors.New("ai task template config format error")
	}
	str, ok := val.(string)
	if !ok {
		return "", errors.New("ai task template id is not a string")
	}

	return d.parseSingleAITaskTemplate(str)
}

func (d *AITaskTemplateParser) ParseList(config entity.DynConfig) (interface{}, error) {
	var templateConfigs []*entity.AITaskTemplateConfig4Show
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
		val, exists := newList[i]["id"]
		if !exists {
			continue
		}
		str, ok := val.(string)
		if !ok {
			continue
		}
		templateConfig, err := d.parseSingleAITaskTemplate(str)
		if err != nil {
			log.Error("Parse dataset config error. err=%v", err)
			continue
		}
		templateConfigs = append(templateConfigs, templateConfig)
	}
	return templateConfigs, nil
}

func (d *AITaskTemplateParser) parseSingleAITaskTemplate(id string) (*entity.AITaskTemplateConfig4Show, error) {
	template, err := models.GetAITaskTemplateByID(id)
	if err != nil {
		log.Error("Get template error.id=%s, err=%v", id, err)
		return nil, errors.New("template id error")
	}
	template.LoadAtrribute(0)
	owner := template.GetOwner()

	if owner == nil {
		log.Error("Get template owner error.id=%s, err=%v", id, err)
		return nil, err
	}

	templateConfig := &entity.AITaskTemplateConfig4Show{
		Name:          template.Name,
		ID:            template.ID,
		Tags:          template.Tags,
		Recommend:     template.Recommend,
		ComputeSource: template.ComputeSource,
		JobType:       template.JobType,
		Owner:         owner.ToFrontFormat(),
		DatasetLists:  template.DatasetList,
		RepoName:      template.RepoName,
		RepoOwnerName: template.RepoOwnerName,
		ModelLists:    template.ModelList,
		ImageID:       template.ImageID,
		ImageName:     template.ImageName,
		ImageUrl:      template.ImageUrl,
	}
	return templateConfig, nil
}
