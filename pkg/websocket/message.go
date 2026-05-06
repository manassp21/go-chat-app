package websocket

import(
	"go-chat-app/pkg/models"
)

type MessageType string

const(
	MessageTypeChatMessage	MessageType = "message"
	MessageTypeUserJoined	MessageType = "user_joined"
	MessageTypeUserLeft		MessageType = "user_left"
	MessageTypeTyping		MessageType = "typing"
	MessageTypeError		MessageType = "error"
)

type WSMessage struct {
	Type      MessageType       `json:"type"`
	Content   string            `json:"content,omitempty"`
	Username  string            `json:"username,omitempty"`
	UserID    int               `json:"user_id,omitempty"`
	RoomID    string            `json:"room_id,omitempty"`
	Timestamp int64             `json:"timestamp,omitempty"`
	Message   *models.Message   `json:"message,omitempty"`
	Error     string            `json:"error,omitempty"`
}