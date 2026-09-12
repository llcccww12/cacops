package access

import (
	"strings"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/routers/utils"
)

func GetCollaborators(ctx *context.Context) {
	users, err := ctx.AccessContext.SubjectAccessContext.GetCollaborators(models.ListOptions{})
	if err != nil {
		ctx.JSON(200, response.OuterBizError(response.NewBizError(err)))
		return
	}

	teams, err := ctx.AccessContext.SubjectAccessContext.GetTeams()
	if err != nil {
		ctx.JSON(200, response.OuterBizError(response.NewBizError(err)))
		return
	}
	collaborators4Front := make([]*models.SubjectCollaborator4Front, 0, len(users))
	for i := 0; i < len(users); i++ {
		collaborators4Front = append(collaborators4Front, users[i].ToFrontFormat())
	}
	teamFronts := make([]*models.Team4Front, 0, len(users))
	var userHasDatasetTeamChangeAccess = ctx.AccessContext.SubjectAccessContext.Owner.DatasetAdminChangeTeamAccess || ctx.AccessContext.IsOwner()
	var userHasAimodelTeamChangeAccess = ctx.AccessContext.SubjectAccessContext.Owner.AimodelAdminChangeTeamAccess || ctx.AccessContext.IsOwner()

	for i := 0; i < len(teams); i++ {
		t := teams[i]
		var allowedToChangeDatasetTeams bool
		var allowedToChangeAimodelTeams bool
		if !t.IncludesAllDatasets {
			allowedToChangeDatasetTeams = userHasDatasetTeamChangeAccess
		}
		if !t.IncludesAllDatasets {
			allowedToChangeAimodelTeams = userHasAimodelTeamChangeAccess
		}

		teamFronts = append(teamFronts, &models.Team4Front{
			ID:                          t.ID,
			Name:                        t.Name,
			NumMembers:                  t.NumMembers,
			NumRepos:                    t.NumRepos,
			NumDatasets:                 t.NumDatasets,
			IncludesAllDatasets:         t.IncludesAllDatasets,
			AllowedToChangeDatasetTeams: allowedToChangeDatasetTeams,
			DatasetAuthorize:            t.DatasetAuthorize,
			OwnerName:                   ctx.AccessContext.SubjectAccessContext.Owner.Name,
			NumAimodels:                 t.NumAimodels,
			IncludesAllAimodels:         t.IncludesAllAimodels,
			AllowedToChangeAimodelTeams: allowedToChangeAimodelTeams,
			AimodelAuthorize:            t.AimodelAuthorize,
		})
	}
	m := map[string]interface{}{
		"Collaborators":         collaborators4Front,
		"Teams":                 teamFronts,
		"CanChangeDatasetTeams": userHasDatasetTeamChangeAccess,
		"CanChangeAimodelTeams": userHasAimodelTeamChangeAccess,
	}
	ctx.JSON(200, response.OuterSuccessWithData(m))
}

func AddCollaborator(ctx *context.Context) {
	name := utils.RemoveUsernameParameterSuffix(strings.ToLower(ctx.Query("collaborator")))
	if len(name) == 0 {
		ctx.JSON(200, response.OuterTrBizError(response.PARAM_ERROR, ctx.Locale))
		return
	}

	u, err := models.GetUserByName(name)
	if err != nil {
		if models.IsErrUserNotExist(err) {
			ctx.JSON(200, response.OuterTrBizError(response.USER_NOT_EXISTS, ctx.Locale))
			return
		}
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}

	// Organization is not allowed to be added as a collaborator.
	if u.IsOrganization() {
		ctx.JSON(200, response.OuterTrBizError(response.ORG_NOT_ALLOWED_TO_BE_COLLABORATOR, ctx.Locale))
		return
	}
	subjectCtx := ctx.AccessContext.SubjectAccessContext
	if subjectCtx.OwnerID == u.ID {
		ctx.JSON(200, response.OuterTrBizError(response.OWNER_NOT_ALLOWED_TO_BE_COLLABORATOR, ctx.Locale))
		return
	}

	if got, err := subjectCtx.IsCollaborator(u.ID); err == nil && got {
		ctx.JSON(200, response.OuterTrBizError(response.ADD_COLLABORATOR_DUPLICATE, ctx.Locale))
		return
	}

	if err = subjectCtx.AddCollaborator(u); err != nil {
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))

		return
	}

	ctx.JSON(200, response.OuterSuccess())
}

func ChangeCollaborationAccessMode(ctx *context.Context) {
	if err := ctx.AccessContext.SubjectAccessContext.ChangeCollaborationAccessMode(
		ctx.QueryInt64("uid"),
		models.AccessMode(ctx.QueryInt("mode"))); err != nil {
		log.Error("ChangeCollaborationAccessMode: %v", err)
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return
	}
	ctx.JSON(200, response.OuterSuccess())
}

func DeleteCollaboration(ctx *context.Context) {
	if err := ctx.AccessContext.SubjectAccessContext.DeleteCollaboration(ctx.QueryInt64("uid")); err != nil {
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
	}

	ctx.JSON(200, response.OuterSuccess())
}

func AddDatasetTeam(ctx *context.Context) {
	if ctx.AccessContext.SubjectAccessContext.Dataset == nil {
		ctx.JSON(200, response.OuterTrBizError(response.INSUFFICIENT_PERMISSION, ctx.Locale))
		return
	}
	if !ctx.AccessContext.SubjectAccessContext.Owner.DatasetAdminChangeTeamAccess && !ctx.AccessContext.IsOwner() {
		ctx.JSON(200, response.OuterTrBizError(response.CHANGE_TEAM_ACCESS_NOT_ALLOWED, ctx.Locale))
		return
	}

	name := utils.RemoveUsernameParameterSuffix(strings.ToLower(ctx.Query("team")))
	if len(name) == 0 {
		ctx.JSON(200, response.OuterTrBizError(response.PARAM_ERROR, ctx.Locale))
		return
	}
	subjectCtx := ctx.AccessContext.SubjectAccessContext
	subjectCtx.GetOwner()
	team, err := subjectCtx.Owner.GetTeam(name)
	if err != nil {
		if models.IsErrTeamNotExist(err) {
			ctx.JSON(200, response.OuterTrBizError(response.TEAM_NOT_EXIST, ctx.Locale))
			return
		}
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return

	}

	if team.OrgID != subjectCtx.OwnerID {
		ctx.JSON(200, response.OuterTrBizError(response.TEAM_NOT_IN_ORGANIZATION, ctx.Locale))
		return
	}

	if models.HasTeamSubject(subjectCtx.OwnerID, team.ID, subjectCtx.SubjectID, subjectCtx.SubjectType) {
		ctx.JSON(200, response.OuterTrBizError(response.ADD_TEAM_DUPLICATE, ctx.Locale))
		return
	}

	if err = team.AddDataset(subjectCtx.Dataset); err != nil {
		ctx.ServerError("team.AddDataset", err)
		return
	}
	ctx.JSON(200, response.OuterSuccess())
}

func DeleteDatasetTeam(ctx *context.Context) {
	if !ctx.AccessContext.SubjectAccessContext.Owner.DatasetAdminChangeTeamAccess && !ctx.AccessContext.IsOwner() {
		ctx.JSON(200, response.OuterTrBizError(response.CHANGE_TEAM_ACCESS_NOT_ALLOWED, ctx.Locale))
		return
	}

	team, err := models.GetTeamByID(ctx.QueryInt64("id"))
	if err != nil {
		ctx.ServerError("GetTeamByID", err)
		return
	}

	if err = team.RemoveDataset(ctx.AccessContext.SubjectAccessContext.SubjectID); err != nil {
		ctx.ServerError("team.RemoveRepositorys", err)
		return
	}

	ctx.JSON(200, response.OuterSuccess())
}

func AddAimodelTeam(ctx *context.Context) {
	if ctx.AccessContext.SubjectAccessContext.Aimodel == nil {
		ctx.JSON(200, response.OuterTrBizError(response.INSUFFICIENT_PERMISSION, ctx.Locale))
		return
	}
	if !ctx.AccessContext.SubjectAccessContext.Owner.AimodelAdminChangeTeamAccess && !ctx.AccessContext.IsOwner() {
		ctx.JSON(200, response.OuterTrBizError(response.CHANGE_TEAM_ACCESS_NOT_ALLOWED, ctx.Locale))
		return
	}

	name := utils.RemoveUsernameParameterSuffix(strings.ToLower(ctx.Query("team")))
	if len(name) == 0 {
		ctx.JSON(200, response.OuterTrBizError(response.PARAM_ERROR, ctx.Locale))
		return
	}
	subjectCtx := ctx.AccessContext.SubjectAccessContext
	subjectCtx.GetOwner()
	team, err := subjectCtx.Owner.GetTeam(name)
	if err != nil {
		if models.IsErrTeamNotExist(err) {
			ctx.JSON(200, response.OuterTrBizError(response.TEAM_NOT_EXIST, ctx.Locale))
			return
		}
		ctx.JSON(200, response.OuterTrBizError(response.NewBizError(err), ctx.Locale))
		return

	}

	if team.OrgID != subjectCtx.OwnerID {
		ctx.JSON(200, response.OuterTrBizError(response.TEAM_NOT_IN_ORGANIZATION, ctx.Locale))
		return
	}

	if models.HasTeamSubject(subjectCtx.OwnerID, team.ID, subjectCtx.SubjectID, subjectCtx.SubjectType) {
		ctx.JSON(200, response.OuterTrBizError(response.ADD_TEAM_DUPLICATE, ctx.Locale))
		return
	}

	if err = team.AddAimodel(subjectCtx.Aimodel); err != nil {
		ctx.ServerError("team.AddAimodel", err)
		return
	}
	ctx.JSON(200, response.OuterSuccess())
}

func DeleteAimodelTeam(ctx *context.Context) {
	if !ctx.AccessContext.SubjectAccessContext.Owner.AimodelAdminChangeTeamAccess && !ctx.AccessContext.IsOwner() {
		ctx.JSON(200, response.OuterTrBizError(response.CHANGE_TEAM_ACCESS_NOT_ALLOWED, ctx.Locale))
		return
	}

	team, err := models.GetTeamByID(ctx.QueryInt64("id"))
	if err != nil {
		ctx.ServerError("GetTeamByID", err)
		return
	}

	if err = team.RemoveAimodel(ctx.AccessContext.SubjectAccessContext.SubjectID); err != nil {
		ctx.ServerError("team.RemoveRepositorys", err)
		return
	}

	ctx.JSON(200, response.OuterSuccess())
}
