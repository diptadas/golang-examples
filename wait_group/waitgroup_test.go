package wait_group

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	stopCh := make(chan struct{})
	go fn(stopCh, &wg)

	go func() {
		time.Sleep(5 * time.Second)
		close(stopCh)
	}()

	wg.Wait()
	t.Logf("end")
}

func fn(stopCh <-chan struct{}, wg *sync.WaitGroup) {
	log.Println("started")
	t := time.NewTicker(1 * time.Second)

loop:
	for i := 1; ; i++ {
		select {
		case <-stopCh:
			log.Println("signal stopCh")
			break loop
		case <-t.C:
			log.Println("main iteration", i)

		}
	}

	log.Println("stopped")
	wg.Done()
}
