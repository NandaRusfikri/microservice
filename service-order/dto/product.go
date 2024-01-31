package dto

import (
	"time"

	"gorm.io/gorm"
)

type ProductExternalResponse struct {
	ID        uint64
	Name      string
	Quantity  uint64
	IsActive  bool
	Price     uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
