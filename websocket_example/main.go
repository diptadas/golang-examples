package main

import (
	"bufio"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"fmt"
	"io"
)

var (
	addr     = "127.0.0.1:8080"
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func wsReader(ws *websocket.Conn, writer io.Writer) {
	for {
		_, data, err := ws.ReadMessage()

		if err != nil {
			fmt.Println(err)
			break
		}
		writer.Write(data)
	}
}

func wsWriter(ws *websocket.Conn, reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		if err := ws.WriteMessage(websocket.TextMessage, scanner.Bytes()); err != nil {
			fmt.Println(err)
			return
		}
	}

	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			fmt.Println(err)
		}
		return
	}

	if err := ws.WriteMessage(websocket.TextMessage, []byte("websocket connected")); err != nil {
		fmt.Println(err)
		return
	}

	go wsWriter(ws, os.Stdin)
	go wsReader(ws, os.Stdout)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Println(err)
	}
}
