package main

import (
	"fmt"
)

type One struct{}
type Two struct{}
type Three struct{}
type Four struct{}
type Five struct{}
type Six struct{}
type Seven struct{}
type Eight struct{}
type Nine struct{}
type Zero struct{}

func Sum(a, b interface{}) interface{} {
	switch a := a.(type) {
	case One:
		switch b := b.(type) {
		case One:
			return new(Two)
		case Two:
			return new(Three)
		case int:
			return b + 1
		default:
			return fmt.Errorf("number not found")
		}
	case Two:
		switch b.(type) {
		case One:
			return new(Three)
		default:
			return fmt.Errorf("number not found")
		}
	case int:
		switch b := b.(type) {
		case One:
			return new(Two)
		case int:
			return a + b
		default:
			return fmt.Errorf("number not found")
		}
	default:
		return fmt.Errorf("number not found")
	}
}

func main() {
	fmt.Printf("%#v\n", Sum(One{}, Two{}))
	fmt.Printf("%d\n", Sum(1, 2))
	fmt.Printf("%d\n", Sum(One{}, 2))
}
