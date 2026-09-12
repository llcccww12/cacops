package storage_limit

import (
	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/services/role"
)

/*
*GetUserStorageLimit 返回用户的存储限制。
该函数用于获取系统或服务为用户设置的存储空间上限。

	返回值为一个整数，表示存储空间的限制,单位为G。
	如果返回值为0，可能表示用户没有存储限制。
*/
func SearchDatasetAttachment(opts *models.SearchDatasetAttachmentOptions) ([]*entity.DatasetAttachmentShow, int64, error) {

	attachmentList, total, err := models.SearchDatasetAttachments(opts)

	return convertAttachmentListToDatasetAttachmentShowList(attachmentList), total, err

}

func convertAttachmentListToDatasetAttachmentShowList(attachmentList []*models.Attachment) []*entity.DatasetAttachmentShow {
	if attachmentList == nil {
		return nil
	}

	var datasetAttachmentShowList []*entity.DatasetAttachmentShow

	for _, attachment := range attachmentList {

		show := &entity.DatasetAttachmentShow{
			ID:              attachment.ID,
			Name:            attachment.Name,
			Size:            attachment.Size,
			UUID:            attachment.UUID,
			DecompressState: attachment.DecompressState,
			CreatedUnix:     int64(attachment.CreatedUnix),
			DatasetID:       attachment.DatasetID,
			IsPrivate:       attachment.IsPrivate,
		}
		if attachment.Dataset != nil {
			show.DatasetName = attachment.Dataset.Title
			if attachment.Dataset.Repo != nil {
				show.RepoName = attachment.Dataset.Repo.Name
				show.RepoOwner = attachment.Dataset.Repo.OwnerName
				show.RepoID = attachment.Dataset.Repo.ID
			}
		}

		datasetAttachmentShowList = append(datasetAttachmentShowList, show)

	}
	return datasetAttachmentShowList
}

func IsFileUploadOverLimit(userId int64, fileSzie int64) bool {

	storageLimit := role.UserStorageNum(userId)

	if storageLimit == -1 {
		return false
	}

	usedDatasetSize, err := models.GetUsedDatasetRegistrySizeByUser(userId)
	if err != nil {
		log.Warn("GetUsedDatasetSizeByUser", err)
		return true

	}
	usedModelSize, err := models.GetUsedModelSizeByOwner(userId)
	if err != nil {
		log.Warn("GetUsedModelSizeByOwner", err)

		return true

	}

	return storageLimit*1024*1024*1024 < usedDatasetSize+usedModelSize+fileSzie

}

func IsSubjectFileUploadOverLimit(fileSzie int64, subject *models.SubjectAccessContext) bool {
	err := subject.GetOwner()
	if err != nil {
		log.Error("IsSubjectFileUploadOverLimit GetOwner error: %v", err)
		return true
	}
	return IsFileUploadOverLimit(subject.OwnerID, fileSzie)
}

func GetUserStorageSummary(user *models.User) (*entity.StorageSummary, error) {
	storageLimit := role.UserStorageNum(user.ID)

	usedDatasetSize, err := models.GetUsedDatasetRegistrySizeByUser(user.ID)
	if err != nil {
		return nil, err
	}
	usedModelSize, err := models.GetUsedModelSizeByOwner(user.ID)
	if err != nil {
		return nil, err

	}
	return &entity.StorageSummary{
		StorageLimit:       storageLimit * 1024 * 1024 * 1024,
		DatasetUsedStorage: usedDatasetSize,
		ModelUsedStorage:   usedModelSize,
		UsedStorage:        usedDatasetSize + usedModelSize,
		RemainingStorage:   storageLimit*1024*1024*1024 - usedDatasetSize - usedModelSize,
	}, nil
}

func GetOrgStorageSummary(user *models.User) (*entity.StorageSummary, error) {
	storageLimit := role.UserStorageNum(user.ID)

	usedDatasetSize, err := models.GetUsedDatasetRegistrySizeByUser(user.ID)
	if err != nil {
		return nil, err
	}
	usedModelSize, err := models.GetUsedModelSizeByOwner(user.ID)
	if err != nil {
		return nil, err

	}
	return &entity.StorageSummary{
		StorageLimit:       storageLimit * 1024 * 1024 * 1024,
		DatasetUsedStorage: usedDatasetSize,
		ModelUsedStorage:   usedModelSize,
		UsedStorage:        usedDatasetSize + usedModelSize,
		RemainingStorage:   storageLimit*1024*1024*1024 - usedDatasetSize - usedModelSize,
	}, nil
}
