package repositorys

import (
	"service-order/dto"
	"service-order/module/order/entity"
)

type RepositoryInterface interface {
	Create(input entities.Order) (*entities.Order, dto.ResponseError)
	GetById(orderId uint64) (*entities.Order, dto.ResponseError)
	GetList() ([]*entities.Order, dto.ResponseError)
	CreateOrderReply(input entities.TopicOrderReply) (bool, dto.ResponseError)
	UpdateState(orderId uint64, state string) (*entities.Order, dto.ResponseError)
}

type ServiceInterface interface {
	Create(input *dto.CreateOrderRequest) (*entities.Order, dto.ResponseError)
	GetById(orderId uint64) (*entities.Order, dto.ResponseError)
	GetList() ([]*entities.Order, dto.ResponseError)
	CreateOrderReply(input dto.CreateOrderReplyRequest) (bool, dto.ResponseError)
}
