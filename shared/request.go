package shared

type Request struct {
	Type string
	Body []byte
}

// CtS = Client to Server
// StC = Server to Client

type CtS_HandshakeRequest struct {
	Nick rune
}

type StC_HandshakeRequest struct {
	MapSize Coordinate
}
