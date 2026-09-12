package models

import (
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
)

type RecommendOrg struct {
	ID          int64              `xorm:"pk autoincr"`
	Order       int64              `xorm:"INDEX NOT NULL unique"`
	OrgID       int64              `xorm:"INDEX NOT NULL unique"`
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
}

type RecommendOrgInfo struct {
	RecommendOrg `xorm:"extends"`
	User `xorm:"extends"`
}

type RecommendOrgList []*RecommendOrg
type RecommendOrgInfoList []*RecommendOrgInfo

func getRecommendOrgs(e Engine) (RecommendOrgList, error) {
	orgs := make(RecommendOrgList, 0, 10)
	err := e.Asc("order").
		Find(&orgs)
	return orgs, err
}

func GetRecommendOrgs() (RecommendOrgList, error) {
	return getRecommendOrgs(x)
}

func getRecommendOrgInfos(e Engine) (RecommendOrgInfoList, error) {
	orgs := make(RecommendOrgInfoList, 0, 10)
	if err := e.Table(&RecommendOrg{}).Join("INNER", "`user`", "`user`.id = `recommend_org`.org_id").
		OrderBy("recommend_org.order").Find(&orgs); err != nil {
		return orgs, err
	}
	return orgs, nil
}

func GetRecommendOrgInfos() (RecommendOrgInfoList, error) {
	return getRecommendOrgInfos(x)
}

func delRecommendOrgs(e Engine) error {
	sql := "delete from recommend_org"
	_, err := e.Exec(sql)
	return err
}

func UpdateRecommendOrgs(orgs RecommendOrgList) error {
	if err := delRecommendOrgs(x); err != nil {
		log.Error("delRecommendOrgs failed:%v", err.Error())
		return err
	}

	if _, err := x.Insert(&orgs); err != nil {
		log.Error("Insert failed:%v", err.Error())
		return err
	}

	return nil
}
