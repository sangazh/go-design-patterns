package main

import (
	"fmt"
	"sync"
	"time"
)

func main1() {
	channel := make(chan string)
	go func() {
		channel <- "Hello World"
		fmt.Println("Finishing goroutine")
	}()

	message := <-channel

	fmt.Println(message)

}

func main2() {
	channel := make(chan string)
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func() {
		channel <- "Hello World!"
		fmt.Println("Finishing goroutine")
		waitGroup.Done()
	}()

	time.Sleep(time.Second)
	message := <-channel
	fmt.Println(message)
	waitGroup.Wait()
}
