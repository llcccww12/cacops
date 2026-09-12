// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
)

type SubjectType int

const (
	RepoSubject    SubjectType = iota //0
	DatasetSubject                    //1
	AimodelSubject                    //2
)

func GetSubjectContext(subjectId string, subjectType SubjectType) (*SubjectAccessContext, error) {
	switch subjectType {
	case DatasetSubject:
		dataset, err := GetDatasetRegistryByID(subjectId)
		if err != nil {
			return nil, err
		}
		return &SubjectAccessContext{
			SubjectType: DatasetSubject,
			SubjectID:   dataset.ID,
			OwnerID:     dataset.OwnerID,
			IsPrivate:   dataset.IsPrivate,
			Dataset:     dataset,
		}, nil
	case AimodelSubject:
		aimodel, err := GetAimodelByID(subjectId)
		if err != nil {
			return nil, err
		}
		return &SubjectAccessContext{
			SubjectType: AimodelSubject,
			SubjectID:   aimodel.ID,
			OwnerID:     aimodel.OwnerID,
			IsPrivate:   aimodel.IsPrivate,
			Aimodel:     aimodel,
		}, nil
	}
	return nil, nil
}

type SubjectAccess struct {
	ID          int64  `xorm:"pk autoincr"`
	UserID      int64  `xorm:"UNIQUE(s) INDEX "`
	SubjectID   string `xorm:"uuid UNIQUE(s) INDEX "`
	SubjectType int
	Mode        AccessMode
}

func subjectAccessLevel(e Engine, user *User, ctx *SubjectAccessContext) (AccessMode, error) {
	mode := AccessModeNone
	var userID int64
	restricted := false

	if user != nil {
		userID = user.ID
		restricted = user.IsRestricted
	}

	if !restricted && !ctx.IsPrivate {
		mode = AccessModeRead
	}

	if userID == 0 {
		return mode, nil
	}

	if userID == ctx.OwnerID {
		return AccessModeOwner, nil
	}

	a := &SubjectAccess{UserID: userID, SubjectID: ctx.SubjectID, SubjectType: int(ctx.SubjectType)}
	if has, err := e.Get(a); !has || err != nil {
		return mode, err
	}
	return a.Mode, nil
}

func GetAllAccessUser(datasetId string, mode AccessMode) ([]*User, error) {
	return getAllAccessUser(x, datasetId, mode)
}

func getAllAccessUser(e Engine, datasetId string, mode AccessMode) ([]*User, error) {
	userIds := make([]int64, 0)
	err := e.Table("subject_access").Cols("user_id").Where("subject_id = ? and subject_type = ? and mode >= ?", datasetId, DatasetSubject, mode).Find(&userIds)
	if err != nil {
		return nil, err
	}
	if len(userIds) == 0 {
		return nil, nil
	}
	users := make([]*User, 0, len(userIds))
	err = e.In("id", userIds).Find(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

type SubjectAccessContext struct {
	SubjectType SubjectType
	SubjectID   string
	OwnerID     int64
	IsPrivate   bool
	Owner       *User
	Dataset     *DatasetRegistry
	Aimodel     *AiModelManage
}

func (ctx *SubjectAccessContext) getOwner(e Engine) (err error) {
	if ctx.Owner != nil {
		return nil
	}

	ctx.Owner, err = getUserByID(e, ctx.OwnerID)
	return err
}

func (ctx *SubjectAccessContext) GetOwner() (err error) {
	return ctx.getOwner(x)
}

func SetSubjectTeamAuthorize(t *Team, subjectType SubjectType, mode AccessMode) {
	switch subjectType {
	case DatasetSubject:
		t.DatasetAuthorize = mode
	case AimodelSubject:
		t.AimodelAuthorize = mode
	}
}

func GetSubjectTeamAuthorize(t *Team, subjectType SubjectType) AccessMode {
	switch subjectType {
	case DatasetSubject:
		return t.DatasetAuthorize
	case AimodelSubject:
		return t.AimodelAuthorize
	}
	return AccessModeNone
}

func (ctx *SubjectAccessContext) recalculateTeamAccesses(e Engine, ignTeamID int64, subjectType SubjectType) (err error) {
	accessMap := make(map[int64]*userAccess, 20)

	if err = ctx.getOwner(e); err != nil {
		return err
	} else if !ctx.Owner.IsOrganization() {
		return fmt.Errorf("owner is not an organization: %d", ctx.OwnerID)
	}

	if err = ctx.refreshCollaboratorAccesses(e, accessMap, subjectType); err != nil {
		return fmt.Errorf("refreshCollaboratorAccesses: %v", err)
	}

	if err = ctx.Owner.getTeams(e); err != nil {
		return err
	}

	for _, t := range ctx.Owner.Teams {
		if t.ID == ignTeamID {
			continue
		}

		// Owner team gets owner access, and skip for teams that do not
		// have relations with repository.
		if t.IsOwnerTeam() {
			//t.DatasetAuthorize = AccessModeOwner
			SetSubjectTeamAuthorize(t, subjectType, AccessModeOwner)
		} else if !t.hasSubject(e, ctx.SubjectID, ctx.SubjectType) {
			continue
		}

		if err = t.getMembers(e); err != nil {
			return fmt.Errorf("getMembers '%d': %v", t.ID, err)
		}
		for _, m := range t.Members {
			//updateUserAccess(accessMap, m, t.DatasetAuthorize)
			subjectAuthorize := GetSubjectTeamAuthorize(t, subjectType)
			updateUserAccess(accessMap, m, subjectAuthorize)
		}
	}

	return ctx.refreshAccesses(e, accessMap)
}

func (ctx *SubjectAccessContext) refreshAccesses(e Engine, accessMap map[int64]*userAccess) (err error) {
	minMode := AccessModeRead
	if !ctx.IsPrivate {
		minMode = AccessModeWrite
	}

	newAccesses := make([]SubjectAccess, 0, len(accessMap))
	for userID, ua := range accessMap {
		if ua.Mode < minMode && !ua.User.IsRestricted {
			continue
		}

		newAccesses = append(newAccesses, SubjectAccess{
			UserID:      userID,
			SubjectID:   ctx.SubjectID,
			SubjectType: int(ctx.SubjectType),
			Mode:        ua.Mode,
		})
	}

	if _, err = e.Delete(&SubjectAccess{SubjectID: ctx.SubjectID, SubjectType: int(ctx.SubjectType)}); err != nil {
		return fmt.Errorf("delete old subject accesses: %v", err)
	}
	if len(newAccesses) == 0 {
		return nil
	}

	if _, err = e.Insert(newAccesses); err != nil {
		return fmt.Errorf("insert new accesses: %v", err)
	}
	return nil
}

// refreshCollaboratorAccesses retrieves repository collaborations with their access modes.
func (ctx *SubjectAccessContext) refreshCollaboratorAccesses(e Engine, accessMap map[int64]*userAccess, subjectType SubjectType) error {
	collaborators, err := ctx.getCollaborators(e, ListOptions{})
	if err != nil {
		return fmt.Errorf("getCollaborations: %v", err)
	}
	for _, c := range collaborators {
		updateUserAccess(accessMap, c.User, c.Collaboration.Mode)
	}
	return nil
}

func (ctx *SubjectAccessContext) recalculateUserAccess(e Engine, uid int64) (err error) {
	minMode := AccessModeRead
	if !ctx.IsPrivate {
		minMode = AccessModeWrite
	}

	accessMode := AccessModeNone
	collaborator, err := ctx.getCollaboration(e, uid)
	if err != nil {
		return err
	} else if collaborator != nil {
		accessMode = collaborator.Mode
	}

	if err = ctx.getOwner(e); err != nil {
		return err
	} else if ctx.Owner.IsOrganization() {
		var teams []Team
		if err := e.Join("INNER", "team_subject", "team_subject.team_id = team.id").
			Join("INNER", "team_user", "team_user.team_id = team.id").
			Where("team.org_id = ?", ctx.OwnerID).
			And("team_subject.subject_id=?", ctx.SubjectID).
			And("team_subject.subject_type=?", ctx.SubjectType).
			And("team_user.uid=?", uid).
			Find(&teams); err != nil {
			return err
		}

		for _, t := range teams {
			teamAccess := GetSubjectTeamAuthorize(&t, ctx.SubjectType)
			if t.IsOwnerTeam() {
				teamAccess = AccessModeOwner
			}
			accessMode = maxAccessMode(accessMode, teamAccess)
		}
	}

	if _, err = e.Delete(&SubjectAccess{SubjectID: ctx.SubjectID, SubjectType: int(ctx.SubjectType), UserID: uid}); err != nil {
		return fmt.Errorf("delete old user subject accesses: %v", err)
	} else if accessMode >= minMode {
		if _, err = e.Insert(&SubjectAccess{SubjectID: ctx.SubjectID, SubjectType: int(ctx.SubjectType), UserID: uid, Mode: accessMode}); err != nil {
			return fmt.Errorf("insert new user subject accesses: %v", err)
		}
	}
	return nil
}

func (ctx *SubjectAccessContext) recalculateAccesses(e Engine) error {
	ctx.getOwner(e)
	if err := ctx.getOwner(e); err != nil {
		return err
	}
	if ctx.Owner.IsOrganization() {
		return ctx.recalculateTeamAccesses(e, 0, ctx.SubjectType)
	}

	accessMap := make(map[int64]*userAccess, 20)
	if err := ctx.refreshCollaboratorAccesses(e, accessMap, ctx.SubjectType); err != nil {
		return fmt.Errorf("refreshCollaboratorAccesses: %v", err)
	}
	return ctx.refreshAccesses(e, accessMap)
}
