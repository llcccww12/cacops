package models

import (
	"code.gitea.io/gitea/modules/timeutil"
)

const (
	RewardAdminLogProcessing = 1
	RewardAdminLogSuccess    = 2
	RewardAdminLogFailed     = 3
)

type RewardAdminLog struct {
	ID           int64   `xorm:"pk autoincr"`
	LogId        string  `xorm:"INDEX NOT NULL"`
	Amount       float64 `xorm:"NOT NULL"`
	RewardType   string
	Remark       string
	Status       int
	TargetUserId int64 `xorm:"INDEX NOT NULL"`
	CreatorId    int64 `xorm:"NOT NULL"`
	CreatorName  string
	CreatedUnix  timeutil.TimeStamp `xorm:"INDEX created"`
}

func (r *RewardAdminLog) ToShow() *RewardAdminLogShow {
	return &RewardAdminLogShow{
		CreatorName: r.CreatorName,
	}
}

type RewardAdminLogShow struct {
	CreatorName string
}

type AdminLogAndUser struct {
	AdminRewardAdminLog RewardAdminLog `xorm:"extends"`
	User                User           `xorm:"extends"`
}

type PointUser struct {
	UserId   string
	UserName string
	Amount   string
	Remark   string
}

type CSVFailedData struct {
	PointUser PointUser
	ErrorMsg  string
}
type BatchRewardPointReturnData struct {
	SuccessNum     int
	FailedNum      int
	CSVFailedDatas []CSVFailedData
}

func getRewardAdminLog(ra *RewardAdminLog) (*RewardAdminLog, error) {
	has, err := x.Get(ra)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrRecordNotExist{}
	}
	return ra, nil
}

func InsertRewardAdminLog(ra *RewardAdminLog) (int64, error) {
	return x.Insert(ra)
}

func UpdateRewardAdminLogStatus(logId string, oldStatus, newStatus int) error {
	_, err := x.Where("log_id = ? and status = ?", logId, oldStatus).Update(&RewardAdminLog{Status: newStatus})
	if err != nil {
		return err
	}
	return nil
}

func GetRewardAdminLogByLogIds(logIds []string) ([]*RewardAdminLog, error) {
	if len(logIds) == 0 {
		return nil, nil
	}
	adminLogs := make([]*AdminLogAndUser, 0)
	err := x.Table("reward_admin_log").Join("LEFT", "user", "reward_admin_log.creator_id = public.user.id").In("reward_admin_log.log_id", logIds).Find(&adminLogs)
	if err != nil {
		return nil, err
	}
	r := make([]*RewardAdminLog, len(adminLogs))
	for i, v := range adminLogs {
		temp := &v.AdminRewardAdminLog
		temp.CreatorName = v.User.Name
		r[i] = temp
	}
	return r, nil
}

func GetRewardAdminLogByLogId(logId string) (*RewardAdminLog, error) {
	if logId == "" {
		return nil, ErrRecordNotExist{}
	}
	adminLog := &AdminLogAndUser{}
	has, err := x.Table("reward_admin_log").Join("LEFT", "user", "reward_admin_log.creator_id = public.user.id").Where("reward_admin_log.log_id = ?", logId).Get(adminLog)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrRecordNotExist{}
	}
	r := &adminLog.AdminRewardAdminLog
	r.CreatorName = adminLog.User.Name
	return r, nil
}
