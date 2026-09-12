package models

import (
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
	"path/filepath"
	"strings"
	"xorm.io/builder"
)

type Badge struct {
	ID          int64 `xorm:"pk autoincr"`
	Name        string
	LightedIcon string `xorm:"varchar(2048)"`
	GreyedIcon  string `xorm:"varchar(2048)"`
	Url         string `xorm:"varchar(2048)"`
	CategoryId  int64
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
	UpdatedUnix timeutil.TimeStamp `xorm:"updated"`
	DeletedAt   timeutil.TimeStamp `xorm:"deleted"`
}

func (m *Badge) ToUserShow() *Badge4UserShow {
	return &Badge4UserShow{
		Name:        m.Name,
		LightedIcon: GetIconOuterLink(m.LightedIcon),
		GreyedIcon:  GetIconOuterLink(m.GreyedIcon),
		Url:         m.Url,
	}
}

type GetBadgeOpts struct {
	BadgeType  BadgeType
	CategoryId int64
	ListOpts   ListOptions
}

type BadgeAndCategory struct {
	Badge    Badge         `xorm:"extends"`
	Category BadgeCategory `xorm:"extends"`
}

func (*BadgeAndCategory) TableName() string {
	return "badge"
}

func (m *BadgeAndCategory) ToShow() *Badge4AdminShow {
	return &Badge4AdminShow{
		ID:           m.Badge.ID,
		Name:         m.Badge.Name,
		LightedIcon:  GetIconOuterLink(m.Badge.LightedIcon),
		GreyedIcon:   GetIconOuterLink(m.Badge.GreyedIcon),
		Url:          m.Badge.Url,
		CategoryName: m.Category.Name,
		CategoryId:   m.Category.ID,
		CreatedUnix:  m.Badge.CreatedUnix,
		UpdatedUnix:  m.Badge.UpdatedUnix,
	}
}

type Badge4AdminShow struct {
	ID           int64
	Name         string
	LightedIcon  string
	GreyedIcon   string
	Url          string
	CategoryName string
	CategoryId   int64
	CreatedUnix  timeutil.TimeStamp
	UpdatedUnix  timeutil.TimeStamp
}

func (m Badge4AdminShow) ToDTO() Badge {
	return Badge{
		Name:        m.Name,
		LightedIcon: m.LightedIcon,
		GreyedIcon:  m.GreyedIcon,
		Url:         m.Url,
		CategoryId:  m.CategoryId,
	}
}

type BadgeOperateReq struct {
	ID          int64
	Name        string
	LightedIcon string
	GreyedIcon  string
	Url         string
	CategoryId  int64
}

func (m BadgeOperateReq) ToDTO() Badge {
	return Badge{
		Name:        m.Name,
		LightedIcon: m.LightedIcon,
		GreyedIcon:  m.GreyedIcon,
		Url:         m.Url,
		CategoryId:  m.CategoryId,
	}
}

type Badge4UserShow struct {
	Name        string
	LightedIcon string
	GreyedIcon  string
	Url         string
}

type BadgeShowWithStatus struct {
	Badge     *Badge4UserShow
	IsLighted bool
}

type UserAllBadgeInCategory struct {
	CategoryName string
	CategoryId   int64
	LightedNum   int
	Badges       []*BadgeShowWithStatus
}

func GetBadgeList(opts GetBadgeOpts) (int64, []*BadgeAndCategory, error) {
	if opts.ListOpts.Page <= 0 {
		opts.ListOpts.Page = 1
	}
	var cond = builder.NewCond()
	if opts.BadgeType > 0 {
		cond = cond.And(builder.Eq{"badge_category.type": opts.BadgeType})
	}
	if opts.CategoryId > 0 {
		cond = cond.And(builder.Eq{"badge_category.id": opts.CategoryId})
	}
	n, err := x.Join("INNER", "badge_category", "badge_category.ID = badge.category_id").Where(cond).Count(&BadgeAndCategory{})
	if err != nil {
		return 0, nil, err
	}
	r := make([]*BadgeAndCategory, 0)
	if err = x.Join("INNER", "badge_category", "badge_category.ID = badge.category_id").Where(cond).OrderBy("badge.created_unix desc").Limit(opts.ListOpts.PageSize, (opts.ListOpts.Page-1)*opts.ListOpts.PageSize).Find(&r); err != nil {
		return 0, nil, err
	}
	return n, r, nil
}

func AddBadge(m Badge) (int64, error) {
	return x.Insert(&m)
}

func UpdateBadgeById(id int64, param Badge) (int64, error) {
	return x.ID(id).Update(&param)
}

func DelBadge(id int64) (int64, error) {
	return x.ID(id).Delete(&Badge{})
}

func GetBadgeById(id int64) (*Badge, error) {
	m := &Badge{}
	has, err := x.ID(id).Get(m)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, &ErrRecordNotExist{}
	}
	return m, nil
}

func GetBadgeByCategoryId(categoryId int64) ([]*Badge, error) {
	r := make([]*Badge, 0)
	err := x.Where("category_id = ?", categoryId).Find(&r)
	return r, err
}

func GetCustomIconByHash(hash string) string {
	if len(hash) == 0 {
		return ""
	}
	return filepath.Join(setting.IconUploadPath, hash)
}

func GetIconOuterLink(hash string) string {
	return strings.TrimRight(setting.AppSubURL, "/") + "/show/icon/" + hash
}
