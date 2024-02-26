package models

import "time"

type Order struct {
	ID        uint      `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	ProductID int       `json:"product_id" db:"product_id"`
	Product   Product   `json:"product" db:"product"`
	UserID    int       `json:"user_id" db:"user_id"`
	User      User      `json:"user" db:"user"`
}
