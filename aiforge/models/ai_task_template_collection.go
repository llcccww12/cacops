package models

import (
	"code.gitea.io/gitea/modules/timeutil"
	"xorm.io/builder"
)

type AITaskTemplateCollection struct {
	ID          int64              `xorm:"pk autoincr"`
	UserID      int64              `xorm:"UNIQUE(s)"`
	TemplateID  string             `xorm:"uuid UNIQUE(s)"`
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
}

func CollectionAITaskTemplate(userID int64, templateID string, collect bool) error {
	sess := x.NewSession()
	defer sess.Close()

	if err := sess.Begin(); err != nil {
		return err
	}

	var err error
	defer func() {
		if err != nil {
			sess.Rollback()
		} else {
			sess.Commit()
		}
	}()

	if collect {
		result, err2 := sess.Exec(
			"INSERT INTO ai_task_template_collection (user_id, template_id) VALUES (?, ?) ON CONFLICT (user_id, template_id) DO NOTHING",
			userID, templateID,
		)
		if err = err2; err != nil {
			return err
		}

		if affected, _ := result.RowsAffected(); affected > 0 {
			_, err = sess.Exec("UPDATE ai_task_template SET num_collections = num_collections + 1 WHERE id = ?", templateID)
		}
	} else {
		result, err2 := sess.Exec(
			"DELETE FROM ai_task_template_collection WHERE user_id = ? AND template_id = ?",
			userID, templateID,
		)
		if err = err2; err != nil {
			return err
		}

		if affected, _ := result.RowsAffected(); affected > 0 {
			_, err = sess.Exec("UPDATE ai_task_template SET num_collections = num_collections - 1 WHERE id = ? AND num_collections > 0", templateID)
		}
	}

	return err
}

func isAITaskTemplateCollecting(e Engine, userID int64, templateID string) bool {
	has, _ := e.Get(&AITaskTemplateCollection{0, userID, templateID, 0})
	return has
}

func IsAITaskTemplateCollecting(userID int64, templateID string) bool {
	return isAITaskTemplateCollecting(x, userID, templateID)
}

func GetAITaskTemplateCollectionsMap(userId int64, templateIds []string) (map[string]bool, error) {
	if len(templateIds) == 0 {
		return nil, nil
	}
	var collections []AITaskTemplateCollection
	err := x.Where(
		builder.NewCond().
			And(builder.Eq{"user_id": userId}).
			And(builder.In("template_id", templateIds))).
		Find(&collections)
	if err != nil {
		return nil, err
	}
	collectionMap := make(map[string]bool)
	for _, templateId := range templateIds {
		for _, collection := range collections {
			if collection.TemplateID == templateId {
				collectionMap[templateId] = true
				break
			}
		}
		if _, exists := collectionMap[templateId]; !exists {
			collectionMap[templateId] = false
		}
	}
	return collectionMap, nil
}
