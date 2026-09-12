package badge

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/badge"
	"errors"
	"fmt"
	"github.com/unknwon/com"
	"net/http"
	"strconv"
	"strings"
)

func GetCustomizeBadgeList(ctx *context.Context) {
	page := ctx.QueryInt("page")
	category := ctx.QueryInt64("category")
	pageSize := 50
	n, r, err := badge.GetBadgeList(models.GetBadgeOpts{CategoryId: category, BadgeType: models.CustomizeBadge, ListOpts: models.ListOptions{PageSize: pageSize, Page: page}})
	if err != nil {
		log.Error("GetCustomizeBadgeList error.%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	m := make(map[string]interface{})
	m["List"] = r
	m["Total"] = n
	m["PageSize"] = pageSize
	ctx.JSON(http.StatusOK, response.SuccessWithData(m))
}

func OperateBadge(ctx *context.Context, req models.BadgeOperateReq) {
	action := ctx.Params(":action")

	var err *response.BizError
	switch action {
	case "edit":
		err = badge.EditBadge(req, ctx.User)
	case "new":
		err = badge.AddBadge(req, ctx.User)
	case "del":
		err = badge.DelBadge(req.ID, ctx.User)
	default:
		err = response.NewBizError(errors.New("action type error"))
	}

	if err != nil {
		log.Error("OperateBadge error ,%v", err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(err, ctx))
		return
	}
	ctx.JSON(http.StatusOK, response.Success())
}

func GetBadgeUsers(ctx *context.Context) {
	page := ctx.QueryInt("page")
	badgeId := ctx.QueryInt64("badge")
	pageSize := 50
	n, r, err := badge.GetBadgeUsers(badgeId, models.ListOptions{PageSize: pageSize, Page: page})
	if err != nil {
		log.Error("GetBadgeUsers error.%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	m := make(map[string]interface{})
	m["List"] = r
	m["Total"] = n
	m["PageSize"] = pageSize
	ctx.JSON(http.StatusOK, response.SuccessWithData(m))
}

func AddOperateBadgeUsers(ctx *context.Context, req models.AddBadgeUsersReq) {
	userStr := req.Users
	if userStr == "" {
		ctx.JSON(http.StatusOK, response.Success())
		return
	}
	userStr = strings.ReplaceAll(userStr, " ", "")
	userStr = strings.ReplaceAll(userStr, "\r", "")
	userStr = strings.TrimSuffix(userStr, "\n")
	userIdStrs := strings.Split(userStr, "\n")
	userIds := make([]int64, 0)
	errList := make([]string, 0)

	total := 0
	for _, s := range userIdStrs {
		if s == "" {
			continue
		}
		total++
		id, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			errList = append(errList, s)
			continue
		}
		userIds = append(userIds, id)
	}
	n, list, err := badge.AddBadgeUsers(req.BadgeId, userIds)
	if err != nil {
		log.Error("AddOperateBadgeUsers error.%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}

	if len(list) > 0 {
		for _, l := range list {
			errList = append(errList, fmt.Sprint(l))
		}
	}

	m := make(map[string]interface{})
	m["Total"] = total
	m["Success"] = n
	m["Exists"] = total - n - len(errList)
	m["ErrorList"] = errList
	ctx.JSON(http.StatusOK, response.SuccessWithData(m))
}

func DelBadgeUsers(ctx *context.Context, req models.DelBadgeUserReq) {
	id := req.ID
	if id <= 0 {
		ctx.JSON(http.StatusOK, response.Success())
		return
	}

	err := badge.DelBadgeUser(id)
	if err != nil {
		log.Error("DelBadgeUsers error.%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success())
}

func UploadIcon(ctx *context.Context, form badge.IconUploadForm) {

	uploader := badge.NewIconUploader(badge.IconUploadConfig{
		FileMaxSize:   setting.BadgeIconMaxFileSize,
		FileMaxWidth:  setting.BadgeIconMaxWidth,
		FileMaxHeight: setting.BadgeIconMaxHeight,
		NeedSquare:    true,
	})
	iconName, err := uploader.Upload(form, ctx.User)
	if err != nil {
		log.Error("UploadIcon error.%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	m := make(map[string]string, 0)
	m["IconName"] = iconName
	ctx.JSON(http.StatusOK, response.SuccessWithData(m))
}

func GetIcon(ctx *context.Context) {
	hash := ctx.Params(":hash")
	if !com.IsFile(models.GetCustomIconByHash(hash)) {
		ctx.NotFound(ctx.Req.URL.RequestURI(), nil)
		return
	}
	ctx.Redirect(setting.AppSubURL + "/icons/" + hash)
}
