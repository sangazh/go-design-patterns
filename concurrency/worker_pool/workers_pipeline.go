package main

import (
	"fmt"
	"log"
	"sync"
)

type Request struct {
	Data    interface{}
	Handler RequestHandler
}

type RequestHandler func(interface{})

func NewStringRequest(s string, id int, wg *sync.WaitGroup) Request {
	return Request{
		Data: fmt.Sprintf(s, id),
		Handler: func(i interface{}) {
			defer wg.Done()
			s, ok := i.(string)
			if !ok {
				log.Fatal("Invalid casting to string")
			}
			fmt.Println(s)
		},
	}
}

func main() {
	bufferSize := 100
	var dispatcher = NewDispatcher(bufferSize)
	workers := 3
	for i := 0; i < workers; i++ {
		var w WorkerLauncher = &PrefixSuffixWorker{
			id:      i,
			PrefixS: fmt.Sprintf("WorkerID: %d -> ", i),
			SuffixS: " World",
		}
		dispatcher.LaunchWorker(w)
	}

	requests := 10

	var wg sync.WaitGroup
	wg.Add(requests)

	for i := 0; i < requests; i++ {
		req := NewStringRequest("(Msg_id %d) -> Hello", i, &wg)
		dispatcher.MakeRequest(req)
	}

	dispatcher.Stop()
	wg.Wait()

}
