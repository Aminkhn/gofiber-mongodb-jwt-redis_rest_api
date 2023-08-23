package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	UserID string             `json:"user_id" bson:"user_id"`
	//user      User    `gorm:"foreignKey:UserId"`
	//ProductID uint     `json:"product_id"`
	//Product   Product `gorm:"foreignKey:ProductId"`
	Products []Product `json:"products" bson:"products"`
}
