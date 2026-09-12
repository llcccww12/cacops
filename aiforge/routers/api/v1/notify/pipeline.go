package notify

import (
	"net/http"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	api "code.gitea.io/gitea/modules/structs"
)

func PipelineNotify(ctx *context.APIContext, form api.PipelineNotification) {

	ctx.JSON(http.StatusOK, models.BaseOKMessageApi)

}
