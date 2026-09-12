package container_builder

import (
	"path"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/context"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
	"code.gitea.io/gitea/services/role"
)

type OutputPathBuilder struct {
	Opts *entity.ContainerBuildOpts
}

func init() {
	o := &OutputPathBuilder{}
	RegisterContainerBuilder(o)
}

func (b *OutputPathBuilder) SetOpts(opts *entity.ContainerBuildOpts) {
	b.Opts = opts
}

func (b *OutputPathBuilder) Build(ctx *context.CreationContext) ([]entity.ContainerData, *response.BizError) {
	log.Info("Start to build output path.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	if b.Opts.Disable {
		return nil, nil
	}
	storageTypes := b.Opts.AcceptStorageType
	if storageTypes == nil || len(storageTypes) == 0 {
		return nil, response.SYSTEM_ERROR
	}

	jobName := ctx.Request.JobName

	uploader := storage_helper.SelectStorageHelperFromStorageType(storageTypes[0])
	remoteDir := path.Join(uploader.GetJobDefaultObjectKeyPrefix(jobName), b.Opts.GetLocalPath())
	if b.Opts.MKDIR {
		err := uploader.MKDIR(remoteDir)
		if err != nil {
			log.Error("MKDIR err.displayJobName = %s err=%v", ctx.Request.DisplayJobName, err)
			return nil, response.NewBizError(err)
		}
	}

	//如果是继续训练，需要将上次的结果拷贝到本次训练任务的输出目录
	if ctx.Request.IsContinueRequest {
		if ctx.SourceCloudbrain == nil {
			log.Error("SourceCloudbrain empty.displayJobName = %s", ctx.Request.DisplayJobName)
			return nil, response.PARAM_ERROR
		}
		sourcePath := getSourceOutputPath(ctx.SourceCloudbrain, uploader, b.Opts.ContainerPath)
		err := uploader.CopyDir(sourcePath, remoteDir, []string{"README", ".txt"})
		if err != nil {
			log.Error("CopyByPath err.displayJobName = %s err=%v", ctx.Request.DisplayJobName, err)
			return nil, response.NewBizError(err)
		}
	}
	log.Info("Build output path success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	outputData := entity.ContainerData{
		ContainerPath:     b.Opts.ContainerPath,
		ReadOnly:          b.Opts.ReadOnly,
		ObjectKey:         remoteDir,
		RealPath:          uploader.GetRealPath(remoteDir),
		Bucket:            uploader.GetBucket(),
		EndPoint:          uploader.GetEndpoint(),
		GetBackEndpoint:   uploader.GetEndpoint(),
		IsDir:             true,
		StorageType:       storageTypes[0],
		S3DownloadUrl:     uploader.GetS3DownloadUrl(remoteDir),
		IsNeedTensorboard: ctx.Request.VisualizeRequired,
	}
	if ctx.User != nil {
		_, outputSize := role.GetUserContainerStorageLimits(ctx.User.ID)
		outputData.SizeLimit = outputSize
	}
	return []entity.ContainerData{outputData}, nil
}

func getSourceOutputPath(sourceCloudbrain *models.Cloudbrain, helper storage_helper.StorageHelper, containerPath string) string {
	c := sourceCloudbrain.GetCloudbrainConfig()
	if c != nil {
		return c.OutputObjectPrefix
	}
	return path.Join(helper.GetJobDefaultObjectKeyPrefix(sourceCloudbrain.JobName), sourceCloudbrain.VersionName, containerPath)
}

func (b *OutputPathBuilder) GetContainerType() entity.ContainerDataType {
	return entity.ContainerOutPutPath
}
