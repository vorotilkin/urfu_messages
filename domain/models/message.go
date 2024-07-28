package models

import "time"

type Message struct {
	ID        int32
	UserID    int32
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
