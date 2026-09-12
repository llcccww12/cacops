package repo

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/git"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/services/mailer"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

const (
	PAGE_SIZE       = 2000
	Excel_File_Path = "/useranalysis/"
	USER_YEAR       = 2023
)

type ProjectsPeriodData struct {
	RecordBeginTime string `json:"recordBeginTime"`
	LastUpdatedTime string `json:"lastUpdatedTime"`
}

func getUserMetricsExcelHeader(ctx *context.Context) map[string]string {
	excelHeader := make([]string, 0)
	excelHeader = append(excelHeader, ctx.Tr("user.metrics.date"))
	excelHeader = append(excelHeader, ctx.Tr("user.metrics.newregistuser"))
	excelHeader = append(excelHeader, ctx.Tr("user.metrics.newregistandactiveuser"))
	excelHeader = append(excelHeader, ctx.Tr("user.metrics.hasactivateuser"))
	excelHeader = append(excelHeader, ctx.Tr("user.metrics.newregistnotactiveuser"))
	excelHeader = append(excelHeader, ctx.Tr("user.metrics.newuseractiveindex"))
	excelHeader = append(excelHeader, ctx.Tr("user.metrics.currentdayactivity"))
	excelHeader = append(excelHeader, ctx.Tr("user.metrics.totalregistuser"))
	excelHeader = append(excelHeader, ctx.Tr("user.metrics.totalactiveduser"))
	excelHeader = append(excelHeader, ctx.Tr("user.metrics.totalhasactivityuser"))

	excelHeaderMap := make(map[string]string, 0)
	var i byte
	i = 0
	for _, value := range excelHeader {
		excelColumn := getColumn(i) + fmt.Sprint(1)
		excelHeaderMap[excelColumn] = value
		i++
	}
	return excelHeaderMap
}

func writeUserMetricsExcel(row int, xlsx *excelize.File, sheetName string, userMetrics *models.UserMetrics) {
	rows := fmt.Sprint(row)
	var tmp byte
	tmp = 0
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userMetrics.DisplayDate)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userMetrics.ActivateRegistUser+userMetrics.NotActivateRegistUser)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userMetrics.ActivateRegistUser)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userMetrics.RegistActivityUser)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userMetrics.NotActivateRegistUser)
	tmp = tmp + 1
	t := userMetrics.ActivateIndex * 100
	value := "-"
	if t < 100 && t > 0 {
		value = fmt.Sprintf("%.2f", t) + "%"
	} else if t >= 100 {
		value = "100%"
	}
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, value)
	tmp = tmp + 1

	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userMetrics.HasActivityUser)
	tmp = tmp + 1

	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userMetrics.TotalUser)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userMetrics.TotalActivateRegistUser)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userMetrics.TotalHasActivityUser)
}

func getExcelHeader(ctx *context.Context) map[string]string {
	excelHeader := make([]string, 0)
	excelHeader = append(excelHeader, ctx.Tr("user.static.id"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.name"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.UserIndex"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.UserIndexPrimitive"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.codemergecount"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.commitcount"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.issuecount"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.commentcount"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.focusrepocount"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.starrepocount"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.logincount"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.loginactioncount"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.watchedcount"))
	//excelHeader = append(excelHeader, ctx.Tr("user.static.commitcodesize"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.solveissuecount"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.encyclopediascount"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.createrepocount"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.openiindex"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.CloudBrainTaskNum"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.CloudBrainRunTime"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.CommitDatasetNum"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.CommitModelCount"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.ModelConvertCount"))

	excelHeader = append(excelHeader, ctx.Tr("user.static.FocusOtherUser"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.CollectDataset"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.CollectedDataset"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.RecommendDataset"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.CollectImage"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.CollectedImage"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.RecommendImage"))

	excelHeader = append(excelHeader, ctx.Tr("user.static.email"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.occupation"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.phone"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.location"))

	excelHeader = append(excelHeader, ctx.Tr("user.static.registdate"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.countdate"))

	excelHeaderMap := make(map[string]string, 0)
	var i byte
	i = 0
	for _, value := range excelHeader {
		excelColumn := getColumn(i) + fmt.Sprint(1)
		log.Info("excelColumn=" + excelColumn)
		excelHeaderMap[excelColumn] = value
		i++
	}
	return excelHeaderMap
}

func writeExcel(row int, xlsx *excelize.File, sheetName string, userRecord *models.UserBusinessAnalysisAll) {
	rows := fmt.Sprint(row)
	var tmp byte
	tmp = 0
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.ID)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.Name)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, fmt.Sprintf("%.2f", userRecord.UserIndex))
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, fmt.Sprintf("%.2f", userRecord.UserIndexPrimitive))
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CodeMergeCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CommitCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.IssueCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CommentCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.FocusRepoCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.StarRepoCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.LoginCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.LoginActionCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.WatchedCount)
	//tmp = tmp + 1
	//xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CommitCodeSize)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.SolveIssueCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.EncyclopediasCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CreateRepoCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, fmt.Sprintf("%.2f", userRecord.OpenIIndex))
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CloudBrainTaskNum)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, fmt.Sprintf("%.2f", float64(userRecord.CloudBrainRunTime)/3600))
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CommitDatasetNum)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CommitModelCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.ModelConvertCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.FocusOtherUser)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CollectDataset)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CollectedDataset)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.RecommendDataset)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CollectImage)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CollectedImage)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.RecommendImage)
	tmp = tmp + 1

	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.Email)
	tmp = tmp + 1

	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, ConvertOccupationToString(userRecord.Occupation))
	tmp = tmp + 1

	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.Phone)
	tmp = tmp + 1

	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.UserLocation)
	tmp = tmp + 1

	formatTime := userRecord.RegistDate.Format("2006-01-02 15:04:05")
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, formatTime[0:len(formatTime)-3])
	tmp = tmp + 1

	formatTime = userRecord.DataDate
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, formatTime)
}

func writeExcelPage(row int, xlsx *excelize.File, sheetName string, userRecord *models.UserBusinessAnalysis) {
	rows := fmt.Sprint(row)
	var tmp byte
	tmp = 0
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.ID)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.Name)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, fmt.Sprintf("%.2f", userRecord.UserIndex))
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, fmt.Sprintf("%.2f", userRecord.UserIndexPrimitive))
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CodeMergeCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CommitCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.IssueCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CommentCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.FocusRepoCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.StarRepoCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.LoginCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.LoginActionCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.WatchedCount)
	//tmp = tmp + 1
	//xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CommitCodeSize)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.SolveIssueCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.EncyclopediasCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CreateRepoCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, fmt.Sprintf("%.2f", userRecord.OpenIIndex))
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CloudBrainTaskNum)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, fmt.Sprintf("%.2f", float64(userRecord.CloudBrainRunTime)/3600))
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CommitDatasetNum)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CommitModelCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.ModelConvertCount)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.FocusOtherUser)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CollectDataset)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CollectedDataset)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.RecommendDataset)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CollectImage)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CollectedImage)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.RecommendImage)
	tmp = tmp + 1

	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.Email)
	tmp = tmp + 1

	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, ConvertOccupationToString(userRecord.Occupation))
	tmp = tmp + 1

	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.Phone)
	tmp = tmp + 1

	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.UserLocation)
	tmp = tmp + 1

	formatTime := userRecord.RegistDate.Format("2006-01-02 15:04:05")
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, formatTime[0:len(formatTime)-3])
	tmp = tmp + 1

	formatTime = userRecord.DataDate
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, formatTime)
}

func getColumn(tmp byte) string {
	var tmpA byte
	tmpA = 'A'
	if tmp < 26 {
		return string(tmpA + tmp)
	} else {
		return "A" + string(tmpA+(tmp-26))
	}
}

func queryUserDataPage(ctx *context.Context, tableName string, queryObj interface{}) {
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := ctx.QueryInt("pageSize")
	if pageSize <= 0 {
		pageSize = setting.UI.IssuePagingNum
	}
	userName := ctx.Query("userName")
	IsReturnFile := ctx.QueryBool("IsReturnFile")

	if IsReturnFile {
		//writer exec file.
		xlsx := excelize.NewFile()
		sheetName := ctx.Tr("user.static.sheetname")
		index := xlsx.NewSheet(sheetName)
		xlsx.DeleteSheet("Sheet1")
		dataHeader := getExcelHeader(ctx)
		for k, v := range dataHeader {
			//设置单元格的值
			xlsx.SetCellValue(sheetName, k, v)
		}
		_, count := models.QueryUserStaticDataByTableName(1, 1, tableName, queryObj, userName)
		var indexTotal int64
		indexTotal = 0
		row := 1
		for {
			re, _ := models.QueryUserStaticDataByTableName(int(indexTotal), PAGE_SIZE, tableName, queryObj, "")
			log.Info("return count=" + fmt.Sprint(count))
			for _, userRecord := range re {
				row++
				writeExcel(row, xlsx, sheetName, userRecord)
			}

			indexTotal += PAGE_SIZE
			if indexTotal >= count {
				break
			}
		}
		//设置默认打开的表单
		xlsx.SetActiveSheet(index)
		filename := sheetName + "_" + ctx.Tr("user.static."+tableName) + ".xlsx"
		ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+url.QueryEscape(filename))
		ctx.Resp.Header().Set("Content-Type", "application/octet-stream")
		if _, err := xlsx.WriteTo(ctx.Resp); err != nil {
			log.Info("writer exel error." + err.Error())
		}
	} else {
		re, count := models.QueryUserStaticDataByTableName((page-1)*pageSize, pageSize, tableName, queryObj, userName)
		mapInterface := make(map[string]interface{})
		mapInterface["data"] = re
		mapInterface["count"] = count

		ctx.JSON(http.StatusOK, mapInterface)
	}
}

func queryMetrics(ctx *context.Context, tableName string, startTime time.Time, endTime time.Time) {

	IsReturnFile := ctx.QueryBool("IsReturnFile")

	var count int64
	result := make([]*models.UserMetrics, 0)
	if tableName == "public.user_business_analysis_current_year" {
		result = models.QueryMetricsForYear(startTime.Unix(), endTime.Unix())
		count = int64(len(result))
	} else if tableName == "public.user_business_analysis_all" {
		result = models.QueryMetricsForAll(startTime.Unix(), endTime.Unix())
		count = int64(len(result))
	} else {
		result, count = models.QueryMetricsPage(startTime.Unix(), endTime.Unix())
	}
	if IsReturnFile {
		//writer exec file.
		xlsx := excelize.NewFile()
		sheetName := ctx.Tr("user.metrics.sheetname")
		index := xlsx.NewSheet(sheetName)
		xlsx.DeleteSheet("Sheet1")
		dataHeader := getUserMetricsExcelHeader(ctx)
		for k, v := range dataHeader {
			//设置单元格的值
			xlsx.SetCellValue(sheetName, k, v)
		}
		row := 1
		log.Info("return count=" + fmt.Sprint(count))
		for _, userRecord := range result {
			row++
			writeUserMetricsExcel(row, xlsx, sheetName, userRecord)
		}
		//设置默认打开的表单
		xlsx.SetActiveSheet(index)
		filename := sheetName + "_" + ctx.Tr("user.static."+tableName) + ".xlsx"
		if tableName == "" {
			filename = sheetName + "_" + getTimeFileName(startTime) + "_" + getTimeFileName(endTime) + ".xlsx"
		}
		ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+url.QueryEscape(filename))
		ctx.Resp.Header().Set("Content-Type", "application/octet-stream")
		if _, err := xlsx.WriteTo(ctx.Resp); err != nil {
			log.Info("writer exel error." + err.Error())
		}
	} else {
		mapInterface := make(map[string]interface{})
		mapInterface["data"] = result
		mapInterface["count"] = count
		if tableName == "public.user_business_analysis_yesterday" {
			mapInterface["datarecordbegintime"] = setting.RadarMap.GrowthBeginTime
			if len(result) > 0 {
				dateTime := time.Unix(result[0].CountDate, 0).AddDate(0, 0, 1)
				mapInterface["lastUpdatedTime"] = dateTime.Format("2006-01-02 15:04:05")
			} else {
				mapInterface["lastUpdatedTime"] = ""
			}
		}
		ctx.JSON(http.StatusOK, mapInterface)
	}
}
func getTimeFileName(t time.Time) string {
	t = t.Local()
	return t.Format("20060102")
}

func QueryRankingList(ctx *context.Context) {
	key := ctx.Query("key")
	tableName := ctx.Query("tableName")
	limit := ctx.QueryInt("limit")

	result, count := models.QueryRankList(key, tableName, limit)
	mapInterface := make(map[string]interface{})
	mapInterface["data"] = result
	mapInterface["count"] = count
	ctx.JSON(http.StatusOK, mapInterface)
}

func DownloadUserDefineFile(ctx *context.Context) {
	filename := ctx.Query("filename")
	length := len(filename)
	if filename[0:1] == "\"" {
		filename = filename[1 : length-1]
	}
	allFilename := setting.AppDataPath + Excel_File_Path + filename
	log.Info("allFilename=" + allFilename)
	_, err := os.Stat(allFilename)
	if err != nil { //文件不存在
		log.Info("file not exist.")
		ctx.JSON(http.StatusOK, "File Not Exist.")
	} else {
		ctx.ServeFile(allFilename, url.QueryEscape(filename))
	}
}

func QueryUserMetricsCurrentMonth(ctx *context.Context) {

	currentTimeNow := time.Now()
	pageEndTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 0, 0, 0, 0, currentTimeNow.Location())
	pageStartTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), 1, 0, 0, 0, 0, currentTimeNow.Location())
	pageStartTime = getStartTime(pageStartTime)
	queryMetrics(ctx, "public.user_business_analysis_current_month", pageStartTime, pageEndTime)
}
func QueryUserStaticCurrentMonth(ctx *context.Context) {
	queryUserDataPage(ctx, "public.user_business_analysis_current_month", new(models.UserBusinessAnalysisCurrentMonth))
}

func getStartTime(pageStartTime time.Time) time.Time {
	t, _ := time.ParseInLocation("2006-01-02", setting.RadarMap.GrowthBeginTime, time.Local)
	t = t.UTC()
	if pageStartTime.Before(t) {
		pageStartTime = t
	}
	return pageStartTime
}

func QueryUserMetricsCurrentWeek(ctx *context.Context) {
	currentTimeNow := time.Now()
	offset := int(time.Monday - currentTimeNow.Weekday())
	if offset > 0 {
		offset = -6
	}
	pageStartTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	pageStartTime = getStartTime(pageStartTime)
	pageEndTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 0, 0, 0, 0, currentTimeNow.Location())
	queryMetrics(ctx, "public.user_business_analysis_current_week", pageStartTime, pageEndTime)
}
func QueryUserStaticCurrentWeek(ctx *context.Context) {
	queryUserDataPage(ctx, "public.user_business_analysis_current_week", new(models.UserBusinessAnalysisCurrentWeek))
}
func QueryUserStaticLastWeek(ctx *context.Context) {
	queryUserDataPage(ctx, "public.user_business_analysis_last_week", new(models.UserBusinessAnalysisLastWeek))
}

func QueryUserMetricsCurrentYear(ctx *context.Context) {
	currentTimeNow := time.Now()
	pageStartTime := time.Date(currentTimeNow.Year(), 1, 1, 0, 0, 0, 0, currentTimeNow.Location())
	pageStartTime = getStartTime(pageStartTime)
	pageEndTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 0, 0, 0, 0, currentTimeNow.Location())
	queryMetrics(ctx, "public.user_business_analysis_current_year", pageStartTime, pageEndTime)
}
func QueryUserStaticCurrentYear(ctx *context.Context) {
	queryUserDataPage(ctx, "public.user_business_analysis_current_year", new(models.UserBusinessAnalysisCurrentYear))
}
func QueryUserMetricsLast30Day(ctx *context.Context) {
	currentTimeNow := time.Now()
	pageStartTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, -30)
	pageStartTime = getStartTime(pageStartTime)
	pageEndTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 0, 0, 0, 0, currentTimeNow.Location())
	queryMetrics(ctx, "public.user_business_analysis_last30_day", pageStartTime, pageEndTime)
}
func QueryUserStaticLast30Day(ctx *context.Context) {
	queryUserDataPage(ctx, "public.user_business_analysis_last30_day", new(models.UserBusinessAnalysisLast30Day))
}
func QueryUserMetricsLastMonth(ctx *context.Context) {
	currentTimeNow := time.Now()
	thisMonth := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), 1, 0, 0, 0, 0, currentTimeNow.Location())
	pageStartTime := thisMonth.AddDate(0, -1, 0)
	pageStartTime = getStartTime(pageStartTime)
	pageEndTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), 1, 23, 59, 59, 0, currentTimeNow.Location()).AddDate(0, 0, -1)
	queryMetrics(ctx, "public.user_business_analysis_last_month", pageStartTime, pageEndTime)
}
func QueryUserStaticLastMonth(ctx *context.Context) {
	queryUserDataPage(ctx, "public.user_business_analysis_last_month", new(models.UserBusinessAnalysisLastMonth))
}
func QueryUserMetricsYesterday(ctx *context.Context) {
	currentTimeNow := time.Now().AddDate(0, 0, -1)
	pageStartTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 0, 0, 0, 0, time.Local)
	pageStartTime = getStartTime(pageStartTime)
	pageEndTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 23, 59, 59, 0, currentTimeNow.Location())
	queryMetrics(ctx, "public.user_business_analysis_yesterday", pageStartTime, pageEndTime)
}
func QueryUserStaticYesterday(ctx *context.Context) {
	queryUserDataPage(ctx, "public.user_business_analysis_yesterday", new(models.UserBusinessAnalysisYesterday))
}
func QueryUserMetricsAll(ctx *context.Context) {
	currentTimeNow := time.Now()
	pageStartTime := time.Date(2022, 4, 5, 0, 0, 0, 0, currentTimeNow.Location())
	pageStartTime = getStartTime(pageStartTime)
	pageEndTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 0, 0, 0, 0, currentTimeNow.Location())
	queryMetrics(ctx, "public.user_business_analysis_all", pageStartTime, pageEndTime)
}
func QueryUserStaticAll(ctx *context.Context) {
	queryUserDataPage(ctx, "public.user_business_analysis_all", new(models.UserBusinessAnalysisAll))
}

func QueryUserMetricDataPage(ctx *context.Context) {
	startDate := ctx.Query("startDate")
	endDate := ctx.Query("endDate")
	startTime, _ := time.ParseInLocation("2006-01-02", startDate, time.Local)
	startTime = startTime.UTC()
	endTime, _ := time.ParseInLocation("2006-01-02", endDate, time.Local)
	startTime = getStartTime(startTime)
	queryMetrics(ctx, "", startTime, endTime)
}

func QueryUserStaticDataPage(ctx *context.Context) {
	startDate := ctx.Query("startDate")
	endDate := ctx.Query("endDate")
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := ctx.QueryInt("pageSize")
	if pageSize <= 0 {
		pageSize = setting.UI.IssuePagingNum
	}
	userName := ctx.Query("userName")
	IsReturnFile := ctx.QueryBool("IsReturnFile")
	log.Info("startDate=" + startDate + "  endDate=" + endDate + " userName=" + userName + " page=" + fmt.Sprint(page))
	var startTime time.Time
	var endTime time.Time
	var isAll bool
	if startDate == "all" {
		isAll = true
		startTime = time.Now()
		endTime = time.Now()
	} else {
		startTime, _ = time.ParseInLocation("2006-01-02", startDate, time.Local)
		startTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, startTime.Location())
		settingStartTime, _ := time.Parse("2006-01-02", setting.RadarMap.RecordBeginTime)
		if startTime.Unix() < settingStartTime.Unix() {
			startTime = settingStartTime
			startDate = settingStartTime.Format("2006-01-02")
		}
		endTime, _ = time.ParseInLocation("2006-01-02", endDate, time.Local)
		endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 23, 59, 59, 0, startTime.Location())

		isAll = false
		log.Info("startTime=" + fmt.Sprint(startTime.Unix()) + "  endDate=" + fmt.Sprint(endTime.Unix()))
	}

	if IsReturnFile {
		page = -1
		pageSize = -1
	}

	pageOpts := &models.UserBusinessAnalysisQueryOptions{
		ListOptions: models.ListOptions{
			Page:     page,
			PageSize: pageSize,
		},
		UserName:  userName,
		StartTime: startTime.Unix(),
		EndTime:   endTime.Unix(),
		IsAll:     isAll,
	}

	if IsReturnFile {
		//re, count := models.QueryUserStaticDataAll(pageOpts)
		wikiMap, _ := queryWikiCountMap(startTime, endTime)
		re, count := models.QueryUserStaticDataForUserDefine(pageOpts, wikiMap)
		sheetName := ctx.Tr("user.static.sheetname")
		filename := sheetName + "_" + startDate + "_" + endDate + ".xlsx"
		os.Remove(setting.AppDataPath + Excel_File_Path + filename)
		go writeFileToDisk(ctx, count, re, filename)
		ctx.JSON(http.StatusOK, ctx.Tr("user.static.downloadinfo")+"/api/v1/download_user_define_file?filename="+filename)
	} else {
		mapInterface := make(map[string]interface{})
		key := startTime.Format("2006-01-02") + endTime.Format("2006-01-02")
		log.Info("db key =" + key)
		re, count := models.QueryDataForUserDefineFromDb(pageOpts, key)
		if count == 0 {
			wikiMap, _ := queryWikiCountMap(startTime, endTime)
			re, count = models.QueryUserStaticDataForUserDefine(pageOpts, wikiMap)
			models.WriteDataToDb(re, key)
		}
		re, count = models.QueryDataForUserDefineFromDb(pageOpts, key)
		mapInterface["data"] = re
		mapInterface["count"] = count
		ctx.JSON(http.StatusOK, mapInterface)
	}
}

func writeFileToDisk(ctx *context.Context, count int64, re []*models.UserBusinessAnalysis, filename string) {
	log.Info("return count=" + fmt.Sprint(count))
	//writer exec file.
	xlsx := excelize.NewFile()
	sheetName := ctx.Tr("user.static.sheetname")
	index := xlsx.NewSheet(sheetName)
	xlsx.DeleteSheet("Sheet1")

	dataHeader := getExcelHeader(ctx)
	for k, v := range dataHeader {
		//设置单元格的值
		xlsx.SetCellValue(sheetName, k, v)
	}

	for i, userRecord := range re {
		row := i + 2
		writeExcelPage(row, xlsx, sheetName, userRecord)
	}

	//设置默认打开的表单
	xlsx.SetActiveSheet(index)

	//ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+url.QueryEscape(filename))
	//ctx.Resp.Header().Set("Content-Type", "application/octet-stream")
	filename = setting.AppDataPath + Excel_File_Path + filename
	os.Mkdir(setting.AppDataPath+Excel_File_Path, 0755)
	if err := xlsx.SaveAs(filename); err != nil {
		log.Info("writer exel error." + err.Error())
	} else {
		log.Info("write to file succeed, filepath=" + filename)
	}
}

func queryWikiCountMap(startTime time.Time, endTime time.Time) (map[string]int, error) {
	wikiMap := make(map[string]int)
	warnEmailMessage := "用户统计信息入库失败，请尽快定位。"
	repoList, err := models.GetAllRepositories()
	if err != nil {
		log.Error("query repo error." + err.Error())
		mailer.SendWarnNotifyMail(setting.Warn_Notify_Mails, warnEmailMessage)
		return nil, err
	}
	log.Info("start to query wiki data")
	for _, repoRecord := range repoList {
		wikiPath := models.WikiPath(repoRecord.OwnerName, repoRecord.Name)
		time, err := git.GetLatestCommitTime(wikiPath)
		if err == nil {
			log.Info("last commit time:" + time.Format("2006-01-02 15:04:05") + " wikiPath=" + wikiPath)
			if time.After(startTime) && time.Before(endTime) {
				wikiRepo, _, err := FindWikiRepoCommitByWikiPath(wikiPath)
				if err != nil {
					log.Error("wiki not exist. wikiPath=" + wikiPath)
				} else {
					log.Info("wiki exist, wikiPath=" + wikiPath)
					list, err := wikiRepo.GetCommitByPathAndDays(wikiPath, 1)
					if err != nil {
						log.Info("err,err=v%", err)
					} else {
						for logEntry := list.Front(); logEntry != nil; logEntry = logEntry.Next() {
							commit := logEntry.Value.(*git.Commit)
							log.Info("commit msg=" + commit.CommitMessage + " time=" + commit.Committer.When.Format("2006-01-02 15:04:05") + " user=" + commit.Committer.Name)
							if _, ok := wikiMap[commit.Committer.Name]; !ok {
								wikiMap[commit.Committer.Name] = 1
							} else {
								wikiMap[commit.Committer.Name] += 1
							}
						}
					}

				}
			}
		}
	}
	return wikiMap, nil
}

func UpdateYearReportCollection() {

	setting.RefreshAnnualCollectionConfig()
	if setting.AnnualCollection.Enabled {
		startYear := time.Date(setting.AnnualCollection.Year, 1, 1, 0, 0, 0, 1, time.Now().Location())
		endYear := startYear.AddDate(1, 0, 0)

		models.RefreshUserYearTable(startYear, endYear)
	}

}

func TimingCountDataByDateAndReCount(date string, isReCount bool) {
	t, _ := time.Parse("2006-01-02", date)

	startTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	startTime = startTime.UTC()
	endTime := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	endTime = endTime.UTC()
	log.Info("startTime time:" + startTime.Format("2006-01-02 15:04:05"))
	log.Info("endTime time:" + endTime.Format("2006-01-02 15:04:05"))
	warnEmailMessage := "用户统计信息入库失败，请尽快定位。"

	//query wiki data
	log.Info("start to time count data")
	wikiMap, err := queryWikiCountMap(startTime, endTime)
	//other user info data
	err = models.CounDataByDateAndReCount(wikiMap, startTime, endTime, isReCount)
	if err != nil {
		log.Error("count user info error." + err.Error())
		mailer.SendWarnNotifyMail(setting.Warn_Notify_Mails, warnEmailMessage)
	}

	setting.RefreshAnnualCollectionConfig()
	if setting.AnnualCollection.Enabled {
		startYear := time.Date(setting.AnnualCollection.Year, 1, 1, 0, 0, 0, 1, t.Location())
		endYear := startYear.AddDate(1, 0, 0)

		models.RefreshUserYearTable(startYear, endYear)
	}

	log.Info("end to count all user info data")
}

func TimingCountDataByDate(date string) {
	TimingCountDataByDateAndReCount(date, true)
}

func TimingCountData() {
	log.Info("start to time count data")
	context.UserActionMapClear()
	currentTimeNow := time.Now()
	log.Info("current time:" + currentTimeNow.Format("2006-01-02 15:04:05"))
	startTime := currentTimeNow.AddDate(0, 0, -1).Format("2006-01-02")
	TimingCountDataByDateAndReCount(startTime, false)
}

func QueryUserActivity(ctx *context.Context) {
	startDate := ctx.Query("beginTime")
	endDate := ctx.Query("endTime")

	t, _ := time.Parse("2006-01-02", startDate)
	startTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	startTime = startTime.UTC()

	t, _ = time.Parse("2006-01-02", endDate)
	endTime := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	endTime = endTime.UTC()

	sheetName := ctx.Tr("user.static.sheetname")
	filename := sheetName + "_" + startDate + "_" + endDate + ".xlsx"
	filePath := setting.AppDataPath + Excel_File_Path + filename
	os.Remove(setting.AppDataPath + Excel_File_Path + filename)

	go writeUserActivityToExcel(startTime, endTime, filePath, ctx)

	ctx.JSON(http.StatusOK, ctx.Tr("user.static.downloadinfo")+"/api/v1/download_user_define_file?filename="+filename)

}

func writeUserActivityToExcel(startTime time.Time, endTime time.Time, filePath string, ctx *context.Context) {
	re := models.QueryDataForActivity(startTime, endTime)
	log.Info("return count=" + fmt.Sprint(len(re)))
	//writer exec file.
	xlsx := excelize.NewFile()
	sheetName := ctx.Tr("user.static.sheetname")
	index := xlsx.NewSheet(sheetName)
	xlsx.DeleteSheet("Sheet1")

	excelHeader := make([]string, 0)
	excelHeader = append(excelHeader, ctx.Tr("user.static.id"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.name"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.codemergecount"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.commitcount"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.issuecount"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.commentcount"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.watchedcount"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.commitcodesize"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.solveissuecount"))

	excelHeader = append(excelHeader, ctx.Tr("user.static.CommitDatasetNum"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.CommitModelCount"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.ModelConvertCount"))
	//model convert

	//new
	excelHeader = append(excelHeader, ctx.Tr("user.static.logincount"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.createrepocount"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.CloudBrainTaskNum"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.CollectDataset"))

	excelHeader = append(excelHeader, ctx.Tr("user.static.email"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.occupation"))

	excelHeader = append(excelHeader, ctx.Tr("user.static.phone"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.registdate"))

	excelHeaderMap := make(map[string]string, 0)
	var j byte
	j = 0
	for _, value := range excelHeader {
		excelColumn := getColumn(j) + fmt.Sprint(1)
		log.Info("excelColumn=" + excelColumn)
		excelHeaderMap[excelColumn] = value
		j++
	}

	for k, v := range excelHeaderMap {
		//设置单元格的值
		xlsx.SetCellValue(sheetName, k, v)
	}

	for i, userRecord := range re {
		row := i + 2
		rows := fmt.Sprint(row)
		var tmp byte
		tmp = 0
		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.ID)
		tmp = tmp + 1
		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.Name)
		tmp = tmp + 1
		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CodeMergeCount)
		tmp = tmp + 1
		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CommitCount)
		tmp = tmp + 1
		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.IssueCount)
		tmp = tmp + 1
		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CommentCount)
		tmp = tmp + 1
		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.WatchedCount)
		tmp = tmp + 1
		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CommitCodeSize)
		tmp = tmp + 1
		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.SolveIssueCount)
		tmp = tmp + 1
		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CommitDatasetNum)
		tmp = tmp + 1
		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CommitModelCount)
		tmp = tmp + 1
		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.ModelConvertCount)
		tmp = tmp + 1

		//new
		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.LoginCount)
		tmp = tmp + 1
		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CreateRepoCount)
		tmp = tmp + 1
		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CloudBrainTaskNum)
		tmp = tmp + 1
		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.CollectDataset)
		tmp = tmp + 1

		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.Email)
		tmp = tmp + 1

		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, ConvertOccupationToString(userRecord.Occupation))
		tmp = tmp + 1

		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.Phone)
		tmp = tmp + 1
		formatTime := userRecord.RegistDate.Format("2006-01-02 15:04:05")
		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, formatTime[0:len(formatTime)-3])
		tmp = tmp + 1
	}

	//设置默认打开的表单
	xlsx.SetActiveSheet(index)

	os.Mkdir(setting.AppDataPath+Excel_File_Path, 0755)
	if err := xlsx.SaveAs(filePath); err != nil {
		log.Info("writer exel error." + err.Error())
	} else {
		log.Info("write to file succeed, filepath=" + filePath)
	}
}

func ConvertOccupationToString(occupation int64) string {
	switch occupation {
	case 1:
		return "其他"
	case 2:
		return "学生"
	case 3:
		return "科研人员"
	case 4:
		return "企业人员"
	}
	return ""
}

// URL:  /api/v1/query_user_login?userId=1,2,3,4
func QueryUserLoginInfo(ctx *context.Context) {
	userId := ctx.Query("userId")
	userIds := strings.Split(userId, ",")
	userIdInt := make([]int64, 0)
	for _, id := range userIds {
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err == nil {
			userIdInt = append(userIdInt, idInt)
		}
	}
	result := models.QueryUserLoginInfo(userIdInt)

	xlsx := excelize.NewFile()
	sheetName := ctx.Tr("用户登录信息")
	index := xlsx.NewSheet(sheetName)
	xlsx.DeleteSheet("Sheet1")

	excelHeader := make([]string, 0)
	excelHeader = append(excelHeader, "用户ID")
	excelHeader = append(excelHeader, "登录IP")
	excelHeader = append(excelHeader, "登录时间")

	excelHeaderMap := make(map[string]string, 0)
	var j byte
	j = 0
	for _, value := range excelHeader {
		excelColumn := getColumn(j) + fmt.Sprint(1)
		log.Info("excelColumn=" + excelColumn)
		excelHeaderMap[excelColumn] = value
		j++
	}
	for k, v := range excelHeaderMap {
		//设置单元格的值
		xlsx.SetCellValue(sheetName, k, v)
	}
	for i, userLogin := range result {
		row := i + 2
		rows := fmt.Sprint(row)
		var tmp byte
		tmp = 0
		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userLogin.UId)
		tmp = tmp + 1
		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userLogin.IpAddr)
		tmp = tmp + 1
		formatTime := userLogin.CreatedUnix.Format("2006-01-02 15:04:05")
		xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, formatTime)
	}
	//设置默认打开的表单
	xlsx.SetActiveSheet(index)
	filename := sheetName + "_" + time.Now().Format("2006-01-02 15:04:05") + ".xlsx"
	ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+url.QueryEscape(filename))
	ctx.Resp.Header().Set("Content-Type", "application/octet-stream")
	if _, err := xlsx.WriteTo(ctx.Resp); err != nil {
		log.Info("writer exel error." + err.Error())
	}
}

func QueryUserAnnualReport(ctx *context.Context) {
	log.Info("start to QueryUserAnnualReport ")
	result := models.QueryUserAnnualReport(ctx.User.ID)
	ctx.JSON(http.StatusOK, result)
}

func getRecordBeginTime() (time.Time, error) {
	return time.ParseInLocation("2006-01-02", setting.RadarMap.RecordBeginTime, time.Local)
}

func QueryUserCountTimeInfo(ctx *context.Context) {
	recordBeginTime, err := getRecordBeginTime()
	if err != nil {
		recordBeginTime = time.Now()
	}
	projectsPeriodData := ProjectsPeriodData{
		RecordBeginTime: recordBeginTime.Format("2006-01-02"),
		LastUpdatedTime: models.GetLastModifyTime(),
	}
	ctx.JSON(http.StatusOK, projectsPeriodData)
}
