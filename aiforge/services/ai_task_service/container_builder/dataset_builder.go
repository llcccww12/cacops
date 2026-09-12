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
	"code.gitea.io/gitea/services/subject_service"
)

type DatasetBuilder struct {
	Opts *entity.ContainerBuildOpts
}

func init() {
	o := &DatasetBuilder{}
	RegisterContainerBuilder(o)
}

func (b *DatasetBuilder) SetOpts(opts *entity.ContainerBuildOpts) {
	b.Opts = opts
}

func (b *DatasetBuilder) GetContainerType() entity.ContainerDataType {
	return entity.ContainerDataset
}

func (b *DatasetBuilder) Build(ctx *context.CreationContext) ([]entity.ContainerData, *response.BizError) {
	log.Info("Start to build dataset.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	form := ctx.Request
	var datasetEntityList []entity.ContainerData

	if b.Opts.Disable {
		return datasetEntityList, nil
	}

	storageTypes := b.Opts.AcceptStorageType
	if storageTypes == nil || len(storageTypes) == 0 {
		return nil, response.SYSTEM_ERROR
	}
	//未选择数据集，跳过此步
	if form.DatasetUUIDStr == "" {
		return datasetEntityList, nil
	}
	uuids := strings.Split(form.DatasetUUIDStr, ";")
	datasetMaps, err := models.QueryDatasetRegistryMapsByIds(uuids)
	if err != nil || len(datasetMaps) == 0 {
		log.Error("Can not find dataset.%v", err)
		return nil, response.DATASET_NOT_EXISTS
	}

	for _, dataset := range datasetMaps {
		subject_service.InitDatasetVersion(dataset.ID)
		data, err := b.buildDatasetData(dataset, form.JobName)
		if err != nil {
			return nil, response.SYSTEM_ERROR
		}
		datasetEntityList = append(datasetEntityList, data)
	}
	log.Info("Build pretrain model success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return datasetEntityList, nil
}

const DATASET_MKDIR_README = "The dataset files have already been loaded into the container and are ready for use.\n"

func (b *DatasetBuilder) buildDatasetData(dataset *models.DatasetRegistry, jobName string) (entity.ContainerData, *response.BizError) {

	uploader := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(dataset.StorageType))
	if uploader == nil {
		return entity.ContainerData{}, response.BuildDefaultBizError("dataset storage type not suport")
	}
	datasetpath := uploader.TrimBucketPrefix(dataset.Path)
	uploader.MKDIR(datasetpath, DATASET_MKDIR_README)
	modelData := entity.ContainerData{
		Name:          dataset.Name,
		Bucket:        uploader.GetBucket(),
		EndPoint:      uploader.GetEndpoint(),
		Id:            uploader.GetDataId(dataset.Path),
		ObjectKey:     datasetpath,
		ReadOnly:      true,
		ContainerPath: path.Join(b.Opts.ContainerPath, dataset.Name),
		RealPath:      uploader.GetRealPath(datasetpath),
		S3DownloadUrl: uploader.GetS3DownloadUrl(datasetpath),
		IsDir:         true,
		Size:          dataset.Size,
		IsOverwrite:   false,
		IsNeedUnzip:   false,
	}
	return modelData, nil
}
