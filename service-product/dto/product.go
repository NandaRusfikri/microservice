package dto

import "time"

type SchemaProduct struct {
	ID        uint64    `json:"id" `
	Name      string    `json:"name" binding:"required,lowercase"`
	Quantity  uint64    `json:"quantity" binding:"required"`
	Price     uint64    `json:"price" binding:"required"`
	IsActive  bool      `json:"is_active" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateStockRequest struct {
	ProductId uint64
	Quantity  uint64
	OrderId   uint64
}
