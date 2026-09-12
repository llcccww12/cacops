package redis_key

const SIGN_UP_PREFIX = "sign_up"

func SignUpWechatKey(signUpId string) string {
	return KeyJoin(SIGN_UP_PREFIX, signUpId, "wechat")
}
