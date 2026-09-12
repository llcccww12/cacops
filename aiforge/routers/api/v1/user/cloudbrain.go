package user

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"net/http"

	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/convert"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/services/reward/point/account"
)

func GetPointAccount(ctx *context.APIContext) {
	a, err := account.GetAccount(ctx.User.ID)
	if err != nil {
		ctx.ServerError("GetPointAccount", err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"pointAccount":        convert.ToPointAccount(a),
		"cloudBrainPaySwitch": setting.CloudBrainPaySwitch,
	})
}

func GetCloudbrainRepo(ctx *context.APIContext) {
	cloudbrains, err := models.GetNewestCloudbrainByUser(ctx.User.ID)
	if err != nil || len(cloudbrains) == 0 {
		log.Warn("GetNewestCloudbrainByUser", err)
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"job_type": "",
			"repo":     "",
		})
		return
	}

	cloudbrain := cloudbrains[0]
	repo, err := models.GetRepositoryByID(cloudbrain.RepoID)
	if err != nil || repo == nil {
		log.Warn("GetRepositoryByID", err)
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"job_type": cloudbrain.JobType,
			"repo":     "",
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"job_type": cloudbrain.JobType,
		"repo":     repo.FullName(),
	})

}
