package repositorys

import (
	"service-order/dto"
	"service-order/module/order/entity"
)

type OrderRepositoryInterface interface {
	Create(input *dto.SchemaOrder) (*entities.EntityOrder, dto.ResponseError)
	GetById(input *dto.SchemaOrder) (*entities.EntityOrder, dto.ResponseError)
	GetList() ([]*entities.EntityOrder, dto.ResponseError)
}

type OrderServiceInterface interface {
	Create(input *dto.SchemaOrder) (*entities.EntityOrder, dto.ResponseError)

	GetById(input *dto.SchemaOrder) (*entities.EntityOrder, dto.ResponseError)

	GetList() ([]*entities.EntityOrder, dto.ResponseError)
}
