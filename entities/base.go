package entities

import "time"

// Base :
type Base struct {
	ID         int64      `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}
