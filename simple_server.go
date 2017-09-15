package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	arg := flag.String("arg", "Dipta", "A string")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		log.Println("Request:", r.Form)

		resp := *arg + ":\n"
		for k, v := range r.Form {
			resp += k + "=" + strings.Join(v, "") + "\n"
		}

		log.Println("Response:", resp)
		fmt.Fprintf(w, resp)
	})

	log.Println("Starting server: Port: 9090 Argument:", *arg)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
