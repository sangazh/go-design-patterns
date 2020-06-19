package visitor

import (
	"testing"
)

func Test_Overall(t *testing.T) {
	helper := new(TestHelper)
	visitor := new(MessageVisitor)

	t.Run("MessageA test", func(t *testing.T) {
		msg := MessageA{
			Msg:    "Hello World",
			Output: helper,
		}

		msg.Accept(visitor)
		msg.Print()

		expected := "A: Hello World (Visited A)"
		if helper.Received != expected {
			t.Errorf("expect: %s, got: %s", expected, helper.Received)
		}
	})
	t.Run("MessageB test", func(t *testing.T) {
		msg := MessageB{
			Msg:    "Hello World",
			Output: helper,
		}

		msg.Accept(visitor)
		msg.Print()

		expected := "B: Hello World (Visited B)"
		if helper.Received != expected {
			t.Errorf("expect: %s, got: %s", expected, helper.Received)
		}
	})
}
