package account

import (
	"bytes"
	"encoding/json"
	"strings"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/redis/redis_lock"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/util"
)

func GetAccount(userId int64) (*models.PointAccount, error) {
	redisKey := redis_key.PointAccountInfo(userId)
	val, _ := redis_client.Get(redisKey)
	if val != "" {
		account := &models.PointAccount{}
		json.Unmarshal([]byte(val), account)
		return account, nil
	}
	account, err := models.GetAccountByUserId(userId)
	if err != nil {
		if models.IsErrRecordNotExist(err) {
			a, err := InitAccount(userId)
			if err != nil {
				log.Error("InitAccount error,err=%v", err)
				return nil, err
			}
			return a, nil
		}
		log.Error("GetAccountByUserId error,err=%v", err)
		return nil, err
	}
	jsonStr, _ := json.Marshal(account)
	redis_client.Setex(redisKey, string(jsonStr), 24*time.Hour)
	return account, nil
}

func InitAccount(userId int64) (*models.PointAccount, error) {
	lock := redis_lock.NewDistributeLock(redis_key.PointAccountInitLock(userId))
	isOk, err := lock.LockWithWait(3*time.Second, 3*time.Second)
	if err != nil {
		log.Error("PointAccountInitLock error,err=%v", err)
		return nil, err
	}
	if isOk {
		defer lock.UnLock()
		account, _ := models.GetAccountByUserId(userId)
		if account == nil {
			models.InsertAccount(&models.PointAccount{
				Balance:       0,
				TotalEarned:   0,
				TotalConsumed: 0,
				UserId:        userId,
				Status:        models.PointAccountNormal,
				Version:       0,
				AccountCode:   util.UUID(),
			})
			return models.GetAccountByUserId(userId)
		}
		return account, nil
	}
	return nil, nil

}

//IsPointBalanceEnough check whether the user's point balance is enough to start the task
func IsPointBalanceEnough(targetUserId int64, c models.PointDeductCondition) bool {
	if !setting.CloudBrainPaySwitch {
		return true
	}
	a, err := GetAccount(targetUserId)
	if err != nil {
		log.Error("IsPointBalanceEnough GetAccount error,err=%v", err)
		return false
	}
	if c.SpecUnitPrice == 0 {
		return true
	}
	if c.WorkServerNumber <= 0 {
		c.WorkServerNumber = 1
	}

	return a.Balance >= c.SpecUnitPrice*float64(c.WorkServerNumber)
}

func SearchPointAccount(opt models.SearchPointAccountOpts) (*models.SearchPointAccountResponse, error) {
	var result = &models.SearchPointAccountResponse{
		Records:  make([]*models.UserPointAccount, 0),
		PageSize: opt.PageSize,
		Page:     opt.Page,
		Total:    0,
	}

	userSearch := &models.SearchUserOptions{
		Type: models.UserTypeIndividual,
		ListOptions: models.ListOptions{
			PageSize: 20,
		},
		SearchByEmail: true,
		OrderBy:       models.SearchOrderByAlphabetically,
	}

	userSearch.Page = opt.Page
	if userSearch.Page <= 0 {
		userSearch.Page = 1
	}
	userSearch.Keyword = strings.Trim(opt.Keyword, " ")
	if len(userSearch.Keyword) == 0 || isKeywordValid(userSearch.Keyword) {
		userSearch.OrderBy = ""
		users, count, err := models.SearchUsers(userSearch)
		if err != nil {
			log.Error("SearchPointAccount SearchUsers error.%v", err)
			return nil, err
		}
		userIds := make([]int64, 0)
		for _, v := range users {
			userIds = append(userIds, v.ID)
		}
		accountMap, err := models.GetPointAccountMapByUserIds(userIds)
		if err != nil {
			return nil, err
		}

		records := make([]*models.UserPointAccount, 0)
		for _, v := range users {
			upa := &models.UserPointAccount{
				UserId:        v.ID,
				UserName:      v.Name,
				Email:         v.Email,
				Balance:       0,
				TotalEarned:   0,
				TotalConsumed: 0,
			}
			a := accountMap[v.ID]
			if a != nil {
				upa.Balance = a.Balance
				upa.TotalConsumed = a.TotalConsumed
				upa.TotalEarned = a.TotalEarned
			}
			records = append(records, upa)
		}
		result.Records = records
		result.Total = count
	}
	return result, nil
}

func isKeywordValid(keyword string) bool {
	return !bytes.Contains([]byte(keyword), []byte{0x00})
}
