package interpreter

import (
	"strconv"
	"strings"
)

const (
	SUM = "sum"
	SUB = "sub"
	MUL = "mul"
	DIV = "div"
)

type polishNotationStack []int

func (p *polishNotationStack) Push(s int) {
	*p = append(*p, s)
}

func (p *polishNotationStack) Pop() int {
	if len(*p) > 0 {
		temp := (*p)[len(*p)-1]
		*p = (*p)[:len(*p)-1]
		return temp
	}
	return 0
}

func Calculate(o string) (int, error) {
	stack := polishNotationStack{}
	operators := strings.Split(o, " ")

	for _, o := range operators {
		if isOperator(o) {
			right := stack.Pop()
			left := stack.Pop()
			f := getOperationFunc(o)
			res := f(left, right)
			stack.Push(res)
		} else {
			val, err := strconv.Atoi(o)
			if err != nil {
				return 0, err
			}
			stack.Push(val)
		}
	}
	return stack.Pop(), nil
}

func isOperator(o string) bool {
	switch o {
	case SUM, SUB, MUL, DIV:
		return true
	}
	return false
}

func getOperationFunc(o string) func(a, b int) int {
	switch o {
	case SUM:
		return func(a, b int) int {
			return a + b
		}
	case SUB:
		return func(a, b int) int {
			return a - b
		}
	case MUL:
		return func(a, b int) int {
			return a * b
		}
	case DIV:
		return func(a, b int) int {
			return a / b
		}
	}
	return nil
}
