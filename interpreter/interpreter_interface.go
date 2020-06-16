package interpreter

import (
	"strconv"
	"strings"
)

type Interpreter interface {
	Read() int
}

type value int

func (v *value) Read() int {
	return int(*v)
}

type operationSum struct {
	Left  Interpreter
	Right Interpreter
}

func (o *operationSum) Read() int {
	return o.Left.Read() + o.Right.Read()
}

type operationSub struct {
	Left  Interpreter
	Right Interpreter
}

func (o *operationSub) Read() int {
	return o.Left.Read() - o.Right.Read()
}

func operatorFactory(o string, left, right Interpreter) Interpreter {
	switch o {
	case SUM:
		return &operationSum{left, right}
	case SUB:
		return &operationSub{left, right}
	}
	return nil
}

type PolishNotationStack []Interpreter

func (p *PolishNotationStack) Push(s Interpreter) {
	*p = append(*p, s)
}

func (p *PolishNotationStack) Pop() Interpreter {
	if len(*p) > 0 {
		temp := (*p)[len(*p)-1]
		*p = (*p)[:len(*p)-1]
		return temp
	}
	return nil
}

func Cal(o string) (int, error) {
	s := PolishNotationStack{}
	operators := strings.Split(o, " ")

	for _, o := range operators {
		if isOperator(o) {
			right := s.Pop()
			left := s.Pop()
			f := operatorFactory(o, left, right)
			res := value(f.Read())
			s.Push(&res)
		} else {
			val, err := strconv.Atoi(o)
			if err != nil {
				panic(err)
			}
			temp := value(val)
			s.Push(&temp)
		}
	}
	return s.Pop().Read(), nil
}
