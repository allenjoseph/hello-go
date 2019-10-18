package model

import "time"

// Woof db struct
type Woof struct {
	ID        string    `db:"id"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
}

// WoofResponse json struct
type WoofResponse struct {
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
}
