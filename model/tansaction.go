package model

// Define entity of a transaction
type Transaction struct {
	ID  int    `json:"id"`
	UserID int  `json:"userId"`
	FromId string `json:"fromId"`
	ToId string`json:"toId"`
	Amount float64 `json:"amount,omitempty"`
	TransKey string `json:"trans_key,omitempty"`
	CreatedAt int64 `json:"created_at,omitempty"`
}
