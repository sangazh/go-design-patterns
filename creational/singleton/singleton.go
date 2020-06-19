package singleton

// not thread safe, cannot handle concurrency
type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = &singleton{count: 0}
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
