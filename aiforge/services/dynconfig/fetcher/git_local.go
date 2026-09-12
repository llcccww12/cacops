//implementation of git local fetcher
package fetcher

import (
	"encoding/json"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/log"
)

type GitLocalFetcher struct {
	ownerName  string
	repoName   string
	branchName string
}

func NewGitLocalFetcher(ownerName, repoName, branchName string) *GitLocalFetcher {
	return &GitLocalFetcher{
		ownerName:  ownerName,
		repoName:   repoName,
		branchName: branchName,
	}
}

func (g *GitLocalFetcher) Fetch(path string) (*entity.DynConfig, error) {
	repo, err := models.GetRepositoryByOwnerAndName(g.ownerName, g.repoName)
	if err != nil {
		log.Error("GetRepositoryByOwnerAndAlias failed, error=%v", err)
		if models.IsErrRepoNotExist(err) {
			return nil, models.ErrNotExist{}
		}
		return nil, err
	}
	repoFile, err := models.ReadLatestFileInRepo(repo.OwnerName, repo.Name, g.branchName, path)
	if err != nil {
		log.Error("ReadLatestFileInRepo failed, error=%v", err)
		if git.IsErrNotExist(err) {
			return nil, models.ErrNotExist{}
		}
		return nil, err
	}
	var config entity.DynConfig
	if err := json.Unmarshal(repoFile.Content, &config); err != nil || config.ConfigType == "" || config.Value == nil {
		config.ConfigType = entity.DynConfigTypeRaw
		config.ValueType = entity.ValueTypeObject
		config.Value = string(repoFile.Content)
	}
	return &config, nil
}
