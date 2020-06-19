package main

import (
	"fmt"
	"time"
)

type CommandInfo interface {
	Info() string
}

type TimePassed struct {
	start time.Time
}

func (t *TimePassed) Info() string{
	return time.Since(t.start).String()
}

type HelloMessage struct{}

func (h HelloMessage) Info() string{
	return "Hello world!"
}

type ChainLogger interface {
	Next(info CommandInfo)
}

type Logger struct {
	NextChain ChainLogger
}

func (f *Logger) Next(c CommandInfo) {
	time.Sleep(time.Second)

	fmt.Printf("Elaspsed time from creation: %s\n", c.Info())

	if f.NextChain != nil {
		f.NextChain.Next(c)
	}
}

func main1() {
	var timeCommand CommandInfo
	timeCommand = &TimePassed{start:time.Now()}

	var helloCommand CommandInfo
	helloCommand = new(HelloMessage)

	time.Sleep(time.Second)

	fmt.Println(timeCommand.Info())
	fmt.Println(helloCommand.Info())
}

func main() {
	second := new(Logger)
	first := Logger{NextChain:second}
	timeCommand := &TimePassed{start:time.Now()}
	first.Next(timeCommand)
}
