package cloudbrain

import (
	"code.gitea.io/gitea/modules/util"

	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
)

var pageSize = 500

func CardStatistic() {
	sevenDayStatistic()
	thirtyDayStatistic()
	daysAllStatistic()

}

func daysAllStatistic() {

	daysStatistic(0, models.TypeAllDays)
}

func sevenDayStatistic() {

	daysStatistic(7, models.TypeSevenDays)

}

func thirtyDayStatistic() {

	daysStatistic(30, models.TypeThirtyDays)

}

func daysStatistic(days int, statisticType int) {
	xpuCardStatisticMap, err := getStatisticInfoMap(days, statisticType)
	if err != nil {
		return
	}

	updateOrInsertStatisticRecord(xpuCardStatisticMap)

}

func updateOrInsertStatisticRecord(xpuCardStatisticMap map[string]*models.XPUInfoStatistic) {
	for _, v := range xpuCardStatisticMap {
		v.UsedCardHour = v.UsedDuration / 3600
		err := models.UpdateOrInsertXPUInfoStatistic(v)
		if err != nil {
			log.Warn("update or insert xpu statistic err", err)
		}
	}
}

func getStatisticInfoMap(days int, statisticType int, beginTime ...int64) (map[string]*models.XPUInfoStatistic, error) {
	endTimeUnix := time.Now().Unix()
	beginTimeUnix := time.Now().AddDate(0, 0, -days).Unix()
	if days == 0 {
		beginTimeUnix = 1
	}
	if len(beginTime) > 0 {
		beginTimeUnix = beginTime[0]
	}

	_, count, err := models.CloudbrainAll(&models.CloudbrainsOptions{
		ListOptions: models.ListOptions{
			Page:     1,
			PageSize: 1,
		},
		NeedRepoInfo:  false,
		BeginTimeUnix: beginTimeUnix,
		EndTimeUnix:   endTimeUnix,
		Type:          -1,
		AccCardsNum:   -1,
	})
	if err != nil {
		log.Error("Get job failed:", err)
		return nil, err
	}

	xpuInfos, err := models.GetXPUInfos()
	if err != nil {
		log.Error("can not get XPU base info", err)
		return nil, err
	}

	var xpuCardStatisticMap = make(map[string]*models.XPUInfoStatistic, len(xpuInfos))
	var xpuCardStatisticUserMap = make(map[string]map[int64]struct{}, len(xpuInfos))
	for _, xpuInfo := range xpuInfos {
		xpuCardStatisticMap[xpuInfo.CardType] = &models.XPUInfoStatistic{
			InfoID:      xpuInfo.ID,
			Type:        statisticType,
			UpdatedUnix: endTimeUnix,
		}
		xpuCardStatisticUserMap[xpuInfo.CardType] = make(map[int64]struct{}, 0)
	}

	totalPage := util.GetTotalPage(count, pageSize)

	for i := 0; i < totalPage; i++ {

		ciTasks, _, err := models.CloudbrainAll(&models.CloudbrainsOptions{
			ListOptions: models.ListOptions{
				Page:     i + 1,
				PageSize: pageSize,
			},
			NeedRepoInfo:  false,
			BeginTimeUnix: beginTimeUnix,
			EndTimeUnix:   endTimeUnix,
			Type:          -1,
			AccCardsNum:   -1,
		})
		if err != nil {
			log.Error("Get job failed:", err)
			return nil, err
		}
		err = models.LoadSpecs4CloudbrainInfo(ciTasks)
		if err != nil {
			log.Error("can not load spec info", err)
			return nil, err
		}

		for _, task := range ciTasks {
			if task.Spec != nil {
				if v, ok := xpuCardStatisticMap[task.Spec.AccCardType]; ok {
					v.TaskCount += 1
					if calculateCardDuration(task.Cloudbrain) > 0 {
						v.UsedDuration += calculateCardDuration(task.Cloudbrain)
					}
					if _, found := xpuCardStatisticUserMap[task.Spec.AccCardType][task.UserID]; !found {

						v.UserCount += 1
						xpuCardStatisticUserMap[task.Spec.AccCardType][task.UserID] = struct{}{}
					}
				}

			}

		}

	}
	return xpuCardStatisticMap, nil
}

func calculateCardDuration(task models.Cloudbrain) int64 {

	cardNum := task.Spec.AccCardsNum

	var workServerNumber int64
	if task.WorkServerNumber >= 1 {
		workServerNumber = int64(task.WorkServerNumber)
	} else {
		workServerNumber = 1
	}
	return workServerNumber * int64(cardNum) * task.Duration
}
