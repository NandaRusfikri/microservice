package entity

import (
	"service-product/constant"
	"time"

	"gorm.io/gorm"
)

type EntityProduct struct {
	ID        uint64    `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Quantity  int64     `gorm:"column:quantity" json:"quantity"`
	IsActive  bool      `gorm:"column:is_active;default:true" json:"is_active"`
	Price     int64     `gorm:"column:price" json:"price"`
	CreatedAt time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (model *EntityProduct) TableName() string {
	return constant.TABLE_PRODUCT
}

func (model *EntityProduct) BeforeUpdate(db *gorm.DB) error {
	model.UpdatedAt = time.Now().Local()
	return nil
}
