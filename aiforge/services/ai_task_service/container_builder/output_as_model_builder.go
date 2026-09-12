package container_builder

import (
	"path"
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/context"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
)

type OutputAsPretrainModelBuilder struct {
	Opts *entity.ContainerBuildOpts
}

func init() {
	o := &OutputAsPretrainModelBuilder{}
	RegisterContainerBuilder(o)
}

func (b *OutputAsPretrainModelBuilder) SetOpts(opts *entity.ContainerBuildOpts) {
	b.Opts = opts
}

func (b *OutputAsPretrainModelBuilder) Build(ctx *context.CreationContext) ([]entity.ContainerData, *response.BizError) {
	log.Info("Start to build output path.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	if b.Opts.Disable {
		return nil, nil
	}
	storageTypes := b.Opts.AcceptStorageType
	if storageTypes == nil || len(storageTypes) == 0 {
		return nil, response.SYSTEM_ERROR
	}

	uploader := storage_helper.SelectStorageHelperFromStorageType(storageTypes[0])

	subPath := models.GetParam(ctx.Request.ParamArray.Parameter, "sub_path")
	if ctx.SourceCloudbrain == nil {
		log.Info("SourceCloudbrain empty.displayJobName = %s", ctx.Request.DisplayJobName)
		return nil, nil
	}
	sourcePath := getSourceOutputPath(ctx.SourceCloudbrain, uploader, b.Opts.ContainerPath)
	if subPath != "" {
		sourcePath = path.Join(sourcePath, subPath)
	}
	if !strings.HasSuffix(sourcePath, "/") {
		sourcePath = sourcePath + "/"
	}

	uploader.MKDIR(sourcePath, MODEL_MKDIR_README)

	log.Info("Build output path success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return []entity.ContainerData{{
		Name:          "finetuned_model",
		Bucket:        uploader.GetBucket(),
		EndPoint:      uploader.GetEndpoint(),
		ObjectKey:     sourcePath,
		ReadOnly:      true,
		ContainerPath: path.Join(b.Opts.ContainerPath, "finetuned_model"),
		RealPath:      uploader.GetRealPath(sourcePath),
		S3DownloadUrl: uploader.GetS3DownloadUrl(sourcePath),
		IsDir:         true,
		IsOverwrite:   true,
		IsNeedUnzip:   false,
	}}, nil

}

func (b *OutputAsPretrainModelBuilder) GetContainerType() entity.ContainerDataType {
	return entity.ContainerOutputAsModel
}
