package entities

import (
	"service-order/constant"
	"time"
)

type TopicOrderReply struct {
	ID          int64     `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedAt   time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	OrderId     uint64    `gorm:"column:order_id" json:"order_id"`
	ServiceName string    `gorm:"column:service_name" json:"service_name"`
}

func (e *TopicOrderReply) TableName() string {
	return constant.TABLE_ORDER_REPLY
}
