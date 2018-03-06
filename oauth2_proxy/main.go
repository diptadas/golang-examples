package main

import (
	"fmt"
	"net/http"
)

const (
	htmlIndex = `<html><body>
HELLO WORLD
</body></html>
`
)

func handleUpstrem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, htmlIndex)
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	url := "http://127.0.0.1:4180"
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func main() {
	fmt.Println("Starting server")
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/upstream", handleUpstrem)
	fmt.Println(http.ListenAndServe(":3000", nil))
}
