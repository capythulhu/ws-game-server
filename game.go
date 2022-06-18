package main

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/thzoid/ws-game-server/shared"
)

var (
	matchMap = new(shared.Map)
	players  = make(map[uuid.UUID]shared.Player)
	conns    = make(map[uuid.UUID]*websocket.Conn)
)

func SpawnPlayer(conn *websocket.Conn, profile shared.Profile) uuid.UUID {
	// Build player
	player := shared.Player{
		UserProfile: profile,
		Position:    shared.Coordinate{X: 0, Y: 0},
		Velocity:    1,
	}

	// Spawn player
	id := uuid.New()
	players[id] = player
	conns[id] = conn
	return id
}
