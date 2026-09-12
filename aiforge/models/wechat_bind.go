package models

import (
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
	"fmt"
	"time"
)

type WechatBindAction int

const (
	WECHAT_BIND WechatBindAction = iota + 1
	WECHAT_UNBIND
)

type WechatBindLog struct {
	ID           int64  `xorm:"pk autoincr"`
	UserID       int64  `xorm:"INDEX"`
	WechatOpenId string `xorm:"INDEX"`
	Action       int
	CreateTime   time.Time `xorm:"INDEX created"`
}

func BindWechatOpenId(userId int64, wechatOpenId string) error {
	sess := x.NewSession()
	defer sess.Close()
	if err := sess.Begin(); err != nil {
		return err
	}

	param := &User{WechatOpenId: wechatOpenId, WechatBindUnix: timeutil.TimeStampNow()}
	n, err := sess.Where("ID = ?", userId).Update(param)
	if err != nil {
		log.Error("update wechat_open_id failed,e=%v", err)
		if e := sess.Rollback(); e != nil {
			log.Error("BindWechatOpenId: sess.Rollback: %v", e)
		}
		return err
	}
	if n == 0 {
		log.Error("update wechat_open_id failed,user not exist,userId=%d", userId)
		if e := sess.Rollback(); e != nil {
			log.Error("BindWechatOpenId: sess.Rollback: %v", e)
		}
		return nil
	}

	logParam := &WechatBindLog{
		UserID:       userId,
		WechatOpenId: wechatOpenId,
		Action:       int(WECHAT_BIND),
	}
	sess.Insert(logParam)
	return sess.Commit()
}

func GetUserWechatOpenId(userId int64) string {
	param := &User{}
	x.Cols("wechat_open_id").Where("ID =?", userId).Get(param)
	return param.WechatOpenId
}

func GetUserByWechatOpenId(wechatOpenId string) *User {
	user := &User{}
	isExists, err := x.Where("wechat_open_id = ?", wechatOpenId).Get(user)
	if err != nil || !isExists {
		return nil
	}
	return user
}

func UnbindWechatOpenId(userId int64, oldWechatOpenID string) error {
	sess := x.NewSession()
	defer sess.Close()
	if err := sess.Begin(); err != nil {
		return err
	}

	n, err := x.Table(new(User)).Where("ID = ? AND wechat_open_id =?", userId, oldWechatOpenID).Update(map[string]interface{}{"wechat_open_id": "", "wechat_bind_unix": nil})
	if err != nil {
		log.Error("update wechat_open_id failed,e=%v", err)
		if e := sess.Rollback(); e != nil {
			log.Error("UnbindWechatOpenId: sess.Rollback: %v", e)
		}
		return err
	}
	if n == 0 {
		log.Error("update wechat_open_id failed,user not exist,userId=%d", userId)
		if e := sess.Rollback(); e != nil {
			log.Error("UnbindWechatOpenId: sess.Rollback: %v", e)
		}
		return nil
	}
	logParam := &WechatBindLog{
		UserID:       userId,
		WechatOpenId: oldWechatOpenID,
		Action:       int(WECHAT_UNBIND),
	}
	sess.Insert(logParam)
	return sess.Commit()
}

func CountWechatBindLog(wechatOpenId string, action WechatBindAction) (int64, error) {
	return x.Where("wechat_open_id = ? and action = ?", wechatOpenId, action).Count(&WechatBindLog{})
}

type QRCodeResponse struct {
	Url           string
	Ticket        string
	SceneStr      string
	ExpireSeconds int
}

func NewWechatBindError(reply string) WechatBindError {
	return WechatBindError{Reply: reply}
}

type QRCode4BindCache struct {
	UserId int64
	Status int
	Type   string
	UUID   string
	Data   map[string]interface{}
}

type CreateQRCodeReq struct {
	UserId int64
	Type   string
	UUID   string
	Data   map[string]interface{}
}

const (
	BIND_STATUS_UNBIND  = 0
	BIND_STATUS_SCANNED = 1
	BIND_STATUS_BOUND   = 2
	BIND_STATUS_EXPIRED = 9
)

const (
	QR_CODE_TYPE_NORMAL  = "normal"
	QR_CODE_TYPE_SING_UP = "sign_up"
	QR_CODE_TYPE_LOGIN   = "login"
)

type WechatBindError struct {
	Reply string
}

func (err WechatBindError) Error() string {
	return fmt.Sprint("wechat bind error,reply=%s", err.Reply)
}
