package ws

import (
	"github.com/gorilla/websocket"
)

// Client wraps one browser's WebSocket connection.
type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

func NewClient(hub *Hub, conn *websocket.Conn) *Client {
	return &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
}

// writePump waits for messages on this client's `send` channel and writes
// them out to the actual browser connection. Runs in its own goroutine.
func (c *Client) writePump() {
	defer c.conn.Close()
	for message := range c.send {
		if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
			return
		}
	}
}

// readPump listens for messages FROM the browser. We don't actually need
// the browser to send us anything for this feature, but this loop is
// still required — it's what detects the browser closing the tab/connection,
// so we can clean up and remove it from the hub.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		if _, _, err := c.conn.ReadMessage(); err != nil {
			break
		}
	}
}

// Start registers this client with the hub and kicks off both pumps.
func (c *Client) Start() {
	c.hub.register <- c
	go c.writePump()
	go c.readPump()
}