package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("client disconnected")
			return
		}

		fmt.Println("message received:", string(p))
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, _ := upgrader.Upgrade(w, r, nil)

	fmt.Println("client connected")
	reader(conn)
}

func listen(port string) {
	fmt.Println("server listening on", port)
	http.HandleFunc("/", wsEndpoint)
	log.Fatal(http.ListenAndServe(port, nil))
}
