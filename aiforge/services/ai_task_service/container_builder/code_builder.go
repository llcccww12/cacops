package container_builder

import (
	"path"
	"strings"

	"code.gitea.io/gitea/models"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/modules/cloudbrain"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/context"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
	"code.gitea.io/gitea/services/role"
)

type CodeBuilder struct {
	Opts *entity.ContainerBuildOpts
}

func init() {
	o := &CodeBuilder{}
	RegisterContainerBuilder(o)
}

func (b *CodeBuilder) SetOpts(opts *entity.ContainerBuildOpts) {
	b.Opts = opts
}

func (b *CodeBuilder) GetContainerType() entity.ContainerDataType {
	return entity.ContainerCode
}

func (b *CodeBuilder) Build(ctx *context.CreationContext) ([]entity.ContainerData, *response.BizError) {
	log.Info("Start to build code.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	opts := b.Opts
	if opts.Disable {
		return nil, nil
	}
	var storageTypes []entity.StorageType
	//切换到海胆2.0
	if setting.FORCE_OLD_CODE_STORAGE {
		storageTypes = opts.AcceptStorageType
		if storageTypes == nil || len(storageTypes) == 0 {
			return nil, response.SYSTEM_ERROR
		}
	} else {
		storageTypes = []entity.StorageType{entity.URCHIN_V2}
	}

	jobName := ctx.Request.JobName
	uploader := storage_helper.SelectStorageHelperFromStorageType(storageTypes[0])
	remoteDir := uploader.GetJobDefaultObjectKeyPrefix(jobName) + opts.GetLocalPath()
	// 在线运行notebook或调试任务，Repository为nil，创建只有SizeLimit的ContainerData
	if ctx.Repository == nil {
		if ctx.Request.IsFileNoteBookRequest && ctx.Request.Cluster == entity.C2Net && (ctx.Request.ComputeSource.Name == models.GPU || ctx.Request.ComputeSource.Name == models.NPU) {
			err := uploader.MKDIR(remoteDir, "read me")
			if err != nil {
				log.Error("MKDIR err.displayJobName = %s err=%v", ctx.Request.DisplayJobName, err)
				return nil, response.LOAD_CODE_FAILED
			}
		}

		codeData := entity.ContainerData{}
		if ctx.User != nil {
			codeSize, _ := role.GetUserContainerStorageLimits(ctx.User.ID)
			codeData.SizeLimit = codeSize
		}
		log.Info("Build code success (no repository).displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
		return []entity.ContainerData{codeData}, nil
	}
	repo := ctx.Repository
	codeLocalPath := setting.JobPath + jobName + cloudbrain.CodeMountPath + "/"
	//再次调试和在线运行notebook不需要下载、上传代码
	if !ctx.Request.IsRestartRequest && !ctx.Request.IsFileNoteBookRequest {
		log.Info("try to download code.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
		if err := DownloadCode(ctx, codeLocalPath, b.Opts.Uncompressed); err != nil {
			log.Error("downloadZipCode failed, server timed out: %s (%v)", repo.FullName(), err)
			return nil, response.LOAD_CODE_FAILED.WithParams(err.Error())
		}
		log.Info("Download code success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
		if opts.VolumeFolder && storageTypes[0] == entity.OBS {
			uploader.MKDIR(remoteDir)
		}
		log.Info("try to upload code dir.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
		path, err := uploader.AllocateDatasetNamespace(ctx.Repository.Name, remoteDir)
		if err != nil {
			log.Error("downloadZipCode AllocateDatasetNamespace,err= %s (%v)", repo.FullName(), err)
			return nil, response.LOAD_CODE_FAILED.WithParams(err.Error())
		}
		if err := uploader.UploadDir(codeLocalPath, path); err != nil {
			log.Error("Failed to UploadDir: %s (%v)", repo.FullName(), err)
			return nil, response.LOAD_CODE_FAILED.WithParams(err.Error())
		}
		if err := uploader.ReportSizeChanged(path); err != nil {
			log.Error("Report code SizeChanged failed, path=%s err=%v", path, err)
		}
		remoteDir = path
		log.Info("Upload code dir success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	}
	if ctx.Request.IsFileNoteBookRequest && ctx.Request.Cluster == entity.C2Net && (ctx.Request.ComputeSource.Name == models.GPU || ctx.Request.ComputeSource.Name == models.NPU) {
		err := uploader.MKDIR(remoteDir, "read me")

		if err != nil {
			log.Error("MKDIR err.displayJobName = %s err=%v", ctx.Request.DisplayJobName, err)
			return nil, response.LOAD_CODE_FAILED.WithParams(err.Error())
		}
	}
	dataID := uploader.GetDataId(remoteDir)

	var codeArchiveName, objectKey string
	if !b.Opts.Uncompressed && !b.Opts.VolumeFolder {
		if dataID == "" {
			//此时表示用的不是海胆2.0
			// 特殊处理：如果不是海胆2.0且dataID为空，说明是压缩包方式，需要手动指定containerPath为压缩包文件路径
			// 目前只能这样处理，后续如果支持文件传输可以优化
			codeArchiveName = cloudbrain.DefaultBranchName + ".zip"
			objectKey = path.Join(remoteDir, codeArchiveName)
		}

	} else {
		objectKey = remoteDir + "/"
	}

	containerPath := ""
	if opts.ContainerPath != "" {
		//如果代码是压缩包，此时的挂载路径是文件
		//如果代码不是压缩包，此时的挂载路径是目录
		containerPath = path.Join(opts.ContainerPath, codeArchiveName)

	}

	codeData := entity.ContainerData{
		Id:            dataID,
		Name:          strings.ToLower(repo.Name),
		Bucket:        uploader.GetBucket(),
		EndPoint:      uploader.GetEndpoint(),
		ObjectKey:     objectKey,
		ReadOnly:      opts.ReadOnly,
		ContainerPath: containerPath,
		RealPath:      uploader.GetRealPath(objectKey),
		IsDir:         b.Opts.Uncompressed,
		S3DownloadUrl: uploader.GetS3DownloadUrl(objectKey),
		IsNeedUnzip:   true,
		StorageType:   storageTypes[0],
	}
	if ctx.User != nil {
		codeSize, _ := role.GetUserContainerStorageLimits(ctx.User.ID)
		codeData.SizeLimit = codeSize
	}
	log.Info("Build code success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return []entity.ContainerData{codeData}, nil
}
