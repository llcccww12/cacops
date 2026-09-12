package reward

import (
	"encoding/json"
	"math"
	"strconv"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/util"
)

func AdminBalanceOperate(req models.AdminRewardOperateReq, doer *models.User) error {
	logId := util.UUID()
	adminLog := &models.RewardAdminLog{
		LogId:        logId,
		Amount:       req.Amount,
		RewardType:   req.RewardType.Name(),
		TargetUserId: req.TargetUserId,
		CreatorId:    doer.ID,
		CreatorName:  doer.Name,
		Remark:       req.Remark,
		Status:       models.RewardAdminLogProcessing,
	}
	_, err := models.InsertRewardAdminLog(adminLog)
	if err != nil {
		log.Error("AdminBalanceOperate InsertRewardAdminLog error.%v", err)
		return err
	}
	var sourceContent string
	if contentByte, err := json.Marshal(adminLog.ToShow()); err == nil {
		sourceContent = string(contentByte)
	}
	//reward
	err = Operate(&models.RewardOperateContext{
		SourceType: models.SourceTypeAdminOperate,
		SourceId:   logId,
		Title:      "管理员操作",
		Reward: models.Reward{
			Amount: req.Amount,
			Type:   req.RewardType,
		},
		TargetUserId:      req.TargetUserId,
		RequestId:         logId,
		OperateType:       req.OperateType,
		Remark:            req.Remark,
		RejectPolicy:      models.JustReject,
		PermittedNegative: true,
		SourceContent:     sourceContent,
	})

	if err != nil {
		log.Error("AdminBalanceOperate operate error.%v", err)
		models.UpdateRewardAdminLogStatus(logId, models.RewardAdminLogProcessing, models.RewardAdminLogFailed)
		return err
	}
	models.UpdateRewardAdminLogStatus(logId, models.RewardAdminLogProcessing, models.RewardAdminLogSuccess)
	return nil
}

func GetCSVFailedDatas(admimUser *models.User, records [][]string) ([]models.CSVFailedData, int) {
	var SuccessNum int
	var CSVFailedDatas []models.CSVFailedData
	var RewardOperateType models.RewardOperateType
	for i, record := range records {
		//跳过csv文件第一行(属性行)并检查
		if i == 0 {
			if len(record) != 4 {
				CSVFailedDatas = append(CSVFailedDatas, models.CSVFailedData{models.PointUser{}, "wrong number of columns"})
				return CSVFailedDatas, 0
			}
			continue
		}

		userId := record[0]
		userName := record[1]
		amount := record[2]
		remark := record[3]
		pointUser := models.PointUser{
			UserId:   userId,
			UserName: userName,
			Amount:   amount,
			Remark:   remark,
		}

		targetUserId, _ := strconv.ParseInt(userId, 10, 64)
		targetUser, err := models.GetUserByID(targetUserId)
		if err != nil {
			CSVFailedDatas = append(CSVFailedDatas, models.CSVFailedData{pointUser, "userid not found"})
			continue
		}
		if userName != targetUser.Name {
			CSVFailedDatas = append(CSVFailedDatas, models.CSVFailedData{pointUser, "userName not found"})
			continue
		}
		amountNum, err := strconv.ParseFloat(amount,
			64)
		if err != nil {
			CSVFailedDatas = append(CSVFailedDatas, models.CSVFailedData{pointUser, "amount is not number"})
			continue
		}

		if amountNum >= 0 {
			RewardOperateType = models.OperateTypeIncrease
		} else {
			//负整数转为正整数
			amountNum = math.Abs(amountNum)
			RewardOperateType = models.OperateTypeDecrease
		}
		if remark == "" {
			CSVFailedDatas = append(CSVFailedDatas, models.CSVFailedData{pointUser, "remark is empty"})
			continue
		}

		err = AdminBalanceOperate(models.AdminRewardOperateReq{
			TargetUserId: targetUserId,
			OperateType:  RewardOperateType,
			Amount:       amountNum,
			Remark:       remark,
			RewardType:   models.RewardTypePoint,
		}, admimUser)
		if err != nil {
			log.Error("OperatePointAccountBalance error.%v", err)
			CSVFailedDatas = append(CSVFailedDatas, models.CSVFailedData{pointUser, "OperatePointAccountBalance error."})
			continue
		}
		SuccessNum++
	}
	return CSVFailedDatas, SuccessNum
}
