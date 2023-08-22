package models

type Order struct {
	ID     string `json:"_id"`
	UserID string `json:"user_id"`
	//user      User    `gorm:"foreignKey:UserId"`
	//ProductID uint     `json:"product_id"`
	//Product   Product `gorm:"foreignKey:ProductId"`
	Products []Product
}
