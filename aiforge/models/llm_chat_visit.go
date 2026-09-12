package models

import (
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
	"fmt"
)

type LlmChatVisit struct {
	ID          int64 `xorm:"pk autoincr"`
	UserId      int64 `xorm:"INDEX"`
	ChatId      string
	ModelName   string
	Agreement   int
	ExpiredTime string
	ExpiredUnix int64
	CreatedUnix timeutil.TimeStamp `xorm:"created"`
}

func SaveVisit(llmChatVisit *LlmChatVisit) error {
	sess := xStatistic.NewSession()
	defer sess.Close()
	re, err := sess.Insert(llmChatVisit)
	if err != nil {
		log.Info("insert llmChatVisit error %s\n", err.Error())
		return err
	}
	log.Info("success to save llmChatVisit db.re=%+v\n", fmt.Sprint(re))
	return nil
}

func QueryFirstVisit(userId int64) int64 {
	sess := xStatistic.NewSession()
	defer sess.Close()
	query := "SELECT SUM(agreement) AS count FROM public.llm_chat_visit WHERE user_id = ?"
	sumList, err := sess.QueryInterface(query, userId)
	if err == nil {
		if len(sumList) == 1 {
			val := convertInterfaceToInt64(sumList[0]["count"])
			return val
		}
	}
	return 0
}

func QueryRunningChat(userId int64, modelName string, currentTime int64) (*LlmChatVisit, error) {
	sess := xStatistic.NewSession()
	defer sess.Close()
	re := new(LlmChatVisit)
	isExist, err := sess.Table(new(LlmChatVisit)).Where("user_id = ? AND model_name = ? AND ? > created_unix AND ? < expired_unix", userId, modelName, currentTime, currentTime).Get(re)
	if err == nil && isExist {
		return re, nil
	}
	return nil, err
}

func QueryByChatId(chatId string) (*LlmChatVisit, error) {
	sess := xStatistic.NewSession()
	defer sess.Close()
	re := new(LlmChatVisit)
	isExist, err := sess.Table(new(LlmChatVisit)).Where("chat_id = ?", chatId).Get(re)
	if err == nil && isExist {
		return re, nil
	}
	return nil, err
}

func UpdateChat(llmChatVisit *LlmChatVisit) error {
	sess := xStatistic.ID(llmChatVisit.ID)
	defer sess.Close()
	re, err := sess.Cols("agreement").Update(llmChatVisit)
	if err != nil {
		return err
	}
	log.Info("update llmChatVisit db.re=" + fmt.Sprint(re))
	return nil
}

func QueryChatVisitStatistics() ([]map[string]interface{}, error) {
	sess := xStatistic.NewSession()
	query := `
		SELECT
			COALESCE(model_name, 'total') as model_name,
			COUNT(DISTINCT chat_id) AS visit,
			COUNT(DISTINCT user_id) AS visit_user
		FROM
			llm_chat_visit
		GROUP BY
			ROLLUP(model_name)`

	results, err := sess.SQL(query).QueryInterface()
	if err != nil {
		return nil, err
	}

	return results, nil
}
