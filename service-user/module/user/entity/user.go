package entity

import (
	"gorm.io/gorm"
	"service-user/constant"
	"time"
)

type Users struct {
	CreatedAt time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	ID        uint64    `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	Fullname  string    `gorm:"column:fullname" json:"fullname"`
	Email     string    `gorm:"column:email" json:"email"`
	Password  string    `gorm:"column:password" json:"password"`
	IsActive  *bool     `gorm:"column:is_active;default:true" json:"is_active"`
	Balance   int64     `gorm:"column:balance" json:"balance"`
}

func (entity *Users) TableName() string {
	return constant.TABLE_USERS
}

func (entity *Users) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Users) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
