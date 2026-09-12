// Copyright 2015 The Gogs Authors. All rights reserved.
// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

import (
	"code.gitea.io/gitea/services/badge"
	"errors"
	"fmt"
	"path"
	"strings"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/util"
	"code.gitea.io/gitea/routers/org"
)

// GetUserByName get user by name
func GetUserByName(ctx *context.Context, name string) *models.User {
	user, err := models.GetUserByName(name)
	if err != nil {
		if models.IsErrUserNotExist(err) {
			ctx.NotFound("GetUserByName", nil)
		} else {
			ctx.ServerError("GetUserByName", err)
		}
		return nil
	}
	return user
}

// GetUserByParams returns user whose name is presented in URL paramenter.
func GetUserByParams(ctx *context.Context) *models.User {
	return GetUserByName(ctx, ctx.Params(":username"))
}

// Profile render user's profile page
func Profile(ctx *context.Context) {
	uname := ctx.Params(":username")
	// Special handle for FireFox requests favicon.ico.
	if uname == "favicon.ico" {
		ctx.ServeFile(path.Join(setting.StaticRootPath, "public/img/favicon.png"))
		return
	} else if strings.HasSuffix(uname, ".png") {
		ctx.Error(404)
		return
	}

	isShowKeys := false
	if strings.HasSuffix(uname, ".keys") {
		isShowKeys = true
		uname = strings.TrimSuffix(uname, ".keys")
	}

	isShowGPG := false
	if strings.HasSuffix(uname, ".gpg") {
		isShowGPG = true
		uname = strings.TrimSuffix(uname, ".gpg")
	}

	ctxUser := GetUserByName(ctx, uname)
	if ctx.Written() {
		return
	}

	// Show SSH keys.
	if isShowKeys {
		ShowSSHKeys(ctx, ctxUser.ID)
		return
	}

	// Show GPG keys.
	if isShowGPG {
		ShowGPGKeys(ctx, ctxUser.ID)
		return
	}

	if ctxUser.IsOrganization() {
		org.Home(ctx)
		return
	}

	// Show OpenID URIs
	openIDs, err := models.GetUserOpenIDs(ctxUser.ID)
	if err != nil {
		ctx.ServerError("GetUserOpenIDs", err)
		return
	}

	// Show user badges
	badges, err := badge.GetUserBadges(ctxUser.ID, models.ListOptions{Page: 1, PageSize: 5})
	if err != nil {
		ctx.ServerError("GetUserBadges", err)
		return
	}
	// Count user badges
	cnt, err := badge.CountUserBadges(ctxUser.ID)
	if err != nil {
		ctx.ServerError("CountUserBadges", err)
		return
	}

	ctx.Data["Title"] = ctxUser.DisplayName()
	ctx.Data["PageIsUserProfile"] = true
	ctx.Data["Owner"] = ctxUser
	ctx.Data["OpenIDs"] = openIDs
	ctx.Data["RecentBadges"] = badges
	ctx.Data["TotalBadges"] = cnt
	ctx.Data["EnableHeatmap"] = setting.Service.EnableUserHeatmap
	ctx.Data["HeatmapUser"] = ctxUser.Name
	showPrivate := ctx.IsSigned && (ctx.User.IsAdmin || ctx.User.ID == ctxUser.ID)

	orgs, err := models.GetOrgsByUserID(ctxUser.ID, showPrivate)
	if err != nil {
		ctx.ServerError("GetOrgsByUserIDDesc", err)
		return
	}

	for _, org := range orgs {
		_, repoCount, err := models.SearchRepository(&models.SearchRepoOptions{
			OwnerID: org.ID,
			Private: ctx.IsSigned,
			Actor:   ctx.User,
		})
		if err != nil {
			ctx.ServerError("SearchRepository", err)
			return
		}

		var opts = models.FindOrgMembersOpts{
			OrgID:      org.ID,
			PublicOnly: true,
		}

		if ctx.User != nil {
			isMember, err := org.IsOrgMember(ctx.User.ID)
			if err != nil {
				ctx.Error(500, "IsOrgMember")
				return
			}
			opts.PublicOnly = !isMember && !ctx.User.IsAdmin
		}

		membersCount, err := models.CountOrgMembers(opts)
		if err != nil {
			ctx.ServerError("CountOrgMembers", err)
			return
		}
		org.NumMembers = int(membersCount)
		org.NumRepos = int(repoCount)
	}

	ctx.Data["Orgs"] = orgs
	ctx.Data["HasOrgsVisible"] = models.HasOrgsVisible(orgs, ctx.User)

	tab := ctx.Query("tab")
	if tab == "" {
		tab = "activity"
	}
	ctx.Data["TabName"] = tab

	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}

	topicOnly := ctx.QueryBool("topic")

	var (
		repos   []*models.Repository
		count   int64
		total   int
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
	case "downloadtimes":
		orderBy = models.SearchOrderByDownloadTimes
	case "moststars":
		orderBy = models.SearchOrderByStarsReverse
	case "feweststars":
		orderBy = models.SearchOrderByStars
	case "mostforks":
		orderBy = models.SearchOrderByForksReverse
	case "fewestforks":
		orderBy = models.SearchOrderByForks
	default:
		ctx.Data["SortType"] = "recentupdate"
		orderBy = models.SearchOrderByRecentUpdated
	}

	keyword := strings.Trim(ctx.Query("q"), " ")
	ctx.Data["Keyword"] = keyword
	switch tab {
	case "followers":
		items, err := ctxUser.GetFollowers(models.ListOptions{
			PageSize: setting.UI.User.RepoPagingNum,
			Page:     page,
		})
		if err != nil {
			ctx.ServerError("GetFollowers", err)
			return
		}
		ctx.Data["Cards"] = items

		total = ctxUser.NumFollowers
	case "following":
		items, err := ctxUser.GetFollowing(models.ListOptions{
			PageSize: setting.UI.User.RepoPagingNum,
			Page:     page,
		})
		if err != nil {
			ctx.ServerError("GetFollowing", err)
			return
		}
		ctx.Data["Cards"] = items

		total = ctxUser.NumFollowing
	case "activity":
		retrieveFeeds(ctx, models.GetFeedsOptions{RequestedUser: ctxUser,
			Actor:           ctx.User,
			IncludePrivate:  showPrivate,
			OnlyPerformedBy: true,
			IncludeDeleted:  false,
		})
		if ctx.Written() {
			return
		}
	case "stars":
		ctx.Data["PageIsProfileStarList"] = true
		repos, count, err = models.SearchRepository(&models.SearchRepoOptions{
			ListOptions: models.ListOptions{
				PageSize: setting.UI.User.RepoPagingNum,
				Page:     page,
			},
			Actor:              ctx.User,
			Keyword:            keyword,
			OrderBy:            orderBy,
			Private:            ctx.IsSigned,
			StarredByID:        ctxUser.ID,
			Collaborate:        util.OptionalBoolFalse,
			TopicOnly:          topicOnly,
			IncludeDescription: setting.UI.SearchRepoDescription,
		})
		if err != nil {
			ctx.ServerError("SearchRepository", err)
			return
		}

		total = int(count)
	case "datasets":
		var isOwner = false
		if ctx.User != nil && ctx.User.ID == ctxUser.ID {
			isOwner = true
		}
		datasetSearchOptions := &models.SearchDatasetOptions{
			Keyword:       keyword,
			OwnerID:       ctxUser.ID,
			SearchOrderBy: orderBy,
			IsOwner:       isOwner,
			ListOptions: models.ListOptions{
				Page:     page,
				PageSize: setting.UI.User.RepoPagingNum,
			},
			CloudBrainType: -1,
		}

		if len(datasetSearchOptions.SearchOrderBy) == 0 {
			datasetSearchOptions.SearchOrderBy = models.SearchOrderByAlphabetically
		}

		datasets, count, err := models.SearchDataset(datasetSearchOptions)
		if err != nil {
			ctx.ServerError("SearchDatasets", err)
		}
		total = int(count)
		ctx.Data["Datasets"] = datasets
	case "repository":
		repos, count, err = models.SearchRepository(&models.SearchRepoOptions{
			ListOptions: models.ListOptions{
				PageSize: setting.UI.User.RepoPagingNum,
				Page:     page,
			},
			Actor:              ctx.User,
			Keyword:            keyword,
			OwnerID:            ctxUser.ID,
			OrderBy:            orderBy,
			Private:            ctx.IsSigned,
			Collaborate:        util.OptionalBoolFalse,
			TopicOnly:          topicOnly,
			IncludeDescription: setting.UI.SearchRepoDescription,
		})
		if err != nil {
			ctx.ServerError("SearchRepository", err)
			return
		}

		total = int(count)
	case "badge":
		allBadges, err := badge.GetUserAllBadges(ctxUser.ID)
		if err != nil {
			ctx.ServerError("GetUserAllBadges", err)
			return
		}
		ctx.Data["AllBadges"] = allBadges
	case "model":

	case "computetasktemplate":

	default:
		ctx.ServerError("tab error", errors.New("tab error"))
		return
	}
	ctx.Data["Repos"] = repos
	ctx.Data["Total"] = total

	pager := context.NewPagination(total, setting.UI.User.RepoPagingNum, page, 5)
	pager.SetDefaultParams(ctx)
	ctx.Data["Page"] = pager

	ctx.Data["ShowUserEmail"] = len(ctxUser.Email) > 0 && ctx.IsSigned && (!ctxUser.KeepEmailPrivate || ctxUser.ID == ctx.User.ID)

	ctx.HTML(200, tplProfile)
}

// Action response for follow/unfollow user request
func Action(ctx *context.Context) {
	u := GetUserByParams(ctx)
	if ctx.Written() {
		return
	}

	var err error
	switch ctx.Params(":action") {
	case "follow":
		err = models.FollowUser(ctx.User.ID, u.ID)
	case "unfollow":
		err = models.UnfollowUser(ctx.User.ID, u.ID)
	}

	if err != nil {
		ctx.ServerError(fmt.Sprintf("Action (%s)", ctx.Params(":action")), err)
		return
	}

	ctx.RedirectToFirst(ctx.Query("redirect_to"), u.HomeLink())
}
