package resource

import (
	"code.gitea.io/gitea/models"
)

func AddResourceScene(req models.ResourceSceneReq) error {

	if err := models.InsertResourceScene(req); err != nil {
		return err
	}
	return nil
}

func UpdateResourceScene(req models.ResourceSceneReq) error {
	if err := models.UpdateResourceScene(req); err != nil {
		return err
	}
	return nil
}

func DeleteResourceScene(id int64) error {
	if err := models.DeleteResourceScene(id); err != nil {
		return err
	}
	return nil
}

func GetResourceSceneList(opts models.SearchResourceSceneOptions) (*models.ResourceSceneListRes, error) {
	n, r, err := models.SearchResourceScene(opts)
	if err != nil {
		return nil, err
	}

	return models.NewResourceSceneListRes(n, r), nil
}
