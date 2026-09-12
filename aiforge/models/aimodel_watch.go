package models

import (
	"fmt"
)

func notifyAimodelWatchers(e Engine, actions ...*Action) error {
	var err error
	var aimodel *AiModelManage
	var users []*User

	for _, act := range actions {
		if act.AimodelID == nil || *act.AimodelID == "" {
			continue
		}

		aimodelChanged := aimodel == nil || aimodel.ID != *act.AimodelID

		if aimodelChanged {
			users, err = getAllAccessUser(e, *act.AimodelID, AccessModeRead)
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

		if (act.AimodelID == nil || *act.AimodelID == "") && act.Aimodel == nil {
			return nil
		}

		if aimodelChanged {
			act.loadAimodel()
			aimodel = act.Aimodel

			if err := act.Aimodel.getOwner(e); err != nil {
				return fmt.Errorf("can't get dataset owner: %v", err)
			}
		} else if act.Aimodel == nil {
			act.Aimodel = aimodel
		}

		// Add feed for organization
		if act.Aimodel.Owner.IsOrganization() && act.ActUserID != act.Aimodel.Owner.ID {
			act.ID = 0
			act.UserID = act.Aimodel.Owner.ID
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
			case ActionAimodelRecommended, ActionCreateAimodel, ActionDeleteAimodel:
				if _, err = e.InsertOne(act); err != nil {
					return fmt.Errorf("insert new action: %v", err)
				}
			}
		}

	}
	return nil
}
