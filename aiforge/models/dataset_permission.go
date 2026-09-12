// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import "code.gitea.io/gitea/modules/log"

// GetUserDataSetPermission returns the user permissions to the data_set
func GetUserDataSetPermission(dataSet *Dataset, user *User) (isPermit bool, err error) {
	isPermit = false

	if user != nil {
		switch dataSet.Status {
		case DatasetStatusDeleted:
			log.Error("the data_set has been deleted")
		case DatasetStatusPrivate:
			if !user.IsAdmin && user.ID != dataSet.UserID {
				log.Error("the user is not admin nor the owner of the data_set")
			}
			isPermit = true

		case DatasetStatusPublic:
			isPermit = true
		default:
			log.Error("the status of data_set is wrong")
		}
	} else if !dataSet.IsPrivate() {
		isPermit = true
	}

	return isPermit, nil

}

func IsUserDatasetAdmin(dataset *DatasetRegistry, user *User) (bool, error) {
	return isUserDatasetAdmin(x, dataset, user)
}

func isUserDatasetAdmin(e Engine, dataset *DatasetRegistry, user *User) (bool, error) {
	if user == nil || dataset == nil {
		return false, nil
	}
	if user.IsAdmin {
		return true, nil
	}

	mode, err := subjectAccessLevel(e, user, dataset.ConvertSubjectAccessContext())
	if err != nil {
		return false, err
	}
	if mode >= AccessModeAdmin {
		return true, nil
	}

	teams, err := getUserSubjectTeams(e, dataset.OwnerID, user.ID, dataset.ID, DatasetSubject)
	if err != nil {
		return false, err
	}

	for _, team := range teams {
		if team.Authorize >= AccessModeAdmin {
			return true, nil
		}
	}
	return false, nil
}
