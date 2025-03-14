package main

import (
	"sync"
	"time"
)

func DoWork() {
	time.Sleep(1 * time.Second)
}

func main() {
	nChan := make(chan int)

	iterations := 1000

	go func() {
		wg := sync.WaitGroup{}
		defer close(nChan)
		for i := range iterations {
			wg.Add(1)
			go func() {
				defer wg.Done()
				DoWork()
				nChan <- i
			}()
		}
		wg.Wait()
	}()

	for n := range nChan {
		println(n)
	}

}
