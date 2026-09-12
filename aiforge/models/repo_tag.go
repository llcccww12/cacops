package models

import (
	"fmt"
	"strconv"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
)

type OfficialTag struct {
	ID          int64              `xorm:"pk autoincr"`
	Name        string             `xorm:"NOT NULL"`
	Code        string             `xorm:"NOT NULL"`
	Limit       int                `xorm:"NOT NULL default(-1)"`
	Status      int                `xorm:"NOT NULL default(0)"`
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
	UpdatedUnix timeutil.TimeStamp `xorm:"updated"`
}

type OfficialTagRepos struct {
	ID          int64              `xorm:"pk autoincr"`
	OrgID       int64              `xorm:"NOT NULL INDEX"`
	TagID       int64              `xorm:"NOT NULL"`
	RepoID      int64              `xorm:"NOT NULL INDEX"`
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
	UpdatedUnix timeutil.TimeStamp `xorm:"updated"`
}

type OfficialModel struct {
	ID          int64              `xorm:"pk autoincr"`
	OrgID       int64              `xorm:"NOT NULL INDEX"`
	ModelID     string             `xorm:"NOT NULL INDEX"`
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
	UpdatedUnix timeutil.TimeStamp `xorm:"updated"`
}

type TagModelSelected struct {
	ModelID   string
	ModelName string
	Selected  bool
}

type OfficialDataset struct {
	ID          int64              `xorm:"pk autoincr"`
	OrgID       int64              `xorm:"NOT NULL INDEX"`
	DatasetId   int64              `xorm:"NOT NULL INDEX"`
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
	UpdatedUnix timeutil.TimeStamp `xorm:"updated"`
}

type TagDatasetSelected struct {
	DatasetId   int64
	DatasetName string
	Selected    bool
}

type OfficialDatasetRegistry struct {
	ID          int64              `xorm:"pk autoincr"`
	OrgID       int64              `xorm:"NOT NULL INDEX"`
	DatasetId   string             `xorm:"uuid NOT NULL INDEX"`
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
	UpdatedUnix timeutil.TimeStamp `xorm:"updated"`
}

type TagDatasetRegistrySelected struct {
	DatasetId   string
	DatasetName string
	Selected    bool
}

type TagReposBrief struct {
	RepoID   int64
	RepoName string
	Alias    string
	TagID    int64
}

type TagReposSelected struct {
	RepoID   int64
	RepoName string
	Selected bool
}

type TagsDetail struct {
	TagId    int64
	TagName  string
	TagLimit int
	RepoList []*Repository
}

func GetTagByID(id int64) (*OfficialTag, error) {
	r := &OfficialTag{
		ID: id,
	}
	has, err := x.Get(r)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrTagNotExist{0}
	}
	return r, nil
}

func UpdateTagDatasetByID(orgID int64, datasetids []string) error {
	sess := x.NewSession()
	defer sess.Close()

	if err := sess.Begin(); err != nil {
		return fmt.Errorf("UpdateTagDatasetByID[orgID: %d,error:%v", orgID, err)
	}
	//delete old tag repos
	r := &OfficialDataset{
		OrgID: orgID,
	}
	_, err := sess.Delete(r)
	if err != nil {
		return err
	}

	if len(datasetids) == 0 {
		return sess.Commit()
	}

	//add new tag repos
	data := make([]*OfficialDataset, 0)
	for _, datasetid := range datasetids {
		datasetidi64, err := strconv.ParseInt(datasetid, 10, 64)
		if err != nil {
			fmt.Printf("convert to 64 error: %v\n", err)
		}
		data = append(data, &OfficialDataset{
			OrgID:     orgID,
			DatasetId: datasetidi64,
		})
	}
	_, err = sess.Insert(&data)
	if err != nil {
		sess.Rollback()
		return err
	}
	return sess.Commit()
}

func UpdateTagDatasetRegistryByID(orgID int64, datasetids []string) error {
	sess := x.NewSession()
	defer sess.Close()

	if err := sess.Begin(); err != nil {
		return fmt.Errorf("UpdateTagDatasetRegistryByID[orgID: %d,error:%v", orgID, err)
	}
	//delete old tag repos
	r := &OfficialDatasetRegistry{
		OrgID: orgID,
	}
	_, err := sess.Delete(r)
	if err != nil {
		return err
	}

	if len(datasetids) == 0 {
		return sess.Commit()
	}

	//add new tag repos
	datas := make([]*OfficialDatasetRegistry, 0)
	for _, datasetid := range datasetids {
		if datasetid == "" {
			continue
		}
		datas = append(datas, &OfficialDatasetRegistry{
			OrgID:     orgID,
			DatasetId: datasetid,
		})
	}
	if len(datas) == 0 {
		return sess.Commit()
	}
	_, err = sess.Insert(&datas)
	if err != nil {
		sess.Rollback()
		return err
	}
	return sess.Commit()
}

func DeleteRepoTag(orgID int64, repoID int64) {
	sess := x.NewSession()
	defer sess.Close()
	//delete old tag repos
	r := &OfficialTagRepos{
		RepoID: repoID,
		OrgID:  orgID,
	}
	_, err := sess.Delete(r)
	if err != nil {
		log.Info("delete failed." + err.Error())
	}
}

func UpdateTagModelByID(orgID int64, modelids []string) error {
	sess := x.NewSession()
	defer sess.Close()

	if err := sess.Begin(); err != nil {
		return fmt.Errorf("UpdateTagReposByID[orgID: %d,error:%v", orgID, err)
	}
	//delete old tag repos
	r := &OfficialModel{
		OrgID: orgID,
	}
	_, err := sess.Delete(r)
	if err != nil {
		return err
	}

	if len(modelids) == 0 {
		return sess.Commit()
	}

	//add new tag repos
	data := make([]*OfficialModel, 0)
	for _, modelId := range modelids {
		data = append(data, &OfficialModel{
			OrgID:   orgID,
			ModelID: modelId,
		})
	}
	_, err = sess.Insert(&data)
	if err != nil {
		sess.Rollback()
		return err
	}
	return sess.Commit()
}

func UpdateTagReposByID(orgID int64, repoIdList []int64) error {
	sess := x.NewSession()
	defer sess.Close()

	if err := sess.Begin(); err != nil {
		return fmt.Errorf("UpdateTagReposByID[orgID: %d,error:%v", orgID, err)
	}
	//delete old tag repos
	r := &OfficialTagRepos{
		TagID: 1,
		OrgID: orgID,
	}
	_, err := sess.Delete(r)
	if err != nil {
		return err
	}

	if len(repoIdList) == 0 {
		return sess.Commit()
	}

	//add new tag repos
	data := make([]*OfficialTagRepos, 0)
	for _, repoId := range repoIdList {
		data = append(data, &OfficialTagRepos{
			OrgID:  orgID,
			TagID:  1,
			RepoID: repoId,
		})
	}
	_, err = sess.Insert(&data)
	if err != nil {
		sess.Rollback()
		return err
	}
	return sess.Commit()
}

func GetTagRepos(orgID int64) ([]TagReposSelected, error) {
	t := make([]TagReposBrief, 0)
	const SQLCmd = "select t1.id as repo_id,t1.name as repo_name,t1.alias,t2.id as tag_id from repository t1 left join official_tag_repos t2 on (t1.id = t2.repo_id) where t1.owner_id = ? and t1.is_private = false order by t1.updated_unix desc"

	if err := x.SQL(SQLCmd, orgID).Find(&t); err != nil {
		return nil, err
	}
	r := make([]TagReposSelected, 0)
	for _, v := range t {
		selected := false
		if v.TagID > 0 {
			selected = true
		}
		repoName := v.Alias
		if v.Alias == "" {
			repoName = v.RepoName
		}
		r = append(r, TagReposSelected{
			RepoID:   v.RepoID,
			RepoName: repoName,
			Selected: selected,
		})
	}
	return r, nil
}

func GetAvailableOrgOfficialDatasetRegistry(orgID int64) ([]TagDatasetRegistrySelected, error) {
	// OfficialDatasetRegistry
	t := make([]TagDatasetRegistrySelected, 0)
	const SQLCmd = "SELECT t1.id AS dataset_id,t1.alias AS dataset_name, CASE WHEN t2.dataset_id IS NOT NULL THEN TRUE ELSE FALSE END AS selected  FROM dataset_registry t1 LEFT JOIN official_dataset_registry t2 ON (t1.id = t2.dataset_id) WHERE t1.owner_id = ? AND t1.is_private = false ORDER BY t1.updated_unix DESC"

	if err := x.SQL(SQLCmd, orgID).Find(&t); err != nil {
		return nil, err
	}

	return t, nil
}

func GetAvailableOrgOfficialAimodel(orgID int64) ([]TagModelSelected, error) {
	t := make([]TagModelSelected, 0)
	SQLCmd := "SELECT t1.id AS model_id,t1.alias AS model_name, CASE WHEN t2.model_id IS NOT NULL THEN TRUE ELSE FALSE END AS selected  FROM ai_model_manage t1 LEFT JOIN official_model t2 ON (t1.id = t2.model_id) WHERE t1.owner_id = ? AND t1.is_private = false ORDER BY t1.updated_unix DESC"
	if err := x.SQL(SQLCmd, orgID).Find(&t); err != nil {
		return nil, err
	}

	return t, nil
}

func GetTagDataset(orgID int64) ([]TagDatasetSelected, error) {
	publicdataset := GetOwnedPublicDatasetsByUserID(orgID)
	r := make([]TagDatasetSelected, 0)
	if len(publicdataset) > 0 {
		selectdatasets, terr := getAllTagDataset(orgID)
		mapselect := make(map[int64]bool, 0)
		if terr == nil {
			for _, v := range selectdatasets {
				mapselect[v.DatasetId] = true
			}
		}
		for _, v := range publicdataset {
			selected := mapselect[v.ID]
			r = append(r, TagDatasetSelected{
				DatasetId:   v.ID,
				DatasetName: v.Title,
				Selected:    selected,
			})
		}
	}
	return r, nil
}

func GetTagDatasetForCard(orgID int64) ([]*Dataset, error) {

	publicdataset := GetOwnedPublicDatasetsByUserID(orgID)
	r := make([]*Dataset, 0)
	if len(publicdataset) > 0 {
		selectdatasets, terr := getAllTagDataset(orgID)
		mapselect := make(map[int64]bool, 0)
		if terr == nil {
			for _, v := range selectdatasets {
				mapselect[v.DatasetId] = true
			}
		}
		for _, v := range publicdataset {
			if mapselect[v.ID] {
				r = append(r, v)
			}
		}
	}
	return r, nil
}

func GetOrgOfficialDatasetRegistryForCard(orgID int64) (DatasetRegistryList, error) {
	publicdataset, err := GetOwnedPublicDatasetRegistrysByUserID(orgID)
	if err != nil {
		return nil, err
	}
	r := make([]*DatasetRegistry, 0)
	if len(publicdataset) > 0 {
		selectdatasets, terr := getAllTagDatasetRegistry(orgID)
		mapselect := make(map[string]bool, 0)
		if terr == nil {
			for _, v := range selectdatasets {
				mapselect[v.DatasetId] = true
			}
		}
		for _, v := range publicdataset {
			if mapselect[v.ID] {
				r = append(r, v)
			}
		}
	}
	return r, nil

}

func GetOrgOfficialAimodelForCard(orgID int64) (AimodelList, error) {
	publicaimodel, err := GetOwnedPublicAimodelsByUserID(orgID)
	if err != nil {
		return nil, err
	}
	r := make([]*AiModelManage, 0)
	if len(publicaimodel) > 0 {
		selectaimodels, terr := getAllTagModels(orgID)
		mapselect := make(map[string]bool, 0)
		if terr == nil {
			for _, v := range selectaimodels {
				mapselect[v.ModelID] = true
			}
		}
		for _, v := range publicaimodel {
			if mapselect[v.ID] {
				r = append(r, v)
			}
		}
	}
	return r, nil

}

func GetTagModel(orgID int64) ([]TagModelSelected, error) {
	t := make([]TagReposBrief, 0)
	const RepoSQLCmd = "select t1.id as repo_id,t1.name as repo_name,t1.alias from repository t1  where t1.owner_id = ? and t1.is_private = false order by t1.updated_unix desc"
	if err := x.SQL(RepoSQLCmd, orgID).Find(&t); err != nil {
		return nil, err
	}
	repoIds := make([]int64, 0)
	for _, v := range t {
		repoIds = append(repoIds, v.RepoID)
	}
	r := make([]TagModelSelected, 0)
	if len(repoIds) > 0 {
		selectModels, terr := getAllTagModels(orgID)
		mapselect := make(map[string]bool, 0)
		if terr == nil {
			for _, v := range selectModels {
				mapselect[v.ModelID] = true
			}
		}
		models := QueryPublicModelByRepoIdS(repoIds)
		for _, v := range models {
			selected := mapselect[v.ID]
			r = append(r, TagModelSelected{
				ModelID:   v.ID,
				ModelName: v.Name,
				Selected:  selected,
			})
		}
	}
	return r, nil
}

func GetTagModelForCard(orgID int64) ([]*AiModelManage, error) {
	t := make([]TagReposBrief, 0)
	const RepoSQLCmd = "select t1.id as repo_id,t1.name as repo_name,t1.alias from repository t1  where t1.owner_id = ? and t1.is_private = false order by t1.updated_unix desc"
	if err := x.SQL(RepoSQLCmd, orgID).Find(&t); err != nil {
		return nil, err
	}
	repoIds := make([]int64, 0)
	for _, v := range t {
		repoIds = append(repoIds, v.RepoID)
	}
	r := make([]*AiModelManage, 0)
	if len(repoIds) > 0 {
		selectModels, terr := getAllTagModels(orgID)
		mapselect := make(map[string]bool, 0)
		if terr == nil {
			for _, v := range selectModels {
				mapselect[v.ModelID] = true
			}
		}
		models := QueryPublicModelByRepoIdS(repoIds)
		for _, v := range models {
			if mapselect[v.ID] {
				r = append(r, v)
			}
		}
	}
	return r, nil
}

func GetOrgModelList(orgID int64, uid int64, keyword, labelFilter string, orderBy SearchOrderBy, page int, pageSize int, isAdmin bool) ([]*AiModelManage, int64, error) {
	// t := make([]TagReposBrief, 0)
	// const RepoSQLCmd = "select t1.id as repo_id,t1.name as repo_name,t1.alias from repository t1  where t1.owner_id = ? and t1.is_private = false order by t1.updated_unix desc"
	// if err := x.SQL(RepoSQLCmd, orgID).Find(&t); err != nil {
	// 	return nil, 0, err
	// }
	// repoIds := make([]int64, 0)
	// for _, v := range t {
	// 	repoIds = append(repoIds, v.RepoID)
	// }
	// r := make([]*AiModelManage, 0)
	// var count int64
	// if len(repoIds) > 0 {
	r, count := QueryPublicModelAndOwnByRepoIdS(orgID, uid, keyword, labelFilter, orderBy, page, pageSize, isAdmin)
	//}
	return r, count, nil
}

func GetOrgModelLabelList(orgID int64, uid int64) ([]*AiModelManage, error) {
	t := make([]TagReposBrief, 0)
	const RepoSQLCmd = "select t1.id as repo_id,t1.name as repo_name,t1.alias from repository t1  where t1.owner_id = ? and t1.is_private = false order by t1.updated_unix desc"
	if err := x.SQL(RepoSQLCmd, orgID).Find(&t); err != nil {
		return nil, err
	}
	repoIds := make([]int64, 0)
	for _, v := range t {
		repoIds = append(repoIds, v.RepoID)
	}
	r := make([]*AiModelManage, 0)
	if len(repoIds) > 0 {
		r = QueryPublicModelLabelAndOwnByRepoIdS(repoIds, uid)
	}
	return r, nil
}

func getAllTagModels(orgID int64) ([]OfficialModel, error) {
	t := make([]OfficialModel, 0)
	const SQLCmd = "select * from official_model  where org_id = ?"
	if err := x.SQL(SQLCmd, orgID).Find(&t); err != nil {
		return nil, err
	}
	return t, nil
}

func getAllTagDataset(orgID int64) ([]OfficialDataset, error) {
	t := make([]OfficialDataset, 0)
	const SQLCmd = "select * from official_dataset  where org_id = ?"
	if err := x.SQL(SQLCmd, orgID).Find(&t); err != nil {
		return nil, err
	}
	return t, nil
}

func getAllTagDatasetRegistry(orgID int64) ([]OfficialDatasetRegistry, error) {
	t := make([]OfficialDatasetRegistry, 0)
	const SQLCmd = "select * from official_dataset_registry  where org_id = ?"
	if err := x.SQL(SQLCmd, orgID).Find(&t); err != nil {
		return nil, err
	}
	return t, nil
}

func GetAllOfficialTagRepos(orgID int64, isOwner bool) ([]TagsDetail, error) {
	result := make([]TagsDetail, 0)

	repos, err := GetOfficialTagDetail(orgID, 1)
	if err != nil {
		return nil, err
	}
	if len(repos) == 0 && !isOwner {
		return nil, fmt.Errorf("parameter error1.")
	}
	result = append(result, TagsDetail{
		TagLimit: 9,
		RepoList: repos,
	})
	return result, nil
}

func GetOfficialTagDetail(orgID, tagId int64) ([]*Repository, error) {
	t := make([]*Repository, 0)
	const SQLCmd = "select t2.* from official_tag_repos t1 inner join repository t2 on t1.repo_id = t2.id where t1.org_id = ? and t1.tag_id=? order by t2.updated_unix desc"

	if err := x.SQL(SQLCmd, orgID, tagId).Find(&t); err != nil {
		return nil, err
	}
	return t, nil
}

func GetAllOfficialTags() ([]OfficialTag, error) {
	//todo  redis?
	o := make([]OfficialTag, 0)
	err := x.Where("status = ?", 0).OrderBy("updated_unix desc").Find(&o)
	if err != nil {
		log.Error("GetAllOfficialTags error,%v", err)
		return nil, err
	}
	return o, nil
}

type FindSelectedReposOpts struct {
	ListOptions
	OrgId      int64
	OnlyPublic bool
}

func GetSelectedRepos(opts FindSelectedReposOpts) ([]*Repository, error) {
	if opts.Page < 1 {
		opts.Page = 1
	}
	var cond = builder.NewCond()
	cond = cond.And(builder.Eq{"official_tag.code": "selected"})
	if opts.OrgId > 0 {
		cond = cond.And(builder.Eq{"official_tag_repos.org_id": opts.OrgId})
	}
	if opts.OnlyPublic {
		cond = cond.And(builder.Eq{"repository.is_private": false})
	}
	t := make([]*Repository, 0)
	err := x.Join("inner", "official_tag_repos", "repository.id = official_tag_repos.repo_id").
		Join("inner", "official_tag", "official_tag.id = official_tag_repos.tag_id").
		Where(cond).OrderBy("repository.updated_unix desc").Limit(opts.PageSize, (opts.Page-1)*opts.PageSize).Find(&t)

	if err != nil {
		return nil, err
	}

	return t, nil
}
