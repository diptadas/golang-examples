package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	stopCh := make(chan struct{})
	go fn(stopCh, &wg)

	go func() {
		time.Sleep(5 * time.Second)
		close(stopCh)
	}()

	wg.Wait()
	fmt.Println("end")
}

func fn(stopCh <-chan struct{}, wg *sync.WaitGroup) {
	fmt.Println("started")
	t := time.NewTicker(1 * time.Second)

loop:
	for i := 1; ; i++ {
		select {
		case <-stopCh:
			fmt.Println("signal stopCh")
			break loop
		case <-t.C:
			fmt.Println("main iteration", i)

		}
	}

	fmt.Println("stopped")
	wg.Done()
}
