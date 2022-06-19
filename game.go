package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/thzoid/ws-game-server/shared"
)

var (
	matchMap = new(shared.Map)
	players  = make(map[uuid.UUID]shared.Player)
	conns    = make(map[uuid.UUID]*websocket.Conn)
)

// Spawn player into world
func SpawnPlayer(conn *websocket.Conn, profile shared.Profile) uuid.UUID {
	rand.Seed(time.Now().Unix())

	// Build player
	player := shared.Player{
		UserProfile: profile,
		Position:    shared.Coordinate{X: rand.Intn(matchMap.Size.X), Y: rand.Intn(matchMap.Size.Y)},
		// Position: shared.Coordinate{X: 0, Y: 0},
		Velocity: 1,
	}

	fmt.Println("player spawned:", player)

	// Spawn player
	id := uuid.New()
	players[id] = player
	conns[id] = conn
	return id
}

// Unspawn player from world
func UnspawnPlayer(id uuid.UUID) {
	delete(players, id)
	delete(conns, id)
}
