package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	arg := flag.String("arg", "Dipta", "A string")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n=============Request=============\n%+v\n*********************************\n", r)
		r.ParseForm()
		resp := fmt.Sprintln("Arg:", *arg)
		resp += fmt.Sprintln("Sender:", r.RemoteAddr)
		resp += fmt.Sprintln("Host:", r.Host)
		resp += fmt.Sprintln("Form:", r.Form)
		fmt.Fprintf(w, resp)
	})

	log.Println("Starting server: Port: 9090 Argument:", *arg)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

/*
Dockerfile:
FROM ubuntu
ADD ./simple_server simple-server
EXPOSE 9090

Build:
$ docker build -t simple-server .

Run:
$ docker run -it simple-server ./simple-server -arg=appscode

Request:
$ curl -p 9090:9090 'localhost:9090?a=1&&b=1'
*/
