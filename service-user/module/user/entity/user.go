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
	Email     string    `json:"email,omitempty" gorm:"type:varchar(255);unique;not null"`
	Password  string    `gorm:"column:password" json:"password"`
	IsActive  *bool     `json:"is_active,omitempty" gorm:"column:created_at;default:true"`
	Balance   int64     `json:"balance,omitempty" gorm:"column:balance"`
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
