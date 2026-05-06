package websocket 

import(
	"log"
	"time"
	"github.com/gorilla/websocket"
)

type Client struct{
	ID       int
	Username string
	Email    string
	RoomID   string
	Hub      *Hub
	Conn     *websocket.Conn
	Send     chan *WSMessage
	Done     chan bool
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		var msg WSMessage
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}

		msg.UserID = c.ID
		msg.Username = c.Username
		msg.RoomID = c.RoomID

		switch msg.Type {
		case MessageTypeChatMessage:
			c.Hub.Broadcast <- &msg

		case MessageTypeTyping:
			c.Hub.Broadcast <- &msg

		default:
			log.Printf("Unknown message type: %v", msg.Type)
		}
	}
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(54 * time.Second)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.Conn.WriteJSON(msg); err != nil {
				return
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}

		case <-c.Done:
			c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}
	}
}
