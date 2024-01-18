package entity

import (
	"service-user/constant"
	"time"
)

type Transaction struct {
	ID        uint64    `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	UserId    uint64    `gorm:"column:user_id" json:"user_id"`
	Amount    int64     `gorm:"column:amount" json:"amount"`
	Type      string    `gorm:"column:type" json:"type"`
}

func (entity *Transaction) TableName() string {
	return constant.TABLE_TRANSACTION
}
