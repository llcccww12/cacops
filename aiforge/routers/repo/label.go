package repo

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
)

const (
	tplLabelIndex base.TplName = "repo/datasets/label/index"
)

func LabelIndex(ctx *context.Context) {

	log.Info("Go Here LabelIndex.")
	uuid := ctx.Params("uuid")
	attach, err := models.GetAttachmentByUUID(uuid)
	if err != nil {
		log.Info("query attach error")
	} else {
		dataset, err := models.GetDatasetByID(attach.DatasetID)
		if err != nil {
			log.Info("query dataset error")
		} else {
			ctx.Data["repoId"] = dataset.RepoID
		}
	}

	attachments := QueryDataSet(ctx)

	ctx.Data["uuid"] = uuid
	ctx.Data["Attachments"] = attachments

	ctx.HTML(200, tplLabelIndex)

}
