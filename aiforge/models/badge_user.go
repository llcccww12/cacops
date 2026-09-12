package models

import (
	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
)

const (
	ActionAddBadgeUser = 1
	ActionDelBadgeUser = 2
)

type BadgeUser struct {
	ID          int64              `xorm:"pk autoincr"`
	UserId      int64              `xorm:"unique(user_badge)"`
	BadgeId     int64              `xorm:"unique(user_badge) index"`
	CreatedUnix timeutil.TimeStamp `xorm:"created index"`
}

type BadgeUserLog struct {
	ID          int64 `xorm:"pk autoincr"`
	UserId      int64 `xorm:"index"`
	BadgeId     int64 `xorm:"index"`
	Action      int
	CreatedUnix timeutil.TimeStamp `xorm:"created index"`
}

type BadgeUserDetail struct {
	BadgeUser BadgeUser `xorm:"extends"`
	User      User      `xorm:"extends"`
}

func (*BadgeUserDetail) TableName() string {
	return "badge_user"
}

func (m *BadgeUserDetail) ToShow() *BadgeUser4SHow {
	return &BadgeUser4SHow{
		ID:          m.BadgeUser.ID,
		UserId:      m.BadgeUser.UserId,
		Name:        m.User.Name,
		Avatar:      m.User.RelAvatarLink(),
		Email:       m.User.Email,
		CreatedUnix: m.BadgeUser.CreatedUnix,
	}
}

type BadgeUser4SHow struct {
	ID          int64
	UserId      int64
	Name        string
	Avatar      string
	Email       string
	CreatedUnix timeutil.TimeStamp
}

type AddBadgeUsersReq struct {
	BadgeId int64
	Users   string
}
type DelBadgeUserReq struct {
	ID int64
}

type GetUserBadgesOpts struct {
	CategoryId int64
	ListOptions
}

func AddBadgeUser(m BadgeUser) (int64, error) {
	sess := x.NewSession()
	defer sess.Close()
	sess.Begin()
	n, err := sess.Insert(&m)
	if err != nil || n == 0 {
		return 0, err
	}
	_, err = sess.Insert(&BadgeUserLog{
		UserId:  m.UserId,
		BadgeId: m.BadgeId,
		Action:  ActionAddBadgeUser,
	})
	if err != nil {
		sess.Rollback()
		return 0, err
	}
	return n, sess.Commit()
}

func DelBadgeUser(id int64) (int64, error) {
	m := BadgeUser{}
	has, err := x.ID(id).Get(&m)
	if err != nil {
		return 0, err
	}
	if !has {
		return 0, ErrRecordNotExist{}
	}
	sess := x.NewSession()
	defer sess.Close()
	sess.Begin()
	n, err := x.ID(m.ID).Delete(&BadgeUser{})
	if err != nil || n == 0 {
		return 0, err
	}
	_, err = sess.Insert(&BadgeUserLog{
		UserId:  m.UserId,
		BadgeId: m.BadgeId,
		Action:  ActionDelBadgeUser,
	})
	if err != nil {
		sess.Rollback()
		return 0, err
	}
	return n, sess.Commit()
}

func GetBadgeUsers(badgeId int64, opts ListOptions) (int64, []BadgeUserDetail, error) {
	n, err := x.Join("LEFT", "public.user", "public.user.ID = badge_user.user_id").Where("badge_user.badge_id = ?", badgeId).Count(&BadgeUserDetail{})
	if err != nil {
		return 0, nil, err
	}
	if opts.Page <= 0 {
		opts.Page = 1
	}
	m := make([]BadgeUserDetail, 0)
	err = x.Join("LEFT", "public.user", "public.user.ID = badge_user.user_id").Where("badge_user.badge_id = ?", badgeId).OrderBy("badge_user.id desc").Limit(opts.PageSize, (opts.Page-1)*opts.PageSize).Find(&m)
	if err != nil {
		return 0, nil, err
	}
	return n, m, nil
}

func GetUserBadgesPaging(userId int64, opts GetUserBadgesOpts) ([]*Badge, error) {
	cond := builder.NewCond()
	cond = cond.And(builder.Eq{"badge_user.user_id": userId})
	if opts.CategoryId > 0 {
		cond = cond.And(builder.Eq{"badge.category_id": opts.CategoryId})
	}

	r := make([]*Badge, 0)
	err := x.Join("INNER", "badge_user", "badge_user.badge_id = badge.id").Where(cond).OrderBy("badge_user.id desc").Limit(opts.PageSize, (opts.Page-1)*opts.PageSize).Find(&r)
	return r, err
}
func CountUserBadges(userId int64) (int64, error) {
	return x.Where("user_id = ?", userId).Count(&BadgeUser{})
}

func GetUserBadges(userId, categoryId int64) ([]*Badge, error) {
	cond := builder.NewCond()
	cond = cond.And(builder.Eq{"badge_user.user_id": userId})
	if categoryId > 0 {
		cond = cond.And(builder.Eq{"badge.category_id": categoryId})
	}

	r := make([]*Badge, 0)
	err := x.Join("INNER", "badge_user", "badge_user.badge_id = badge.id").Where(cond).OrderBy("badge_user.created_unix desc").Find(&r)
	return r, err
}
