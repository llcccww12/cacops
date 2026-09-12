package models

import (
	"fmt"
	"strconv"
	"time"

	"code.gitea.io/gitea/modules/timeutil"
)

var DomainMap = map[string]int{
	"大模型":      0,
	"ai开发工具":   1,
	"计算机视觉":    2,
	"自然语言处理":   3,
	"机器学习":     4,
	"神经网络":     5,
	"自动驾驶":     6,
	"机器人":      7,
	"联邦学习":     8,
	"数据挖掘":     9,
	"risc-v开发": 10,
}

type SummaryStatistic struct {
	ID                 int64              `xorm:"pk autoincr"`
	Date               string             `xorm:"unique(s) NOT NULL"`
	NumUsers           int64              `xorm:"NOT NULL DEFAULT 0"`
	RepoSize           int64              `xorm:"NOT NULL DEFAULT 0"`
	DatasetSize        int64              `xorm:"NOT NULL DEFAULT 0"`
	NumOrganizations   int64              `xorm:"NOT NULL DEFAULT 0"`
	NumModels          int64              `xorm:"NOT NULL DEFAULT 0"`
	NumRepos           int64              `xorm:"NOT NULL DEFAULT 0"`
	NumRepoBigModel    int                `xorm:"NOT NULL DEFAULT 0"`
	NumRepoAI          int                `xorm:"NOT NULL DEFAULT 0"`
	NumRepoVision      int                `xorm:"NOT NULL DEFAULT 0"`
	NumRepoNLP         int                `xorm:"NOT NULL DEFAULT 0"`
	NumRepoML          int                `xorm:"NOT NULL DEFAULT 0"`
	NumRepoNN          int                `xorm:"NOT NULL DEFAULT 0"`
	NumRepoAutoDrive   int                `xorm:"NOT NULL DEFAULT 0"`
	NumRepoRobot       int                `xorm:"NOT NULL DEFAULT 0"`
	NumRepoLeagueLearn int                `xorm:"NOT NULL DEFAULT 0"`
	NumRepoDataMining  int                `xorm:"NOT NULL DEFAULT 0"`
	NumRepoRISC        int                `xorm:"NOT NULL DEFAULT 0"`
	NumRepoPublic      int64              `xorm:"NOT NULL DEFAULT 0"`
	NumRepoPrivate     int64              `xorm:"NOT NULL DEFAULT 0"`
	NumRepoFork        int64              `xorm:"NOT NULL DEFAULT 0"`
	NumRepoMirror      int64              `xorm:"NOT NULL DEFAULT 0"`
	NumRepoSelf        int64              `xorm:"NOT NULL DEFAULT 0"`
	NumRepoOrg         int64              `xorm:"NOT NULL DEFAULT 0"`
	CreatedUnix        timeutil.TimeStamp `xorm:"INDEX created"`
	UpdatedUnix        timeutil.TimeStamp `xorm:"INDEX updated"`
}

func DeleteSummaryStatisticDaily(date string) error {
	sess := xStatistic.NewSession()
	defer sess.Close()
	if err := sess.Begin(); err != nil {
		return fmt.Errorf("Begin: %v", err)
	}

	if _, err := sess.Where("date = ?", date).Delete(&SummaryStatistic{}); err != nil {
		return fmt.Errorf("Delete: %v", err)
	}

	if err := sess.Commit(); err != nil {
		sess.Close()
		return fmt.Errorf("Commit: %v", err)
	}

	sess.Close()
	return nil
}

func GetLatest2SummaryStatistic() ([]*SummaryStatistic, error) {
	summaryStatistics := make([]*SummaryStatistic, 0)
	err := xStatistic.Desc("created_unix").Limit(2).Find(&summaryStatistics)
	return summaryStatistics, err
}

func GetSummaryStatisticByTimeCount(beginTime time.Time, endTime time.Time) (int64, error) {
	summaryStatistics := new(SummaryStatistic)
	total, err := xStatistic.Asc("created_unix").Where("created_unix>=" + strconv.FormatInt(beginTime.Unix(), 10) + " and created_unix<" + strconv.FormatInt(endTime.Unix(), 10)).Count(summaryStatistics)
	return total, err
}

func GetSummaryStatisticByDateCount(dates []string) (int64, error) {
	summaryStatistics := new(SummaryStatistic)
	total, err := xStatistic.Asc("created_unix").In("date", dates).Count(summaryStatistics)
	return total, err
}


func GetAllSummaryStatisticByTime(beginTime time.Time, endTime time.Time) ([]*SummaryStatistic, error) {
	summaryStatistics := make([]*SummaryStatistic, 0)
	err := xStatistic.Asc("created_unix").Where("created_unix>=" + strconv.FormatInt(beginTime.Unix(), 10) + " and created_unix<" + strconv.FormatInt(endTime.Unix(), 10)).Find(&summaryStatistics)

	return summaryStatistics, err
}

func GetSummaryStatisticByTime(beginTime time.Time, endTime time.Time, page int, pageSize int) ([]*SummaryStatistic, error) {
	summaryStatistics := make([]*SummaryStatistic, 0)
	err := xStatistic.Asc("created_unix").Limit(pageSize+1, (page-1)*pageSize).Where("created_unix>=" + strconv.FormatInt(beginTime.Unix(), 10) + " and created_unix<" + strconv.FormatInt(endTime.Unix(), 10)).Find(&summaryStatistics)

	return summaryStatistics, err
}

func GetAllSummaryStatisticByDates(dates []string) ([]*SummaryStatistic, error) {
	summaryStatistics := make([]*SummaryStatistic, 0)
	err := xStatistic.Asc("created_unix").In("date", dates).Find(&summaryStatistics)
	return summaryStatistics, err
}

func GetSummaryStatisticByDates(dates []string, page int, pageSize int) ([]*SummaryStatistic, error) {
	summaryStatistics := make([]*SummaryStatistic, 0)
	err := xStatistic.Asc("created_unix").In("date", dates).Limit(pageSize+1, (page-1)*pageSize).Find(&summaryStatistics)
	return summaryStatistics, err
}

func InsertSummaryStatistic(summaryStatistic *SummaryStatistic) (int64, error) {
	return xStatistic.Insert(summaryStatistic)
}
