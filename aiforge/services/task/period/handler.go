package period

import (
	"code.gitea.io/gitea/models"
	"errors"
	"time"
)

var PeriodHandlerMap = map[string]PeriodHandler{
	models.PeriodNotCycle: new(NoCycleHandler),
	models.PeriodDaily:    new(DailyHandler),
}

type PeriodHandler interface {
	GetCurrentPeriod() *models.PeriodResult
}

type NoCycleHandler struct {
}

func (l *NoCycleHandler) GetCurrentPeriod() *models.PeriodResult {
	return nil
}

type DailyHandler struct {
}

func (l *DailyHandler) GetCurrentPeriod() *models.PeriodResult {
	t := time.Now()
	startTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	endTime := startTime.Add(24 * time.Hour)
	leftTime := endTime.Sub(t)
	return &models.PeriodResult{
		StartTime: startTime,
		EndTime:   endTime,
		LeftTime:  leftTime,
	}
}

func getPeriodHandler(refreshRateype string) PeriodHandler {
	return PeriodHandlerMap[refreshRateype]
}

func GetPeriod(refreshRate string) (*models.PeriodResult, error) {
	handler := getPeriodHandler(refreshRate)
	if handler == nil {
		return nil, errors.New("task config incorrect")
	}
	return handler.GetCurrentPeriod(), nil
}
