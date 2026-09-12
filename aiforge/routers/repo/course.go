package repo

import (
	"net/http"
	"strings"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/auth"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	repo_service "code.gitea.io/gitea/services/repository"
)

const (
	tplCreateCourse base.TplName = "repo/createCourse"
)

func CreateCourse(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("new_course")
	org, _ := models.GetUserByName(setting.Course.OrgName)

	ctx.Data["Owner"] = org
	ctx.Data["IsCourse"] = true

	ctx.HTML(200, tplCreateCourse)

}

func CreateCoursePost(ctx *context.Context, form auth.CreateCourseForm) {
	ctx.Data["Title"] = ctx.Tr("new_course")

	if ctx.Written() {
		return
	}

	org, _ := models.GetUserByName(setting.Course.OrgName)

	ctx.Data["Owner"] = org
	ctx.Data["IsCourse"] = true

	var topics = make([]string, 0)
	var topicsStr = strings.TrimSpace(form.Topics)
	if len(topicsStr) > 0 {
		topics = strings.Split(topicsStr, ",")
	}

	validTopics, invalidTopics := models.SanitizeAndValidateTopics(topics)

	if len(validTopics) > 25 {
		ctx.RenderWithErr(ctx.Tr("repo.topic.count_prompt"), tplCreateCourse, form)
		return
	}

	if len(invalidTopics) > 0 {
		ctx.RenderWithErr(ctx.Tr("repo.topic.format_prompt"), tplCreateCourse, form)
		return
	}

	var repo *models.Repository
	var err error

	if setting.Course.OrgName == "" || setting.Course.TeamName == "" {
		log.Error("then organization name or team name of course is empty.")
		ctx.RenderWithErr(ctx.Tr("repo.failed_to_create_course"), tplCreateCourse, form)
		return
	}

	org, team, err := getOrgAndTeam()

	if err != nil {
		log.Error("Failed to get team from db.", err)
		ctx.RenderWithErr(ctx.Tr("repo.failed_to_create_course"), tplCreateCourse, form)
		return
	}
	isInTeam, err := models.IsUserInTeams(ctx.User.ID, []int64{team.ID})
	if err != nil {
		log.Error("Failed to get user in team from db.")
		ctx.RenderWithErr(ctx.Tr("repo.failed_to_create_course"), tplCreateCourse, form)
		return
	}

	if !isInTeam {
		err = models.AddTeamMember(team, ctx.User.ID)
		if err != nil {
			log.Error("Failed to add user to team.")
			ctx.RenderWithErr(ctx.Tr("repo.failed_to_create_course"), tplCreateCourse, form)
			return
		}
	}

	if ctx.HasError() {
		ctx.HTML(200, tplCreateCourse)
		return
	}

	repo, err = repo_service.CreateRepository(ctx.User, org, models.CreateRepoOptions{
		Name:          form.RepoName,
		Alias:         form.Alias,
		Description:   form.Description,
		Gitignores:    "",
		IssueLabels:   "",
		License:       "",
		Readme:        "Default",
		IsPrivate:     false,
		DefaultBranch: "master",
		AutoInit:      true,
		IsCourse:      true,
		Topics:        validTopics,
	})
	if err == nil {
		log.Trace("Repository created [%d]: %s/%s", repo.ID, org.Name, repo.Name)
		ctx.Redirect(setting.AppSubURL + "/" + org.Name + "/" + repo.Name)
		return
	}

	handleCreateCourseError(ctx, org, err, "CreateCoursePost", tplCreateCourse, &form)
}

func AddCourseOrg(ctx *context.Context) {

	_, team, err := getOrgAndTeam()

	if err != nil {
		log.Error("Failed to get team from db.", err)
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    1,
			"message": ctx.Tr("repo.addCourseOrgFail"),
		})
		return
	}
	isInTeam, err := models.IsUserInTeams(ctx.User.ID, []int64{team.ID})
	if err != nil {
		log.Error("Failed to get user in team from db.", err)
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    1,
			"message": ctx.Tr("repo.add_course_org_fail"),
		})
		return
	}

	if !isInTeam {
		err = models.AddTeamMember(team, ctx.User.ID)
		if err != nil {
			log.Error("Failed to add user to team.", err)
			ctx.JSON(http.StatusOK, map[string]interface{}{
				"code":    1,
				"message": ctx.Tr("repo.add_course_org_fail"),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "",
	})

}

func getOrgAndTeam() (*models.User, *models.Team, error) {
	org, err := models.GetUserByName(setting.Course.OrgName)

	if err != nil {
		log.Error("Failed to get organization from db.", err)
		return nil, nil, err
	}

	team, err := models.GetTeam(org.ID, setting.Course.TeamName)

	if err != nil {
		log.Error("Failed to get team from db.", err)

		return nil, nil, err
	}
	return org, team, nil
}

func handleCreateCourseError(ctx *context.Context, owner *models.User, err error, name string, tpl base.TplName, form interface{}) {
	switch {
	case models.IsErrReachLimitOfRepo(err):
		ctx.RenderWithErr(ctx.Tr("repo.form.reach_limit_of_course_creation", owner.MaxCreationLimit()), tpl, form)
	case models.IsErrRepoAlreadyExist(err):
		ctx.Data["Err_RepoName"] = true
		ctx.RenderWithErr(ctx.Tr("form.course_name_been_taken"), tpl, form)
	case models.IsErrNameReserved(err):
		ctx.Data["Err_RepoName"] = true
		ctx.RenderWithErr(ctx.Tr("repo.form.course_name_reserved", err.(models.ErrNameReserved).Name), tpl, form)
	case models.IsErrNamePatternNotAllowed(err):
		ctx.Data["Err_RepoName"] = true
		ctx.RenderWithErr(ctx.Tr("repo.form.course_name_pattern_not_allowed", err.(models.ErrNamePatternNotAllowed).Pattern), tpl, form)
	default:
		ctx.ServerError(name, err)
	}
}
