package entity

import (
	"service-product/constant"
	"time"
)

type Transaction struct {
	ID        uint64    `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	ProductId uint64    `gorm:"column:product_id" json:"product_id"`
	Price     uint64    `gorm:"column:price" json:"price"`
	Quantity  uint64    `gorm:"column:quantity" json:"quantity"`
	OrderId   uint64    `gorm:"column:order_id" json:"order_id"`
	Type      string    `gorm:"column:type" json:"type"`
}

func (entity *Transaction) TableName() string {
	return constant.TABLE_TRANSACTION
}
