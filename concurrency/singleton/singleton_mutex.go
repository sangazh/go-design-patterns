package singleton

import (
	"sync"
)

type singleton2 struct {
	count int
	sync.RWMutex
}

var instance2 singleton2

func GetInstance2() *singleton2 {
	return &instance2
}

func (s *singleton2) AddOne() {
	s.Lock()
	defer s.Unlock()
	s.count++
}

func (s *singleton2) GetCount() int {
	s.RLock()
	defer s.RUnlock()
	return s.count
}
