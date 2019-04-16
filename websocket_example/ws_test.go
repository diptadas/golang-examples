package websocket

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
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
			log.Println(err)
			return
		}
		writer.Write(data)
	}
}

func wsWriter(ws *websocket.Conn, reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if err := ws.WriteMessage(websocket.TextMessage, scanner.Bytes()); err != nil {
			log.Println(scanner.Err())
			return
		}
	}
	if scanner.Err() != nil {
		log.Println(scanner.Err())
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving /ws")

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	if err := ws.WriteMessage(websocket.TextMessage, []byte("websocket connected")); err != nil {
		log.Println(err)
		return
	}
	go wsWriter(ws, strings.NewReader("hello world"))
	go wsReader(ws, os.Stdout)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving /")
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

func TestWebsocket(t *testing.T) {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
	if err := http.ListenAndServe(addr, nil); err != nil {
		t.Fatal(err)
	}
}
