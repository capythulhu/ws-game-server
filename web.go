package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/thzoid/ws-game-server/shared"
)

// Websocket upgrader
var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	hbRate = 100
)

// Output function
func heartbeat(conn *websocket.Conn) {
	for range time.Tick(time.Millisecond * time.Duration(hbRate)) {
		if err := shared.WriteMessage(conn, "heartbeat",
			shared.HeartbeatResponse{
				Players: players,
			},
		); err != nil {
			return
		}
	}
}

// Input function
func reader(conn *websocket.Conn) {
	var playerID uuid.UUID
	for {
		// Read message from client
		m, err := shared.ReadMessage(conn)
		if err != nil {
			fmt.Println("client disconnected")
			UnspawnPlayer(playerID)
			return
		}

		switch m.Type {
		case "handshake":
			// Read handshake from client
			hsC := &shared.HandshakeRequest{}
			json.Unmarshal(m.Body, hsC)

			// Send handshake to client
			hsS := &shared.HandshakeResponse{
				MatchMap: *matchMap,
			}
			shared.WriteMessage(conn, "handshake", hsS)

			// Spawn player in world
			playerID = SpawnPlayer(conn, hsC.UserProfile)

			// Start sending heartbeat
			go heartbeat(conn)
		default:
			fmt.Println("message received:", m)
		}
	}
}

func listen(port string) {
	fmt.Println("server listening on", port)
	// Handle websocket connection
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			upgrader.CheckOrigin = func(r *http.Request) bool { return true }
			conn, _ := upgrader.Upgrade(w, r, nil)

			fmt.Println("client connected")
			reader(conn)
		},
	)
	log.Fatal(http.ListenAndServe(port, nil))
}
