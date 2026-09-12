package models

import (
	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
)

type LimitType string

const (
	LimitTypeTask        LimitType = "TASK"
	LimitTypeRewardPoint LimitType = "REWARD_POINT"
)

func (l LimitType) Name() string {
	switch l {
	case LimitTypeTask:
		return "TASK"
	case LimitTypeRewardPoint:
		return "REWARD_POINT"
	default:
		return ""
	}
}

type LimitScope string

const (
	LimitScopeAllUsers   LimitScope = "ALL_USERS"
	LimitScopeSingleUser LimitScope = "SINGLE_USER"
)

func (l LimitScope) Name() string {
	switch l {
	case LimitScopeAllUsers:
		return "ALL_USERS"
	case LimitScopeSingleUser:
		return "SINGLE_USER"
	default:
		return ""
	}
}

type LimiterRejectPolicy string

const (
	JustReject    LimiterRejectPolicy = "JUST_REJECT"
	PermittedOnce LimiterRejectPolicy = "PERMITTED_ONCE"
	FillUp        LimiterRejectPolicy = "FillUp"
)

type LimitConfig struct {
	ID          int64 `xorm:"pk autoincr"`
	Title       string
	RefreshRate string  `xorm:"NOT NULL"`
	Scope       string  `xorm:"NOT NULL"`
	LimitNum    float64 `xorm:"NOT NULL"`
	LimitCode   string
	LimitType   string `xorm:"NOT NULL"`
	RelatedId   int64  `xorm:"INDEX"`
	CreatorId   int64  `xorm:"NOT NULL"`
	CreatorName string
	DeleterId   int64
	DeleterName string
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
	DeletedAt   timeutil.TimeStamp `xorm:"deleted"`
}

type LimitConfigQueryOpts struct {
	RefreshRate string
	Scope       LimitScope
	LimitCode   string
	LimitType   LimitType
}

type LimitConfigVO struct {
	ID          int64
	Title       string
	RefreshRate string
	Scope       string
	LimitNum    float64
	LimitCode   string
	LimitType   string
	Creator     string
	CreatedUnix timeutil.TimeStamp
}

func (l *LimitConfig) ToLimitConfigVO() *LimitConfigVO {
	return &LimitConfigVO{
		ID:          l.ID,
		Title:       l.Title,
		RefreshRate: l.RefreshRate,
		Scope:       l.Scope,
		LimitNum:    l.LimitNum,
		LimitCode:   l.LimitCode,
		LimitType:   l.LimitType,
		Creator:     l.CreatorName,
		CreatedUnix: l.CreatedUnix,
	}
}

func GetLimitConfigByLimitType(limitType LimitType) ([]LimitConfig, error) {
	r := make([]LimitConfig, 0)
	err := x.Where(" limit_type = ?", limitType.Name()).Find(&r)
	if err != nil {
		return nil, err
	} else if len(r) == 0 {
		return nil, ErrRecordNotExist{}
	}
	return r, nil
}

func GetLimitersByRelatedIdWithDeleted(limitType LimitType) ([]LimitConfig, error) {
	r := make([]LimitConfig, 0)
	err := x.Unscoped().Where("  = ?", limitType.Name()).Find(&r)
	if err != nil {
		return nil, err
	} else if len(r) == 0 {
		return nil, ErrRecordNotExist{}
	}
	return r, nil
}

func AddLimitConfig(l *LimitConfig) error {
	sess := x.NewSession()
	defer sess.Close()

	//delete old limit config
	cond := builder.NewCond()
	cond = cond.And(builder.Eq{"limit_type": l.LimitType})
	cond = cond.And(builder.Eq{"scope": l.Scope})
	if l.LimitCode == "" {
		subCond := builder.NewCond()
		subCond = subCond.Or(builder.IsNull{"limit_code"})
		subCond = subCond.Or(builder.Eq{"limit_code": ""})
		cond = cond.And(subCond)
	} else {
		cond = cond.And(builder.Eq{"limit_code": l.LimitCode})
	}
	_, err := sess.Where(cond).Delete(&LimitConfig{})
	if err != nil {
		sess.Rollback()
		return err
	}

	//add new config
	_, err = sess.Insert(l)
	if err != nil {
		sess.Rollback()
		return err
	}

	sess.Commit()
	return nil
}

func DeleteLimitConfig(config LimitConfig, deleterId int64, deleterName string) error {
	sess := x.NewSession()
	defer sess.Close()

	_, err := x.ID(config.ID).Update(&LimitConfig{DeleterName: deleterName, DeleterId: deleterId})
	if err != nil {
		sess.Rollback()
		return err
	}
	_, err = x.ID(config.ID).Delete(&LimitConfig{})
	if err != nil {
		sess.Rollback()
		return err
	}
	sess.Commit()
	return nil
}

func GetLimitConfigById(id int64) (*LimitConfig, error) {
	r := &LimitConfig{}
	isOk, err := x.ID(id).Get(r)
	if err != nil {
		return nil, err
	} else if !isOk {
		return nil, nil
	}
	return r, nil
}
