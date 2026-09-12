package models

import (
	"fmt"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
)

// Follow represents relations of user and his/her followers.
type Invitation struct {
	ID                int64              `xorm:"pk autoincr"`
	SrcUserID         int64              `xorm:"NOT NULL DEFAULT 0"`
	UserID            int64              `xorm:"NOT NULL DEFAULT 0"`
	Phone             string             `xorm:"INDEX"`
	Email             string             `xorm:"-"`
	Avatar            string             `xorm:"-"`
	Name              string             `xorm:"-"`
	InvitationUserNum int                `xorm:"-"`
	IsActive          bool               `xorm:"-"`
	CreatedUnix       timeutil.TimeStamp `xorm:"created"`
}

func QueryInvitaionByPhone(phone string) []*Invitation {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()
	cond := "phone ='" + phone + "'"
	invitationList := make([]*Invitation, 0)
	if err := statictisSess.Table(new(Invitation)).Where(cond).
		Find(&invitationList); err != nil {
		return nil
	} else {
		return invitationList
	}
}

func GetAllUserName() map[int64]string {
	sess := x.NewSession()
	defer sess.Close()
	sess.Select("id,name").Table("user")
	userList := make([]*User, 0)
	reMap := make(map[int64]string)
	sess.Find(&userList)
	for _, user := range userList {
		reMap[user.ID] = user.Name
	}
	return reMap
}

func QueryInvitaionPage(start int, pageSize int) ([]*Invitation, int64) {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()
	//cond := "created_unix >=" + fmt.Sprint(startTime) + " and created_unix <=" + fmt.Sprint(endTime)

	allCount, err := statictisSess.Count(new(Invitation))
	if err != nil {
		log.Info("query error." + err.Error())
		return nil, 0
	}
	invitationList := make([]*Invitation, 0)
	if err := statictisSess.Table(new(Invitation)).OrderBy("created_unix desc").Limit(pageSize, start).
		Find(&invitationList); err != nil {
		return nil, 0
	}
	return invitationList, allCount
}

func QueryInvitaionByTime(startTime int64, endTime int64) []*Invitation {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()
	cond := "created_unix >=" + fmt.Sprint(startTime) + " and created_unix <=" + fmt.Sprint(endTime)
	invitationList := make([]*Invitation, 0)
	if err := statictisSess.Table(new(Invitation)).Where(cond).OrderBy("created_unix desc").
		Find(&invitationList); err != nil {
		return nil
	}
	return invitationList
}

func InsertInvitaion(invitationUser *Invitation) error {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()
	_, err := statictisSess.Insert(invitationUser)
	return err
}

func QueryInvitaionBySrcUserId(srcUserId int64, start int, pageSize int) ([]*Invitation, int64) {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()
	cond := "src_user_id =" + fmt.Sprint(srcUserId)
	allCount, err := statictisSess.Where(cond).Count(new(Invitation))
	if err != nil {
		log.Info("query error." + err.Error())
		return nil, 0
	}
	invitationList := make([]*Invitation, 0)

	if err := statictisSess.Table(new(Invitation)).Where(cond).OrderBy("created_unix desc").Limit(pageSize, start).
		Find(&invitationList); err != nil {
		return nil, 0
	}
	return invitationList, allCount
}
