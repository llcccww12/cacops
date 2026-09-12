package dync_parser

import (
	"errors"
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
)

func init() {
	RegisterParser(entity.DynConfigTypeRepo, &RepoParser{})
}

type RepoParser struct {
}

func (d *RepoParser) ParseObject(config entity.DynConfig) (interface{}, error) {
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

	return d.parseSingleRepo(str)
}

func (d *RepoParser) ParseList(config entity.DynConfig) (interface{}, error) {
	var repoConfigs []*entity.RepoConfig4Show
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
		repoConfig, err := d.parseSingleRepo(str)
		if err != nil {
			log.Error("Parse dataset config error. err=%v", err)
			continue
		}
		repoConfigs = append(repoConfigs, repoConfig)
	}
	return repoConfigs, nil
}

func (d *RepoParser) parseSingleRepo(name string) (*entity.RepoConfig4Show, error) {
	trimmed := strings.Trim(name, "/")

	parts := strings.Split(trimmed, "/")
	if len(parts) != 2 {
		return nil, errors.New("repo name format error")
	}
	repo, err := models.GetRepositoryByOwnerAndName(parts[0], parts[1])
	if err != nil {
		log.Error("Get repo error. err=%v", err)
		return nil, errors.New("repo name error")
	}
	// 组装数据集配置
	repoConfig := &entity.RepoConfig4Show{
		Name:        repo.DisplayName(),
		Description: repo.Description,
		RepoLink:    repo.FullName(),
		Avatar:      repo.RelAvatarLink(),
		NumWatches:  repo.NumWatches,
		NumStars:    repo.NumStars,
		NumForks:    repo.NumForks,
	}
	return repoConfig, nil
}
