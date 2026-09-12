package otel

import (
	"bytes"
	"runtime"
	"strconv"
	"sync"
)

// 存储结构：key为goroutine ID + 自定义key，value为上下文
type storeKey struct {
	GID uint64 // goroutine ID
	Key string // 自定义的键名（如"trace_ctx"）
}

var (
	// 全局存储，绑定goroutine ID和数据
	goroutineStore sync.Map
)

// GlsSet 将数据存入当前goroutine的本地存储
// key：自定义键名（如"trace_ctx"）
// value：要存储的数据（如r.Context()）
func GlsSet(key string, value any) any {
	gid := GetGID()
	storeKey := storeKey{
		GID: gid,
		Key: key,
	}
	oldVal, ok := goroutineStore.Load(storeKey)
	goroutineStore.Store(storeKey, value)
	if ok {
		return oldVal
	} else {
		return nil
	}
}

// GlsGet 从当前goroutine的本地存储提取数据
func GlsGet(key string) (any, bool) {
	gid := GetGID()
	storeKey := storeKey{
		GID: gid,
		Key: key,
	}
	return goroutineStore.Load(storeKey)
}

// GlsDelete 清理当前goroutine本地存储的指定键数据
func GlsDelete(key string, oldVal any) {
	gid := GetGID()
	storeKey := storeKey{
		GID: gid,
		Key: key,
	}
	goroutineStore.Delete(storeKey)
	if oldVal != nil {
		goroutineStore.Store(storeKey, oldVal)
	}
}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
