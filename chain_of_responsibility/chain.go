package chain_of_responsibility

import (
	"fmt"
	"io"
	"strings"
)

type ChainLogger interface {
	Next(string)
}

type FirstLogger struct {
	NextChain ChainLogger
}

func (l *FirstLogger) Next(s string) {
	fmt.Printf("First Logger:%s \n", s)

	if l.NextChain != nil {
		l.NextChain.Next(s)
	}
}

type SecondLogger struct {
	NextChain ChainLogger
}

func (l *SecondLogger) Next(s string) {
	if !strings.Contains(s, "hello") {
		fmt.Println("Finishing in second logging")
		return
	}

	fmt.Printf("Second Logger:%s \n", s)

	if l.NextChain != nil {
		l.NextChain.Next(s)
	}
}

type WriterLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

func (l *WriterLogger) Next(s string) {
	if l.Writer != nil {
		l.Writer.Write([]byte("WriterLogger: "+ s))
	}

	if l.NextChain != nil {
		l.NextChain.Next(s)
	}

}

type myTestWriter struct {
	receivedMessage *string
}

func (m *myTestWriter) Write(p []byte) (int, error) {
	if m.receivedMessage == nil {
		m.receivedMessage = new(string)
	}
	tmpMsg := fmt.Sprintf("%s%s", *m.receivedMessage, p)
	m.receivedMessage = &tmpMsg
	return len(p), nil
}

func (m *myTestWriter) Next(s string) {
	m.Write([]byte(s))
}


type ClosureChain struct {
	NextChain ChainLogger
	Closure func(string)
}

func (c *ClosureChain) Next(s string) {
	if c.Closure!= nil {
		c.Closure(s)
	}

	if c.NextChain != nil {
		c.Next(s)
	}
}
