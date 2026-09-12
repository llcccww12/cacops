package repository

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
)

func ConvertToDatasetWithStar(ctx *context.Context, datasets []*models.Dataset) []*models.DatasetWithStar {
	var datasetsWithStar []*models.DatasetWithStar
	for _, dataset := range datasets {
		if !ctx.IsSigned {
			datasetsWithStar = append(datasetsWithStar, &models.DatasetWithStar{Dataset: *dataset, IsStaring: false})
		} else {
			datasetsWithStar = append(datasetsWithStar, &models.DatasetWithStar{Dataset: *dataset, IsStaring: models.IsDatasetStaring(ctx.User.ID, dataset.ID)})
		}

	}
	return datasetsWithStar
}
