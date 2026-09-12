// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/routers/repo"

	"code.gitea.io/gitea/routers/response"

	"code.gitea.io/gitea/services/repository"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
	code_indexer "code.gitea.io/gitea/modules/indexer/code"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/modules/util"
	"code.gitea.io/gitea/routers/user"
)

const (
	// tplHome home page template
	tplHome base.TplName = "home"
	// tplExploreRepos explore repositories page template
	tplExploreRepos base.TplName = "explore/repos"
	// tplExploreDataset explore datasets page template
	tplExploreDataset base.TplName = "explore/datasets"
	// tplExploreDatasetDetail explore datasets page template
	tplExploreDatasetDetail base.TplName = "explore/datasets_detail"
	// tplExploreModel explore model page template
	tplExploreModel base.TplName = "explore/models"
	// tplExploreModelDetail explore model page template
	tplExploreModelDetail base.TplName = "explore/models_detail"
	// tplExploreUsers explore users page template
	tplExploreUsers base.TplName = "explore/users"
	// tplExploreOrganizations explore organizations page template
	tplExploreOrganizations base.TplName = "explore/organizations"
	// tplExploreCode explore code page template
	tplExploreCode      base.TplName = "explore/code"
	tplExploreImages    base.TplName = "explore/images"
	tplExploreOrg       base.TplName = "explore/org_my"
	tplIde              base.TplName = "explore/ide"
	tplHomeTerm         base.TplName = "terms"
	tplHomeAnnual       base.TplName = "annual_privacy"
	tplHomeWenxin       base.TplName = "wenxin_privacy"
	tplHomePrivacy      base.TplName = "privacy"
	tplHomeAuthTerm     base.TplName = "auth_terms"
	tplResoruceDesc     base.TplName = "resource_desc"
	tplRepoSquare       base.TplName = "explore/repos/square"
	tplRepoSearch       base.TplName = "explore/repos/search"
	tplRoshmci          base.TplName = "explore/ros-hmci"
	tplExploreCenterMap base.TplName = "explore/center_map"
	tplExploreC2Net     base.TplName = "explore/c2net"

	tplComputingPowerDemand   base.TplName = "computingpower/demand"
	tplComputingPowerDomestic base.TplName = "computingpower/domestic"
	MODEL_MAX_SIZE                         = 1024 * 1024 * 1024
)

// Home render home page
func Home(ctx *context.Context) {
	ctx.Data["PageIsHome"] = true
	ctx.Data["IsRepoIndexerEnabled"] = setting.Indexer.RepoIndexerEnabled
	setRecommendURL(ctx)
	ctx.HTML(200, tplHome)
}

func setRecommendURLOnly(ctx *context.Context) {
	addr := setting.RecommentRepoAddr[10:]
	start := strings.Index(addr, "/")
	end := strings.Index(addr, "raw")
	if start != -1 && end != -1 {
		ctx.Data["RecommendURL"] = addr[start:end]
	} else {
		ctx.Data["RecommendURL"] = setting.RecommentRepoAddr
	}
}

func setRecommendURL(ctx *context.Context) {
	setRecommendURLOnly(ctx)
	ctx.Data["page_title"] = ctx.Tr("home.page_title")
	ctx.Data["page_small_title"] = ctx.Tr("home.page_small_title")
	ctx.Data["page_description"] = ctx.Tr("home.page_description")
	ctx.Data["page_use"] = ctx.Tr("home.page_use")
	ctx.Data["page_only_dynamic"] = ctx.Tr("home.page_only_dynamic")
	ctx.Data["page_recommend_org"] = ctx.Tr("home.page_recommend_org")
	ctx.Data["page_recommend_org_desc"] = ctx.Tr("home.page_recommend_org_desc")
	ctx.Data["page_recommend_org_commit"] = ctx.Tr("home.page_recommend_org_commit")
	ctx.Data["page_recommend_org_more"] = ctx.Tr("home.page_recommend_org_more")
	ctx.Data["page_recommend_repo"] = ctx.Tr("home.page_recommend_repo")
	ctx.Data["page_recommend_repo_desc"] = ctx.Tr("home.page_recommend_repo_desc")
	ctx.Data["page_recommend_repo_commit"] = ctx.Tr("home.page_recommend_repo_commit")
	ctx.Data["page_recommend_repo_go"] = ctx.Tr("home.page_recommend_repo_go")
	ctx.Data["page_recommend_repo_more"] = ctx.Tr("home.page_recommend_repo_more")
	ctx.Data["page_dev_env"] = ctx.Tr("home.page_dev_env")
	ctx.Data["page_dev_env_desc"] = ctx.Tr("home.page_dev_env_desc")
	ctx.Data["page_dev_env_desc_title"] = ctx.Tr("home.page_dev_env_desc_title")
	ctx.Data["page_dev_env_desc_desc"] = ctx.Tr("home.page_dev_env_desc_desc")
	ctx.Data["page_dev_env_desc1_title"] = ctx.Tr("home.page_dev_env_desc1_title")
	ctx.Data["page_dev_env_desc1_desc"] = ctx.Tr("home.page_dev_env_desc1_desc")
	ctx.Data["page_dev_env_desc2_title"] = ctx.Tr("home.page_dev_env_desc2_title")
	ctx.Data["page_dev_env_desc2_desc"] = ctx.Tr("home.page_dev_env_desc2_desc")
	ctx.Data["page_dev_env_desc3_title"] = ctx.Tr("home.page_dev_env_desc3_title")
	ctx.Data["page_dev_env_desc3_desc"] = ctx.Tr("home.page_dev_env_desc3_desc")
	ctx.Data["page_dev_yunlao"] = ctx.Tr("home.page_dev_yunlao")
	ctx.Data["page_dev_yunlao_desc1"] = ctx.Tr("home.page_dev_yunlao_desc1")
	ctx.Data["page_dev_yunlao_desc2"] = ctx.Tr("home.page_dev_yunlao_desc2")
	ctx.Data["page_dev_yunlao_desc3"] = ctx.Tr("home.page_dev_yunlao_desc3")
	ctx.Data["page_dev_yunlao_desc4"] = ctx.Tr("home.page_dev_yunlao_desc4")
	ctx.Data["page_dev_yunlao_desc5"] = ctx.Tr("home.page_dev_yunlao_desc5")
	ctx.Data["page_dev_yunlao_apply"] = ctx.Tr("home.page_dev_yunlao_apply")
	ctx.Data["page_recommend_activity"] = ctx.Tr("home.page_recommend_activity")
	ctx.Data["page_recommend_activity_desc"] = ctx.Tr("home.page_recommend_activity_desc")
}

func Dashboard(ctx *context.Context) {
	if ctx.IsSigned {
		// 注意：也需要对应概览页面的ActivityImageOverview结构体
		pictureInfo, err := getImageInfo("dashboard-picture")
		if err == nil && len(pictureInfo) > 0 {
			log.Info("set image info=" + pictureInfo[0]["url"])
			ctx.Data["image_url"] = pictureInfo[0]["url"]
			ctx.Data["image_link"] = pictureInfo[0]["image_link"]

			if len(pictureInfo) > 1 {
				ctx.Data["invite_image_url"] = pictureInfo[1]["url"]
				ctx.Data["invite_image_link"] = pictureInfo[1]["image_link"]
			}
		}
		if ctx.User.ProhibitLogin {
			log.Info("Failed authentication attempt for %s from %s", ctx.User.Name, ctx.RemoteAddr())
			ctx.Data["Title"] = ctx.Tr("auth.prohibit_login")
			ctx.HTML(200, "user/auth/prohibit_login")
		} else if ctx.User.MustChangePassword {
			ctx.Data["Title"] = ctx.Tr("auth.must_change_password")
			ctx.Data["ChangePasscodeLink"] = setting.AppSubURL + "/user/change_password"
			ctx.SetCookie("redirect_to", setting.AppSubURL+ctx.Req.URL.RequestURI(), 0, setting.AppSubURL)
			ctx.Redirect(setting.AppSubURL + "/user/settings/change_password")
		} else if setting.PhoneService.Enabled && ctx.User.PhoneNumber == "" {
			ctx.Data["Title"] = ctx.Tr("phone.bind_phone")
			ctx.HTML(200, "user/auth/bind_phone")
			return
		} else {
			user.Dashboard(ctx)
		}
		return
		// Check non-logged users landing page.
	} else if setting.LandingPageURL != setting.LandingPageHome {
		ctx.Redirect(setting.AppSubURL + string(setting.LandingPageURL))
		return
	} else {
		ctx.Redirect(setting.AppSubURL + "/user/login")
		return
	}

	// Check auto-login.
	uname := ctx.GetCookie(setting.CookieUserName)
	if len(uname) != 0 {
		ctx.Redirect(setting.AppSubURL + "/user/login")
		return
	}
	setRecommendURL(ctx)
	ctx.Data["PageIsHome"] = true
	ctx.Data["IsRepoIndexerEnabled"] = setting.Indexer.RepoIndexerEnabled
	ctx.HTML(200, tplHome)
}
func Dashboard1(ctx *context.Context) {
	if ctx.IsSigned {
		// 注意：也需要对应概览页面的ActivityImageOverview结构体
		pictureInfo, err := getImageInfo("dashboard-picture")
		if err == nil && len(pictureInfo) > 0 {
			log.Info("set image info=" + pictureInfo[0]["url"])
			ctx.Data["image_url"] = pictureInfo[0]["url"]
			ctx.Data["image_link"] = pictureInfo[0]["image_link"]

			if len(pictureInfo) > 1 {
				ctx.Data["invite_image_url"] = pictureInfo[1]["url"]
				ctx.Data["invite_image_link"] = pictureInfo[1]["image_link"]
			}
		}
		if ctx.User.ProhibitLogin {
			log.Info("Failed authentication attempt for %s from %s", ctx.User.Name, ctx.RemoteAddr())
			ctx.Data["Title"] = ctx.Tr("auth.prohibit_login")
			ctx.HTML(200, "user/auth/prohibit_login")
		} else if ctx.User.MustChangePassword {
			ctx.Data["Title"] = ctx.Tr("auth.must_change_password")
			ctx.Data["ChangePasscodeLink"] = setting.AppSubURL + "/user/change_password"
			ctx.SetCookie("redirect_to", setting.AppSubURL+ctx.Req.URL.RequestURI(), 0, setting.AppSubURL)
			ctx.Redirect(setting.AppSubURL + "/user/settings/change_password")
		} else if setting.PhoneService.Enabled && ctx.User.PhoneNumber == "" {
			ctx.Data["Title"] = ctx.Tr("phone.bind_phone")
			ctx.HTML(200, "user/auth/bind_phone")
			return
		} else {
			user.Dashboard1(ctx)
		}
		return
		// Check non-logged users landing page.
	} else if setting.LandingPageURL != setting.LandingPageHome {
		ctx.Redirect(setting.AppSubURL + string(setting.LandingPageURL))
		return
	} else {
		ctx.Redirect(setting.AppSubURL + "/user/login")
		return
	}

	// Check auto-login.
	uname := ctx.GetCookie(setting.CookieUserName)
	if len(uname) != 0 {
		ctx.Redirect(setting.AppSubURL + "/user/login")
		return
	}
	setRecommendURL(ctx)
	ctx.Data["PageIsHome"] = true
	ctx.Data["IsRepoIndexerEnabled"] = setting.Indexer.RepoIndexerEnabled
	ctx.HTML(200, tplHome)
}
func OverView(ctx *context.Context) {
	if ctx.IsSigned {
		if ctx.User.ProhibitLogin {
			log.Info("Failed authentication attempt for %s from %s", ctx.User.Name, ctx.RemoteAddr())
			ctx.Data["Title"] = ctx.Tr("auth.prohibit_login")
			ctx.HTML(200, "user/auth/prohibit_login")
		} else if ctx.User.MustChangePassword {
			ctx.Data["Title"] = ctx.Tr("auth.must_change_password")
			ctx.Data["ChangePasscodeLink"] = setting.AppSubURL + "/user/change_password"
			ctx.SetCookie("redirect_to", setting.AppSubURL+ctx.Req.URL.RequestURI(), 0, setting.AppSubURL)
			ctx.Redirect(setting.AppSubURL + "/user/settings/change_password")
		} else if setting.PhoneService.Enabled && ctx.User.PhoneNumber == "" {
			ctx.Data["Title"] = ctx.Tr("phone.bind_phone")
			ctx.HTML(200, "user/auth/bind_phone")
			return
		} else {
			user.OverView(ctx)
		}
		return
	} else if setting.LandingPageURL != setting.LandingPageHome {
		ctx.Redirect(setting.AppSubURL + string(setting.LandingPageURL))
		return
	} else {
		ctx.Redirect(setting.AppSubURL + "/user/login")
		return
	}

	uname := ctx.GetCookie(setting.CookieUserName)
	if len(uname) != 0 {
		ctx.Redirect(setting.AppSubURL + "/user/login")
		return
	}
}

// RepoSearchOptions when calling search repositories
type RepoSearchOptions struct {
	OwnerID    int64
	Private    bool
	Restricted bool
	PageSize   int
	TplName    base.TplName
	Course     util.OptionalBool
}

var (
	nullByte = []byte{0x00}
)

func isKeywordValid(keyword string) bool {
	return !bytes.Contains([]byte(keyword), nullByte)
}

// RenderRepoSearch render repositories search page
func RenderRepoSearch(ctx *context.Context, opts *RepoSearchOptions) {
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}

	var (
		repos   []*models.Repository
		count   int64
		err     error
		orderBy models.SearchOrderBy
	)

	ctx.Data["SortType"] = ctx.Query("sort")
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
	case "reversesize":
		orderBy = models.SearchOrderBySizeReverse
	case "size":
		orderBy = models.SearchOrderBySize
	case "moststars":
		orderBy = models.SearchOrderByStarsReverse
	case "feweststars":
		orderBy = models.SearchOrderByStars
	case "mostforks":
		orderBy = models.SearchOrderByForksReverse
	case "fewestforks":
		orderBy = models.SearchOrderByForks
	case "hot":
		orderBy = models.SearchOrderByHot
	case "active":
		orderBy = models.SearchOrderByActive

	default:
		ctx.Data["SortType"] = "hot"
		orderBy = models.SearchOrderByHot
	}
	orderBy = orderBy + ",id"
	//todo:support other topics
	keyword := strings.Trim(ctx.Query("q"), " ")
	topic := strings.Trim(ctx.Query("topic"), " ")

	repos, count, err = models.SearchRepository(&models.SearchRepoOptions{
		ListOptions: models.ListOptions{
			Page:     page,
			PageSize: opts.PageSize,
		},
		Actor:              ctx.User,
		OrderBy:            orderBy,
		Private:            opts.Private,
		Keyword:            keyword,
		OwnerID:            opts.OwnerID,
		AllPublic:          true,
		AllLimited:         true,
		TopicName:          topic,
		IncludeDescription: setting.UI.SearchRepoDescription,
		Course:             opts.Course,
	})
	if err != nil {
		ctx.ServerError("SearchRepository", err)
		return
	}

	for _, repo := range repos {
		repo.Hot = int64(repo.NumWatches) + int64(repo.NumStars) + int64(repo.NumForks) + int64(repo.CloneCnt)
		repo.Active = int64(repo.NumIssues) + int64(repo.NumPulls) + int64(repo.NumCommit)
	}
	ctx.Data["Keyword"] = keyword
	ctx.Data["Topic"] = topic

	ctx.Data["Total"] = count
	ctx.Data["Repos"] = repos
	ctx.Data["IsRepoIndexerEnabled"] = setting.Indexer.RepoIndexerEnabled

	pager := context.NewPagination(int(count), opts.PageSize, page, 5)
	pager.SetDefaultParams(ctx)
	pager.AddParam(ctx, "topic", "TopicOnly")
	ctx.Data["Page"] = pager

	recommendOrgs, err := models.GetRecommendOrgInfos()
	if err != nil {
		log.Error("GetRecommendOrgInfos failed:%v", err.Error(), ctx.Data["MsgID"])
		ctx.ServerError("GetRecommendOrgInfos", err)
		return
	}
	ctx.Data["RecommendOrgs"] = recommendOrgs

	ctx.HTML(http.StatusOK, opts.TplName)
}

// ExploreRepos render explore repositories page
func ExploreRepos(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("explore")
	ctx.Data["PageIsExplore"] = true
	ctx.Data["PageIsExploreRepositories"] = true
	ctx.Data["IsRepoIndexerEnabled"] = setting.Indexer.RepoIndexerEnabled
	pictureInfo, err := getImageInfo("explore-user-picture")
	if err == nil && len(pictureInfo) > 0 {
		ctx.Data["image_url"] = pictureInfo[0]["url"]
		ctx.Data["image_link"] = pictureInfo[0]["image_link"]
	}
	var ownerID int64
	if ctx.User != nil && !ctx.User.IsAdmin {
		ownerID = ctx.User.ID
	}

	RenderRepoSearch(ctx, &RepoSearchOptions{
		PageSize: setting.UI.ExplorePagingNum,
		OwnerID:  ownerID,
		Private:  ctx.User != nil,
		TplName:  tplExploreRepos,
	})
}

func GetRepoSquarePage(ctx *context.Context) {
	ctx.Data["SquareBanners"] = repository.GetBanners()
	ctx.Data["SquareTopics"] = repository.GetTopics()
	ctx.Data["SquareRecommendRepos"] = repository.GetRecommendRepos()

	repos, _ := repository.GetPreferredRepos()
	ctx.Data["SquarePreferredRepos"] = repos
	ctx.HTML(200, tplRepoSquare)
}
func GetRepoSearchPage(ctx *context.Context) {
	ctx.Data["SquareTopics"] = repository.GetTopics()
	ctx.HTML(200, tplRepoSearch)
}

func RepoSquare(ctx *context.Context) {
	var result []*models.Repository4Card
	var err error
	switch ctx.Query("type") {
	case "preferred":
		result, err = repository.GetPreferredRepos()
	case "incubation":
		result, err = repository.GetIncubationRepos()
	case "hot-paper":
		result, err = repository.GetHotPaperRepos()
	default:
		result, err = repository.GetPreferredRepos()
	}
	if err != nil {
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}
	resultMap := make(map[string]interface{}, 0)
	resultMap["Repos"] = result
	ctx.JSON(http.StatusOK, response.SuccessWithData(resultMap))
}

func ActiveUser(ctx *context.Context) {
	var err error
	var currentUserId int64
	if ctx.User != nil {
		currentUserId = ctx.User.ID
	}
	result, err := repository.GetActiveUser4Square(currentUserId)
	if err != nil {
		log.Error("ActiveUser err. %v", err)
		ctx.JSON(http.StatusOK, response.Success())
		return
	}
	resultMap := make(map[string]interface{}, 0)
	resultMap["Users"] = result
	ctx.JSON(http.StatusOK, response.SuccessWithData(resultMap))
}
func ActiveOrg(ctx *context.Context) {
	result, err := repository.GetActiveOrgs()
	if err != nil {
		log.Error("ActiveOrg err. %v", err)
		ctx.JSON(http.StatusOK, response.Success())
		return
	}
	resultMap := make(map[string]interface{}, 0)
	resultMap["Orgs"] = result
	ctx.JSON(http.StatusOK, response.SuccessWithData(resultMap))
}

func RepoFind(ctx *context.Context) {
	keyword := strings.Trim(ctx.Query("q"), " ")
	topic := strings.Trim(ctx.Query("topic"), " ")
	sort := strings.Trim(ctx.Query("sort"), " ")
	page := ctx.QueryInt("page")
	pageSize := ctx.QueryInt("pageSize")
	if pageSize == 0 {
		pageSize = 15
	}
	if pageSize > 100 {
		ctx.JSON(http.StatusOK, response.ServerError("pageSize illegal"))
		return
	}
	if page <= 0 {
		page = 1
	}

	var ownerID int64
	if ctx.User != nil && !ctx.User.IsAdmin {
		ownerID = ctx.User.ID
	}

	topics, isIntersect, err := handleTopic4RepoFind(topic)
	if err != nil {
		log.Error("RepoFind handleTopic4RepoFind err.topicStr=%s err=%v", topic, err)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}

	result, err := repository.FindRepos(repository.FindReposOptions{
		ListOptions:      models.ListOptions{Page: page, PageSize: pageSize},
		Actor:            ctx.User,
		Sort:             sort,
		Keyword:          keyword,
		Topics:           topics,
		IsTopicIntersect: isIntersect,
		Private:          false,
		OwnerID:          ownerID,
	})
	if err != nil {
		log.Error("RepoFind error. %v", err)
		ctx.JSON(http.StatusOK, response.ResponseError(err))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessWithData(result))
}

const UnionSeparator = "|"
const IntersectSeparator = "*"

func handleTopic4RepoFind(topicStr string) ([]string, bool, error) {
	topicStr = strings.ReplaceAll(topicStr, " ", "")
	if topicStr == "" {
		return []string{}, false, nil
	}
	if strings.Contains(topicStr, UnionSeparator) && strings.Contains(topicStr, IntersectSeparator) {
		return nil, false, errors.New("topic format error")
	}
	if strings.Contains(topicStr, IntersectSeparator) {
		return strings.Split(topicStr, IntersectSeparator), true, nil
	}
	return strings.Split(topicStr, UnionSeparator), false, nil

}

func ExploreMyDatasets(ctx *context.Context) {
	opts := &models.SearchDatasetOptions{
		UploadAttachmentByMe: true,
		NeedAttachment:       false,
		CloudBrainType:       -1,
		SearchOrderBy:        getDatasetOrderBy(ctx),
		JustNeedZipFile:      false,
		DatasetIDs:           models.GetOwnedDatasetIdsByUserID(ctx.User.ID),
		ExploreMine:          true,
	}
	repo.DatasetMultiple(ctx, opts)
}
func ExploreFavoriteDatasets(ctx *context.Context) {
	opts := &models.SearchDatasetOptions{
		StarByMe:        true,
		DatasetIDs:      models.GetDatasetIdsStarByUser(ctx.User.ID),
		NeedAttachment:  false,
		CloudBrainType:  -1,
		SearchOrderBy:   getDatasetOrderBy(ctx),
		JustNeedZipFile: false,
	}
	repo.DatasetMultiple(ctx, opts)

}
func ExploreDatasets(ctx *context.Context) {
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}

	var (
		datasets []*models.Dataset
		count    int64
		err      error
		orderBy  models.SearchOrderBy
	)
	orderBy = getDatasetOrderBy(ctx)

	keyword := strings.Trim(ctx.Query("q"), " ")

	category := ctx.Query("category")
	task := ctx.Query("task")
	license := ctx.Query("license")
	var ownerID int64
	if ctx.User != nil && !ctx.User.IsAdmin {
		ownerID = ctx.User.ID
	}
	var datasetsIds []int64
	if ownerID > 0 {

		collaboratorDatasetsIds := models.GetCollaboratorDatasetIdsByUserID(ownerID)
		teamDatasetsIds := models.GetTeamDatasetIdsByUserID(ownerID)
		datasetsIds = append(collaboratorDatasetsIds, teamDatasetsIds...)
	}

	opts := &models.SearchDatasetOptions{
		Keyword:        keyword,
		IncludePublic:  true,
		SearchOrderBy:  orderBy,
		Category:       category,
		Task:           task,
		License:        license,
		OwnerID:        ownerID,
		DatasetIDs:     datasetsIds,
		RecommendOnly:  ctx.QueryBool("recommend"),
		CloudBrainType: -1,
		ListOptions: models.ListOptions{
			Page:     page,
			PageSize: 30,
		},
	}

	datasets, count, err = models.SearchDataset(opts)

	if err != nil {
		ctx.ServerError("SearchDatasets", err)
		return
	}

	datasetsWithStar := repository.ConvertToDatasetWithStar(ctx, datasets)

	data, err := json.Marshal(datasetsWithStar)

	if err != nil {
		log.Error("json.Marshal failed:", err.Error())
		ctx.JSON(200, map[string]string{
			"result_code": "-1",
			"error_msg":   err.Error(),
			"data":        "",
		})
		return
	}
	ctx.JSON(200, map[string]string{
		"result_code": "0",
		"data":        string(data),
		"count":       strconv.FormatInt(count, 10),
	})

}

func ExploreDatasetsUI(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("explore")
	ctx.Data["PageIsExplore"] = true
	ctx.Data["PageIsExploreDatasets"] = true
	ctx.Data["PageIsDatasets"] = true
	ctx.HTML(200, tplExploreDataset)
}

func ExploreDatasetsDetailUI(ctx *context.Context) {
	ctx.Data["max_model_size"] = int64(setting.MaxModelSize * MODEL_MAX_SIZE)
	ctx.HTML(200, tplExploreDatasetDetail)
}
func ExploreModelsUI(ctx *context.Context) {
	ctx.HTML(200, tplExploreModel)
}
func ExploreModelsDetailUI(ctx *context.Context) {
	ctx.Data["max_model_size"] = int64(setting.MaxModelSize * MODEL_MAX_SIZE)
	ctx.HTML(200, tplExploreModelDetail)
}
func CenterMapUI(ctx *context.Context) {

	ctx.HTML(200, tplExploreCenterMap)
}

func C2NetUI(ctx *context.Context) {

	ctx.HTML(200, tplExploreC2Net)
}

func getDatasetOrderBy(ctx *context.Context) models.SearchOrderBy {
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
	case "reversesize":
		orderBy = models.SearchOrderBySizeReverse
	case "downloadtimes":
		orderBy = models.SearchOrderByDownloadTimes
	case "moststars":
		orderBy = models.SearchOrderByStarsReverse
	case "feweststars":
		orderBy = models.SearchOrderByStars
	case "mostusecount":
		orderBy = models.SearchOrderByUseCountReverse
	case "fewestusecount":
		orderBy = models.SearchOrderByUseCount
	case "default":
		orderBy = models.SearchOrderByDefault
	default:
		ctx.Data["SortType"] = "default"
		orderBy = models.SearchOrderByDefault
	}
	return orderBy
}

// RenderUserSearch render user search page
func RenderUserSearch(ctx *context.Context, opts *models.SearchUserOptions, tplName base.TplName) {
	opts.Page = ctx.QueryInt("page")
	if opts.Page <= 1 {
		opts.Page = 1
	}

	var (
		users   []*models.User
		count   int64
		err     error
		orderBy models.SearchOrderBy
	)

	ctx.Data["SortType"] = ctx.Query("sort")
	switch ctx.Query("sort") {
	case "newest":
		orderBy = models.SearchOrderByIDReverse
	case "oldest":
		orderBy = models.SearchOrderByID
	case "recentupdate":
		orderBy = models.SearchOrderByRecentUpdated
	case "leastupdate":
		orderBy = models.SearchOrderByLeastUpdated
	case "reversealphabetically":
		orderBy = models.SearchOrderByAlphabeticallyReverse
	case "alphabetically":
		orderBy = models.SearchOrderByAlphabetically
	default:
		ctx.Data["SortType"] = "alphabetically"
		orderBy = models.SearchOrderByAlphabetically
	}

	opts.Keyword = strings.Trim(ctx.Query("q"), " ")
	opts.OrderBy = orderBy
	if len(opts.Keyword) == 0 || isKeywordValid(opts.Keyword) {
		users, count, err = models.SearchUsers(opts)
		if err != nil {
			ctx.ServerError("SearchUsers", err)
			return
		}
		// for _, user := range users {
		// 	//user.IsSubscriber = role.UserHasRole(user.ID, models.Subscriber)
		// 	user.OperRole, user.ResourceRole = role.QueryUserRole(user.ID)
		// }

	}
	ctx.Data["Keyword"] = opts.Keyword
	ctx.Data["Total"] = count
	ctx.Data["Users"] = users
	ctx.Data["ShowUserEmail"] = setting.UI.ShowUserEmail
	ctx.Data["IsRepoIndexerEnabled"] = setting.Indexer.RepoIndexerEnabled

	pager := context.NewPagination(int(count), opts.PageSize, opts.Page, 5)
	pager.SetDefaultParams(ctx)
	pager.AddParam(ctx, "adminType", "AdminType")
	pager.AddParam(ctx, "subscriberType", "SubscriberType")
	ctx.Data["Page"] = pager

	ctx.HTML(200, tplName)
}

// ExploreUsers render explore users page
func ExploreUsers(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("explore")
	ctx.Data["PageIsExplore"] = true
	ctx.Data["PageIsExploreUsers"] = true
	ctx.Data["IsRepoIndexerEnabled"] = setting.Indexer.RepoIndexerEnabled
	pictureInfo, err := getImageInfo("explore-user-picture")
	if err == nil && len(pictureInfo) > 0 {
		ctx.Data["image_url"] = pictureInfo[0]["url"]
		ctx.Data["image_link"] = pictureInfo[0]["image_link"]
	}
	RenderUserSearch(ctx, &models.SearchUserOptions{
		Actor:       ctx.User,
		Type:        models.UserTypeIndividual,
		ListOptions: models.ListOptions{PageSize: setting.UI.ExplorePagingNum},
		IsActive:    util.OptionalBoolTrue,
		Visible:     []structs.VisibleType{structs.VisibleTypePublic, structs.VisibleTypeLimited, structs.VisibleTypePrivate},
	}, tplExploreUsers)
}

// ExploreOrganizations render explore organizations page
func ExploreOrganizations(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("explore")
	ctx.Data["PageIsExplore"] = true
	ctx.Data["PageIsExploreOrganizations"] = true
	ctx.Data["IsRepoIndexerEnabled"] = setting.Indexer.RepoIndexerEnabled

	N := 10
	starInfo, err := models.FindTopNStarsOrgs(N)
	if err != nil {
		log.Error("GetStarOrgInfos failed:%v", err.Error(), ctx.Data["MsgID"])
		ctx.ServerError("GetStarOrgInfos", err)
		return
	}
	memberInfo, err := models.FindTopNMembersOrgs(N)
	if err != nil {
		log.Error("GetMemberOrgInfos failed:%v", err.Error(), ctx.Data["MsgID"])
		ctx.ServerError("GetMemberOrgInfos", err)
		return
	}
	openIInfo, err := models.FindTopNOpenIOrgs(N)
	if err != nil {
		log.Error("GetOpenIOrgInfos failed:%v", err.Error(), ctx.Data["MsgID"])
		ctx.ServerError("GetOpenIOrgInfos", err)
		return
	}

	recommendOrgs, err := getRecommendOrg()
	if err != nil {
		log.Error("GetRecommendOrgInfos failed:%v", err.Error(), ctx.Data["MsgID"])
		ctx.ServerError("GetRecommendOrgInfos", err)
		return
	}
	setRecommendURLOnly(ctx)
	ctx.Data["RecommendOrgs"] = recommendOrgs
	ctx.Data["StarOrgs"] = starInfo
	ctx.Data["MemberOrgs"] = memberInfo
	ctx.Data["ActiveOrgs"] = openIInfo

	ctx.HTML(http.StatusOK, tplExploreOrganizations)
}

// ExploreCode render explore code page
func ExploreCode(ctx *context.Context) {
	if !setting.Indexer.RepoIndexerEnabled {
		ctx.Redirect(setting.AppSubURL+"/explore", 302)
		return
	}

	ctx.Data["IsRepoIndexerEnabled"] = setting.Indexer.RepoIndexerEnabled
	ctx.Data["Title"] = ctx.Tr("explore")
	ctx.Data["PageIsExplore"] = true
	ctx.Data["PageIsExploreCode"] = true

	language := strings.TrimSpace(ctx.Query("l"))
	keyword := strings.TrimSpace(ctx.Query("q"))
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}

	var (
		repoIDs []int64
		err     error
		isAdmin bool
		userID  int64
	)
	if ctx.User != nil {
		userID = ctx.User.ID
		isAdmin = ctx.User.IsAdmin
	}

	// guest user or non-admin user
	if ctx.User == nil || !isAdmin {
		repoIDs, err = models.FindUserAccessibleRepoIDs(ctx.User)
		if err != nil {
			ctx.ServerError("SearchResults", err)
			return
		}
	}

	var (
		total                 int
		searchResults         []*code_indexer.Result
		searchResultLanguages []*code_indexer.SearchResultLanguages
	)

	// if non-admin login user, we need check UnitTypeCode at first
	if ctx.User != nil && len(repoIDs) > 0 {
		repoMaps, err := models.GetRepositoriesMapByIDs(repoIDs)
		if err != nil {
			ctx.ServerError("SearchResults", err)
			return
		}

		var rightRepoMap = make(map[int64]*models.Repository, len(repoMaps))
		repoIDs = make([]int64, 0, len(repoMaps))
		for id, repo := range repoMaps {
			if repo.CheckUnitUser(userID, isAdmin, models.UnitTypeCode) {
				rightRepoMap[id] = repo
				repoIDs = append(repoIDs, id)
			}
		}

		ctx.Data["RepoMaps"] = rightRepoMap

		total, searchResults, searchResultLanguages, err = code_indexer.PerformSearch(repoIDs, language, keyword, page, setting.UI.RepoSearchPagingNum)
		if err != nil {
			ctx.ServerError("SearchResults", err)
			return
		}
		// if non-login user or isAdmin, no need to check UnitTypeCode
	} else if (ctx.User == nil && len(repoIDs) > 0) || isAdmin {
		total, searchResults, searchResultLanguages, err = code_indexer.PerformSearch(repoIDs, language, keyword, page, setting.UI.RepoSearchPagingNum)
		if err != nil {
			ctx.ServerError("SearchResults", err)
			return
		}

		var loadRepoIDs = make([]int64, 0, len(searchResults))
		for _, result := range searchResults {
			var find bool
			for _, id := range loadRepoIDs {
				if id == result.RepoID {
					find = true
					break
				}
			}
			if !find {
				loadRepoIDs = append(loadRepoIDs, result.RepoID)
			}
		}

		repoMaps, err := models.GetRepositoriesMapByIDs(loadRepoIDs)
		if err != nil {
			ctx.ServerError("SearchResults", err)
			return
		}

		ctx.Data["RepoMaps"] = repoMaps
	}

	ctx.Data["Keyword"] = keyword
	ctx.Data["Language"] = language
	ctx.Data["SearchResults"] = searchResults
	ctx.Data["SearchResultLanguages"] = searchResultLanguages
	ctx.Data["RequireHighlightJS"] = true
	ctx.Data["PageIsViewCode"] = true

	pager := context.NewPagination(total, setting.UI.RepoSearchPagingNum, page, 5)
	pager.SetDefaultParams(ctx)
	pager.AddParam(ctx, "l", "Language")
	ctx.Data["Page"] = pager

	ctx.HTML(200, tplExploreCode)
}

func ExploreImages(ctx *context.Context) {
	ctx.HTML(200, tplExploreImages)
}

func ExploreOrgMy(ctx *context.Context) {
	ctx.HTML(200, tplExploreOrg)
}
func ComputingPowerDemand(ctx *context.Context) {
	ctx.HTML(200, tplComputingPowerDemand)
}

func ComputingPowerDomestic(ctx *context.Context) {
	ctx.HTML(200, tplComputingPowerDomestic)
}

func IdeProject(ctx *context.Context) {
	ctx.Data["url_params"] = time.Now().Unix()
	ctx.HTML(200, tplIde)
}

func ExploreRosHmci(ctx *context.Context) {
	ctx.HTML(200, tplRoshmci)
}

// NotFound render 404 page
func NotFound(ctx *context.Context) {
	ctx.Data["Title"] = "Page Not Found"
	ctx.NotFound("home.NotFound", nil)
}

func getRecommendOrg() ([]map[string]interface{}, error) {
	url := setting.RecommentRepoAddr + "home/organizations"
	result, err := repository.RecommendFromPromote(url)

	if err != nil {
		return nil, err
	}
	names := make([]string, 0)
	for _, userName := range result {
		names = append(names, userName)
	}
	users, _ := models.GetUsersByNames(names)
	userMap := make(map[string]*models.User, 0)
	for _, user := range users {
		userMap[user.Name] = user
	}
	resultOrg := make([]map[string]interface{}, 0)
	for _, userName := range result {
		user := userMap[userName]
		if user != nil {
			userMap := make(map[string]interface{})
			userMap["Name"] = user.Name
			userMap["Description"] = user.Description
			userMap["FullName"] = user.FullName
			userMap["HomeLink"] = user.HomeLink()
			userMap["ID"] = user.ID
			userMap["Avatar"] = user.RelAvatarLink()
			userMap["NumRepos"] = user.NumRepos
			userMap["NumTeams"] = user.NumTeams
			userMap["NumMembers"] = user.NumMembers
			resultOrg = append(resultOrg, userMap)
		} else {
			log.Info("the user not exist," + userName)
		}
	}
	return resultOrg, nil
}

func getImageInfo(filename string) ([]map[string]string, error) {
	url := setting.RecommentRepoAddr + filename
	result, err := repository.RecommendFromPromote(url)

	if err != nil {
		return nil, err
	}
	imageInfo := make([]map[string]string, 0)
	for i := 0; i < (len(result) - 1); i++ {
		line := result[i]
		imageMap := make(map[string]string)
		if line[0:4] == "url=" {
			url := line[4:]
			imageMap["url"] = url
			if result[i+1][0:11] == "image_link=" {
				image_link := result[i+1][11:]
				imageMap["image_link"] = image_link
			}
		}
		imageInfo = append(imageInfo, imageMap)
		i = i + 1
	}
	return imageInfo, nil
}

func GetMapInfo(ctx *context.Context) {
	filename := ctx.Query("filename")
	if strings.HasPrefix(filename, "ad-") || strings.HasPrefix(filename, "invita") {
		ctx.JSON(http.StatusOK, "{}")
		return
	}
	url := setting.RecommentRepoAddr + filename
	result, err := repository.RecommendContentFromPromote(url)
	if err != nil {
		log.Info("get file error:" + err.Error())
	}
	ctx.JSON(http.StatusOK, result)
}

func GetRankUser(index string) ([]map[string]interface{}, error) {
	url := setting.RecommentRepoAddr + "user_rank/user_rank_" + index
	result, err := repository.RecommendFromPromote(url)

	if err != nil {
		return nil, err
	}
	resultOrg := make([]map[string]interface{}, 0)
	for _, userRank := range result {
		tmpIndex := strings.Index(userRank, " ")
		userName := userRank
		score := 0
		label := ""
		if tmpIndex != -1 {
			userName = userRank[0:tmpIndex]
			left := userRank[tmpIndex+1:]
			tmpIndex1 := strings.Index(left, " ")
			if tmpIndex1 != -1 {
				tmpScore, err := strconv.Atoi(left[0:tmpIndex1])
				if err != nil {
					log.Info("convert to int error.")
				}
				score = tmpScore
				label = left[tmpIndex1+1:]
			} else {
				tmpScore, err := strconv.Atoi(left[tmpIndex+1:])
				if err != nil {
					log.Info("convert to int error.")
				}
				score = tmpScore
			}
		}
		user, err := models.GetUserByName(userName)
		if err == nil {
			userMap := make(map[string]interface{})
			userMap["Name"] = user.Name
			userMap["Description"] = user.Description
			userMap["FullName"] = user.FullName
			userMap["HomeLink"] = user.HomeLink()
			userMap["ID"] = user.ID
			userMap["Label"] = label
			userMap["Avatar"] = user.RelAvatarLink()
			userMap["Score"] = score
			resultOrg = append(resultOrg, userMap)
		} else {
			log.Info("query user error," + err.Error())
		}
	}
	return resultOrg, nil
}

func GetUserRankFromPromote(ctx *context.Context) {
	index := ctx.Params("index")
	resultUserRank, err := GetRankUser(index)
	if err != nil {
		ctx.ServerError("500", err)
		return
	}
	ctx.JSON(200, resultUserRank)
}

func getMapContent(fileName string) []map[string]string {
	url := setting.RecommentRepoAddr + fileName
	result, err := repository.RecommendContentFromPromote(url)
	remap := make([]map[string]string, 0)
	if err == nil {
		json.Unmarshal([]byte(result), &remap)
	}
	return remap
}

func HomeNoticeTmpl(ctx *context.Context) {
	ctx.Data["url_params"] = ""
	ctx.HTML(200, "notice")
}

func RecommendHomeInfo(ctx *context.Context) {
	resultOrg, err := getRecommendOrg()
	if err != nil {
		log.Info("error." + err.Error())
	}
	repoMap := getMapContent("home/projects")
	resultRepo, err := repository.GetRecommendRepoFromPromote(repoMap)
	if err != nil {
		log.Info("error." + err.Error())
	}
	resultActivityInfo := getMapContent("home/activity_info")
	mapInterface := make(map[string]interface{})
	mapInterface["org"] = resultOrg
	mapInterface["repo"] = resultRepo
	mapInterface["activity"] = resultActivityInfo

	user_experience := getMapContent("home/user_experience")
	for _, amap := range user_experience {
		userId := amap["userid"]
		userIntId, _ := strconv.Atoi(userId)
		user, err := models.GetUserByID(int64(userIntId))
		if err == nil {
			amap["name"] = user.Name
			amap["fullname"] = user.FullName
			amap["detail"] = user.Description
			amap["avatar"] = user.AvatarLink()
		}
	}
	mapInterface["user_experience"] = user_experience
	dataset, err := models.QueryDatasetGroupByTask(7)
	if err == nil {
		mapInterface["dataset"] = dataset
	}
	ctx.JSON(http.StatusOK, mapInterface)
}

func HomeTerm(ctx *context.Context) {
	ctx.HTML(200, tplHomeTerm)
}
func HomeAnnual(ctx *context.Context) {
	ctx.HTML(200, tplHomeAnnual)
}
func HomeWenxin(ctx *context.Context) {
	ctx.HTML(200, tplHomeWenxin)
}
func HomePrivacy(ctx *context.Context) {
	ctx.HTML(200, tplHomePrivacy)
}
func HomeAuthTerm(ctx *context.Context) {
	ctx.HTML(200, tplHomeAuthTerm)
}

func HomeResoruceDesc(ctx *context.Context) {
	ctx.HTML(200, tplResoruceDesc)
}
