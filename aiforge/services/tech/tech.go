package tech

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/services/repository"
)

func FindTech(opt models.FindTechOpt) ([]*models.TechConvergeBrief, error) {
	techList, err := models.FindTech(opt)
	if err != nil {
		return nil, err
	}
	r := make([]*models.TechConvergeBrief, len(techList))
	for i := 0; i < len(techList); i++ {
		r[i] = techList[i].Brief()
	}
	return r, nil
}
func GetMyRepoInfo(opts *models.SearchUserRepoOpt) ([]*models.TechRepoInfoUser, int64, error) {
	infos, total, err := models.GetTechRepoInfoForUser(opts)
	if err != nil {
		return nil, total, err
	}
	result := make([]*models.TechRepoInfoUser, 0, total)
	for _, info := range infos {

		techRepoInfoUser := &models.TechRepoInfoUser{
			ID:                      info.ID,
			Url:                     info.Url,
			Status:                  info.Status,
			ContributionInstitution: info.Institution,
			CreatedUnix:             info.CreatedUnix,
			UpdatedUnix:             info.UpdatedUnix,
		}
		if info.User != nil {
			techRepoInfoUser.UserName = info.User.Name
		}
		if info.BaseInfo != nil {
			baseInfo := info.BaseInfo
			techRepoInfoUser.ProjectNumber = baseInfo.ProjectNumber
			techRepoInfoUser.ProjectName = baseInfo.ProjectName
			techRepoInfoUser.Institution = baseInfo.Institution

			techRepoInfoUser.AllInstitution = baseInfo.AllInstitution
		}
		if info.Repo != nil {
			techRepoInfoUser.RepoName = info.Repo.Name
			techRepoInfoUser.RepoOwnerName = info.Repo.OwnerName
		}

		result = append(result, techRepoInfoUser)
	}
	return result, total, nil
}
func GetAdminRepoInfo(opts *models.SearchUserRepoOpt) ([]*models.TechRepoInfoAdmin, int64, error) {
	infos, total, err := models.GetTechRepoInfoForUser(opts)
	if err != nil {
		return nil, total, err
	}
	result := make([]*models.TechRepoInfoAdmin, 0, total)
	for _, info := range infos {

		techRepoInfoAdmin := &models.TechRepoInfoAdmin{
			ID:                      info.ID,
			Url:                     info.Url,
			Status:                  info.Status,
			ContributionInstitution: info.Institution,
			CreatedUnix:             info.CreatedUnix,
			UpdatedUnix:             info.UpdatedUnix,
		}
		if info.User != nil {
			techRepoInfoAdmin.UserName = info.User.Name
		}
		if info.BaseInfo != nil {
			baseInfo := info.BaseInfo
			techRepoInfoAdmin.ProjectNumber = baseInfo.ProjectNumber
			techRepoInfoAdmin.ProjectName = baseInfo.ProjectName
			techRepoInfoAdmin.Institution = baseInfo.Institution
			techRepoInfoAdmin.ApplyYear = baseInfo.ApplyYear
			techRepoInfoAdmin.Province = baseInfo.Province
			techRepoInfoAdmin.Category = baseInfo.Category
			techRepoInfoAdmin.Owner = baseInfo.Owner
			techRepoInfoAdmin.Recommend = baseInfo.Recommend
			techRepoInfoAdmin.Phone = baseInfo.Phone
			techRepoInfoAdmin.Email = baseInfo.Email

			techRepoInfoAdmin.Contact = baseInfo.Contact
			techRepoInfoAdmin.ContactPhone = baseInfo.ContactPhone
			techRepoInfoAdmin.ContactEmail = baseInfo.ContactEmail
			techRepoInfoAdmin.ExecuteMonth = baseInfo.ExecuteMonth
			techRepoInfoAdmin.ExecutePeriod = baseInfo.ExecutePeriod
			techRepoInfoAdmin.Type = baseInfo.Type
			techRepoInfoAdmin.StartUp = baseInfo.StartUp
			techRepoInfoAdmin.NumberTopic = baseInfo.NumberTopic
			techRepoInfoAdmin.Topic1 = baseInfo.Topic1
			techRepoInfoAdmin.Topic2 = baseInfo.Topic2
			techRepoInfoAdmin.Topic3 = baseInfo.Topic3
			techRepoInfoAdmin.Topic4 = baseInfo.Topic4
			techRepoInfoAdmin.Topic5 = baseInfo.Topic5
			techRepoInfoAdmin.AllInstitution = baseInfo.AllInstitution
		}
		if info.Repo != nil {
			techRepoInfoAdmin.RepoName = info.Repo.Name
			techRepoInfoAdmin.RepoOwnerName = info.Repo.OwnerName
		}

		result = append(result, techRepoInfoAdmin)
	}
	return result, total, nil

}

func SearchRepoInfoWithInstitution(opt *models.SearchRepoOpt) ([]*models.RepoWithInstitution, int64, error) {
	infos, err := models.GetAvailableRepoConvergeInfo(opt)
	if err != nil {
		return nil, 0, err
	}
	if len(infos) == 0 {
		return []*models.RepoWithInstitution{}, 0, nil
	}
	repoIds, institutionMap := getRepoIdAndInstitutionMap(infos)
	topics := make([]string, 0)
	if opt.Topic != "" {
		topics = append(topics, opt.Topic)
	}

	result, err := repository.FindRepos(repository.FindReposOptions{
		ListOptions: models.ListOptions{Page: opt.Page, PageSize: opt.PageSize},
		Sort:        opt.OrderBy,
		Keyword:     opt.Q,
		Topics:      topics,
		RepoIds:     repoIds,
	})

	if err != nil {
		return nil, 0, err
	}
	repoResult := make([]*models.RepoWithInstitution, 0, len(result.Repos))

	for _, repo := range result.Repos {
		repoResult = append(repoResult, &models.RepoWithInstitution{
			ID:            repo.ID,
			OwnerID:       repo.OwnerID,
			OwnerName:     repo.OwnerName,
			Name:          repo.Name,
			Alias:         repo.Alias,
			Topics:        repo.Topics,
			Description:   repo.Description,
			Institution:   institutionMap[repo.ID],
			RelAvatarLink: repo.RelAvatarLink,
			UpdatedUnix:   repo.UpdatedUnix,
		})
	}

	return repoResult, result.Total, nil

}

func getRepoIdAndInstitutionMap(infos []*models.RepoConvergeInfo) ([]int64, map[int64]string) {
	repoIds := make([]int64, 0, len(infos))
	institutionMap := make(map[int64]string, 0)
	// We only need the keys
	for _, info := range infos {
		repoIds = append(repoIds, info.RepoID)
		institutionMap[info.RepoID] = info.Institution
	}
	return repoIds, institutionMap
}

func IsValidRepoConvergeExists(repoId, baseInfoId int64) (bool, error) {
	list, err := models.GetRepoConverge(models.GetRepoConvergeOpts{
		Status:     []int{models.TechHide, models.TechShow, models.TechMigrating},
		RepoId:     repoId,
		BaseInfoId: baseInfoId,
	})
	if err != nil {
		return false, err
	}
	return len(list) > 0, nil
}

func UpdateTechStatus() {
	err := UpdateTechRepoDeleteStatus()
	if err != nil {
		log.Error("update delete status fail", err)
	}

	err = UpdateTechMigrateStatus()
	if err != nil {
		log.Error("update migrate status fail", err)
	}
}

func UpdateTechRepoDeleteStatus() error {
	needDeleteCheckRepos, err := models.GetRepoConverge(models.GetRepoConvergeOpts{
		Status: []int{models.TechShow, models.TechHide},
	})
	if err != nil {
		return err
	}

	for _, r := range needDeleteCheckRepos {
		_, err := models.GetRepositoryByID(r.RepoID)
		if err != nil && models.IsErrRepoNotExist(err) {
			models.UpdateRepoConvergeStatus(r.ID, models.TechNotExist)
			continue
		}

	}
	return nil

}

func UpdateTechMigrateStatus() error {
	migratingRepos, err := models.GetRepoConverge(models.GetRepoConvergeOpts{
		Status: []int{models.TechMigrating},
	})
	if err != nil {
		return err
	}
	for _, r := range migratingRepos {
		repo, err := models.GetRepositoryByID(r.RepoID)
		if err != nil {
			if models.IsErrRepoNotExist(err) {
				models.UpdateRepoConvergeStatus(r.ID, models.TechMigrateFailed)
				continue
			}
			continue
		}
		if repo.Status == models.RepositoryReady {
			models.UpdateRepoConvergeStatus(r.ID, models.DefaultTechApprovedStatus)
			continue
		}
	}
	return nil
}
