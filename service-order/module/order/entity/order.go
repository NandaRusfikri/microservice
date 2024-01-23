package entities

import (
	"service-order/constant"
	"time"

	"gorm.io/gorm"
)

type EntityOrder struct {
	ID int64 `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`

	UserId    uint64 `gorm:"column:user_id" json:"user_id"`
	ProductId uint64 `gorm:"column:product_id" json:"product_id"`
	Amount    int64  `gorm:"column:amount" json:"amount"`

	CreatedAt time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (e *EntityOrder) TableName() string {
	return constant.TABLE_ORDER
}

func (e *EntityOrder) BeforeUpdate(db *gorm.DB) error {
	e.UpdatedAt = time.Now().Local()
	return nil
}
