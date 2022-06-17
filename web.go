package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/thzoid/ws-game-server/shared"
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

		r := &shared.Request{}
		json.Unmarshal(p, r)
		switch r.Type {
		case "handshake":
			// Read handshake from client
			hsC := &shared.CtS_HandshakeRequest{}
			json.Unmarshal(r.Body, hsC)
			fmt.Println("handshake received.", "client nick:", string(hsC.Nick))

			// Send handshake to client
			hsS := &shared.StC_HandshakeRequest{
				MapSize: mapSize,
			}
			shared.WriteRequest(conn, "handshake", hsS)
		default:
			fmt.Println("message received:", string(p))
		}
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
