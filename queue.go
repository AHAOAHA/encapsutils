package encapsutils

import (
	"sync"
)

// Queue queue interface.
type Queue interface {
	Push(v interface{}) bool
	Pop() interface{}
	Back() interface{}
	Front() interface{}
	Size() int
	Empty() bool
}

type queue struct {
	mtx *sync.RWMutex
	bus []interface{}
}

// NewQueue create a new queue.
func NewQueue() Queue {
	return &queue{
		mtx: &sync.RWMutex{},
		bus: make([]interface{}, 0),
	}
}

func (q *queue) Push(v interface{}) bool {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	q.bus = append([]interface{}{v}, q.bus...)
	return true
}

func (q *queue) Pop() interface{} {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	if len(q.bus) > 0 {
		rv := q.bus[0]
		q.bus = q.bus[1:]
		return rv
	}
	return nil
}

func (q *queue) Back() interface{} {
	q.mtx.RLock()
	defer q.mtx.RUnlock()
	if len(q.bus) > 0 {
		return q.bus[len(q.bus)-1]
	}
	return nil
}

func (q *queue) Front() interface{} {
	q.mtx.RLock()
	defer q.mtx.RUnlock()
	if len(q.bus) > 0 {
		return q.bus[0]
	}
	return nil
}

func (q *queue) Size() int {
	q.mtx.RLock()
	defer q.mtx.RUnlock()
	return len(q.bus)
}

func (q *queue) Empty() bool {
	q.mtx.RLock()
	defer q.mtx.RUnlock()
	return q.Size() == 0
}
