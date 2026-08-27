// Package ws implements a simple WebSocket broadcast hub: every connected
// browser registers itself here, and any message sent to Broadcast gets
// pushed out to all of them.
package ws

// Hub tracks connected clients and fans out broadcast messages to all of
// them. There is exactly one Hub per running server — every websocket
// connection registers with the same instance.
type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	Broadcast  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		Broadcast:  make(chan []byte),
	}
}

// Run must be started once, in its own goroutine (see main.go). It loops
// forever, reacting to clients joining, leaving, or a message needing to
// go out to everyone.
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}

		case message := <-h.Broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					// The client's outgoing buffer is full/stuck — drop it
					// rather than let one slow client block everyone else.
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}