package models

import (
	"fmt"
	"time"

	"code.gitea.io/gitea/modules/timeutil"
)

// RepoStatistic statistic info of all repository
type RepoStatistic struct {
	ID                   int64              `xorm:"pk autoincr" json:"-"`
	RepoID               int64              `xorm:"unique(s) NOT NULL"  json:"repo_id"`
	Name                 string             `xorm:"INDEX"  json:"name"`
	Alias                string             `xorm:"INDEX"  json:"alias"`
	OwnerName            string             `json:"ownerName"`
	IsPrivate            bool               `json:"isPrivate"`
	IsMirror             bool               `json:"isMirror"`
	IsFork               bool               `json:"isFork"`
	RepoCreatedUnix      timeutil.TimeStamp `xorm:"NOT NULL DEFAULT 0" json:"createUnix"`
	Date                 string             `xorm:"unique(s) NOT NULL" json:"date"`
	NumWatches           int64              `xorm:"NOT NULL DEFAULT 0" json:"watch"`
	NumWatchesAdded      int64              `xorm:"NOT NULL DEFAULT 0" json:"-"`
	NumStars             int64              `xorm:"NOT NULL DEFAULT 0" json:"star"`
	NumStarsAdded        int64              `xorm:"NOT NULL DEFAULT 0" json:"-"`
	NumForks             int64              `xorm:"NOT NULL DEFAULT 0" json:"fork"`
	NumForksAdded        int64              `xorm:"NOT NULL DEFAULT 0" json:"-"`
	NumDownloads         int64              `xorm:"NOT NULL DEFAULT 0" json:"download"`
	NumDownloadsAdded    int64              `xorm:"NOT NULL DEFAULT 0" json:"-"`
	NumComments          int64              `xorm:"NOT NULL DEFAULT 0" json:"comment"`
	NumCommentsAdded     int64              `xorm:"NOT NULL DEFAULT 0" json:"-"`
	NumVisits            int64              `xorm:"NOT NULL DEFAULT 0" json:"view"`
	NumClosedIssues      int64              `xorm:"NOT NULL DEFAULT 0" json:"issueClosed"`
	NumClosedIssuesAdded int64              `xorm:"NOT NULL DEFAULT 0" json:"-"`
	NumVersions          int64              `xorm:"NOT NULL DEFAULT 0" json:"-"`
	NumDevMonths         int64              `xorm:"NOT NULL DEFAULT 0" json:"-"`
	RepoSize             int64              `xorm:"NOT NULL DEFAULT 0" json:"-"`
	DatasetSize          int64              `xorm:"NOT NULL DEFAULT 0" json:"-"`
	NumModels            int64              `xorm:"NOT NULL DEFAULT 0" json:"model"`
	NumWikiViews         int64              `xorm:"NOT NULL DEFAULT 0" json:"-"`
	NumCommits           int64              `xorm:"NOT NULL DEFAULT 0" json:"commit"`
	NumCommitsAdded      int64              `xorm:"NOT NULL DEFAULT 0" json:"-"`
	NumIssues            int64              `xorm:"NOT NULL DEFAULT 0" json:"issue"`
	NumIssuesAdded       int64              `xorm:"NOT NULL DEFAULT 0" json:"-"`
	NumPulls             int64              `xorm:"NOT NULL DEFAULT 0" json:"pr"`
	NumPullsAdded        int64              `xorm:"NOT NULL DEFAULT 0" json:"-"`
	IssueFixedRate       float32            `xorm:"NOT NULL" json:"issueClosedRatio"`
	NumContributor       int64              `xorm:"NOT NULL DEFAULT 0" json:"contributor"`
	NumContributorAdded  int64              `xorm:"NOT NULL DEFAULT 0" json:"-"`
	NumKeyContributor    int64              `xorm:"NOT NULL DEFAULT 0" json:"-"`

	NumContributorsGrowth int64 `xorm:"NOT NULL DEFAULT 0" json:"-"`
	NumCommitsGrowth      int64 `xorm:"NOT NULL DEFAULT 0" json:"-"`
	NumCommitLinesGrowth  int64 `xorm:"NOT NULL DEFAULT 0" json:"-"`
	NumIssuesGrowth       int64 `xorm:"NOT NULL DEFAULT 0" json:"-"`
	NumCommentsGrowth     int64 `xorm:"NOT NULL DEFAULT 0" json:"-"`

	NumDatasetFile  int64 `xorm:"NOT NULL DEFAULT 0" json:"datasetFiles"`
	NumCloudbrain   int64 `xorm:"NOT NULL DEFAULT 0" json:"cloudbrains"`
	NumModelConvert int64 `xorm:"NOT NULL DEFAULT 0" json:"modelConverts"`

	NumDatasetFileAdded  int64 `xorm:"NOT NULL DEFAULT 0" json:"-"`
	NumCloudbrainAdded   int64 `xorm:"NOT NULL DEFAULT 0" json:"-"`
	NumModelConvertAdded int64 `xorm:"NOT NULL DEFAULT 0" json:"- "`
	NumModelsAdded       int64 `xorm:"NOT NULL DEFAULT 0" json:"- "`

	Impact        float64            `xorm:"NOT NULL DEFAULT 0" json:"impact"`
	Completeness  float64            `xorm:"NOT NULL DEFAULT 0" json:"completeness"`
	Liveness      float64            `xorm:"NOT NULL DEFAULT 0" json:"liveness"`
	ProjectHealth float64            `xorm:"NOT NULL DEFAULT 0" json:"projectHealth"`
	TeamHealth    float64            `xorm:"NOT NULL DEFAULT 0" json:"teamHealth"`
	Growth        float64            `xorm:"NOT NULL DEFAULT 0" json:"growth"`
	RadarTotal    float64            `xorm:"NOT NULL DEFAULT 0" json:"openi"`
	CreatedUnix   timeutil.TimeStamp `xorm:"INDEX created" json:"-"`
	UpdatedUnix   timeutil.TimeStamp `xorm:"INDEX updated" json:"-"`
}

func (repo *RepoStatistic) DisplayName() string {
	if repo.Alias == "" {
		return repo.Name
	}
	return repo.Alias
}

func getOpenIByRepoId(repoId int64) float64 {
	repoStatistic := new(RepoStatistic)
	has, err := xStatistic.Cols("radar_total").Where("repo_id=?", repoId).Desc("id").Limit(1).Get(repoStatistic)
	if !has || err != nil {
		return 0
	}
	return repoStatistic.RadarTotal

}

func DeleteRepoStatDaily(date string) error {
	sess := xStatistic.NewSession()
	defer sess.Close()
	if err := sess.Begin(); err != nil {
		return fmt.Errorf("Begin: %v", err)
	}

	if _, err := sess.Where("date = ?", date).Delete(&RepoStatistic{}); err != nil {
		return fmt.Errorf("Delete: %v", err)
	}

	if err := sess.Commit(); err != nil {
		sess.Close()
		return fmt.Errorf("Commit: %v", err)
	}

	sess.Close()
	return nil
}

func CountRepoStatByRawSql(sql string) (int64, error) {

	return xStatistic.SQL(sql).Count()

}

func GetRepoStatisticByRawSql(sql string) []*RepoStatistic {
	repoStatistics := make([]*RepoStatistic, 0)
	xStatistic.SQL(sql).Find(&repoStatistics)
	return repoStatistics
}

func GetRepoStatLastUpdatedTime(repoId ...string) (string, string, error) {

	repoStatistic := new(RepoStatistic)
	var has bool
	var err error
	if len(repoId) == 0 {
		has, err = xStatistic.Desc("created_unix").Limit(1).Cols("created_unix", "date").Get(repoStatistic)

	} else {
		has, err = xStatistic.Where("repo_id=?", repoId[0]).Desc("created_unix").Limit(1).Cols("created_unix", "date").Get(repoStatistic)
	}

	if err != nil {
		return "", "", err
	} else {
		if has {
			return repoStatistic.CreatedUnix.Format("2006-01-02 15:04:05"), repoStatistic.Date, nil
		} else {
			return "", "", fmt.Errorf("Can not get the latest record.")
		}
	}

}

func GetRepoStatisticByDateAndRepoId(date string, repoId int64) (*RepoStatistic, error) {
	repoStatistic := new(RepoStatistic)
	has, err := xStatistic.Where("date=? and repo_id=?", date, repoId).Get(repoStatistic)
	if err != nil {
		return nil, err
	} else {
		if has {
			return repoStatistic, nil
		} else {
			return nil, fmt.Errorf("The num of return records is 0.")
		}
	}

}

func GetRepoStatisticByDate(date string, repoId int64) ([]*RepoStatistic, error) {
	repoStatistics := make([]*RepoStatistic, 0)
	err := xStatistic.Where("date = ? and repo_id=?", date, repoId).Find(&repoStatistics)
	return repoStatistics, err

}

func GetOneRepoStatisticBeforeTime(time time.Time) (*RepoStatistic, error) {
	repoStatistics := make([]*RepoStatistic, 0)
	err := xStatistic.Where("created_unix >= ?", time.Unix()).OrderBy("created_unix").Limit(1).Find(&repoStatistics)
	if err != nil {
		return nil, err
	} else {
		if len(repoStatistics) == 0 {
			return nil, fmt.Errorf("the repo statistic record count is 0")
		} else {
			return repoStatistics[0], nil
		}
	}

}

func InsertRepoStat(repoStat *RepoStatistic) (int64, error) {
	return xStatistic.Insert(repoStat)
}

func RestoreRepoStatFork(numForks int64, repoId int64) error {
	sql := "update repo_statistic set num_forks=?  where repo_id=?"

	_, err := xStatistic.Exec(sql, numForks, repoId)
	return err
}

func UpdateRepoStat(repoStat *RepoStatistic) error {
	sql := "update repo_statistic set impact=?,completeness=?,liveness=?,project_health=?,team_health=?,growth=?,radar_total=?  where repo_id=? and date=?"

	_, err := xStatistic.Exec(sql, repoStat.Impact, repoStat.Completeness, repoStat.Liveness, repoStat.ProjectHealth, repoStat.TeamHealth, repoStat.Growth, repoStat.RadarTotal, repoStat.RepoID, repoStat.Date)
	return err
}

func UpdateRepoStatVisits(repoStat *RepoStatistic) error {
	sql := "update repo_statistic set num_visits=? where repo_id=? and date=?"

	_, err := xStatistic.Exec(sql, repoStat.NumVisits, repoStat.RepoID, repoStat.Date)
	return err
}

func SumRepoStatColumn(begin, end time.Time, repoId int64, columnName string) (int64, error) {
	res, err := xStatistic.Where("created_unix <= ? and created_unix >= ? and repo_id = ? ", end.Unix(), begin.Unix(), repoId).Sum(&RepoStatistic{}, columnName)
	if err != nil {
		return 0, err
	}
	return int64(res), nil
}

func SumLastMonthNumVisits(repoId int64) (int64, error) {
	end := time.Now()
	begin := end.AddDate(0, 0, -30)
	return SumRepoStatColumn(begin, end, repoId, "num_visits")
}

func SumLastFourMonthNumCommits(repoId int64) (int64, error) {
	end := time.Now()
	begin := end.AddDate(0, 0, -120)
	return SumRepoStatColumn(begin, end, repoId, "num_commits_added")
}
