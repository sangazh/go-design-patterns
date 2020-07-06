package main

import (
	"fmt"
	"time"
)

func main6() {
	helloCh := make(chan string, 1)
	goodbyeCh := make(chan string, 1)
	quitCh := make(chan bool)
	go receiver(helloCh, goodbyeCh, quitCh)

	go sendString(helloCh, "hello!")

	time.Sleep(time.Second)

	go sendString(goodbyeCh, "goodbye!")

	<-quitCh

}

func sendString(ch chan<- string, s string) {
	ch <- s
}

func receiver(helloCh, goodbyeCh <-chan string, quitCh chan<- bool) {
	for {
		select {
		case msg := <-helloCh:
			fmt.Println(msg)
		case msg := <-goodbyeCh:
			fmt.Println(msg)
		case <-time.After(time.Second * 2):
			fmt.Println("Nothing received in 2 second.Exiting")
			quitCh <- true
			break
		}
	}
}
