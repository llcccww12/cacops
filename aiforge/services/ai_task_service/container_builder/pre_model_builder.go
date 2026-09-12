package container_builder

import (
	"code.gitea.io/gitea/services/subject_service"
	"path"
	"strings"

	"code.gitea.io/gitea/routers/response"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/services/ai_task_service/context"
	"code.gitea.io/gitea/services/ai_task_service/storage_helper"
)

type PretrainModelBuilder struct {
	Opts *entity.ContainerBuildOpts
}

func init() {
	o := &PretrainModelBuilder{}
	RegisterContainerBuilder(o)
}

func (b *PretrainModelBuilder) SetOpts(opts *entity.ContainerBuildOpts) {
	b.Opts = opts
}

func (b *PretrainModelBuilder) GetContainerType() entity.ContainerDataType {
	return entity.ContainerPreTrainModel
}

func (b *PretrainModelBuilder) Build(ctx *context.CreationContext) ([]entity.ContainerData, *response.BizError) {
	log.Info("Start to build pretrain model.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	form := ctx.Request
	var preTrainModelEntity []entity.ContainerData
	//if ctx.Request.Cluster == entity.C2Net && (ctx.Request.JobType == models.JobTypeDebug || ctx.Request.JobType == models.JobTypeTrain) && ctx.Request.ComputeSource.Name == models.GPU {
	//	//挂载一个文件夹保证容器内pretrainmodel目录提交镜像时不被打包
	//	uploader := storage_helper.SelectStorageHelperFromStorageType(entity.OBS)
	//	objectKey := path.Join(uploader.GetJobDefaultObjectKeyPrefix(form.JobName), "pretrain_model_mount")
	//	uploader.MKDIR(objectKey, "pretrain model folder")
	//	preTrainModelEntity = append(preTrainModelEntity, entity.ContainerData{
	//		Name:          "pretrain_model_mount",
	//		Bucket:        uploader.GetBucket(),
	//		EndPoint:      uploader.GetEndpoint(),
	//		ObjectKey:     objectKey + "/",
	//		ReadOnly:      false,
	//		ContainerPath: b.Opts.ContainerPath,
	//		RealPath:      uploader.GetRealPath(objectKey),
	//		S3DownloadUrl: uploader.GetS3DownloadUrl(objectKey),
	//		IsDir:         true,
	//		IsOverwrite:   true,
	//		IsNeedUnzip:   false,
	//	})
	//}

	if b.Opts.Disable {
		log.Info("Build pretrain model b.Opts.Disable.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
		return preTrainModelEntity, nil
	}

	storageTypes := b.Opts.AcceptStorageType
	if storageTypes == nil || len(storageTypes) == 0 {
		return nil, response.SYSTEM_ERROR
	}
	//未选择预训练模型，跳过此步
	if form.PretrainModelId == "" {
		log.Info("Build pretrain model form.PretrainModelId empty displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
		return preTrainModelEntity, nil
	}
	//查出模型数据
	uuids := strings.Split(form.PretrainModelId, ";")
	modelInfoMaps, err := models.QueryModelMapsByIds(uuids)
	if err != nil || len(modelInfoMaps) == 0 {
		log.Error("Can not find model", err)
		return nil, response.MODEL_NOT_EXISTS
	}

	for _, aimodel := range modelInfoMaps {
		subject_service.InitAimodelVersion(aimodel.ID)
		data, err := b.buildAimodelData(aimodel, form.JobName)
		if err != nil {
			return nil, response.SYSTEM_ERROR
		}
		preTrainModelEntity = append(preTrainModelEntity, data)
	}
	log.Info("Build pretrain model success.displayJobName=%s jobType=%s cluster=%s", ctx.Request.DisplayJobName, ctx.Request.JobType, ctx.Request.Cluster)
	return preTrainModelEntity, nil
}

const MODEL_MKDIR_README = "The model files have already been loaded into the container and are ready for use.\n"

func (b *PretrainModelBuilder) buildAimodelData(aimodel *models.AiModelManage, jobName string) (entity.ContainerData, *response.BizError) {
	log.Info("Start build aimodel data.model=%+v displayJobName=%s", aimodel, jobName)
	uploader := storage_helper.SelectStorageHelperFromStorageType(entity.StorageType(aimodel.StorageType))
	if uploader == nil {
		return entity.ContainerData{}, response.BuildDefaultBizError("aimodel storage type not suport")
	}
	aimodelpath := uploader.TrimBucketPrefix(aimodel.Path)
	uploader.MKDIR(aimodelpath, MODEL_MKDIR_README)
	modelData := entity.ContainerData{
		Name:          aimodel.Name,
		Bucket:        uploader.GetBucket(),
		EndPoint:      uploader.GetEndpoint(),
		Id:            uploader.GetDataId(aimodel.Path),
		ObjectKey:     aimodelpath,
		ReadOnly:      true,
		ContainerPath: path.Join(b.Opts.ContainerPath, aimodel.Name),
		RealPath:      uploader.GetRealPath(aimodelpath),
		S3DownloadUrl: uploader.GetS3DownloadUrl(aimodelpath),
		IsDir:         true,
		Size:          aimodel.Size,
		IsOverwrite:   false,
		IsNeedUnzip:   false,
	}
	log.Info("Start build aimodel data success.displayJobName=%s", jobName)
	return modelData, nil
}

//func (b *PretrainModelBuilder) buildModelData(m *models.AiModelManage, jobName string) (entity.ContainerData, *response.BizError) {
//	oldStorageType := entity.GetStorageTypeFromCloudbrainType(m.Type)
//	if oldStorageType == "" {
//		log.Error("model storage type error.modelId=%d", m.ID)
//		return entity.ContainerData{}, response.SYSTEM_ERROR
//	}
//	oldStorageHelper := storage_helper.SelectStorageHelperFromStorageType(oldStorageType)
//
//	preTrainModelPath := getPreTrainModelPath(m.Path)
//	storageType := oldStorageType
//	if !b.Opts.IsStorageTypeIn(oldStorageType) {
//		//意味着模型之前存储的位置不符合要求，需要转存到指定存储
//		newStorageType := b.Opts.AcceptStorageType[0]
//		newStorageHelper := storage_helper.SelectStorageHelperFromStorageType(newStorageType)
//		files, err := oldStorageHelper.GetAllObjectsUnderDir(preTrainModelPath)
//		newObjectPrefix := path.Join(newStorageHelper.GetJobDefaultObjectKeyPrefix(jobName), b.Opts.GetLocalPath(), m.Name)
//		for _, file := range files {
//			newFilePath := path.Join(newObjectPrefix, file.FileName)
//			err = storage_helper.CopyFileBetweenStorage(oldStorageHelper, newStorageHelper, file.RelativePath, newFilePath)
//			if err != nil {
//				log.Error("transfer file between storage error.model=%+v file=%+v err=%v", m, file, err)
//				return entity.ContainerData{}, response.SYSTEM_ERROR
//			}
//		}
//		preTrainModelPath = newObjectPrefix
//		storageType = newStorageType
//	}
//
//	uploader := storage_helper.SelectStorageHelperFromStorageType(storageType)
//	uploader.MKDIR(preTrainModelPath, MODEL_MKDIR_README)
//	modelData := entity.ContainerData{
//		Name:          m.Name,
//		Bucket:        uploader.GetBucket(),
//		EndPoint:      uploader.GetEndpoint(),
//		ObjectKey:     preTrainModelPath,
//		ReadOnly:      true,
//		ContainerPath: path.Join(b.Opts.ContainerPath, m.Name),
//		RealPath:      uploader.GetRealPath(preTrainModelPath),
//		S3DownloadUrl: uploader.GetS3DownloadUrl(preTrainModelPath),
//		IsDir:         true,
//		Size:          m.Size,
//		IsOverwrite:   false,
//		IsNeedUnzip:   false,
//	}
//	log.Info("buildModelData modelData=%+v", modelData)
//	return modelData, nil
//}
//
//func getPreTrainModelPath(pretrainModelDir string) string {
//	index := strings.Index(pretrainModelDir, "/")
//	if index > 0 {
//		filterBucket := pretrainModelDir[index+1:]
//		return filterBucket
//	} else {
//		return ""
//	}
//
//}
