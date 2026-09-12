//implementation of git local fetcher
package fetcher

import (
	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/models"
)

type DatasetStatisticFetcher struct {
}

func NewDatasetStatisticFetcher() *DatasetStatisticFetcher {
	return &DatasetStatisticFetcher{}
}

func (g *DatasetStatisticFetcher) Fetch(path string) (*entity.DynConfig, error) {
	datasetMap, _ := models.QueryDatasetGroupByTask(8)
	if len(datasetMap) == 0 {
		return nil, nil
	}
	return &entity.DynConfig{
		ConfigType: entity.DynConfigTypeRaw,
		ValueType:  entity.ValueTypeObject,
		Value:      datasetMap,
	}, nil
}
