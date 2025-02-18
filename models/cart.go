package models

import "time"

// Cart represents the cart entity
type Cart struct {
	ID         int       `json:"id" db:"id"`
	UserID     int       `json:"user_id" db:"user_id"`
	ProductIDs []int     `json:"product_ids" db:"product_ids"`
	Quantity   []int     `json:"quantity" db:"quantity"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
