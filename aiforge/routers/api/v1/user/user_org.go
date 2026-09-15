package user

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"net/http"
	"strings"
)

// UserOrganization 获取当前用户组织数据
func UserOrganization(ctx *context.APIContext) {
	var (
		keyword = strings.Trim(ctx.Query("q"), " ")
		resp    = make([]*models.PlatFormOrg, 0, 0)
	)

	// This endpoint is "my organizations". Always include private memberships.
	// Team add-member defaults org_user.is_public=false, so filtering to public
	// memberships (the old admin-only showPrivate) made joined orgs invisible
	// to every non-admin user.
	log.Info("UserOrganization start reqId[%v] keyword[%v]", ctx.ReqId, keyword)

	orgs, err := models.GetOrgsByName(ctx.User.ID, true, keyword)
	if err != nil {
		ctx.ServerError("UserOrganization", err)
		return
	}

	for _, org := range orgs {
		resp = append(resp, &models.PlatFormOrg{
			OrgID:         org.ID,
			OrgName:       org.Name,
			RelAvatarLink: org.AvatarLink(),
			Description:   org.Description,
		})
	}

	ctx.JSON(http.StatusOK, response.SuccessWithData(resp))
	return
}
