package model

import "gorm.io/gorm"

// Define entity of a transaction
type Transaction struct {
	gorm.Model
	Password string `json:"password"`
	UserID  uint
	FromWalletId string `json:"fromWalletId"`
	ToWalletId string`json:"toWalletId"`
	Amount float64 `json:"amount"`
}
