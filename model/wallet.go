package model

type Wallet struct {
	WalletId string `json:"walletId,omitempty"`
	UserId string `json:"userId,omitempty"`
	Balance int `json:"balance,omitempty"`
	UpdatedAt int64  `json:"updatedAt,omitempty"`
	CreatedAt int64  `json:"createdAt,omitempty"`
}