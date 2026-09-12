package point

import (
	"bufio"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/reward"
	"code.gitea.io/gitea/services/reward/point/account"
	"code.gitea.io/gitea/services/task"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"net/http"
	"net/url"
	"strconv"
)

const tplPoint base.TplName = "reward/point"
const tplPointRule base.TplName = "reward/point/rule"

type AccountResponse struct {
	Balance       float64
	TotalEarned   float64
	TotalConsumed float64
}

func GetPointAccount(ctx *context.Context) {
	userId := ctx.User.ID
	a, err := account.GetAccount(userId)
	if err != nil {
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	res := &AccountResponse{
		Balance:       a.Balance,
		TotalEarned:   a.TotalEarned,
		TotalConsumed: a.TotalConsumed,
	}
	ctx.JSON(http.StatusOK, response.SuccessWithData(res))
}

func GetPointRecordList(ctx *context.Context) {
	operateType := ctx.Query("Operate")
	page := ctx.QueryInt("Page")
	var orderBy models.RewardOperateOrderBy
	switch ctx.Query("sort") {
	default:
		orderBy = models.RewardOrderByIDDesc
	}
	t := models.GetRewardOperateTypeInstance(operateType)
	if t == "" {
		ctx.JSON(http.StatusOK, response.ServerError("param error"))
		return
	}

	r, err := reward.GetRewardRecordList(&models.RewardRecordListOpts{
		ListOptions: models.ListOptions{PageSize: 10, Page: page},
		UserId:      ctx.User.ID,
		OperateType: t,
		RewardType:  models.RewardTypePoint,
		OrderBy:     orderBy,
		IsAdmin:     false,
		UserName:    ctx.User.Name,
	})
	if err != nil {
		log.Error("GetPointRecordList error.%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessWithData(r))
	return
}

func OperatePointAccountBalance(ctx *context.Context, req models.AdminRewardOperateReq) {
	req.RewardType = models.RewardTypePoint
	if req.OperateType.Name() == "" || req.Remark == "" {
		ctx.JSON(http.StatusOK, "param error")
		return
	}
	err := reward.AdminBalanceOperate(req, ctx.User)
	if err != nil {
		log.Error("OperatePointAccountBalance error.%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.Success())
}

func BatchOperatePointAccountBalance(ctx *context.Context) {
	file, header, err := ctx.Req.FormFile("csvFile")
	if err != nil {
		log.Error("FormFile error.%v", err)
		ctx.Error(500, fmt.Sprintf("FormFile: %v", err))
		return
	}
	defer file.Close()

	if header.Header.Get("Content-Type") != "text/csv" {
		log.Error("Only CSV files are allowed error.%v", err)
		ctx.Error(500, "Only CSV files are allowed")
		return
	}
	// 创建一个转换器，将GBK编码转换为UTF-8编码
	reader := bufio.NewReader(file)
	decoder := simplifiedchinese.GBK.NewDecoder()
	convertedReader := decoder.Reader(reader)
	csvReader := csv.NewReader(convertedReader)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Error("FormFile error.%v", err)
		ctx.Error(500, "Error reading records")
		return
	}

	CSVFailedDatas, SuccessNum := reward.GetCSVFailedDatas(ctx.User, records)

	batchRewardPointReturnData := models.BatchRewardPointReturnData{
		SuccessNum:     SuccessNum,
		FailedNum:      len(CSVFailedDatas),
		CSVFailedDatas: CSVFailedDatas,
	}
	ctx.JSON(http.StatusOK, response.SuccessWithData(batchRewardPointReturnData))
}

func GetPointPage(ctx *context.Context) {
	ctx.HTML(200, tplPoint)
}

func GetRulePage(ctx *context.Context) {
	ctx.HTML(200, tplPointRule)
}

func GetRuleConfig(ctx *context.Context) {
	r, err := task.GetPointRule()
	if err != nil {
		log.Error("GetRuleConfig error.%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessWithData(r))
}

func GetAdminRewardList(ctx *context.Context) {
	opts, err := buildAdminRewardRecordListOpts(ctx)
	if err != nil {
		log.Error("buildAdminRewardRecordListOpts error.%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}

	username := ctx.Query("userName")
	if username != "" {
		user, err := models.GetUserByName(username)
		if err != nil {
			log.Error("GetUserByName error.%v", err)
			if models.IsErrUserNotExist(err) {
				ctx.JSON(http.StatusOK, response.ServerError("user not exist"))
			} else {
				ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
			}
			return
		}
		opts.UserId = user.ID
		opts.UserName = user.Name
	}

	r, err := reward.GetRewardRecordList(opts)
	if err != nil {
		log.Error("GetRewardRecordList error.%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessWithData(r))
}

func ExportAdminRewardList(ctx *context.Context) {
	opts, err := buildAdminRewardRecordListOpts(ctx)
	if err != nil {
		log.Error("buildAdminRewardRecordListOpts error.%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}

	username := ctx.Query("userName")
	if username != "" {
		user, err := models.GetUserByName(username)
		if err != nil {
			log.Error("GetUserByName error.%v", err)
			if models.IsErrUserNotExist(err) {
				ctx.JSON(http.StatusOK, response.ServerError("user not exist"))
			} else {
				ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
			}
			return
		}
		opts.UserId = user.ID
		opts.UserName = user.Name
	}

	xlsx, fileName, err := reward.GenerateAdminRewardExcel(opts)
	if err != nil {
		log.Error("GenerateAdminReward2Response error.%v", err)
		ctx.JSON(http.StatusOK, response.ServerError(err.Error()))
		return
	}
	ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+url.QueryEscape(fileName))
	ctx.Resp.Header().Set("Content-Type", "application/octet-stream")
	if _, err := xlsx.WriteTo(ctx.Resp); err != nil {
		log.Info("writer reward exel error." + err.Error())
	}
}

func buildAdminRewardRecordListOpts(ctx *context.Context) (*models.RewardRecordListOpts, error) {
	operateType := ctx.Query("operate")
	sourceType := ctx.Query("source")
	taskType := ctx.Query("action")
	serialNo := ctx.Query("serialNo")
	status := ctx.Query("status")

	page := ctx.QueryInt("page")
	var orderBy models.RewardOperateOrderBy
	switch ctx.Query("sort") {
	default:
		orderBy = models.RewardOrderByIDDesc
	}
	t := models.GetRewardOperateTypeInstance(operateType)
	if t == "" {
		return nil, errors.New("param error")
	}
	opts := &models.RewardRecordListOpts{
		ListOptions: models.ListOptions{PageSize: 10, Page: page},
		OperateType: t,
		RewardType:  models.RewardTypePoint,
		OrderBy:     orderBy,
		SourceType:  sourceType,
		TaskType:    taskType,
		SerialNo:    serialNo,
		IsAdmin:     true,
		Status:      status,
	}
	return opts, nil
}

func AddHistoricSourceContent(ctx *context.Context) {
	list := ctx.Query("list")
	array := make([]map[string]string, 0)
	err := json.Unmarshal([]byte(list), &array)
	if err != nil {
		ctx.JSON(http.StatusOK, response.PARAM_ERROR)
		return
	}

	for i := 0; i < len(array); i++ {
		item := array[i]
		recordType := item["type"]
		content := item["content"]
		var sourceContent string
		var sourceId string
		var sourceType = models.SourceType(recordType)
		//1、查询记录
		switch sourceType {
		case models.SourceTypeAccomplishTask:
			act := &models.Action{}
			err = json.Unmarshal([]byte(content), act)
			if err != nil {
				log.Error("json.Unmarshal err %v", err)
				continue
			}
			models.ActionList([]*models.Action{act}).LoadAllAttributes()
			contentByte, err := json.Marshal(act.ToShow())
			if err != nil {
				log.Error("act.ToShow() err %v", err)
				continue
			}
			sourceContent = string(contentByte)
			sourceId = fmt.Sprint(act.ID)
		case models.SourceTypeAdminOperate:
			adminLog := &models.RewardAdminLog{}
			err = json.Unmarshal([]byte(content), adminLog)
			if err != nil {
				log.Error("json.Unmarshal err %v", err)
				continue
			}
			contentByte, err := json.Marshal(adminLog.ToShow())
			if err != nil {
				log.Error("adminLog.ToShow() err %v", err)
				continue
			}
			sourceContent = string(contentByte)
			sourceId = adminLog.LogId
		case models.SourceTypeRunCloudbrainTask:
			cloudbrain := &models.Cloudbrain{}
			err = json.Unmarshal([]byte(content), cloudbrain)
			if err != nil {
				log.Error("json.Unmarshal err %v", err)
				continue
			}
			contentByte, err := json.Marshal(cloudbrain.ToShow())
			if err != nil {
				log.Error("cloudbrain.ToShow() err %v", err)
				continue
			}
			sourceContent = string(contentByte)
			sourceId = fmt.Sprint(cloudbrain.ID)
		}
		if sourceContent == "" {
			continue
		}
		err = models.UpdateSourceContent(sourceId, sourceType, sourceContent)
		if err != nil {
			log.Error("UpdateSourceContent err %v", err)
		}
	}

	ctx.JSON(http.StatusOK, response.Success())
}

func HandleHistoricSourceContent(ctx *context.Context) {
	var lastId int64
	for i := 0; i < 100000; i++ {
		array, err := models.GetNoSourceContentRewardRecord(lastId)
		if err != nil {
			log.Error("GetNoSourceContentRewardRecord err.%v", err)
			ctx.JSON(http.StatusOK, response.NewBizError(err))
		}

		for j := 0; j < len(array); j++ {

			var sourceContent string
			var sourceId string
			record := array[j]
			if lastId < record.ID {
				lastId = record.ID
			}
			switch record.SourceType {
			case "ACCOMPLISH_TASK":
				id, tmpErr := strconv.ParseInt(record.SourceId, 10, 64)
				if tmpErr != nil {
					log.Error("ACCOMPLISH_TASK strconv.ParseInt err.%v", tmpErr)
					continue
				}
				act, tmpErr := models.GetActionById(id)
				if tmpErr != nil || act == nil {
					log.Error("ACCOMPLISH_TASK GetActionById err or action is empty.%v", tmpErr)
					continue
				}
				contentByte, tmpErr := json.Marshal(act.ToShow())
				if tmpErr != nil {
					log.Error("ACCOMPLISH_TASK act.ToShow() err %v", err)
					continue
				}
				sourceContent = string(contentByte)
				sourceId = fmt.Sprint(act.ID)
			case "ADMIN_OPERATE":
				adminLog, tmpErr := models.GetRewardAdminLogByLogId(record.SourceId)
				if tmpErr != nil {
					log.Error("ACCOMPLISH_TASK strconv.ParseInt err.%v", tmpErr)
					continue
				}
				contentByte, err := json.Marshal(adminLog.ToShow())
				if err != nil {
					log.Error("adminLog.ToShow() err %v", err)
					continue
				}
				sourceContent = string(contentByte)
				sourceId = adminLog.LogId
			case "RUN_CLOUDBRAIN_TASK":
				cloudbrain, tmpErr := models.GetCloudbrainByIDWithDeleted(record.SourceId)
				if tmpErr != nil {
					log.Error("RUN_CLOUDBRAIN_TASK GetCloudbrainByCloudbrainID err.%v", tmpErr)
					continue
				}
				contentByte, err := json.Marshal(cloudbrain.ToShow())
				if err != nil {
					log.Error("cloudbrain.ToShow() err %v", err)
					continue
				}
				sourceContent = string(contentByte)
				sourceId = fmt.Sprint(cloudbrain.ID)
			}
			err = models.UpdateSourceContent(sourceId, models.SourceType(record.SourceType), sourceContent)
			if err != nil {
				log.Error("UpdateSourceContent err %v", err)
			}
		}
		if len(array) < 100 {
			log.Info("array < 100 .break")
			break
		}
	}
	ctx.JSON(http.StatusOK, response.Success())
}
