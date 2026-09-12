package repo

import (
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func QueryInvitationCurrentMonth(ctx *context.Context) {
	// userName := ctx.Query("userName")
	// currentTimeNow := time.Now()
	// pageEndTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 0, 0, 0, 0, currentTimeNow.Location())
	// pageStartTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), 1, 0, 0, 0, 0, currentTimeNow.Location())

	//queryUserDataPage(ctx, "public.user_business_analysis_current_month", new(models.UserBusinessAnalysisCurrentMonth))
	//_, count := models.QueryUserStaticDataByTableName(1, 1, "public.user_business_analysis_current_month", new(models.UserBusinessAnalysisCurrentMonth), userName, 1)

	queryDataFromStaticTable(ctx, "public.user_business_analysis_current_month", new(models.UserBusinessAnalysisCurrentMonth))
}

func getInvitationExcelHeader(ctx *context.Context) map[string]string {
	excelHeader := make([]string, 0)
	excelHeader = append(excelHeader, ctx.Tr("user.static.id"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.name"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.invitationNum"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.phone"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.registdate"))

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

func getInvitationDetailExcelHeader(ctx *context.Context) map[string]string {
	excelHeader := make([]string, 0)
	excelHeader = append(excelHeader, ctx.Tr("user.static.id"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.name"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.email"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.phone"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.registdate"))
	excelHeader = append(excelHeader, ctx.Tr("user.static.srcUserId"))

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

func writeInvitationExcel(row int, xlsx *excelize.File, sheetName string, userRecord *models.UserBusinessAnalysisAll) {
	rows := fmt.Sprint(row)
	var tmp byte
	tmp = 0
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.ID)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.Name)
	tmp = tmp + 1

	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.InvitationUserNum)
	tmp = tmp + 1

	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.Phone)
	tmp = tmp + 1

	formatTime := userRecord.RegistDate.Format("2006-01-02 15:04:05")
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, formatTime[0:len(formatTime)-3])

}

func writeInvitationDetailExcel(row int, xlsx *excelize.File, sheetName string, userRecord *models.Invitation) {
	rows := fmt.Sprint(row)
	var tmp byte
	tmp = 0
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.UserID)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.Name)
	tmp = tmp + 1
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.Email)
	tmp = tmp + 1

	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.Phone)
	tmp = tmp + 1

	formatTime := userRecord.CreatedUnix.Format("2006-01-02 15:04:05")
	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, formatTime[0:len(formatTime)-3])
	tmp = tmp + 1

	xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.SrcUserID)
}

func DownloadInvitationDetail(ctx *context.Context) {
	xlsx := excelize.NewFile()
	sheetName := ctx.Tr("user.static.invitationdetailsheetname")
	index := xlsx.NewSheet(sheetName)
	xlsx.DeleteSheet("Sheet1")
	excelHeader := getInvitationDetailExcelHeader(ctx)
	for k, v := range excelHeader {
		//设置单元格的值
		xlsx.SetCellValue(sheetName, k, v)
	}
	userNameMap := models.GetAllUserName()
	_, count := models.QueryInvitaionPage(1, 1)
	var indexTotal int64
	indexTotal = 0
	row := 1
	for {
		re, _ := models.QueryInvitaionPage(int(indexTotal), PAGE_SIZE)
		log.Info("return count=" + fmt.Sprint(count))
		for _, userRecord := range re {
			row++
			userRecord.Name = userNameMap[userRecord.UserID]
			if userRecord.Name == "" {
				userRecord.Name = "已注销"
			}
			writeInvitationDetailExcel(row, xlsx, sheetName, userRecord)
		}
		indexTotal += PAGE_SIZE
		if indexTotal >= count {
			break
		}
	}
	//设置默认打开的表单
	xlsx.SetActiveSheet(index)
	filename := sheetName + ".xlsx"
	ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+url.QueryEscape(filename))
	ctx.Resp.Header().Set("Content-Type", "application/octet-stream")
	if _, err := xlsx.WriteTo(ctx.Resp); err != nil {
		log.Info("writer exel error." + err.Error())
	}
}

func queryDataFromStaticTable(ctx *context.Context, tableName string, queryObj interface{}) {
	page, pageSize := getPageInfo(ctx)
	userName := ctx.Query("userName")
	IsReturnFile := ctx.QueryBool("IsReturnFile")

	if IsReturnFile {
		//writer exec file.
		xlsx := excelize.NewFile()
		sheetName := ctx.Tr("user.static.invitationsheetname")
		index := xlsx.NewSheet(sheetName)
		xlsx.DeleteSheet("Sheet1")
		excelHeader := getInvitationExcelHeader(ctx)
		for k, v := range excelHeader {
			//设置单元格的值
			xlsx.SetCellValue(sheetName, k, v)
		}
		_, count := models.QueryUserInvitationDataByTableName(1, 1, tableName, queryObj, "", 1)
		var indexTotal int64
		indexTotal = 0
		row := 1
		for {
			re, _ := models.QueryUserInvitationDataByTableName(int(indexTotal), PAGE_SIZE, tableName, queryObj, "", 1)
			log.Info("return count=" + fmt.Sprint(count))
			for _, userRecord := range re {
				row++
				writeInvitationExcel(row, xlsx, sheetName, userRecord)
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
		resultRecord, count := models.QueryUserInvitationDataByTableName((page-1)*pageSize, pageSize, tableName, queryObj, userName, 1)
		result := make([]models.Invitation, 0)
		for _, record := range resultRecord {
			invi := models.Invitation{
				SrcUserID:         record.ID,
				Name:              record.Name,
				InvitationUserNum: record.InvitationUserNum,
				Phone:             record.Phone,
				CreatedUnix:       record.RegistDate,
			}
			result = append(result, invi)
		}
		mapInterface := make(map[string]interface{})
		mapInterface["data"] = result
		mapInterface["count"] = count
		ctx.JSON(http.StatusOK, mapInterface)
	}
}

func QueryInvitationCurrentWeek(ctx *context.Context) {
	// currentTimeNow := time.Now()
	// offset := int(time.Monday - currentTimeNow.Weekday())
	// if offset > 0 {
	// 	offset = -6
	// }
	// pageStartTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	// pageEndTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 0, 0, 0, 0, currentTimeNow.Location())
	// queryData(ctx, pageStartTime.Unix(), pageEndTime.Unix())
	queryDataFromStaticTable(ctx, "public.user_business_analysis_current_week", new(models.UserBusinessAnalysisCurrentWeek))
}

func QueryInvitationLastWeek(ctx *context.Context) {
	// currentTimeNow := time.Now()
	// offset := int(time.Monday - currentTimeNow.Weekday())
	// if offset > 0 {
	// 	offset = -6
	// }
	// pageEndTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	// pageStartTime := pageEndTime.AddDate(0, 0, -7)
	// queryData(ctx, pageStartTime.Unix(), pageEndTime.Unix())

	queryDataFromStaticTable(ctx, "public.user_business_analysis_last_week", new(models.UserBusinessAnalysisLastWeek))
}

func QueryInvitationCurrentYear(ctx *context.Context) {
	// currentTimeNow := time.Now()
	// pageStartTime := time.Date(currentTimeNow.Year(), 1, 1, 0, 0, 0, 0, currentTimeNow.Location())
	// pageEndTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 0, 0, 0, 0, currentTimeNow.Location())
	// queryData(ctx, pageStartTime.Unix(), pageEndTime.Unix())

	queryDataFromStaticTable(ctx, "public.user_business_analysis_current_year", new(models.UserBusinessAnalysisCurrentYear))
}

func QueryInvitationLast30Day(ctx *context.Context) {
	// currentTimeNow := time.Now()
	// pageStartTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, -30)
	// pageEndTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 0, 0, 0, 0, currentTimeNow.Location())
	// queryData(ctx, pageStartTime.Unix(), pageEndTime.Unix())

	queryDataFromStaticTable(ctx, "public.user_business_analysis_last30_day", new(models.UserBusinessAnalysisLast30Day))
}

func QueryInvitationLastMonth(ctx *context.Context) {
	// currentTimeNow := time.Now()
	// thisMonth := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), 1, 0, 0, 0, 0, currentTimeNow.Location())
	// pageStartTime := thisMonth.AddDate(0, -1, 0)
	// pageEndTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), 1, 23, 59, 59, 0, currentTimeNow.Location()).AddDate(0, 0, -1)
	// queryData(ctx, pageStartTime.Unix(), pageEndTime.Unix())

	queryDataFromStaticTable(ctx, "public.user_business_analysis_last_month", new(models.UserBusinessAnalysisLastMonth))
}

func QueryInvitationYesterday(ctx *context.Context) {
	// currentTimeNow := time.Now().AddDate(0, 0, -1)
	// pageStartTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 0, 0, 0, 0, time.Local)
	// pageEndTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 23, 59, 59, 0, currentTimeNow.Location())
	// queryData(ctx, pageStartTime.Unix(), pageEndTime.Unix())

	queryDataFromStaticTable(ctx, "public.user_business_analysis_yesterday", new(models.UserBusinessAnalysisYesterday))
}

func QueryInvitationAll(ctx *context.Context) {
	// currentTimeNow := time.Now()
	// pageStartTime := time.Date(2022, 8, 5, 0, 0, 0, 0, currentTimeNow.Location())
	// pageEndTime := time.Date(currentTimeNow.Year(), currentTimeNow.Month(), currentTimeNow.Day(), 0, 0, 0, 0, currentTimeNow.Location())
	// queryData(ctx, pageStartTime.Unix(), pageEndTime.Unix())

	queryDataFromStaticTable(ctx, "public.user_business_analysis_all", new(models.UserBusinessAnalysisAll))
}

func QueryUserDefineInvitationPage(ctx *context.Context) {
	startDate := ctx.Query("startDate")
	endDate := ctx.Query("endDate")
	startTime, _ := time.ParseInLocation("2006-01-02", startDate, time.Local)
	//startTime = startTime.UTC()
	endTime, _ := time.ParseInLocation("2006-01-02", endDate, time.Local)

	queryData(ctx, startTime, endTime)
}

func queryData(ctx *context.Context, startTime time.Time, endTime time.Time) {
	page, pageSize := getPageInfo(ctx)
	IsReturnFile := ctx.QueryBool("IsReturnFile")

	dbResult := models.QueryInvitaionByTime(startTime.Unix(), endTime.Unix())

	invitaionNumMap := make(map[int64]int, 0)
	allUserIds := make([]int64, 0)
	for _, record := range dbResult {
		if _, ok := invitaionNumMap[record.SrcUserID]; !ok {
			invitaionNumMap[record.SrcUserID] = 1
		} else {
			invitaionNumMap[record.SrcUserID] = invitaionNumMap[record.SrcUserID] + 1
		}
	}
	invitaionNumList := make([]models.Invitation, 0)
	for key, value := range invitaionNumMap {
		invi := models.Invitation{
			SrcUserID:         key,
			InvitationUserNum: value,
		}
		invitaionNumList = append(invitaionNumList, invi)
		allUserIds = append(allUserIds, key)
	}
	sort.Slice(invitaionNumList, func(i, j int) bool {
		return invitaionNumList[i].InvitationUserNum > invitaionNumList[j].InvitationUserNum
	})
	if IsReturnFile {
		xlsx := excelize.NewFile()
		sheetName := ctx.Tr("user.static.invitationsheetname")
		index := xlsx.NewSheet(sheetName)
		xlsx.DeleteSheet("Sheet1")
		excelHeader := getInvitationExcelHeader(ctx)
		for k, v := range excelHeader {
			//设置单元格的值
			xlsx.SetCellValue(sheetName, k, v)
		}
		end := 100
		userMap := make(map[int64]*models.User, 0)
		log.Info("len(allUserIds)=" + fmt.Sprint(len(allUserIds)))
		for i := 0; i < len(allUserIds); i += 100 {
			if end >= len(allUserIds) {
				end = len(allUserIds)
			}
			log.Info("i=" + fmt.Sprint(i) + " end=" + fmt.Sprint(end))
			if i == end {
				break
			}
			userList, err := models.GetUsersByIDs(allUserIds[i:end])
			if err == nil {
				for _, tmp := range userList {
					userMap[tmp.ID] = tmp
				}
			} else {

			}
			end = end + 100
		}
		row := 1
		log.Info("len(userMap)=" + fmt.Sprint(len(userMap)))
		for _, userRecord := range invitaionNumList {
			row++
			rows := fmt.Sprint(row)
			var tmp byte
			tmp = 0

			xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.SrcUserID)
			tmp = tmp + 1
			name := "已注销"
			if userMap[userRecord.SrcUserID] != nil {
				name = userMap[userRecord.SrcUserID].Name
			}
			xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, name)
			tmp = tmp + 1

			xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, userRecord.InvitationUserNum)
			tmp = tmp + 1
			Phone := ""
			if userMap[userRecord.SrcUserID] != nil {
				Phone = userMap[userRecord.SrcUserID].PhoneNumber
			}
			xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, Phone)
			tmp = tmp + 1

			formatTime := ""
			if userMap[userRecord.SrcUserID] != nil {
				formatTime = userMap[userRecord.SrcUserID].CreatedUnix.Format("2006-01-02 15:04:05")
				formatTime = formatTime[0 : len(formatTime)-3]
			}
			xlsx.SetCellValue(sheetName, getColumn(tmp)+rows, formatTime)
		}
		//设置默认打开的表单
		xlsx.SetActiveSheet(index)
		filename := sheetName + "_" + getTimeFileName(startTime) + "_" + getTimeFileName(endTime) + ".xlsx"
		//filename := sheetName + "_" + ctx.Tr("user.static."+tableName) + ".xlsx"
		ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+url.QueryEscape(filename))
		ctx.Resp.Header().Set("Content-Type", "application/octet-stream")
		if _, err := xlsx.WriteTo(ctx.Resp); err != nil {
			log.Info("writer exel error." + err.Error())
		}
	} else {
		result := make([]*models.Invitation, 0)
		userIds := make([]int64, 0)
		end := len(invitaionNumList) - 1
		for start := (page - 1) * pageSize; start <= end; start++ {
			invi := invitaionNumList[start]
			//todo name phone,createunix
			result = append(result, &invi)
			userIds = append(userIds, invi.SrcUserID)
			if len(result) == pageSize {
				break
			}
		}
		userList, err := models.GetUsersByIDs(userIds)
		if err == nil {
			for _, invi := range result {
				tmpUser := userList[0]
				for _, tmp := range userList {
					if tmp.ID == invi.SrcUserID {
						tmpUser = tmp
						break
					}
				}
				if invi.SrcUserID == tmpUser.ID {
					invi.Name = tmpUser.Name
					invi.Phone = tmpUser.PhoneNumber
					invi.CreatedUnix = tmpUser.CreatedUnix
					invi.Email = tmpUser.Email
				} else {
					invi.Name = "已注销"
				}
			}
		} else {
			log.Info("query user error." + err.Error())
		}
		mapInterface := make(map[string]interface{})
		mapInterface["data"] = result
		mapInterface["count"] = len(invitaionNumList)
		ctx.JSON(http.StatusOK, mapInterface)
	}
}

func getPageInfo(ctx *context.Context) (int, int) {
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pageSize := ctx.QueryInt("pageSize")
	if pageSize <= 0 {
		pageSize = setting.UI.IssuePagingNum
	}
	return page, pageSize
}
