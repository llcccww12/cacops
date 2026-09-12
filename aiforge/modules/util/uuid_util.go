package util

import (
	gouuid "github.com/satori/go.uuid"
	"strings"
)

func UUID() string {
	return strings.ReplaceAll(gouuid.NewV4().String(), "-", "")
}
