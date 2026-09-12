package models

import (
	"fmt"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
)

const (
	TypeWenXin_Model int = iota
	TypeSD_Model
	TypeLLMT_Model
)

type ModelApp struct {
	ID                    string `xorm:"pk"`
	Desc                  string `xorm:"varchar(1000)"`
	Count                 int
	UserId                int64 `xorm:"INDEX"`
	ExternalID            string
	Picture               string `xorm:"text NULL"` //picture base64
	TaskId                int
	Height                int
	Width                 int
	Seed                  int
	Steps                 int
	Num_images_per_prompt int
	Negative_prompt       string
	Scheduler_name        string
	Guidance_scale        int
	Status                int
	Type                  int
	Url                   string
	CreatedUnix           timeutil.TimeStamp `xorm:"created"`
	UpdatedUnix           timeutil.TimeStamp `xorm:"updated"`
}

func SaveModelApp(modelApp *ModelApp) error {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()
	re, err := statictisSess.Insert(modelApp)
	if err != nil {
		log.Info("insert modelApp error." + err.Error())
		return err
	}
	log.Info("success to save modelApp db.re=" + fmt.Sprint((re)))
	return nil
}

func UpdateModelApp(modelApp *ModelApp) error {
	statictisSess := xStatistic.ID(modelApp.ID)
	defer statictisSess.Close()
	re, err := statictisSess.Cols("picture", "status").Update(modelApp)
	if err != nil {
		return err
	}
	log.Info("update modelApp db.re=" + fmt.Sprint((re)))
	return nil
}

func QueryModelAppById(id string) *ModelApp {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()
	re := new(ModelApp)
	isExist, err := statictisSess.Table(new(ModelApp)).ID(id).Get(re)
	if err == nil && isExist {
		return re
	}
	return nil
}

func QueryModelAppCount(userId int64) int64 {
	statictisSess := xStatistic.NewSession()
	defer statictisSess.Close()
	sumList, err := statictisSess.QueryInterface("select sum(count) as count from public.model_app where user_id=" + fmt.Sprint(userId))
	if err == nil {
		if len(sumList) == 1 {
			return convertInterfaceToInt64(sumList[0]["count"])
		}
	}
	return 0
}
