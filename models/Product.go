package models

import "time"

type Product struct {
	ID           string `json:"_id"`
	CreatedAt    time.Time
	Title        string `json:"title"`
	Description  string `json:"description"`
	Amount       int    `json:"amount"`
	SerialNumber string `json:"serial_number"`
	OrderID      uint
}
