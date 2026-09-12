package models

import (
	"net/http"

	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
)

type UserLoginLog struct {
	ID          int64              `xorm:"pk autoincr"`
	UId         int64              `xorm:"NOT NULL"`
	IpAddr      string             `xorm:"default NULL"`
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
}

func GetIpByUID(uid int64) string {
	userLoginLog := new(UserLoginLog)
	has, err := xStatistic.Where("u_id=?", uid).Desc("id").Limit(1).Get(userLoginLog)
	if err != nil || !has {
		return ""
	}
	return userLoginLog.IpAddr
}

func SaveLoginInfoToDb(r *http.Request, u *User) {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()

	var dateRecord UserLoginLog

	dateRecord.UId = u.ID
	dateRecord.IpAddr = getIP(r)

	statictisSess.Insert(&dateRecord)
}

func getIP(r *http.Request) string {
	if setting.CustomRealUrlHeader != "" {
		realIp := r.Header.Get(setting.CustomRealUrlHeader)
		if realIp != "" {
			return realIp
		}

	}

	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
