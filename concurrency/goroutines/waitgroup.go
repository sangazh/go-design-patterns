package main

import (
	"fmt"
	"sync"
)

func main3() {
	var wait sync.WaitGroup
	wait.Add(1)
	go func() {
		fmt.Println("Hello World!")
		wait.Done()
	}()
	wait.Wait()
}

func main() {
	var wait sync.WaitGroup

	goRoutines := 5
	wait.Add(goRoutines)

	for i := 0; i < goRoutines; i ++ {
		go func(id int) {
			fmt.Printf("id: %d: Hello goroutines! \n", id)
			wait.Done()
		}(i)
	}
	wait.Wait()

}
