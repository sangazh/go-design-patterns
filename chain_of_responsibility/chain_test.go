package chain_of_responsibility

import (
	"fmt"
	"strings"
	"testing"
)

func TestCreateDefaultChain(t *testing.T) {
	myWriter := new(myTestWriter)
	writerLogger := &WriterLogger{Writer:myWriter}
	second := &SecondLogger{NextChain:writerLogger}
	chain := FirstLogger{NextChain:second}

	t.Run("3 loggers", func(t *testing.T) {
		chain.Next("message that breaks the chain\n")

		if myWriter.receivedMessage != nil {
			t.Fatal("Last link should not receive any message")
		}

		chain.Next("hello\n")

		if !strings.Contains(*myWriter.receivedMessage, "hello") {
			t.Fatal("Last link didn't received expected message")
		}
	})

	t.Run("use closure", func(t *testing.T) {
		myWriter = new(myTestWriter)
		closureLogger := &ClosureChain{
			Closure: func(s string) {
				fmt.Printf("My closure logger! Message: %s\n", s)
				myWriter.receivedMessage = &s
			},
		}
		writerLogger.NextChain = closureLogger

		chain.Next("hello closure logger")
		if *myWriter.receivedMessage != "hello closure logger" {
			t.Fatal("not expect")
		}
	})

}
