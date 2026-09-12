package models

import (
	"code.gitea.io/gitea/modules/timeutil"
)

type UserOtherInfo struct {
	ID          int64 `xorm:"pk autoincr"`
	UId         int64 `xorm:"NOT NULL"`
	IsNewAgree  bool
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
}

func GetNewAgreeByUID(uid int64) bool {
	userOtherInfo := new(UserOtherInfo)
	has, err := xStatistic.Where("u_id=?", uid).Desc("id").Limit(1).Get(userOtherInfo)
	if err != nil || !has {
		return false
	}
	return userOtherInfo.IsNewAgree
}

func SaveUserOtherInfoToDb(uid int64) {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()

	userOtherInfo := &UserOtherInfo{
		UId:        uid,
		IsNewAgree: true,
	}
	statictisSess.Insert(userOtherInfo)
}
