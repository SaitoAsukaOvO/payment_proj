package model

import "gorm.io/gorm"

// Define entity of a transaction
type Transaction struct {
	gorm.Model
	Password string `json:"password"`
	UserID  uint `json:"userID"`
	ToUserID uint `json:"toUserId"`
	FromWalletId int `json:"fromWalletId"`
	ToWalletId int`json:"toWalletId"`
	Amount float64 `json:"amount"`
}
