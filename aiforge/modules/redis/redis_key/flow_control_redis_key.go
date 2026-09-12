package redis_key

import "fmt"

const FLOW_CONTROL_REDIS_PREFIX = "flow_control"

//全网并发上传文件数
func GlobalFileNumLimitKey(subjectType int) string {
	return KeyJoin(FLOW_CONTROL_REDIS_PREFIX, "all", "24h", fmt.Sprint(subjectType))
}

//单用户上传文件列表
func SingleUser24hFileKey(userId int64, subjectType int) string {
	return KeyJoin(FLOW_CONTROL_REDIS_PREFIX, "24h", fmt.Sprint(subjectType), fmt.Sprint(userId))
}

//单用户10分钟上传文件id列表
func SingleUser10mFileNumLimitKey(userId int64, subjectType int) string {
	return KeyJoin(FLOW_CONTROL_REDIS_PREFIX, "10m", "num", fmt.Sprint(subjectType), fmt.Sprint(userId))
}
