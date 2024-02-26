package models

import "time"

type Product struct {
	ID           uint      `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	Name         string    `json:"name" db:"name"`
	SerialNumber string    `json:"serial_number" db:"serial_number"`
}
