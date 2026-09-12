package util

import (
	"sync/atomic"
)

type NonBlockingLock struct {
	state int32
}

func (m *NonBlockingLock) TryLock() bool {
	return atomic.CompareAndSwapInt32(&m.state, 0, 1)
}

func (m *NonBlockingLock) Unlock() {
	atomic.StoreInt32(&m.state, 0)
}
