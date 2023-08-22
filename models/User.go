package models

import "time"

type User struct {
	ID        string `json:"_id"`
	Name      string `json:"name"`
	Family    string `json:"family"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt time.Time
	Orders    []Order `json:"orders"`
}
