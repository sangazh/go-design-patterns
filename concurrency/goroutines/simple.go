package main

import (
	"fmt"
	"time"
)

func main1() {
	go func() {
		fmt.Println("Hello World")
	}()

	time.Sleep(time.Second)
}

func main2() {
	messagePrinter := func(msg string) {
		fmt.Println(msg)
	}
	go messagePrinter("Hello World")
	go messagePrinter("Hello goroutine")

	time.Sleep(time.Second)
}
