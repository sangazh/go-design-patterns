package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	sync.RWMutex
	value int
}

func main() {
	counter := new(Counter)
	for i := 0; i < 10; i++ {
		go func(i int) {
			counter.Lock()
			counter.value++
			fmt.Printf("%d counter.value %d\n", i, counter.value)
			defer counter.Unlock()
		}(i)
	}

	time.Sleep(time.Second)
	counter.Lock()
	defer counter.Unlock() // this lock and unlock used to lock the read(print)

	fmt.Println(counter.value) // print
}
