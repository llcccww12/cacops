package models

import (
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
)

const (
	TechShow          = 1
	TechHide          = 2
	TechMigrating     = 3
	TechMigrateFailed = 4
	TechNotExist      = 5
)

const DefaultTechApprovedStatus = TechShow

type TechConvergeBaseInfo struct {
	ID               int64  `xorm:"pk autoincr"`
	ProjectNumber    string `xorm:"UNIQUE NOT NULL"` //项目立项编号
	ProjectName      string //科技项目名称
	Institution      string //项目承担单位
	ApplyYear        int    //申报年度
	Province         string //所属省（省市）
	Category         string //单位性质
	Recommend        string //推荐单位
	Owner            string //项目负责人
	Phone            string //负责人电话
	Email            string //负责人邮箱
	Contact          string //项目联系人
	ContactPhone     string //联系人电话
	ContactEmail     string //联系人邮箱
	ExecuteMonth     int    //执行周期(月)
	ExecuteStartYear int    //执行开始年份
	ExecuteEndYear   int    //执行结束年份
	ExecutePeriod    string //执行期限
	Type             string //项目类型
	StartUp          string //启动会时间
	NumberTopic      int
	Topic1           string
	Topic2           string
	Topic3           string
	Topic4           string
	Topic5           string
	Topic6           string
	Topic7           string
	AllInstitution   string             `xorm:"TEXT"`
	CreatedUnix      timeutil.TimeStamp `xorm:"INDEX created"`
	UpdatedUnix      timeutil.TimeStamp `xorm:"INDEX updated"`
}

func (t *TechConvergeBaseInfo) Brief() *TechConvergeBrief {
	return &TechConvergeBrief{
		ProjectNumber:  t.ProjectNumber,
		ProjectName:    t.ProjectName,
		Institution:    t.Institution,
		AllInstitution: t.AllInstitution,
	}
}
func (t *TechConvergeBaseInfo) IsValidInstitution(institution string) bool {
	if t.AllInstitution == "" && t.Institution == "" {
		return false
	}

	allInstitution := make([]string, 0)
	if t.AllInstitution != "" {
		allInstitution = strings.Split(t.AllInstitution, ",")
	}
	if t.Institution != "" {
		allInstitution = append(allInstitution, t.Institution)
	}

	newInstitution := strings.Split(institution, ",")
	total := len(newInstitution)
	matched := 0
	for _, n := range newInstitution {
		for _, s := range allInstitution {
			if s == n {
				matched++
				break
			}
		}
	}
	if matched == total {
		return true
	}
	return false
}

type RepoConvergeInfo struct {
	ID          int64 `xorm:"pk autoincr"`
	RepoID      int64
	Url         string
	BaseInfoID  int64
	Institution string
	UID         int64
	Status      int
	CreatedUnix timeutil.TimeStamp    `xorm:"INDEX created"`
	UpdatedUnix timeutil.TimeStamp    `xorm:"INDEX updated"`
	User        *User                 `xorm:"-"`
	Repo        *Repository           `xorm:"-"`
	BaseInfo    *TechConvergeBaseInfo `xorm:"-"`
}

func (r *RepoConvergeInfo) InsertOrUpdate() error {
	if r.ID != 0 {
		_, err := x.ID(r.ID).Update(r)
		return err
	} else {
		_, err := x.InsertOne(r)
		return err
	}
}

func GetTechConvergeBaseInfoByProjectNumber(projectNumber string) (*TechConvergeBaseInfo, error) {
	tb := &TechConvergeBaseInfo{ProjectNumber: projectNumber}
	return getTechConvergeBaseInfo(tb)
}

func GetTechConvergeBaseInfoById(id int64) (*TechConvergeBaseInfo, error) {
	tb := &TechConvergeBaseInfo{ID: id}
	return getTechConvergeBaseInfo(tb)
}

func (baseInfo *TechConvergeBaseInfo) InsertOrUpdate() error {
	if baseInfo.ID != 0 {
		_, err := x.ID(baseInfo.ID).Update(baseInfo)
		return err
	} else {
		_, err := x.InsertOne(baseInfo)
		return err
	}
}

type ErrTechConvergeBaseInfoNotExist struct {
	ID string
}

func (err ErrTechConvergeBaseInfoNotExist) Error() string {
	return "tech.tech_not_exist"
}

func IsErrTechConvergeBaseInfoNotExist(err error) bool {
	_, ok := err.(ErrTechConvergeBaseInfoNotExist)
	return ok
}

func getTechConvergeBaseInfo(tb *TechConvergeBaseInfo) (*TechConvergeBaseInfo, error) {
	has, err := x.Get(tb)
	if err != nil {
		return nil, err
	} else if !has {
		if tb.ProjectNumber != "" {
			return nil, ErrTechConvergeBaseInfoNotExist{tb.ProjectNumber}
		} else {
			return nil, ErrTechConvergeBaseInfoNotExist{strconv.FormatInt(tb.ID, 10)}
		}
	}
	return tb, nil
}

func GetProjectNames() []string {
	var names []string
	x.Table("tech_converge_base_info").Distinct("project_name").Find(&names)
	return names
}
func GetIdByProjectName(name string) []string {
	var ids []int64
	x.Table("tech_converge_base_info").Cols("id").Where("project_name=?", name).Find(&ids)

	idStrs := make([]string, 0, len(ids))
	for _, id := range ids {
		idStrs = append(idStrs, strconv.FormatInt(id, 10))
	}
	return idStrs
}
func GetSummitRepoIds() []int64 {
	var ids []int64
	x.Table("repo_converge_info").Cols("repo_id").Find(&ids)
	return ids
}

func GetTechRepoTopics(limit int) []string {

	repoIds := GetSummitRepoIds()
	if len(repoIds) == 0 {
		return []string{}
	}

	//select name, repo_count from topic a, repo_topic b
	//where a.id=b.topic_id and repo_id in (1,3) order by repo_count desc
	inCondition := "repo_id in ("

	const MaxINItems = 1000
	for i := 0; i < len(repoIds); i++ {
		if i == len(repoIds)-1 {
			inCondition += strconv.FormatInt(repoIds[i], 10)
		} else if (i+1)%MaxINItems == 0 {
			inCondition += strconv.FormatInt(repoIds[i], 10) + ") or repo_id in ("
		} else {
			inCondition += strconv.FormatInt(repoIds[i], 10) + ","
		}

	}
	inCondition += ")"

	sql := "select distinct name, repo_count from topic a, repo_topic b where a.id=b.topic_id and (" + inCondition + ") order by repo_count desc"
	if limit > 0 {
		sql += "limit " + strconv.Itoa(limit)
	}

	result, err := x.QueryString(sql)
	if err != nil {
		return []string{}
	}
	var topics []string
	for _, record := range result {
		topics = append(topics, record["name"])
	}
	return topics

}

func GetProjectTypes() []string {

	sql := "SELECT COUNT(id) AS theCount, type from tech_converge_base_info GROUP BY type ORDER BY theCount DESC"
	result, err := x.QueryString(sql)
	if err != nil {
		return []string{}
	}
	var projectTypes []string
	for _, record := range result {
		projectTypes = append(projectTypes, record["type"])
	}
	return projectTypes

}
func GetApplyExecuteYears() ([]int, []int) {
	apply, executeStart, executeEnd := GetYearInfos()
	applyEnd := time.Now().Year()

	var applyArray []int
	var executeArray []int

	for i := apply; i <= applyEnd; i++ {
		applyArray = append(applyArray, i)

	}
	for i := executeStart; i <= executeEnd; i++ {
		executeArray = append(executeArray, i)

	}
	return applyArray, executeArray
}
func GetYearInfos() (int, int, int) {

	sql := "select min(apply_year) as apply_year,min(CASE  WHEN execute_start_year != 0 THEN execute_start_year END) as execute_start_year,max(execute_end_year) as execute_end_year from tech_converge_base_info"
	result, err := x.QueryString(sql)
	if err != nil {
		return 2018, 2019, 2024
	}

	for _, record := range result {
		apply, _ := strconv.Atoi(record["apply_year"])
		executeStart, _ := strconv.Atoi(record["execute_start_year"])
		executeEnd, _ := strconv.Atoi(record["execute_end_year"])
		return apply, executeStart, executeEnd
	}
	return 2018, 2019, 2024
}

func GetAllInstitutions() []string {
	var names []string
	x.Table("tech_converge_base_info").Cols("all_institution").Find(&names)
	var allNames []string
	for _, name := range names {
		singleNames := strings.Split(name, ",")
		for _, singleName := range singleNames {
			if singleName != "" {
				if !contains(allNames, singleName) {
					allNames = append(allNames, singleName)
				}
			}
		}

	}
	return allNames
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type SearchTechOpt struct {
	Q           string //科技项目名称
	ProjectType string
	Institution string
	ApplyYear   int
	ExecuteYear int
	OrderBy     string
	ListOptions
}

type SearchRepoOpt struct {
	Q           string //项目名称，简介
	ProjectName string
	Topic       string
	Institution string
	OrderBy     string
	ListOptions
}

type SearchUserRepoOpt struct {
	User *User
	ListOptions
}

type TechRepoInfoUser struct {
	ID                      int64              `json:"id"`
	ProjectNumber           string             `json:"no"`
	ProjectName             string             `json:"name"`
	Institution             string             `json:"institution"`
	AllInstitution          string             `json:"all_institution"`
	Url                     string             `json:"url"`
	RepoName                string             `json:"repo_name"`
	RepoOwnerName           string             `json:"repo_owner_name"`
	ContributionInstitution string             `json:"contribution_institution"`
	UserName                string             `json:"user_name"`
	Status                  int                `json:"status"`
	CreatedUnix             timeutil.TimeStamp `json:"created_unix"`
	UpdatedUnix             timeutil.TimeStamp `json:"updated_unix"`
}
type TechRepoInfoAdmin struct {
	ID                      int64              `json:"id"`
	Url                     string             `json:"url"`
	ContributionInstitution string             `json:"contribution_institution"`
	Status                  int                `json:"status"`
	ProjectNumber           string             `json:"no"`
	ProjectName             string             `json:"name"`
	ApplyYear               int                `json:"apply_year"`
	Institution             string             `json:"institution"`
	Province                string             `json:"province"`
	Category                string             `json:"category"`
	Recommend               string             `json:"recommend"`
	Owner                   string             `json:"owner"`
	Phone                   string             `json:"phone"`
	Email                   string             `json:"email"`
	Contact                 string             `json:"contact"`
	ContactPhone            string             `json:"contact_phone"`
	ContactEmail            string             `json:"contact_email"`
	ExecuteMonth            int                `json:"execute_month"`
	ExecutePeriod           string             `json:"execute_period"`
	Type                    string             `json:"type"`
	StartUp                 string             `json:"start_up"`
	NumberTopic             int                `json:"number_topic"`
	Topic1                  string             `json:"topic1"`
	Topic2                  string             `json:"topic2"`
	Topic3                  string             `json:"topic3"`
	Topic4                  string             `json:"topic4"`
	Topic5                  string             `json:"topic5"`
	AllInstitution          string             `json:"all_institution"`
	RepoName                string             `json:"repo_name"`
	RepoOwnerName           string             `json:"repo_owner_name"`
	UserName                string             `json:"user_name"`
	CreatedUnix             timeutil.TimeStamp `json:"created_unix"`
	UpdatedUnix             timeutil.TimeStamp `json:"updated_unix"`
}

type RepoWithInstitution struct {
	ID            int64              `json:"id"`
	OwnerID       int64              `json:"owner_id"`
	OwnerName     string             `json:"owner_name"`
	Name          string             `json:"name"`
	Alias         string             `json:"alias"`
	Topics        []string           `json:"topics"`
	Description   string             `json:"description"`
	Institution   string             `json:"institution"`
	RelAvatarLink string             `json:"rel_avatar_link"`
	UpdatedUnix   timeutil.TimeStamp `json:"updated_unix"`
}

type TechRepoInfo struct {
	ID             int64  `json:"id"`
	ProjectName    string `json:"project_name"`
	Institution    string `json:"institution"`
	Type           string `json:"type"`
	ApplyYear      int    `json:"apply_year"`
	ExecutePeriod  string `json:"execute_period"`
	AllInstitution string `json:"all_institution"`
	RepoCount      int    `json:"repo_numer"`
	Repos          []*RepoWithInstitution
}

func ShowTechRepo(ids []int64, show bool) error {

	status := TechShow
	if !show {
		status = TechHide
	}

	idStrs := make([]string, 0, len(ids))
	for _, id := range ids {
		idStrs = append(idStrs, strconv.FormatInt(id, 10))
	}

	sql := "update repo_converge_info  set status=? where id in (" + strings.Join(idStrs, ",") + ") and (status=" + strconv.Itoa(TechShow) + " or status=" + strconv.Itoa(TechHide) + ")"
	_, err := x.Exec(sql, status)
	return err
}

func GetTechRepoInfoForUser(opts *SearchUserRepoOpt) ([]*RepoConvergeInfo, int64, error) {
	cond := buildTechRepoForUserCondition(opts)
	total, err := x.Where(cond).Count(new(RepoConvergeInfo))
	if err != nil {
		return nil, 0, err
	}
	repoConvergeInfos := make([]*RepoConvergeInfo, 0)
	err = x.Where(cond).Limit(opts.PageSize, (opts.Page-1)*opts.PageSize).Desc("status").Find(&repoConvergeInfos)
	if err != nil {
		return nil, 0, err
	}

	loadAttributes(repoConvergeInfos)

	return repoConvergeInfos, total, nil

}

func loadAttributes(infos []*RepoConvergeInfo) {
	for _, info := range infos {
		info.User, _ = GetUserByID(info.UID)
		info.Repo, _ = GetRepositoryByID(info.RepoID)
		info.BaseInfo, _ = GetTechConvergeBaseInfoById(info.BaseInfoID)
	}
}

func buildTechRepoForUserCondition(opts *SearchUserRepoOpt) builder.Cond {
	var cond = builder.NewCond()
	if opts.User != nil {
		cond = cond.And(builder.Eq{"uid": opts.User.ID})
	}
	return cond
}

func GetAvailableRepoConvergeInfo(opt *SearchRepoOpt) ([]*RepoConvergeInfo, error) {

	repos := make([]*RepoConvergeInfo, 0)
	err := x.Table("repo_converge_info").Where(buildRepoFilterCond(opt)).Find(&repos)
	return repos, err

}

func buildRepoFilterCond(opt *SearchRepoOpt) string {

	sql := "status=" + strconv.Itoa(TechShow)

	if opt.Institution != "" {
		sql += " and (institution like '%" + opt.Institution + ",%'" + " or institution like '%," + opt.Institution + "%'" + " or institution = '" + opt.Institution + "')"
	}

	if opt.ProjectName != "" {
		baseInfoIds := GetIdByProjectName(opt.ProjectName)
		if len(baseInfoIds) > 0 {

			sql += " and base_info_id in (" + strings.Join(baseInfoIds, ",") + ")"
		}

	}
	return sql
}

func SearchTechRepoInfo(opt *SearchTechOpt) ([]*TechRepoInfo, int64, error) {

	sql := `select a.*,COALESCE(b.count,0) as repo_count,  COALESCE(b.max,0) as max from tech_converge_base_info a left join
		(select base_info_id,count(id),max(updated_unix) from repo_converge_info where status=` + strconv.Itoa(TechShow) + ` GROUP BY base_info_id ) b
			on a.id=b.base_info_id`
	totalSql := "select count(*) from (" + sql + ") c" + buildTechFilterCond(opt)
	total, err := x.SQL(totalSql).Count(new(TechConvergeBaseInfo))
	resultList := make([]*TechRepoInfo, 0)
	if err != nil {
		return resultList, total, err

	}
	resultSql := "select id,project_name, institution,type,apply_year,execute_period,all_institution,repo_count from (" +
		sql + ") c " + buildTechFilterCond(opt) + opt.OrderBy + " offset " + strconv.Itoa((opt.Page-1)*opt.PageSize) + " limit " + strconv.Itoa(opt.PageSize)

	resultMap, err := x.QueryInterface(resultSql)
	if err == nil {
		for _, record := range resultMap {
			resultList = append(resultList, &TechRepoInfo{
				ID:             record["id"].(int64),
				ProjectName:    record["project_name"].(string),
				Institution:    record["institution"].(string),
				Type:           record["type"].(string),
				ApplyYear:      int(record["apply_year"].(int64)),
				ExecutePeriod:  record["execute_period"].(string),
				AllInstitution: record["all_institution"].(string),
				RepoCount:      int(record["repo_count"].(int64)),
			})
		}
	}

	loadRepoInfoForTech(resultList)

	return resultList, total, err

}

func buildTechFilterCond(opt *SearchTechOpt) string {

	sql := ""

	if opt.Q != "" {
		sql += getWherePrefix(sql) + " project_name like '%" + opt.Q + "%'"

	}
	if opt.ProjectType != "" {
		sql += getWherePrefix(sql) + " type ='" + opt.ProjectType + "'"
	}
	if opt.ApplyYear != 0 {
		sql += getWherePrefix(sql) + " apply_year =" + strconv.Itoa(opt.ApplyYear)
	}
	if opt.Institution != "" {
		sql += getWherePrefix(sql) + " (all_institution like '%" + opt.Institution + ",%'" + " or all_institution like '%," + opt.Institution + "%'" + " or all_institution = '" + opt.Institution + "')"
	}

	if opt.ExecuteYear != 0 {
		sql += getWherePrefix(sql) + " execute_start_year <=" + strconv.Itoa(opt.ExecuteYear) + " and execute_end_year >=" + strconv.Itoa(opt.ExecuteYear)
	}
	return sql
}

func getWherePrefix(sql string) string {
	if sql == "" {
		return " where "
	}
	return " and "
}

func loadRepoInfoForTech(list []*TechRepoInfo) {

	for _, techRepo := range list {
		techRepo.Repos = []*RepoWithInstitution{}
		if techRepo.RepoCount > 0 {
			var repoIds []int64
			x.Table("repo_converge_info").Cols("repo_id").Where("base_info_id=? and status=?", techRepo.ID, TechShow).Limit(2).Desc("updated_unix").Find(&repoIds)

			resultMap, err := GetRepositoriesMapByIDs(repoIds)

			if err == nil {

				for _, repoId := range repoIds {
					repo, ok := resultMap[repoId]
					if ok {
						techRepo.Repos = append(techRepo.Repos, &RepoWithInstitution{
							ID:            repo.ID,
							Institution:   techRepo.Institution,
							OwnerID:       repo.OwnerID,
							OwnerName:     repo.OwnerName,
							Name:          repo.Name,
							Alias:         repo.Alias,
							Topics:        repo.Topics,
							Description:   repo.Description,
							RelAvatarLink: repo.RelAvatarLink(),
							UpdatedUnix:   repo.UpdatedUnix,
						})
					}

				}

			}
		}
	}

}

type TechConvergeBrief struct {
	ProjectNumber  string `json:"no"`          //项目立项编号
	ProjectName    string `json:"name"`        //科技项目名称
	Institution    string `json:"institution"` //项目承担单位
	AllInstitution string `json:"all_institution"`
}

type FindTechOpt struct {
	TechNo      string
	ProjectName string
	Institution string
}

func FindTech(opt FindTechOpt) ([]*TechConvergeBaseInfo, error) {
	var cond = builder.NewCond()
	if opt.TechNo != "" {
		cond = cond.And(builder.Like{"project_number", opt.TechNo})
	}
	if opt.ProjectName != "" {
		cond = cond.And(builder.Like{"project_name", opt.ProjectName})
	}
	if opt.Institution != "" {
		cond = cond.And(builder.Like{"institution", opt.Institution}.Or(builder.Like{"all_institution", opt.Institution}))
	}

	r := make([]*TechConvergeBaseInfo, 0)
	err := x.Where(cond).OrderBy("updated_unix desc").Find(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetTechByTechNo(techNo string) (*TechConvergeBaseInfo, error) {
	var tech = &TechConvergeBaseInfo{}
	has, err := x.Where("project_number = ?", techNo).Get(tech)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrTechConvergeBaseInfoNotExist{}
	}
	return tech, nil

}

type GetRepoConvergeOpts struct {
	RepoId     int64
	BaseInfoId int64
	Status     []int
}

func GetRepoConverge(opts GetRepoConvergeOpts) ([]*RepoConvergeInfo, error) {
	r := make([]*RepoConvergeInfo, 0)
	cond := builder.NewCond()
	if opts.RepoId > 0 {
		cond = cond.And(builder.Eq{"repo_id": opts.RepoId})
	}
	if opts.BaseInfoId > 0 {
		cond = cond.And(builder.Eq{"base_info_id": opts.BaseInfoId})
	}
	if len(opts.Status) > 0 {
		cond = cond.And(builder.In("status", opts.Status))
	}
	err := x.Where(cond).Find(&r)
	if err != nil {
		return nil, err
	}
	return r, nil

}

func UpdateRepoConvergeStatus(id int64, status int) (int64, error) {
	return x.ID(id).Update(&RepoConvergeInfo{
		Status: status,
	})
}
