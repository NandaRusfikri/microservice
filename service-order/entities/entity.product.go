package entities

import (
	"time"

	"gorm.io/gorm"
)

type EntityProduct struct {
	//ID        string    `json:"id" gorm:"primary_key;"`
	ID uint64 `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`

	Name      string         `json:"name,omitempty" gorm:"type:varchar(255);not null"`
	Quantity  int64          `json:"quantity,omitempty" gorm:"type:int;default:0"`
	IsActive  bool           `json:"is_active,omitempty" gorm:"type:bool;default:false"`
	Price     int64          `json:"price,omitempty" gorm:"column:price"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (model *EntityProduct) TableName() string {
	return "products"
}
func (model *EntityProduct) BeforeCreate(db *gorm.DB) error {
	//model.ID = uuid.New().String()
	model.CreatedAt = time.Now().Local()
	return nil
}

func (model *EntityProduct) BeforeUpdate(db *gorm.DB) error {
	model.UpdatedAt = time.Now().Local()
	return nil
}
