package interpreter

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	type testCase struct {
		input  string
		expect int
	}
	tests := []testCase{
		{"3 4 sum 2 sub", 5},
		//{"5 3 sub 8 mul 4 sum 5 div", 4},
	}
	for i, test := range tests {
		res, err := Cal(test.input)
		if err != nil {
			t.Error(err)
		}
		if res != test.expect {
			t.Errorf("%d expect %d got %d",i, test.expect, res)
		}
	}

}

