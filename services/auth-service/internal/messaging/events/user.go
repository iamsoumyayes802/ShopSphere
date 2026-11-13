package events

import "time"

type UserCreatedEvent struct {
	UserID string `json:"user_id"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}	