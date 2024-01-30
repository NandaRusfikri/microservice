package repositorys

import (
	"service-order/dto"
	"service-order/module/order/entity"
)

type RepositoryInterface interface {
	Create(input *dto.SchemaOrder) (*entities.Order, dto.ResponseError)
	GetById(orderId uint64) (*entities.Order, dto.ResponseError)
	GetList() ([]*entities.Order, dto.ResponseError)
}

type OrderServiceInterface interface {
	Create(input *dto.SchemaOrder) (*entities.Order, dto.ResponseError)

	GetById(orderId uint64) (*entities.Order, dto.ResponseError)

	GetList() ([]*entities.Order, dto.ResponseError)
}
