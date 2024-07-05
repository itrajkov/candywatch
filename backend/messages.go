package backend

type Message struct {
	sender UserSession
	payload []byte
}
