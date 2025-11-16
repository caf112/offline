package internal

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}

	fmt.Println("🔥 New WebSocket connection!")

	client := &Client{Conn: conn, Send: make(chan []byte)}
	HubInstance.Register <- client

	go client.Read()
	go client.Write()
}

func (c *Client) Read() {
	defer func() {
		HubInstance.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("📩 Received: %s\n", msg)
		HubInstance.Broadcast <- msg
	}
}

func (c *Client) Write() {
	for msg := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
}
