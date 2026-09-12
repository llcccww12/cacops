package monitor

import (
	"net/http"
	"strconv"
	"strings"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/convert"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/repo"
	"code.gitea.io/gitea/services/cloudbrain/resource"
)

func GetTaskMonitorSummary(ctx *context.APIContext) {

	summary, err := GetCurrentMonitorSummary()
	if err != nil {
		log.Warn("resty GetCurrentMonitorSummary: %v", err)
		ctx.JSON(http.StatusNotFound, nil)

	} else {
		ctx.JSON(http.StatusOK, summary)
	}

}

func GetAlerts(ctx *context.APIContext) {

	listOption := models.ListOptions{
		Page:     ctx.QueryInt("page"),
		PageSize: ctx.QueryInt("page_size"),
	}

	opt := SearchAlertsOption{AlertType: ctx.Query("alert_type"), JobType: ctx.Query("job_type"), Keyword: ctx.Query("q"), ListOptions: listOption}
	alerts, err := GetTaskAlerts(opt)

	if err != nil {
		log.Warn("resty GetCurrentMonitorSummary: %v", err)
		ctx.JSON(http.StatusNotFound, nil)

	} else {
		loadAttributes(ctx, alerts)
		ctx.JSON(http.StatusOK, alerts)
	}

}

func loadAttributes(ctx *context.APIContext, alerts *AlertsResponse) {
	if alerts.Code == 0 {
		for _, alert := range alerts.Data.Alerts {
			taskID, err := strconv.ParseInt(alert.TaskID, 10, 64)
			if err != nil {
				log.Warn("task id is not integer: %v", err)
				continue
			}
			cloudbrain, err := models.GetCloudbrainByCloudbrainID(taskID)
			if err != nil {
				log.Warn("get cloudbrain err: %v", err)
				continue
			}
			spec, err := resource.GetCloudbrainSpec(taskID)
			if err != nil {
				log.Warn("get spec err: %v", err)
				continue
			}
			if spec != nil {
				alert.Spec = convert.ToSpecification(spec)
				alert.CardType = alert.Spec.AccCardType
			}

			alert.JobName = cloudbrain.JobName
			if alert.UserName == "" {
				user, _ := models.GetUserByID(cloudbrain.UserID)
				if user != nil {
					alert.UserName = user.Name
				}
			}
			alert.AICenter = repo.GetAiCenterNameByCode(strings.Split(cloudbrain.AiCenter, "+")[0], ctx.Language())
			alert.Status = cloudbrain.Status

		}
	}

}
