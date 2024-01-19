package usecase

import (
	"service-product/dto"
	"service-product/module/product/entity"
	"service-product/module/product/repository"
)

type ServiceProductInterface interface {
	Create(input *dto.SchemaProduct) (*entity.EntityProduct, dto.SchemaError)
	GetByID(id uint64) (entity.EntityProduct, dto.SchemaError)
	GetList() ([]*entity.EntityProduct, dto.SchemaError)
	Update(input *dto.SchemaProduct) (*entity.EntityProduct, dto.SchemaError)
}

type serviceProduct struct {
	repository repository.ProductRepositoryInterface
}

func NewServiceProduct(repository repository.ProductRepositoryInterface) *serviceProduct {
	return &serviceProduct{repository: repository}
}

func (s *serviceProduct) GetByID(id uint64) (entity.EntityProduct, dto.SchemaError) {

	res, err := s.repository.GetById(id)
	return res, err
}

func (s *serviceProduct) Create(input *dto.SchemaProduct) (*entity.EntityProduct, dto.SchemaError) {

	var product dto.SchemaProduct
	product.Name = input.Name
	product.Quantity = input.Quantity
	product.IsActive = input.IsActive
	product.Price = input.Price

	res, err := s.repository.Create(&product)
	return res, err
}

func (s *serviceProduct) GetList() ([]*entity.EntityProduct, dto.SchemaError) {

	res, err := s.repository.GetList()
	return res, err
}

func (s *serviceProduct) Update(input *dto.SchemaProduct) (*entity.EntityProduct, dto.SchemaError) {

	Product, err := s.repository.GetById(input.ID)
	if Product.ID < 1 {
		return nil, dto.SchemaError{StatusCode: 404, Error: err.Error}
	}

	var dataProduct dto.SchemaProduct
	dataProduct.ID = Product.ID
	dataProduct.Name = input.Name
	dataProduct.Quantity = input.Quantity
	dataProduct.IsActive = input.IsActive
	dataProduct.Price = input.Price

	res, err := s.repository.Update(&dataProduct)
	return res, err
}
