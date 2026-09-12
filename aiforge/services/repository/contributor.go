package repository

import (
	"encoding/json"
	"math/rand"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"github.com/patrickmn/go-cache"
)

var repoContributorCache = cache.New(5*time.Minute, 1*time.Minute)

type ContributorCacheVal struct {
	Contributors []*models.ContributorInfo
	Total        int
}

func GetRepoTopNContributors(repo *models.Repository, N int) ([]*models.ContributorInfo, int) {
	val, _ := redis_client.Get(redis_key.RepoTopNContributors(repo.ID, N))
	if val != "" {
		log.Info("Get RepoTopNContributors from redis,repo.ID = %d value = %v", repo.ID, val)
		temp := &ContributorCacheVal{}
		json.Unmarshal([]byte(val), temp)
		return temp.Contributors, temp.Total
	}

	contributorInfos, total := getRepoTopNContributorsFromDisk(repo, N)
	log.Info("Get RepoTopNContributors from disk,repo.ID = %d ", repo.ID)
	jsonVal, err := json.Marshal(&ContributorCacheVal{Contributors: contributorInfos, Total: total})
	if err == nil {
		redis_client.Setex(redis_key.RepoTopNContributors(repo.ID, N), string(jsonVal), 30*24*time.Hour+time.Duration(rand.Intn(10))*time.Hour)
	}
	return contributorInfos, total
}

func getRepoTopNContributorsFromDisk(repo *models.Repository, N int) ([]*models.ContributorInfo, int) {
	contributorInfos := make([]*models.ContributorInfo, 0)

	branchName := GetDefaultBranchName(repo)
	if branchName == "" {
		return contributorInfos, 0
	}

	contributors, err := git.GetContributors(repo.RepoPath(), branchName)
	if err == nil && contributors != nil {
		contributorInfoHash := make(map[string]*models.ContributorInfo)
		for _, c := range contributors {
			if len(contributorInfos) >= N {
				break
			}
			if c.Email == "" {
				continue
			}
			// get user info from committer email
			user, err := models.GetUserByActivateEmail(c.Email)
			if err == nil {
				// committer is system user, get info through user's primary email
				if existedContributorInfo, ok := contributorInfoHash[user.Email]; ok {
					// existed: same primary email, different committer name
					existedContributorInfo.CommitCnt += c.CommitCnt
				} else {
					// new committer info
					var newContributor = &models.ContributorInfo{
						user.RelAvatarLink(), user.Name, "", c.CommitCnt,
					}
					contributorInfos = append(contributorInfos, newContributor)
					contributorInfoHash[user.Email] = newContributor
				}
			} else {
				// committer is not system user
				if existedContributorInfo, ok := contributorInfoHash[c.Email]; ok {
					// existed: same primary email, different committer name
					existedContributorInfo.CommitCnt += c.CommitCnt
				} else {
					var newContributor = &models.ContributorInfo{
						"", "", c.Email, c.CommitCnt,
					}
					contributorInfos = append(contributorInfos, newContributor)
					contributorInfoHash[c.Email] = newContributor
				}
			}
		}
	}
	return contributorInfos, len(contributors)
}
