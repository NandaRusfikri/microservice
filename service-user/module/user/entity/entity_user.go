package entity

import (
	"gorm.io/gorm"
	"service-user/pkg"
	"time"
)

type Users struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        int64     `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	Fullname  string    `json:"fullname,omitempty" gorm:"type:varchar(255);unique;not null"`
	Email     string    `json:"email,omitempty" gorm:"type:varchar(255);unique;not null"`
	Password  string    `json:"password,omitempty" gorm:"type:varchar(255);not null"`
	IsActive  bool      `json:"is_active,omitempty" gorm:"type:bool;not null"`
	Balance   int64     `json:"balance,omitempty" gorm:"type:int;default:0"`
}

func (model *Users) TableName() string {
	return "users"
}

func (model *Users) BeforeCreate(db *gorm.DB) error {
	//model.ID = uuid.New().String()
	model.Password = pkg.HashPassword(model.Password)
	model.CreatedAt = time.Now().Local()
	return nil
}

func (model *Users) BeforeUpdate(db *gorm.DB) error {
	model.UpdatedAt = time.Now().Local()
	model.Password = pkg.HashPassword(model.Password)
	return nil
}
