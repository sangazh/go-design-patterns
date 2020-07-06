package singleton

import (
	"fmt"
	"testing"
	"time"
)

func TestGetInstance(t *testing.T) {
	s1 := GetInstance()
	s2 := GetInstance()

	n := 5000

	for i := 0; i < n; i++ {
		go s1.AddOne()
		go s2.AddOne()
	}

	fmt.Printf("current count is: %d \n", s1.GetCount())

	var val int
	for val != n*2 {
		val = s1.GetCount()
		time.Sleep(10 * time.Millisecond)
	}
	s1.Stop()
}

func TestGetInstance2(t *testing.T) {
	s1 := GetInstance2()
	s2 := GetInstance2()

	n := 5000

	for i := 0; i < n; i++ {
		go s1.AddOne()
		go s2.AddOne()
	}

	fmt.Printf("current count is: %d \n", s1.GetCount())

	var val int
	for val != n*2 {
		val = s1.GetCount()
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Printf("final count is: %d \n", s1.GetCount())
}
