// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cron

import (
	"context"
	"time"

	"code.gitea.io/gitea/routers/api/v1/hf_model"

	"code.gitea.io/gitea/routers/api/v1/finetune"
	"code.gitea.io/gitea/services/ai_task_service/schedule"
	"code.gitea.io/gitea/services/ai_task_service/task"

	"code.gitea.io/gitea/modules/setting"

	"code.gitea.io/gitea/modules/urfs_client/urchin"
	cloudbrainService "code.gitea.io/gitea/services/cloudbrain"

	"code.gitea.io/gitea/modules/attachment"
	"code.gitea.io/gitea/modules/modelarts"
	"code.gitea.io/gitea/services/cloudbrain/resource"
	"code.gitea.io/gitea/services/reward"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/migrations"
	repository_service "code.gitea.io/gitea/modules/repository"
	"code.gitea.io/gitea/routers/repo"
	mirror_service "code.gitea.io/gitea/services/mirror"
	subject_service "code.gitea.io/gitea/services/subject_service"
)

func registerUpdateMirrorTask() {
	RegisterTaskFatal("update_mirrors", &BaseConfig{
		Enabled:    true,
		RunAtStart: false,
		Schedule:   "@every 10m",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		return mirror_service.Update(ctx)
	})
}

func registerRepoHealthCheck() {
	type RepoHealthCheckConfig struct {
		BaseConfig
		Timeout time.Duration
		Args    []string `delim:" "`
	}
	RegisterTaskFatal("repo_health_check", &RepoHealthCheckConfig{
		BaseConfig: BaseConfig{
			Enabled:    true,
			RunAtStart: false,
			Schedule:   "@every 24h",
		},
		Timeout: 60 * time.Second,
		Args:    []string{},
	}, func(ctx context.Context, _ *models.User, config Config) error {
		rhcConfig := config.(*RepoHealthCheckConfig)
		return repository_service.GitFsck(ctx, rhcConfig.Timeout, rhcConfig.Args)
	})
}

func registerCheckRepoStats() {
	RegisterTaskFatal("check_repo_stats", &BaseConfig{
		Enabled:    true,
		RunAtStart: true,
		Schedule:   "@every 24h",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		return models.CheckRepoStats(ctx)
	})
}

func registerArchiveCleanup() {
	RegisterTaskFatal("archive_cleanup", &OlderThanConfig{
		BaseConfig: BaseConfig{
			Enabled:    true,
			RunAtStart: true,
			Schedule:   "@every 24h",
		},
		OlderThan: 24 * time.Hour,
	}, func(ctx context.Context, _ *models.User, config Config) error {
		acConfig := config.(*OlderThanConfig)
		return models.DeleteOldRepositoryArchives(ctx, acConfig.OlderThan)
	})
}

func registerSyncExternalUsers() {
	RegisterTaskFatal("sync_external_users", &UpdateExistingConfig{
		BaseConfig: BaseConfig{
			Enabled:    true,
			RunAtStart: false,
			Schedule:   "@every 24h",
		},
		UpdateExisting: true,
	}, func(ctx context.Context, _ *models.User, config Config) error {
		realConfig := config.(*UpdateExistingConfig)
		return models.SyncExternalUsers(ctx, realConfig.UpdateExisting)
	})
}

func registerDeletedBranchesCleanup() {
	RegisterTaskFatal("deleted_branches_cleanup", &OlderThanConfig{
		BaseConfig: BaseConfig{
			Enabled:    true,
			RunAtStart: true,
			Schedule:   "@every 24h",
		},
		OlderThan: 24 * time.Hour,
	}, func(ctx context.Context, _ *models.User, config Config) error {
		realConfig := config.(*OlderThanConfig)
		models.RemoveOldDeletedBranches(ctx, realConfig.OlderThan)
		return nil
	})
}

func registerUpdateMigrationPosterID() {
	RegisterTaskFatal("update_migration_poster_id", &BaseConfig{
		Enabled:    true,
		RunAtStart: true,
		Schedule:   "@every 24h",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		return migrations.UpdateMigrationPosterID(ctx)
	})
}

func registerHandleUnDecompressAttachment() {
	RegisterTaskFatal("handle_undecompress_attachment", &BaseConfig{
		Enabled:    true,
		RunAtStart: true,
		Schedule:   "@every 10m",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		repo.HandleUnDecompressAttachment()
		return nil
	})
}

func registerHandleModelSafetyTask() {
	RegisterTaskFatal("handle_modelsafety_task", &BaseConfig{
		Enabled:    true,
		RunAtStart: true,
		Schedule:   "@every 5m",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		repo.TimerHandleModelSafetyTestTask()
		return nil
	})
}

func registerHandleBlockChainUnSuccessUsers() {
	RegisterTaskFatal("handle_blockchain_unsuccess_users", &BaseConfig{
		Enabled:    true,
		RunAtStart: true,
		Schedule:   "@every 10m",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		repo.HandleBlockChainUnSuccessUsers()
		return nil
	})
}

func registerHandleBlockChainUnSuccessRepos() {
	RegisterTaskFatal("handle_blockchain_unsuccess_repos", &BaseConfig{
		Enabled:    true,
		RunAtStart: true,
		Schedule:   "@every 10m",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		repo.HandleBlockChainUnSuccessRepos()
		return nil
	})
}

func registerHandleBlockChainMergedPulls() {
	RegisterTaskFatal("handle_blockchain_merged_pull", &BaseConfig{
		Enabled:    true,
		RunAtStart: true,
		Schedule:   "@every 10m",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		repo.HandleBlockChainMergedPulls()
		return nil
	})
}

func registerHandleBlockChainUnSuccessCommits() {
	RegisterTaskFatal("handle_blockchain_unsuccess_commits", &BaseConfig{
		Enabled:    true,
		RunAtStart: true,
		Schedule:   "@every 10m",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		repo.HandleBlockChainUnSuccessCommits()
		return nil
	})
}

func registerHandleRepoAndUserStatistic() {
	RegisterTaskFatal("handle_repo_and_user_statistic", &BaseConfig{
		Enabled:    true,
		RunAtStart: false,
		Schedule:   "@daily",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		repo.StatisticAuto()
		return nil
	})
}

func registerHandleClearCloudbrainResult() {
	RegisterTaskFatal("handle_cloudbrain_one_result_clear", &BaseConfig{
		Enabled:    true,
		RunAtStart: setting.ClearStrategy.RunAtStart,
		Schedule:   setting.ClearStrategy.Cron,
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		cloudbrainService.ClearCloudbrainResultSpace()
		return nil
	})
}

func registerHandleClearNotebook() {
	RegisterTaskFatal("handle_notebook_clear", &BaseConfig{
		Enabled:    true,
		RunAtStart: setting.NotebookStrategy.RunAtStart,
		Schedule:   setting.NotebookStrategy.Cron,
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		task.ClearNotebook()
		return nil
	})
}

func registerHandleSummaryStatistic() {
	RegisterTaskFatal("handle_summary_statistic", &BaseConfig{
		Enabled:    true,
		RunAtStart: false,
		Schedule:   "@daily",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		repo.SummaryStatistic()
		return nil
	})
}
func registerHandleCardStatistic() {
	RegisterTaskFatal("handle_card_statistic", &BaseConfig{
		Enabled:    setting.CardStatistic.Enabled,
		RunAtStart: setting.CardStatistic.RunAtStart,
		Schedule:   setting.CardStatistic.Cron,
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		cloudbrainService.CardStatistic()
		return nil
	})

}

func registerHandleOrgStatistic() {
	RegisterTaskFatal("handle_org_statistic", &BaseConfig{
		Enabled:    true,
		RunAtStart: false,
		Schedule:   "0 0 2 * * ?",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		models.UpdateOrgStatistics()
		return nil
	})
}

func registerSyncCloudbrainStatus() {
	RegisterTaskFatal("sync_cloudbrain_status", &BaseConfig{
		Enabled:    true,
		RunAtStart: false,
		Schedule:   setting.SyncCloudbrainStatusDuration,
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		repo.SyncCloudbrainStatus()
		return nil
	})
}

func registerSyncCommitImageStatus() {
	RegisterTaskFatal("sync_commit_image_status", &BaseConfig{
		Enabled:    true,
		RunAtStart: false,
		Schedule:   "@every 1m",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		repo.SyncCommitImageStatus()
		return nil
	})
}

func registerHandleScheduleRecord() {
	RegisterTaskFatal("handle_schedule_record", &BaseConfig{
		Enabled:    true,
		RunAtStart: false,
		Schedule:   "@every 1m",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		urchin.HandleScheduleRecords()
		return nil
	})
}

func registerHandleModelMigrateRecord() {
	RegisterTaskFatal("handle_model_migrate_record", &BaseConfig{
		Enabled:    true,
		RunAtStart: false,
		Schedule:   "@every 1m",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		schedule.HandleUnfinishedMigrateRecords()
		return nil
	})
}

func registerHandleNoJobIdAITasks() {
	RegisterTaskFatal("handle_no_job_id_ai_task", &BaseConfig{
		Enabled:    true,
		RunAtStart: false,
		Schedule:   "@every 1m",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		task.HandleNoJobIdAITasks()
		return nil
	})
}

func registerRewardPeriodTask() {
	RegisterTaskFatal("reward_period_task", &BaseConfig{
		Enabled:    true,
		RunAtStart: true,
		Schedule:   "@every 1m",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		reward.StartRewardTask()
		return nil
	})
}

func registerCloudbrainPointDeductTask() {
	RegisterTaskFatal("cloudbrain_point_deduct_task", &BaseConfig{
		Enabled:    true,
		RunAtStart: true,
		Schedule:   "@every 1m",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		reward.StartCloudbrainPointDeductTask()
		return nil
	})
}

func registerSyncResourceSpecs() {
	RegisterTaskFatal("sync_grampus_specs", &BaseConfig{
		Enabled:    true,
		RunAtStart: true,
		Schedule:   "0 0 1 * * ?",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		resource.SyncGrampusQueueAndSpecs()
		return nil
	})
}

func registerSyncModelArtsTempJobs() {
	RegisterTaskFatal("sync_model_arts_temp_jobs", &BaseConfig{
		Enabled:    true,
		RunAtStart: false,
		Schedule:   "@every 1m",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		modelarts.SyncTempStatusJob()
		return nil
	})
}

func registerHandleCloudbrainDurationStatistic() {
	RegisterTaskFatal("handle_cloudbrain_duration_statistic", &BaseConfig{
		Enabled:    true,
		RunAtStart: false,
		Schedule:   "1 1 * * * ?",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		repo.CloudbrainDurationStatisticHour()
		return nil
	})
}

func registerSyncFinetunePanguDeployStatus() {
	RegisterTaskFatal("sync_pangu_deploy_status", &BaseConfig{
		Enabled:    false,
		RunAtStart: false,
		Schedule:   "@every 1m",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		finetune.SyncPanguDeployStatus()
		return nil
	})
}

func registerFinetunePanguServiceCreateQueue() {
	RegisterTaskFatal("pangu_service_create_queue", &BaseConfig{
		Enabled:    false,
		RunAtStart: false,
		Schedule:   "@every 5m",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		finetune.PanguServiceCreateQueue()
		return nil
	})
}

func registerCheckHfHeartbeat() {
	RegisterTaskFatal("hf_file_transfer_heartbeat", &BaseConfig{
		Enabled:    true,
		RunAtStart: true,
		Schedule:   setting.ExternalTransfer.HeartbeatCron,
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		hf_model.CheckHfHeartbeats()
		return nil
	})
}

func registerClearTemporaryAttachment() {
	RegisterTaskFatal("handle_clear_temporary_attachment", &BaseConfig{
		Enabled:    true,
		RunAtStart: true,
		Schedule:   "@daily",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		attachment.ClearTemporaryAttachment()
		return nil
	})
}

func registerProcessSubjectSizeChangedQueue() {
	RegisterTaskFatal("process_subject_size_changed_queue", &BaseConfig{
		Enabled:    true,
		RunAtStart: false,
		Schedule:   "@every 1m",
	}, func(ctx context.Context, _ *models.User, _ Config) error {
		subject_service.ProcessAimodelSizeChangedQueue()
		subject_service.ProcessDatasetSizeChangedQueue()
		return nil
	})
}

func initBasicTasks() {
	registerUpdateMirrorTask()
	registerRepoHealthCheck()
	registerCheckRepoStats()
	registerArchiveCleanup()
	registerSyncExternalUsers()
	registerDeletedBranchesCleanup()
	registerUpdateMigrationPosterID()

	registerHandleUnDecompressAttachment()
	//registerHandleBlockChainUnSuccessUsers()
	//registerHandleBlockChainUnSuccessRepos()
	//registerHandleBlockChainMergedPulls()
	//registerHandleBlockChainUnSuccessCommits()

	registerHandleRepoAndUserStatistic()
	registerHandleSummaryStatistic()
	registerHandleClearCloudbrainResult()
	registerHandleClearNotebook()

	registerSyncCloudbrainStatus()
	registerSyncCommitImageStatus()
	registerHandleOrgStatistic()
	registerSyncResourceSpecs()
	registerSyncModelArtsTempJobs()

	//registerRewardPeriodTask()
	registerCloudbrainPointDeductTask()

	registerHandleModelSafetyTask()

	//registerHandleScheduleRecord()
	registerHandleCloudbrainDurationStatistic()

	registerHandleModelMigrateRecord()

	registerHandleNoJobIdAITasks()

	registerSyncFinetunePanguDeployStatus()
	registerFinetunePanguServiceCreateQueue()
	registerHandleCardStatistic()

	registerCheckHfHeartbeat()
	registerClearTemporaryAttachment()

	registerProcessSubjectSizeChangedQueue()

}
