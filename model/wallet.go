package model

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	UserID  uint `json:"userID"`
	Balance int `json:"balance"`
}