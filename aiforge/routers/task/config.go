package task

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/task"
	"errors"
	"net/http"
)

func GetTaskConfigList(ctx *context.Context) {
	page := ctx.QueryInt("Page")
	status := ctx.QueryInt("Status")
	action := ctx.Query("Action")
	r, err := task.GetTaskConfigWithLimitList(models.GetTaskConfigOpts{
		ListOptions: models.ListOptions{PageSize: 20, Page: page},
		Status:      status,
		TaskType:    action,
	})
	if err != nil {
		log.Error("GetTaskConfigList error.%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessWithData(r))
}

func OperateTaskConfig(ctx *context.Context, config models.TaskConfigWithLimit) {
	action := ctx.Params(":action")

	var err error
	switch action {
	case "edit":
		err = task.EditTaskConfig(config, ctx.User)
	case "new":
		err = task.AddTaskConfig(config, ctx.User)
	case "del":
		err = task.DelTaskConfig(config.ID, ctx.User)
	default:
		err = errors.New("action type error")
	}

	if err != nil {
		log.Error("OperateTaskConfig error ,%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success())
}
func BatchAddTaskConfig(ctx *context.Context, list models.BatchLimitConfigVO) {
	successCount := 0
	failCount := 0
	for _, config := range list.ConfigList {
		err := task.AddTaskConfig(config, ctx.User)
		if err != nil {
			failCount++
		} else {
			successCount++
		}
	}
	r := make(map[string]int, 2)
	r["successCount"] = successCount
	r["failCount"] = failCount
	log.Debug("BatchAddTaskConfig success.result=%v", r)
	ctx.JSON(http.StatusOK, response.SuccessWithData(r))
}
