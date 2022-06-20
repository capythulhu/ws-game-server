package shared

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Message struct {
	Type string
	Body []byte
}

func WriteMessage(conn *websocket.Conn, messageType string, messageBody interface{}) error {
	body, _ := json.Marshal(messageBody)
	req := Message{
		Type: messageType,
		Body: body,
	}
	return conn.WriteJSON(req)
}

func ReadMessage(conn *websocket.Conn) (*Message, error) {
	_, p, err := conn.ReadMessage()
	if err != nil {
		return nil, err
	}

	r := new(Message)
	json.Unmarshal(p, r)
	return r, nil
}

type HandshakeRequest struct {
	UserProfile Profile
}

type HandshakeResponse struct {
	PlayerID uuid.UUID
	MatchMap Map
}

type MoveRequest struct {
	Direction Coordinate
}

type ShootRequest struct{}

type HeartbeatResponse struct {
	Players map[uuid.UUID]Player
}
