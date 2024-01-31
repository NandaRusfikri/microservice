package entity

import (
	"service-product/constant"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint64    `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Stock     uint64    `gorm:"column:stock" json:"stock"`
	IsActive  bool      `gorm:"column:is_active;default:true" json:"is_active"`
	Price     uint64    `gorm:"column:price" json:"price"`
	CreatedAt time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (model *Product) TableName() string {
	return constant.TABLE_PRODUCT
}

func (model *Product) BeforeUpdate(db *gorm.DB) error {
	model.UpdatedAt = time.Now().Local()
	return nil
}
