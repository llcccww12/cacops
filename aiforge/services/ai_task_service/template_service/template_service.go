package template_service

import (
	"regexp"
	"strings"

	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"

	gouuid "github.com/satori/go.uuid"
)

var AlphaDashDotChinese = regexp.MustCompile(`^[\p{Han}A-Za-z0-9\-_.]{1,100}$`)

func CreateAITaskTemplate(req entity.CreateAITaskTemplateReq) (string, *response.BizError) {
	if req.JobType == "" || req.Cluster == "" || req.ComputeSource == "" || req.Name == "" {
		return "", response.PARAM_ERROR
	}

	if !AlphaDashDotChinese.MatchString(req.Name) {
		return "", response.AI_TAKS_TEMPLATE_NAME_INVALID
	}

	dataId := gouuid.NewV4().String()

	//补全信息
	var repo *models.Repository
	if req.RepoName != "" && req.RepoOwnerName != "" {
		r, err := models.GetRepositoryByOwnerAndName(req.RepoOwnerName, req.RepoName)
		if err != nil {
			if models.IsErrRepoNotExist(err) {
				return "", response.REPOSITORY_NOT_EXISTS
			}
			return "", response.NewBizError(err)
		}
		// if !req.IsPrivate && r.IsPrivate {
		// 	return "", response.PRIVATE_REPOSITORY_IN_PUBLIC_TEMPLATE
		// }
		repo = r
	}

	var modelList []*models.TemplateModelInfo
	if len(req.ModelIDs) > 0 {
		r, err := models.QueryModelByIds(req.ModelIDs)
		if err != nil {
			return "", response.NewBizError(err)
		}
		if len(r) != len(req.ModelIDs) {
			return "", response.MODEL_NOT_EXISTS
		}
		for _, m := range r {
			// //设置为公开模板时需要数据集也为公开
			// if !req.IsPrivate && m.IsPrivate {
			// 	return "", response.PRIVATE_MODEL_IN_PUBLIC_TEMPLATE
			// }

			var ownerName string
			if m.Owner != nil {
				ownerName = m.Owner.Name
			}
			modelList = append(modelList, &models.TemplateModelInfo{
				TemplateID: dataId,
				ModelID:    m.ID,
				ModelName:  m.Name,
				ModelAlias: m.Alias,
				OwnerName:  ownerName,
			})
		}
	}

	var datasetList []*models.TemplateDatasetInfo
	if len(req.DatasetIDs) > 0 {
		r, err := models.GetDatasetRegistryListByIDs(req.DatasetIDs)
		if err != nil {
			return "", response.NewBizError(err)
		}
		if len(r) != len(req.DatasetIDs) {
			return "", response.DATASET_NOT_EXISTS
		}
		for _, d := range r {
			// if !req.IsPrivate && d.IsPrivate {
			// 	return "", response.PRIVATE_DATASET_IN_PUBLIC_TEMPLATE
			// }
			var ownerName string
			if d.Owner != nil {
				ownerName = d.Owner.Name
			}
			datasetList = append(datasetList, &models.TemplateDatasetInfo{
				TemplateID:   dataId,
				DatasetID:    d.ID,
				DatasetName:  d.Name,
				DatasetAlias: d.Alias,
				OwnerName:    ownerName,
			})
		}
	}

	templateDTO := &models.AITaskTemplate{
		ID:                dataId,
		Name:              req.Name,
		LowerName:         strings.ToLower(req.Name),
		Description:       req.Description,
		Tags:              req.Tags,
		IsPrivate:         req.IsPrivate,
		OwnerID:           req.OwnerID,
		JobType:           req.JobType,
		Cluster:           req.Cluster,
		ComputeSource:     req.ComputeSource,
		HasInternet:       req.HasInternet,
		VisualizeRequired: req.VisualizeRequired,
		Parameters:        req.Parameters,
		AccCardsNum:       req.AccCardsNum,
		AccCardType:       req.AccCardType,
		CpuCores:          req.CpuCores,
		MemGiB:            req.MemGiB,
		GPUMemGiB:         req.GPUMemGiB,
		ShareMemGiB:       req.ShareMemGiB,
		ImageID:           req.ImageID,
		ImageName:         req.ImageName,
		ImageUrl:          req.ImageUrl,
		DatasetList:       datasetList,
		ModelList:         modelList,
	}

	if repo != nil {
		templateDTO.RepoId = repo.ID
		templateDTO.RepoName = repo.Name
		templateDTO.RepoOwnerName = req.RepoOwnerName
		templateDTO.BranchName = req.BranchName
		templateDTO.BootFile = req.BootFile
	}

	err := models.CreateAITaskTemplate(templateDTO)
	if err != nil {
		log.Error("CreateDatasetRegistry failed, name=%s, err=%v", req.Name, err)
		return "", response.NewBizError(err)
	}

	return templateDTO.ID, nil
}

func EditAITaskTemplate(req entity.CreateAITaskTemplateReq, id string) *response.BizError {
	if req.JobType == "" || req.Cluster == "" || req.ComputeSource == "" || req.Name == "" {
		return response.PARAM_ERROR
	}

	if !AlphaDashDotChinese.MatchString(req.Name) {
		return response.AI_TAKS_TEMPLATE_NAME_INVALID
	}

	//补全信息
	var repo *models.Repository
	if req.RepoName != "" && req.RepoOwnerName != "" {
		r, err := models.GetRepositoryByOwnerAndName(req.RepoOwnerName, req.RepoName)
		if err != nil {
			if models.IsErrRepoNotExist(err) {
				return response.REPOSITORY_NOT_EXISTS
			}
			return response.NewBizError(err)
		}
		// if !req.IsPrivate && r.IsPrivate {
		// 	return response.PRIVATE_REPOSITORY_IN_PUBLIC_TEMPLATE
		// }
		repo = r
	}

	var modelList []*models.TemplateModelInfo
	if len(req.ModelIDs) > 0 {
		r, err := models.QueryModelByIds(req.ModelIDs)
		if err != nil {
			return response.NewBizError(err)
		}
		if len(r) != len(req.ModelIDs) {
			return response.MODEL_NOT_EXISTS
		}
		for _, m := range r {
			// //设置为公开模板时需要模型也为公开
			// if !req.IsPrivate && m.IsPrivate {
			// 	return response.PRIVATE_MODEL_IN_PUBLIC_TEMPLATE
			// }
			var ownerName string
			if m.Owner != nil {
				ownerName = m.Owner.Name
			}
			modelList = append(modelList, &models.TemplateModelInfo{
				TemplateID: id,
				ModelID:    m.ID,
				ModelName:  m.Name,
				ModelAlias: m.Alias,
				OwnerName:  ownerName,
			})
		}
	}

	var datasetList []*models.TemplateDatasetInfo
	if len(req.DatasetIDs) > 0 {
		r, err := models.GetDatasetRegistryListByIDs(req.DatasetIDs)
		if err != nil {
			return response.NewBizError(err)
		}
		if len(r) != len(req.DatasetIDs) {
			return response.DATASET_NOT_EXISTS
		}
		for _, d := range r {
			// //设置为公开模板时需要数据集也为公开
			// if !req.IsPrivate && d.IsPrivate {
			// 	return response.PRIVATE_DATASET_IN_PUBLIC_TEMPLATE
			// }
			var ownerName string
			if d.Owner != nil {
				ownerName = d.Owner.Name
			}
			datasetList = append(datasetList, &models.TemplateDatasetInfo{
				TemplateID:   id,
				DatasetID:    d.ID,
				DatasetName:  d.Name,
				DatasetAlias: d.Alias,
				OwnerName:    ownerName,
			})
		}
	}

	templateDTO := &models.AITaskTemplate{
		ID:                id,
		Name:              req.Name,
		LowerName:         strings.ToLower(req.Name),
		Description:       req.Description,
		Tags:              req.Tags,
		IsPrivate:         req.IsPrivate,
		OwnerID:           req.OwnerID,
		JobType:           req.JobType,
		Cluster:           req.Cluster,
		ComputeSource:     req.ComputeSource,
		HasInternet:       req.HasInternet,
		VisualizeRequired: req.VisualizeRequired,
		Parameters:        req.Parameters,
		AccCardsNum:       req.AccCardsNum,
		AccCardType:       req.AccCardType,
		CpuCores:          req.CpuCores,
		MemGiB:            req.MemGiB,
		GPUMemGiB:         req.GPUMemGiB,
		ShareMemGiB:       req.ShareMemGiB,
		ImageID:           req.ImageID,
		ImageName:         req.ImageName,
		ImageUrl:          req.ImageUrl,
		DatasetList:       datasetList,
		ModelList:         modelList,
	}

	if repo != nil {
		templateDTO.RepoId = repo.ID
		templateDTO.RepoName = repo.Name
		templateDTO.RepoOwnerName = req.RepoOwnerName
		templateDTO.BranchName = req.BranchName
		templateDTO.BootFile = req.BootFile
	}

	err := models.UpdateAITaskTemplate(templateDTO)
	if err != nil {
		log.Error("UpdateAITaskTemplate failed, name=%s, err=%v", req.Name, err)
		return response.NewBizError(err)
	}

	return nil
}

func SearchAITaskTemplates(req models.SearchAITaskTemplateReq) ([]*entity.AITaskTemplateDetailInfo, int64, *response.BizError) {
	templates, total, err := models.SearchAITaskTempalte(req)
	if err != nil {
		log.Error("SearchAITaskTemplate failed, req=%+v, err=%v", req, err)
		return nil, 0, response.NewBizError(err)
	}
	if templates == nil || len(templates) == 0 {
		log.Info("SearchAITaskTemplate no data found, req=%+v", req)
		return nil, 0, nil
	}
	results := make([]*entity.AITaskTemplateDetailInfo, 0, len(templates))
	for i := 0; i < len(templates); i++ {
		templateDTO := templates[i]
		results = append(results, entity.BuildAITaskTemplateDetaiInfo(templateDTO, req.Operator))
	}

	return results, total, nil

}

func DelAITaskTemplate(id string) error {
	return models.DelAITaskTemplateByID(id)
}

func CollectAITaskTemplate(userId int64, templateID string) error {
	return models.CollectionAITaskTemplate(userId, templateID, true)
}

func UnCollectAITaskTemplate(userId int64, templateID string) error {
	return models.CollectionAITaskTemplate(userId, templateID, false)
}

func RecommendAITaskTemplate(templateID string) error {
	var err error
	err = models.RecommendAITaskTemplate(templateID, true)
	if err != nil {
		return err
	}

	return nil
}

func UnRecommendAITaskTemplate(templateID string) error {
	var err error
	err = models.RecommendAITaskTemplate(templateID, false)
	if err != nil {
		return err
	}
	return nil
}
