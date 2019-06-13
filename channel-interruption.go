package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	done := make(chan interface{}, 1)

	fmt.Println("NumCPU", runtime.NumCPU())
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0))

	go func() {
		defer wg.Done()
		i := 0
		for {
			select {
			case <-done:
				fmt.Println("gorutine 1 finished with:", i)
				return
			default:
				i++
			}
		}
	}()

	go func() {
		defer wg.Done()

		//heavy operation
		time.Sleep(2 * time.Second)

		fmt.Println("gorutine 2 finished")

		close(done)
	}()

	wg.Wait()
}
