package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		t.Errorf("expected pointer to Singleton after Calling GetInstance not be nil")
	}
	expectedCounter := counter1

	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling for the first time, should be 1, but is %d\n", currentCount)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Error("Expected same instance")
	}

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("should be 2, but was %d \n", currentCount)
	}
}
