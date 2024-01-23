package order

import (
	"context"
	"service-order/dto"
	orders "service-order/module/order"
	"service-order/module/order/entity"
	repoproduct "service-order/module/product/repository"
)

type orderService struct {
	OrderRepository orders.OrderRepositoryInterface
	OrderRepoRPC    repoproduct.ProductRepositoryGRPC
}

func NewOrderService(repository orders.OrderRepositoryInterface, repoProduct repoproduct.ProductRepositoryGRPC) orders.OrderServiceInterface {
	return &orderService{OrderRepository: repository, OrderRepoRPC: repoProduct}
}

func (s *orderService) Create(input *dto.SchemaOrder) (*entities.EntityOrder, dto.ResponseError) {

	var product dto.SchemaOrder

	Product, err := s.OrderRepoRPC.FindProductByIdRepository(context.Background(), input.ProductId)
	if Product.ID < 1 {
		return nil, err
	}

	product.ProductId = Product.ID
	product.UserId = input.UserId
	product.Amount = Product.Price

	res, err := s.OrderRepository.Create(&product)
	return res, err
}

func (s *orderService) GetById(input *dto.SchemaOrder) (*entities.EntityOrder, dto.ResponseError) {

	var student dto.SchemaOrder
	student.ID = input.ID

	res, err := s.OrderRepository.GetById(&student)
	return res, err
}

func (s *orderService) GetList() ([]*entities.EntityOrder, dto.ResponseError) {

	res, err := s.OrderRepository.GetList()
	return res, err
}
