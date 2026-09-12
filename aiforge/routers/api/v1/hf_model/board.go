package hf_model

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"encoding/csv"
	"encoding/json"
	"net/http"
	"strconv"
)

func getCtxModelQueryOptions(ctx *context.APIContext) *models.ModelQueryOptions {
	status := 100
	if ctx.Query("status") != "" {
		status = ctx.QueryInt("status")
	}

	return &models.ModelQueryOptions{
		Page:             ctx.QueryInt("page"),
		PageSize:         ctx.QueryInt("pageSize"),
		CreatedUnixStart: ctx.QueryInt("createdUnixStart"),
		CreatedUnixEnd:   ctx.QueryInt("createdUnixEnd"),
		Keyword:          ctx.Query("q"),
		KeywordType:      ctx.QueryInt("q_type"),
		Status:           status,
	}
}

func getModelStats(opt *models.ModelQueryOptions) ([]*models.ModelQueryData, int64, error) {
	res, total, err := models.QueryHfModelStats(opt)
	if err != nil {
		return res, total, err
	}
	return res, total, nil
}

func GetModelStats(ctx *context.APIContext) {
	opt := getCtxModelQueryOptions(ctx)
	modelList, total, err := getModelStats(opt)
	if err != nil {
		ctx.Error(500, "QueryHfModelStats", err.Error())
		return
	}

	if len(modelList) == 0 {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"count": total,
			"data":  make([]string, 0),
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"count": total,
		"data":  modelList,
	})
}

func DownloadModelStats(ctx *context.APIContext) {
	opt := getCtxModelQueryOptions(ctx)
	modelList, _, err := getModelStats(opt)
	if err != nil {
		ctx.Error(500, "QueryHfModelStats", err.Error())
		return
	}
	exportCSV(ctx, modelList)
}

// Export to CSV and send as response
func exportCSV(ctx *context.APIContext, modelList []*models.ModelQueryData) {
	ctx.Resp.Header().Set("Content-Type", "text/csv")
	ctx.Resp.Header().Set("Content-Disposition", "attachment; filename=hf_model_stats.csv")

	writer := csv.NewWriter(ctx.Resp)
	defer writer.Flush()

	// Write Header
	header := []string{"ID", "Name", "HFRepoID", "OpenIRepoID", "Source", "UserId", "Username",
		"TotalFiles", "FinishedFiles", "TotalSize", "FinishedSize", "CreatedUnix",
		"HfModelStatus", "OperatorList"}
	writer.Write(header)

	// Write Data
	for _, model := range modelList {
		operatorListJSON, _ := json.Marshal(model.OperatorsList)
		log.Info("OperatorList: %s", string(operatorListJSON))
		record := []string{
			model.ID, model.Name, model.HFRepoID, model.OpenIRepoID, model.Source,
			strconv.FormatInt(model.UserId, 10), model.Username,
			strconv.FormatInt(model.TotalFiles, 10), strconv.FormatInt(model.FinishedFiles, 10),
			strconv.FormatInt(model.TotalSize, 10), strconv.FormatInt(model.FinishedSize, 10),
			strconv.FormatInt(model.CreatedUnix, 10), strconv.FormatInt(model.AiModelStatus, 10),
			string(operatorListJSON),
		}
		writer.Write(record)
	}
}
