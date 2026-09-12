package repo

import (
	"errors"
	"sync"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/normalization"
	"code.gitea.io/gitea/modules/repository"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/services/mailer"

	"gitea.com/macaron/macaron"
	"github.com/Jeffail/tunny"
)

var statisticMutex sync.Mutex
var statisticWg sync.WaitGroup

type StatisticInput struct {
	Date string
	Repo *models.Repository
	T    time.Time
}

func StatisticAuto() {
	go RepoStatisticAuto()
	go TimingCountData()
}

// auto daily
func RepoStatisticAuto() {
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	setting.UpdateRadarMap()
	RepoStatisticDaily(yesterday)
}

func RepoStatisticDaily(date string) {
	log.Info("%s", date)
	log.Info("begin Repo Statistic")
	t, _ := time.Parse("2006-01-02", date)
	warnEmailMessage := "项目统计信息入库失败，请尽快定位。"
	if err := models.DeleteRepoStatDaily(date); err != nil {
		log.Error("DeleteRepoStatDaily failed: %v", err.Error())
		mailer.SendWarnNotifyMail(setting.Warn_Notify_Mails, warnEmailMessage)
		return
	}

	repos, err := models.GetAllRepositories()
	if err != nil {
		log.Error("GetAllRepositories failed: %v", err.Error())
		mailer.SendWarnNotifyMail(setting.Warn_Notify_Mails, warnEmailMessage)
		return
	}

	var reposRadar = make([]*models.RepoStatistic, 0)

	var minRepoRadar models.RepoStatistic
	var maxRepoRadar models.RepoStatistic

	isInitMinMaxRadar := false

	var error_projects = make([]string, 0)

	pool := tunny.NewFunc(8, func(input interface{}) interface{} {
		defer statisticWg.Done()

		statisticInput := input.(*StatisticInput)
		repo := statisticInput.Repo
		projectName := getDistinctProjectName(repo)
		log.Info("start statistic: %s", projectName)
		var numDevMonths, numWikiViews, numContributor, numKeyContributor, numCommitsGrowth, numCommitLinesGrowth, numContributorsGrowth, numCommits int64
		dayBeforeYestoday := statisticInput.T.AddDate(0, 0, -1)

		repoStatisticBeforeYestoday, err := models.GetRepoStatisticByDateAndRepoId(dayBeforeYestoday.Format("2006-01-02"), repo.ID)
		if err == nil && repo.UpdatedUnix <= timeutil.TimeStamp(dayBeforeYestoday.Unix()) && isGitItemGrowthIsZero(repoStatisticBeforeYestoday) {
			numDevMonths = repoStatisticBeforeYestoday.NumDevMonths
			numKeyContributor = repoStatisticBeforeYestoday.NumKeyContributor
			numWikiViews = repoStatisticBeforeYestoday.NumWikiViews
			numContributor = repoStatisticBeforeYestoday.NumContributor
			numCommitsGrowth = repoStatisticBeforeYestoday.NumCommitsGrowth
			numCommitLinesGrowth = repoStatisticBeforeYestoday.NumCommitLinesGrowth
			numContributorsGrowth = repoStatisticBeforeYestoday.NumContributorsGrowth
			numCommits = repoStatisticBeforeYestoday.NumCommits
		} else {

			repoGitStat, err := models.GetRepoKPIStats(repo)

			if err != nil {
				log.Error("GetRepoKPIStats failed: %s", projectName)
			} else {
				numDevMonths = repoGitStat.DevelopAge
				numKeyContributor = repoGitStat.KeyContributors
				numWikiViews = repoGitStat.WikiPages
				numContributor = repoGitStat.Contributors
				numCommitsGrowth = repoGitStat.CommitsAdded
				numCommitLinesGrowth = repoGitStat.CommitLinesModified
				numContributorsGrowth = repoGitStat.ContributorsAdded
				numCommits = repoGitStat.Commits

			}
		}

		var issueFixedRate float32
		if repo.NumIssues != 0 {
			issueFixedRate = float32(repo.NumClosedIssues) / float32(repo.NumIssues)
		} else {
			issueFixedRate = float32(setting.RadarMap.ProjectHealth0IssueCloseRatio)
		}

		var numVersions int64
		numVersions, err = models.GetReleaseCountByRepoID(repo.ID, models.FindReleasesOptions{})
		if err != nil {
			log.Error("GetReleaseCountByRepoID failed(%s): %v", projectName, err)
		}

		var datasetSize int64
		datasetSize, err = getDatasetSize(repo)
		if err != nil {
			log.Error("getDatasetSize failed(%s): %v", projectName, err)
		}

		var numComments int64
		numComments, err = models.GetCommentCountByRepoID(repo.ID)
		if err != nil {
			log.Error("GetCommentCountByRepoID failed(%s): %v", projectName, err)
		}

		beginTime, endTime := getStatTime(statisticInput.Date)
		var numVisits int
		numVisits, err = repository.AppointProjectView(repo.OwnerName, repo.Name, beginTime, endTime)
		if err != nil {
			log.Error("AppointProjectView failed(%s): %v", projectName, err)
		}

		repoStat := models.RepoStatistic{
			RepoID:                repo.ID,
			Date:                  statisticInput.Date,
			Name:                  repo.Name,
			Alias:                 repo.Alias,
			IsPrivate:             repo.IsPrivate,
			IsMirror:              repo.IsMirror,
			IsFork:                repo.IsFork,
			RepoCreatedUnix:       repo.CreatedUnix,
			OwnerName:             repo.OwnerName,
			NumWatches:            int64(repo.NumWatches),
			NumStars:              int64(repo.NumStars),
			NumForks:              int64(repo.NumForks),
			NumDownloads:          repo.CloneCnt,
			NumComments:           numComments,
			NumVisits:             int64(numVisits),
			NumClosedIssues:       int64(repo.NumClosedIssues),
			NumVersions:           numVersions,
			NumDevMonths:          numDevMonths,
			RepoSize:              repo.Size,
			DatasetSize:           datasetSize,
			NumModels:             repo.ModelCnt,
			NumWikiViews:          numWikiViews,
			NumCommits:            numCommits,
			NumIssues:             int64(repo.NumIssues),
			NumPulls:              int64(repo.NumPulls),
			IssueFixedRate:        issueFixedRate,
			NumContributor:        numContributor,
			NumKeyContributor:     numKeyContributor,
			NumCommitsGrowth:      numCommitsGrowth,
			NumCommitLinesGrowth:  numCommitLinesGrowth,
			NumContributorsGrowth: numContributorsGrowth,
			NumCloudbrain:         repo.AiTaskCnt,
			NumDatasetFile:        repo.DatasetCnt,
			NumModelConvert:       models.QueryModelConvertCountByRepoID(repo.ID),
		}

		dayBeforeDate := statisticInput.T.AddDate(0, 0, -1).Format("2006-01-02")
		repoStatisticsBefore, err := models.GetRepoStatisticByDate(dayBeforeDate, repo.ID)

		if err != nil {
			log.Error("get data of day before the date failed ", err)
		} else {
			if len(repoStatisticsBefore) > 0 {
				repoStatisticBefore := repoStatisticsBefore[0]
				repoStat.NumWatchesAdded = repoStat.NumWatches - repoStatisticBefore.NumWatches
				repoStat.NumStarsAdded = repoStat.NumStars - repoStatisticBefore.NumStars
				repoStat.NumForksAdded = repoStat.NumForks - repoStatisticBefore.NumForks
				repoStat.NumDownloadsAdded = repoStat.NumDownloads - repoStatisticBefore.NumDownloads
				repoStat.NumCommentsAdded = repoStat.NumComments - repoStatisticBefore.NumComments
				repoStat.NumClosedIssuesAdded = repoStat.NumClosedIssues - repoStatisticBefore.NumClosedIssues
				repoStat.NumCommitsAdded = repoStat.NumCommits - repoStatisticBefore.NumCommits
				repoStat.NumIssuesAdded = repoStat.NumIssues - repoStatisticBefore.NumIssues
				repoStat.NumPullsAdded = repoStat.NumPulls - repoStatisticBefore.NumPulls
				repoStat.NumContributorAdded = repoStat.NumContributor - repoStatisticBefore.NumContributor
				repoStat.NumModelsAdded = repoStat.NumModels - repoStatisticBefore.NumModels
				repoStat.NumCloudbrainAdded = repoStat.NumCloudbrain - repoStatisticBefore.NumCloudbrain
				repoStat.NumModelConvertAdded = repoStat.NumModelConvert - repoStatisticBefore.NumModelConvert
				repoStat.NumDatasetFileAdded = repoStat.NumDatasetFile - repoStatisticBefore.NumDatasetFile
			} else if (t.Year() == time.Unix(int64(repo.CreatedUnix), 0).Year() && t.Month() == time.Unix(int64(repo.CreatedUnix), 0).Month() && t.Day() == time.Unix(int64(repo.CreatedUnix), 0).Day()) || (!repo.IsFork && !repo.IsMirror) {
				repoStat.NumWatchesAdded = repoStat.NumWatches
				repoStat.NumStarsAdded = repoStat.NumStars
				repoStat.NumForksAdded = repoStat.NumForks
				repoStat.NumDownloadsAdded = repoStat.NumDownloads
				repoStat.NumCommentsAdded = repoStat.NumComments
				repoStat.NumClosedIssuesAdded = repoStat.NumClosedIssues
				repoStat.NumCommitsAdded = repoStat.NumCommits
				repoStat.NumIssuesAdded = repoStat.NumIssues
				repoStat.NumPullsAdded = repoStat.NumPulls
				repoStat.NumContributorAdded = repoStat.NumContributor
				repoStat.NumModelsAdded = repoStat.NumModels
				repoStat.NumCloudbrainAdded = repoStat.NumCloudbrain
				repoStat.NumModelConvertAdded = repoStat.NumModelConvert
				repoStat.NumDatasetFileAdded = repoStat.NumDatasetFile

			}
		}
		day4MonthsAgo := statisticInput.T.AddDate(0, -4, 0)
		repoStatisticFourMonthsAgo, err := models.GetOneRepoStatisticBeforeTime(day4MonthsAgo)
		if err != nil {
			log.Error("Get data of 4 moth ago failed.", err)
		} else {
			repoStat.NumCommentsGrowth = repoStat.NumComments - repoStatisticFourMonthsAgo.NumComments
			repoStat.NumIssuesGrowth = repoStat.NumIssues - repoStatisticFourMonthsAgo.NumIssues
		}

		models.SyncStatDataToRepo(repo)

		if _, err = models.InsertRepoStat(&repoStat); err != nil {
			log.Error("InsertRepoStat failed(%s): %v", projectName, err)
			log.Error("failed statistic: %s", projectName)
			statisticMutex.Lock()
			{
				error_projects = append(error_projects, projectName)
			}
			statisticMutex.Unlock()
			return nil
		}

		statisticMutex.Lock()
		{

			tempRepoStat := models.RepoStatistic{
				RepoID:        repoStat.RepoID,
				Date:          repoStat.Date,
				IsMirror:      repoStat.IsMirror,
				Impact:        normalization.GetImpactInitValue(repoStat.NumWatches, repoStat.NumStars, repoStat.NumForks, repoStat.NumDownloads, repoStat.NumComments, repoStat.NumVisits),
				Completeness:  normalization.GetCompleteInitValue(repoStat.NumClosedIssues, repoStat.NumVersions, repoStat.NumDevMonths, repoStat.DatasetSize, repoStat.NumModels, repoStat.NumWikiViews),
				Liveness:      normalization.GetLivenessInitValue(repoStat.NumCommits, repoStat.NumIssues, repoStat.NumPulls, repoStat.NumVisits),
				ProjectHealth: normalization.GetProjectHealthInitValue(repoStat.IssueFixedRate),
				TeamHealth:    normalization.GetTeamHealthInitValue(repoStat.NumContributor, repoStat.NumKeyContributor, repoStat.NumContributorsGrowth),
				Growth:        normalization.GetRepoGrowthInitValue(repoStat.NumCommitLinesGrowth, repoStat.NumIssuesGrowth, repoStat.NumCommitsGrowth, repoStat.NumContributorsGrowth, repoStat.NumCommentsGrowth),
			}

			reposRadar = append(reposRadar, &tempRepoStat)

			if !isInitMinMaxRadar {

				if !setting.RadarMap.IgnoreMirrorRepo || (setting.RadarMap.IgnoreMirrorRepo && !tempRepoStat.IsMirror) {
					minRepoRadar = tempRepoStat
					maxRepoRadar = tempRepoStat
					isInitMinMaxRadar = true
				}

			} else {
				if !setting.RadarMap.IgnoreMirrorRepo || (setting.RadarMap.IgnoreMirrorRepo && !tempRepoStat.IsMirror) {
					if tempRepoStat.Impact < minRepoRadar.Impact {
						minRepoRadar.Impact = tempRepoStat.Impact
					}

					if tempRepoStat.Impact > maxRepoRadar.Impact {
						maxRepoRadar.Impact = tempRepoStat.Impact
					}

					if tempRepoStat.Completeness < minRepoRadar.Completeness {
						minRepoRadar.Completeness = tempRepoStat.Completeness
					}

					if tempRepoStat.Completeness > maxRepoRadar.Completeness {
						maxRepoRadar.Completeness = tempRepoStat.Completeness
					}

					if tempRepoStat.Liveness < minRepoRadar.Liveness {
						minRepoRadar.Liveness = tempRepoStat.Liveness
					}

					if tempRepoStat.Liveness > maxRepoRadar.Liveness {
						maxRepoRadar.Liveness = tempRepoStat.Liveness
					}

					if tempRepoStat.ProjectHealth < minRepoRadar.ProjectHealth {
						minRepoRadar.ProjectHealth = tempRepoStat.ProjectHealth
					}

					if tempRepoStat.ProjectHealth > maxRepoRadar.ProjectHealth {
						maxRepoRadar.ProjectHealth = tempRepoStat.ProjectHealth
					}

					if tempRepoStat.TeamHealth < minRepoRadar.TeamHealth {
						minRepoRadar.TeamHealth = tempRepoStat.TeamHealth
					}

					if tempRepoStat.TeamHealth > maxRepoRadar.TeamHealth {
						maxRepoRadar.TeamHealth = tempRepoStat.TeamHealth
					}

					if tempRepoStat.Growth < minRepoRadar.Growth {
						minRepoRadar.Growth = tempRepoStat.Growth
					}

					if tempRepoStat.Growth > maxRepoRadar.Growth {
						maxRepoRadar.Growth = tempRepoStat.Growth
					}

				}

			}
		}
		statisticMutex.Unlock()

		log.Info("finish statistic: %s", getDistinctProjectName(repo))
		return nil
	})
	defer pool.Close()
	statisticWg.Add(len(repos))
	for _, repo := range repos {
		input := &StatisticInput{
			Date: date,
			Repo: repo,
			T:    t,
		}
		pool.Process(input)
	}
	statisticWg.Wait()

	if len(error_projects) > 0 {
		mailer.SendWarnNotifyMail(setting.Warn_Notify_Mails, warnEmailMessage)
	}

	//radar  map
	log.Info("begin statistic radar")
	for _, radarInit := range reposRadar {
		if radarInit.IsMirror && setting.RadarMap.IgnoreMirrorRepo {
			radarInit.Impact = 0
			radarInit.Completeness = 0
			radarInit.Liveness = 0
			radarInit.ProjectHealth = 0
			radarInit.TeamHealth = 0
			radarInit.Growth = 0
			radarInit.RadarTotal = 0
		} else {
			radarInit.Impact = normalization.Normalization(radarInit.Impact, minRepoRadar.Impact, maxRepoRadar.Impact)
			radarInit.Completeness = normalization.Normalization(radarInit.Completeness, minRepoRadar.Completeness, maxRepoRadar.Completeness)
			radarInit.Liveness = normalization.Normalization(radarInit.Liveness, minRepoRadar.Liveness, maxRepoRadar.Liveness)
			radarInit.ProjectHealth = normalization.Normalization(radarInit.ProjectHealth, minRepoRadar.ProjectHealth, maxRepoRadar.ProjectHealth)
			radarInit.TeamHealth = normalization.Normalization(radarInit.TeamHealth, minRepoRadar.TeamHealth, maxRepoRadar.TeamHealth)
			radarInit.Growth = normalization.Normalization(radarInit.Growth, minRepoRadar.Growth, maxRepoRadar.Growth)
			radarInit.RadarTotal = normalization.GetRadarValue(radarInit.Impact, radarInit.Completeness, radarInit.Liveness, radarInit.ProjectHealth, radarInit.TeamHealth, radarInit.Growth)
		}
		models.UpdateRepoStat(radarInit)
	}

	log.Info("finish statistic: radar")

}

func isGitItemGrowthIsZero(repoStatisticBeforeYestoday *models.RepoStatistic) bool {

	return repoStatisticBeforeYestoday.NumCommitsGrowth == 0 &&
		repoStatisticBeforeYestoday.NumCommitLinesGrowth == 0 &&
		repoStatisticBeforeYestoday.NumContributorsGrowth == 0
}

func getDistinctProjectName(repo *models.Repository) string {
	return repo.OwnerName + "/" + repo.Alias
}

func getDatasetSize(repo *models.Repository) (int64, error) {
	dataset, err := models.GetDatasetByRepo(repo)
	if err != nil {
		return 0, err
	}

	return models.GetAttachmentSizeByDatasetID(dataset.ID)
}

func getStatTime(timeStr string) (string, string) {
	t, _ := time.Parse("2006-01-02", timeStr)
	timeNumber := t.Unix()
	beginTimeNumber := timeNumber - 8*60*60
	endTimeNumber := timeNumber + 16*60*60
	beginTime := time.Unix(beginTimeNumber, 0).Format(time.RFC3339)
	endTime := time.Unix(endTimeNumber, 0).Format(time.RFC3339)
	log.Info("%s, %s", beginTime, endTime)

	return beginTime, endTime
}

func UpdateRepoVisits(ctx *macaron.Context, repo *models.Repository, date string) error {
	beginTime, endTime := getStatTime(date)
	var numVisits int
	numVisits, err := repository.AppointProjectView(repo.OwnerName, repo.Name, beginTime, endTime)
	if err != nil {
		log.Error("AppointProjectView failed(%s): %v", getDistinctProjectName(repo), err)
		return err
	}

	repoStat, err := models.GetRepoStatisticByDate(date, repo.ID)
	if err != nil {
		log.Error("GetRepoStatisticByDate failed(%s): %v", getDistinctProjectName(repo), err)
		return err
	} else if len(repoStat) != 1 {
		log.Error("GetRepoStatisticByDate failed(%s): %v", getDistinctProjectName(repo), err)
		return errors.New("not find repo")
	}
	repoStat[0].NumVisits = int64(numVisits)

	if err = models.UpdateRepoStatVisits(repoStat[0]); err != nil {
		log.Error("UpdateRepoStatVisits failed(%s): %v", getDistinctProjectName(repo), err)
		return err
	}
	return nil
}
