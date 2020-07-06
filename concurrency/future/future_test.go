package future

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func TestStringOrError_Execute(t *testing.T) {
	future := new(MaybeString)
	t.Run("Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		//Timeout
		go timeout(t, &wg)
		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
			wg.Done()
		})

		future.Execute(func() (s string, e error) {
			return "Hello World!", nil
		})
		wg.Wait()
	})

	t.Run("Error result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		future.Success(func(s string) {
			t.Fail()
			wg.Done()
		}).Fail(func(e error) {
			t.Log(e.Error())
			wg.Done()
		})

		future.Execute(func() (s string, e error) {
			return "", errors.New("error")
		})
		wg.Wait()
	})

}

func timeout(t *testing.T, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	t.Log("Timeout!")

	t.Fail()
	wg.Done()
}
