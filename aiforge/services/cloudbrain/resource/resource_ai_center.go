package resource

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/grampus"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/redis/redis_lock"
	"code.gitea.io/gitea/modules/timeutil"
	"encoding/json"
	"sort"
	"strings"
	"time"
)

func GetActiveAICenterList(req models.ActiveAICenterReq) ([]models.AICenterDetailInfo, error) {
	allCenterInfo, err := GetAllActiveAICenterList()
	if err != nil {
		log.Error("GetAllActiveAICenterList err.%v", err)
		return nil, err
	}
	if req.ComputeSource == "" && req.AccCardType == "" {
		return allCenterInfo, nil
	}
	result := make([]models.AICenterDetailInfo, 0)
	for i := 0; i < len(allCenterInfo); i++ {
		if req.ComputeSource != "" {
			if req.ComputeSource == models.ILUVATAR || req.ComputeSource == models.METAX {
				req.ComputeSource = models.GPGPU
			}
			matched := false
			for _, v := range allCenterInfo[i].ComputeSourceList {
				if strings.ToUpper(v) == strings.ToUpper(req.ComputeSource) {
					matched = true
					break
				}
			}
			if !matched {
				continue
			}
		}
		if req.AccCardType != "" {
			matched := false
			for _, v := range allCenterInfo[i].AccCardTypeList {
				if strings.ToUpper(v) == strings.ToUpper(req.AccCardType) {
					matched = true
					break
				}
			}
			if !matched {
				continue
			}
		}
		result = append(result, allCenterInfo[i])
	}
	return result, nil
}

func GetAllActiveAICenterList() ([]models.AICenterDetailInfo, error) {
	list := make([]models.AICenterDetailInfo, 0)

	now := time.Now()
	everyDayRedisKey := redis_key.EveryDayAICenterList(now)
	val, _ := redis_client.Get(everyDayRedisKey)
	if val != "" {
		json.Unmarshal([]byte(val), &list)
		return list, nil
	}
	return loadAICenterRedis(now)

}

func loadAICenterRedis(now time.Time) ([]models.AICenterDetailInfo, error) {
	lock := redis_lock.NewDistributeLock(redis_key.UpdateEveryDayAICenterListLock(now))
	isOk, _ := lock.LockWithWait(5*time.Second, 5*time.Second)
	if !isOk {
		return []models.AICenterDetailInfo{}, nil
	}
	defer lock.UnLock()

	list := make([]models.AICenterDetailInfo, 0)
	everyDayRedisKey := redis_key.EveryDayAICenterList(now)
	val, _ := redis_client.Get(everyDayRedisKey)
	if val != "" {
		json.Unmarshal([]byte(val), &list)
		return list, nil
	}
	list, err := getAICenterList(now)
	if err != nil {
		log.Error("getAICenterList error,err=%v", err)
		return nil, err
	}
	jsonStr, _ := json.Marshal(list)
	redis_client.Setex(redis_key.EveryDayAICenterList(now), string(jsonStr), 24*time.Hour)
	return list, nil
}

func getAICenterList(now time.Time) ([]models.AICenterDetailInfo, error) {
	endTime := timeutil.GetMidNight(now)
	centerQueryStartTime := endTime.Add(-24 * 365 * time.Hour)
	//获取一年内有被使用的智算中心列表
	centerCodeList, err := models.GetActiveAICenterList(centerQueryStartTime, endTime)
	if err != nil {
		log.Error("GetActiveAICenterList err.%v", err)
		return nil, err
	}
	if len(centerCodeList) == 0 {
		log.Info("GetActiveAICenterList result empty")
		return []models.AICenterDetailInfo{}, nil
	}
	//获取一年内智算中心计算资源信息
	centerWithComputeResourceMap, err := models.GetAICenterComputeResourceMap(centerCodeList, centerQueryStartTime, endTime)
	if err != nil {
		log.Error("GetAICenterComputeResourceMap err.%v", err)
		return nil, err
	}
	//获取一年内智算中心卡类型信息
	centerWithCardTypeMap, err := models.GetAICenterCardTypeMap(centerCodeList, centerQueryStartTime, endTime)
	if err != nil {
		log.Error("GetAICenterComputeResourceMap err.%v", err)
		return nil, err
	}
	cardsNumQueryStartTime := endTime.Add(-24 * 7 * time.Hour)
	//获取七天内智算中心使用总卡时
	centerWithUsageMap, err := models.GetAICenterTotalCardsUsageMap(centerCodeList, cardsNumQueryStartTime, endTime)
	if err != nil {
		log.Error("GetAICenterTotalCardsUsage err.%v", err)
		return nil, err
	}
	//获取智算中心基础信息
	centerInfoMap, err := models.GetAICenterInfoMap(centerCodeList)
	if err != nil {
		log.Error("GetAICenterInfoMap err.%v", err)
		return nil, err
	}

	result := make([]models.AICenterDetailInfo, 0)
	//组装结果
	for k, v := range centerInfoMap {
		computeSourceList := centerWithComputeResourceMap[k]
		if computeSourceList != nil {
			sort.Strings(computeSourceList)
		}
		cardTypeList := centerWithCardTypeMap[k]
		if cardTypeList != nil {
			sort.Strings(cardTypeList)
		}
		result = append(result, models.AICenterDetailInfo{
			AICenterCode:       k,
			AICenterName:       v.AICenterName,
			City:               v.City,
			Province:           v.Province,
			AccessTime:         v.AccessTime,
			ComputeScale:       v.ComputeScale,
			TotalCardsUsageNum: centerWithUsageMap[k],
			ComputeSourceList:  computeSourceList,
			AccCardTypeList:    cardTypeList,
		})
	}
	//排序
	sort.Sort(models.AICenterDetailInfoList(result))
	return result, nil
}

func GetActiveAccCardInfo() ([]models.AccCardInfo, error) {
	list := make([]models.AccCardInfo, 0)
	now := time.Now()
	everyDayRedisKey := redis_key.EveryDayCardInfo(now)
	val, _ := redis_client.Get(everyDayRedisKey)
	if val != "" {
		json.Unmarshal([]byte(val), &list)
		return list, nil
	}

	lock := redis_lock.NewDistributeLock(redis_key.UpdateCardInfoLock(now))
	isOk, _ := lock.LockWithWait(5*time.Second, 5*time.Second)
	if !isOk {
		return []models.AccCardInfo{}, nil
	}
	defer lock.UnLock()
	val, _ = redis_client.Get(everyDayRedisKey)
	if val != "" {
		json.Unmarshal([]byte(val), &list)
		return list, nil
	}

	endTime := timeutil.GetMidNight(now)
	startTime := endTime.Add(-24 * 365 * time.Hour)
	list, err := models.GetActiveAccCardType(startTime, endTime)
	if err != nil {
		log.Error("GetActiveAccCardType err.%v", err)
		return nil, err
	}
	jsonStr, _ := json.Marshal(list)
	redis_client.Setex(everyDayRedisKey, string(jsonStr), 24*time.Hour)
	return list, nil

}

func SyncGrampusAICenter() error {
	r, err := grampus.GetResourceAICenter()
	if err != nil {
		log.Error("Get grampus AI Centers failed.err=%v", err)
		return err
	}
	log.Info("SyncGrampusAICenter result = %+v", r)
	centerUpdateList := make([]models.ResourceAICenter, 0)
	centerInsertList := make([]models.ResourceAICenter, 0)

	for _, center := range r {
		oldCenter, err := models.GetResourceAICenter(&models.ResourceAICenter{
			AICenterCode: center.Id,
		})
		if err != nil {
			log.Error("GetResourceAICenter err.%v", err)
			return err
		}
		city := center.City
		province := center.Province
		cityArray := strings.Split(city, "-")
		if len(cityArray) >= 2 {
			province = cityArray[0]
			city = cityArray[1]
		}
		if oldCenter == nil {
			centerInsertList = append(centerInsertList, models.ResourceAICenter{
				AICenterCode: center.Id,
				AICenterName: center.Name,
				City:         city,
				Province:     province,
				AccessTime:   timeutil.TimeStamp(center.AccessTime),
				ComputeScale: center.ComputeScale,
			})
		} else {
			centerUpdateList = append(centerUpdateList, models.ResourceAICenter{
				AICenterCode: center.Id,
				AICenterName: center.Name,
				City:         city,
				Province:     province,
				AccessTime:   timeutil.TimeStamp(center.AccessTime),
				ComputeScale: center.ComputeScale,
			})
		}

	}
	err = models.SyncGrampusAICenters(centerUpdateList, centerInsertList)
	if err == nil {
		delAICenterRelatedCache()
	}
	return err
}

func delAICenterRelatedCache() {
	redis_client.Del(redis_key.EveryDayAICenterList(time.Now()))
	redis_client.Del(redis_key.EveryDayCardInfo(time.Now()))
}
