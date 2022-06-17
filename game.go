package main

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/thzoid/ws-game-server/shared"
)

type controlledPlayer struct {
	shared.Player
	conn *websocket.Conn
}

var (
	matchMap = shared.Map{}
	players  = map[uuid.UUID]controlledPlayer{}
)

func SpawnPlayer(conn *websocket.Conn, player shared.Player) uuid.UUID {
	id := uuid.New()
	players[id] = controlledPlayer{Player: player, conn: conn}
	return id
}
