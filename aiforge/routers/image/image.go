package image

import (
	"net/http"
	"strconv"
	"strings"

	"code.gitea.io/gitea/modules/auth"

	"code.gitea.io/gitea/modules/notification"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
)

func UserAction(ctx *context.Context) {
	imageId, _ := strconv.ParseInt(ctx.Params(":id"), 10, 64)
	image, err := models.GetImageByID(imageId)
	if err != nil {
		ctx.Error(http.StatusNotFound)
		return
	}
	switch ctx.Params(":action") {

	case "star":
		err = models.StarImage(ctx.User.ID, imageId, true)
	case "unstar":
		err = models.StarImage(ctx.User.ID, imageId, false)
	}
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("repo.star_fail", ctx.Params(":action"))))
	} else {
		notification.NotifyImageRecommend(ctx.User, image, ctx.Params(":action"))
		ctx.JSON(http.StatusOK, models.BaseOKMessage)
	}
}
func Action(ctx *context.Context, form auth.ReviewImageForm) {
	imageId, _ := strconv.ParseInt(ctx.Params(":id"), 10, 64)
	image, err := models.GetImageByID(imageId)
	if err != nil {
		ctx.Error(http.StatusNotFound)
		return
	}
	switch ctx.Params(":action") {
	case "recommend":
		if !ctx.User.IsAdmin {
			ctx.Error(http.StatusForbidden)
			return
		}
		if image.ApplyStatus == models.FailApply {
			ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("repo.failApplyCannotRecommend")))
			return
		}
		if image.ApplyStatus == models.Applying {
			image.ApplyStatus = models.OKApply
		}

		err = models.RecommendImage(imageId, true, image.ApplyStatus, "")
	case "unrecommend":
		if !ctx.User.IsAdmin {
			ctx.Error(http.StatusForbidden)
			return
		}
		if strings.TrimSpace(form.Message) == "" {
			ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("repo.reviewCannotBeEmpty")))
			return
		}
		oldStatus := image.ApplyStatus
		if oldStatus == models.OKApply {
			image.ApplyStatus = models.NoneApply
		}
		if oldStatus == models.Applying {
			image.ApplyStatus = models.FailApply
		}

		err = models.RecommendImage(imageId, false, image.ApplyStatus, strings.TrimSpace(form.Message))
	}
	if err != nil {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("repo.star_fail", ctx.Params(":action"))))
	} else {

		notification.NotifyImageRecommend(ctx.User, image, ctx.Params(":action"))

		ctx.JSON(http.StatusOK, models.BaseOKMessage)
	}
}
