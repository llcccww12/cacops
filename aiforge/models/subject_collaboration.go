// Copyright 2016 The Gogs Authors. All rights reserved.
// Copyright 2020 The Gitea Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"

	"code.gitea.io/gitea/modules/log"
)

type SubjectCollaboration struct {
	ID          int64  `xorm:"pk autoincr"`
	SubjectID   string `xorm:"uuid UNIQUE(s) INDEX NOT NULL"`
	SubjectType int
	UserID      int64      `xorm:"UNIQUE(s) INDEX NOT NULL"`
	Mode        AccessMode `xorm:"DEFAULT 2 NOT NULL"`
}

func (ctx *SubjectAccessContext) getCollaborations(e Engine, listOptions ListOptions) ([]*SubjectCollaboration, error) {
	if listOptions.Page == 0 {
		collaborations := make([]*SubjectCollaboration, 0, 8)
		return collaborations, e.Find(&collaborations, &SubjectCollaboration{SubjectID: ctx.SubjectID, SubjectType: int(ctx.SubjectType)})
	}

	e = listOptions.setEnginePagination(e)

	collaborations := make([]*SubjectCollaboration, 0, listOptions.PageSize)
	return collaborations, e.Find(&collaborations, &SubjectCollaboration{SubjectID: ctx.SubjectID, SubjectType: int(ctx.SubjectType)})
}

// Collaborator represents a user with collaboration details.
type SubjectCollaborator struct {
	*User
	Collaboration *SubjectCollaboration
}

type SubjectCollaborator4Front struct {
	User          *User4Front
	Collaboration *SubjectCollaboration
}

func (s *SubjectCollaborator) ToFrontFormat() *SubjectCollaborator4Front {
	var user *User4Front
	if s.User != nil {
		user = s.User.ToFrontFormat()
	}
	return &SubjectCollaborator4Front{
		User:          user,
		Collaboration: s.Collaboration,
	}
}

func (ctx *SubjectAccessContext) getCollaborators(e Engine, listOptions ListOptions) ([]*SubjectCollaborator, error) {
	collaborations, err := ctx.getCollaborations(e, listOptions)
	if err != nil {
		return nil, fmt.Errorf("getCollaborations: %v", err)
	}

	collaborators := make([]*SubjectCollaborator, 0, len(collaborations))
	for i := 0; i < len(collaborations); i++ {
		c := collaborations[i]
		user, err := getUserByID(e, c.UserID)
		if err != nil {
			if IsErrUserNotExist(err) {
				log.Error("getCollaborators err.user not exists.userid = %d subject_id = %s", c.UserID, ctx.SubjectID)
				continue
			}
			return nil, err
		}
		collaborators = append(collaborators, &SubjectCollaborator{
			User:          user,
			Collaboration: c,
		})
	}
	return collaborators, nil
}

func (ctx *SubjectAccessContext) GetCollaborators(listOptions ListOptions) ([]*SubjectCollaborator, error) {
	return ctx.getCollaborators(x, listOptions)
}

func (ctx *SubjectAccessContext) getCollaboration(e Engine, uid int64) (*SubjectCollaboration, error) {
	collaboration := &SubjectCollaboration{
		SubjectID:   ctx.SubjectID,
		SubjectType: int(ctx.SubjectType),
		UserID:      uid,
	}
	has, err := e.Get(collaboration)
	if !has {
		collaboration = nil
	}
	return collaboration, err
}

func (ctx *SubjectAccessContext) isCollaborator(e Engine, userID int64) (bool, error) {
	return e.Get(&SubjectCollaboration{
		SubjectID:   ctx.SubjectID,
		SubjectType: int(ctx.SubjectType),
		UserID:      userID,
	})
}

func (ctx *SubjectAccessContext) IsCollaborator(userID int64) (bool, error) {
	return ctx.isCollaborator(x, userID)
}

func (ctx *SubjectAccessContext) AddCollaborator(u *User) error {
	sess := x.NewSession()
	defer sess.Close()
	if err := sess.Begin(); err != nil {
		return err
	}

	if err := ctx.addCollaborator(sess, u); err != nil {
		return err
	}

	return sess.Commit()
}

func (ctx *SubjectAccessContext) addCollaborator(e Engine, u *User) error {
	collaboration := &SubjectCollaboration{
		SubjectID:   ctx.SubjectID,
		SubjectType: int(ctx.SubjectType),
		UserID:      u.ID,
	}

	has, err := e.Get(collaboration)
	if err != nil {
		return err
	} else if has {
		return nil
	}
	collaboration.Mode = AccessModeWrite

	if _, err = e.InsertOne(collaboration); err != nil {
		return err
	}

	return ctx.recalculateUserAccess(e, u.ID)
}

func (ctx *SubjectAccessContext) ChangeCollaborationAccessMode(uid int64, mode AccessMode) error {
	sess := x.NewSession()
	defer sess.Close()
	if err := sess.Begin(); err != nil {
		return err
	}

	if err := ctx.changeCollaborationAccessMode(sess, uid, mode); err != nil {
		return err
	}

	return sess.Commit()
}

func (ctx *SubjectAccessContext) changeCollaborationAccessMode(e Engine, uid int64, mode AccessMode) error {
	// Discard invalid input
	if mode <= AccessModeNone || mode > AccessModeOwner {
		return nil
	}

	collaboration := &SubjectCollaboration{
		SubjectID:   ctx.SubjectID,
		SubjectType: int(ctx.SubjectType),
		UserID:      uid,
	}
	has, err := e.Get(collaboration)
	if err != nil {
		return fmt.Errorf("get collaboration: %v", err)
	} else if !has {
		return nil
	}

	if collaboration.Mode == mode {
		return nil
	}
	collaboration.Mode = mode

	if _, err = e.
		ID(collaboration.ID).
		Cols("mode").
		Update(collaboration); err != nil {
		return fmt.Errorf("update collaboration: %v", err)
	} else if _, err = e.Exec("UPDATE subject_access SET mode = ? WHERE user_id = ? AND subject_id = ? AND subject_type=?", mode, uid, ctx.SubjectID, ctx.SubjectType); err != nil {
		return fmt.Errorf("update subject access table: %v", err)
	}

	return nil
}

func (ctx *SubjectAccessContext) DeleteCollaboration(uid int64) (err error) {
	collaboration := &SubjectCollaboration{
		SubjectID:   ctx.SubjectID,
		SubjectType: int(ctx.SubjectType),
		UserID:      uid,
	}

	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}

	if has, err := sess.Delete(collaboration); err != nil || has == 0 {
		return err
	} else if err = ctx.recalculateAccesses(sess); err != nil {
		return err
	}

	return sess.Commit()
}

func (ctx *SubjectAccessContext) getTeams(e Engine) (teams []*Team, err error) {
	return teams, e.
		Join("INNER", "team_subject", "team_subject.team_id = team.id").
		Where("team.org_id = ?", ctx.OwnerID).
		And("team_subject.subject_id=?", ctx.SubjectID).
		And("team_subject.subject_type=?", ctx.SubjectType).
		OrderBy("CASE WHEN name LIKE '" + ownerTeamName + "' THEN '' ELSE name END").
		Find(&teams)
}

func (ctx *SubjectAccessContext) GetTeams() ([]*Team, error) {
	return ctx.getTeams(x)
}
