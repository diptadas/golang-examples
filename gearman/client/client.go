package main

import (
	"log"
	"os"

	"github.com/appscode/g2/client"
	"github.com/appscode/g2/pkg/runtime"
	"github.com/mikespook/golib/signal"
)

func main() {

	c, err := client.New("tcp4", "127.0.0.1:4730")
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	//error handling
	c.ErrorHandler = func(e error) {
		log.Println(e)
	}

	echo := []byte("Hello world")
	echoMsg, err := c.Echo(echo)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Echo:", string(echoMsg))

	jobHandler := func(resp *client.Response) {
		log.Printf("Response: %s", resp.Data)
	}

	c.Do("ToUpper", echo, runtime.JobNormal, jobHandler)
	c.Do("ToUpperTimeOut5", echo, runtime.JobNormal, jobHandler)

	//Ctrl-C to exit
	signal.Bind(os.Interrupt, func() uint { return signal.BreakExit })
	signal.Wait()
}
