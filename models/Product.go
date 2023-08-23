package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt    time.Time
	Title        string `json:"title" bson:"title"`
	Description  string `json:"description" bson:"description"`
	Amount       int    `json:"amount" bson:"amount"`
	SerialNumber string `json:"serial_number" bson:"serial_number"`
	OrderID      uint
}
