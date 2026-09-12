package admin

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"

	"code.gitea.io/gitea/modules/modelarts"
	"code.gitea.io/gitea/routers/repo"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	cloudbrainService "code.gitea.io/gitea/services/cloudbrain"
)

const (
	tplCloudBrains     base.TplName = "admin/cloudbrain/list"
	tplImages          base.TplName = "admin/cloudbrain/images"
	tplCommitImages    base.TplName = "admin/cloudbrain/imagecommit"
	EXCEL_DATE_FORMAT               = "20060102150405"
	CREATE_TIME_FORMAT              = "2006/01/02 15:04:05"
)

func CloudBrains(ctx *context.Context) {
	ctx.Data["PageIsAdmin"] = true
	ctx.Data["PageIsAdminCloudBrains"] = true
	ctx.HTML(200, tplCloudBrains)
}

func Images(ctx *context.Context) {
	ctx.Data["PageIsAdminImages"] = true
	ctx.HTML(200, tplImages)

}

func CloudBrainCommitImageShow(ctx *context.Context) {
	ctx.Data["PageIsAdminImages"] = true
	ctx.HTML(200, tplCommitImages)

}

func DownloadCloudBrains(ctx *context.Context) {

	page := 1

	pageSize := 300

	var cloudBrain = ctx.Tr("repo.cloudbrain")
	fileName := getFileName(cloudBrain)

	_, total, err := models.Cloudbrains(&models.CloudbrainsOptions{
		ListOptions: models.ListOptions{
			Page:     page,
			PageSize: 1,
		},
		Type:            models.TypeCloudBrainAll,
		NeedRepoInfo:    false,
		IsLatestVersion: modelarts.IsLatestVersion,
	})

	if err != nil {
		log.Warn("Can not get cloud brain info", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.cloudbrain_query_fail"))
		return
	}

	totalPage := getTotalPage(total, pageSize)

	f := excelize.NewFile()

	index := f.NewSheet(cloudBrain)
	f.DeleteSheet("Sheet1")

	for k, v := range allHeader(ctx) {
		f.SetCellValue(cloudBrain, k, v)
	}

	var row = 2
	for i := 0; i < totalPage; i++ {

		pageRecords, _, err := models.Cloudbrains(&models.CloudbrainsOptions{
			ListOptions: models.ListOptions{
				Page:     page,
				PageSize: pageSize,
			},
			Type:            models.TypeCloudBrainAll,
			NeedRepoInfo:    true,
			IsLatestVersion: modelarts.IsLatestVersion,
		})
		if err != nil {
			log.Warn("Can not get cloud brain info", err)
			continue
		}
		models.LoadSpecs4CloudbrainInfo(pageRecords)
		for _, record := range pageRecords {
			record = cloudbrainService.UpdateCloudbrainAiCenter(record)
			record.Cloudbrain.AiCenter = repo.GetAiCenterNameByCode(record.Cloudbrain.AiCenter, ctx.Language())
			for k, v := range allValues(row, record, ctx) {
				f.SetCellValue(cloudBrain, k, v)
			}
			row++

		}

		page++
	}
	f.SetActiveSheet(index)

	ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+url.QueryEscape(fileName))
	ctx.Resp.Header().Set("Content-Type", "application/octet-stream")

	f.WriteTo(ctx.Resp)
}

func allValues(row int, rs *models.CloudbrainInfo, ctx *context.Context) map[string]string {
	return map[string]string{getCellName("A", row): rs.DisplayJobName, getCellName("B", row): repo.GetCloudbrainCluster(rs.Cloudbrain, ctx),
		getCellName("C", row): rs.JobType, getCellName("D", row): rs.Status, getCellName("E", row): time.Unix(int64(rs.Cloudbrain.CreatedUnix), 0).Format(CREATE_TIME_FORMAT),
		getCellName("F", row): getDurationTime(rs), getCellName("G", row): rs.ComputeResource,
		getCellName("H", row): rs.Cloudbrain.AiCenter, getCellName("I", row): getCloudbrainCardType(rs),
		getCellName("J", row): rs.Name, getCellName("K", row): getRepoPathName(rs), getCellName("L", row): rs.JobName,
	}
}

func getCloudbrainCardType(rs *models.CloudbrainInfo) string {
	if rs.Cloudbrain.Spec != nil {
		return rs.Cloudbrain.Spec.AccCardType
	} else {
		return ""
	}
}

func getRepoPathName(rs *models.CloudbrainInfo) string {
	if rs.Repo != nil {
		return rs.Repo.OwnerName + "/" + rs.Repo.Alias
	}
	return ""
}

func getDurationTime(rs *models.CloudbrainInfo) string {
	if rs.JobType == "TRAIN" || rs.JobType == "INFERENCE" || rs.JobType == string(models.JobTypeSuperCompute) {
		return rs.TrainJobDuration
	} else {
		return "-"
	}
}

func getFileName(baseName string) string {
	return baseName + "_" + time.Now().Format(EXCEL_DATE_FORMAT) + ".xlsx"

}

func getTotalPage(total int64, pageSize int) int {

	another := 0
	if int(total)%pageSize != 0 {
		another = 1
	}
	return int(total)/pageSize + another

}

func allHeader(ctx *context.Context) map[string]string {

	return map[string]string{"A1": ctx.Tr("repo.cloudbrain_task"), "B1": ctx.Tr("repo.modelarts.cluster"),
		"C1": ctx.Tr("repo.cloudbrain_task_type"), "D1": ctx.Tr("repo.modelarts.status"), "E1": ctx.Tr("repo.modelarts.createtime"),
		"F1": ctx.Tr("repo.modelarts.train_job.dura_time"), "G1": ctx.Tr("repo.modelarts.computing_resources"),
		"H1": ctx.Tr("repo.modelarts.ai_center"), "I1": ctx.Tr("repo.modelarts.card_type"), "J1": ctx.Tr("repo.cloudbrain_creator"),
		"K1": ctx.Tr("repo.repo_name"), "L1": ctx.Tr("repo.cloudbrain_task_name")}

}

func getCellName(col string, row int) string {
	return col + strconv.Itoa(row)
}
