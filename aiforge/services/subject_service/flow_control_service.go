package subject_service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"code.gitea.io/gitea/modules/labelmsg"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/services/role"
)

func CheckFlow(userID int64, fileName string, fileSize int64, subjectId string, subjectTye int) error {
	if !isUserUnderFlowControl(userID) {
		return nil
	}
	// 1. 检查全局并发数
	countGlobal24h, err := getGlobal24hFileCount(subjectTye)
	if err != nil {
		return err
	}
	if countGlobal24h >= int64(setting.FLOW_CONTROL.ALL_ATTACHEMENT_NUM_SDK) {
		return errors.New("The number of files uploaded using the SDK simultaneously cannot exceed " + fmt.Sprint(setting.FLOW_CONTROL.ALL_ATTACHEMENT_NUM_SDK))
	}

	// 2. 检查用户10分钟文件数
	count10m, err := getUser10mFileCount(userID, subjectTye, subjectId, fileName)
	if err != nil {
		rollbackGlobalConcurrent(subjectTye, subjectId, fileName)
		return err
	}
	if count10m >= int64(setting.FLOW_CONTROL.ATTACHEMENT_NUM_A_USER_LAST10M) {
		rollbackGlobalConcurrent(subjectTye, subjectId, fileName)
		return errors.New("A single user cannot upload more than " + fmt.Sprint(setting.FLOW_CONTROL.ATTACHEMENT_NUM_A_USER_LAST10M) + " files within the past 10 minutes.")
	}

	// 3. 检查用户24小时文件数
	count24h, err := getUser24hFileCount(userID, subjectTye, subjectId, fileName)
	if err != nil {
		rollbackGlobalConcurrent(subjectTye, subjectId, fileName)
		return err
	}
	if count24h >= int64(setting.FLOW_CONTROL.ATTACHEMENT_NUM_A_USER_LAST24HOUR) {
		rollbackGlobalConcurrent(subjectTye, subjectId, fileName)
		return errors.New("A single user cannot upload more than " + fmt.Sprint(setting.FLOW_CONTROL.ATTACHEMENT_NUM_A_USER_LAST24HOUR) + " files within the last 24 hours.")
	}

	// 4. 检查用户24小时上传总量
	totalSize, err := getUser24hTotalSize(userID, subjectTye)
	if err != nil {
		rollbackGlobalConcurrent(subjectTye, subjectId, fileName)
		return err
	}
	if totalSize+fileSize > setting.FLOW_CONTROL.ATTACHEMENT_SIZE_A_USER*1024*1024*1024 {
		rollbackGlobalConcurrent(subjectTye, subjectId, fileName)
		return errors.New("The total file size uploaded by a single user within the past 24 hours cannot exceed " + fmt.Sprint(setting.FLOW_CONTROL.ATTACHEMENT_SIZE_A_USER) + "G.")
	}

	return nil
}

func rollbackGlobalConcurrent(subjectType int, subjectId string, fileName ...string) {
	key := redis_key.GlobalFileNumLimitKey(subjectType)
	if len(fileName) == 0 {
		return
	}
	fileIdList := make([]string, 0, len(fileName))
	for _, f := range fileName {
		fileIdList = append(fileIdList, subjectId+"_"+f)
	}
	redis_client.BatchZRemSafe(key, fileIdList)
}

func getGlobal24hFileCount(subjectType int) (int64, error) {
	key := redis_key.GlobalFileNumLimitKey(subjectType)
	now := time.Now().Unix()
	start := now - 24*60*60

	// 清理窗口外的记录
	if err := redis_client.ZRemRangeByScore(key, 0, float64(start)); err != nil {
		return 0, err
	}

	// 获取当前数量
	count, err := redis_client.ZCard(key)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func getUser10mFileCount(userId int64, subjectType int, subjectId string, fileName string) (int64, error) {
	key := redis_key.SingleUser10mFileNumLimitKey(userId, subjectType)
	now := time.Now().Unix()
	start := now - 10*60

	// 清理窗口外的记录
	if err := redis_client.ZRemRangeByScore(key, 0, float64(start)); err != nil {
		return 0, err
	}

	// 获取当前数量
	count, err := redis_client.ZCard(key)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func getUser24hFileCount(userId int64, subjectType int, subjectId string, fileName string) (int64, error) {
	key := redis_key.SingleUser24hFileKey(userId, subjectType)
	now := time.Now().Unix()
	start := now - 24*60*60

	// 清理窗口外的记录
	if err := redis_client.ZRemRangeByScore(key, 0, float64(start)); err != nil {
		return 0, err
	}

	// 获取当前数量
	count, err := redis_client.ZCard(key)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func getUser24hTotalSize(userID int64, subjectType int) (int64, error) {
	now := time.Now().Unix()
	start := now - 24*60*60

	key := redis_key.SingleUser24hFileKey(userID, subjectType)

	// 清理24小时前的记录
	if err := redis_client.ZRemRangeByScore(key, 0, float64(start)); err != nil {
		return 0, err
	}

	// 获取所有文件大小并求和
	values, err := redis_client.ZRangeWithScoresStruct(key, 0, -1)
	if err != nil {
		return 0, err
	}

	var total int64
	for i := 0; i < len(values); i++ {

		s := values[i].Member
		if s == "" {
			continue
		}
		parts := strings.Split(s, ",")
		if len(parts) != 2 {
			continue
		}
		size, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			continue
		}
		total += size
	}

	return total, nil
}

func SetFlowCache(userId int64, subjectType int, subjectId string, fileName string, fileSize int64) error {
	r := labelmsg.Get()
	defer r.Close()

	now := float64(time.Now().Unix())

	keyGlobal24h := redis_key.GlobalFileNumLimitKey(subjectType)
	key10mCount := redis_key.SingleUser10mFileNumLimitKey(userId, subjectType)
	key24h := redis_key.SingleUser24hFileKey(userId, subjectType)

	fileId := subjectId + "_" + fileName

	r.Send("MULTI")
	r.Send("ZADD", key10mCount, now, fileId)
	r.Send("EXPIRE", key10mCount, 20*60)
	r.Send("ZADD", key24h, now, fmt.Sprintf("%d,%s", fileSize, fileId))
	r.Send("EXPIRE", key24h, 48*60*60)
	r.Send("ZADD", keyGlobal24h, now, fileId)

	_, err := r.Do("EXEC")
	if err != nil {
		return err
	}

	return nil
}

func UpdateFlowCacheAfterUploaded(subjectType int, subjectId string, fileName ...string) {
	defer func() {
		if err := recover(); err != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:%v", combinedErr)
		}
	}()
	rollbackGlobalConcurrent(subjectType, subjectId, fileName...)
}

//判断用户是否受流控策略影响
func isUserUnderFlowControl(userId int64) bool {
	if setting.FLOW_CONTROL.IGNORE_FLAG != "" && role.UserHasOper(userId, role.ROLE_OPER_IGNORE_FLOW_CONTROL) {
		return false
	}
	return true
}
