package models

import (
	"strings"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
	"xorm.io/xorm"
)

type SearchRecord struct {
	ID int64 `xorm:"pk autoincr"`
	//user
	Keyword string `xorm:"NOT NULL"`
	//
	CreatedUnix timeutil.TimeStamp `xorm:"INDEX created"`
}

func SaveSearchKeywordToDb(keyword string) error {
	record := &SearchRecord{
		Keyword: keyword,
	}
	sess := x.NewSession()
	defer sess.Close()
	_, err := sess.Insert(record)
	if err != nil {
		log.Info("insert error." + err.Error())
		return err
	}
	return nil
}

func setIssueQueryCondition(sess *xorm.Session, Keyword string, isPull bool, userId int64) {
	sess.And("issue.poster_id=?", userId)
	sess.And("issue.is_pull=?", isPull)
	sess.And("(issue.name like '%" + Keyword + "%' or issue.content like '%" + Keyword + "%')")
	sess.Join("INNER", "repository", "issue.repo_id = repository.id").And("repository.is_private = ?", true)
}

func SearchPrivateIssueOrPr(Page int, PageSize int, Keyword string, isPull bool, userId int64) ([]*Issue, int64, error) {
	sess := x.NewSession()
	defer sess.Close()
	setIssueQueryCondition(sess, Keyword, isPull, userId)
	count, err := sess.Count(new(Issue))
	if err != nil {
		return nil, 0, err
	}

	setIssueQueryCondition(sess, Keyword, isPull, userId)
	sess.Desc("issue.created_unix")
	sess.Limit(PageSize, (Page-1)*PageSize)
	issues := make([]*Issue, 0)
	if err := sess.Find(&issues); err != nil {
		return nil, 0, err
	} else {
		return issues, count, nil
	}
}

func setDataSetQueryCondition(Keyword string, userId int64, isAdmin bool) builder.Cond {
	query := builder.NewCond()
	if Keyword != "" {
		query = query.And(builder.Or(builder.Like{"lower_alias", "%" + strings.ToLower(Keyword) + "%"}, builder.Like{"lower_name", "%" + strings.ToLower(Keyword) + "%"}))
	}
	if !isAdmin {
		query = query.And(
			builder.Or(
				builder.Eq{"owner_id": userId},
				builder.Or(
					builder.In("id", builder.Select("`subject_access`.subject_id").
						From("subject_access").
						Where(builder.Eq{"`subject_access`.user_id": userId}.
							And(builder.Eq{"`subject_access`.subject_type": DatasetSubject}).
							And(builder.In("`subject_access`.mode", int(AccessModeOwner), int(AccessModeRead), int(AccessModeWrite), int(AccessModeAdmin))))),
					builder.In("id", builder.Select("`team_subject`.subject_id").
						From("team_subject").
						Where(builder.Eq{"`team_user`.uid ": userId}.
							And(builder.Neq{"`team`.dataset_authorize": int(AccessModeOwner)})).
						Join("INNER", "team_user", "`team_user`.team_id = `team_subject`.team_id").
						Join("INNER", "team", "`team`.id = `team_subject`.team_id")),
				)),
		)
	}
	query = query.And(
		builder.Eq{"is_private": true},
	)
	return query
	//sess.And("(creator_id=" + fmt.Sprint(userId) + " or owner_id=" + fmt.Sprint(userId) + ")")
	//sess.And("(lower_name like '%" + strings.ToLower(Keyword) + "%' and is_private=true)")
}

func SearchDatasetBySQL(Page int, PageSize int, Keyword string, userId int64, isAdmin bool) ([]*DatasetRegistry, int64, error) {

	sess := x.NewSession()
	defer sess.Close()
	query := setDataSetQueryCondition(Keyword, userId, isAdmin)
	count, err := sess.Where(query).Count(new(DatasetRegistry))
	if err != nil {
		return nil, 0, err
	}

	//setDataSetQueryCondition(sess, Keyword, userId)
	sess.Desc("created_unix")
	sess.Limit(PageSize, (Page-1)*PageSize)
	datasets := make([]*DatasetRegistry, 0)
	if err := sess.Where(query).Find(&datasets); err != nil {
		return nil, 0, err
	} else {
		return datasets, count, nil
	}

}
