package model

// Define entity of a transaction
type Transaction struct {
	ID  int    `json:"id"`
	UserID int  `json:"userId"`
	FromWalletId string `json:"fromWalletId"`
	ToWalletId string`json:"toWalletId"`
	Amount float64 `json:"amount,omitempty"`
	CreatedAt int64 `json:"created_at,omitempty"`
}
