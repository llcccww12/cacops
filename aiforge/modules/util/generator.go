package util

import (
	"crypto/rand"
	"fmt"
	"io"
	"time"
)

// NewReqId 请求ID生成器
func NewReqId() string {
	timestamp := time.Now().UnixNano()

	// 生成随机数
	randomBytes := make([]byte, 4)
	io.ReadFull(rand.Reader, randomBytes)

	return fmt.Sprintf("%x%x", timestamp, randomBytes)
}
