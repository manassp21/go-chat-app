package models

import(
	"time"
)

type Message struct{
	ID int `json:"id"`
	UserID int `json:"user_id"`
	UserName string `json:"username"`
	Content string `json:"content"`
	RoomID int `json:"room_id"`
	CreatedAt time.Time `json:"created_at"`
}