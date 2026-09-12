package repo

import (
	"net/http"

	"code.gitea.io/gitea/models"

	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/convert"
	"code.gitea.io/gitea/modules/log"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/routers/api/v1/utils"
)

//标注任务可分配人员
func ListTagger(ctx *context.APIContext) {

	taggers := make([]*api.Tagger, 0)
	userRemember := make(map[string]string)
	collaborators, err := ctx.Repo.Repository.GetCollaborators(utils.GetListOptions(ctx))
	if err != nil {
		log.Warn("ListCollaborators", err)
		ctx.JSON(http.StatusOK, taggers)
		return
	}
	for _, collaborator := range collaborators {
		taggers = append(taggers, convert.ToTagger(collaborator.User))
		userRemember[collaborator.User.Name] = ""
	}

	teams, err := ctx.Repo.Repository.GetRepoTeams()
	if err != nil {
		log.Warn("ListTeams", err)
		ctx.JSON(http.StatusOK, taggers)
		return
	}

	for _, team := range teams {
		team.GetMembers(&models.SearchMembersOptions{})
		for _, user := range team.Members {
			if _, ok := userRemember[user.Name]; !ok {
				taggers = append(taggers, convert.ToTagger(user))
				userRemember[user.Name] = ""
			}
		}
	}
	if !ctx.Repo.Owner.IsOrganization() {
		if _, ok := userRemember[ctx.Repo.Owner.Name]; !ok {
			taggers = append(taggers, convert.ToTagger(ctx.Repo.Owner))

		}
	}
	ctx.JSON(http.StatusOK, taggers)

}
func GetRight(ctx *context.APIContext) {
	right := "none"

	if ctx.IsUserRepoReaderSpecific(models.UnitTypeCode) {
		right = "read"
	}

	if ctx.IsUserRepoWriter([]models.UnitType{models.UnitTypeCode}) || ctx.IsUserRepoAdmin() {
		right = "write"
	}

	ctx.JSON(http.StatusOK, map[string]string{
		"right": right,
	})

}

func GetCloudBrainJobId(ctx *context.APIContext) {
	cloudbrains, err := models.GetCloudbrainsByDisplayJobName(ctx.Repo.Repository.ID, ctx.Query("jobType"), ctx.Query("name"))
	if err != nil {
		log.Warn("get cloudbrain by display name failed", err)
		ctx.JSON(http.StatusOK, map[string]string{"jobId": ""})
		return
	}
	if len(cloudbrains) > 0 {
		ctx.JSON(http.StatusOK, map[string]string{"jobId": cloudbrains[0].JobID})
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{"jobId": ""})
}
