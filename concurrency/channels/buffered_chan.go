package main

import (
	"fmt"
	"time"
)

func main3() {
	channel := make(chan string, 1)
	go func() {
		channel <- "Hello World!"
		fmt.Println("Finishing goroutine")
	}()

	time.Sleep(time.Second)

	message := <-channel
	fmt.Println(message)
}

func main4() {
	channel := make(chan string, 1)
	go func() {
		channel <- "Hello World! 1"
		channel <- "Hello World! 2"
		fmt.Println("Finishing goroutine")
	}()

	time.Sleep(time.Second)

	message := <-channel
	fmt.Println(message)
}
