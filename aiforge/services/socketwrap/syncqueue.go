package socketwrap

import (
	"container/list"
	"sync"
)

type SyncQueue struct {
	Queue   *list.List
	Mutex   *sync.RWMutex
	MaxSize int
}

func (q *SyncQueue) Push(value interface{}) {
	q.Mutex.Lock()
	{
		if q.Queue.Len() < q.MaxSize {
			q.Queue.PushBack(value)
		} else {
			q.Queue.PushBack(value)
			q.Queue.Remove(q.Queue.Front())
		}
	}

	q.Mutex.Unlock()
}

func (q *SyncQueue) List() []interface{} {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()

	if q.Queue.Len() > 0 {
		list := make([]interface{}, q.Queue.Len())
		i := 0
		for e := q.Queue.Front(); e != nil; e = e.Next() {
			list[i] = e.Value
			i++
		}
		return list
	}
	return nil

}

func NewSyncQueue(maxSize int) *SyncQueue {
	return &SyncQueue{
		list.New(),
		&sync.RWMutex{},
		maxSize,
	}
}
