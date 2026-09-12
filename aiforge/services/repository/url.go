package repository

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/setting"
	"errors"
	"net/url"
	"strings"
)

func FindRepoByUrl(url string) (*models.Repository, error) {
	ownerName, repoName := parseOpenIUrl(url)
	if ownerName == "" || repoName == "" {
		return nil, errors.New("tech.incorrect_openi_format")
	}
	r, err := models.GetRepositoryByOwnerAndName(ownerName, repoName)
	if err != nil {
		if models.IsErrRepoNotExist(err) {
			return nil, errors.New("tech.openi_repo_not_exist")
		}
		return nil, err
	}
	return r, nil
}

//parseOpenIUrl parse openI repo url,return ownerName and repoName
func parseOpenIUrl(u string) (string, string) {
	newUrl := strings.TrimSuffix(u, ".git")
	url, err := url.Parse(newUrl)
	if err != nil {
		return "", ""
	}

	appUrl, err := url.Parse(setting.AppURL)
	if err != nil {
		return "", ""
	}
	if appUrl.Host != url.Host {
		return "", ""
	}

	array := strings.Split(url.Path, "/")
	if len(array) < 3 {
		return "", ""
	}

	ownerName := array[1]
	repoName := array[2]
	return ownerName, repoName
}
