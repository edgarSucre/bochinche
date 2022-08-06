package api

type chatterRecord struct {
	room   string
	client *Client
}

type broadCastSignal struct {
	message []byte
	room    string
}

type Hub struct {
	clients map[string]map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan broadCastSignal

	// Register requests from the clients.
	register chan chatterRecord

	// Unregister requests from clients.
	unregister chan chatterRecord
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan broadCastSignal),
		register:   make(chan chatterRecord),
		unregister: make(chan chatterRecord),
		clients:    make(map[string]map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case reg := <-h.register:
			if _, ok := h.clients[reg.room]; !ok {
				h.clients[reg.room] = make(map[*Client]bool)
			}
			h.clients[reg.room][reg.client] = true
		case reg := <-h.unregister:
			if _, ok := h.clients[reg.room]; ok {
				delete(h.clients[reg.room], reg.client)
				close(reg.client.send)
			}
		case signal := <-h.broadcast:
			for client := range h.clients[signal.room] {
				select {
				case client.send <- signal.message:
				default:
					close(client.send)
					delete(h.clients[signal.room], client)
				}
			}
		}
	}
}
