package redis_key

import "strings"

const KEY_SEPARATE = ":"

const EMPTY_REDIS_VAL = "Nil"

func KeyJoin(keys ...string) string {
	var build strings.Builder
	for _, v := range keys {
		build.WriteString(v)
		build.WriteString(KEY_SEPARATE)
	}
	s := build.String()
	s = strings.TrimSuffix(s, KEY_SEPARATE)
	return s
}
