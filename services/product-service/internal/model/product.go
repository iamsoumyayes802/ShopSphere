package model

import "time"

type Product struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updatedAt"`
}