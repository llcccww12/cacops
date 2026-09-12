package oauth2_service

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
)

func GetOauth2ApplicationList(req models.SearchOauth2ApplicationReq) ([]*models.Oauth2Application4Show, int64, *response.BizError) {
	list, total, err := models.SearchOauth2ApplicationList(req)
	if err != nil {
		log.Error("SearchOauth2ApplicationList failed, req=%+v, err=%v", req, err)
		return nil, 0, response.NewBizError(err)
	}
	if list == nil || len(list) == 0 {
		log.Info("SearchOauth2ApplicationList no data found, req=%+v", req)
		return nil, 0, nil
	}

	return list, total, nil

}
