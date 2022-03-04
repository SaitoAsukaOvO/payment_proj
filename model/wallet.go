package model

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	WalletNo string `json:"walletNo"`
	UserID  uint
	Balance int `json:"balance"`
}