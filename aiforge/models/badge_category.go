package models

import "code.gitea.io/gitea/modules/timeutil"

type BadgeType int

const (
	CustomizeBadge = iota + 1
	SystemBadge
)

type BadgeCategory struct {
	ID          int64 `xorm:"pk autoincr"`
	Name        string
	Position    int64
	Type        BadgeType
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
	UpdatedUnix timeutil.TimeStamp `xorm:"updated"`
	DeletedAt   timeutil.TimeStamp `xorm:"deleted"`
}

func (m *BadgeCategory) ToShow() *BadgeCategory4Show {
	return &BadgeCategory4Show{
		ID:          m.ID,
		Name:        m.Name,
		Position:    m.Position,
		Type:        m.Type,
		CreatedUnix: m.CreatedUnix,
	}
}

type BadgeCategory4Show struct {
	ID          int64 `xorm:"pk autoincr"`
	Name        string
	Position    int64
	Type        BadgeType
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
}

func (m BadgeCategory4Show) ToDTO() BadgeCategory {
	return BadgeCategory{
		ID:          m.ID,
		Name:        m.Name,
		Position:    m.Position,
		Type:        m.Type,
		CreatedUnix: m.CreatedUnix,
	}
}

func GetBadgeCategoryListPaging(opts ListOptions) (int64, []*BadgeCategory, error) {
	n, err := x.Count(&BadgeCategory{})
	if err != nil {
		return 0, nil, err
	}
	if opts.Page <= 0 {
		opts.Page = 1
	}
	r := make([]*BadgeCategory, 0)
	if err := x.OrderBy("position asc,created_unix desc").Limit(opts.PageSize, (opts.Page-1)*opts.PageSize).Find(&r); err != nil {
		return 0, nil, err
	}
	return n, r, nil
}

func GetBadgeCategoryList() ([]*BadgeCategory, error) {
	r := make([]*BadgeCategory, 0)
	if err := x.OrderBy("position asc,created_unix desc").Find(&r); err != nil {
		return nil, err
	}
	return r, nil
}

func AddBadgeCategory(m BadgeCategory) (int64, error) {
	return x.Insert(&m)
}

func UpdateBadgeCategoryById(id int64, param BadgeCategory) (int64, error) {
	return x.ID(id).Update(&param)
}

func DelBadgeCategory(id int64) (int64, error) {
	return x.ID(id).Delete(&BadgeCategory{})
}

func GetBadgeCategoryById(id int64) (*BadgeCategory, error) {
	m := &BadgeCategory{}
	has, err := x.ID(id).Get(m)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrRecordNotExist{}
	}
	return m, nil
}
