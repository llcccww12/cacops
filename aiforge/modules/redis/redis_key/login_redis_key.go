package redis_key

const LOGIN_PREFIX = "login"

func LoginWechatKey(loginId string) string {
	return KeyJoin(LOGIN_PREFIX, loginId, "wechat")
}
