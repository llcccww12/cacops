package models

import (
	"code.gitea.io/gitea/modules/log"
	"xorm.io/xorm"
)

type CustomMigration struct {
	Description string
	Migrate     func(*xorm.Engine) error
}

type CustomMigrationStatic struct {
	Description string
	Migrate     func(*xorm.Engine, *xorm.Engine) error
}

var customMigrations = []CustomMigration{
	//手机号功能可以不启用，不启用时手机号为空串，为了避免启用时唯一性约束报错，定制唯一性约束（对null和空字符串不做唯一性检查）
	{"set phone number  unique index", setPhoneNumberUniqueIndex},
}

var customMigrationsStatic []CustomMigrationStatic



func MigrateCustom(x *xorm.Engine) {

	for _, m := range customMigrations {
		log.Info("Migration: %s", m.Description)
		if err := m.Migrate(x); err != nil {

			log.Error("Migration: %v", err)

		}
	}
}

func MigrateCustomStatic(x *xorm.Engine, static *xorm.Engine) {
	for _, m := range customMigrationsStatic {
		log.Info("Migration: %s", m.Description)
		if err := m.Migrate(x, static); err != nil {
			log.Error("Migration: %v", err)
		}
	}
}

func setPhoneNumberUniqueIndex(x *xorm.Engine) error {

	query := "CREATE UNIQUE INDEX IF NOT EXISTS \"UQE_user_phone_number\" ON \"user\"  (phone_number) WHERE phone_number IS NOT NULL and phone_number !='';"

	_, err := x.Exec(query)
	return err
}


func syncTopicStruct(x *xorm.Engine) error {

	query := "ALTER TABLE topic ALTER COLUMN name TYPE varchar(105);"

	_, err := x.Exec(query)
	return err
}

func updateIssueFixedRate(x *xorm.Engine, static *xorm.Engine) error {
	updateSQL := "update  repo_statistic  set issue_fixed_rate=1.0 where num_issues=0"
	_, err := static.Exec(updateSQL)
	return err
}
