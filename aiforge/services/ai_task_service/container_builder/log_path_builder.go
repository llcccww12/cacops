package container_builder

import (
	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/context"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
	"path"
)

type LogPathBuilder struct {
	Opts *entity.ContainerBuildOpts
}

func init() {
	o := &LogPathBuilder{}
	RegisterContainerBuilder(o)
}

func (b *LogPathBuilder) SetOpts(opts *entity.ContainerBuildOpts) {
	b.Opts = opts
}

func (b *LogPathBuilder) Build(ctx *context.CreationContext) ([]entity.ContainerData, *response.BizError) {
	log.Info("start to build log path.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
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
	log.Info("Build log path success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return []entity.ContainerData{{
		ContainerPath: b.Opts.ContainerPath,
		ReadOnly:      b.Opts.ReadOnly,
		ObjectKey:     remoteDir,
		RealPath:      uploader.GetRealPath(remoteDir),
		Bucket:        uploader.GetBucket(),
		EndPoint:      uploader.GetEndpoint(),
		IsDir:         true,
		StorageType:   storageTypes[0],
	}}, nil
}

func (b *LogPathBuilder) GetContainerType() entity.ContainerDataType {
	return entity.ContainerLogPath
}
