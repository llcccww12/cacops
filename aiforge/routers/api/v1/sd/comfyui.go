package sd

import (
	"code.gitea.io/gitea/modules/context"
)

func GetComfyuiExperienceUrl(ctx *context.APIContext) {
	cloudbrainId := ctx.QueryInt64("task_id")
	tokenUrl := tokenUrlByTaskId(cloudbrainId) + "/"

	ctx.JSON(200, map[string]interface{}{
		"url": tokenUrl,
	})
}
