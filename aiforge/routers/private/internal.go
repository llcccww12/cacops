// Copyright 2017 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package private includes all internal routes. The package name internal is ideal but Golang is not allowed, so we use private as package name instead.
package private

import (
	"code.gitea.io/gitea/services/cloudbrain/resource"
	"code.gitea.io/gitea/services/repository"
	"strings"

	"code.gitea.io/gitea/routers/admin"

	"code.gitea.io/gitea/routers/repo"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/private"
	"code.gitea.io/gitea/modules/setting"

	"gitea.com/macaron/binding"
	"gitea.com/macaron/macaron"
)

// CheckInternalToken check internal token is set
func CheckInternalToken(ctx *macaron.Context) {
	tokens := ctx.Req.Header.Get("Authorization")
	fields := strings.Fields(tokens)
	if len(fields) != 2 || fields[0] != "Bearer" || fields[1] != setting.InternalToken {
		log.Debug("Forbidden attempt to access internal url: Authorization header: %s", tokens)
		ctx.Error(403)
	}
}

// RegisterRoutes registers all internal APIs routes to web application.
// These APIs will be invoked by internal commands for example `gitea serv` and etc.
func RegisterRoutes(m *macaron.Macaron) {
	bind := binding.Bind

	m.Group("/", func() {
		m.Post("/ssh/authorized_keys", AuthorizedPublicKeyByContent)
		m.Post("/ssh/:id/update/:repoid", UpdatePublicKeyInRepo)
		m.Post("/hook/pre-receive/:owner/:repo", bind(private.HookOptions{}), HookPreReceive)
		m.Post("/hook/post-receive/:owner/:repo", bind(private.HookOptions{}), HookPostReceive)
		m.Post("/hook/set-default-branch/:owner/:repo/:branch", SetDefaultBranch)
		m.Get("/hook/env/:owner/:repo", HookEnv)
		m.Get("/serv/none/:keyid", ServNoCommand)
		m.Get("/serv/command/:keyid/:owner/:repo", ServCommand)
		m.Post("/manager/shutdown", Shutdown)
		m.Post("/manager/restart", Restart)
		m.Post("/manager/flush-queues", bind(private.FlushOptions{}), FlushQueues)
		m.Post("/tool/update_all_repo_commit_cnt", UpdateAllRepoCommitCnt)
		m.Post("/tool/repo_stat/:date", RepoStatisticManually)
		//m.Post("/tool/user_stat/:date", UserStatisticManually)
		m.Get("/tool/year_report", YearReportCollection)
		m.Get("/tool/org_stat", OrgStatisticManually)
		m.Post("/tool/update_repo_visit/:date", UpdateRepoVisit)
		m.Post("/task/history_handle/duration", repo.HandleTaskWithNoDuration)
		m.Post("/task/history_handle/aicenter", repo.HandleTaskWithAiCenter)
		m.Post("/resources/specification/handle_historical_task", admin.RefreshHistorySpec)
		m.Post("/repos/cnt_stat/handle_historical_task", admin.RefreshHistorySpec)
		m.Post("/cloudbrain_duration_statisctic/history_handle", repo.CloudbrainDurationUpdateHistoryData)
		m.Post("/cloudbrain_task_num_statisctic/history_handle", repo.CloudbrainTaskNumUpdateHistoryData)
		m.Post("/queue_duration_statisctic/history_handle", repo.QueueTaskDurationUpdateHistoryData)
		m.Post("/queue_task_num_statisctic/history_handle", repo.QueueTaskNumUpdateHistoryData)
		m.Post("/square/repo/stat/refresh", repository.RefreshRepoStatData)
		m.Get("/setting/refresh", RefreshSetting)
		m.Get("/resource/SyncGrampusSpecPools", resource.SyncGrampusSpecPools)

	}, CheckInternalToken)
}
