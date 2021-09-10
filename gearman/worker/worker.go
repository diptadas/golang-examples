package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/appscode/g2/worker"
	"github.com/mikespook/golib/signal"
)

func ToUpper(job worker.Job) ([]byte, error) {
	log.Printf("ToUpper: Data=[%s]\n", job.Data())
	data := []byte(strings.ToUpper(string(job.Data())))
	time.Sleep(10 * time.Second)
	return data, nil
}

func main() {

	log.Println("Worker starting...")
	defer log.Println("Worker shutdown")

	w := worker.New(worker.OneByOne)
	defer w.Close()

	//error handling
	w.ErrorHandler = func(e error) {
		log.Println(e)
	}

	w.AddServer("tcp4", "127.0.0.1:4730")

	// Use worker.Unlimited (0) if you want no timeout
	w.AddFunc("ToUpper", ToUpper, worker.Unlimited)
	// This will give a timeout of 5 seconds
	w.AddFunc("ToUpperTimeOut5", ToUpper, 5)

	if err := w.Ready(); err != nil {
		log.Fatal(err)
		return
	}

	go w.Work()

	//Ctrl-C to exit
	signal.Bind(os.Interrupt, func() uint { return signal.BreakExit })
	signal.Wait()
}
