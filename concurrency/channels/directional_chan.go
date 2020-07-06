package main

import (
	"fmt"
	"time"
)

func main5() {
	c := make(chan string, 1)
	go func(ch chan<- string) {
		ch <- "Hello World!"
		fmt.Println("Finishing goroutine")
	}(c)

	time.Sleep(time.Second)
	message := <-c
	fmt.Println(message)
}

func receivingChan(ch <-chan string) {
	msg := <-ch
	fmt.Println(msg)
}
