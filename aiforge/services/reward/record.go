package reward

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"math"
)

type RecordResponse struct {
	Records  []*models.RewardOperateRecordShow
	Total    int64
	PageSize int
	Page     int
}

func GenerateAdminRewardExcel(opts *models.RewardRecordListOpts) (*excelize.File, string, error) {
	opts.Page = 1
	opts.PageSize = 100
	recordList, total, err := models.GetAdminRewardRecordShowList(opts)
	if err != nil {
		log.Error("export increase point excel error.opts=%+v  err=%v", opts, err)
		return nil, "", err
	}
	if len(recordList) == 0 {
		log.Error("no content to export.opts=%+v ", opts)
		return nil, "", errors.New("no content to export")
	}
	xlsx := excelize.NewFile()
	sheetName := "积分获取明细"
	if opts.OperateType == models.OperateTypeDecrease {
		sheetName = "积分消耗明细"
	}
	xlsx.NewSheet(sheetName)
	xlsx.DeleteSheet("Sheet1")
	for i, v := range recordList[0].ConvertToExcelColumn() {
		cellRef := fmt.Sprintf("%s%d", excelColumnName(i), 1)
		xlsx.SetCellValue(sheetName, cellRef, v.Name)
	}

	currentRow := int64(1)
	currentRow = writeIncreasePointList2Excel(recordList, sheetName, xlsx, currentRow)
	totalPage := int64(math.Ceil(float64(total) / float64(opts.PageSize)))
	for page := 2; int64(page) <= totalPage; page++ {
		opts.Page = page
		recordList, _, err = models.GetAdminRewardRecordShowList(opts)
		if err != nil {
			log.Error("export increase point excel error.opts=%+v  err=%v", opts, err)
			return nil, "", err
		}
		if len(recordList) == 0 {
			break
		}
		currentRow = writeIncreasePointList2Excel(recordList, sheetName, xlsx, currentRow)
	}
	filename := sheetName + ".xlsx"

	return xlsx, filename, nil
}

func writeIncreasePointList2Excel(list models.RewardRecordShowList, sheetName string, xlsx *excelize.File, currentRow int64) int64 {
	for i := 0; i < len(list); i++ {
		currentRow = writeIncreasePoint2Excel(list[i], sheetName, xlsx, currentRow)
	}
	return currentRow
}

func writeIncreasePoint2Excel(record *models.RewardOperateRecordShow, sheetName string, xlsx *excelize.File, currentRow int64) int64 {
	columns := record.ConvertToExcelColumn()
	currentRow++
	for i := 0; i < len(columns); i++ {
		cellRef := fmt.Sprintf("%s%d", excelColumnName(i), currentRow)
		xlsx.SetCellValue(sheetName, cellRef, columns[i].Value)
	}
	return currentRow
}

func initIncreasePointExcel() (string, *excelize.File) {
	xlsx := excelize.NewFile()
	sheetName := "积分明细"
	xlsx.NewSheet(sheetName)
	xlsx.DeleteSheet("Sheet1")
	dataHeader := getIncreasePointExcelHeader()
	for i, v := range dataHeader {
		cellRef := fmt.Sprintf("%s%d", excelColumnName(i), 1)
		xlsx.SetCellValue(sheetName, cellRef, v)
	}
	return sheetName, xlsx
}

// 将列索引转换为 Excel 列名
func excelColumnName(index int) string {
	dividend := index + 1
	columnName := ""
	for dividend > 0 {
		modulo := (dividend - 1) % 26
		columnName = string(rune('A'+modulo)) + columnName
		dividend = (dividend - modulo) / 26
	}
	return columnName
}

func getIncreasePointExcelHeader() []string {
	excelHeader := make([]string, 0)
	excelHeader = append(excelHeader, "流水号")
	excelHeader = append(excelHeader, "用户名")
	excelHeader = append(excelHeader, "时间")
	excelHeader = append(excelHeader, "场景")
	excelHeader = append(excelHeader, "积分行为")
	excelHeader = append(excelHeader, "数量")
	excelHeader = append(excelHeader, "积分余额")

	return excelHeader
}

func GetRewardRecordList(opts *models.RewardRecordListOpts) (*RecordResponse, error) {
	var l models.RewardRecordShowList
	var n int64
	var err error
	if opts.IsAdmin {
		l, n, err = models.GetAdminRewardRecordShowList(opts)
	} else {
		l, n, err = models.GetRewardRecordShowList(opts)
	}
	if err != nil {
		log.Error("GetRewardRecordList error. %v", err)
		return nil, err
	}
	if len(l) == 0 {
		return &RecordResponse{Records: make([]*models.RewardOperateRecordShow, 0), Total: n, Page: opts.Page, PageSize: opts.PageSize}, nil
	}
	return &RecordResponse{Records: l, Total: n, Page: opts.Page, PageSize: opts.PageSize}, nil
}

func handleRecordResponse(opts *models.RewardRecordListOpts, list models.RewardRecordShowList) {
	if opts.IsAdmin {
		for _, v := range list {
			v.UserName = opts.UserName
		}
	} else {
		for _, v := range list {
			if v.Cloudbrain != nil {
				v.Cloudbrain.AiCenter = ""
			}
		}
	}
}
