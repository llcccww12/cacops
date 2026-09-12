// Copyright 2018 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
	"strconv"

	"xorm.io/builder"
)

// ActionList defines a list of actions
type ActionList []*Action

func (actions ActionList) getUserIDs() []int64 {
	userIDs := make(map[int64]struct{}, len(actions))
	for _, action := range actions {
		if _, ok := userIDs[action.ActUserID]; !ok {
			userIDs[action.ActUserID] = struct{}{}
		}
	}
	return keysInt64(userIDs)
}

func (actions ActionList) loadUsers(e Engine) ([]*User, error) {
	if len(actions) == 0 {
		return nil, nil
	}

	userIDs := actions.getUserIDs()
	userMaps := make(map[int64]*User, len(userIDs))
	if len(userIDs) == 0 {
		return make([]*User, 0), nil
	}
	err := e.
		In("id", userIDs).
		Find(&userMaps)
	if err != nil {
		return nil, fmt.Errorf("find user: %v", err)
	}

	for _, action := range actions {
		action.ActUser = userMaps[action.ActUserID]
	}
	return valuesUser(userMaps), nil
}

// LoadUsers loads actions' all users
func (actions ActionList) LoadUsers() ([]*User, error) {
	return actions.loadUsers(x)
}

func (actions ActionList) getRepoIDs() []int64 {
	repoIDs := make(map[int64]struct{}, len(actions))
	for _, action := range actions {
		if action.RepoID == 0 {
			continue
		}
		if _, ok := repoIDs[action.RepoID]; !ok {
			repoIDs[action.RepoID] = struct{}{}
		}
	}
	return keysInt64(repoIDs)
}

func (actions ActionList) getDatasetIDs() []string {
	datasetIDs := make(map[string]struct{}, len(actions))
	for _, action := range actions {
		if action.DatasetID == nil || *action.DatasetID == "" {
			continue
		}
		if _, ok := datasetIDs[*action.DatasetID]; !ok {
			datasetIDs[*action.DatasetID] = struct{}{}
		}
	}
	return keysString(datasetIDs)
}

func (actions ActionList) getAimodelIDs() []string {
	aimodelIDs := make(map[string]struct{}, len(actions))
	for _, action := range actions {
		if action.AimodelID == nil || *action.AimodelID == "" {
			continue
		}
		if _, ok := aimodelIDs[*action.AimodelID]; !ok {
			aimodelIDs[*action.AimodelID] = struct{}{}
		}
	}
	return keysString(aimodelIDs)
}

func (actions ActionList) loadRepositories(e Engine) ([]*Repository, error) {
	if len(actions) == 0 {
		return nil, nil
	}

	repoIDs := actions.getRepoIDs()
	repoMaps := make(map[int64]*Repository, len(repoIDs))
	if len(repoIDs) == 0 {
		return make([]*Repository, 0), nil
	}
	err := e.
		In("id", repoIDs).
		Find(&repoMaps)
	if err != nil {
		return nil, fmt.Errorf("find repository: %v", err)
	}

	for _, action := range actions {
		action.Repo = repoMaps[action.RepoID]
	}
	return valuesRepository(repoMaps), nil
}

// LoadRepositories loads actions' all repositories
func (actions ActionList) LoadRepositories() ([]*Repository, error) {
	return actions.loadRepositories(x)
}

func (actions ActionList) loadDatasets(e Engine) ([]*DatasetRegistry, error) {
	if len(actions) == 0 {
		return nil, nil
	}

	datasetIDs := actions.getDatasetIDs()
	datasetMaps := make(map[string]*DatasetRegistry, len(datasetIDs))
	if len(datasetIDs) == 0 {
		return make([]*DatasetRegistry, 0), nil
	}
	err := e.
		In("id", datasetIDs).
		Find(&datasetMaps)
	if err != nil {
		return nil, fmt.Errorf("find datasets: %v", err)
	}

	for _, action := range actions {
		if action.DatasetID == nil || *action.DatasetID == "" {
			continue
		}
		dataset := datasetMaps[*action.DatasetID]
		if dataset != nil {
			dataset.GetOwner()
		}
		action.Dataset = dataset
	}
	return valuesDataset(datasetMaps), nil
}

// LoadRepositories loads actions' all repositories
func (actions ActionList) LoadDatasets() ([]*DatasetRegistry, error) {
	return actions.loadDatasets(x)
}

// LoadRepositories loads actions' all repositories
func (actions ActionList) LoadAimodels() ([]*AiModelManage, error) {
	return actions.loadAimodels(x)
}

func (actions ActionList) loadAimodels(e Engine) ([]*AiModelManage, error) {
	if len(actions) == 0 {
		return nil, nil
	}

	aimodelIDs := actions.getAimodelIDs()
	aimodelMaps := make(map[string]*AiModelManage, len(aimodelIDs))
	if len(aimodelIDs) == 0 {
		return make([]*AiModelManage, 0), nil
	}
	err := e.
		In("id", aimodelIDs).
		Find(&aimodelMaps)
	if err != nil {
		return nil, fmt.Errorf("find datasets: %v", err)
	}

	for _, action := range actions {
		if action.AimodelID == nil || *action.AimodelID == "" {
			continue
		}
		aimodel := aimodelMaps[*action.AimodelID]
		if aimodel != nil {
			aimodel.GetOwner()
		}
		action.Aimodel = aimodel
	}
	return valuesAimodel(aimodelMaps), nil
}

func (actions ActionList) getCommentIDs() []int64 {
	commentIDs := make(map[int64]struct{}, len(actions))
	for _, action := range actions {
		if action.CommentID == 0 {
			continue
		}
		if _, ok := commentIDs[action.CommentID]; !ok {
			commentIDs[action.CommentID] = struct{}{}
		}
	}
	return keysInt64(commentIDs)
}

func (actions ActionList) loadComments(e Engine) ([]*Comment, error) {
	if len(actions) == 0 {
		return nil, nil
	}

	commentIDs := actions.getCommentIDs()

	commentMaps := make(map[int64]*Comment, len(commentIDs))
	if len(commentIDs) == 0 {
		return make([]*Comment, 0), nil
	}
	err := e.
		In("id", commentIDs).
		Find(&commentMaps)
	if err != nil {
		return nil, fmt.Errorf("find comment: %v", err)
	}

	for _, action := range actions {
		if action.CommentID > 0 {
			action.Comment = commentMaps[action.CommentID]
		}
	}
	return valuesComment(commentMaps), nil
}

// LoadComments loads actions' all comments
func (actions ActionList) LoadComments() ([]*Comment, error) {
	return actions.loadComments(x)
}

func (actions ActionList) getCloudbrainIDs() []int64 {
	cloudbrainIDs := make(map[int64]struct{}, 0)
	for _, action := range actions {
		if !action.IsCloudbrainAction() {
			continue
		}
		cloudbrainId, _ := strconv.ParseInt(action.Content, 10, 64)
		if _, ok := cloudbrainIDs[cloudbrainId]; !ok {
			cloudbrainIDs[cloudbrainId] = struct{}{}
		}
	}
	return keysInt64(cloudbrainIDs)
}

func (actions ActionList) getCloudbrainJobIDs() []string {
	cloudbrainJobIDs := make(map[string]struct{}, 0)
	for _, action := range actions {
		if !action.IsCloudbrainAction() {
			continue
		}
		if _, ok := cloudbrainJobIDs[action.Content]; !ok {
			cloudbrainJobIDs[action.Content] = struct{}{}
		}
	}
	return keysString(cloudbrainJobIDs)
}

func (actions ActionList) loadCloudbrains(e Engine) ([]*Cloudbrain, error) {
	if len(actions) == 0 {
		return nil, nil
	}
	cloudbrainIDs := actions.getCloudbrainIDs()
	cloudbrainJobIDs := actions.getCloudbrainJobIDs()

	cloudbrainMaps := make(map[int64]*Cloudbrain, len(cloudbrainIDs))
	if len(cloudbrainIDs) == 0 {
		return make([]*Cloudbrain, 0), nil
	}
	//由于各个类型的云脑任务在发布action的时候，content字段保存的ID含义不同，部分取的是ID,部分取的是jobId
	//所以在查询action对应的cloudbrain对象时，以这两个字段做为条件查询
	cond := builder.Or(builder.In("id", cloudbrainIDs)).Or(builder.In("job_id", cloudbrainJobIDs))
	err := e.
		Where(cond).Unscoped().
		Find(&cloudbrainMaps)
	if err != nil {
		return nil, fmt.Errorf("find cloudbrain: %v", err)
	}

	cloudBrainJobIdMap := make(map[string]*Cloudbrain, len(cloudbrainIDs))
	for _, v := range cloudbrainMaps {
		cloudBrainJobIdMap[v.JobID] = v
	}

	for _, action := range actions {
		if !action.IsCloudbrainAction() {
			continue
		}
		cloudbrainId, _ := strconv.ParseInt(action.Content, 10, 64)
		if cloudbrainId > 0 {
			if c, ok := cloudbrainMaps[cloudbrainId]; ok {
				if c.DisplayJobName == action.RefName || c.JobName == action.RefName {
					action.Cloudbrain = c
					continue
				}

			}
		}
		if c, ok := cloudBrainJobIdMap[action.Content]; ok {
			if c.DisplayJobName == action.RefName || c.JobName == action.RefName {
				action.Cloudbrain = c
				continue
			}

		}
	}
	return valuesCloudbrain(cloudbrainMaps), nil
}

// LoadComments loads actions' all comments
func (actions ActionList) LoadCloudbrains() ([]*Comment, error) {
	return actions.loadComments(x)
}

// loadAttributes loads all attributes
func (actions ActionList) loadAttributes(e Engine) (err error) {
	if _, err = actions.loadUsers(e); err != nil {
		return
	}

	if _, err = actions.loadRepositories(e); err != nil {
		return
	}

	return nil
}

// LoadAttributes loads attributes of the actions
func (actions ActionList) LoadAttributes() error {
	return actions.loadAttributes(x)
}

// LoadAllAttributes loads all attributes of the actions
// compare with LoadAttributes() ,LoadAllAttributes() loads Comment and Cloudbrain attribute
func (actions ActionList) LoadAllAttributes() error {
	return actions.loadAllAttributes(x)
}

// loadAllAttributes
func (actions ActionList) loadAllAttributes(e Engine) (err error) {
	if _, err = actions.loadUsers(e); err != nil {
		return
	}

	if _, err = actions.loadRepositories(e); err != nil {
		return
	}

	if _, err = actions.loadComments(e); err != nil {
		return
	}

	if _, err = actions.loadCloudbrains(e); err != nil {
		return
	}
	if _, err = actions.loadDatasets(e); err != nil {
		return
	}
	if _, err = actions.loadAimodels(e); err != nil {
		return
	}

	return nil
}
