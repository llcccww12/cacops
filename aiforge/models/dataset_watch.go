package models

import (
	"fmt"
)

func notifyDatasetWatchers(e Engine, actions ...*Action) error {
	var err error
	var dataset *DatasetRegistry
	var users []*User

	for _, act := range actions {
		if act.DatasetID == nil || *act.DatasetID == "" {
			continue
		}

		datasetChanged := dataset == nil || dataset.ID != *act.DatasetID

		if datasetChanged {
			users, err = getAllAccessUser(e, *act.DatasetID, AccessModeRead)
			if err != nil {
				return fmt.Errorf("get watchers: %v", err)
			}
		}
		// Add feed for actioner.
		act.UserID = act.ActUserID
		if _, err = e.InsertOne(act); err != nil {
			return fmt.Errorf("insert new actioner: %v", err)
		}
		// Send the act to task chan
		ActionChan4Task <- *act

		if (act.DatasetID == nil || *act.DatasetID == "") && act.Dataset == nil {
			return nil
		}

		if datasetChanged {
			act.loadDataset()
			dataset = act.Dataset

			if err := act.Dataset.getOwner(e); err != nil {
				return fmt.Errorf("can't get dataset owner: %v", err)
			}
		} else if act.Dataset == nil {
			act.Dataset = dataset
		}

		// Add feed for organization
		if act.Dataset.Owner.IsOrganization() && act.ActUserID != act.Dataset.Owner.ID {
			act.ID = 0
			act.UserID = act.Dataset.Owner.ID
			if _, err = e.InsertOne(act); err != nil {
				return fmt.Errorf("insert new actioner: %v", err)
			}
		}

		for _, user := range users {
			if act.ActUserID == user.ID {
				continue
			}
			act.ID = 0
			act.UserID = user.ID

			switch act.OpType {
			case ActionDatasetRecommended, ActionCreateDataset, ActionDeleteDataset:
				if _, err = e.InsertOne(act); err != nil {
					return fmt.Errorf("insert new action: %v", err)
				}
			}
		}

	}
	return nil
}
