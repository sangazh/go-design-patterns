package adapter

import (
	"fmt"
	"io"
	"strconv"
)

type Counter struct{}

func (f *Counter) Count(n uint64) uint64 {
	if n == 0 {
		fmt.Println(strconv.Itoa(0))
		return 0
	}

	cur := n
	fmt.Println(strconv.FormatUint(cur, 10))
	return f.Count(n - 1)
}

type CounterWriter struct {
	Writer io.Writer
}

func (f *CounterWriter) Count(n uint64) uint64 {
	if n == 0 {
		f.Writer.Write([]byte(strconv.Itoa(0) + "\n"))
		return 0
	}

	cur := n
	f.Writer.Write([]byte(strconv.FormatUint(cur, 10) + "\n"))
	return f.Count(n - 1)
}
