package models

import "time"

type Message struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
