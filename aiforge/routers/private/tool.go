// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package private

import (
	"net/http"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/repo"
	"code.gitea.io/gitea/routers/response"
	"gitea.com/macaron/macaron"
)

func UpdateAllRepoCommitCnt(ctx *macaron.Context) {
	repos, err := models.GetAllRepositories()
	if err != nil {
		log.Error("GetAllRepositories failed:%v", err.Error(), ctx.Data["MsgID"])
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error_msg": "GetAllRepositories failed",
		})
		return
	}

	for i, repo := range repos {
		log.Info("%d:begin updateRepoCommitCnt(id = %d, name = %s)", i, repo.ID, repo.Name)
		if err = updateRepoCommitCnt(ctx, repo); err != nil {
			log.Error("updateRepoCommitCnt(id = %d, name = %s) failed:%v", repo.ID, repo.Name, err.Error(), ctx.Data["MsgID"])
			continue
		}
		log.Info("%d:finish updateRepoCommitCnt(id = %d, name = %s)", i, repo.ID, repo.Name)
	}

	ctx.JSON(http.StatusOK, map[string]string{
		"error_msg": "",
	})
}

func YearReportCollection(ctx *macaron.Context) {

	go repo.UpdateYearReportCollection()
	ctx.JSON(http.StatusOK, map[string]string{
		"error_msg": "",
	})
}

func RepoStatisticManually(ctx *macaron.Context) {
	date := ctx.Params("date")
	repo.RepoStatisticDaily(date)
	//repo.SummaryStatisticDaily(date)
	//repo.TimingCountDataByDate(date)
}

func UserStatisticManually(ctx *macaron.Context) {
	date := ctx.Query("date")
	log.Info("UserStatisticManually api received; date=%s", date)
	if date == "" {
		ctx.JSON(http.StatusBadRequest, response.OuterError(400, "'date' parameter is empty"))
	}
	repo.TimingCountDataByDate(date)
	ctx.JSON(http.StatusOK, response.OuterSuccess())
}

func OrgStatisticManually() {
	models.UpdateOrgStatistics()
}

func UpdateRepoVisit(ctx *macaron.Context) {
	date := ctx.Params("date")
	log.Info("date(%s)", date)

	repos, err := models.GetAllRepositories()
	if err != nil {
		log.Error("GetAllRepositories failed:%v", err.Error(), ctx.Data["MsgID"])
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error_msg": "GetAllRepositories failed",
		})
		return
	}

	for i, repoStat := range repos {
		log.Info("%d:begin UpdateRepoVisits(id = %d, name = %s)", i, repoStat.ID, repoStat.Name)
		if err = repo.UpdateRepoVisits(ctx, repoStat, date); err != nil {
			log.Error("UpdateRepoVisits(id = %d, name = %s) failed:%v", repoStat.ID, repoStat.Name, err.Error())
			continue
		}
		log.Info("%d:finish UpdateRepoVisits(id = %d, name = %s)", i, repoStat.ID, repoStat.Name)
	}

	ctx.JSON(http.StatusOK, map[string]string{
		"error_msg": "",
	})
}
