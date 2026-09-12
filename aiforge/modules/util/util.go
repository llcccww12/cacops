// Copyright 2017 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// OptionalBool a boolean that can be "null"
type OptionalBool byte

const (
	// OptionalBoolNone a "null" boolean value
	OptionalBoolNone = iota
	// OptionalBoolTrue a "true" boolean value
	OptionalBoolTrue
	// OptionalBoolFalse a "false" boolean value
	OptionalBoolFalse
)

// IsTrue return true if equal to OptionalBoolTrue
func (o OptionalBool) IsTrue() bool {
	return o == OptionalBoolTrue
}

// IsFalse return true if equal to OptionalBoolFalse
func (o OptionalBool) IsFalse() bool {
	return o == OptionalBoolFalse
}

// IsNone return true if equal to OptionalBoolNone
func (o OptionalBool) IsNone() bool {
	return o == OptionalBoolNone
}

// OptionalBoolOf get the corresponding OptionalBool of a bool
func OptionalBoolOf(b bool) OptionalBool {
	if b {
		return OptionalBoolTrue
	}
	return OptionalBoolFalse
}

// Max max of two ints
func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Min min of two ints
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// IsEmptyString checks if the provided string is empty
func IsEmptyString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// NormalizeEOL will convert Windows (CRLF) and Mac (CR) EOLs to UNIX (LF)
func NormalizeEOL(input []byte) []byte {
	var right, left, pos int
	if right = bytes.IndexByte(input, '\r'); right == -1 {
		return input
	}
	length := len(input)
	tmp := make([]byte, length)

	// We know that left < length because otherwise right would be -1 from IndexByte.
	copy(tmp[pos:pos+right], input[left:left+right])
	pos += right
	tmp[pos] = '\n'
	left += right + 1
	pos++

	for left < length {
		if input[left] == '\n' {
			left++
		}

		right = bytes.IndexByte(input[left:], '\r')
		if right == -1 {
			copy(tmp[pos:], input[left:])
			pos += length - left
			break
		}
		copy(tmp[pos:pos+right], input[left:left+right])
		pos += right
		tmp[pos] = '\n'
		left += right + 1
		pos++
	}
	return tmp[:pos]
}

func AddZero(t int64) (m string) {
	if t < 10 {
		m = "0" + strconv.FormatInt(t, 10)
		return m
	} else {
		return strconv.FormatInt(t, 10)
	}
}

func ConvertDisplayJobNameToJobName(DisplayName string) (JobName string) {
	t := time.Now()
	JobName = strings.ToLower(cutNameString(DisplayName, 15)) + "t" + t.Format("20060102150405")[10:] + strconv.Itoa(int(rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000)))
	return JobName
}

func cutNameString(str string, lens int) string {
	if len(str) < lens {
		return str
	}
	return str[:lens]
}

func GetTotalPage(total int64, pageSize int) int {

	another := 0
	if int(total)%pageSize != 0 {
		another = 1
	}
	return int(total)/pageSize + another

}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

const (
	PUBULIC_CHUNK_SIZE = 200
)

// ChunkInt64 将 int64 切片分成指定大小的批次
func ChunkInt64(items []int64, chunkSize int) [][]int64 {
	if chunkSize <= 0 {
		chunkSize = 500 // 默认批次大小
	}

	var chunks [][]int64
	total := len(items)

	for start := 0; start < total; start += chunkSize {
		end := start + chunkSize
		if end > total {
			end = total
		}
		chunks = append(chunks, items[start:end])
	}

	return chunks
}

func JsonToString(unMarshalJson interface{}) string {
	toString, _ := json.Marshal(unMarshalJson)

	return string(toString)
}
