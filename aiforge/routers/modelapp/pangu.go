package modelapp

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/services/cloudbrain/cloudbrainTask"
)

var tplModelBase base.TplName = "model/base/index"
var tplPanguFinetuneHome base.TplName = "model/base/panguIndex"
var tplPanguFinetuneCreate base.TplName = "model/base/panguCreate"
var tplPanguInference base.TplName = "model/base/panguInference"

func ModelBaseUI(ctx *context.Context) {
	ctx.HTML(200, tplModelBase)
}

func PanguFinetuneUI(ctx *context.Context) {
	ctx.HTML(200, tplPanguFinetuneHome)
}

func PanguFinetuneCreateUI(ctx *context.Context) {
	NotStopTaskCount, _ := cloudbrainTask.GetNotFinalStatusTaskCount(ctx.User.ID, string(models.JobTypeTrain))
	ctx.Data["NotStopTaskCount"] = NotStopTaskCount
	ctx.HTML(200, tplPanguFinetuneCreate)
}

func PanguInferenceUI(ctx *context.Context) {
	ctx.HTML(200, tplPanguInference)
}
