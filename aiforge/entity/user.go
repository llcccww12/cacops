package entity

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/timeutil"
)

type UserBriefInfo struct {
	ID            int64              `json:"id"`
	LowerName     string             `json:"lower_name"`
	Name          string             `json:"name"`
	FullName      string             `json:"full_name"`
	Email         string             `json:"email"`
	Language      string             `json:"language"`
	Description   string             `json:"description"`
	RelAvatarLink string             `json:"rel_avatar_link"`
	NumMembers    int                `json:"num_members"`
	CreatedUnix   timeutil.TimeStamp `json:"created_unix"`
	UpdatedUnix   timeutil.TimeStamp `json:"updated_unix"`
}

func ConvertUserToBrief(u *models.User) *UserBriefInfo {
	fullName := u.Name
	if u.FullName != "" {
		fullName = u.FullName
	}
	uf := &UserBriefInfo{
		ID:          u.ID,
		LowerName:   u.LowerName,
		Name:        u.Name,
		FullName:    fullName,
		Email:       u.Email,
		Language:    u.Language,
		Description: u.Description,
		CreatedUnix: u.CreatedUnix,
		UpdatedUnix: u.UpdatedUnix,
		NumMembers:  u.NumMembers,
	}
	if !u.KeepEmailPrivate {
		uf.Email = u.Email
	}
	uf.RelAvatarLink = u.RelAvatarLink()
	return uf
}
