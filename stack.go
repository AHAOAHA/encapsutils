package encapsutils

import (
	"sync"
)

type Stack interface {
	Push(v interface{}) bool
	Top() interface{}
	Pop() interface{}
	Size() int
	Empty() bool
}

type stack struct {
	mtx *sync.RWMutex
	bus []interface{}
}

func NewStack() Stack {
	return &stack{
		mtx: &sync.RWMutex{},
		bus: make([]interface{}, 0),
	}
}

func (s *stack) Push(v interface{}) bool {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.bus = append(s.bus, v)
	return true
}

func (s *stack) Pop() interface{} {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if len(s.bus) > 0 {
		rv := s.bus[len(s.bus)-1]
		s.bus = s.bus[:len(s.bus)-1]
		return rv
	}
	return nil
}

func (s *stack) Top() interface{} {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	if len(s.bus) > 0 {
		return s.bus[len(s.bus)-1]
	}
	return nil
}

func (s *stack) Size() int {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return len(s.bus)
}

func (s *stack) Empty() bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.Size() == 0
}
