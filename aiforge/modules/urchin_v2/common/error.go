package common

import (
	"errors"
)

const (
	SuccessCode = 0
	ErrorSystem = 1
)

var ErrAbort = errors.New("AbortError")
