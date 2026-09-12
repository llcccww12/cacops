package tech

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/auth"
	"code.gitea.io/gitea/modules/context"
	api "code.gitea.io/gitea/modules/structs"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/repository"
	techService "code.gitea.io/gitea/services/tech"
	"errors"
	"fmt"
	"net/http"
)

//CommitOpenIRepo 新建启智项目申请页面提交
func CommitOpenIRepo(ctx *context.APIContext, form api.OpenITechRepo) {
	//解析项目路径，查找项目
	repo, err := repository.FindRepoByUrl(form.Url)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterServerError(ctx.Tr(err.Error())))
		return
	}
	//查找项目编号
	tech, err := models.GetTechByTechNo(form.TechNo)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterServerError(ctx.Tr(err.Error())))
		return
	}

	//判断承接单位是否在科技项目的参与单位中
	if !tech.IsValidInstitution(form.Institution) {
		ctx.JSON(http.StatusOK, response.OuterServerError(ctx.Tr("tech.institution_not_valid")))
		return
	}

	//判断是否已经存在了该项目
	exist, err := techService.IsValidRepoConvergeExists(repo.ID, tech.ID)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterServerError(ctx.Tr(err.Error())))
		return
	}
	if exist {
		ctx.JSON(http.StatusOK, response.OuterServerError(fmt.Sprintf(ctx.Tr("tech.repo_converge_exists"), tech.ProjectName, repo.Alias)))
		return
	}

	//写入数据库
	rci := &models.RepoConvergeInfo{
		RepoID:      repo.ID,
		Url:         form.Url,
		BaseInfoID:  tech.ID,
		Institution: form.Institution,
		UID:         ctx.User.ID,
		Status:      models.DefaultTechApprovedStatus,
	}
	err = rci.InsertOrUpdate()
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterServerError(ctx.Tr(err.Error())))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

//CommitNotOpenIRepo 新建非启智项目申请页面提交
func CommitNotOpenIRepo(ctx *context.APIContext, form api.NotOpenITechRepo) {
	//触发更新迁移状态
	go techService.UpdateTechMigrateStatus()

	//查找项目编号
	tech, err := models.GetTechByTechNo(form.TechNo)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterServerError(ctx.Tr(err.Error())))
		return
	}

	//判断承接单位是否在科技项目的参与单位中
	if !tech.IsValidInstitution(form.Institution) {
		ctx.JSON(http.StatusOK, response.OuterServerError(ctx.Tr("tech.institution_not_valid")))
		return
	}

	//调用迁移接口
	bizErr := repository.MigrateSubmit(ctx.User, auth.MigrateRepoForm{
		CloneAddr:   form.Url,
		UID:         form.UID,
		RepoName:    form.RepoName,
		Alias:       form.Alias,
		Description: form.Description,
		Labels:      false,
		Mirror:      true,
	})
	if bizErr != nil {
		if bizErr.Code == 3 {
			ctx.JSON(http.StatusOK, response.OuterError(bizErr.Code, ctx.Tr("tech.to_migrate_repo_exists")))
			return
		}
		ctx.JSON(http.StatusOK, response.OuterError(bizErr.Code, ctx.Tr(bizErr.DefaultMsg)))
		return
	}

	repo, err := models.GetRepositoryByName(form.UID, form.RepoName)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterServerError(ctx.Tr(err.Error())))
		return
	}

	//判断是否已经存在了该项目
	exist, err := techService.IsValidRepoConvergeExists(repo.ID, tech.ID)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterServerError(ctx.Tr(err.Error())))
		return
	}
	if exist {
		ctx.JSON(http.StatusOK, response.OuterServerError(fmt.Sprintf(ctx.Tr("tech.repo_converge_exists"), tech.ProjectName, repo.Alias)))
		return
	}

	//给仓库加标签
	updateTopics(repo.ID, form.Topics)

	//写入数据库
	rci := &models.RepoConvergeInfo{
		RepoID:      repo.ID,
		Url:         form.Url,
		BaseInfoID:  tech.ID,
		Institution: form.Institution,
		UID:         ctx.User.ID,
		Status:      models.TechMigrating,
	}
	err = rci.InsertOrUpdate()
	if err != nil {
		ctx.JSON(http.StatusOK, response.OuterServerError(ctx.Tr(err.Error())))
		return
	}
	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func updateTopics(repoId int64, topicNames []string) error {
	validTopics, invalidTopics := models.SanitizeAndValidateTopics(topicNames)

	if len(validTopics) > 25 {
		return errors.New("Exceeding maximum number of topics per repo")
	}

	if len(invalidTopics) > 0 {
		return errors.New("Topic names are invalid")
	}

	err := models.SaveTopics(repoId, validTopics...)
	if err != nil {
		return err
	}
	return nil
}
