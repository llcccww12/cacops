package repo

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"code.gitea.io/gitea/routers/response"

	"code.gitea.io/gitea/modules/util"

	"code.gitea.io/gitea/modules/setting"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/routers/repo"
	cloudbrainService "code.gitea.io/gitea/services/cloudbrain"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type DateCloudbrainNum struct {
	Date                      string         `json:"date"`
	CloudOneJobTypeRes        map[string]int `json:"cloudOneJobTypeRes"`
	CloudTwoJobTypeRes        map[string]int `json:"cloudTwoJobTypeRes"`
	IntelligentNetJobTypeRes  map[string]int `json:"intelligentNetJobTypeRes"`
	CDCenterJobTypeRes        map[string]int `json:"cDCenterJobTypeRes"`
	CloudBrainPeriodNum       map[int]int    `json:"cloudBrainPeriodNum"`
	CloudBrainComputeResource map[string]int `json:"cloudBrainComputeResource"`
}
type DateCloudbrainInfo struct {
	Date                      string         `json:"date"`
	CloudBrainPeriodNum       map[int]int    `json:"cloudBrainPeriodNum"`
	CloudBrainComputeResource map[string]int `json:"cloudBrainComputeResource"`
}
type CloudbrainsStatusAnalysis struct {
	JobWaitingCount   int64 `json:"jobWaitingCount"`
	JobRunningCount   int64 `json:"jobRunningCount"`
	JobStoppedCount   int64 `json:"jobStoppedCount"`
	JobCompletedCount int64 `json:"jobCompletedCount"`
	JobFailedCount    int64 `json:"jobFailedCount"`
	JobKilledCount    int64 `json:"jobKilledCount"`
	JobInitCount      int64 `json:"jobInitCount"`
}

func GetAllCloudbrainsOverview(ctx *context.Context) {
	recordCloudbrain, err := models.GetRecordBeginTime()
	if err != nil {
		log.Error("Can not get recordCloudbrain", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.record_begintime_get_err"))
		return
	}
	recordBeginTime := recordCloudbrain[0].Cloudbrain.CreatedUnix
	now := time.Now()
	beginTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endTime := now
	todayCreatorCount, err := models.GetTodayCreatorCount(beginTime, endTime)
	if err != nil {
		log.Error("Can not query todayCreatorCount.", err)
		return
	}
	cloudbrainTypeCount, err := models.GetCloudbrainTypeCount()
	if err != nil {
		log.Error("Can not query cloudbrainTypeCount.", err)
		return
	}

	todayCloudbrainCount, err := models.GetTodayCloudbrainCount(beginTime, endTime)
	if err != nil {
		log.Error("Can not query todayCloudbrainCount.", err)
		return
	}

	todayRunningCount, err := models.GetTodayRunningCount(beginTime, endTime)
	if err != nil {
		log.Error("Can not query todayRunningCount.", err)
		return
	}

	todayWaitingCount, err := models.GetTodayWaitingCount(beginTime, endTime)
	if err != nil {
		log.Error("Can not query todayWaitingCount.", err)
		return
	}

	todayCompletedCount := todayCloudbrainCount - todayRunningCount - todayWaitingCount

	creatorCount, err := models.GetCreatorCount()
	if err != nil {
		log.Error("Can not query creatorCount.", err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"recordBeginTime":     recordBeginTime,
		"updateTime":          now.Unix(),
		"todayCreatorCount":   todayCreatorCount,
		"creatorCount":        creatorCount,
		"todayRunningCount":   todayRunningCount,
		"todayCompletedCount": todayCompletedCount,
		"todayWaitingCount":   todayWaitingCount,
		"todayNewJobCount":    todayCloudbrainCount,
		"cloudbrainTypeCount": cloudbrainTypeCount,
	})
}
func GetOverviewDuration(ctx *context.Context) {
	durationSum := 0
	cardDurationSum := 0

	cloudBrainOneCardDuSum := 0
	cloudBrainTwoCardDuSum := 0
	c2NetCardDuSum := 0
	cDNetCardDuSum := 0

	cloudBrainOneDuration := 0
	cloudBrainTwoDuration := 0
	c2NetDuration := 0
	cDCenterDuration := 0

	cloudbrainTypeDuration, err := models.GetCloudbrainTypeCardDuration()
	if err != nil {
		log.Error("GetCloudbrainTypeCardDuration err!", err)
		return
	}
	for _, result := range cloudbrainTypeDuration {
		if result.Type == models.TypeCloudBrainOne {
			cloudBrainOneDuration = result.DurationSum
			cloudBrainOneCardDuSum = result.CardDurationSum
		}
		if result.Type == models.TypeCloudBrainTwo {
			cloudBrainTwoDuration = result.DurationSum
			cloudBrainTwoCardDuSum = result.CardDurationSum
		}
		if result.Type == models.TypeC2Net {
			c2NetDuration = result.DurationSum
			c2NetCardDuSum = result.CardDurationSum
		}
		if result.Type == models.TypeCDCenter {
			cDCenterDuration = result.DurationSum
			cDNetCardDuSum = result.CardDurationSum
		}
	}
	cloudbrainAllDuration, err := models.GetCloudbrainAllCardDuration()
	if err != nil {
		log.Error("GetCloudbrainAllCardDuration err!", err)
		return
	}
	durationSum = cloudbrainAllDuration.DurationSum
	cardDurationSum = cloudbrainAllDuration.CardDurationSum

	openiCardDuSum := cloudBrainOneCardDuSum + cloudBrainTwoCardDuSum + cDNetCardDuSum
	openiDuration := cloudBrainOneDuration + cloudBrainTwoDuration + cDCenterDuration
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"cloudBrainOneCardDuSum": cloudBrainOneCardDuSum,
		"cloudBrainTwoCardDuSum": cloudBrainTwoCardDuSum,
		"cDNetCardDuSum":         cDNetCardDuSum,
		"c2NetCardDuSum":         c2NetCardDuSum,
		"openiCardDuSum":         openiCardDuSum,
		"cardDuSum":              cardDurationSum,

		"cloudBrainOneDuration": cloudBrainOneDuration,
		"cloudBrainTwoDuration": cloudBrainTwoDuration,
		"cDCenterDuration":      cDCenterDuration,
		"c2NetDuration":         c2NetDuration,
		"openiDuration":         openiDuration,
		"durationSum":           durationSum,
	})
}

func GetCloubrainOverviewGroupByAiCenter(ctx *context.Context) {

	cloudbrainCardTimeAndCountArray, err := models.GetCloudbrainCardTimeAndCountGroupByAICenter()
	if err != nil {
		log.Error("Can not query CardTimeAndCount.", err)
	}

	cardTimeMap, maxCardTime, _ := getCenterCardTimeInfo(cloudbrainCardTimeAndCountArray)

	var aiCenterLocationInfos = make(map[string][]*cloudbrainService.AiCenterLocationInfo, 0)

	const AI_CENTER = "智算中心"
	for _, value := range setting.AiCenterCodeAndNameAndLocMapInfo {
		long, lat := getLongLat(value.Loc)
		aicenterArray, ok := aiCenterLocationInfos[value.Type]
		if !ok {
			aicenterArray = make([]*cloudbrainService.AiCenterLocationInfo, 0)

		}
		if value.Type == "超算中心" || value.Type == "东数西算" {

			aiCenterLocationInfos[value.Type] = append(aicenterArray, &cloudbrainService.AiCenterLocationInfo{
				Name:      cloudbrainService.GetAiCenterShowByAiCenterId(value.Name, ctx),
				Longitude: long,
				Latitude:  lat,
				Value:     setting.ScreenMap.MinValue,
			})
		} else if value.Type == AI_CENTER {

			aiCenterLocationInfos[value.Type] = append(aicenterArray, &cloudbrainService.AiCenterLocationInfo{
				Name:      cloudbrainService.GetAiCenterShowByAiCenterId(value.Name, ctx),
				Longitude: long,
				Latitude:  lat,
				Value:     getAiCenterSize(value.Name, cardTimeMap, maxCardTime, 0),
			})

		}

	}
	sort.SliceStable(aiCenterLocationInfos[AI_CENTER], func(i, j int) bool {
		return aiCenterLocationInfos[AI_CENTER][i].Value > aiCenterLocationInfos[AI_CENTER][j].Value
	})

	if setting.ScreenMap.ShowData || ctx.IsUserSiteAdmin() {

		var cloudbrainCardTimeAndCountFilterArray = make([]map[string]string, 0)

		for _, cloudbrainCardTimeAndCountMap := range cloudbrainCardTimeAndCountArray {
			centerId := cloudbrainCardTimeAndCountMap["ai_center"]
			centerShow := cloudbrainService.GetAiCenterShowByAiCenterId(centerId, ctx)
			cloudbrainCardTimeAndCountMap["ai_center"] = centerShow

			if len(setting.ScreenMap.ExcludeCenter) > 0 && isExcludeCenter(centerId) {
				continue

			}
			cloudbrainCardTimeAndCountFilterArray = append(cloudbrainCardTimeAndCountFilterArray, cloudbrainCardTimeAndCountMap)
		}
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"cardAndJobCount": cloudbrainCardTimeAndCountFilterArray,
			"locationInfo":    aiCenterLocationInfos,
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"cardAndJobCount": []map[string]string{},
		"locationInfo":    aiCenterLocationInfos,
	})
	return

}

func isExcludeCenter(id string) bool {
	for _, centerId := range setting.ScreenMap.ExcludeCenter {
		if id == centerId {
			return true
		}
	}
	return false
}

func getAiCenterSize(name string, timeMap map[string]int64, MaxCardTime int64, MinCardTime int64) int {
	cardTime, _ := timeMap[name]
	if cardTime == 0 {
		return setting.ScreenMap.MinValue
	} else {
		if MaxCardTime == MinCardTime {
			return setting.ScreenMap.MaxValue
		} else {
			return int(float64(cardTime-MinCardTime)/float64(MaxCardTime-MinCardTime)*float64(setting.ScreenMap.MaxValue-setting.ScreenMap.MinValue)) + setting.ScreenMap.MinValue
		}
	}

}

func getLongLat(loc string) (string, string) {
	longLat := strings.Split(loc, ",")
	if len(longLat) != 2 {
		return "", ""
	}
	return longLat[0], longLat[1]
}

func getCenterCardTimeInfo(cloudbrainCardTimeAndCountArray []map[string]string) (map[string]int64, int64, int64) {
	var centerCardTimeMap = make(map[string]int64, len(cloudbrainCardTimeAndCountArray))
	var maxCardTime int64 = 0
	var minCardTime int64 = 0
	for i, cloudbrainCardTimeAndCount := range cloudbrainCardTimeAndCountArray {

		cardTime, _ := strconv.ParseInt(cloudbrainCardTimeAndCount["card_duration"], 10, 64)
		if i == 0 {
			maxCardTime = cardTime
		}
		if i == len(cloudbrainCardTimeAndCountArray)-1 {
			minCardTime = cardTime
		}
		centerCardTimeMap[cloudbrainCardTimeAndCount["ai_center"]] = cardTime
	}
	return centerCardTimeMap, maxCardTime, minCardTime
}

func GetAllCloudbrainsTrend(ctx *context.Context) {

	queryType := ctx.QueryTrim("type")
	now := time.Now()

	beginTimeStr := ctx.QueryTrim("beginTime")
	endTimeStr := ctx.QueryTrim("endTime")
	var beginTime time.Time
	var endTime time.Time
	dateCloudbrainNum := make([]DateCloudbrainNum, 0)
	var err error
	if queryType != "" {
		var targetDate string
		if queryType == "all" {
			if setting.CloudbrainNumTargetTime == "" {
				targetDate = "2025/08"
			} else {
				targetDate = setting.CloudbrainNumTargetTime
			}
			//判断表里是否有targetDate之前的记录，若无，则执行下面开始入库，若有，则targetDate开始，之后的照常执行
			if models.CheckDateExistsInCloudbrainNumJson(targetDate) > 0 {
				// 获取targetDate之前所有数据
				target, err := time.Parse("2006/01", targetDate)
				if err != nil {
					log.Error("解析目标日期失败", err)
					ctx.Error(http.StatusBadRequest, ctx.Tr("get_targetDate_error"))
					return
				}

				BeforeDateCloudbrainNum, _ := GetAllCloudbrainNums()

				// 先加一个月，然后将日期设置为1号，时间设置为0点
				beginTime := time.Date(
					target.Year(),
					target.Month()+1, // 下个月
					1,                // 第一天
					0, 0, 0, 0,       // 0时0分0秒
					time.UTC, // 时区
				)

				endTime = now
				dateCloudbrainNum, err = getMonthCloudbrainNum(beginTime, endTime)
				if err != nil {
					log.Error("Can not query getMonthCloudbrainNum.", err)
					ctx.Error(http.StatusBadRequest, ctx.Tr("getMonthCloudbrainNum_get_error"))
					return
				}
				dateCloudbrainNum = MergeByAppend(BeforeDateCloudbrainNum, dateCloudbrainNum)
			} else {
				recordCloudbrain, err := models.GetRecordBeginTime()
				if err != nil {
					log.Error("Can not get recordCloudbrain", err)
					ctx.Error(http.StatusBadRequest, ctx.Tr("repo.record_begintime_get_err"))
					return
				}
				brainRecordBeginTime := recordCloudbrain[0].Cloudbrain.CreatedUnix.AsTime()
				beginTime = brainRecordBeginTime
				endTime = now
				dateCloudbrainNum, err = getMonthCloudbrainNum(beginTime, endTime)
				if err != nil {
					log.Error("Can not query getMonthCloudbrainNum.", err)
					ctx.Error(http.StatusBadRequest, ctx.Tr("getMonthCloudbrainNum_get_error"))
					return
				}
			}
		} else if queryType == "today" {
			beginTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
			// endTimeTemp = now
			endTime = beginTime.AddDate(0, 0, 1)
			dateCloudbrainNum, err = getDayCloudbrainNum(beginTime, endTime)
			if err != nil {
				log.Error("Can not query getDayCloudbrainNum.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("getDayCloudbrainNum_get_error"))
				return
			}
		} else if queryType == "yesterday" {
			beginTime = now.AddDate(0, 0, -1)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
			dateCloudbrainNum, err = getDayCloudbrainNum(beginTime, endTime)
			if err != nil {
				log.Error("Can not query getDayCloudbrainNum.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("getDayCloudbrainNum_get_error"))
				return
			}
		} else if queryType == "last_7day" {
			beginTime = now.AddDate(0, 0, -7)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())

			endTime = now.AddDate(0, 0, 0)
			endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, now.Location())
			dateCloudbrainNum, err = getDayCloudbrainNum(beginTime, endTime)
			if err != nil {
				log.Error("Can not query getDayCloudbrainNum.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("getDayCloudbrainNum_get_error"))
				return
			}
		} else if queryType == "last_30day" {
			beginTime = now.AddDate(0, 0, -30)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())

			endTime = now.AddDate(0, 0, 0)
			endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, now.Location())
			dateCloudbrainNum, err = getDayCloudbrainNum(beginTime, endTime)
			if err != nil {
				log.Error("Can not query getDayCloudbrainNum.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("getDayCloudbrainNum_get_error"))
				return
			}
		} else if queryType == "current_month" {
			endTime = now
			beginTime = time.Date(endTime.Year(), endTime.Month(), 1, 0, 0, 0, 0, now.Location())
			dateCloudbrainNum, err = getDayCloudbrainNum(beginTime, endTime)
			if err != nil {
				log.Error("Can not query getDayCloudbrainNum.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("getDayCloudbrainNum_get_error"))
				return
			}
		} else if queryType == "current_year" {
			endTime = now
			beginTime = time.Date(endTime.Year(), 1, 1, 0, 0, 0, 0, now.Location())
			// endTimeTemp = beginTime.AddDate(0, 1, 0)
			dateCloudbrainNum, err = getMonthCloudbrainNum(beginTime, endTime)
			if err != nil {
				log.Error("Can not query getMonthCloudbrainNum.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("getMonthCloudbrainNum_get_error"))
				return
			}
		} else if queryType == "last_month" {

			lastMonthTime := now.AddDate(0, -1, 0)
			beginTime = time.Date(lastMonthTime.Year(), lastMonthTime.Month(), 1, 0, 0, 0, 0, now.Location())
			endTime = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
			dateCloudbrainNum, err = getDayCloudbrainNum(beginTime, endTime)
			if err != nil {
				log.Error("Can not query getDayCloudbrainNum.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("getDayCloudbrainNum_get_error"))
				return
			}
		}

	} else {
		if beginTimeStr == "" || endTimeStr == "" {
			//如果查询类型和开始时间结束时间都未设置，按queryType=all处理
			recordCloudbrain, err := models.GetRecordBeginTime()
			if err != nil {
				log.Error("Can not get recordCloudbrain", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("repo.record_begintime_get_err"))
				return
			}
			brainRecordBeginTime := recordCloudbrain[0].Cloudbrain.CreatedUnix.AsTime()
			beginTime = brainRecordBeginTime
			endTime = now
			dateCloudbrainNum, err = getMonthCloudbrainNum(beginTime, endTime)
			if err != nil {
				log.Error("Can not query getMonthCloudbrainNum.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("getMonthCloudbrainNum_get_error"))
				return
			}
		} else {
			beginTime, err = time.ParseInLocation("2006-01-02", beginTimeStr, time.Local)
			if err != nil {
				log.Error("Can not ParseInLocation.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("ParseInLocation_get_error"))
				return
			}
			endTime, err = time.ParseInLocation("2006-01-02", endTimeStr, time.Local)
			if err != nil {
				log.Error("Can not ParseInLocation.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("ParseInLocation_get_error"))
				return
			}
			if endTime.After(time.Now()) {
				endTime = time.Now()
			}
			dateCloudbrainNum, err = getDayCloudbrainNum(beginTime, endTime)
			if err != nil {
				log.Error("Can not query getDayCloudbrainNum.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("getDayCloudbrainNum_get_error"))
				return
			}
		}

	}
	totalCloudbrainNum := SumDateCloudbrainNums(dateCloudbrainNum)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"timeCloudbrainNum":  dateCloudbrainNum,
		"totalCloudbrainNum": totalCloudbrainNum,
	})

}

func MergeByAppend(a, b []DateCloudbrainNum) []DateCloudbrainNum {
	return append(a, b...)
}

// GetAllCloudbrainNums 从数据库获取所有记录并解析为DateCloudbrainNum切片
func GetAllCloudbrainNums() ([]DateCloudbrainNum, error) {
	// 1. 从数据库获取所有原始记录
	rawRecords, err := models.GetAllCloudbrainNumJson()
	if err != nil {
		return nil, fmt.Errorf("获取数据库记录失败: %v", err)
	}

	var results []DateCloudbrainNum

	for _, record := range rawRecords {
		var cloudbrainNum DateCloudbrainNum

		err := json.Unmarshal([]byte(record.CloudbrainNumJson), &cloudbrainNum)
		if err != nil {
			log.Error("解析JSON失败，日期: %s, 错误: %v", record.Date, err)
			continue // 跳过解析失败的记录，或根据需求返回错误
		}
		if cloudbrainNum.Date == "" {
			cloudbrainNum.Date = record.Date
		}

		results = append(results, cloudbrainNum)
	}
	// 按日期升序排序（"YYYY/MM"格式可直接按字符串比较）
	sort.Slice(results, func(i, j int) bool {
		return results[i].Date < results[j].Date
	})

	return results, nil
}

func SumDateCloudbrainNums(entries []DateCloudbrainNum) DateCloudbrainNum {
	if len(entries) == 0 {
		return DateCloudbrainNum{}
	}

	// 初始化结果结构体
	result := DateCloudbrainNum{
		Date:                      "TOTAL",
		CloudOneJobTypeRes:        make(map[string]int),
		CloudTwoJobTypeRes:        make(map[string]int),
		IntelligentNetJobTypeRes:  make(map[string]int),
		CDCenterJobTypeRes:        make(map[string]int),
		CloudBrainPeriodNum:       make(map[int]int),
		CloudBrainComputeResource: make(map[string]int),
	}

	// 遍历所有条目进行累加
	for _, entry := range entries {
		// 累加CloudOneJobTypeRes
		for k, v := range entry.CloudOneJobTypeRes {
			result.CloudOneJobTypeRes[k] += v
		}

		// 累加CloudTwoJobTypeRes
		for k, v := range entry.CloudTwoJobTypeRes {
			result.CloudTwoJobTypeRes[k] += v
		}

		// 累加IntelligentNetJobTypeRes
		for k, v := range entry.IntelligentNetJobTypeRes {
			result.IntelligentNetJobTypeRes[k] += v
		}

		// 累加CDCenterJobTypeRes
		for k, v := range entry.CDCenterJobTypeRes {
			result.CDCenterJobTypeRes[k] += v
		}

		// 累加CloudBrainPeriodNum
		for k, v := range entry.CloudBrainPeriodNum {
			result.CloudBrainPeriodNum[k] += v
		}

		// 累加CloudBrainComputeResource
		for k, v := range entry.CloudBrainComputeResource {
			result.CloudBrainComputeResource[k] += v
		}
	}

	return result
}

func GetAllCloudbrainsTrendDetail(ctx *context.Context) {
	queryType := ctx.QueryTrim("type")
	now := time.Now()

	beginTimeStr := ctx.QueryTrim("beginTime")
	endTimeStr := ctx.QueryTrim("endTime")
	var beginTime time.Time
	var endTime time.Time
	var endTimeTemp time.Time
	dayCloudbrainInfo := make([]DateCloudbrainInfo, 0)
	var err error
	var count int
	if queryType != "" {
		if queryType == "all" {
			recordCloudbrain, err := models.GetRecordBeginTime()
			if err != nil {
				log.Error("Can not get recordCloudbrain", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("repo.record_begintime_get_err"))
				return
			}
			brainRecordBeginTime := recordCloudbrain[0].Cloudbrain.CreatedUnix.AsTime()
			beginTime = brainRecordBeginTime
			endTime = now
			dayCloudbrainInfo, count, err = getMonthCloudbrainInfo(beginTime, endTime)
			if err != nil {
				log.Error("Can not query getDayCloudbrainInfo.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("getDayCloudbrainInfo_get_error"))
				return
			}
		} else if queryType == "today" {
			beginTime = now.AddDate(0, 0, 0)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = now
			dayCloudbrainInfo, count, err = getDayCloudbrainInfo(beginTime, endTime)
			if err != nil {
				log.Error("Can not query getDayCloudbrainInfo.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("getDayCloudbrainInfo_get_error"))
				return
			}

		} else if queryType == "yesterday" {
			beginTime = now.AddDate(0, 0, -1)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
			dayCloudbrainInfo, count, err = getDayCloudbrainInfo(beginTime, endTime)
			if err != nil {
				log.Error("Can not query getDayCloudbrainInfo.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("getDayCloudbrainInfo_get_error"))
				return
			}

		} else if queryType == "last_7day" {
			beginTime = now.AddDate(0, 0, -7)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = now.AddDate(0, 0, 0)
			endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, now.Location())
			endTimeTemp = time.Date(endTimeTemp.Year(), endTimeTemp.Month(), endTimeTemp.Day(), 0, 0, 0, 0, now.Location())
			dayCloudbrainInfo, count, err = getDayCloudbrainInfo(beginTime, endTime)
			if err != nil {
				log.Error("Can not query getDayCloudbrainInfo.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("getDayCloudbrainInfo_get_error"))
				return
			}
		} else if queryType == "last_30day" {
			beginTime = now.AddDate(0, 0, -30)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = now.AddDate(0, 0, 0)
			endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, now.Location())
			dayCloudbrainInfo, count, err = getDayCloudbrainInfo(beginTime, endTime)
			if err != nil {
				log.Error("Can not query getDayCloudbrainInfo.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("getDayCloudbrainInfo_get_error"))
				return
			}
		} else if queryType == "current_month" {
			endTime = now
			beginTime = time.Date(endTime.Year(), endTime.Month(), 1, 0, 0, 0, 0, now.Location())
			dayCloudbrainInfo, count, err = getDayCloudbrainInfo(beginTime, endTime)
			if err != nil {
				log.Error("Can not query getDayCloudbrainInfo.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("getDayCloudbrainInfo_get_error"))
				return
			}

		} else if queryType == "current_year" {
			endTime = now
			beginTime = time.Date(endTime.Year(), 1, 1, 0, 0, 0, 0, now.Location())
			dayCloudbrainInfo, count, err = getMonthCloudbrainInfo(beginTime, endTime)
			if err != nil {
				log.Error("Can not query getDayCloudbrainInfo.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("getDayCloudbrainInfo_get_error"))
				return
			}

		} else if queryType == "last_month" {

			lastMonthTime := now.AddDate(0, -1, 0)
			beginTime = time.Date(lastMonthTime.Year(), lastMonthTime.Month(), 1, 0, 0, 0, 0, now.Location())
			endTime = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
			dayCloudbrainInfo, count, err = getDayCloudbrainInfo(beginTime, endTime)
			if err != nil {
				log.Error("Can not query getDayCloudbrainInfo.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("getDayCloudbrainInfo_get_error"))
				return
			}

		}

	} else {
		if beginTimeStr == "" || endTimeStr == "" {
			//如果查询类型和开始时间结束时间都未设置，按queryType=all处理
			recordCloudbrain, err := models.GetRecordBeginTime()
			if err != nil {
				log.Error("Can not get recordCloudbrain", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("repo.record_begintime_get_err"))
				return
			}
			brainRecordBeginTime := recordCloudbrain[0].Cloudbrain.CreatedUnix.AsTime()
			beginTime = brainRecordBeginTime
			endTime = now
			dayCloudbrainInfo, count, err = getMonthCloudbrainInfo(beginTime, endTime)
			if err != nil {
				log.Error("Can not query getDayCloudbrainInfo.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("getDayCloudbrainInfo_get_error"))
				return
			}
		} else {
			beginTime, err = time.ParseInLocation("2006-01-02", beginTimeStr, time.Local)
			if err != nil {
				log.Error("Can not ParseInLocation.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("ParseInLocation_get_error"))
				return
			}
			endTime, err = time.ParseInLocation("2006-01-02", endTimeStr, time.Local)
			if err != nil {
				log.Error("Can not ParseInLocation.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("ParseInLocation_get_error"))
				return
			}
			if endTime.After(time.Now()) {
				endTime = time.Now()
			}
			endTimeTemp = beginTime.AddDate(0, 0, 1)
			dayCloudbrainInfo, count, err = getDayCloudbrainInfo(beginTime, endTime)
			if err != nil {
				log.Error("Can not query getDayCloudbrainInfo.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("getDayCloudbrainInfo_get_error"))
				return
			}
		}

	}

	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pagesize := ctx.QueryInt("pagesize")
	if pagesize <= 0 {
		pagesize = 5
	}
	pageDateCloudbrainInfo := getPageDateCloudbrainInfo(dayCloudbrainInfo, page, pagesize)

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"totalCount":         count,
		"timeCloudbrainInfo": pageDateCloudbrainInfo,
	})

}
func getPageDateCloudbrainInfo(dateCloudbrainInfo []DateCloudbrainInfo, page int, pagesize int) []DateCloudbrainInfo {
	begin := (page - 1) * pagesize
	end := (page) * pagesize

	if begin > len(dateCloudbrainInfo)-1 {
		return nil
	}
	if end > len(dateCloudbrainInfo)-1 {
		return dateCloudbrainInfo[begin:]
	} else {
		return dateCloudbrainInfo[begin:end]
	}

}

func getPageDateCloudbrainDuration(dateUsageStatistic []models.DateUsageStatistic, page int, pagesize int) []models.DateUsageStatistic {
	begin := (page - 1) * pagesize
	end := (page) * pagesize

	if begin > len(dateUsageStatistic)-1 {
		return nil
	}
	if end > len(dateUsageStatistic)-1 {
		return dateUsageStatistic[begin:]
	} else {
		return dateUsageStatistic[begin:end]
	}

}

func getPageDateCloudbrainTask(dateTaskStatistic []models.DateCenterTaskStatistic, page int, pagesize int) []models.DateCenterTaskStatistic {
	begin := (page - 1) * pagesize
	end := (page) * pagesize

	if begin > len(dateTaskStatistic)-1 {
		return nil
	}
	if end > len(dateTaskStatistic)-1 {
		return dateTaskStatistic[begin:]
	} else {
		return dateTaskStatistic[begin:end]
	}

}

func getPageDateQueueTask(dateTaskStatistic []models.DateQueueTaskStatistic, page int, pagesize int) []models.DateQueueTaskStatistic {
	begin := (page - 1) * pagesize
	end := (page) * pagesize

	if begin > len(dateTaskStatistic)-1 {
		return nil
	}
	if end > len(dateTaskStatistic)-1 {
		return dateTaskStatistic[begin:]
	} else {
		return dateTaskStatistic[begin:end]
	}

}

func GetAllCloudbrainsPeriodDistribution(ctx *context.Context) {
	queryType := ctx.QueryTrim("type")
	beginTimeStr := ctx.QueryTrim("beginTime")
	endTimeStr := ctx.QueryTrim("endTime")
	recordCloudbrain, err := models.GetRecordBeginTime()
	if err != nil {
		log.Error("Can not get record begin time", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.record_begintime_get_err"))
		return
	}
	recordBeginTime := time.Unix(int64(recordCloudbrain[0].Cloudbrain.CreatedUnix), 0)
	beginTime, endTime, err := getCloudbrainTimePeroid(ctx, recordBeginTime)
	if err != nil {
		log.Error("getCloudbrainTimePeroid error:", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.parameter_is_wrong"))
		return
	}
	cloudOneJobTypeRes := make(map[string]int)
	cloudTwoJobTypeRes := make(map[string]int)
	intelligentNetJobTypeRes := make(map[string]int)
	cDCenterJobTypeRes := make(map[string]int)
	cloudBrainPeriodNum := make(map[int]int)
	cloudBrainComputeResource := make(map[string]int)
	beginTimeTemp := beginTime.Unix()
	if queryType == "all" || (queryType == "" && (beginTimeStr == "" || endTimeStr == "")) {
		beginTimeTemp = int64(recordCloudbrain[0].Cloudbrain.CreatedUnix)
	}

	page := 1
	pagesize := 10000
	count := pagesize
	//Each time a maximum of 10000 pieces of data are detected to the memory, batch processing
	for count == pagesize && count != 0 {
		cloudbrains, _, err := models.CloudbrainAllStatic(&models.CloudbrainsOptions{
			ListOptions: models.ListOptions{
				Page:     page,
				PageSize: pagesize,
			},
			Type:          models.TypeCloudBrainAll,
			BeginTimeUnix: beginTimeTemp,
			EndTimeUnix:   endTime.Unix(),
		})
		if err != nil {
			ctx.ServerError("Get cloudbrains failed:", err)
			return
		}

		for _, cloudbrain := range cloudbrains {
			if cloudbrain.Cloudbrain.Type == models.TypeCloudBrainOne {
				if _, ok := cloudOneJobTypeRes[cloudbrain.JobType]; !ok {
					cloudOneJobTypeRes[cloudbrain.JobType] = 1
				} else {
					cloudOneJobTypeRes[cloudbrain.JobType] += 1
				}
			}
			if cloudbrain.Cloudbrain.Type == models.TypeCloudBrainTwo {
				if _, ok := cloudTwoJobTypeRes[cloudbrain.JobType]; !ok {
					cloudTwoJobTypeRes[cloudbrain.JobType] = 1
				} else {
					cloudTwoJobTypeRes[cloudbrain.JobType] += 1
				}
			}
			if cloudbrain.Cloudbrain.Type == models.TypeC2Net {
				if _, ok := intelligentNetJobTypeRes[cloudbrain.JobType]; !ok {
					intelligentNetJobTypeRes[cloudbrain.JobType] = 1
				} else {
					intelligentNetJobTypeRes[cloudbrain.JobType] += 1
				}
			}
			if cloudbrain.Cloudbrain.Type == models.TypeCDCenter {
				if _, ok := cDCenterJobTypeRes[cloudbrain.JobType]; !ok {
					cDCenterJobTypeRes[cloudbrain.JobType] = 1
				} else {
					cDCenterJobTypeRes[cloudbrain.JobType] += 1
				}
			}

			if _, ok := cloudBrainPeriodNum[cloudbrain.Cloudbrain.Type]; !ok {
				cloudBrainPeriodNum[cloudbrain.Cloudbrain.Type] = 1
			} else {
				cloudBrainPeriodNum[cloudbrain.Cloudbrain.Type] += 1
			}

			if _, ok := cloudBrainComputeResource[cloudbrain.Cloudbrain.ComputeResource]; !ok {
				cloudBrainComputeResource[cloudbrain.Cloudbrain.ComputeResource] = 1
			} else {
				cloudBrainComputeResource[cloudbrain.Cloudbrain.ComputeResource] += 1
			}
		}
		count = len(cloudbrains)
		page += 1
	}

	jobTypeList := models.AllJobType()
	for _, v := range jobTypeList {
		if _, ok := cloudOneJobTypeRes[v]; !ok {
			cloudOneJobTypeRes[v] = 0
		}
		if _, ok := cloudTwoJobTypeRes[v]; !ok {
			cloudTwoJobTypeRes[v] = 0
		}
		if _, ok := intelligentNetJobTypeRes[v]; !ok {
			intelligentNetJobTypeRes[v] = 0
		}
		if _, ok := cDCenterJobTypeRes[v]; !ok {
			cDCenterJobTypeRes[v] = 0
		}
	}
	cloudBrainTypeList := []int{0, 1, 2, 3}
	for _, v := range cloudBrainTypeList {
		if _, ok := cloudBrainPeriodNum[v]; !ok {
			cloudBrainPeriodNum[v] = 0
		}
	}

	ComputeResourceList := []string{"CPU/GPU", "NPU", "GCU", "MLU"}
	for _, v := range ComputeResourceList {
		if _, ok := cloudBrainComputeResource[v]; !ok {
			cloudBrainComputeResource[v] = 0
		}
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"cloudOneJobTypeRes":        cloudOneJobTypeRes,
		"cloudTwoJobTypeRes":        cloudTwoJobTypeRes,
		"intelligentNetJobTypeRes":  intelligentNetJobTypeRes,
		"cDCenterJobTypeRes":        cDCenterJobTypeRes,
		"cloudBrainPeriodNum":       cloudBrainPeriodNum,
		"cloudBrainComputeResource": cloudBrainComputeResource,
	})
}

func GetCloudbrainsStatusAnalysis(ctx *context.Context) {
	cloudbrainStatusCount, err := models.GetCloudbrainStatusCount()
	log.Info("cloudbrainStatusCount:", cloudbrainStatusCount)
	if err != nil {
		log.Error("Can not query cloudbrainStatusCount.", err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"cloudbrainStatusCount": cloudbrainStatusCount,
	})
}

func GetCloudbrainsDetailData(ctx *context.Context) {
	BeginTime, EndTime := getTaskBeginAndEndTime(ctx)

	listType := ctx.Query("listType")
	jobType := ctx.Query("jobType")
	jobStatus := ctx.Query("jobStatus")
	cloudBrainType := ctx.QueryInt("Type")
	aiCenter := ctx.Query("aiCenter")
	needDeleteInfo := ctx.Query("needDeleteInfo")

	accCardType := ctx.Query("accCardType")
	accCardsNum := ctx.QueryInt("accCardsNum")
	workServerNumber := ctx.QueryInt("workServerNumber")
	queueId := ctx.QueryInt64("queueId")

	if cloudBrainType == models.TypeCloudBrainOne && aiCenter == models.AICenterOfCloudBrainOne {
		aiCenter = ""
	}
	if cloudBrainType == models.TypeCloudBrainTwo && aiCenter == models.AICenterOfCloudBrainTwo {
		aiCenter = ""
	}
	if cloudBrainType == models.TypeCDCenter && aiCenter == models.AICenterOfChengdu {
		aiCenter = ""
	}
	if cloudBrainType == models.TypeCloudBrainAll {
		if aiCenter == models.AICenterOfCloudBrainOne {
			cloudBrainType = models.TypeCloudBrainOne
			aiCenter = ""
		}
		if aiCenter == models.AICenterOfCloudBrainTwo {
			cloudBrainType = models.TypeCloudBrainTwo
			aiCenter = ""
		}
		if aiCenter == models.AICenterOfChengdu {
			cloudBrainType = models.TypeCDCenter
			aiCenter = ""
		}
	}

	page := ctx.QueryInt("page")
	pageSize := ctx.QueryInt("pagesize")
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	var jobTypes []string
	jobTypeNot := false
	if jobType == string(models.JobTypeBenchmark) {
		jobTypes = models.AllBenchMarkJobType()
	} else if jobType != "all" && jobType != "" {
		jobTypes = append(jobTypes, jobType)
	}

	var jobStatuses []string
	jobStatusNot := false
	if jobStatus == "other" {
		jobStatusNot = true
		jobStatuses = append(jobStatuses, string(models.ModelArtsTrainJobWaiting), string(models.ModelArtsTrainJobFailed), string(models.ModelArtsRunning), string(models.ModelArtsTrainJobCompleted),
			string(models.ModelArtsStarting), string(models.ModelArtsRestarting), string(models.ModelArtsStartFailed),
			string(models.ModelArtsStopping), string(models.ModelArtsStopped), string(models.JobSucceeded))
	} else if jobStatus != "all" && jobStatus != "" {
		jobStatuses = append(jobStatuses, jobStatus)
	}

	keyword := strings.Trim(ctx.Query("q"), " ")

	ciTasks, count, err := models.CloudbrainAll(&models.CloudbrainsOptions{
		ListOptions: models.ListOptions{
			Page:     page,
			PageSize: pageSize,
		},
		Keyword:          keyword,
		Type:             cloudBrainType,
		ComputeResource:  listType,
		JobTypeNot:       jobTypeNot,
		JobStatusNot:     jobStatusNot,
		JobStatus:        jobStatuses,
		JobTypes:         jobTypes,
		NeedRepoInfo:     true,
		BeginTimeUnix:    BeginTime.Unix(),
		EndTimeUnix:      EndTime.Unix(),
		AiCenter:         aiCenter,
		NeedDeleteInfo:   needDeleteInfo,
		AccCardType:      accCardType,
		AccCardsNum:      accCardsNum,
		WorkServerNumber: workServerNumber,
		QueueId:          queueId,
	})
	if err != nil {
		ctx.ServerError("Get job failed:", err)
		return
	}
	models.LoadSpecs4CloudbrainInfo(ciTasks)
	nilTime := time.Time{}
	tasks := []models.TaskDetail{}
	for i, task := range ciTasks {
		task = cloudbrainService.UpdateCloudbrainAiCenter(task)
		var taskDetail models.TaskDetail
		taskDetail.ID = ciTasks[i].Cloudbrain.ID
		taskDetail.JobID = ciTasks[i].Cloudbrain.JobID
		taskDetail.JobName = ciTasks[i].JobName
		taskDetail.DisplayJobName = ciTasks[i].DisplayJobName
		taskDetail.Status = ciTasks[i].Status
		taskDetail.DetailedStatus = ciTasks[i].DetailedStatus
		taskDetail.JobType = ciTasks[i].JobType
		taskDetail.CreatedUnix = ciTasks[i].Cloudbrain.CreatedUnix
		taskDetail.RunTime = ciTasks[i].Cloudbrain.TrainJobDuration
		taskDetail.StartTime = ciTasks[i].StartTime
		taskDetail.EndTime = ciTasks[i].EndTime
		taskDetail.ComputeResource = ciTasks[i].ComputeResource
		taskDetail.Type = ciTasks[i].Cloudbrain.Type
		taskDetail.UserName = ciTasks[i].User.Name
		taskDetail.RepoID = ciTasks[i].RepoID
		// 只有任务成功发起并返回了AiCenter信息后才显示计算中心
		if task.Cloudbrain.AiCenter != "" {
			taskDetail.AiCenter = repo.GetAiCenterNameByCode(task.Cloudbrain.AiCenter, ctx.Language())
		}
		// 只有任务成功发起并返回了AiCenter信息后才显示资源池队列
		if ciTasks[i].Cloudbrain.AiCenter == "" && ciTasks[i].Spec != nil {
			ciTasks[i].Spec.AiCenterName = ""
			ciTasks[i].Spec.QueueName = ""
			ciTasks[i].Spec.AiCenterCode = ""
		}
		if ciTasks[i].Repo != nil {
			taskDetail.RepoName = ciTasks[i].Repo.OwnerName + "/" + ciTasks[i].Repo.Name
			taskDetail.RepoAlias = ciTasks[i].Repo.OwnerName + "/" + ciTasks[i].Repo.Alias
		}
		if ciTasks[i].Cloudbrain.WorkServerNumber >= 1 {
			taskDetail.WorkServerNum = int64(ciTasks[i].Cloudbrain.WorkServerNumber)
		} else {
			taskDetail.WorkServerNum = 1
		}
		taskDetail.CardDuration = repo.GetCloudbrainCardDuration(ciTasks[i].Cloudbrain)
		taskDetail.WaitTime = repo.GetCloudbrainWaitTime(ciTasks[i].Cloudbrain)

		if ciTasks[i].Cloudbrain.DeletedAt != nilTime {
			taskDetail.IsDelete = true
		} else {
			taskDetail.IsDelete = false
		}
		taskDetail.Spec = ciTasks[i].Spec
		tasks = append(tasks, taskDetail)
	}
	pager := context.NewPagination(int(count), pageSize, page, util.GetTotalPage(count, pageSize))
	pager.SetDefaultParams(ctx)
	pager.AddParam(ctx, "listType", "ListType")

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Title":   ctx.Tr("kanban.cloudBrains"),
		"Tasks":   tasks,
		"Keyword": keyword,
		"pager":   pager,
		"count":   count,
	})
}

func GetCartStatisticData(ctx *context.Context) {

	dataType := ctx.QueryTrim("type")
	category := ctx.QueryTrim("category")

	if _, ok := models.XpuInfoType[dataType]; !ok {
		ctx.Error(http.StatusBadRequest)
		return
	}
	if _, ok := models.XpuInfoCategories[category]; !ok {
		ctx.Error(http.StatusBadRequest)
		return
	}

	result, err := models.GetXPUStatisticInfos(models.XpuInfoType[dataType], category)
	if err != nil {
		log.Error("can not get statistic info", err)
		ctx.JSON(http.StatusOK, response.OuterSuccessWithData([]*models.XPUInfoStatisticShow{}))
	}
	ctx.JSON(http.StatusOK, response.OuterSuccessWithData(result))
}

func GetCloudbrainsCreateHoursData(ctx *context.Context) {
	recordCloudbrain, err := models.GetRecordBeginTime()
	if err != nil {
		log.Error("Can not get record begin time", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.record_begintime_get_err"))
		return
	}
	recordBeginTime := time.Unix(int64(recordCloudbrain[0].Cloudbrain.CreatedUnix), 0)
	// recordBeginTime, _ := time.Parse("2006-01-02 15:04:05", recordTime)
	now := time.Now()
	queryType := ctx.QueryTrim("type")
	beginTimeStr := ctx.QueryTrim("beginTime")
	endTimeStr := ctx.QueryTrim("endTime")
	var beginTime time.Time
	var endTime time.Time
	if queryType != "" {
		if queryType == "all" {
			beginTime = recordBeginTime
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = now.AddDate(0, 0, 1)
			endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, now.Location())
		} else if queryType == "today" {
			beginTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
			endTime = beginTime.AddDate(0, 0, 1)
			endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, now.Location())
		} else if queryType == "yesterday" {
			beginTime = now.AddDate(0, 0, -1)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		} else if queryType == "current_week" {
			beginTime = now.AddDate(0, 0, -int(time.Now().Weekday())+1) //begin from monday
			endTime = now
		} else if queryType == "current_month" {
			beginTime = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
			endTime = now.AddDate(0, 0, 1)
			endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, now.Location())
		} else if queryType == "monthly" {
			endTime = now
			beginTime = now.AddDate(0, -1, 0)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
		} else if queryType == "current_year" {
			beginTime = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
			endTime = now.AddDate(0, 0, 1)
			endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, now.Location())
		} else if queryType == "last_month" {
			lastMonthTime := now.AddDate(0, -1, 0)
			beginTime = time.Date(lastMonthTime.Year(), lastMonthTime.Month(), 1, 0, 0, 0, 0, now.Location())
			endTime = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		} else if queryType == "last_7day" {
			beginTime = now.AddDate(0, 0, -7)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = now.AddDate(0, 0, 0)
			endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, now.Location())
		} else if queryType == "last_30day" {
			beginTime = now.AddDate(0, 0, -30)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = now.AddDate(0, 0, 0)
			endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, now.Location())
		}
	} else {
		if beginTimeStr == "" || endTimeStr == "" {
			//如果查询类型和开始时间结束时间都未设置，按queryType=all处理
			beginTime = recordBeginTime
			endTime = now.AddDate(0, 0, 1)
			endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, now.Location())
		} else {
			beginTime, err = time.ParseInLocation("2006-01-02", beginTimeStr, time.Local)
			if err != nil {
				log.Error("Can not ParseInLocation.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("ParseInLocation_get_error"))
				return
			}
			endTime, err = time.ParseInLocation("2006-01-02", endTimeStr, time.Local)
			if err != nil {
				log.Error("Can not ParseInLocation.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("ParseInLocation_get_error"))
				return
			}
			endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, now.Location())
		}
	}
	runPeriodCount := make(map[string]int64)
	createPeriodCount := make(map[string]int64)
	page := 1
	pagesize := 10000
	createCounts := pagesize
	//Each time a maximum of 10000 pieces of data are detected to the memory, batch processing
	for createCounts == pagesize && createCounts != 0 {
		createCloudbrains, _, err := models.CloudbrainAllStatic(&models.CloudbrainsOptions{
			ListOptions: models.ListOptions{
				Page:     page,
				PageSize: pagesize,
			},
			Type:          models.TypeCloudBrainAll,
			BeginTimeUnix: beginTime.Unix(),
			EndTimeUnix:   endTime.Unix(),
		})
		if err != nil {
			log.Error("Can not query cloudbrain.", err)
			ctx.Error(http.StatusBadRequest, ctx.Tr("cloudbrain_get_error"))
			return
		}
		dateBegin := endTime.AddDate(0, 0, -1)
		dateEnd := endTime
		for beginTime.Before(dateBegin) || beginTime.Equal(dateBegin) {
			createHourPeriodCount, err := models.GetCreateHourPeriodCount(createCloudbrains, dateBegin, dateEnd)
			if err != nil {
				log.Error("Can not query runHourPeriodCount.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("runHourPeriodCount_get_error"))
				return
			}
			var slice = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"}
			for _, v := range slice {
				createPeriodCount[v] = createPeriodCount[v] + createHourPeriodCount[v]
			}
			dateEnd = dateBegin
			dateBegin = dateEnd.AddDate(0, 0, -1)
		}
		createCounts = len(createCloudbrains)
		page += 1
	}
	runPage := 1
	runCounts := pagesize
	for runCounts == pagesize && runCounts != 0 {
		var jobStatuses []string
		jobStatuses = append(jobStatuses, string(models.JobRunning))
		runCloudbrains, _, err := models.CloudbrainAllStatic(&models.CloudbrainsOptions{
			ListOptions: models.ListOptions{
				Page:     runPage,
				PageSize: pagesize,
			},
			Type:              models.TypeCloudBrainAll,
			DateBeginTimeUnix: beginTime.Unix(),
			JobStatus:         jobStatuses,
		})
		if err != nil {
			log.Error("Can not query cloudbrain.", err)
			ctx.Error(http.StatusBadRequest, ctx.Tr("cloudbrain_get_error"))
			return
		}
		dateBegin := endTime.AddDate(0, 0, -1)
		dateEnd := endTime
		for beginTime.Before(dateBegin) || beginTime.Equal(dateBegin) {
			runHourPeriodCount, err := models.GetRunHourPeriodCount(runCloudbrains, dateBegin, dateEnd)
			if err != nil {
				log.Error("Can not query runHourPeriodCount.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("runHourPeriodCount_get_error"))
				return
			}
			var slice = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"}
			for _, v := range slice {
				runPeriodCount[v] = runPeriodCount[v] + runHourPeriodCount[v]
			}
			dateEnd = dateBegin
			dateBegin = dateEnd.AddDate(0, 0, -1)
		}
		runCounts = len(runCloudbrains)
		runPage += 1
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"recordBeginTime":       recordCloudbrain[0].Cloudbrain.CreatedUnix,
		"updateTime":            now.Unix(),
		"createHourPeriodCount": createPeriodCount,
		"runHourPeriodCount":    runPeriodCount,
	})

}
func GetWaittingTop(ctx *context.Context) {
	ciTasks, err := models.GetWaittingTop()
	if err != nil {
		log.Error("Can not get record begin time", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.record_begintime_get_err"))
		return
	}
	tasks := []models.TaskDetail{}
	for i, task := range ciTasks {
		ciTasks[i].Cloudbrain.ComputeResource = task.ComputeResource

		var taskDetail models.TaskDetail
		taskDetail.DisplayJobName = ciTasks[i].DisplayJobName
		taskDetail.UserName = ciTasks[i].User.Name
		taskDetail.RepoID = ciTasks[i].RepoID
		if ciTasks[i].Repo != nil {
			taskDetail.RepoName = ciTasks[i].Repo.OwnerName + "/" + ciTasks[i].Repo.Name
		} else {
			taskDetail.RepoName = ""
		}
		WaitTimeInt := time.Now().Unix() - ciTasks[i].Cloudbrain.CreatedUnix.AsTime().Unix()
		taskDetail.WaitTime = models.ConvertDurationToStr(WaitTimeInt)

		if WaitTimeInt < 0 {
			taskDetail.WaitTime = "00:00:00"
		}

		taskDetail.ID = ciTasks[i].Cloudbrain.ID
		taskDetail.ComputeResource = ciTasks[i].Cloudbrain.ComputeResource
		taskDetail.JobType = ciTasks[i].Cloudbrain.JobType
		taskDetail.JobID = ciTasks[i].Cloudbrain.JobID
		taskDetail.Type = ciTasks[i].Cloudbrain.Type

		tasks = append(tasks, taskDetail)
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"tasks": tasks,
	})
}
func GetRunningTop(ctx *context.Context) {
	ciTasks, err := models.GetRunningTop()
	if err != nil {
		log.Error("Can not get record begin time", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.record_begintime_get_err"))
		return
	}
	tasks := []models.TaskDetail{}
	for i, task := range ciTasks {
		ciTasks[i].Cloudbrain.ComputeResource = task.ComputeResource

		var taskDetail models.TaskDetail
		taskDetail.DisplayJobName = ciTasks[i].DisplayJobName
		taskDetail.RunTime = ciTasks[i].Cloudbrain.TrainJobDuration
		taskDetail.UserName = ciTasks[i].User.Name
		taskDetail.RepoID = ciTasks[i].RepoID
		if ciTasks[i].Repo != nil {
			taskDetail.RepoName = ciTasks[i].Repo.OwnerName + "/" + ciTasks[i].Repo.Name
		}

		taskDetail.ID = ciTasks[i].Cloudbrain.ID
		taskDetail.ComputeResource = ciTasks[i].Cloudbrain.ComputeResource
		taskDetail.JobType = ciTasks[i].Cloudbrain.JobType
		taskDetail.JobID = ciTasks[i].Cloudbrain.JobID
		taskDetail.Type = ciTasks[i].Cloudbrain.Type

		tasks = append(tasks, taskDetail)
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"tasks": tasks,
	})
}

func getCloudbrainCount(beginTime time.Time, endTime time.Time, cloudbrains []*models.CloudbrainInfo) (map[string]int, map[string]int, map[string]int, map[string]int, map[int]int, map[string]int) {

	cloudOneJobTypeRes := make(map[string]int)
	cloudTwoJobTypeRes := make(map[string]int)
	intelligentNetJobTypeRes := make(map[string]int)
	cDCenterJobTypeRes := make(map[string]int)
	cloudBrainPeriodNum := make(map[int]int)
	cloudBrainComputeResource := make(map[string]int)
	for _, cloudbrain := range cloudbrains {
		if int64(cloudbrain.Cloudbrain.CreatedUnix) >= beginTime.Unix() && int64(cloudbrain.Cloudbrain.CreatedUnix) < endTime.Unix() {
			if cloudbrain.Cloudbrain.Type == models.TypeCloudBrainOne {
				if _, ok := cloudOneJobTypeRes[cloudbrain.Cloudbrain.JobType]; !ok {
					cloudOneJobTypeRes[cloudbrain.Cloudbrain.JobType] = 1
				} else {
					cloudOneJobTypeRes[cloudbrain.Cloudbrain.JobType] += 1
				}
			}
			if cloudbrain.Cloudbrain.Type == models.TypeCloudBrainTwo {
				if _, ok := cloudTwoJobTypeRes[cloudbrain.JobType]; !ok {
					cloudTwoJobTypeRes[cloudbrain.Cloudbrain.JobType] = 1
				} else {
					cloudTwoJobTypeRes[cloudbrain.Cloudbrain.JobType] += 1
				}
			}

			if _, ok := cloudBrainPeriodNum[cloudbrain.Cloudbrain.Type]; !ok {
				cloudBrainPeriodNum[cloudbrain.Cloudbrain.Type] = 1
			} else {
				cloudBrainPeriodNum[cloudbrain.Cloudbrain.Type] += 1
			}

			if _, ok := cloudBrainComputeResource[cloudbrain.Cloudbrain.ComputeResource]; !ok {
				cloudBrainComputeResource[cloudbrain.Cloudbrain.ComputeResource] = 1
			} else {
				cloudBrainComputeResource[cloudbrain.Cloudbrain.ComputeResource] += 1
			}

		}
	}

	ComputeResourceList := []string{"CPU/GPU", "NPU", "GCU", "MLU", "CPU"}
	for _, v := range ComputeResourceList {
		if _, ok := cloudBrainComputeResource[v]; !ok {
			cloudBrainComputeResource[v] = 0
		}
	}

	jobTypeList := []string{"DEBUG", "BENCHMARK", "INFERENCE", "TRAIN", "SNN4IMAGENET", "BRAINSCORE"}
	cloudBrainTypeList := []int{0, 1, 2, 3}
	for _, v := range jobTypeList {
		if _, ok := cloudOneJobTypeRes[v]; !ok {
			cloudOneJobTypeRes[v] = 0
		}
		if _, ok := cloudTwoJobTypeRes[v]; !ok {
			cloudTwoJobTypeRes[v] = 0
		}
		if _, ok := intelligentNetJobTypeRes[v]; !ok {
			intelligentNetJobTypeRes[v] = 0
		}
		if _, ok := cDCenterJobTypeRes[v]; !ok {
			cDCenterJobTypeRes[v] = 0
		}
	}
	for _, v := range cloudBrainTypeList {
		if _, ok := cloudBrainPeriodNum[v]; !ok {
			cloudBrainPeriodNum[v] = 0
		}
	}
	cloudBrainPeriodNum[-1] = cloudBrainPeriodNum[0] + cloudBrainPeriodNum[1] + cloudBrainPeriodNum[2] + cloudBrainPeriodNum[3]
	return cloudOneJobTypeRes, cloudTwoJobTypeRes, intelligentNetJobTypeRes, cDCenterJobTypeRes, cloudBrainPeriodNum, cloudBrainComputeResource
}

func getDayCloudbrainNum(beginTime time.Time, endTime time.Time) ([]DateCloudbrainNum, error) {
	var endTimeTemp time.Time
	endTimeTemp = beginTime.AddDate(0, 0, 1)
	cloudbrains, _, err := models.CloudbrainAllStatic(&models.CloudbrainsOptions{
		Type:          models.TypeCloudBrainAll,
		BeginTimeUnix: beginTime.Unix(),
		EndTimeUnix:   endTime.Unix(),
	})
	if err != nil {
		log.Error("Get cloudbrains failed:", err)
		return nil, err
	}
	dayCloudbrainNum := make([]DateCloudbrainNum, 0)
	for endTimeTemp.Before(endTime) || endTimeTemp.Equal(endTime) {
		cloudOneJobTypeRes, cloudTwoJobTypeRes, intelligentNetJobTypeRes, cDCenterJobTypeRes, cloudBrainPeriodNum, cloudBrainComputeResource := getCloudbrainCount(beginTime, endTimeTemp, cloudbrains)
		dayCloudbrainNum = append(dayCloudbrainNum, DateCloudbrainNum{
			Date:                      beginTime.Format("2006/01/02"),
			CloudOneJobTypeRes:        cloudOneJobTypeRes,
			CloudTwoJobTypeRes:        cloudTwoJobTypeRes,
			IntelligentNetJobTypeRes:  intelligentNetJobTypeRes,
			CDCenterJobTypeRes:        cDCenterJobTypeRes,
			CloudBrainPeriodNum:       cloudBrainPeriodNum,
			CloudBrainComputeResource: cloudBrainComputeResource,
		})
		if endTime.Before(endTimeTemp.AddDate(0, 0, 1)) && endTimeTemp.Before(endTime) {
			beginTime = endTimeTemp
			endTimeTemp = endTime
		} else {
			beginTime = endTimeTemp
			endTimeTemp = beginTime.AddDate(0, 0, 1)
		}
	}
	return dayCloudbrainNum, nil
}
func getMonthCloudbrainNum(beginTime time.Time, endTime time.Time) ([]DateCloudbrainNum, error) {
	var endTimeTemp time.Time
	now := time.Now()
	endTimeTemp = beginTime.AddDate(0, 1, 0)
	endTimeTemp = time.Date(endTimeTemp.Year(), endTimeTemp.Month(), 1, 0, 0, 0, 0, now.Location())
	monthCloudbrainNum := make([]DateCloudbrainNum, 0)
	cloudbrains, _, err := models.CloudbrainAllStatic(&models.CloudbrainsOptions{
		Type:          models.TypeCloudBrainAll,
		BeginTimeUnix: beginTime.Unix(),
		EndTimeUnix:   endTime.Unix(),
	})
	if err != nil {
		log.Error("Getcloudbrains failed:%v", err)
		return nil, err
	}
	if endTime.Before(endTimeTemp) {
		endTimeTemp = endTime
	}
	for endTimeTemp.Before(endTime) || endTimeTemp.Equal(endTime) {
		cloudOneJobTypeRes, cloudTwoJobTypeRes, intelligentNetJobTypeRes, cDCenterJobTypeRes, cloudBrainPeriodNum, cloudBrainComputeResource := getCloudbrainCount(beginTime, endTimeTemp, cloudbrains)
		dateCloudbrainNum := DateCloudbrainNum{
			Date:                      beginTime.Format("2006/01"),
			CloudOneJobTypeRes:        cloudOneJobTypeRes,
			CloudTwoJobTypeRes:        cloudTwoJobTypeRes,
			IntelligentNetJobTypeRes:  intelligentNetJobTypeRes,
			CDCenterJobTypeRes:        cDCenterJobTypeRes,
			CloudBrainPeriodNum:       cloudBrainPeriodNum,
			CloudBrainComputeResource: cloudBrainComputeResource,
		}
		monthCloudbrainNum = append(monthCloudbrainNum, dateCloudbrainNum)

		// 检查数据库中是否已存在该日期的记录
		counts := models.CheckDateExistsInCloudbrainNumJson(beginTime.Format("2006/01"))
		if err != nil {
			log.Error("Check date exists failed: %v", err.Error())
			return nil, err
		}

		cutoffTime, err := time.Parse("2006/01", "2025/06")
		if err != nil {
			log.Error("Parse cutoff time failed: %v", err.Error())
			return nil, err
		}

		if counts == 0 && beginTime.Before(cutoffTime) {
			data, err := json.Marshal(dateCloudbrainNum)
			if err != nil {
				log.Error("json.Marshal dateCloudbrainNum failed: %v", err.Error())
				return nil, err
			}
			if _, err = models.InsertCloudbrainTaskNumJson(&models.DateCloudbrainNumJson{
				Date:              beginTime.Format("2006/01"),
				CloudbrainNumJson: string(data),
			}); err != nil {
				log.Error("Insert CloudbrainTaskNumJson failed: %v", err.Error())
			}
		}
		if endTime.Before(endTimeTemp.AddDate(0, 1, 0)) && endTimeTemp.Before(endTime) {
			beginTime = endTimeTemp
			endTimeTemp = endTime
		} else {
			beginTime = endTimeTemp
			endTimeTemp = beginTime.AddDate(0, 1, 0)
		}
	}
	return monthCloudbrainNum, nil
}

func getDayCloudbrainInfo(beginTime time.Time, endTime time.Time) ([]DateCloudbrainInfo, int, error) {
	now := time.Now()
	endTimeTemp := time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, now.Location())
	if endTimeTemp.Equal(endTime) {
		endTimeTemp = endTimeTemp.AddDate(0, 0, -1)
	}
	cloudbrains, _, err := models.CloudbrainAllStatic(&models.CloudbrainsOptions{
		Type:          models.TypeCloudBrainAll,
		BeginTimeUnix: beginTime.Unix(),
		EndTimeUnix:   endTime.Unix(),
	})
	if err != nil {
		log.Error("Getcloudbrains failed:%v", err)
		return nil, 0, err
	}
	dayCloudbrainInfo := make([]DateCloudbrainInfo, 0)
	count := 0
	for beginTime.Before(endTimeTemp) || beginTime.Equal(endTimeTemp) {
		_, _, _, _, cloudBrainPeriodNum, cloudBrainComputeResource := getCloudbrainCount(endTimeTemp, endTime, cloudbrains)
		dayCloudbrainInfo = append(dayCloudbrainInfo, DateCloudbrainInfo{
			Date:                      endTimeTemp.Format("2006/01/02"),
			CloudBrainPeriodNum:       cloudBrainPeriodNum,
			CloudBrainComputeResource: cloudBrainComputeResource,
		})
		endTime = endTimeTemp
		endTimeTemp = endTimeTemp.AddDate(0, 0, -1)
		count += 1
	}
	return dayCloudbrainInfo, count, nil
}

func getMonthCloudbrainInfo(beginTime time.Time, endTime time.Time) ([]DateCloudbrainInfo, int, error) {
	now := time.Now()
	endTimeTemp := time.Date(endTime.Year(), endTime.Month(), 1, 0, 0, 0, 0, now.Location())
	if endTimeTemp.Equal(endTime) {
		endTimeTemp = endTimeTemp.AddDate(0, -1, 0)
	}
	cloudbrains, _, err := models.CloudbrainAllStatic(&models.CloudbrainsOptions{
		Type:          models.TypeCloudBrainAll,
		BeginTimeUnix: beginTime.Unix(),
		EndTimeUnix:   endTime.Unix(),
	})
	if err != nil {
		log.Error("Getcloudbrains failed:%v", err)
		return nil, 0, err
	}
	dayCloudbrainInfo := make([]DateCloudbrainInfo, 0)
	count := 0
	for beginTime.Before(endTimeTemp) || beginTime.Equal(endTimeTemp) || (endTimeTemp.Before(beginTime) && beginTime.Before(endTime)) {
		_, _, _, _, cloudBrainPeriodNum, cloudBrainComputeResource := getCloudbrainCount(endTimeTemp, endTime, cloudbrains)
		dayCloudbrainInfo = append(dayCloudbrainInfo, DateCloudbrainInfo{
			Date:                      endTimeTemp.Format("2006/01"),
			CloudBrainPeriodNum:       cloudBrainPeriodNum,
			CloudBrainComputeResource: cloudBrainComputeResource,
		})
		endTime = endTimeTemp
		endTimeTemp = endTimeTemp.AddDate(0, -1, 0)
		count += 1
	}
	return dayCloudbrainInfo, count, nil
}

var (
	downloadAllSem     chan struct{}
	downloadAllSemOnce sync.Once
)

// tryAcquireDownloadAll 尝试获取一个全量下载槽位。
// 返回 ok=false 表示已被占用，调用方应立即返回 429。
// 返回 ok=true 时必须调用 release() 释放（建议 defer）。
func tryAcquireDownloadAll() (ok bool, release func()) {
	downloadAllSemOnce.Do(func() {
		if downloadAllSem == nil {
			n := setting.CloudbrainDownloadAllConcurrency
			if n > 0 {
				downloadAllSem = make(chan struct{}, n)
			}
		}
	})
	if downloadAllSem == nil {
		return true, func() {}
	}
	select {
	case downloadAllSem <- struct{}{}:
		return true, func() { <-downloadAllSem }
	default:
		return false, nil
	}
}

func DownloadCloudBrainBoard(ctx *context.Context) {
	if ok, release := tryAcquireDownloadAll(); !ok {
		ctx.Error(http.StatusTooManyRequests, ctx.Tr("repo.cloudbrain_download_busy"))
		return
	} else {
		defer release()
	}

	queryType := ctx.QueryTrim("type")
	now := time.Now()

	beginTimeStr := ctx.QueryTrim("beginTimeStr")
	endTimeStr := ctx.QueryTrim("endTimeStr")
	var beginTime int64
	var endTime int64

	page := 1
	pageSize := 300

	var cloudBrain = ctx.Tr("repo.cloudbrain")
	fileName := getCloudbrainFileName(queryType + cloudBrain)

	if queryType != "" {
		if queryType == "select_time" {
			if beginTimeStr == "" || endTimeStr == "" {
				//默认近30天的数据
				beginTimeUTC := now.AddDate(0, 0, -30)
				beginTimeUTC = time.Date(beginTimeUTC.Year(), beginTimeUTC.Month(), beginTimeUTC.Day(), 0, 0, 0, 0, now.Location())
				endTimeUTC := now.AddDate(0, 0, 0)
				endTimeUTC = time.Date(endTimeUTC.Year(), endTimeUTC.Month(), endTimeUTC.Day(), 0, 0, 0, 0, now.Location())
				beginTime = beginTimeUTC.Unix()
				endTime = endTimeUTC.Unix()
			} else {
				loc, _ := time.LoadLocation("Local") // 或者指定时区如"Asia/Shanghai"
				beginTimeUTC, err := time.ParseInLocation("2006-01-02", beginTimeStr, loc)
				if err != nil {
					log.Info("Error parsing time:", err)
				}
				endTimeUTC, err := time.ParseInLocation("2006-01-02", endTimeStr, loc)
				if err != nil {
					log.Info("Error parsing time:", err)
				}
				beginTime = beginTimeUTC.Unix()
				endTime = endTimeUTC.Unix()
			}
		} else if queryType == "last_30day" {
			beginTimeUTC := now.AddDate(0, 0, -30)
			beginTimeUTC = time.Date(beginTimeUTC.Year(), beginTimeUTC.Month(), beginTimeUTC.Day(), 0, 0, 0, 0, now.Location())
			endTimeUTC := now.AddDate(0, 0, 0)
			endTimeUTC = time.Date(endTimeUTC.Year(), endTimeUTC.Month(), endTimeUTC.Day(), 0, 0, 0, 0, now.Location())
			beginTime = beginTimeUTC.Unix()
			endTime = endTimeUTC.Unix()
		} else if queryType == "current_month" {
			beginTimeUTC := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
			endTimeUTC := now
			beginTime = beginTimeUTC.Unix()
			endTime = endTimeUTC.Unix()
		} else if queryType == "current_year" {
			beginTimeUTC := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
			endTimeUTC := now
			beginTime = beginTimeUTC.Unix()
			endTime = endTimeUTC.Unix()
		}

	} else {
		//默认是近30天的数据
		beginTimeUTC := now.AddDate(0, 0, -30)
		beginTimeUTC = time.Date(beginTimeUTC.Year(), beginTimeUTC.Month(), beginTimeUTC.Day(), 0, 0, 0, 0, now.Location())
		endTimeUTC := now.AddDate(0, 0, 0)
		endTimeUTC = time.Date(endTimeUTC.Year(), endTimeUTC.Month(), endTimeUTC.Day(), 0, 0, 0, 0, now.Location())
		beginTime = beginTimeUTC.Unix()
		endTime = endTimeUTC.Unix()
	}

	total, err := models.CloudbrainTotalForDashBoard(&models.CloudbrainsOptions{
		ListOptions: models.ListOptions{
			Page:     page,
			PageSize: pageSize,
		},
		Type:          models.TypeCloudBrainAll,
		BeginTimeUnix: beginTime,
		EndTimeUnix:   endTime,
		AccCardsNum:   models.AccCardsNumAll,
	})

	if err != nil {
		log.Warn("Can not get cloud brain info", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.cloudbrain_query_fail"))
		return
	}
	totalPage := util.GetTotalPage(total, pageSize)
	f := excelize.NewFile()

	index := f.NewSheet(cloudBrain)
	f.DeleteSheet("Sheet1")

	for k, v := range allCloudbrainHeader(ctx) {
		f.SetCellValue(cloudBrain, k, v)
	}
	var row = 2
	userMap := models.GetAllUserName()
	for i := 0; i < totalPage; i++ {
		log.Info("DownloadCloudBrainBoard now=" + fmt.Sprint(i) + "/" + fmt.Sprint(totalPage))
		pageRecords, err := models.CloudbrainAllForDashBoard(&models.CloudbrainsOptions{
			ListOptions: models.ListOptions{
				Page:     page,
				PageSize: pageSize,
			},
			Type:          models.TypeCloudBrainAll,
			BeginTimeUnix: beginTime,
			EndTimeUnix:   endTime,
			NeedRepoInfo:  true,
			AccCardsNum:   models.AccCardsNumAll,
		})
		if err != nil {
			log.Warn("Can not get cloud brain info", err)
			continue
		}

		models.LoadSpecs4CloudbrainInfo(pageRecords)
		for _, record := range pageRecords {
			record = cloudbrainService.UpdateCloudbrainAiCenter(record)
			record.Cloudbrain.AiCenter = repo.GetAiCenterNameByCode(record.Cloudbrain.AiCenter, ctx.Language())
			for k, v := range allCloudbrainValues(row, record, ctx, userMap) {
				f.SetCellValue(cloudBrain, k, v)
			}
			row++
		}
		page++
	}
	f.SetActiveSheet(index)

	ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+url.QueryEscape(fileName))
	ctx.Resp.Header().Set("Content-Type", "application/octet-stream")

	f.WriteTo(ctx.Resp)
}

func DownloadAitask4Admin(ctx *context.Context) {
	BeginTime, EndTime := getTaskBeginAndEndTime(ctx)
	queryType := ctx.QueryTrim("type")
	page := 1
	pageSize := 300

	var cloudBrain = ctx.Tr("repo.cloudbrain")
	fileName := getCloudbrainFileName(queryType + cloudBrain)

	total, err := models.CloudbrainTotalForDashBoard(&models.CloudbrainsOptions{
		ListOptions: models.ListOptions{
			Page:     page,
			PageSize: pageSize,
		},
		Type:          models.TypeCloudBrainAll,
		BeginTimeUnix: BeginTime.Unix(),
		EndTimeUnix:   EndTime.Unix(),
		AccCardsNum:   models.AccCardsNumAll,
	})

	if err != nil {
		log.Warn("Can not get cloud brain info", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.cloudbrain_query_fail"))
		return
	}
	totalPage := util.GetTotalPage(total, pageSize)
	f := excelize.NewFile()

	index := f.NewSheet(cloudBrain)
	f.DeleteSheet("Sheet1")

	for k, v := range allCloudbrainHeader(ctx) {
		f.SetCellValue(cloudBrain, k, v)
	}
	var row = 2
	userMap := models.GetAllUserName()
	for i := 0; i < totalPage; i++ {
		log.Info("DownloadCloudBrainBoard now=" + fmt.Sprint(i) + "/" + fmt.Sprint(totalPage))
		pageRecords, err := models.CloudbrainAllForDashBoard(&models.CloudbrainsOptions{
			ListOptions: models.ListOptions{
				Page:     page,
				PageSize: pageSize,
			},
			Type:          models.TypeCloudBrainAll,
			BeginTimeUnix: BeginTime.Unix(),
			EndTimeUnix:   EndTime.Unix(),
			NeedRepoInfo:  true,
			AccCardsNum:   models.AccCardsNumAll,
		})
		if err != nil {
			log.Warn("Can not get cloud brain info", err)
			continue
		}

		models.LoadSpecs4CloudbrainInfo(pageRecords)
		for _, record := range pageRecords {
			record = cloudbrainService.UpdateCloudbrainAiCenter(record)
			record.Cloudbrain.AiCenter = repo.GetAiCenterNameByCode(record.Cloudbrain.AiCenter, ctx.Language())
			for k, v := range allCloudbrainValues(row, record, ctx, userMap) {
				f.SetCellValue(cloudBrain, k, v)
			}
			row++
		}
		page++
	}
	f.SetActiveSheet(index)

	ctx.Resp.Header().Set("Content-Disposition", "attachment; filename="+url.QueryEscape(fileName))
	ctx.Resp.Header().Set("Content-Type", "application/octet-stream")

	f.WriteTo(ctx.Resp)
}

func getTaskBeginAndEndTime(ctx *context.Context) (time.Time, time.Time) {
	queryType := ctx.QueryTrim("type")
	now := time.Now()
	beginTimeStr := ctx.QueryTrim("beginTime")
	endTimeStr := ctx.QueryTrim("endTime")

	var beginTime time.Time
	var endTime time.Time
	if queryType != "" {
		if queryType == "all" {
			recordCloudbrain, err := models.GetRecordBeginTime()
			if err != nil {
				log.Error("Can not get recordCloudbrain", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("repo.record_begintime_get_err"))
			}
			brainRecordBeginTime := recordCloudbrain[0].Cloudbrain.CreatedUnix.AsTime()
			beginTime = brainRecordBeginTime
			endTime = now
		} else if queryType == "today" {
			beginTime = now.AddDate(0, 0, 0)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = now

		} else if queryType == "yesterday" {
			beginTime = now.AddDate(0, 0, -1)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		} else if queryType == "last_7day" {
			beginTime = now.AddDate(0, 0, -7)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = now
		} else if queryType == "last_30day" {
			beginTime = now.AddDate(0, 0, -30)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = now
		} else if queryType == "current_month" {
			endTime = now
			beginTime = time.Date(endTime.Year(), endTime.Month(), 1, 0, 0, 0, 0, now.Location())

		} else if queryType == "current_year" {
			endTime = now
			beginTime = time.Date(endTime.Year(), 1, 1, 0, 0, 0, 0, now.Location())
		} else if queryType == "last_month" {
			lastMonthTime := now.AddDate(0, -1, 0)
			beginTime = time.Date(lastMonthTime.Year(), lastMonthTime.Month(), 1, 0, 0, 0, 0, now.Location())
			endTime = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		} else if queryType == "last_year" {
			lastYear := now.Year() - 1                                           // 获取去年的年份
			beginTime = time.Date(lastYear, 1, 1, 0, 0, 0, 0, now.Location())    // 去年1月1日 00:00:00
			endTime = time.Date(lastYear, 12, 31, 23, 59, 59, 0, now.Location()) // 去年12月31日 23:59:59
		}

	} else {
		if beginTimeStr == "" || endTimeStr == "" {
			//如果查询类型和开始时间结束时间都未设置，按queryType=last_30day处理
			beginTime = now.AddDate(0, 0, -30)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = now
		} else {
			var err error
			beginTime, err = time.ParseInLocation("2006-01-02", beginTimeStr, time.Local)
			if err != nil {
				log.Error("Can not ParseInLocation.", err)
			}
			endTime, err = time.ParseInLocation("2006-01-02", endTimeStr, time.Local)
			if err != nil {
				log.Error("Can not ParseInLocation.", err)
			}
		}

	}
	return beginTime, endTime
}

func getCloudbrainFileName(baseName string) string {
	return baseName + "_" + time.Now().Format(EXCEL_DATE_FORMAT) + ".xlsx"

}
func allCloudbrainHeader(ctx *context.Context) map[string]string {

	return map[string]string{"A1": ctx.Tr("repo.cloudbrain_task"), "B1": ctx.Tr("repo.cloudbrain_type"), "C1": ctx.Tr("repo.modelarts.status"), "D1": ctx.Tr("repo.cloudbrain_task_type"),
		"E1": ctx.Tr("repo.modelarts.createtime"), "F1": ctx.Tr("repo.modelarts.train_job.wait_time"), "G1": ctx.Tr("repo.modelarts.train_job.dura_time"),
		"H1": ctx.Tr("cloudbrain.card_duration"),
		"I1": ctx.Tr("repo.modelarts.train_job.start_time"), "J1": ctx.Tr("repo.modelarts.train_job.end_time"),
		"K1": ctx.Tr("repo.modelarts.computing_resources"), "L1": ctx.Tr("cloudbrain.card_type"),
		"M1": ctx.Tr("repo.modelarts.train_job.amount_of_compute_node"), "N1": ctx.Tr("repo.grampus.train_job.ai_center"),
		"O1": ctx.Tr("cloudbrain.resource_specification"), "P1": ctx.Tr("repo.cloudbrain_creator"), "Q1": ctx.Tr("repo.repo_name"),
		"R1": ctx.Tr("repo.cloudbrain_task_name"), "S1": ctx.Tr("repo.modelarts.deletetime")}
}
func allCloudbrainValues(row int, rs *models.CloudbrainInfo, ctx *context.Context, userMap map[int64]string) map[string]string {
	return map[string]string{getCellName("A", row): rs.DisplayJobName, getCellName("B", row): getCloudbrainType(rs, ctx), getCellName("C", row): rs.Status, getCellName("D", row): rs.JobType,
		getCellName("E", row): time.Unix(int64(rs.Cloudbrain.CreatedUnix), 0).Format(CREATE_TIME_FORMAT), getCellName("F", row): repo.GetCloudbrainWaitTime(rs.Cloudbrain),
		getCellName("G", row): rs.TrainJobDuration, getCellName("H", row): repo.GetCloudbrainCardDuration(rs.Cloudbrain),
		getCellName("I", row): getBrainStartTime(rs),
		getCellName("J", row): getBrainEndTime(rs), getCellName("K", row): rs.ComputeResource, getCellName("L", row): getCloudbrainCardType(rs),
		getCellName("M", row): getWorkServerNum(rs), getCellName("N", row): rs.Cloudbrain.AiCenter,
		getCellName("O", row): getCloudbrainFlavorName(rs), getCellName("P", row): userMap[rs.UserID],
		getCellName("Q", row): getBrainRepo(rs), getCellName("R", row): rs.JobName, getCellName("S", row): getBrainDeleteTime(rs),
	}
}
func getWorkServerNum(rs *models.CloudbrainInfo) string {
	if rs.Cloudbrain.WorkServerNumber >= 1 {
		return fmt.Sprint(rs.Cloudbrain.WorkServerNumber)
	} else {
		return "1"
	}
}
func getBrainRepo(rs *models.CloudbrainInfo) string {
	if rs.Repo != nil {
		return rs.Repo.OwnerName + "/" + rs.Repo.Alias
	}
	return ""
}
func getBrainStartTime(rs *models.CloudbrainInfo) string {
	timeString := time.Unix(int64(rs.Cloudbrain.StartTime), 0).Format(CREATE_TIME_FORMAT)
	if timeString != "1970/01/01 08:00:00" {
		return timeString
	} else {
		return "0"
	}

}
func getBrainEndTime(rs *models.CloudbrainInfo) string {
	timeString := time.Unix(int64(rs.Cloudbrain.EndTime), 0).Format(CREATE_TIME_FORMAT)
	if timeString != "1970/01/01 08:00:00" {
		return timeString
	} else {
		return "0"
	}

}
func getCloudbrainType(rs *models.CloudbrainInfo, ctx *context.Context) string {
	if rs.Cloudbrain.Type == models.TypeCloudBrainOne {
		return ctx.Tr("repo.cloudbrain1")
	} else if rs.Cloudbrain.Type == models.TypeCloudBrainTwo {
		return ctx.Tr("repo.cloudbrain2")
	} else if rs.Cloudbrain.Type == models.TypeC2Net {
		return ctx.Tr("repo.intelligent_net")
	} else {
		return ctx.Tr("repo.cloudbrain_untype")
	}
}
func getCloudbrainCardType(rs *models.CloudbrainInfo) string {
	if rs.Cloudbrain.Spec != nil {
		return rs.Cloudbrain.Spec.AccCardType
	} else {
		return ""
	}
}
func getCloudbrainFlavorName(rs *models.CloudbrainInfo) string {
	flavorName := repo.GetCloudbrainFlavorName(rs.Cloudbrain)
	return flavorName
}

func getBrainDeleteTime(rs *models.CloudbrainInfo) string {
	nilTime := time.Time{}
	if rs.Cloudbrain.DeletedAt != nilTime {
		return rs.Cloudbrain.DeletedAt.Format("2006-01-02 15:04:05")
	} else {
		return ""
	}
}
func getCloudbrainTimePeroid(ctx *context.Context, recordBeginTime time.Time) (time.Time, time.Time, error) {
	queryType := ctx.QueryTrim("type")
	beginTimeStr := ctx.QueryTrim("beginTime")
	endTimeStr := ctx.QueryTrim("endTime")
	now := time.Now()

	var beginTime time.Time
	var endTime time.Time
	var err error
	if queryType != "" {

		if queryType == "all" {
			beginTime = recordBeginTime
			endTime = now
		} else if queryType == "today" {
			endTime = now
			beginTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, now.Location())
		} else if queryType == "yesterday" {
			endTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
			beginTime = endTime.AddDate(0, 0, -1)

		} else if queryType == "last_7day" {
			beginTime = now.AddDate(0, 0, -7)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = now.AddDate(0, 0, 0)
			endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, now.Location())
		} else if queryType == "last_30day" {
			beginTime = now.AddDate(0, 0, -30)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = now.AddDate(0, 0, 0)
			endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, now.Location())
		} else if queryType == "current_month" {
			endTime = now
			beginTime = time.Date(endTime.Year(), endTime.Month(), 1, 0, 0, 0, 0, now.Location())
		} else if queryType == "monthly" {
			endTime = now
			beginTime = now.AddDate(0, -1, 1)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())

		} else if queryType == "current_year" {
			endTime = now
			beginTime = time.Date(endTime.Year(), 1, 1, 0, 0, 0, 0, now.Location())

		} else if queryType == "last_month" {

			lastMonthTime := now.AddDate(0, -1, 0)
			beginTime = time.Date(lastMonthTime.Year(), lastMonthTime.Month(), 1, 0, 0, 0, 0, now.Location())
			endTime = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

		} else {
			return now, now, fmt.Errorf("The value of type parameter is wrong.")

		}

	} else {
		if beginTimeStr == "" || endTimeStr == "" {
			//如果查询类型和开始时间结束时间都未设置，按queryType=all处理
			beginTime = recordBeginTime
			endTime = now

		} else {

			beginTime, err = time.ParseInLocation("2006-01-02", beginTimeStr, time.Local)
			if err != nil {
				return now, now, err
			}

			endTime, err = time.ParseInLocation("2006-01-02", endTimeStr, time.Local)
			if err != nil {
				return now, now, err
			}
		}

	}

	if beginTime.Before(recordBeginTime) {
		beginTime = recordBeginTime
	}

	return beginTime, endTime, nil
}

func GetCloudbrainResourceOverview(ctx *context.Context) {
	var recordBeginTime timeutil.TimeStamp
	recordCloudbrainDuration, err := models.GetDurationRecordBeginTime()
	if err != nil {
		log.Error("Can not get GetDurationRecordBeginTime", err)
		return
	}
	if len(recordCloudbrainDuration) > 0 && err == nil {
		recordBeginTime = recordCloudbrainDuration[0].DateTimeUnix
	} else {
		recordBeginTime = timeutil.TimeStamp(time.Now().Unix())
	}
	recordUpdateTime := time.Now().Unix()
	resourceQueues, err := models.GetCanUseCardInfo()
	if err != nil {
		log.Info("GetCanUseCardInfo err: %v", err)
		return
	}
	OpenIResourceDetail := []models.ResourceDetail{}
	C2NetResourceDetail := []models.ResourceDetail{}
	for _, resourceQueue := range resourceQueues {
		if resourceQueue.Cluster == models.OpenICluster {
			aiCenterName := repo.GetAiCenterNameByCode(resourceQueue.AiCenterCode, ctx.Language())
			var resourceDetail models.ResourceDetail
			resourceDetail.QueueCode = resourceQueue.QueueCode
			resourceDetail.Cluster = resourceQueue.Cluster
			resourceDetail.AiCenterCode = resourceQueue.AiCenterCode
			resourceDetail.AiCenterName = resourceQueue.AiCenterCode + "/" + aiCenterName
			resourceDetail.ComputeResource = resourceQueue.ComputeResource
			resourceDetail.AccCardType = resourceQueue.AccCardType + "(" + resourceQueue.ComputeResource + ")"

			resourceDetail.IsAutomaticSync = resourceQueue.IsAutomaticSync
			if resourceQueue.IsAvailable && resourceQueue.DeletedTime == 0 {
				resourceDetail.CardsTotalNum = resourceQueue.CardsTotalNum
				OpenIResourceDetail = append(OpenIResourceDetail, resourceDetail)
			} else {
				resourceDetail.CardsTotalNum = 0
			}
		}
		if resourceQueue.Cluster == models.C2NetCluster {
			aiCenterName := repo.GetAiCenterNameByCode(resourceQueue.AiCenterCode, ctx.Language())
			var resourceDetail models.ResourceDetail
			resourceDetail.QueueCode = resourceQueue.QueueCode
			resourceDetail.Cluster = resourceQueue.Cluster
			resourceDetail.AiCenterCode = resourceQueue.AiCenterCode
			resourceDetail.AiCenterName = resourceQueue.AiCenterCode + "/" + aiCenterName
			resourceDetail.ComputeResource = resourceQueue.ComputeResource
			resourceDetail.AccCardType = resourceQueue.AccCardType + "(" + resourceQueue.ComputeResource + ")"
			resourceDetail.IsAutomaticSync = resourceQueue.IsAutomaticSync
			if resourceQueue.IsAvailable && resourceQueue.DeletedTime == 0 {
				resourceDetail.CardsTotalNum = resourceQueue.CardsTotalNum
				C2NetResourceDetail = append(C2NetResourceDetail, resourceDetail)
			} else {
				resourceDetail.CardsTotalNum = 0
			}
		}
	}
	openIResourceNum := make(map[string]map[string]int)

	for _, openIResourceDetail := range OpenIResourceDetail {
		if _, ok := openIResourceNum[openIResourceDetail.AiCenterName]; !ok {
			openIResourceNum[openIResourceDetail.AiCenterName] = make(map[string]int)
		}
		if _, ok := openIResourceNum[openIResourceDetail.AiCenterName][openIResourceDetail.AccCardType]; !ok {
			openIResourceNum[openIResourceDetail.AiCenterName][openIResourceDetail.AccCardType] = openIResourceDetail.CardsTotalNum
		} else {
			openIResourceNum[openIResourceDetail.AiCenterName][openIResourceDetail.AccCardType] += openIResourceDetail.CardsTotalNum
		}
	}

	c2NetResourceNum := make(map[string]map[string]int)
	for _, c2NetResourceDetail := range C2NetResourceDetail {
		if _, ok := c2NetResourceNum[c2NetResourceDetail.AiCenterName]; !ok {
			c2NetResourceNum[c2NetResourceDetail.AiCenterName] = make(map[string]int)
		}
		if _, ok := c2NetResourceNum[c2NetResourceDetail.AiCenterName][c2NetResourceDetail.AccCardType]; !ok {
			c2NetResourceNum[c2NetResourceDetail.AiCenterName][c2NetResourceDetail.AccCardType] = c2NetResourceDetail.CardsTotalNum
		} else {
			c2NetResourceNum[c2NetResourceDetail.AiCenterName][c2NetResourceDetail.AccCardType] += c2NetResourceDetail.CardsTotalNum
		}

	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"openI":            openIResourceNum,
		"c2Net":            c2NetResourceNum,
		"recordUpdateTime": recordUpdateTime,
		"recordBeginTime":  recordBeginTime,
	})
}

func GetCloudbrainResourceUsageDetail(ctx *context.Context) {
	aiCenterCode := ctx.QueryTrim("aiCenterCode")
	if aiCenterCode == "" {
		aiCenterCode = models.AICenterOfCloudBrainOne
	}
	beginTime, endTime := getBeginAndEndTime(ctx)
	dayCloudbrainDuration, count, err := getDayCloudbrainDuration(beginTime, endTime, aiCenterCode)
	if err != nil {
		log.Error("Can not query dayCloudbrainDuration.", err)
		return
	}
	hourCloudbrainDuration, err := getHourCloudbrainDuration(beginTime, endTime, aiCenterCode)
	if err != nil {
		log.Error("Can not query hourCloudbrainDuration.", err)
		return
	}
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pagesize := ctx.QueryInt("pagesize")
	if pagesize <= 0 {
		pagesize = 36500
	}
	pageDateCloudbrainDuration := getPageDateCloudbrainDuration(dayCloudbrainDuration, page, pagesize)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"totalCount":                 count,
		"pageDateCloudbrainDuration": pageDateCloudbrainDuration,
		"hourCloudbrainDuration":     hourCloudbrainDuration,
	})
}

func GetQueueDurationUsageDetail(ctx *context.Context) {
	queueID := ctx.QueryInt64("queueID")
	beginTime, endTime := getBeginAndEndTime(ctx)

	dayQueueTaskInfo, count, err := getDayQueueTaskDuration(beginTime, endTime, queueID)
	if err != nil {
		log.Error("Can not query dayQueueTaskInfo.", err)
		return
	}
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pagesize := ctx.QueryInt("pagesize")
	if pagesize <= 0 {
		pagesize = 36500
	}
	pageDateQueueDuration := getPageDateQueueTask(dayQueueTaskInfo, page, pagesize)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"totalCount":            count,
		"pageDateQueueDuration": pageDateQueueDuration,
	})
}

func GetComputeCenterTaskDetail(ctx *context.Context) {
	var dayComputeCenterTaskNum []models.DateCenterTaskStatistic
	var dayQueueTaskNum []models.DateQueueTaskStatistic
	var count int
	var err error
	beginTime, endTime := getBeginAndEndTime(ctx)
	queryType := ctx.QueryTrim("type")
	if queryType != "" {
		if queryType == "all" {
			dayComputeCenterTaskNum, dayQueueTaskNum, count, err = getMonthComputeCenterTaskNum(beginTime, endTime)
			if err != nil {
				log.Error("Can not query getMonthComputeCenterTaskNum.", err)
				return
			}
		} else if queryType == "current_year" {
			dayComputeCenterTaskNum, dayQueueTaskNum, count, err = getMonthComputeCenterTaskNum(beginTime, endTime)
			if err != nil {
				log.Error("Can not query getMonthComputeCenterTaskNum.", err)
				return
			}
		} else {
			dayComputeCenterTaskNum, dayQueueTaskNum, count, err = getDayComputeCenterTaskNum(beginTime, endTime)
			if err != nil {
				log.Error("Can not query dayComputeCenterTaskNum.", err)
				return
			}
		}
	} else {
		dayComputeCenterTaskNum, dayQueueTaskNum, count, err = getDayComputeCenterTaskNum(beginTime, endTime)
		if err != nil {
			log.Error("Can not query dayComputeCenterTaskNum.", err)
			return
		}
	}
	// dayComputeCenterTaskNum, dayQueueTaskNum, count, err = getDayComputeCenterTaskNum(beginTime, endTime)
	// if err != nil {
	// 	log.Error("Can not query dayComputeCenterTaskNum.", err)
	// 	return
	// }
	page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pagesize := ctx.QueryInt("pagesize")
	if pagesize <= 0 {
		pagesize = 36500
	}
	pageDateComputeCenterTask := getPageDateCloudbrainTask(dayComputeCenterTaskNum, page, pagesize)
	pageDateQueueTask := getPageDateQueueTask(dayQueueTaskNum, page, pagesize)

	resourceQueues, err := models.GetCanUseCardInfo()
	if err != nil {
		log.Error("GetCanUseCardInfo err: %v", err)
	}
	allAiCenterCodeMap := make(map[string]string)
	allQueueMap := make(map[int64]*models.ResourceQueue)
	for _, resourceQueue := range resourceQueues {
		key := resourceQueue.AiCenterCode
		if _, ok := allAiCenterCodeMap[key]; !ok {
			if resourceQueue.AiCenterName != "" {
				allAiCenterCodeMap[key] = resourceQueue.AiCenterName
			} else {
				allAiCenterCodeMap[key] = resourceQueue.AiCenterCode
			}
		}

		key_id := resourceQueue.ID
		if _, ok := allQueueMap[key_id]; !ok {
			allQueueMap[key_id] = resourceQueue
		}
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"count":              count,
		"pageDateCenterTask": pageDateComputeCenterTask,
		"pageDateQueueTask":  pageDateQueueTask,
		"allAiCenterCodeMap": allAiCenterCodeMap,
		"allQueueMap":        allQueueMap,
	})
}

func GetDurationRateStatistic(ctx *context.Context) {
	beginTime, endTime := getBeginAndEndTime(ctx)
	OpenIDurationRate, C2NetDurationRate, totalUsageRate := getDurationStatistic(beginTime, endTime)

	// 检查 UsageRate 中的 NaN 值并替换为合适的默认值
	for key := range OpenIDurationRate.UsageRate {
		if math.IsNaN(OpenIDurationRate.UsageRate[key]) {
			OpenIDurationRate.UsageRate[key] = 0.0 // 或者其他默认值
		}
	}

	for key := range C2NetDurationRate.UsageRate {
		if math.IsNaN(C2NetDurationRate.UsageRate[key]) {
			C2NetDurationRate.UsageRate[key] = 0.0 // 或者其他默认值
		}
	}
	if math.IsNaN(totalUsageRate) {
		totalUsageRate = 0.0 // 或者其他默认值
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"openIDurationRate": OpenIDurationRate,
		"c2NetDurationRate": C2NetDurationRate,
		"totalUsageRate":    totalUsageRate,
	})

}

func GetQueueDurationUsageStat(ctx *context.Context) {
	beginTime, endTime := getBeginAndEndTime(ctx)
	queueTaskDurationStatistics, err := models.GetQueueTaskDurationStatistics(&models.DurationStatisticOptions{
		BeginTime: beginTime,
		EndTime:   endTime,
	})
	if err != nil {
		log.Error("GetQueueTaskDurationStatistics error:", err)
		return
	}

	var queueTaskNumStatistics []*models.QueueTaskNumStatistic
	_, QueueUsageDuration, QueueUsageTotalDuration, _ := getQueueTaskNumAndDuration(beginTime, endTime, queueTaskDurationStatistics, queueTaskNumStatistics)
	resourceQueues, err := models.GetCanUseCardInfo()
	if err != nil {
		log.Error("GetCanUseCardInfo err: %v", err)
	}
	allAiCenterCodeMap := make(map[string]string)
	allQueueMap := make(map[int64]*models.ResourceQueue)
	for _, resourceQueue := range resourceQueues {
		key := resourceQueue.AiCenterCode
		if _, ok := allAiCenterCodeMap[key]; !ok {
			allAiCenterCodeMap[key] = resourceQueue.AiCenterName
		}

		key_id := resourceQueue.ID
		if _, ok := allQueueMap[key_id]; !ok {
			allQueueMap[key_id] = resourceQueue
		}
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"queueUsageDuration":      QueueUsageDuration,
		"queueTotalUsageDuration": QueueUsageTotalDuration,
		// "queueInfo":          QueueInfo,
		"allQueueMap": allQueueMap,
	})
}

func CloudbrainDurationStatisticForTest(ctx *context.Context) {
	repo.CloudbrainDurationStatisticHour()
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": 0,
	})
}

func getBeginAndEndTime(ctx *context.Context) (time.Time, time.Time) {
	queryType := ctx.QueryTrim("type")
	now := time.Now()
	beginTimeStr := ctx.QueryTrim("beginTime")
	endTimeStr := ctx.QueryTrim("endTime")
	var brainRecordBeginTime time.Time

	var beginTime time.Time
	var endTime time.Time
	var err error
	if queryType != "" {
		if queryType == "all" {
			recordCloudbrainDuration, err := models.GetDurationRecordBeginTime()
			if err != nil {
				log.Error("Can not get GetDurationRecordBeginTime", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("repo.record_begintime_get_err"))
				return beginTime, endTime
			}
			if len(recordCloudbrainDuration) > 0 && err == nil {
				brainRecordBeginTime = recordCloudbrainDuration[0].DateTimeUnix.AsTime()
			} else {
				brainRecordBeginTime = now
			}

			beginTime = brainRecordBeginTime
			endTime = now
		} else if queryType == "today" {
			beginTime = now.AddDate(0, 0, 0)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = now

		} else if queryType == "yesterday" {
			beginTime = now.AddDate(0, 0, -1)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		} else if queryType == "last_7day" {
			beginTime = now.AddDate(0, 0, -7)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = now.AddDate(0, 0, 0)
			endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, now.Location())
		} else if queryType == "last_30day" {
			beginTime = now.AddDate(0, 0, -30)
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, now.Location())
			endTime = now.AddDate(0, 0, 0)
			endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, now.Location())
		} else if queryType == "current_month" {
			endTime = now
			beginTime = time.Date(endTime.Year(), endTime.Month(), 1, 0, 0, 0, 0, now.Location())

		} else if queryType == "current_year" {
			endTime = now
			beginTime = time.Date(endTime.Year(), 1, 1, 0, 0, 0, 0, now.Location())
		} else if queryType == "last_month" {
			lastMonthTime := now.AddDate(0, -1, 0)
			beginTime = time.Date(lastMonthTime.Year(), lastMonthTime.Month(), 1, 0, 0, 0, 0, now.Location())
			endTime = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		}

	} else {
		if beginTimeStr == "" || endTimeStr == "" {
			//如果查询类型和开始时间结束时间都未设置，按queryType=all处理
			recordCloudbrainDuration, err := models.GetDurationRecordBeginTime()
			if err != nil {
				log.Error("Can not get recordCloudbrain", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("repo.record_begintime_get_err"))
				return beginTime, endTime
			}
			if len(recordCloudbrainDuration) > 0 && err == nil {
				brainRecordBeginTime = recordCloudbrainDuration[0].DateTimeUnix.AsTime()
			} else {
				brainRecordBeginTime = now
			}
			beginTime = brainRecordBeginTime
			endTime = now
		} else {
			beginTime, err = time.ParseInLocation("2006-01-02", beginTimeStr, time.Local)
			if err != nil {
				log.Error("Can not ParseInLocation.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("ParseInLocation_get_error"))
				return beginTime, endTime
			}
			endTime, err = time.ParseInLocation("2006-01-02", endTimeStr, time.Local)
			if err != nil {
				log.Error("Can not ParseInLocation.", err)
				ctx.Error(http.StatusBadRequest, ctx.Tr("ParseInLocation_get_error"))
				return beginTime, endTime
			}
			if endTime.After(time.Now()) {
				endTime = time.Now()
			}
		}

	}
	return beginTime, endTime
}

func getAiCenterUsageDuration(beginTime time.Time, endTime time.Time, cloudbrainStatistics []*models.CloudbrainDurationStatistic) (int, int, float64) {
	totalDuration := int(0)
	usageDuration := int(0)
	usageRate := float64(0)

	for _, cloudbrainStatistic := range cloudbrainStatistics {
		if int64(cloudbrainStatistic.DateTimeUnix) >= beginTime.Unix() && int64(cloudbrainStatistic.DateTimeUnix) < endTime.Unix() {
			totalDuration += cloudbrainStatistic.CardsTotalDuration
			usageDuration += cloudbrainStatistic.CardsUseDuration
		}
	}
	if totalDuration == 0 || usageDuration == 0 {
		usageRate = 0
	} else {
		usageRate = float64(usageDuration) / float64(totalDuration)
	}

	return totalDuration, usageDuration, usageRate
}

func getComputeCenterTaskNumAndDuration(beginTime time.Time, endTime time.Time, cloudbrainStatistics []*models.CloudbrainDurationStatistic, cloudbrainNumStatistics []*models.CloudbrainTaskNumStatistic) (map[string]int, map[string]int) {
	CenterTaskNum := make(map[string]int)
	CenterDuration := make(map[string]int)

	for _, cloudbrainStatistic := range cloudbrainStatistics {
		if int64(cloudbrainStatistic.DateTimeUnix) >= beginTime.Unix() && int64(cloudbrainStatistic.DateTimeUnix) < endTime.Unix() {
			if _, ok := CenterDuration[cloudbrainStatistic.AiCenterCode]; !ok {
				if cloudbrainStatistic.CardsUseDuration != 0 {
					CenterDuration[cloudbrainStatistic.AiCenterCode] = cloudbrainStatistic.CardsUseDuration
				}
			} else {
				if cloudbrainStatistic.CardsUseDuration != 0 {
					CenterDuration[cloudbrainStatistic.AiCenterCode] += cloudbrainStatistic.CardsUseDuration
				}
			}

		}
	}

	for _, cloudbrainNumStatistic := range cloudbrainNumStatistics {
		if int64(cloudbrainNumStatistic.DateTimeUnix) >= beginTime.Unix() && int64(cloudbrainNumStatistic.DateTimeUnix) < endTime.Unix() {
			if _, ok := CenterTaskNum[cloudbrainNumStatistic.AiCenterCode]; !ok {
				if cloudbrainNumStatistic.TaskNum != 0 {
					CenterTaskNum[cloudbrainNumStatistic.AiCenterCode] = cloudbrainNumStatistic.TaskNum
				}
			} else {
				if cloudbrainNumStatistic.TaskNum != 0 {
					CenterTaskNum[cloudbrainNumStatistic.AiCenterCode] += cloudbrainNumStatistic.TaskNum
				}
			}
		}
	}

	return CenterTaskNum, CenterDuration
}

func getQueueTaskNumAndDuration(beginTime time.Time, endTime time.Time, queueTaskDurationStatistics []*models.QueueTaskDurationStatistic, queueTaskNumStatistics []*models.QueueTaskNumStatistic) (map[int64]int, map[int64]int, map[int64]int, map[int64]*models.ResourceQueue) {
	QueueTaskNum := make(map[int64]int)
	QueueDuration := make(map[int64]int)
	QueueTotalDuration := make(map[int64]int)
	QueueInfo := make(map[int64]*models.ResourceQueue)

	for _, queueTaskNumStat := range queueTaskNumStatistics {
		if int64(queueTaskNumStat.DateTimeUnix) >= beginTime.Unix() && int64(queueTaskNumStat.DateTimeUnix) < endTime.Unix() {
			if _, ok := QueueTaskNum[queueTaskNumStat.QueueID]; !ok {
				QueueTaskNum[queueTaskNumStat.QueueID] = queueTaskNumStat.TaskNum
				//获取资源池队列信息
				if _, ok := QueueInfo[queueTaskNumStat.QueueID]; !ok {
					resourceQueue, err := models.GetResourceQueueByID(queueTaskNumStat.QueueID)
					if err != nil {
						log.Error("GetResourceQueue error.%v", err)
					}
					QueueInfo[queueTaskNumStat.QueueID] = resourceQueue
				}
			} else {
				QueueTaskNum[queueTaskNumStat.QueueID] += queueTaskNumStat.TaskNum
			}
		}
	}

	for _, queueTaskDurationStat := range queueTaskDurationStatistics {
		if int64(queueTaskDurationStat.DateTimeUnix) >= beginTime.Unix() && int64(queueTaskDurationStat.DateTimeUnix) < endTime.Unix() {
			if _, ok := QueueDuration[queueTaskDurationStat.QueueID]; !ok {
				QueueDuration[queueTaskDurationStat.QueueID] = queueTaskDurationStat.CardsUseDuration
				QueueTotalDuration[queueTaskDurationStat.QueueID] = queueTaskDurationStat.CardsTotalDuration
				//获取资源池队列信息
				if _, ok := QueueInfo[queueTaskDurationStat.QueueID]; !ok {
					resourceQueue, err := models.GetResourceQueueByID(queueTaskDurationStat.QueueID)
					if err != nil {
						log.Error("GetResourceQueue error.%v", err)
					}
					QueueInfo[queueTaskDurationStat.QueueID] = resourceQueue
				}
			} else {
				QueueDuration[queueTaskDurationStat.QueueID] += queueTaskDurationStat.CardsUseDuration
				QueueTotalDuration[queueTaskDurationStat.QueueID] += queueTaskDurationStat.CardsTotalDuration
			}

		}
	}

	return QueueTaskNum, QueueDuration, QueueTotalDuration, QueueInfo
}

func getDurationStatistic(beginTime time.Time, endTime time.Time) (models.DurationRateStatistic, models.DurationRateStatistic, float64) {
	OpenITotalDuration := make(map[string]int)
	OpenIUsageDuration := make(map[string]int)
	OpenIUsageRate := make(map[string]float64)

	C2NetTotalDuration := make(map[string]int)
	C2NetUsageDuration := make(map[string]int)

	C2NetUsageDurationTable := make(map[string]int)
	OpenIDurationRate := models.DurationRateStatistic{}
	C2NetDurationRate := models.DurationRateStatistic{}
	cardDurationStatistics, err := models.CloudbrainDurationStatistics(&models.DurationStatisticOptions{
		BeginTime: beginTime,
		EndTime:   endTime,
	})
	if err != nil {
		log.Error("CloudbrainDurationStatistics error:", err)
		return OpenIDurationRate, C2NetDurationRate, 0
	}
	for _, cloudbrainStatistic := range cardDurationStatistics {
		aiCenterName := cloudbrainStatistic.AiCenterCode + "/" + repo.GetAiCenterNameByCode(cloudbrainStatistic.AiCenterCode, "zh-CN")
		if cloudbrainStatistic.Cluster == models.OpenICluster {
			if _, ok := OpenITotalDuration[aiCenterName]; !ok {
				OpenITotalDuration[aiCenterName] = cloudbrainStatistic.CardsTotalDuration
			} else {
				OpenITotalDuration[aiCenterName] += cloudbrainStatistic.CardsTotalDuration
			}
			if _, ok := OpenIUsageDuration[aiCenterName]; !ok {
				OpenIUsageDuration[aiCenterName] = cloudbrainStatistic.CardsUseDuration
			} else {
				OpenIUsageDuration[aiCenterName] += cloudbrainStatistic.CardsUseDuration
			}
		}
		if cloudbrainStatistic.Cluster == models.C2NetCluster {
			if _, ok := C2NetTotalDuration[aiCenterName]; !ok {
				C2NetTotalDuration[aiCenterName] = cloudbrainStatistic.CardsTotalDuration
			} else {
				C2NetTotalDuration[aiCenterName] += cloudbrainStatistic.CardsTotalDuration
			}
			if _, ok := C2NetUsageDuration[aiCenterName]; !ok {
				C2NetUsageDuration[aiCenterName] = cloudbrainStatistic.CardsUseDuration
			} else {
				C2NetUsageDuration[aiCenterName] += cloudbrainStatistic.CardsUseDuration
			}
		}
	}

	//增加算力中心使用情况表格
	for _, cloudbrainStatistic := range cardDurationStatistics {
		aiCenterAndCard := cloudbrainStatistic.AiCenterCode + "/" + repo.GetAiCenterNameByCode(cloudbrainStatistic.AiCenterCode, "zh-CN") + "/" + cloudbrainStatistic.AccCardType
		if cloudbrainStatistic.Cluster == models.C2NetCluster {
			if _, ok := C2NetUsageDurationTable[aiCenterAndCard]; !ok {
				C2NetUsageDurationTable[aiCenterAndCard] = cloudbrainStatistic.CardsUseDuration
			} else {
				C2NetUsageDurationTable[aiCenterAndCard] += cloudbrainStatistic.CardsUseDuration
			}
		}
	}

	ResourceAiCenterRes, err := models.GetResourceAiCenters()
	if err != nil {
		log.Error("Can not get ResourceAiCenterRes.", err)
		return OpenIDurationRate, C2NetDurationRate, 0
	}
	for _, v := range ResourceAiCenterRes {
		aiCenterName := v.AiCenterCode + "/" + repo.GetAiCenterNameByCode(v.AiCenterCode, "zh-CN")
		if cutString(v.AiCenterCode, 4) == cutString(models.AICenterOfCloudBrainOne, 4) {
			if _, ok := OpenIUsageDuration[aiCenterName]; !ok {
				OpenIUsageDuration[aiCenterName] = 0
			}
			if _, ok := OpenITotalDuration[aiCenterName]; !ok {
				OpenITotalDuration[aiCenterName] = 0
			}
		} else {
			if _, ok := C2NetUsageDuration[aiCenterName]; !ok {
				C2NetUsageDuration[aiCenterName] = 0
			}
		}
	}
	totalCanUse := float64(0)
	totalUse := float64(0)
	totalUsageRate := float64(0)
	for k, v := range OpenITotalDuration {
		for i, j := range OpenIUsageDuration {
			if k == i {
				OpenIUsageRate[k] = float64(j) / float64(v)
			}
		}
	}
	for _, v := range OpenITotalDuration {
		totalCanUse += float64(v)
	}
	for _, v := range OpenIUsageDuration {
		totalUse += float64(v)
	}
	if totalCanUse == 0 || totalUse == 0 {
		totalUsageRate = 0
	} else {
		totalUsageRate = totalUse / totalCanUse
	}
	delete(C2NetUsageDuration, "/")

	OpenIDurationRate.AiCenterTotalDurationStat = OpenITotalDuration
	OpenIDurationRate.AiCenterUsageDurationStat = OpenIUsageDuration
	OpenIDurationRate.UsageRate = OpenIUsageRate
	C2NetDurationRate.AiCenterTotalDurationStat = C2NetTotalDuration
	C2NetDurationRate.AiCenterUsageDurationStat = C2NetUsageDuration
	C2NetDurationRate.C2NetUsageDurationTable = C2NetUsageDurationTable

	return OpenIDurationRate, C2NetDurationRate, totalUsageRate
}

func cutString(str string, lens int) string {
	if len(str) < lens {
		return str
	}
	return str[:lens]
}

func getDayCloudbrainDuration(beginTime time.Time, endTime time.Time, aiCenterCode string) ([]models.DateUsageStatistic, int, error) {
	now := time.Now()
	endTimeTemp := time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, now.Location())
	if endTimeTemp.Equal(endTime) {
		endTimeTemp = endTimeTemp.AddDate(0, 0, -1)
	}
	cardDurationStatistics, err := models.CloudbrainDurationStatistics(&models.DurationStatisticOptions{
		BeginTime:    beginTime,
		EndTime:      endTime,
		AiCenterCode: aiCenterCode,
	})
	if err != nil {
		log.Error("CloudbrainDurationStatistics error:", err)
		return nil, 0, err
	}

	dayCloudbrainInfo := make([]models.DateUsageStatistic, 0)
	count := 0
	for beginTime.Before(endTimeTemp) || beginTime.Equal(endTimeTemp) {
		TotalDuration, UsageDuration, UsageRate := getAiCenterUsageDuration(endTimeTemp, endTime, cardDurationStatistics)
		dayCloudbrainInfo = append(dayCloudbrainInfo, models.DateUsageStatistic{
			Date:          endTimeTemp.Format("2006/01/02"),
			UsageDuration: UsageDuration,
			TotalDuration: TotalDuration,
			UsageRate:     UsageRate,
		})
		endTime = endTimeTemp
		endTimeTemp = endTimeTemp.AddDate(0, 0, -1)
		if endTimeTemp.Before(beginTime) && beginTime.Before(endTime) {
			endTimeTemp = beginTime
		}
		count += 1
	}
	return dayCloudbrainInfo, count, nil
}

func getDayQueueTaskDuration(beginTime time.Time, endTime time.Time, queueID int64) ([]models.DateQueueTaskStatistic, int, error) {
	now := time.Now()
	endTimeTemp := time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, now.Location())
	if endTimeTemp.Equal(endTime) {
		endTimeTemp = endTimeTemp.AddDate(0, 0, -1)
	}
	queueTaskDurationStatistics, err := models.GetQueueTaskDurationStatistics(&models.DurationStatisticOptions{
		BeginTime: beginTime,
		EndTime:   endTime,
		QueueID:   queueID,
	})
	if err != nil {
		log.Error("GetQueueTaskDurationStatistics error:", err)
		return nil, 0, err
	}
	var queueTaskNumStatistics []*models.QueueTaskNumStatistic
	dayQueueTaskInfo := make([]models.DateQueueTaskStatistic, 0)
	count := 0
	for beginTime.Before(endTimeTemp) || beginTime.Equal(endTimeTemp) {
		QueueTaskNum, QueueUsageDuration, QueueTotalDuration, QueueInfo := getQueueTaskNumAndDuration(endTimeTemp, endTime, queueTaskDurationStatistics, queueTaskNumStatistics)
		dayQueueTaskInfo = append(dayQueueTaskInfo, models.DateQueueTaskStatistic{
			Date:               endTimeTemp.Format("2006/01/02"),
			QueueTaskNum:       QueueTaskNum,
			QueueTotalDuration: QueueTotalDuration,
			QueueUsageDuration: QueueUsageDuration,
			QueueInfo:          QueueInfo,
		})
		endTime = endTimeTemp
		endTimeTemp = endTimeTemp.AddDate(0, 0, -1)
		if endTimeTemp.Before(beginTime) && beginTime.Before(endTime) {
			endTimeTemp = beginTime
		}
		count += 1
	}
	return dayQueueTaskInfo, count, nil
}

func getDayComputeCenterTaskNum(beginTime time.Time, endTime time.Time) ([]models.DateCenterTaskStatistic, []models.DateQueueTaskStatistic, int, error) {
	now := time.Now()
	endTimeTemp := time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, now.Location())
	if endTimeTemp.Equal(endTime) {
		endTimeTemp = endTimeTemp.AddDate(0, 0, -1)
	}
	cloudbrainDurationStatistics, err := models.CloudbrainDurationStatistics(&models.DurationStatisticOptions{
		BeginTime: beginTime,
		EndTime:   endTime,
	})
	if err != nil {
		log.Error("CloudbrainDurationStatistics error:", err)
		return nil, nil, 0, err
	}

	cloudbrainNumStatistics, err := models.CloudbrainTaskNumStatistics(&models.DurationStatisticOptions{
		BeginTime: beginTime,
		EndTime:   endTime,
	})
	if err != nil {
		log.Error("CloudbrainNumStatistics error:", err)
		return nil, nil, 0, err
	}

	queueTaskDurationStatistics, err := models.GetQueueTaskDurationStatistics(&models.DurationStatisticOptions{
		BeginTime: beginTime,
		EndTime:   endTime,
	})
	if err != nil {
		log.Error("GetQueueTaskDurationStatistics error:", err)
		return nil, nil, 0, err
	}

	queueTaskNumStatistics, err := models.GetQueueTaskNumStatistics(&models.DurationStatisticOptions{
		BeginTime: beginTime,
		EndTime:   endTime,
	})
	if err != nil {
		log.Error("queueTaskNumStatistics error:", err)
		return nil, nil, 0, err
	}

	dayCenterTaskInfo := make([]models.DateCenterTaskStatistic, 0)
	dayQueueTaskInfo := make([]models.DateQueueTaskStatistic, 0)
	count := 0
	for beginTime.Before(endTimeTemp) || beginTime.Equal(endTimeTemp) {
		CenterTaskNum, CenterUsageDuration := getComputeCenterTaskNumAndDuration(endTimeTemp, endTime, cloudbrainDurationStatistics, cloudbrainNumStatistics)
		QueueTaskNum, QueueUsageDuration, QueueTotalDuration, QueueInfo := getQueueTaskNumAndDuration(endTimeTemp, endTime, queueTaskDurationStatistics, queueTaskNumStatistics)
		dayCenterTaskInfo = append(dayCenterTaskInfo, models.DateCenterTaskStatistic{
			Date:          endTimeTemp.Format("2006/01/02"),
			TaskNum:       CenterTaskNum,
			UsageDuration: CenterUsageDuration,
		})
		dayQueueTaskInfo = append(dayQueueTaskInfo, models.DateQueueTaskStatistic{
			Date:               endTimeTemp.Format("2006/01/02"),
			QueueTotalDuration: QueueTotalDuration,
			QueueTaskNum:       QueueTaskNum,
			QueueUsageDuration: QueueUsageDuration,
			QueueInfo:          QueueInfo,
		})
		endTime = endTimeTemp
		endTimeTemp = endTimeTemp.AddDate(0, 0, -1)
		if endTimeTemp.Before(beginTime) && beginTime.Before(endTime) {
			endTimeTemp = beginTime
		}
		count += 1
	}
	return dayCenterTaskInfo, dayQueueTaskInfo, count, nil
}

func getMonthComputeCenterTaskNum(beginTime time.Time, endTime time.Time) ([]models.DateCenterTaskStatistic, []models.DateQueueTaskStatistic, int, error) {
	now := time.Now()
	endTimeTemp := time.Date(endTime.Year(), endTime.Month(), 1, 0, 0, 0, 0, now.Location())
	if endTimeTemp.Equal(endTime) {
		endTimeTemp = endTimeTemp.AddDate(0, -1, 0)
	}
	cardDurationStatistics, err := models.CloudbrainDurationStatistics(&models.DurationStatisticOptions{
		BeginTime: beginTime,
		EndTime:   endTime,
	})
	if err != nil {
		log.Error("CloudbrainDurationStatistics error:", err)
		return nil, nil, 0, err
	}

	cloudbrainNumStatistics, err := models.CloudbrainTaskNumStatistics(&models.DurationStatisticOptions{
		BeginTime: beginTime,
		EndTime:   endTime,
	})
	if err != nil {
		log.Error("cloudbrainNumStatistics error:", err)
		return nil, nil, 0, err
	}

	queueTaskDurationStatistics, err := models.GetQueueTaskDurationStatistics(&models.DurationStatisticOptions{
		BeginTime: beginTime,
		EndTime:   endTime,
	})
	if err != nil {
		log.Error("GetQueueTaskDurationStatistics error:", err)
		return nil, nil, 0, err
	}

	queueTaskNumStatistics, err := models.GetQueueTaskNumStatistics(&models.DurationStatisticOptions{
		BeginTime: beginTime,
		EndTime:   endTime,
	})
	if err != nil {
		log.Error("queueTaskNumStatistics error:", err)
		return nil, nil, 0, err
	}

	dayCenterTaskInfo := make([]models.DateCenterTaskStatistic, 0)
	dayQueueTaskInfo := make([]models.DateQueueTaskStatistic, 0)
	count := 0
	for beginTime.Before(endTimeTemp) || beginTime.Equal(endTimeTemp) || (endTimeTemp.Before(beginTime) && beginTime.Before(endTime)) {
		CenterTaskNum, CenterUsageDuration := getComputeCenterTaskNumAndDuration(endTimeTemp, endTime, cardDurationStatistics, cloudbrainNumStatistics)
		QueueTaskNum, QueueUsageDuration, QueueTotalDuration, QueueInfo := getQueueTaskNumAndDuration(endTimeTemp, endTime, queueTaskDurationStatistics, queueTaskNumStatistics)
		dayCenterTaskInfo = append(dayCenterTaskInfo, models.DateCenterTaskStatistic{
			Date:          endTimeTemp.Format("2006/01"),
			TaskNum:       CenterTaskNum,
			UsageDuration: CenterUsageDuration,
		})
		dayQueueTaskInfo = append(dayQueueTaskInfo, models.DateQueueTaskStatistic{
			Date:               endTimeTemp.Format("2006/01"),
			QueueTaskNum:       QueueTaskNum,
			QueueTotalDuration: QueueTotalDuration,
			QueueUsageDuration: QueueUsageDuration,
			QueueInfo:          QueueInfo,
		})
		endTime = endTimeTemp
		endTimeTemp = endTimeTemp.AddDate(0, -1, 0)
		if endTimeTemp.Before(beginTime) && beginTime.Before(endTime) {
			endTimeTemp = beginTime
		}
		count += 1
	}
	return dayCenterTaskInfo, dayQueueTaskInfo, count, nil
}

func getHourCloudbrainDuration(beginTime time.Time, endTime time.Time, aiCenterCode string) (models.HourTimeStatistic, error) {
	hourTimeTotalDuration := make(map[string]int)
	hourTimeUsageDuration := make(map[string]int)
	hourTimeUsageRate := make(map[string]float64)
	hourTimeStatistic := models.HourTimeStatistic{}

	cardDurationStatistics, err := models.CloudbrainDurationStatistics(&models.DurationStatisticOptions{
		BeginTime: beginTime,
		EndTime:   endTime,
	})
	if err != nil {
		log.Error("CloudbrainDurationStatistics error:", err)
		return hourTimeStatistic, err
	}
	for _, cloudbrainStatistic := range cardDurationStatistics {
		if cloudbrainStatistic.AiCenterCode == aiCenterCode {
			if _, ok := hourTimeTotalDuration[strconv.Itoa(cloudbrainStatistic.HourTime)]; !ok {
				hourTimeTotalDuration[strconv.Itoa(cloudbrainStatistic.HourTime)] = cloudbrainStatistic.CardsTotalDuration
			} else {
				hourTimeTotalDuration[strconv.Itoa(cloudbrainStatistic.HourTime)] += cloudbrainStatistic.CardsTotalDuration
			}
			if _, ok := hourTimeUsageDuration[strconv.Itoa(cloudbrainStatistic.HourTime)]; !ok {
				hourTimeUsageDuration[strconv.Itoa(cloudbrainStatistic.HourTime)] = cloudbrainStatistic.CardsUseDuration
			} else {
				hourTimeUsageDuration[strconv.Itoa(cloudbrainStatistic.HourTime)] += cloudbrainStatistic.CardsUseDuration
			}
		}
	}
	hourTimeList := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"}
	for _, v := range hourTimeList {
		if _, ok := hourTimeUsageDuration[v]; !ok {
			hourTimeUsageDuration[v] = 0
		}
		if _, ok := hourTimeTotalDuration[v]; !ok {
			hourTimeTotalDuration[v] = 0
		}
	}

	for k, v := range hourTimeTotalDuration {
		for i, j := range hourTimeUsageDuration {
			if k == i {
				if v == 0 || j == 0 {
					hourTimeUsageRate[k] = 0
				} else {
					hourTimeUsageRate[k] = float64(j) / float64(v)
				}
			}
		}
	}

	hourTimeStatistic.HourTimeTotalDuration = hourTimeTotalDuration
	hourTimeStatistic.HourTimeUsageDuration = hourTimeUsageDuration
	hourTimeStatistic.HourTimeUsageRate = hourTimeUsageRate
	return hourTimeStatistic, nil
}

func CloudbrainUpdateAiCenter(ctx *context.Context) {
	repo.CloudbrainDurationStatisticHour()
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": 0,
	})
}

func GetResourceQueues(ctx *context.Context) {
	resourceQueues, err := models.GetCanUseCardInfo()
	if err != nil {
		log.Error("GetCanUseCardInfo err: %v", err)
		return
	}
	Resource := make([]*models.ResourceQueue, 0)
	aiCenterCodeMap := make(map[string]string)
	for _, resourceQueue := range resourceQueues {
		if _, ok := aiCenterCodeMap[resourceQueue.AiCenterCode]; !ok {
			// resourceQueue.AiCenterName = cloudbrainService.GetAiCenterShow(resourceQueue.AiCenterCode, ctx)
			aiCenterCodeMap[resourceQueue.AiCenterCode] = resourceQueue.AiCenterCode
			Resource = append(Resource, resourceQueue)
		}
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"resourceQueues": Resource,
	})
}

func GetComputeCenterTaskStatus(ctx *context.Context) {
	jobStatus := ctx.QueryTrim("jobStatus")
	computeResource := ctx.QueryTrim("computeResource")
	jobType := ctx.QueryTrim("jobType")

	ciTasks, err := models.GetQueueTask(&models.QueueCloudbrainOptions{
		ComputeResource: computeResource,
		JobTypes:        jobType,
		JobStatus:       jobStatus,
	})
	if err != nil {
		log.Error("Can not get ciTasks", err)
		ctx.Error(http.StatusBadRequest, ctx.Tr("repo.citask_get_err"))
		return
	}

	models.LoadSpecs4CloudbrainInfo(ciTasks)
	ComputeCenterTaskNum := make(map[string]map[string]map[string]map[string]int)
	QueueTaskNum := make(map[int64]map[string]map[string]map[string]int)
	QueueInfo := make(map[int64]*models.ResourceQueue)
	for _, taskInfo := range ciTasks {
		task := taskInfo.Cloudbrain
		task.AiCenter = repo.GetAiCenterNameByCode(cloudbrainService.GetAiCenterCode(task.AiCenter), ctx.Language())
		//按照计算中心
		if _, ok := ComputeCenterTaskNum[task.AiCenter]; !ok {
			ComputeCenterTaskNum[task.AiCenter] = make(map[string]map[string]map[string]int)
		}
		if _, ok := ComputeCenterTaskNum[task.AiCenter][task.Status]; !ok {
			ComputeCenterTaskNum[task.AiCenter][task.Status] = make(map[string]map[string]int)
		}
		if _, ok := ComputeCenterTaskNum[task.AiCenter][task.Status][task.ComputeResource]; !ok {
			ComputeCenterTaskNum[task.AiCenter][task.Status][task.ComputeResource] = make(map[string]int)
		}
		if _, ok := ComputeCenterTaskNum[task.AiCenter][task.Status][task.ComputeResource][task.JobType]; !ok {
			ComputeCenterTaskNum[task.AiCenter][task.Status][task.ComputeResource][task.JobType] = 1
		} else {
			ComputeCenterTaskNum[task.AiCenter][task.Status][task.ComputeResource][task.JobType] += 1
		}

		//按照资源池队列
		if _, ok := QueueTaskNum[task.Spec.QueueId]; !ok {
			QueueTaskNum[task.Spec.QueueId] = make(map[string]map[string]map[string]int)
		}
		if _, ok := QueueTaskNum[task.Spec.QueueId][task.Status]; !ok {
			QueueTaskNum[task.Spec.QueueId][task.Status] = make(map[string]map[string]int)
		}
		if _, ok := QueueTaskNum[task.Spec.QueueId][task.Status][task.ComputeResource]; !ok {
			QueueTaskNum[task.Spec.QueueId][task.Status][task.ComputeResource] = make(map[string]int)
		}
		if _, ok := QueueTaskNum[task.Spec.QueueId][task.Status][task.ComputeResource][task.JobType]; !ok {
			QueueTaskNum[task.Spec.QueueId][task.Status][task.ComputeResource][task.JobType] = 1
		} else {
			QueueTaskNum[task.Spec.QueueId][task.Status][task.ComputeResource][task.JobType] += 1
		}
		//获取资源池队列信息

		if _, ok := QueueInfo[task.Spec.QueueId]; !ok {
			resourceQueue, err := models.GetResourceQueueByID(task.Spec.QueueId)
			if err != nil {
				log.Error("GetResourceQueue error.%v", err)
			}
			QueueInfo[task.Spec.QueueId] = resourceQueue
		}

	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"computeCenterTaskNum": ComputeCenterTaskNum,
		"queueTaskNum":         QueueTaskNum,
		"queueInfo":            QueueInfo,
	})
}
