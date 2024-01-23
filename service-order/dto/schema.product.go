package dto

import "time"

type SchemaOrder struct {
	ID        int64     `json:"id" `
	UserId    uint64    `json:"user_id" `
	ProductId uint64    `json:"product_id" `
	Amount    int64     `json:"amount" `
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
