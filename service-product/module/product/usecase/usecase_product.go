package usecase

import (
	"service-product/constant"
	"service-product/dto"
	"service-product/module/product/entity"
	"service-product/module/product/repository"
	"service-product/pkg/kafka"
)

type UsecaseInterface interface {
	Create(input *dto.SchemaProduct) (*entity.Product, dto.ResponseError)
	GetByID(id uint64) (entity.Product, dto.ResponseError)
	GetList() ([]*entity.Product, dto.ResponseError)
	UpdateStock(input dto.UpdateStockRequest) (*entity.Product, dto.ResponseError)
}

type ServiceProduct struct {
	repository repository.ProductRepositoryInterface
	Kafka      *kafka.Producer
}

func NewServiceProduct(repository repository.ProductRepositoryInterface, kafka *kafka.Producer) *ServiceProduct {
	return &ServiceProduct{repository: repository, Kafka: kafka}
}

func (s *ServiceProduct) GetByID(id uint64) (entity.Product, dto.ResponseError) {

	res, err := s.repository.GetById(id)
	return res, err
}

func (s *ServiceProduct) Create(input *dto.SchemaProduct) (*entity.Product, dto.ResponseError) {

	var product dto.SchemaProduct
	product.Name = input.Name
	product.Quantity = input.Quantity
	product.IsActive = input.IsActive
	product.Price = input.Price

	res, err := s.repository.Create(&product)
	return res, err
}

func (s *ServiceProduct) GetList() ([]*entity.Product, dto.ResponseError) {

	res, err := s.repository.GetList()
	return res, err
}

func (s *ServiceProduct) UpdateStock(input dto.UpdateStockRequest) (*entity.Product, dto.ResponseError) {

	res, resErr := s.repository.UpdateStock(input)

	if resErr.Error != nil {
		return res, resErr
	} else {
		data := map[string]interface{}{
			"status":       true,
			"service_name": constant.SERVICE_NAME,
			"order_id":     input.OrderId,
		}

		err := s.Kafka.SendMessage(constant.TOPIC_ORDER_REPLY, data, 1)
		if err != nil {
			return nil, dto.ResponseError{
				Error: err,
			}
		}
	}

	return res, resErr
}
