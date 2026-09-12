package convert

import (
	"code.gitea.io/gitea/models"
	api "code.gitea.io/gitea/modules/structs"
)

func ToCloudBrain(task *models.Cloudbrain) *api.Cloudbrain {
	return &api.Cloudbrain{
		ID:               task.ID,
		JobID:            task.JobID,
		JobType:          task.JobType,
		Type:             task.Type,
		DisplayJobName:   task.DisplayJobName,
		Status:           task.Status,
		CreatedUnix:      int64(task.CreatedUnix),
		RepoID:           task.RepoID,
		Duration:         task.Duration,
		TrainJobDuration: task.TrainJobDuration,
		ImageID:          task.ImageID,
		Image:            task.Image,
		Uuid:             task.Uuid,
		DatasetName:      task.DatasetName,
		ComputeResource:  task.ComputeResource,
		AiCenter:         task.AiCenter,
		BranchName:       task.BranchName,
		Parameters:       task.Parameters,
		BootFile:         task.BootFile,
		Description:      task.Description,
		ModelName:        task.ModelName,
		VersionName:      task.VersionName,
		ModelVersion:     task.ModelVersion,
		CkptName:         task.CkptName,
		ModelId:          task.ModelId,
		StartTime:        int64(task.StartTime),
		EndTime:          int64(task.EndTime),
		Spec:             ToSpecification(task.Spec),
	}
}
func ToAttachment(attachment *models.Attachment) *api.AttachmentShow {
	return &api.AttachmentShow{
		ID:              attachment.ID,
		UUID:            attachment.UUID,
		DatasetID:       attachment.DatasetID,
		ReleaseID:       attachment.ReleaseID,
		UploaderID:      attachment.UploaderID,
		CommentID:       attachment.CommentID,
		Name:            attachment.Name,
		Description:     attachment.Description,
		DownloadCount:   attachment.DownloadCount,
		UseNumber:       attachment.UseNumber,
		Size:            attachment.Size,
		IsPrivate:       attachment.IsPrivate,
		DecompressState: attachment.DecompressState,
		Type:            attachment.Type,
		CreatedUnix:     int64(attachment.CreatedUnix),
	}
}

func ToDataset(dataset *models.Dataset) *api.Dataset {
	var convertAttachments []*api.AttachmentShow
	for _, attachment := range dataset.Attachments {
		convertAttachments = append(convertAttachments, ToAttachment(attachment))
	}
	return &api.Dataset{
		ID:            dataset.ID,
		Title:         dataset.Title,
		Status:        dataset.Status,
		Category:      dataset.Category,
		Description:   dataset.Description,
		DownloadTimes: dataset.DownloadTimes,
		UseCount:      dataset.UseCount,
		NumStars:      dataset.NumStars,
		Recommend:     dataset.Recommend,
		License:       dataset.License,
		Task:          dataset.Task,
		ReleaseID:     dataset.ReleaseID,
		UserID:        dataset.UserID,
		RepoID:        dataset.RepoID,
		Repo: &api.RepositoryShow{
			OwnerName: dataset.Repo.OwnerName,
			Name:      dataset.Repo.Name,
		},
		CreatedUnix: int64(dataset.CreatedUnix),
		UpdatedUnix: int64(dataset.UpdatedUnix),
		Attachments: convertAttachments,
	}
}

func ToSpecification(s *models.Specification) *api.SpecificationShow {
	if s == nil {
		return nil
	}
	return &api.SpecificationShow{
		ID:                  s.ID,
		AccCardsNum:         s.AccCardsNum,
		AccCardType:         s.AccCardType,
		CpuCores:            s.CpuCores,
		MemGiB:              s.MemGiB,
		GPUMemGiB:           s.GPUMemGiB,
		ShareMemGiB:         s.ShareMemGiB,
		ComputeResource:     s.ComputeResource,
		UnitPrice:           s.UnitPrice,
		SourceSpecId:        s.SourceSpecId,
		HasInternet:         s.HasInternet,
		EnableVisualization: s.EnableVisualization,
		AICenterCode:        s.AiCenterCode,
	}
}

func ToSpecificationFromAggregate(s *models.AggregateSpecification) *api.SpecificationShow {
	if s == nil {
		return nil
	}
	return &api.SpecificationShow{
		ID:                            s.ID,
		AccCardsNum:                   s.AccCardsNum,
		AccCardType:                   s.AccCardType,
		CpuCores:                      s.CpuCores,
		MemGiB:                        s.MemGiB,
		GPUMemGiB:                     s.GPUMemGiB,
		ShareMemGiB:                   s.ShareMemGiB,
		ComputeResource:               s.ComputeResource,
		UnitPrice:                     s.UnitPrice,
		SourceSpecId:                  s.SourceSpecId,
		HasInternet:                   s.HasInternet,
		EnableVisualization:           s.EnableVisualization,
		AICenterCode:                  s.AiCenterCode,
		NetworkCapableQueuesExist:     s.Capability.NetworkCapableQueuesExist,
		NetworkIncapableQueuesExist:   s.Capability.NetworkIncapableQueuesExist,
		VisualizeCapableQueuesExist:   s.Capability.VisualizeCapableQueuesExist,
		VisualizeIncapableQueuesExist: s.Capability.VisualizeIncapableQueuesExist,
	}
}

func ToPointAccount(s *models.PointAccount) *api.PointAccountShow {
	if s == nil {
		return nil
	}
	return &api.PointAccountShow{
		ID:            s.ID,
		AccountCode:   s.AccountCode,
		Balance:       s.Balance,
		TotalEarned:   s.TotalEarned,
		TotalConsumed: s.TotalConsumed,
		Status:        s.Status,
		Version:       s.Version,
		CreatedUnix:   int64(s.CreatedUnix),
		UpdatedUnix:   int64(s.UpdatedUnix),
	}
}

func ToTagger(user *models.User) *api.Tagger {
	return &api.Tagger{
		ID:           user.ID,
		Name:         user.Name,
		RelAvatarURL: user.RelAvatarLink(),
		Email:        user.Email,
	}
}
