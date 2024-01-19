package entity

import (
	"service-product/constant"
	"time"
)

type Transaction struct {
	ID        uint64    `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	ProductId uint64    `gorm:"column:product_id" json:"product_id"`
	Amount    int64     `gorm:"column:amount" json:"amount"`
	Type      string    `gorm:"column:type" json:"type"`
}

func (entity *Transaction) TableName() string {
	return constant.TABLE_TRANSACTION
}
