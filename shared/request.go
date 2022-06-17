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

// CtS = Client to Server
// StC = Server to Client

type CtS_HandshakeRequest struct {
	Nick rune
}

type StC_HandshakeRequest struct {
	MatchMap Map
}
