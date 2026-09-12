package models

import (
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
)

type PointAccountStatus int

// Possible PointAccountStatus types.
const (
	PointAccountNormal  int = iota + 1 // 1
	PointAccountFreeze                 // 2
	PointAccountDeleted                // 3
)

type PointAccount struct {
	ID            int64              `xorm:"pk autoincr"`
	AccountCode   string             `xorm:"INDEX NOT NULL"`
	Balance       float64            `xorm:"NOT NULL DEFAULT 0"`
	TotalEarned   float64            `xorm:"NOT NULL DEFAULT 0"`
	TotalConsumed float64            `xorm:"NOT NULL DEFAULT 0"`
	UserId        int64              `xorm:"INDEX NOT NULL"`
	Status        int                `xorm:"NOT NULL"`
	Version       int64              `xorm:"NOT NULL"`
	CreatedUnix   timeutil.TimeStamp `xorm:"INDEX created"`
	UpdatedUnix   timeutil.TimeStamp `xorm:"INDEX updated"`
}

func (account *PointAccount) Increase(amount float64, sourceId string) error {
	sess := x.NewSession()
	defer sess.Close()
	sql := "update point_account set balance = balance + ?,total_earned = total_earned + ? ,version = version + 1,updated_unix = ? where account_code = ? "
	_, err := sess.Exec(sql, amount, amount, timeutil.TimeStampNow(), account.AccountCode)
	if err != nil {
		sess.Rollback()
		return err
	}
	accountLog := &PointAccountLog{
		AccountCode:    account.AccountCode,
		UserId:         account.UserId,
		Type:           IncreaseAccountBalance,
		SourceId:       sourceId,
		PointsAmount:   amount,
		BalanceBefore:  account.Balance,
		BalanceAfter:   account.Balance + amount,
		AccountVersion: account.Version,
	}
	_, err = sess.Insert(accountLog)
	if err != nil {
		sess.Rollback()
		return err
	}
	sess.Commit()
	return nil
}

func (account *PointAccount) Decrease(amount float64, sourceId string) error {
	sess := x.NewSession()
	defer sess.Close()
	sql := "update point_account set balance = balance - ?,total_consumed = total_consumed + ? ,version = version + 1 ,updated_unix = ? where account_code = ? "
	_, err := sess.Exec(sql, amount, amount, timeutil.TimeStampNow(), account.AccountCode)
	if err != nil {
		sess.Rollback()
		return err
	}
	accountLog := &PointAccountLog{
		AccountCode:    account.AccountCode,
		UserId:         account.UserId,
		Type:           DecreaseAccountBalance,
		SourceId:       sourceId,
		PointsAmount:   amount,
		BalanceBefore:  account.Balance,
		BalanceAfter:   account.Balance - amount,
		AccountVersion: account.Version,
	}
	_, err = sess.Insert(accountLog)
	if err != nil {
		sess.Rollback()
		return err
	}
	sess.Commit()
	return nil
}

func GetAccountByUserId(userId int64) (*PointAccount, error) {
	p := &PointAccount{}
	has, err := x.Where("user_id = ?", userId).Get(p)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrRecordNotExist{}
	}
	return p, nil
}

func InsertAccount(tl *PointAccount) (int64, error) {
	return x.Insert(tl)
}

type SearchPointAccountOpts struct {
	ListOptions
	Keyword string
}

type SearchPointAccountResponse struct {
	Records  []*UserPointAccount
	PageSize int
	Page     int
	Total    int64
}

type UserPointAccount struct {
	UserId        int64
	UserName      string
	Email         string
	Balance       float64
	TotalEarned   float64
	TotalConsumed float64
}

func (UserPointAccount) TableName() string {
	return "user"
}

func GetPointAccountMapByUserIds(userIds []int64) (map[int64]*PointAccount, error) {
	if len(userIds) == 0 {
		return make(map[int64]*PointAccount, 0), nil
	}
	accounts := make([]*PointAccount, 0)
	err := x.In("user_id", userIds).Find(&accounts)
	if err != nil {
		log.Error("GetPointAccountMapByUserIds error.%v", err)
		return nil, err
	}
	accountMap := make(map[int64]*PointAccount, 0)
	for _, v := range accounts {
		accountMap[v.UserId] = v
	}
	return accountMap, nil
}

type PointDeductCondition struct {
	SpecUnitPrice    float64
	WorkServerNumber int
}
