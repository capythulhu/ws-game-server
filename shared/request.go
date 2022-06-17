package shared

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Request struct {
	Type string
	Body []byte
}

func WriteRequest(conn *websocket.Conn, requestType string, requestBody interface{}) {
	body, _ := json.Marshal(requestBody)
	req := Request{
		Type: requestType,
		Body: body,
	}
	conn.WriteJSON(req)
}

type HandshakeRequest struct {
	Nick rune
}

type HandshakeResponse struct {
	MatchMap Map
}

type MoveRequest struct {
	Direction Coordinate
}

type ShootRequest struct{}
