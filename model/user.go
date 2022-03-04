package model

import "gorm.io/gorm"

//Entity of World Wallet user
type User struct{
	gorm.Model
	Name string `json:"name"`
	Email string     `json:"email"`
	Phone string      `json:"phone"`
	Password string `json:"password"`
	Wallets  []Wallet  `json:"wallets"`
	Transactions []Transaction `json:"transactions"`
}

