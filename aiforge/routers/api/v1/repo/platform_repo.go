package repo

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/modules/util"
	"code.gitea.io/gitea/routers/api/v1/utils"
	"fmt"
	"net/http"
	"strings"
)

// PlatFormRepoSearch 工作台的项目页面查询
func PlatFormRepoSearch(ctx *context.APIContext) {
	opts := buildParams(ctx)

	// 查询对应的仓库
	repos, count, err := models.SearchRepository(opts)
	if err != nil {
		log.Error("PlatFormRepoSearch SearchRepository reqId [%v] err [%+v]", ctx.ReqId, err)
		ctx.JSON(http.StatusInternalServerError, api.SearchError{
			OK:    false,
			Error: err.Error(),
		})
		return
	}

	results := make([]*models.Repository4Card, len(repos))

	for i, repo := range repos {
		if err = repo.GetOwner(); err != nil {
			log.Error("PlatFormRepoSearch GetOwner reqId [%v] err [%+v]", ctx.ReqId, err)
			ctx.JSON(http.StatusInternalServerError, api.SearchError{
				OK:    false,
				Error: err.Error(),
			})
			return
		}

		t := repo.ToCardFormat()
		results[i] = t

	}

	// 补充置顶记录
	if err = appendPinnedRecord(results, ctx.User.ID); err != nil {
		// 补充失败就算了，还是正常返回
		log.Error("PlatFormRepoSearch appendPinnedRecord reqId [%v] err [%+v]", ctx.ReqId, err)
	}

	ctx.JSON(http.StatusOK, models.PlatformSearchResults{
		Data: models.PlatformSearch{
			Repos:    results,
			Total:    count,
			Page:     opts.Page,
			PageSize: opts.PageSize,
		},
	})
	return
}

// appendPinnedRecord 批量添加置顶记录
func appendPinnedRecord(record []*models.Repository4Card, userId int64) error {
	// 查询置顶关系
	pinnedMap, err := models.GetPinnedRecordByUserId(userId)
	if err != nil {
		return err
	}
	if len(pinnedMap) == 0 {
		return nil
	}

	for i, repo := range record {
		record[i].HasPinned = pinnedMap[repo.ID] != nil
	}

	return nil
}

// buildParams 构建参数
func buildParams(ctx *context.APIContext) (opts *models.SearchRepoOptions) {
	opts = &models.SearchRepoOptions{
		ListOptions:        utils.GetListOptions(ctx),         // 分页
		Actor:              ctx.User,                          // 用户
		Keyword:            strings.Trim(ctx.Query("q"), " "), // 查询的字段
		OwnerID:            ctx.QueryInt64("uid"),             // 项目归属人
		PriorityOwnerID:    ctx.QueryInt64("priority_owner_id"),
		TopicOnly:          ctx.QueryBool("topic"),
		Collaborate:        util.OptionalBoolNone,
		Private:            ctx.IsSigned && (ctx.Query("private") == "" || ctx.QueryBool("private")),
		OnlyPrivate:        ctx.IsSigned && ctx.QueryBool("onlyPrivate"),
		Template:           util.OptionalBoolNone,
		StarredByID:        ctx.QueryInt64("starredBy"),
		IncludeDescription: ctx.QueryBool("includeDesc"),
		HasPinned:          true, // 存在置顶
	}

	if ctx.Query("template") != "" {
		opts.Template = util.OptionalBoolOf(ctx.QueryBool("template"))
	}

	if ctx.QueryBool("exclusive") {
		opts.Collaborate = util.OptionalBoolFalse
	}

	var mode = ctx.Query("mode")
	switch mode {
	case "source":
		opts.Fork = util.OptionalBoolFalse
		opts.Mirror = util.OptionalBoolFalse
	case "fork":
		opts.Fork = util.OptionalBoolTrue
	case "mirror":
		opts.Mirror = util.OptionalBoolTrue
	case "collaborative":
		opts.Collaborate = util.OptionalBoolTrue
	case "":
	default:
		ctx.Error(http.StatusUnprocessableEntity, "", fmt.Errorf("Invalid search mode: \"%s\"", mode))
		return
	}

	if ctx.Query("archived") != "" {
		opts.Archived = util.OptionalBoolOf(ctx.QueryBool("archived"))
	}

	var sortMode = ctx.Query("sort")
	if len(sortMode) > 0 {
		var sortOrder = ctx.Query("order")
		if len(sortOrder) == 0 {
			sortOrder = "asc"
		}
		if searchModeMap, ok := searchOrderByMap[sortOrder]; ok {
			if orderBy, ok := searchModeMap[sortMode]; ok {
				opts.OrderBy = orderBy
			} else {
				ctx.Error(http.StatusUnprocessableEntity, "", fmt.Errorf("Invalid sort mode: \"%s\"", sortMode))
				return
			}
		} else {
			ctx.Error(http.StatusUnprocessableEntity, "", fmt.Errorf("Invalid sort order: \"%s\"", sortOrder))
			return
		}
	}

	return
}
