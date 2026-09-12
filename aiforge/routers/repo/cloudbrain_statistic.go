package repo

import (
	"net/http"
	"strings"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
	cloudbrainService "code.gitea.io/gitea/services/cloudbrain"
)

func CloudbrainDurationStatisticHour() {
	if setting.IsCloudbrainTimingEnabled {
		log.Info("CloudbrainDurationStatisticHour start")
		var statisticTime time.Time
		var count int64
		var countQueueNum int64
		recordDurationUpdateTime, err := models.GetDurationRecordUpdateTime()
		if err != nil {
			log.Error("Can not get GetDurationRecordBeginTime", err)
		}
		now := time.Now()
		currentTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
		if err == nil && len(recordDurationUpdateTime) > 0 {
			statisticTime = time.Unix(int64(recordDurationUpdateTime[0].DateTimeUnix), 0).Add(+1 * time.Hour)
		} else {
			statisticTime = currentTime
		}

		err = models.DeleteCloudbrainDurationStatistic(timeutil.TimeStamp(statisticTime.Unix()), timeutil.TimeStamp(currentTime.Unix()))
		if err != nil {
			log.Error("DeleteCloudbrainDurationStatistic failed", err)
		}

		err = models.DeleteQueueTaskDurationStatistic(timeutil.TimeStamp(statisticTime.Unix()), timeutil.TimeStamp(currentTime.Unix()))
		if err != nil {
			log.Error("DeleteQueueTaskDurationStatistic failed", err)
		}
		statisticTime = statisticTime.Add(+1 * time.Hour)
		for statisticTime.Before(currentTime) || statisticTime.Equal(currentTime) {
			countEach := summaryCloudbrainDurationStat(statisticTime)
			summaryCloudbrainTaskNumStat(statisticTime)
			countQueue := summaryQueueTaskDurationStat(statisticTime)
			summaryQueueTaskNumStat(statisticTime)
			count += countEach
			countQueueNum += countQueue
			statisticTime = statisticTime.Add(+1 * time.Hour)
		}
	}
}

func UpdateDurationStatisticHistoryData(beginTime time.Time, endTime time.Time) int64 {
	var count int64
	statisticTime := beginTime
	currentTime := endTime
	for statisticTime.Before(currentTime) || statisticTime.Equal(currentTime) {
		countEach := summaryCloudbrainDurationStat(statisticTime)
		count += countEach
		statisticTime = statisticTime.Add(+1 * time.Hour)
	}
	return count
}

func UpdateTaskNumStatisticHistoryData(beginTime time.Time, endTime time.Time) int64 {
	var count int64
	statisticTime := beginTime
	currentTime := endTime
	for statisticTime.Before(currentTime) || statisticTime.Equal(currentTime) {
		countEach := summaryCloudbrainTaskNumStat(statisticTime)
		count += countEach
		statisticTime = statisticTime.Add(+1 * time.Hour)
	}
	return count
}

func UpdateQueueTaskDurationStatisticHistoryData(beginTime time.Time, endTime time.Time) int64 {
	var count int64
	statisticTime := beginTime
	currentTime := endTime
	for statisticTime.Before(currentTime) || statisticTime.Equal(currentTime) {
		countEach := summaryQueueTaskDurationStat(statisticTime)
		count += countEach
		statisticTime = statisticTime.Add(+1 * time.Hour)
	}
	return count
}

func UpdateQueueTaskNumStatisticHistoryData(beginTime time.Time, endTime time.Time) int64 {
	var count int64
	statisticTime := beginTime
	currentTime := endTime
	for statisticTime.Before(currentTime) || statisticTime.Equal(currentTime) {
		countEach := summaryQueueTaskNumStat(statisticTime)
		count += countEach
		statisticTime = statisticTime.Add(+1 * time.Hour)
	}
	return count
}

// statisticTime是当前的时辰，比如当前是2019-01-01 12:01:01，那么statisticTime就是2019-01-01 12:00:00
func summaryCloudbrainDurationStat(statisticTime time.Time) int64 {
	var count int64
	dateTimeUnix := timeutil.TimeStamp(statisticTime.Add(-1 * time.Hour).Unix())
	beginTime := statisticTime.Add(-1 * time.Hour).Unix()
	dayTime := statisticTime.Add(-1 * time.Hour).Format("2006-01-02")
	hourTime := statisticTime.Add(-1 * time.Hour).Hour()
	endTime := statisticTime.Unix()

	ciTasks, err := models.GetCloudbrainByTime(beginTime, endTime)
	if err != nil {
		log.Error("GetCloudbrainByTime err: %v", err)
		return 0
	}
	err = models.LoadSpecs4CloudbrainInfo(ciTasks)
	if err != nil {
		log.Error("LoadSpecs4CloudbrainInfo err: %v", err)
	}
	cloudBrainCenterCodeAndCardTypeInfo, _ := getcloudBrainCenterCodeAndCardTypeInfo(ciTasks, int(beginTime), int(endTime))

	//需要记录总卡时
	resourceQueues, err := models.GetCanUseCardInfo()
	if err != nil {
		log.Error("GetCanUseCardInfo err: %v", err)
		return 0
	}

	cardsTotalDurationMap := make(map[string]int)
	for _, resourceQueue := range resourceQueues {
		key := resourceQueue.Cluster + "/" + resourceQueue.AiCenterCode + "/" + resourceQueue.ComputeResource + "/" + resourceQueue.AccCardType
		if _, ok := cardsTotalDurationMap[key]; !ok {
			cardsTotalDurationMap[key] = resourceQueue.CardsTotalNum * 1 * 60 * 60
		} else {
			cardsTotalDurationMap[key] += resourceQueue.CardsTotalNum * 1 * 60 * 60
		}
	}
	for key, cardsTotalDuration := range cardsTotalDurationMap {
		if _, ok := cloudBrainCenterCodeAndCardTypeInfo[strings.Split(key, "/")[0]+"/"+strings.Split(key, "/")[1]+"/"+strings.Split(key, "/")[2]][strings.Split(key, "/")[3]]; ok {
			cloudbrainDurationStat := models.CloudbrainDurationStatistic{
				DateTimeUnix:       dateTimeUnix,
				DayTime:            dayTime,
				HourTime:           hourTime,
				Cluster:            strings.Split(key, "/")[0],
				AiCenterName:       GetAiCenterNameByCode(strings.Split(key, "/")[1], "zh-CN"),
				AiCenterCode:       strings.Split(key, "/")[1],
				ComputeResource:    strings.Split(key, "/")[2],
				AccCardType:        strings.Split(key, "/")[3],
				CardsUseDuration:   cloudBrainCenterCodeAndCardTypeInfo[strings.Split(key, "/")[0]+"/"+strings.Split(key, "/")[1]+"/"+strings.Split(key, "/")[2]][strings.Split(key, "/")[3]],
				CardsTotalDuration: cardsTotalDuration,
				CardsTotalNum:      cardsTotalDuration / 1 / 60 / 60,
				CreatedUnix:        timeutil.TimeStampNow(),
			}
			if _, err = models.InsertCloudbrainDurationStatistic(&cloudbrainDurationStat); err != nil {
				log.Error("Insert cloudbrainDurationStat failed: %v", err.Error())
			}
			count++
		}
	}
	return count
}

// statisticTime是当前的时辰，比如当前是2019-01-01 12:01:01，那么statisticTime就是2019-01-01 12:00:00
func summaryCloudbrainTaskNumStat(statisticTime time.Time) int64 {
	var count int64
	dateTimeUnix := timeutil.TimeStamp(statisticTime.Add(-1 * time.Hour).Unix())
	beginTime := statisticTime.Add(-1 * time.Hour).Unix()
	dayTime := statisticTime.Add(-1 * time.Hour).Format("2006-01-02")
	hourTime := statisticTime.Add(-1 * time.Hour).Hour()
	endTime := statisticTime.Unix()

	computeCenterCiTasks, err := models.GetCloudbrainByCreateTime(beginTime, endTime)
	if err != nil {
		log.Error("GetCloudbrainByCreateTime err: %v", err)
		return 0
	}
	err = models.LoadSpecs4CloudbrainInfo(computeCenterCiTasks)
	if err != nil {
		log.Error("LoadSpecs4CloudbrainInfo err: %v", err)
	}
	ComputeCenterTaskNumMap, _ := getComputeCenterTaskNumMap(computeCenterCiTasks)
	for aicenter, ComputeCenterTaskNum := range ComputeCenterTaskNumMap {
		for cardType, TaskNum := range ComputeCenterTaskNum {
			CloudbrainTaskNumStat := models.CloudbrainTaskNumStatistic{
				DateTimeUnix: dateTimeUnix,
				DayTime:      dayTime,
				HourTime:     hourTime,
				Cluster:      strings.Split(aicenter, "/")[0],
				AiCenterName: GetAiCenterNameByCode(aicenter, "zh-CN"),
				AiCenterCode: aicenter,
				AccCardType:  cardType,
				CreatedUnix:  timeutil.TimeStampNow(),
				TaskNum:      TaskNum,
			}
			if _, err = models.InsertCloudbrainTaskNumStatistic(&CloudbrainTaskNumStat); err != nil {
				log.Error("Insert cloudbrainDurationStat failed: %v", err.Error())
			}
			count++
		}

	}
	return count
}

// statisticTime是当前的时辰，比如当前是2019-01-01 12:01:01，那么statisticTime就是2019-01-01 12:00:00
func summaryQueueTaskDurationStat(statisticTime time.Time) int64 {
	var count int64
	dateTimeUnix := timeutil.TimeStamp(statisticTime.Add(-1 * time.Hour).Unix())
	beginTime := statisticTime.Add(-1 * time.Hour).Unix()
	dayTime := statisticTime.Add(-1 * time.Hour).Format("2006-01-02")
	hourTime := statisticTime.Add(-1 * time.Hour).Hour()
	endTime := statisticTime.Unix()

	ciTasks, err := models.GetCloudbrainByTime(beginTime, endTime)
	if err != nil {
		log.Error("GetCloudbrainByTime err: %v", err)
		return 0
	}
	err = models.LoadSpecs4CloudbrainInfo(ciTasks)
	if err != nil {
		log.Error("LoadSpecs4CloudbrainInfo err: %v", err)
	}
	_, queueTaskCardDuration := getcloudBrainCenterCodeAndCardTypeInfo(ciTasks, int(beginTime), int(endTime))

	//需要记录总卡时
	resourceQueues, err := models.GetCanUseCardInfo()
	if err != nil {
		log.Error("GetCanUseCardInfo err: %v", err)
		return 0
	}

	cardsTotalDurationMap := make(map[int64]int)
	for _, resourceQueue := range resourceQueues {
		key := resourceQueue.ID
		if _, ok := cardsTotalDurationMap[key]; !ok {
			cardsTotalDurationMap[key] = resourceQueue.CardsTotalNum * 1 * 60 * 60
		} else {
			cardsTotalDurationMap[key] += resourceQueue.CardsTotalNum * 1 * 60 * 60
		}
	}
	for key, cardsTotalDuration := range cardsTotalDurationMap {
		queueId := key
		if _, ok := queueTaskCardDuration[queueId]; ok {
			queueTaskDurationStat := models.QueueTaskDurationStatistic{
				QueueID:            queueId,
				DateTimeUnix:       dateTimeUnix,
				DayTime:            dayTime,
				HourTime:           hourTime,
				CardsUseDuration:   queueTaskCardDuration[queueId],
				CardsTotalDuration: cardsTotalDuration,
				CardsTotalNum:      cardsTotalDuration / 1 / 60 / 60,
				CreatedUnix:        timeutil.TimeStampNow(),
			}
			if _, err = models.InsertQueueTaskDurationStatistic(&queueTaskDurationStat); err != nil {
				log.Error("Insert queueTaskDurationStat failed: %v", err.Error())
			}
			count++
		}
	}
	return count
}

// statisticTime是当前的时辰，比如当前是2019-01-01 12:01:01，那么statisticTime就是2019-01-01 12:00:00
func summaryQueueTaskNumStat(statisticTime time.Time) int64 {
	var count int64
	dateTimeUnix := timeutil.TimeStamp(statisticTime.Add(-1 * time.Hour).Unix())
	beginTime := statisticTime.Add(-1 * time.Hour).Unix()
	dayTime := statisticTime.Add(-1 * time.Hour).Format("2006-01-02")
	hourTime := statisticTime.Add(-1 * time.Hour).Hour()
	endTime := statisticTime.Unix()

	computeCenterCiTasks, err := models.GetCloudbrainByCreateTime(beginTime, endTime)
	if err != nil {
		log.Error("GetCloudbrainByCreateTime err: %v", err)
		return 0
	}
	err = models.LoadSpecs4CloudbrainInfo(computeCenterCiTasks)
	if err != nil {
		log.Error("LoadSpecs4CloudbrainInfo err: %v", err)
	}
	_, QueueTaskNumMap := getComputeCenterTaskNumMap(computeCenterCiTasks)
	for queueId, QueueTaskNum := range QueueTaskNumMap {
		queueTaskNumStat := models.QueueTaskNumStatistic{
			QueueID:      queueId,
			DateTimeUnix: dateTimeUnix,
			DayTime:      dayTime,
			HourTime:     hourTime,
			CreatedUnix:  timeutil.TimeStampNow(),
			TaskNum:      QueueTaskNum,
		}
		if _, err = models.InsertQueueTaskNumStatistic(&queueTaskNumStat); err != nil {
			log.Error("Insert InsertQueueTaskNumStatistic failed: %v", err.Error())
		}
		count++
	}
	return count
}

func GetAiCenterNameByCode(centerCode string, language string) string {
	ResourceAiCenterRes, err := models.GetResourceAiCenters()
	if err != nil {
		log.Error("Can not get ResourceAiCenterRes.", err)
		return centerCode
	}
	for _, v := range ResourceAiCenterRes {
		if v.AiCenterCode == centerCode {
			return v.AiCenterName
		}
	}
	return centerCode
}

func getcloudBrainCenterCodeAndCardTypeInfo(ciTasks []*models.CloudbrainInfo, hourBeginTime int, hourEndTime int) (map[string]map[string]int, map[int64]int) {
	var WorkServerNumber int
	var AccCardsNum int
	cloudBrainCenterCodeAndCardType := make(map[string]map[string]int)
	queueTaskCardDuration := make(map[int64]int)
	for _, cloudbrain := range ciTasks {
		cloudbrain = cloudbrainService.UpdateCloudbrainAiCenter(cloudbrain)
		cloudbrain = cloudbrainService.UpdateCloudbrainComputeResource(cloudbrain)
		if cloudbrain.Cloudbrain.StartTime == 0 {
			cloudbrain.Cloudbrain.StartTime = cloudbrain.Cloudbrain.CreatedUnix
		}
		if cloudbrain.Cloudbrain.EndTime == 0 {
			if cloudbrain.Cloudbrain.Status == string(models.JobRunning) {
				cloudbrain.Cloudbrain.EndTime = timeutil.TimeStamp(time.Now().Unix())
			} else {
				cloudbrain.Cloudbrain.EndTime = cloudbrain.Cloudbrain.StartTime + timeutil.TimeStamp(cloudbrain.Cloudbrain.Duration)
			}
		}
		if cloudbrain.Cloudbrain.WorkServerNumber >= 1 {
			WorkServerNumber = cloudbrain.Cloudbrain.WorkServerNumber
		} else {
			WorkServerNumber = 1
		}
		if cloudbrain.Cloudbrain.Spec == nil {
			AccCardsNum = 1
		} else {
			AccCardsNum = cloudbrain.Cloudbrain.Spec.AccCardsNum
		}

		var aiCenterCode string
		if cloudbrain.Cloudbrain.AiCenter != "" {
			aiCenterCode = cloudbrain.Cloudbrain.AiCenter
		} else {
			if cloudbrain.Cloudbrain.Spec != nil {
				aiCenterCode = cloudbrain.Cloudbrain.Spec.AiCenterCode
			} else {
				aiCenterCode = "other"
			}
		}

		key := cloudbrain.Cloudbrain.Cluster + "/" + aiCenterCode + "/" + cloudbrain.Cloudbrain.ComputeResource
		if _, ok := cloudBrainCenterCodeAndCardType[key]; !ok {
			cloudBrainCenterCodeAndCardType[key] = make(map[string]int)
		}
		taskStartTime := int(cloudbrain.Cloudbrain.StartTime)
		taskEndTime := int(cloudbrain.Cloudbrain.EndTime)
		if cloudbrain.Cloudbrain.Spec != nil {
			if _, ok := cloudBrainCenterCodeAndCardType[key][cloudbrain.Cloudbrain.Spec.AccCardType]; !ok {
				if taskStartTime < hourBeginTime && taskEndTime >= hourBeginTime && taskEndTime <= hourEndTime {
					cloudBrainCenterCodeAndCardType[key][cloudbrain.Cloudbrain.Spec.AccCardType] = AccCardsNum * WorkServerNumber * (taskEndTime - hourBeginTime)
				} else if taskStartTime < hourBeginTime && taskEndTime > hourEndTime {
					cloudBrainCenterCodeAndCardType[key][cloudbrain.Cloudbrain.Spec.AccCardType] = AccCardsNum * WorkServerNumber * (hourEndTime - hourBeginTime)
				} else if taskStartTime >= hourBeginTime && taskStartTime <= hourEndTime && taskEndTime >= hourBeginTime && taskEndTime <= hourEndTime {
					cloudBrainCenterCodeAndCardType[key][cloudbrain.Cloudbrain.Spec.AccCardType] = AccCardsNum * WorkServerNumber * (taskEndTime - taskStartTime)
				} else if taskStartTime >= hourBeginTime && taskStartTime <= hourEndTime && taskEndTime > hourEndTime {
					cloudBrainCenterCodeAndCardType[key][cloudbrain.Cloudbrain.Spec.AccCardType] = AccCardsNum * WorkServerNumber * (hourEndTime - taskStartTime)
				}
			} else {
				if taskStartTime < hourBeginTime && taskEndTime >= hourBeginTime && taskEndTime <= hourEndTime {
					cloudBrainCenterCodeAndCardType[key][cloudbrain.Cloudbrain.Spec.AccCardType] += AccCardsNum * WorkServerNumber * (taskEndTime - hourBeginTime)
				} else if taskStartTime < hourBeginTime && taskEndTime > hourEndTime {
					cloudBrainCenterCodeAndCardType[key][cloudbrain.Cloudbrain.Spec.AccCardType] += AccCardsNum * WorkServerNumber * (hourEndTime - hourBeginTime)
				} else if taskStartTime >= hourBeginTime && taskStartTime <= hourEndTime && taskEndTime >= hourBeginTime && taskEndTime <= hourEndTime {
					cloudBrainCenterCodeAndCardType[key][cloudbrain.Cloudbrain.Spec.AccCardType] += AccCardsNum * WorkServerNumber * (taskEndTime - taskStartTime)
				} else if taskStartTime >= hourBeginTime && taskStartTime <= hourEndTime && taskEndTime > hourEndTime {
					cloudBrainCenterCodeAndCardType[key][cloudbrain.Cloudbrain.Spec.AccCardType] += AccCardsNum * WorkServerNumber * (hourEndTime - taskStartTime)
				}
			}

			if _, ok := queueTaskCardDuration[cloudbrain.Cloudbrain.Spec.QueueId]; !ok {
				if taskStartTime < hourBeginTime && taskEndTime >= hourBeginTime && taskEndTime <= hourEndTime {
					queueTaskCardDuration[cloudbrain.Cloudbrain.Spec.QueueId] = AccCardsNum * WorkServerNumber * (taskEndTime - hourBeginTime)
				} else if taskStartTime < hourBeginTime && taskEndTime > hourEndTime {
					queueTaskCardDuration[cloudbrain.Cloudbrain.Spec.QueueId] = AccCardsNum * WorkServerNumber * (hourEndTime - hourBeginTime)
				} else if taskStartTime >= hourBeginTime && taskStartTime <= hourEndTime && taskEndTime >= hourBeginTime && taskEndTime <= hourEndTime {
					queueTaskCardDuration[cloudbrain.Cloudbrain.Spec.QueueId] = AccCardsNum * WorkServerNumber * (taskEndTime - taskStartTime)
				} else if taskStartTime >= hourBeginTime && taskStartTime <= hourEndTime && taskEndTime > hourEndTime {
					queueTaskCardDuration[cloudbrain.Cloudbrain.Spec.QueueId] = AccCardsNum * WorkServerNumber * (hourEndTime - taskStartTime)
				}
			} else {
				if taskStartTime < hourBeginTime && taskEndTime >= hourBeginTime && taskEndTime <= hourEndTime {
					queueTaskCardDuration[cloudbrain.Cloudbrain.Spec.QueueId] += AccCardsNum * WorkServerNumber * (taskEndTime - hourBeginTime)
				} else if taskStartTime < hourBeginTime && taskEndTime > hourEndTime {
					queueTaskCardDuration[cloudbrain.Cloudbrain.Spec.QueueId] += AccCardsNum * WorkServerNumber * (hourEndTime - hourBeginTime)
				} else if taskStartTime >= hourBeginTime && taskStartTime <= hourEndTime && taskEndTime >= hourBeginTime && taskEndTime <= hourEndTime {
					queueTaskCardDuration[cloudbrain.Cloudbrain.Spec.QueueId] += AccCardsNum * WorkServerNumber * (taskEndTime - taskStartTime)
				} else if taskStartTime >= hourBeginTime && taskStartTime <= hourEndTime && taskEndTime > hourEndTime {
					queueTaskCardDuration[cloudbrain.Cloudbrain.Spec.QueueId] += AccCardsNum * WorkServerNumber * (hourEndTime - taskStartTime)
				}
			}
		}
	}
	return cloudBrainCenterCodeAndCardType, queueTaskCardDuration
}

func getComputeCenterTaskNumMap(ciTasks []*models.CloudbrainInfo) (map[string]map[string]int, map[int64]int) {
	ComputeCenterTaskNumMap := make(map[string]map[string]int)
	QueueTaskNumMap := make(map[int64]int)
	var aiCenterCode string
	for _, cloudbrain := range ciTasks {
		cloudbrain = cloudbrainService.UpdateCloudbrainAiCenter(cloudbrain)
		if cloudbrain.Cloudbrain.AiCenter != "" {
			aiCenterCode = cloudbrain.Cloudbrain.AiCenter
			if _, ok := ComputeCenterTaskNumMap[aiCenterCode]; !ok {
				ComputeCenterTaskNumMap[aiCenterCode] = make(map[string]int)
			}
		} else {
			if cloudbrain.Cloudbrain.Spec != nil {
				aiCenterCode = cloudbrain.Cloudbrain.Spec.AiCenterCode
				ComputeCenterTaskNumMap[aiCenterCode] = make(map[string]int)
			} else {
				aiCenterCode = "other"
				if _, ok := ComputeCenterTaskNumMap[aiCenterCode]; !ok {
					ComputeCenterTaskNumMap[aiCenterCode] = make(map[string]int)
				}
			}
		}
		if cloudbrain.Cloudbrain.Spec != nil {
			if _, ok := ComputeCenterTaskNumMap[aiCenterCode][cloudbrain.Cloudbrain.Spec.AccCardType]; !ok {
				ComputeCenterTaskNumMap[aiCenterCode][cloudbrain.Cloudbrain.Spec.AccCardType] = 1
			} else {
				ComputeCenterTaskNumMap[aiCenterCode][cloudbrain.Cloudbrain.Spec.AccCardType] += 1
			}

			if _, ok := QueueTaskNumMap[cloudbrain.Cloudbrain.Spec.QueueId]; !ok {
				QueueTaskNumMap[cloudbrain.Cloudbrain.Spec.QueueId] = 1
			} else {
				QueueTaskNumMap[cloudbrain.Cloudbrain.Spec.QueueId] += 1
			}
		}
	}
	return ComputeCenterTaskNumMap, QueueTaskNumMap
}

func CloudbrainDurationUpdateHistoryData(ctx *context.Context) {
	beginTimeStr := ctx.QueryTrim("beginTime")
	endTimeStr := ctx.QueryTrim("endTime")
	var count int64
	var err error
	if beginTimeStr != "" && endTimeStr != "" {
		beginTime, _ := time.ParseInLocation("2006-01-02 15:04:05", beginTimeStr, time.Local)
		endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", endTimeStr, time.Local)
		if time.Now().Before(endTime) {
			endTime = time.Now()
		}
		beginTimeUnix := timeutil.TimeStamp(beginTime.Unix())
		endTimeUnix := timeutil.TimeStamp(endTime.Unix())

		err = models.DeleteCloudbrainDurationStatistic(beginTimeUnix, endTimeUnix)
		count = UpdateDurationStatisticHistoryData(beginTime.Add(+1*time.Hour), endTime.Add(+1*time.Hour))
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": 0,
		"count":   count,
		"err":     err,
	})
}

func CloudbrainTaskNumUpdateHistoryData(ctx *context.Context) {
	beginTimeStr := ctx.QueryTrim("beginTime")
	endTimeStr := ctx.QueryTrim("endTime")
	var count int64
	var err error
	if beginTimeStr != "" && endTimeStr != "" {
		beginTime, _ := time.ParseInLocation("2006-01-02 15:04:05", beginTimeStr, time.Local)
		endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", endTimeStr, time.Local)
		if time.Now().Before(endTime) {
			endTime = time.Now()
		}
		beginTimeUnix := timeutil.TimeStamp(beginTime.Unix())
		endTimeUnix := timeutil.TimeStamp(endTime.Unix())

		err = models.DeleteCloudbrainTaskNumStatistic(beginTimeUnix, endTimeUnix)
		count = UpdateTaskNumStatisticHistoryData(beginTime.Add(+1*time.Hour), endTime.Add(+1*time.Hour))
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": 0,
		"count":   count,
		"err":     err,
	})
}

func QueueTaskDurationUpdateHistoryData(ctx *context.Context) {
	beginTimeStr := ctx.QueryTrim("beginTime")
	endTimeStr := ctx.QueryTrim("endTime")
	var count int64
	var err error
	if beginTimeStr != "" && endTimeStr != "" {
		beginTime, _ := time.ParseInLocation("2006-01-02 15:04:05", beginTimeStr, time.Local)
		endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", endTimeStr, time.Local)
		if time.Now().Before(endTime) {
			endTime = time.Now()
		}
		beginTimeUnix := timeutil.TimeStamp(beginTime.Unix())
		endTimeUnix := timeutil.TimeStamp(endTime.Unix())

		err = models.DeleteQueueTaskDurationStatistic(beginTimeUnix, endTimeUnix)
		count = UpdateQueueTaskDurationStatisticHistoryData(beginTime.Add(+1*time.Hour), endTime.Add(+1*time.Hour))
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": 0,
		"count":   count,
		"err":     err,
	})
}

func QueueTaskNumUpdateHistoryData(ctx *context.Context) {
	beginTimeStr := ctx.QueryTrim("beginTime")
	endTimeStr := ctx.QueryTrim("endTime")
	var count int64
	var err error
	if beginTimeStr != "" && endTimeStr != "" {
		beginTime, _ := time.ParseInLocation("2006-01-02 15:04:05", beginTimeStr, time.Local)
		endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", endTimeStr, time.Local)
		if time.Now().Before(endTime) {
			endTime = time.Now()
		}
		beginTimeUnix := timeutil.TimeStamp(beginTime.Unix())
		endTimeUnix := timeutil.TimeStamp(endTime.Unix())

		err = models.DeleteQueueTaskNumStatistic(beginTimeUnix, endTimeUnix)
		count = UpdateQueueTaskNumStatisticHistoryData(beginTime.Add(+1*time.Hour), endTime.Add(+1*time.Hour))
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": 0,
		"count":   count,
		"err":     err,
	})
}
