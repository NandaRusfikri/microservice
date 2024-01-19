package entity

import (
	"service-product/constant"
	"time"

	"gorm.io/gorm"
)

type EntityProduct struct {
	//ID        string    `json:"id" gorm:"primary_key;"`
	ID uint64 `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`

	Name      string    `json:"name,omitempty" gorm:"type:varchar(255);not null"`
	Quantity  int64     `json:"quantity,omitempty" gorm:"type:int;default:0"`
	IsActive  bool      `gorm:"column:is_active;default:true" json:"is_active"`
	Price     int64     `json:"price,omitempty" gorm:"column:price"`
	CreatedAt time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (model *EntityProduct) TableName() string {
	return constant.TABLE_PRODUCT
}

func (model *EntityProduct) BeforeUpdate(db *gorm.DB) error {
	model.UpdatedAt = time.Now().Local()
	return nil
}
