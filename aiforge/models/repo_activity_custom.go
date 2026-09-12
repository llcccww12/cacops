package models

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/log"

	"code.gitea.io/gitea/modules/git"
)

type ContributorWithUserId struct {
	git.Contributor
	UserId        int64
	IsAdmin       bool
	RelAvatarLink string
}

func GetRepoKPIStats(repo *Repository) (*git.RepoKPIStats, error) {
	wikiPath := ""
	if repo.HasWiki() {
		wikiPath = repo.WikiPath()
	}
	repoCreated := time.Unix(int64(repo.CreatedUnix), 0)
	return getRepoKPIStats(repo.RepoPath(), repoCreated, wikiPath)
}

func getRepoKPIStats(repoPath string, repoCreated time.Time, wikiPath string) (*git.RepoKPIStats, error) {
	stats := &git.RepoKPIStats{}

	contributors, err := git.GetContributorsDetail(repoPath, repoCreated)
	if err != nil {
		return nil, err
	}
	timeUntil := time.Now()
	fourMonthAgo := timeUntil.AddDate(0, -4, 0)
	if fourMonthAgo.Before(repoCreated) {
		fourMonthAgo = repoCreated
	}
	recentlyContributors, err := git.GetContributorsDetail(repoPath, fourMonthAgo)
	newContributersDict := make(map[string]struct{})
	if err != nil {
		return nil, err
	}

	if contributors != nil {
		contributorDistinctDict := make(map[string]int, 0)
		keyContributorsDict := make(map[string]struct{}, 0)
		var commitsCount int64
		for _, contributor := range contributors {
			if strings.Compare(contributor.Email, "") == 0 {
				continue
			}

			user, err := GetUserByActivateEmail(contributor.Email)
			if err == nil {
				value, ok := contributorDistinctDict[user.Email]
				if !ok {
					contributorDistinctDict[user.Email] = contributor.CommitCnt
				} else {
					contributorDistinctDict[user.Email] = value + contributor.CommitCnt
				}
				setKeyContributerDict(contributorDistinctDict, user.Email, keyContributorsDict)

			} else {
				value, ok := contributorDistinctDict[contributor.Email]
				if !ok {
					contributorDistinctDict[contributor.Email] = contributor.CommitCnt
				} else {
					contributorDistinctDict[contributor.Email] = value + contributor.CommitCnt
				}
				setKeyContributerDict(contributorDistinctDict, contributor.Email, keyContributorsDict)
			}

			commitsCount += int64(contributor.CommitCnt)

		}

		if recentlyContributors != nil {
			resentlyContributorDistinctDict := make(map[string]int, 0)
			for _, recentlyContributor := range recentlyContributors {

				user, err := GetUserByActivateEmail(recentlyContributor.Email)

				if err == nil {
					value, ok := resentlyContributorDistinctDict[user.Email]
					if !ok {
						resentlyContributorDistinctDict[user.Email] = recentlyContributor.CommitCnt
					} else {
						resentlyContributorDistinctDict[user.Email] = value + recentlyContributor.CommitCnt
					}

				} else {
					value, ok := resentlyContributorDistinctDict[recentlyContributor.Email]
					if !ok {
						resentlyContributorDistinctDict[recentlyContributor.Email] = recentlyContributor.CommitCnt
					} else {
						resentlyContributorDistinctDict[recentlyContributor.Email] = value + recentlyContributor.CommitCnt
					}

				}

			}

			for k, v := range resentlyContributorDistinctDict {
				count, ok := contributorDistinctDict[k]
				if ok && count == v {
					stats.ContributorsAdded++
				}

			}

		}

		stats.Contributors = int64(len(contributorDistinctDict))
		stats.KeyContributors = int64(len(keyContributorsDict))
		stats.Commits = int64(commitsCount)
	}

	err = git.SetDevelopAge(repoPath, stats, repoCreated)
	if err != nil {
		return nil, fmt.Errorf("FillFromGit: %v", err)
	}
	err = git.SetRepoKPIStats(repoPath, fourMonthAgo, stats, newContributersDict)

	if err != nil {
		return nil, fmt.Errorf("FillFromGit: %v", err)
	}

	git.SetWikiPages(wikiPath, stats)
	return stats, nil

}

func GetTop10Contributor(repoPath string) ([]*ContributorWithUserId, error) {
	contributors, err := git.GetContributors(repoPath)
	if err != nil {
		return make([]*ContributorWithUserId, 0), err
	}
	contributorDistinctDict := make(map[string]*ContributorWithUserId, 0)
	if contributors != nil {
		for _, contributor := range contributors {
			if strings.Compare(contributor.Email, "") == 0 {
				continue
			}

			user, err := GetUserByActivateEmail(contributor.Email)
			if err == nil {

				value, ok := contributorDistinctDict[user.Email]
				if !ok {
					contributorDistinctDict[user.Email] = &ContributorWithUserId{
						git.Contributor{
							contributor.CommitCnt,
							user.Name,
							user.Email,
						},
						user.ID,
						user.IsAdmin,
						user.RelAvatarLink(),
					}
				} else {

					value.CommitCnt += contributor.CommitCnt
				}

			} else {
				value, ok := contributorDistinctDict[contributor.Email]
				if !ok {
					contributorDistinctDict[contributor.Email] = &ContributorWithUserId{
						contributor,
						-1,
						false,
						"",
					}
				} else {
					value.CommitCnt += contributor.CommitCnt
				}

			}

		}
		v := make([]*ContributorWithUserId, 0, len(contributorDistinctDict))
		for _, value := range contributorDistinctDict {
			v = append(v, value)
		}

		sort.Slice(v, func(i, j int) bool {
			return v[i].CommitCnt > v[j].CommitCnt
		})

		if len(v) <= 10 {
			return v, nil
		} else {
			return v[0:10], nil
		}
	}
	return make([]*ContributorWithUserId, 0), nil
}

func setKeyContributerDict(contributorDistinctDict map[string]int, email string, keyContributorsDict map[string]struct{}) {
	if contributorDistinctDict[email] >= 3 {
		_, ok := keyContributorsDict[email]
		if !ok {
			keyContributorsDict[email] = struct{}{}

		}

	}
}

func GetAllUserPublicRepoKPIStats(startTime time.Time, endTime time.Time) (map[string]*git.UserKPIStats, error) {
	authors := make(map[string]*git.UserKPIStats)
	repositorys, err := GetAllRepositoriesByFilterCols("owner_name", "name", "is_private")
	if err != nil {
		return nil, err
	}

	for _, repository := range repositorys {
		if repository.IsPrivate {
			continue
		}
		authorsOneRepo, err1 := git.GetUserKPIStats(repository.RepoPath(), startTime, endTime)
		if err1 != nil {
			log.Warn("get user kpi status err:"+repository.RepoPath(), err1.Error())
			continue
		}

		for key, value := range authorsOneRepo {
			if _, ok := authors[key]; !ok {
				authors[key] = &git.UserKPIStats{

					Name:        value.Name,
					Email:       value.Email,
					Commits:     0,
					CommitLines: 0,
				}
			}
			authors[key].Commits += value.Commits
			authors[key].CommitLines += value.CommitLines
		}

	}
	return authors, nil
}

func GetAllUserKPIStats(startTime time.Time, endTime time.Time) (map[string]*git.UserKPIStats, error) {
	authors := make(map[string]*git.UserKPIStats)
	repositorys, err := GetAllRepositoriesByFilterCols("owner_name", "name")
	if err != nil {
		return nil, err
	}

	for _, repository := range repositorys {
		authorsOneRepo, err1 := git.GetUserKPIStats(repository.RepoPath(), startTime, endTime)
		if err1 != nil {
			log.Warn("get user kpi status err:"+repository.RepoPath(), err1.Error())
			continue
		}
		// if repository.Name == "yolov5" {
		// 	log.Info("repoName=" + repository.Name + "  owner=" + repository.RepoPath())
		// 	authorsOneRepoJson, _ := json.Marshal(authorsOneRepo)
		// 	log.Info("authorsOneRepoJson=" + string(authorsOneRepoJson))
		// }
		for key, value := range authorsOneRepo {
			if _, ok := authors[key]; !ok {
				authors[key] = &git.UserKPIStats{

					Name:        value.Name,
					Email:       value.Email,
					Commits:     0,
					CommitLines: 0,
				}
			}
			authors[key].Commits += value.Commits
			authors[key].CommitLines += value.CommitLines

		}

	}
	return authors, nil
}
