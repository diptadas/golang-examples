package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	server8080 := http.NewServeMux()
	server8080.HandleFunc("/", handler8080)
	fmt.Println("Starting server: Port: 8080")
	go http.ListenAndServe(":8080", server8080)

	server8081 := http.NewServeMux()
	server8081.HandleFunc("/foo", handlerFoo8081)
	server8081.HandleFunc("/bar", handlerBar8081)
	fmt.Println("Starting server: Port: 8081")
	go http.ListenAndServe(":8081", server8081)

	select {}
}

func handler8080(w http.ResponseWriter, r *http.Request) {
	fmt.Println("===============Request: 8080===============")
	fmt.Println(*r)

	req, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println("Error:", err)
	}
	if _, err = fmt.Fprintf(w, string(req)); err != nil {
		log.Println("Error:", err)
	}

	fmt.Println("*******************************************")
}

func handlerFoo8081(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=============Request: 8081/foo=============")
	fmt.Println(*r)

	req, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println("Error:", err)
	}
	if _, err = fmt.Fprintf(w, string(req)); err != nil {
		log.Println("Error:", err)
	}

	fmt.Println("*******************************************")
}

func handlerBar8081(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=============Request: 8081/bar=============")
	fmt.Println(*r)

	req, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println("Error:", err)
	}
	if _, err = fmt.Fprintf(w, string(req)); err != nil {
		log.Println("Error:", err)
	}

	fmt.Println("*******************************************")
}
