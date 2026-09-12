// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2020 The Gitea Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package org

import (
	"errors"
	"fmt"
	"strings"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/auth"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
)

const DefaultOrgTagLimit = 9

// SubmitTags submit repos of org tag
func SubmitTags(ctx *context.Context, form auth.SubmitReposOfTagForm) {
	if !ctx.Org.IsOwner {
		ctx.ServerError("UpdateTagReposByID", errors.New("no access to submit tags"))
		return
	}
	// tag := getTagFromContext(ctx)
	// if ctx.Written() {
	// 	return
	// }

	if len(form.RepoList) > 9 {
		ctx.ServerError("UpdateTagReposByID", errors.New("tags size over limit"))
		return
	}
	for _, v := range form.RepoList {
		log.Info("v=" + fmt.Sprint(v))
	}

	err := models.UpdateTagReposByID(ctx.Org.Organization.ID, form.RepoList)
	if err != nil {
		ctx.ServerError("UpdateTagReposByID", err)
		return
	}

	ctx.JSON(200, map[string]interface{}{
		"code": "00",
		"msg":  "success",
	})
}

func SubmitTagsToDataset(ctx *context.Context, form auth.SubmitDatasetOfTagForm) {
	if !ctx.Org.IsOwner {
		ctx.ServerError("UpdateTagReposByID", errors.New("no access to submit tags"))
		return
	}

	datasetid := form.Datasetids
	log.Info("datasetids=" + datasetid)
	datasetids := strings.Split(datasetid, ",")
	if len(datasetids) > 9 {
		ctx.JSON(200, map[string]interface{}{
			"code": "-1",
			"msg":  "Submit tags cannot over 9.",
		})
		return
	}

	err := models.UpdateTagDatasetByID(ctx.Org.Organization.ID, datasetids)
	if err != nil {
		ctx.ServerError("UpdateTagDatasetByID", err)
		return
	}

	ctx.JSON(200, map[string]interface{}{
		"code": "0",
		"msg":  "success",
	})
}

func SubmitTagsToModel(ctx *context.Context, form auth.SubmitModelOfTagForm) {
	if !ctx.Org.IsOwner {
		ctx.ServerError("UpdateTagReposByID", errors.New("no access to submit tags"))
		return
	}

	modelId := form.Modelids
	log.Info("modelids=" + modelId)
	modelIds := strings.Split(modelId, ",")
	if len(modelIds) > 9 {
		ctx.JSON(200, map[string]interface{}{
			"code": "-1",
			"msg":  "Submit tags cannot over 9.",
		})
		return
	}

	err := models.UpdateTagModelByID(ctx.Org.Organization.ID, modelIds)
	if err != nil {
		ctx.ServerError("UpdateTagModelByID", err)
		return
	}

	ctx.JSON(200, map[string]interface{}{
		"code": "0",
		"msg":  "success",
	})
}

// GetTagModel get repos under org tag
func GetTagModel(ctx *context.Context) {
	if !ctx.Org.IsOwner {
		ctx.ServerError("GetTagModel", errors.New("no access to get tags"))
		return
	}

	r, err := models.GetTagModel(ctx.Org.Organization.ID)
	if err != nil {
		ctx.ServerError("GetTagModel", err)
		return
	}

	ctx.JSON(200, map[string]interface{}{
		"code": "0",
		"msg":  "success",
		"data": r,
	})
}

// GetTagModel get repos under org tag
func GetTagDataset(ctx *context.Context) {
	if !ctx.Org.IsOwner {
		ctx.ServerError("GetTagDataset", errors.New("no access to get tags"))
		return
	}

	r, err := models.GetTagDataset(ctx.Org.Organization.ID)
	if err != nil {
		ctx.ServerError("GetTagDataset", err)
		return
	}

	ctx.JSON(200, map[string]interface{}{
		"code": "0",
		"msg":  "success",
		"data": r,
	})

}

func GetOrgDatasetList(ctx *context.Context) {

	var orderBy models.SearchOrderBy

	switch ctx.Query("sort") {
	case "newest":
		orderBy = models.SearchOrderByNewest
	case "oldest":
		orderBy = models.SearchOrderByOldest
	case "recentupdate":
		orderBy = models.SearchOrderByRecentUpdated
	case "leastupdate":
		orderBy = models.SearchOrderByLeastUpdated
	case "reversealphabetically":
		orderBy = models.SearchOrderByTitleAlphabeticallyReverse
	case "alphabetically":
		orderBy = models.SearchOrderByTitleAlphabetically
	default:
		orderBy = models.SearchOrderByRecentUpdated
	}

	keyword := strings.Trim(ctx.Query("q"), " ")
	pageSize := ctx.QueryInt("pageSize")
	if pageSize <= 0 {
		pageSize = setting.UI.User.RepoPagingNum
	}
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	categoryFilter := strings.Trim(ctx.Query("categoryFilter"), " ")
	taskFilter := strings.Trim(ctx.Query("taskFilter"), " ")
	licenseFilter := strings.Trim(ctx.Query("licenseFilter"), " ")
	var uid int64
	uid = -1
	var isAdmin bool
	isAdmin = false
	if ctx.User != nil {
		uid = ctx.User.ID
		if ctx.User.IsAdmin || ctx.User.IsUserOrgOwner(ctx.Org.Organization.ID) {
			isAdmin = true
		}
	}
	log.Info("categoryFilter=" + categoryFilter + " taskFilter=" + taskFilter + " licenseFilter=" + licenseFilter)
	publicdataset, count := models.GetOwnedPublicDatasetsByUserIDPage(ctx.Org.Organization.ID, keyword, categoryFilter, taskFilter, licenseFilter, orderBy, page, pageSize, uid, isAdmin)

	datasetLabelSet := models.GetAllOwnedPublicDatasetsByOrgID(ctx.Org.Organization.ID)
	orgTopics := make(map[string][]string, 0)
	categoryList := make([]string, 0)
	categoryMap := make(map[string]bool, 0)

	taskList := make([]string, 0)
	taskMap := make(map[string]bool, 0)

	licenseList := make([]string, 0)
	licenseMap := make(map[string]bool, 0)

	for _, v := range datasetLabelSet {
		categoryMap[v.Category] = true
		taskMap[v.Task] = true
		licenseMap[v.License] = true
	}

	for k, _ := range categoryMap {
		categoryList = append(categoryList, k)
	}
	for k, _ := range taskMap {
		taskList = append(taskList, k)
	}
	for k, _ := range licenseMap {
		licenseList = append(licenseList, k)
	}
	orgTopics["category"] = categoryList
	orgTopics["task"] = taskList
	orgTopics["license"] = licenseList

	ctx.JSON(200, map[string]interface{}{
		"code":      "0",
		"msg":       "success",
		"total":     count,
		"data":      publicdataset,
		"orgTopics": orgTopics,
	})

}

func GetOrgDatasetCardList(ctx *context.Context) {
	datasetCard, err := models.GetTagDatasetForCard(ctx.Org.Organization.ID)
	if err == nil {
		ctx.JSON(200, map[string]interface{}{
			"code": "0",
			"msg":  "success",
			"data": datasetCard,
		})
	} else {
		ctx.JSON(200, map[string]interface{}{
			"code": "-1",
			"msg":  "Failed",
		})
	}
}

func GetOrgModelList(ctx *context.Context) {

	var orderBy models.SearchOrderBy

	switch ctx.Query("sort") {
	case "newest":
		orderBy = models.SearchOrderByNewest
	case "oldest":
		orderBy = models.SearchOrderByOldest
	case "recentupdate":
		orderBy = models.SearchOrderByRecentUpdated
	case "leastupdate":
		orderBy = models.SearchOrderByLeastUpdated
	case "reversealphabetically":
		orderBy = models.SearchOrderByAlphabeticallyReverse
	case "alphabetically":
		orderBy = models.SearchOrderByAlphabetically
	default:
		orderBy = models.SearchOrderByRecentUpdated
	}

	keyword := strings.Trim(ctx.Query("q"), " ")
	labelFilter := strings.Trim(ctx.Query("filter"), " ")
	pageSize := ctx.QueryInt("pageSize")
	if pageSize <= 0 {
		pageSize = setting.UI.User.RepoPagingNum
	}
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	var uid int64
	uid = -1
	var isAdmin bool
	isAdmin = false
	if ctx.User != nil {
		uid = ctx.User.ID
		if ctx.User.IsAdmin || ctx.User.IsUserOrgOwner(ctx.Org.Organization.ID) {
			isAdmin = true
		}
	}
	modelResult, count, err := models.GetOrgModelList(ctx.Org.Organization.ID, uid, keyword, labelFilter, orderBy, page, pageSize, isAdmin)
	if err == nil {
		setModelOtherInfo(modelResult, ctx)
		orgTopics := make([]string, 0)
		allModelLabel, err1 := models.GetOrgModelLabelList(ctx.Org.Organization.ID, uid)
		if err1 == nil {
			labelMap := make(map[string]bool, 0)
			for _, t := range allModelLabel {
				if t.Label != "" {
					labels := strings.Split(t.Label, " ")
					for _, t := range labels {
						t = strings.TrimSpace(t)
						if t != "" {
							labelMap[t] = true
						}
					}
				}
			}
			for k, _ := range labelMap {
				orgTopics = append(orgTopics, k)
			}
		}
		ctx.JSON(200, map[string]interface{}{
			"code":      "0",
			"msg":       "success",
			"total":     count,
			"data":      modelResult,
			"orgTopics": orgTopics,
		})
	} else {
		ctx.JSON(200, map[string]interface{}{
			"code": "-1",
			"msg":  "Failed",
		})
	}
}

func setModelOtherInfo(modelResult []*models.AiModelManage, ctx *context.Context) {
	userIds := make([]int64, len(modelResult))
	modelIds := make([]string, len(modelResult))
	repoIds := make([]int64, len(modelResult))
	for i, model := range modelResult {
		userIds[i] = model.UserId
		modelIds[i] = model.ID
		repoIds[i] = model.RepoId
	}
	repoInfo, _ := queryRepoInfoByIds(repoIds)
	userNameMap := models.QueryUserName(userIds)
	var modelCollect map[string]*models.AiModelCollect

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
		if ctx.User != nil && modelCollect != nil {
			value := modelCollect[model.ID]
			if value != nil {
				model.IsCollected = true
			}
		} else {
			model.IsCollected = false
		}
	}
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
func GetOrgModelCardList(ctx *context.Context) {
	modelList, err := models.GetTagModelForCard(ctx.Org.Organization.ID)
	if err == nil {
		setModelOtherInfo(modelList, ctx)
		ctx.JSON(200, map[string]interface{}{
			"code": "0",
			"msg":  "success",
			"data": modelList,
		})
	} else {
		ctx.JSON(200, map[string]interface{}{
			"code": "-1",
			"msg":  "Failed",
		})
	}
}

func GetOrgRepoList(ctx *context.Context) {
	org := ctx.Org.Organization
	var orderBy models.SearchOrderBy
	switch ctx.Query("sort") {
	case "newest":
		orderBy = models.SearchOrderByNewest
	case "oldest":
		orderBy = models.SearchOrderByOldest
	case "recentupdate":
		orderBy = models.SearchOrderByRecentUpdated
	case "leastupdate":
		orderBy = models.SearchOrderByLeastUpdated
	case "reversealphabetically":
		orderBy = models.SearchOrderByAlphabeticallyReverse
	case "alphabetically":
		orderBy = models.SearchOrderByAlphabetically
	case "moststars":
		orderBy = models.SearchOrderByStarsReverse
	case "feweststars":
		orderBy = models.SearchOrderByStars
	case "mostforks":
		orderBy = models.SearchOrderByForksReverse
	case "fewestforks":
		orderBy = models.SearchOrderByForks
	default:
		orderBy = models.SearchOrderByRecentUpdated
	}
	orderBy = orderBy + ",id"
	keyword := strings.Trim(ctx.Query("q"), " ")
	topicName := strings.Trim(ctx.Query("filter"), " ")

	pageSize := ctx.QueryInt("pageSize")
	if pageSize <= 0 {
		pageSize = setting.UI.User.RepoPagingNum
	}
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}

	repos, count, err := models.SearchRepository(&models.SearchRepoOptions{
		ListOptions: models.ListOptions{
			PageSize: pageSize,
			Page:     page,
		},
		Keyword:            keyword,
		OwnerID:            org.ID,
		OrderBy:            orderBy,
		TopicName:          topicName,
		Private:            ctx.IsSigned,
		Actor:              ctx.User,
		IncludeDescription: setting.UI.SearchRepoDescription,
	})

	result := make([]*models.Repository4Card, len(repos))
	for i, r := range repos {
		t := r.ToCardFormat()
		// contributors, _ := repository.GetRepoTopNContributors(r, 6)
		// t.Contributors = contributors
		result[i] = t
	}

	orgTopics, err := models.GetOrgTopics(org.ID)
	if err != nil {
		ctx.Error(500, "GetOrgTopics failed")
		return
	}

	if err != nil {
		ctx.JSON(200, map[string]interface{}{
			"code": "-1",
			"msg":  "Failed",
		})
	} else {
		ctx.JSON(200, map[string]interface{}{
			"code":      "0",
			"msg":       "success",
			"total":     count,
			"data":      result,
			"orgTopics": orgTopics,
		})
	}
}

func GetOrgRepoCardList(ctx *context.Context) {

	tags, err := models.GetAllOfficialTagRepos(ctx.Org.Organization.ID, ctx.Org.IsOwner)
	if err == nil {
		ctx.JSON(200, map[string]interface{}{
			"code": "0",
			"msg":  "success",
			"data": tags,
		})
	} else {
		ctx.JSON(200, map[string]interface{}{
			"code": "-1",
			"msg":  "failed.",
		})
	}

}

// GetTagRepos get repos under org tag
func GetTagRepos(ctx *context.Context) {
	if !ctx.Org.IsOwner {
		ctx.ServerError("GetTagRepos", errors.New("no access to get tags"))
		return
	}
	// tag := getTagFromContext(ctx)
	// if ctx.Written() {
	// 	return
	// }

	r, err := models.GetTagRepos(ctx.Org.Organization.ID)
	if err != nil {
		ctx.ServerError("GetTagRepos", err)
		return
	}

	ctx.JSON(200, map[string]interface{}{
		"code": "00",
		"msg":  "success",
		"data": r,
	})
}

// // getTagFromContext finds out tag info From context.
// func getTagFromContext(ctx *context.Context) *models.OfficialTag {
// 	var tag *models.OfficialTag
// 	var err error

// 	tagIdStr := ctx.Query("tagId")
// 	if len(tagIdStr) == 0 {
// 		ctx.ServerError("GetTagInfo", errors.New("tag is not exist"))
// 		return nil
// 	}
// 	tagId, _ := strconv.ParseInt(tagIdStr, 10, 32)
// 	tag, err = models.GetTagByID(tagId)
// 	if err != nil {
// 		if models.IsErrTagNotExist(err) {
// 			ctx.NotFound("GetTagInfo", err)
// 		} else {
// 			ctx.ServerError("GetTagInfo", err)
// 		}
// 		return nil
// 	}

// 	return tag
// }
