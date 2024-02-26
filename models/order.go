package models

import "time"

type Order struct {
	ID        uint      `json:"id" db:"id"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Product   Product   `json:"product" db:"product"`
	User      User      `json:"user" db:"user"`
}
