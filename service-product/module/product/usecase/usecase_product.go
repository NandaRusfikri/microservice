package usecase

import (
	"service-product/dto"
	"service-product/module/product/entity"
	"service-product/module/product/repository"
)

type ServiceProductInterface interface {
	Create(input *dto.SchemaProduct) (*entity.Product, dto.SchemaError)
	GetByID(id uint64) (entity.Product, dto.SchemaError)
	GetList() ([]*entity.Product, dto.SchemaError)
	UpdateStock(input dto.UpdateStockRequest) (*entity.Product, dto.SchemaError)
}

type serviceProduct struct {
	repository repository.ProductRepositoryInterface
}

func NewServiceProduct(repository repository.ProductRepositoryInterface) *serviceProduct {
	return &serviceProduct{repository: repository}
}

func (s *serviceProduct) GetByID(id uint64) (entity.Product, dto.SchemaError) {

	res, err := s.repository.GetById(id)
	return res, err
}

func (s *serviceProduct) Create(input *dto.SchemaProduct) (*entity.Product, dto.SchemaError) {

	var product dto.SchemaProduct
	product.Name = input.Name
	product.Quantity = input.Quantity
	product.IsActive = input.IsActive
	product.Price = input.Price

	res, err := s.repository.Create(&product)
	return res, err
}

func (s *serviceProduct) GetList() ([]*entity.Product, dto.SchemaError) {

	res, err := s.repository.GetList()
	return res, err
}

func (s *serviceProduct) UpdateStock(input dto.UpdateStockRequest) (*entity.Product, dto.SchemaError) {

	res, err := s.repository.UpdateStock(input)
	return res, err
}
