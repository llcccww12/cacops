package repo

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/util"

	"github.com/360EntSecGroup-Skylar/excelize/v2"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/repository"

	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/setting"
)

const DEFAULT_PAGE_SIZE = 10
const DATE_FORMAT = "2006-01-02"
const MONTH_FORMAT = "2006-01"
const EXCEL_DATE_FORMAT = "20060102"
const CREATE_TIME_FORMAT = "2006/01/02 15:04:05"
const UPDATE_TIME_FORMAT = "2006-01-02 15:04:05"

type ProjectsPeriodData struct {
	RecordBeginTime string                  `json:"recordBeginTime"`
	LastUpdatedTime string                  `json:"lastUpdatedTime"`
	PageSize        int                     `json:"pageSize"`
	TotalPage       int                     `json:"totalPage"`
	TotalCount      int64                   `json:"totalCount"`
	PageRecords     []*models.RepoStatistic `json:"pageRecords"`
}

type UserInfo struct {
	User          string `json:"user"`
	Mode          int    `json:"mode"`
	PR            int64  `json:"pr"`
	Commit        int    `json:"commit"`
	RelAvatarLink string `json:"relAvatarLink"`
	Email         string `json:"email"`
}

type ProjectLatestData struct {
	RecordBeginTime  string     `json:"recordBeginTime"`
	LastUpdatedTime  string     `json:"lastUpdatedTime"`
	CreatTime        string     `json:"creatTime"`
	OpenI            float64    `json:"openi"`
	Comment          int64      `json:"comment"`
	View             int64      `json:"view"`
	Download         int64      `json:"download"`
	IssueClosedRatio float32    `json:"issueClosedRatio"`
	Impact           float64    `json:"impact"`
	Completeness     float64    `json:"completeness"`
	Liveness         float64    `json:"liveness"`
	ProjectHealth    float64    `json:"projectHealth"`
	TeamHealth       float64    `json:"teamHealth"`
	Growth           float64    `json:"growth"`
	Description      string     `json:"description"`
	Top10            []UserInfo `json:"top10"`
}

type ProjectSummaryBaseData struct {
	NumReposAdd       int64  `json:"numReposAdd"`
	NumRepoPublicAdd  int64  `json:"numRepoPublicAdd"`
	NumRepoPrivateAdd int64  `json:"numRepoPrivateAdd"`
	NumRepoForkAdd    int64  `json:"numRepoForkAdd"`
	NumRepoMirrorAdd  int64  `json:"numRepoMirrorAdd"`
	NumRepoSelfAdd    int64  `json:"numRepoSelfAdd"`
	NumRepos          int64  `json:"numRepos"`
	CreatTime         string `json:"creatTime"`
}

type ProjectSummaryData struct {
	ProjectSummaryBaseData
	NumRepoPublic  int64 `json:"numRepoPublic"`
	NumRepoPrivate int64 `json:"numRepoPrivate"`
	NumRepoFork    int64 `json:"numRepoFork"`
	NumRepoMirror  int64 `json:"numRepoMirror"`
	NumRepoSelf    int64 `json:"numRepoSelf"`

	NumRepoOrgAdd    int64 `json:"numRepoOrgAdd"`
	NumRepoNotOrgAdd int64 `json:"numRepoNotOrgAdd"`

	NumRepoOrg    int64 `json:"numRepoOrg"`
	NumRepoNotOrg int64 `json:"numRepoNotOrg"`
}

type ProjectSummaryPeriodData struct {
	RecordBeginTime string                    `json:"recordBeginTime"`
	TotalCount      int64                     `json:"totalCount"`
	PageRecords     []*ProjectSummaryBaseData `json:"pageRecords"`
}

func RestoreForkNumber(ctx *context.Context) {
	repos, err := models.GetAllRepositories()
	if err != nil {
		log.Error("GetAllRepositories failed: %v", err.Error())
		return
	}
	for _, repo := range repos {
		models.RestoreRepoStatFork(int64(repo.NumForks), repo.ID)
	}

	ctx.JSON(http.StatusOK, struct{}{})
}

func GetLatestProjectsSummaryData(ctx *context.Context) {
	stat, err := models.GetLatest2SummaryStatistic()
	data := ProjectSummaryData{}
	if err == nil && len(stat) > 0 {
		data.NumRepos = stat[0].NumRepos
		data.NumRepoOrg = stat[0].NumRepoOrg
		data.NumRepoNotOrg = stat[0].NumRepos - stat[0].NumRepoOrg
		data.NumRepoFork = stat[0].NumRepoFork
		data.NumRepoMirror = stat[0].NumRepoMirror
		data.NumRepoSelf = stat[0].NumRepoSelf
		data.NumRepoPrivate = stat[0].NumRepoPrivate
		data.NumRepoPublic = stat[0].NumRepoPublic
		data.CreatTime = stat[0].CreatedUnix.Format(UPDATE_TIME_FORMAT)
		if len(stat) == 2 {
			data.NumReposAdd = stat[0].NumRepos - stat[1].NumRepos
			data.NumRepoOrgAdd = stat[0].NumRepoOrg - stat[1].NumRepoOrg
			data.NumRepoNotOrgAdd = (stat[0].NumRepos - stat[0].NumRepoOrg) - (stat[1].NumRepos - stat[1].NumRepoOrg)
			data.NumRepoForkAdd = stat[0].NumRepoFork - stat[1].NumRepoFork
			data.NumRepoMirrorAdd = stat[0].NumRepoMirror - stat[1].NumRepoMirror
			data.NumRepoSelfAdd = stat[0].NumRepoSelf - stat[1].NumRepoSelf
			data.NumRepoPrivateAdd = stat[0].NumRepoPrivate - stat[1].NumRepoPrivate
			data.NumRepoPublicAdd = stat[0].NumRepoPublic - stat[1].NumRepoPublic
		}
	}
	ctx.JSON(200, data)
}

func GetProjectsSummaryData(ctx *context.Context) {

	var datas = make([]*ProjectSummaryBaseData, 0)

	recordBeginTime, err := getGrowthRecordBeginTime()
	if err != nil {
		log.Error("Can not get record begin time", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.record_begintime_get_err"))
		return
	}
	beginTime, endTime, err := getTimePeroid(ctx, recordBeginTime)

	beginTime = beginTime.AddDate(0, 0, -1)

	queryType := ctx.QueryTrim("type")

	var count int64

	if queryType == "all" || queryType == "current_year" {
		dates := getEndOfMonthDates(beginTime, endTime)
		count, _ = models.GetSummaryStatisticByDateCount(dates)
		stats, err := models.GetAllSummaryStatisticByDates(dates)
		if err != nil {
			log.Warn("can not get summary data", err)
		} else {

			for i, v := range stats {
				if i == 0 {
					continue
				}
				data := ProjectSummaryBaseData{}
				setStatisticsData(&data, v, stats[i-1])
				createTime, _ := time.Parse(DATE_FORMAT, v.Date)
				data.CreatTime = createTime.Format(MONTH_FORMAT)
				datas = append(datas, &data)
			}
		}

	} else {
		count, _ = models.GetSummaryStatisticByTimeCount(beginTime, endTime)
		stats, err := models.GetAllSummaryStatisticByTime(beginTime, endTime)
		if err != nil {
			log.Warn("can not get summary data", err)
		} else {

			for i, v := range stats {
				if i == 0 {
					continue
				}
				data := ProjectSummaryBaseData{}
				setStatisticsData(&data, v, stats[i-1])
				data.CreatTime = v.Date
				datas = append(datas, &data)
			}
		}

	}

	projectSummaryPeriodData := ProjectSummaryPeriodData{
		TotalCount:      count - 1,
		RecordBeginTime: recordBeginTime.Format(DATE_FORMAT),
		PageRecords:     reverse(datas),
	}

	ctx.JSON(200, projectSummaryPeriodData)

}

func reverse(datas []*ProjectSummaryBaseData) []*ProjectSummaryBaseData {
	for i := 0; i < len(datas)/2; i++ {
		j := len(datas) - i - 1
		datas[i], datas[j] = datas[j], datas[i]
	}
	return datas
}

func setStatisticsData(data *ProjectSummaryBaseData, v *models.SummaryStatistic, stats *models.SummaryStatistic) {
	data.NumReposAdd = v.NumRepos - stats.NumRepos
	data.NumRepoPublicAdd = v.NumRepoPublic - stats.NumRepoPublic
	data.NumRepoPrivateAdd = v.NumRepoPrivate - stats.NumRepoPrivate
	data.NumRepoMirrorAdd = v.NumRepoMirror - stats.NumRepoMirror
	data.NumRepoForkAdd = v.NumRepoFork - stats.NumRepoFork
	data.NumRepoSelfAdd = v.NumRepoSelf - stats.NumRepoSelf

	data.NumRepos = v.NumRepos
}

func getEndOfMonthDates(beginTime time.Time, endTime time.Time) []string {
	var dates = []string{}
	date := endOfMonth(beginTime.AddDate(0, -1, 0))
	dates = append(dates, date.Format(DATE_FORMAT))

	tempDate := endOfMonth(beginTime)

	for {
		if tempDate.Before(endTime) {
			dates = append(dates, tempDate.Format(DATE_FORMAT))
			tempDate = endOfMonth(tempDate.AddDate(0, 0, 1))
		} else {
			break
		}
	}

	dates = append(dates, endTime.AddDate(0, 0, -1).Format(DATE_FORMAT))

	return dates
}

func endOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}

func GetAllProjectsPeriodStatistics(ctx *context.Context) {

	recordBeginTime, err := getRecordBeginTime()
	if err != nil {
		log.Error("Can not get record begin time", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.record_begintime_get_err"))
		return
	}
	beginTime, endTime, err := getTimePeroid(ctx, recordBeginTime)
	if err != nil {
		log.Error("Parameter is wrong", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.parameter_is_wrong"))
		return
	}
	q := ctx.QueryTrim("q")
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := ctx.QueryInt("pagesize")
	if pageSize <= 0 {
		pageSize = DEFAULT_PAGE_SIZE
	}
	orderBy := getOrderBy(ctx)

	latestUpdatedTime, latestDate, err := models.GetRepoStatLastUpdatedTime()
	if err != nil {
		log.Error("Can not query the last updated time.", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.last_update_time_error"))
		return
	}

	countSql := generateCountSql(beginTime, endTime, latestDate, q)
	total, err := models.CountRepoStatByRawSql(countSql)
	if err != nil {
		log.Error("Can not query total count.", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.total_count_get_error"))
		return
	}
	sql := generateSqlByType(ctx, beginTime, endTime, latestDate, q, orderBy, page, pageSize)

	projectsPeriodData := ProjectsPeriodData{
		RecordBeginTime: recordBeginTime.Format(DATE_FORMAT),
		PageSize:        pageSize,
		TotalPage:       util.GetTotalPage(total, pageSize),
		TotalCount:      total,
		LastUpdatedTime: latestUpdatedTime,
		PageRecords:     models.GetRepoStatisticByRawSql(sql),
	}

	ctx.JSON(http.StatusOK, projectsPeriodData)

}

func generateSqlByType(ctx *context.Context, beginTime time.Time, endTime time.Time, latestDate string, q string, orderBy string, page int, pageSize int) string {
	sql := ""
	if ctx.QueryTrim("type") == "all" {
		sql = generateTypeAllSql(beginTime, endTime, latestDate, q, orderBy, page, pageSize)
	} else {
		sql = generatePageSql(beginTime, endTime, latestDate, q, orderBy, page, pageSize)
	}
	return sql
}

func ServeAllProjectsPeriodStatisticsFile(ctx *context.Context) {

	recordBeginTime, err := getRecordBeginTime()
	if err != nil {
		log.Error("Can not get record begin time", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.record_begintime_get_err"))
		return
	}
	beginTime, endTime, err := getTimePeroid(ctx, recordBeginTime)
	if err != nil {
		log.Error("Parameter is wrong", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.parameter_is_wrong"))
		return
	}
	q := ctx.QueryTrim("q")
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := 1000
	orderBy := getOrderBy(ctx)

	_, latestDate, err := models.GetRepoStatLastUpdatedTime()
	if err != nil {
		log.Error("Can not query the last updated time.", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.last_update_time_error"))
		return
	}

	countSql := generateCountSql(beginTime, endTime, latestDate, q)
	total, err := models.CountRepoStatByRawSql(countSql)
	if err != nil {
		log.Error("Can not query total count.", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.total_count_get_error"))
		return
	}
	var projectAnalysis = ctx.Tr("repo.repo_stat_inspect")
	fileName := getFileName(ctx, beginTime, endTime, projectAnalysis)

	totalPage := util.GetTotalPage(total, pageSize)

	f := excelize.NewFile()

	index := f.NewSheet(projectAnalysis)
	f.DeleteSheet("Sheet1")

	for k, v := range allProjectsPeroidHeader(ctx) {
		f.SetCellValue(projectAnalysis, k, v)
	}

	var row = 2
	for i := 0; i <= totalPage; i++ {

		pageRecords := models.GetRepoStatisticByRawSql(generateSqlByType(ctx, beginTime, endTime, latestDate, q, orderBy, i+1, pageSize))
		for _, record := range pageRecords {

			for k, v := range allProjectsPeroidValues(row, record, ctx) {
				f.SetCellValue(projectAnalysis, k, v)
			}
			row++

		}

	}
	f.SetActiveSheet(index)

	ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+url.QueryEscape(fileName))
	ctx.Resp.Header().Set("Content-Type", "application/octet-stream")

	f.WriteTo(ctx.Resp)

}

func GetProjectsSummaryDataFile(ctx *context.Context) {

	recordBeginTime, err := getGrowthRecordBeginTime()
	if err != nil {
		log.Error("Can not get record begin time", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.record_begintime_get_err"))
		return
	}
	beginTime, endTime, err := getTimePeroid(ctx, recordBeginTime)
	beginTime = beginTime.AddDate(0, 0, -1)
	if err != nil {
		log.Error("Parameter is wrong", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.parameter_is_wrong"))
		return
	}

	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := 100

	if err != nil {
		log.Error("Can not query the last updated time.", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.last_update_time_error"))
		return
	}

	var projectAnalysis = ctx.Tr("repo.repo_stat_develop")
	fileName := getSummaryFileName(ctx, beginTime, endTime, projectAnalysis)

	f := excelize.NewFile()

	index := f.NewSheet(projectAnalysis)
	f.DeleteSheet("Sheet1")

	for k, v := range allProjectsPeriodSummaryHeader(ctx) {
		f.SetCellValue(projectAnalysis, k, v)
	}

	var total int64
	queryType := ctx.QueryTrim("type")

	var datas = make([]*ProjectSummaryBaseData, 0)

	if queryType == "all" || queryType == "current_year" {
		dates := getEndOfMonthDates(beginTime, endTime)
		total, _ = models.GetSummaryStatisticByDateCount(dates)
		totalPage := util.GetTotalPage(total, pageSize)

		for i := 0; i < totalPage; i++ {

			stats, err := models.GetSummaryStatisticByDates(dates, i+1, pageSize)
			if err != nil {
				log.Warn("can not get summary data", err)
			} else {
				for j, v := range stats {
					if j == 0 {
						continue
					}
					data := ProjectSummaryBaseData{}
					setStatisticsData(&data, v, stats[j-1])
					createTime, _ := time.Parse(DATE_FORMAT, v.Date)
					data.CreatTime = createTime.Format(MONTH_FORMAT)

					datas = append(datas, &data)

				}

			}

		}

	} else {
		total, _ = models.GetSummaryStatisticByTimeCount(beginTime, endTime)
		totalPage := util.GetTotalPage(total, pageSize)

		for i := 0; i < totalPage; i++ {

			stats, err := models.GetSummaryStatisticByTime(beginTime, endTime, i+1, pageSize)
			if err != nil {
				log.Warn("can not get summary data", err)
			} else {
				for j, v := range stats {
					if j == 0 {
						continue
					}
					data := ProjectSummaryBaseData{}
					setStatisticsData(&data, v, stats[j-1])
					data.CreatTime = v.Date
					datas = append(datas, &data)

				}

			}

		}
	}
	row := 2
	datas = reverse(datas)
	for _, data := range datas {
		for k, v := range allProjectsPeriodSummaryValues(row, data, ctx) {
			f.SetCellValue(projectAnalysis, k, v)
		}
		row++
	}

	f.SetActiveSheet(index)

	ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+url.QueryEscape(fileName))
	ctx.Resp.Header().Set("Content-Type", "application/octet-stream")

	f.WriteTo(ctx.Resp)

}

func ServeAllProjectsOpenIStatisticsFile(ctx *context.Context) {

	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := 1000

	_, latestDate, err := models.GetRepoStatLastUpdatedTime()

	if err != nil {
		log.Error("Can not query the last updated time.", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.last_update_time_error"))
		return
	}

	date := ctx.QueryTrim("date")
	if date == "" {
		date = latestDate
	}

	countSql := generateOpenICountSql(date)
	total, err := models.CountRepoStatByRawSql(countSql)
	if err != nil {
		log.Error("Can not query total count.", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.total_count_get_error"))
		return
	}
	var projectAnalysis = ctx.Tr("repo.repo_stat_inspect")
	fileName := "项目分析_OPENI_" + date + ".xlsx"

	totalPage := util.GetTotalPage(total, pageSize)

	f := excelize.NewFile()

	index := f.NewSheet(projectAnalysis)
	f.DeleteSheet("Sheet1")

	for k, v := range allProjectsOpenIHeader() {
		f.SetCellValue(projectAnalysis, k, v)
	}

	var row = 2
	for i := 0; i <= totalPage; i++ {

		pageRecords := models.GetRepoStatisticByRawSql(generateTypeAllOpenISql(date, i+1, pageSize))
		for _, record := range pageRecords {

			for k, v := range allProjectsOpenIValues(row, record, ctx) {
				f.SetCellValue(projectAnalysis, k, v)
			}
			row++

		}

	}
	f.SetActiveSheet(index)

	ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+url.QueryEscape(fileName))
	ctx.Resp.Header().Set("Content-Type", "application/octet-stream")

	f.WriteTo(ctx.Resp)

}

func getFileName(ctx *context.Context, beginTime time.Time, endTime time.Time, projectAnalysis string) string {
	baseName := projectAnalysis + "_"

	if ctx.QueryTrim("q") != "" {
		baseName = baseName + ctx.QueryTrim("q") + "_"
	}
	if ctx.QueryTrim("type") == "all" {
		baseName = baseName + ctx.Tr("repo.all")
	} else if ctx.QueryTrim("type") == "last_month" {
		baseName = baseName + beginTime.AddDate(0, 0, -1).Format(EXCEL_DATE_FORMAT) + "_" + endTime.AddDate(0, 0, -2).Format(EXCEL_DATE_FORMAT)

	} else {
		baseName = baseName + beginTime.AddDate(0, 0, -1).Format(EXCEL_DATE_FORMAT) + "_" + endTime.AddDate(0, 0, -1).Format(EXCEL_DATE_FORMAT)
	}
	frontName := baseName + ".xlsx"
	return frontName
}

func getSummaryFileName(ctx *context.Context, beginTime time.Time, endTime time.Time, projectAnalysis string) string {
	baseName := projectAnalysis + "_"

	if ctx.QueryTrim("type") == "all" {
		baseName = baseName + ctx.Tr("repo.all")
	} else if ctx.QueryTrim("type") == "current_year" {
		baseName = baseName + ctx.Tr("repo.current_year")
	} else if ctx.QueryTrim("type") == "last_month" {
		baseName = baseName + beginTime.Format(EXCEL_DATE_FORMAT) + "_" + endTime.AddDate(0, 0, -2).Format(EXCEL_DATE_FORMAT)
	} else {
		baseName = baseName + beginTime.Format(EXCEL_DATE_FORMAT) + "_" + endTime.AddDate(0, 0, -1).Format(EXCEL_DATE_FORMAT)
	}
	frontName := baseName + ".xlsx"
	return frontName
}

func allProjectsPeroidHeader(ctx *context.Context) map[string]string {

	return map[string]string{"A1": ctx.Tr("admin.repos.id"), "B1": ctx.Tr("admin.repos.projectName"), "C1": ctx.Tr("repo.owner"), "D1": ctx.Tr("admin.repos.isPrivate"), "E1": ctx.Tr("admin.repos.openi"), "F1": ctx.Tr("admin.repos.visit"), "G1": ctx.Tr("admin.repos.download"), "H1": ctx.Tr("admin.repos.pr"), "I1": ctx.Tr("admin.repos.commit"),
		"J1": ctx.Tr("admin.repos.watches"), "K1": ctx.Tr("admin.repos.stars"), "L1": ctx.Tr("admin.repos.forks"), "M1": ctx.Tr("admin.repos.issues"), "N1": ctx.Tr("admin.repos.closedIssues"), "O1": ctx.Tr("admin.repos.contributor"), "P1": ctx.Tr("admin.repos.numDataset"), "Q1": ctx.Tr("admin.repos.numCloudbrain"), "R1": ctx.Tr("admin.repos.numModel"), "S1": ctx.Tr("admin.repos.numModelConvert"), "T1": ctx.Tr("admin.repos.isFork"), "U1": ctx.Tr("admin.repos.isMirror"), "V1": ctx.Tr("admin.repos.create")}

}

func allProjectsPeriodSummaryHeader(ctx *context.Context) map[string]string {

	return map[string]string{"A1": ctx.Tr("repo.date"), "B1": ctx.Tr("repo.repo_add"), "C1": ctx.Tr("repo.repo_total"), "D1": ctx.Tr("repo.repo_public_add"), "E1": ctx.Tr("repo.repo_private_add"), "F1": ctx.Tr("repo.repo_self_add"), "G1": ctx.Tr("repo.repo_fork_add"), "H1": ctx.Tr("repo.repo_mirror_add")}

}

func allProjectsPeriodSummaryValues(row int, rs *ProjectSummaryBaseData, ctx *context.Context) map[string]string {

	return map[string]string{getCellName("A", row): rs.CreatTime, getCellName("B", row): strconv.FormatInt(rs.NumReposAdd, 10), getCellName("C", row): strconv.FormatInt(rs.NumRepos, 10), getCellName("D", row): strconv.FormatInt(rs.NumRepoPublicAdd, 10), getCellName("E", row): strconv.FormatInt(rs.NumRepoPrivateAdd, 10),
		getCellName("F", row): strconv.FormatInt(rs.NumRepoSelfAdd, 10), getCellName("G", row): strconv.FormatInt(rs.NumRepoForkAdd, 10), getCellName("H", row): strconv.FormatInt(rs.NumRepoMirrorAdd, 10),
	}
}

func allProjectsPeroidValues(row int, rs *models.RepoStatistic, ctx *context.Context) map[string]string {

	return map[string]string{getCellName("A", row): strconv.FormatInt(rs.RepoID, 10), getCellName("B", row): rs.DisplayName(), getCellName("C", row): rs.OwnerName, getCellName("D", row): getBoolDisplay(rs.IsPrivate, ctx), getCellName("E", row): strconv.FormatFloat(rs.RadarTotal, 'f', 2, 64),
		getCellName("F", row): strconv.FormatInt(rs.NumVisits, 10), getCellName("G", row): strconv.FormatInt(rs.NumDownloads, 10), getCellName("H", row): strconv.FormatInt(rs.NumPulls, 10), getCellName("I", row): strconv.FormatInt(rs.NumCommits, 10),
		getCellName("J", row): strconv.FormatInt(rs.NumWatches, 10), getCellName("K", row): strconv.FormatInt(rs.NumStars, 10), getCellName("L", row): strconv.FormatInt(rs.NumForks, 10), getCellName("M", row): strconv.FormatInt(rs.NumIssues, 10),
		getCellName("N", row): strconv.FormatInt(rs.NumClosedIssues, 10), getCellName("O", row): strconv.FormatInt(rs.NumContributor, 10), getCellName("P", row): strconv.FormatInt(rs.NumDatasetFile, 10), getCellName("Q", row): strconv.FormatInt(rs.NumCloudbrain, 10), getCellName("R", row): strconv.FormatInt(rs.NumModels, 10), getCellName("S", row): strconv.FormatInt(rs.NumModelConvert, 10), getCellName("T", row): getBoolDisplay(rs.IsFork, ctx), getCellName("U", row): getBoolDisplay(rs.IsMirror, ctx), getCellName("V", row): time.Unix(int64(rs.RepoCreatedUnix), 0).Format(CREATE_TIME_FORMAT),
	}

}

func allProjectsOpenIHeader() map[string]string {

	return map[string]string{"A1": "ID", "B1": "项目名称", "C1": "拥有者", "D1": "私有", "E1": "OpenI指数",
		"F1": "影响力", "G1": "成熟度", "H1": "活跃度", "I1": "项目健康度", "J1": "团队健康度", "K1": "项目发展趋势",
		"L1": "关注数", "M1": "点赞数", "N1": "派生数", "O1": "代码下载量", "P1": "评论数", "Q1": "浏览量", "R1": "已解决任务数", "S1": "版本发布数量", "T1": "有效开发年龄",
		"U1": "数据集", "V1": "模型数", "W1": "百科页面数量", "X1": "提交数", "Y1": "任务数", "Z1": "PR数", "AA1": "版本发布数量", "AB1": "任务完成比例", "AC1": "贡献者数", "AD1": "关键贡献者数",
		"AE1": "新人增长量", "AF1": "代码规模增长量", "AG1": "任务增长量", "AH1": "新人增长量", "AI1": "提交增长量", "AJ1": "评论增长量", "AK1": "迁移", "AL1": "镜像", "AM1": "项目创建时间",
	}

}

func allProjectsOpenIValues(row int, rs *models.RepoStatistic, ctx *context.Context) map[string]string {

	return map[string]string{getCellName("A", row): strconv.FormatInt(rs.RepoID, 10), getCellName("B", row): rs.DisplayName(), getCellName("C", row): rs.OwnerName, getCellName("D", row): getBoolDisplay(rs.IsPrivate, ctx), getCellName("E", row): strconv.FormatFloat(rs.RadarTotal, 'f', 2, 64),
		getCellName("F", row): strconv.FormatFloat(rs.Impact, 'f', 2, 64), getCellName("G", row): strconv.FormatFloat(rs.Completeness, 'f', 2, 64), getCellName("H", row): strconv.FormatFloat(rs.Liveness, 'f', 2, 64), getCellName("I", row): strconv.FormatFloat(rs.ProjectHealth, 'f', 2, 64), getCellName("J", row): strconv.FormatFloat(rs.TeamHealth, 'f', 2, 64), getCellName("K", row): strconv.FormatFloat(rs.Growth, 'f', 2, 64),
		getCellName("L", row): strconv.FormatInt(rs.NumWatches, 10), getCellName("M", row): strconv.FormatInt(rs.NumStars, 10), getCellName("N", row): strconv.FormatInt(rs.NumForks, 10), getCellName("O", row): strconv.FormatInt(rs.NumDownloads, 10),

		getCellName("P", row): strconv.FormatInt(rs.NumComments, 10), getCellName("Q", row): strconv.FormatInt(rs.NumVisits, 10), getCellName("R", row): strconv.FormatInt(rs.NumClosedIssues, 10), getCellName("S", row): strconv.FormatInt(rs.NumVersions, 10),
		getCellName("T", row): strconv.FormatInt(rs.NumDevMonths, 10), getCellName("U", row): strconv.FormatInt(rs.DatasetSize, 10), getCellName("V", row): strconv.FormatInt(rs.NumModels, 10), getCellName("W", row): strconv.FormatInt(rs.NumWikiViews, 10),
		getCellName("X", row): strconv.FormatInt(rs.NumCommits, 10), getCellName("Y", row): strconv.FormatInt(rs.NumIssues, 10), getCellName("Z", row): strconv.FormatInt(rs.NumPulls, 10), getCellName("AA", row): strconv.FormatInt(rs.NumVersions, 10),
		getCellName("AB", row): strconv.FormatFloat(float64(rs.IssueFixedRate), 'f', 2, 64), getCellName("AC", row): strconv.FormatInt(rs.NumContributor, 10), getCellName("AD", row): strconv.FormatInt(rs.NumKeyContributor, 10), getCellName("AE", row): strconv.FormatInt(rs.NumContributorsGrowth, 10),
		getCellName("AF", row): strconv.FormatInt(rs.NumCommitLinesGrowth, 10), getCellName("AG", row): strconv.FormatInt(rs.NumIssuesGrowth, 10), getCellName("AH", row): strconv.FormatInt(rs.NumContributorsGrowth, 10), getCellName("AI", row): strconv.FormatInt(rs.NumCommitsGrowth, 10), getCellName("AJ", row): strconv.FormatInt(rs.NumCommentsGrowth, 10), getCellName("AK", row): getBoolDisplay(rs.IsFork, ctx), getCellName("AL", row): getBoolDisplay(rs.IsMirror, ctx), getCellName("AM", row): time.Unix(int64(rs.RepoCreatedUnix), 0).Format(CREATE_TIME_FORMAT),
	}

}

func getCellName(col string, row int) string {
	return col + strconv.Itoa(row)
}

func getBoolDisplay(value bool, ctx *context.Context) string {
	if value {
		return ctx.Tr("admin.repos.yes")
	} else {
		return ctx.Tr("admin.repos.no")
	}
}

func GetProjectLatestStatistics(ctx *context.Context) {
	repoId := ctx.Params(":id")
	recordBeginTime, err := getRecordBeginTime()
	if err != nil {
		log.Error("Can not get record begin time", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.record_begintime_get_err"))
		return
	}
	latestUpdatedTime, latestDate, err := models.GetRepoStatLastUpdatedTime(repoId)
	repoIdInt, _ := strconv.ParseInt(repoId, 10, 64)
	repoStat, err := models.GetRepoStatisticByDateAndRepoId(latestDate, repoIdInt)
	if err != nil {
		log.Error("Can not get the repo statistics "+repoId, err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.get_repo_stat_error"))
		return
	}

	repository, err := models.GetRepositoryByID(repoIdInt)
	if err != nil {
		log.Error("Can not get the repo info "+repoId, err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.get_repo_info_error"))
		return
	}
	projectLatestData := ProjectLatestData{
		RecordBeginTime:  recordBeginTime.Format(DATE_FORMAT),
		CreatTime:        time.Unix(int64(repository.CreatedUnix), 0).Format(DATE_FORMAT),
		LastUpdatedTime:  latestUpdatedTime,
		OpenI:            repoStat.RadarTotal,
		Comment:          repoStat.NumComments,
		View:             repoStat.NumVisits,
		Download:         repoStat.NumDownloads,
		IssueClosedRatio: repoStat.IssueFixedRate,
		Impact:           repoStat.Impact,
		Completeness:     repoStat.Completeness,
		Liveness:         repoStat.Liveness,
		ProjectHealth:    repoStat.ProjectHealth,
		TeamHealth:       repoStat.TeamHealth,
		Growth:           repoStat.Growth,
		Description:      repository.Description,
	}

	contributors, err := models.GetTop10Contributor(repository.RepoPath())
	if err != nil {
		log.Error("can not get contributors", err)
	}
	users := make([]UserInfo, 0)

	for _, contributor := range contributors {
		mode := repository.GetCollaboratorMode(contributor.UserId)
		if mode == -1 {
			if contributor.IsAdmin {
				mode = int(models.AccessModeAdmin)
			}
			if contributor.UserId == repository.OwnerID {
				mode = int(models.AccessModeOwner)
			}
		}

		pr := models.GetPullCountByUserAndRepoId(repoIdInt, contributor.UserId)
		userInfo := UserInfo{
			User:          contributor.Committer,
			Commit:        contributor.CommitCnt,
			Mode:          mode,
			PR:            pr,
			RelAvatarLink: contributor.RelAvatarLink,
			Email:         contributor.Email,
		}
		users = append(users, userInfo)

	}

	projectLatestData.Top10 = users

	ctx.JSON(http.StatusOK, projectLatestData)

}

func GetProjectPeriodStatistics(ctx *context.Context) {
	repoId := ctx.Params(":id")
	recordBeginTime, err := getRecordBeginTime()
	if err != nil {
		log.Error("Can not get record begin time", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.record_begintime_get_err"))
		return
	}

	repoIdInt, _ := strconv.ParseInt(repoId, 10, 64)

	if err != nil {
		log.Error("Can not get record begin time", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.record_begintime_get_err"))
		return
	}
	beginTime, endTime, err := getTimePeroid(ctx, recordBeginTime)
	isOpenI := ctx.QueryBool("openi")
	var repositorys []*models.RepoStatistic
	if isOpenI {
		repositorys = models.GetRepoStatisticByRawSql(generateRadarSql(beginTime, endTime, repoIdInt))
	} else {
		repositorys = models.GetRepoStatisticByRawSql(generateTargetSql(beginTime, endTime, repoIdInt))
	}
	ctx.JSON(http.StatusOK, repositorys)
}

func generateRadarSql(beginTime time.Time, endTime time.Time, repoId int64) string {
	sql := "SELECT date, impact, completeness, liveness, project_health, team_health, growth, radar_total FROM repo_statistic" +
		" where repo_id=" + strconv.FormatInt(repoId, 10) + " and created_unix >=" + strconv.FormatInt(beginTime.Unix(), 10) +
		" and created_unix<" + strconv.FormatInt(endTime.Unix(), 10) + " order by created_unix"

	return sql
}

func generateTargetSql(beginTime time.Time, endTime time.Time, repoId int64) string {
	sql := "SELECT date, num_visits,num_downloads_added as num_downloads,num_commits_added as num_commits FROM repo_statistic" +
		" where repo_id=" + strconv.FormatInt(repoId, 10) + " and created_unix >=" + strconv.FormatInt(beginTime.Unix(), 10) +
		" and created_unix<" + strconv.FormatInt(endTime.Unix(), 10) + " order by created_unix desc"

	return sql
}

func generateCountSql(beginTime time.Time, endTime time.Time, latestDate string, q string) string {
	countSql := "SELECT count(*) FROM  " +
		"(SELECT repo_id FROM repo_statistic where created_unix >=" + strconv.FormatInt(beginTime.Unix(), 10) +
		" and created_unix<" + strconv.FormatInt(endTime.Unix(), 10) + " group by repo_id) A," +
		"(SELECT repo_id,name,alias,is_private,radar_total from public.repo_statistic where date='" + latestDate + "') B" +
		" where A.repo_id=B.repo_id"
	if q != "" {
		countSql = countSql + " and LOWER(B.alias) like '%" + strings.ToLower(q) + "%'"
	}
	return countSql
}

func generateOpenICountSql(latestDate string) string {
	countSql := "SELECT count(*) FROM  " +
		"public.repo_statistic where date='" + latestDate + "'"

	return countSql
}

func generateTypeAllSql(beginTime time.Time, endTime time.Time, latestDate string, q string, orderBy string, page int, pageSize int) string {
	sql := "SELECT A.repo_id,name,alias,owner_name,is_private,is_mirror,is_fork,repo_created_unix,radar_total,num_watches,num_visits,num_downloads,num_pulls,num_commits,num_stars,num_forks,num_issues,num_closed_issues,num_contributor,num_models,num_model_convert,num_cloudbrain,num_dataset_file FROM  " +
		"(SELECT repo_id,sum(num_visits) as num_visits " +
		" FROM repo_statistic where created_unix >=" + strconv.FormatInt(beginTime.Unix(), 10) +
		" and created_unix<" + strconv.FormatInt(endTime.Unix(), 10) + " group by repo_id) A," +
		"(SELECT repo_id,name,alias,owner_name,is_private,is_mirror,is_fork,repo_created_unix,radar_total,num_watches,num_downloads,num_pulls,num_commits,num_stars,num_forks,num_issues,num_closed_issues,num_contributor,num_models,num_model_convert,num_cloudbrain,num_dataset_file from public.repo_statistic where date='" + latestDate + "') B" +
		" where A.repo_id=B.repo_id"

	if q != "" {
		sql = sql + " and LOWER(alias) like '%" + strings.ToLower(q) + "%'"
	}
	sql = sql + " order by " + orderBy + " desc,repo_id" + " limit " + strconv.Itoa(pageSize) + " offset " + strconv.Itoa((page-1)*pageSize)
	return sql
}

func generateTypeAllOpenISql(latestDate string, page int, pageSize int) string {
	sql := "SELECT id, repo_id, date, num_watches, num_stars, num_forks, num_downloads, num_comments, num_visits, num_closed_issues, num_versions, num_dev_months, repo_size, dataset_size, num_models, num_wiki_views, num_commits, num_issues, num_pulls, issue_fixed_rate, num_contributor, num_key_contributor, num_contributors_growth, num_commits_growth, num_commit_lines_growth, num_issues_growth, num_comments_growth, impact, completeness, liveness, project_health, team_health, growth, radar_total, name,alias, is_private,is_mirror,is_fork,repo_created_unix, owner_name FROM  " +
		" public.repo_statistic where date='" + latestDate + "'"

	sql = sql + " order by radar_total desc,repo_id" + " limit " + strconv.Itoa(pageSize) + " offset " + strconv.Itoa((page-1)*pageSize)
	return sql
}

func generatePageSql(beginTime time.Time, endTime time.Time, latestDate string, q string, orderBy string, page int, pageSize int) string {

	sql := "SELECT A.repo_id,name,alias,owner_name,is_private,is_mirror,is_fork,repo_created_unix,radar_total,num_watches,num_visits,num_downloads,num_pulls,num_commits,num_stars,num_forks,num_issues,num_closed_issues,num_contributor,num_models,num_model_convert,num_cloudbrain,num_dataset_file FROM  " +
		"(SELECT repo_id,sum(num_watches_added) as num_watches,sum(num_visits) as num_visits, sum(num_downloads_added) as num_downloads,sum(num_pulls_added) as num_pulls,sum(num_commits_added) as num_commits,sum(num_stars_added) as num_stars,sum(num_forks_added) num_forks,sum(num_issues_added) as num_issues,sum(num_closed_issues_added) as num_closed_issues,sum(num_contributor_added) as num_contributor,sum(num_models_added) as num_models,sum(num_model_convert_added) as num_model_convert,sum(num_dataset_file_added) as num_dataset_file, sum(num_cloudbrain_added) as num_cloudbrain " +
		" FROM repo_statistic where created_unix >=" + strconv.FormatInt(beginTime.Unix(), 10) +
		" and created_unix<" + strconv.FormatInt(endTime.Unix(), 10) + " group by repo_id) A," +
		"(SELECT repo_id,name,alias,owner_name,is_private,is_mirror,is_fork,repo_created_unix,radar_total from public.repo_statistic where date='" + latestDate + "') B" +
		" where A.repo_id=B.repo_id"
	if q != "" {
		sql = sql + " and LOWER(B.alias) like '%" + strings.ToLower(q) + "%'"
	}
	sql = sql + " order by " + orderBy + " desc,A.repo_id" + " limit " + strconv.Itoa(pageSize) + " offset " + strconv.Itoa((page-1)*pageSize)
	return sql
}

func getOrderBy(ctx *context.Context) string {
	orderBy := ""
	switch ctx.Query("sort") {
	case "openi":
		orderBy = "B.radar_total"
	case "view":
		orderBy = "A.num_visits"
	case "download":
		orderBy = "A.num_downloads"
	case "pr":
		orderBy = "A.num_pulls"
	case "commit":
		orderBy = "A.num_commits"
	case "watch":
		orderBy = "A.num_watches"
	case "star":
		orderBy = "A.num_stars"
	case "fork":
		orderBy = "A.num_forks"
	case "issue":
		orderBy = "A.num_issues"
	case "issue_closed":
		orderBy = "A.num_closed_issues"
	case "contributor":
		orderBy = "A.num_contributor"
	default:
		orderBy = "B.radar_total"
	}
	return orderBy
}

func getTimePeroid(ctx *context.Context, recordBeginTime time.Time) (time.Time, time.Time, error) {
	queryType := ctx.QueryTrim("type")
	now := time.Now()
	recordBeginTimeTemp := recordBeginTime.AddDate(0, 0, 1)

	beginTimeStr := ctx.QueryTrim("beginTime")
	endTimeStr := ctx.QueryTrim("endTime")
	var beginTime time.Time
	var endTime time.Time
	var err error
	if queryType != "" {
		if queryType == "all" {
			beginTime = recordBeginTimeTemp
			endTime = now
		} else if queryType == "last_year" {
			beginTime = time.Date(now.Year()-1, 1, 2, 0, 0, 0, 0, now.Location())
			endTime = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
		} else if queryType == "yesterday" {
			endTime = now
			beginTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, now.Location())

		} else if queryType == "current_week" {
			beginTime = now.AddDate(0, 0, -int(time.Now().Weekday())+2) //begin from monday
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = now
		} else if queryType == "current_month" {
			endTime = now
			beginTime = time.Date(endTime.Year(), endTime.Month(), 2, 0, 0, 0, 0, now.Location())
		} else if queryType == "monthly" {
			endTime = now
			beginTime = now.AddDate(0, -1, 1)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())

		} else if queryType == "current_year" {
			endTime = now
			beginTime = time.Date(endTime.Year(), 1, 2, 0, 0, 0, 0, now.Location())

		} else if queryType == "last_month" {

			lastMonthTime := now.AddDate(0, -1, 0)
			beginTime = time.Date(lastMonthTime.Year(), lastMonthTime.Month(), 2, 0, 0, 0, 0, now.Location())
			endTime = time.Date(now.Year(), now.Month(), 2, 0, 0, 0, 0, now.Location())

		} else {
			return now, now, fmt.Errorf("The value of type parameter is wrong.")

		}

	} else {
		if beginTimeStr == "" || endTimeStr == "" {
			//如果查询类型和开始时间结束时间都未设置，按queryType=all处理
			beginTime = recordBeginTimeTemp
			endTime = now

		} else {

			beginTime, err = time.ParseInLocation("2006-01-02", beginTimeStr, time.Local)
			if err != nil {
				return now, now, err
			}

			endTime, err = time.ParseInLocation("2006-01-02", endTimeStr, time.Local)
			if err != nil {
				return now, now, err
			}

			beginTime = beginTime.AddDate(0, 0, 1)
			endTime = endTime.AddDate(0, 0, 1)
		}

	}

	if beginTime.Before(recordBeginTimeTemp) {
		beginTime = recordBeginTimeTemp
	}

	return beginTime, endTime, nil

}

func getRecordBeginTime() (time.Time, error) {
	return time.ParseInLocation(DATE_FORMAT, setting.RadarMap.RecordBeginTime, time.Local)
}

func getGrowthRecordBeginTime() (time.Time, error) {
	return time.ParseInLocation(DATE_FORMAT, setting.RadarMap.GrowthBeginTime, time.Local)
}

func ProjectNumVisit(ctx *context.APIContext) {
	var (
		err error
	)

	var userName = ctx.Query("user")
	var projectName = ctx.Query("project")
	var beginTime = ctx.Query("begintime")
	var endTime = ctx.Query("endtime")

	var ProjectNumVisits int
	ProjectNumVisits, err = repository.AppointProjectView(userName, projectName, beginTime, endTime) //访问量
	if err != nil {
		ctx.NotFound(err)
	}
	log.Info("ProjectNumVisits is:", ProjectNumVisits)

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"ProjectNumVisits": ProjectNumVisits,
		"StatusOK":         0,
	})
}
