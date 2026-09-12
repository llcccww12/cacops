// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2020 The Gitea Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package operation

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/routers"
)

const (
	tplOrgs base.TplName = "admin/org/list"
)

type UpdateRecommendOrgs struct {
	OrgInfos string `binding:"required"`
}

type OrgInfo struct {
	OrgID int64 `json:"org_id"`
	Order int64 `json:"order"`
}

type OrgInfos struct {
	OrgInfo []OrgInfo `json:"org_infos"`
}

// Organizations show all the organizations recommended
func Organizations(ctx *context.Context) {
	ctx.Data["PageIsAdmin"] = true
	ctx.Data["PageIsAdminOrganizations"] = true

	routers.RenderUserSearch(ctx, &models.SearchUserOptions{
		Type: models.UserTypeOrganization,
		ListOptions: models.ListOptions{
			PageSize: setting.UI.Admin.OrgPagingNum,
		},
		Visible: []structs.VisibleType{structs.VisibleTypePublic, structs.VisibleTypeLimited, structs.VisibleTypePrivate},
	}, tplOrgs)
}

// UpdateRecommendOrganizations update the organizations recommended
func UpdateRecommendOrganizations(ctx *context.Context, req OrgInfos) {
	orgs := make(models.RecommendOrgList, 0, 10)
	for _, org := range req.OrgInfo {
		orgs = append(orgs, &models.RecommendOrg{
			OrgID: org.OrgID,
			Order: org.Order,
		})
	}

	if err := models.UpdateRecommendOrgs(orgs); err != nil {
		log.Error("UpdateRecommendOrgs failed:%v", err.Error(), ctx.Data["MsgID"])
		ctx.ServerError("UpdateRecommendOrgs failed", err)
		return
	}
}
