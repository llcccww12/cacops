package admin

import (
	"fmt"
	"strings"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
)

func AdminModelManage(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("admin.models")
	ctx.Data["PageIsAdmin"] = true
	ctx.Data["PageIsAdminModels"] = true

	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := ctx.QueryInt("pageSize")
	if pageSize <= 0 {
		pageSize = setting.UI.IssuePagingNum
	}

	var (
		count   int64
		err     error
		orderBy string
	)

	ctx.Data["SortType"] = ctx.Query("sort")
	switch ctx.Query("sort") {
	case "newest":
		orderBy = "created_unix DESC"
	case "oldest":
		orderBy = "created_unix ASC"
	case "recentupdate":
		orderBy = "updated_unix DESC"
	case "leastupdate":
		orderBy = "updated_unix ASC"
	case "reversealphabetically":
		orderBy = "name DESC"
	case "alphabetically":
		orderBy = "name ASC"
	case "reversesize":
		orderBy = "size DESC"
	case "size":
		orderBy = "size ASC"
	case "downloadtimes":
		orderBy = "download_count DESC"
	case "mostusecount":
		orderBy = "reference_count DESC"
	case "fewestusecount":
		orderBy = "reference_count ASC"
	case "derivativeCount":
		orderBy = "derivative_count DESC"
	default:
		ctx.Data["SortType"] = "recentupdate"
		orderBy = "created_unix DESC"
	}

	keyword := strings.Trim(ctx.Query("q"), " ")

	modelResult, count, err := models.QueryModel(&models.AiModelQueryOptions{
		ListOptions: models.ListOptions{
			Page:     page,
			PageSize: setting.UI.ExplorePagingNum,
		},
		Type:                  -1,
		New:                   -1,
		Status:                -1,
		IsQueryPrivate:        true,
		IsRecommend:           ctx.QueryBool("recommend"),
		UserID:                -1,
		IsCollected:           false,
		CollectedUserId:       -1,
		LabelFilter:           "",
		FrameFilter:           -1,
		ComputeResourceFilter: "",
		Namelike:              keyword,
		SortType:              orderBy,
		RepoID:                -1,
	})
	if err != nil {
		ctx.ServerError("Cloudbrain", err)
		return
	}
	userIds := make([]int64, len(modelResult))
	modelIds := make([]string, len(modelResult))
	repoIds := make([]int64, len(modelResult))
	for i, model := range modelResult {
		userIds[i] = model.UserId
		modelIds[i] = model.ID
		repoIds[i] = model.RepoId
	}
	repoInfo, err := queryRepoInfoByIds(repoIds)
	userNameMap := queryUserName(userIds)

	for _, model := range modelResult {
		//removeIpInfo(model)
		model.TrainTaskInfo = ""
		value := userNameMap[model.UserId]
		if value != nil {
			model.UserName = value.Name
			model.UserRelAvatarLink = value.RelAvatarLink()
		}
		if repoInfo != nil {
			repo := repoInfo[model.RepoId]
			if repo != nil {
				model.RepoName = repo.Name
				model.RepoOwnerName = repo.OwnerName
				model.RepoDisplayName = repo.DisplayName()
			}
		}
	}

	ctx.Data["Keyword"] = keyword
	ctx.Data["Total"] = count
	ctx.Data["models"] = modelResult
	ctx.Data["Recommend"] = ctx.QueryBool("recommend")
	pager := context.NewPagination(int(count), setting.UI.ExplorePagingNum, page, 5)
	pager.SetDefaultParams(ctx)
	ctx.Data["Page"] = pager

	ctx.HTML(200, tplAdminModelManage)
}

func ModifyModelRecommend(ctx *context.Context) {
	id := ctx.Query("id")
	isRecommend := ctx.QueryInt("recommend")
	re := map[string]string{
		"code": "-1",
	}
	task, err := models.QueryModelById(id)
	if err != nil || task == nil {
		re["msg"] = err.Error()
		log.Error("no such model!", err.Error())
		ctx.JSON(200, re)
		return
	}
	if ctx.User == nil || !ctx.User.IsAdmin {
		re["msg"] = "No right to operation."
		ctx.JSON(200, re)
		return
	}

	err = models.ModifyModelRecommend(id, isRecommend)
	if err == nil {
		re["code"] = "0"
		ctx.JSON(200, re)
		log.Info("modify success.")
	} else {
		re["msg"] = err.Error()
		ctx.JSON(200, re)
		log.Info("Failed to modify.id=" + id + "  isprivate=" + fmt.Sprint(isRecommend) + " error:" + err.Error())
	}
}

func queryUserName(intSlice []int64) map[int64]*models.User {
	keys := make(map[int64]string)
	uniqueElements := []int64{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = ""
			uniqueElements = append(uniqueElements, entry)
		}
	}
	result := make(map[int64]*models.User)
	userLists, err := models.GetUsersByIDs(uniqueElements)
	if err == nil {
		for _, user := range userLists {
			result[user.ID] = user
		}
	}
	return result
}

func queryRepoInfoByIds(intSlice []int64) (map[int64]*models.Repository, error) {
	keys := make(map[int64]string)
	uniqueElements := []int64{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = ""
			uniqueElements = append(uniqueElements, entry)
		}
	}
	re, err := models.GetRepositoriesMapByIDs(uniqueElements)
	return re, err
}
