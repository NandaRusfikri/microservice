package entities

import (
	"time"

	"gorm.io/gorm"
)

type EntityOrder struct {
	ID        int64      `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`

	UserId uint64 `json:"user_id" gorm:"column:user_id"`
	ProductId uint64 `json:"product_id" gorm:"column:product_id"`
	Amount int64 `json:"amount" gorm:"column:amount"`


	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (e *EntityOrder) TableName() string {
	return "orders"
}
func (e *EntityOrder) BeforeCreate(db *gorm.DB) error {
	//model.ID = uuid.New().String()
	e.CreatedAt = time.Now().Local()
	return nil
}

func (e *EntityOrder) BeforeUpdate(db *gorm.DB) error {
	e.UpdatedAt = time.Now().Local()
	return nil
}
