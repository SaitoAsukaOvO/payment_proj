package model

//Entity of World Wallet user
type User struct{
	ID int  `json:"id,omitempty"`
	FirstName string `json:"firstName"`
	LastName string  `json:"lastName"`
	Email string     `json:"email,omitempty"`
	Phone string      `json:"phone,omitempty"`
	CreatedAt int64 `json:"createdAt,omitempty"` // use unix timestamp
	UpdatedAt int64  `json:"updatedAt,omitempty"`
	DeletedAt int64   `json:"deletedAt,omitempty"`
	Wallet Wallet     `json:"wallet,omitempty"`
}

