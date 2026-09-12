// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package org

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/setting"
)

const (
	tplOrgHome       base.TplName = "org/home"
	tplOrgCourseHome base.TplName = "org/home_courses"
)

// Home show organization home page
func Home(ctx *context.Context) {
	ctx.SetParams(":org", ctx.Params(":username"))
	ctx.Data["PageIsOrgHome"] = true
	ctx.Data["PageIsOrgHomeModel"] = ctx.Query("model")
	ctx.Data["PageIsOrgHomeDataset"] = ctx.Query("dataset")
	ctx.Data["PageIsOrgHomeStorage"] = ctx.Query("storage")
	context.HandleOrgAssignment(ctx)
	if ctx.Written() {
		return
	}

	org := ctx.Org.Organization

	if !models.HasOrgVisible(org, ctx.User) {
		ctx.NotFound("HasOrgVisible", nil)
		return
	}

	ctx.Data["Title"] = org.DisplayName()

	var opts = models.FindOrgMembersOpts{
		OrgID:       org.ID,
		PublicOnly:  true,
		ListOptions: models.ListOptions{Page: 1, PageSize: 25},
	}

	if ctx.User != nil {
		isMember, err := org.IsOrgMember(ctx.User.ID)
		if err != nil {
			ctx.Error(500, "IsOrgMember")
			return
		}
		opts.PublicOnly = !isMember && !ctx.User.IsAdmin
	}

	members, _, err := models.FindOrgMembers(&opts)
	if err != nil {
		ctx.ServerError("FindOrgMembers", err)
		return
	}

	membersCount, err := models.CountOrgMembers(opts)
	if err != nil {
		ctx.ServerError("CountOrgMembers", err)
		return
	}

	ctx.Data["MembersTotal"] = membersCount
	ctx.Data["Members"] = members
	ctx.Data["Teams"] = org.Teams

	if setting.Course.OrgName == org.Name {
		ctx.HTML(200, tplOrgCourseHome)
	} else {
		ctx.HTML(200, tplOrgHome)
	}
}
