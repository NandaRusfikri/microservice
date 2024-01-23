package order

import (
	"context"
	"fmt"
	"service-order/entities"
	repositorys "service-order/repositorys/order"
	"service-order/schemas"
)



type orderServiceImpl struct {
	OrderRepository repositorys.OrderRepository
	OrderRepoRPC repositorys.OrderRepositoryGRPC
}

func NewOrderService(repository repositorys.OrderRepository, repoRPC repositorys.OrderRepositoryGRPC) OrderService {
	return &orderServiceImpl{OrderRepository: repository , OrderRepoRPC: repoRPC}
}

func (s *orderServiceImpl) CreateOrderService(input *schemas.SchemaOrder) (*entities.EntityOrder, schemas.SchemaDatabaseError) {

	var product schemas.SchemaOrder

	Product,err := s.OrderRepoRPC.FindProductByIdRepository(context.Background(),input.ProductId)
	fmt.Println("disini")
	if Product.ID < 1 {
		err.Message = "Product Not Found"
		return nil, err
	}
	fmt.Println("bukan disini")
	//fmt.Printf("Product %+v\n",Product)

	product.ProductId = Product.ID
	product.UserId = input.UserId
	product.Amount = Product.Price

	res, err := s.OrderRepository.CreateOrderRepository(&product)
	return res, err
}

func (s *orderServiceImpl) DeleteOrderService(input *schemas.SchemaOrder) (*entities.EntityOrder, schemas.SchemaDatabaseError) {

	var student schemas.SchemaOrder
	student.ID = input.ID

	res, err := s.OrderRepository.DeleteOrderRepository(&student)
	return res, err
}

func (s *orderServiceImpl) ResultOrderService(input *schemas.SchemaOrder) (*entities.EntityOrder, schemas.SchemaDatabaseError) {

	var student schemas.SchemaOrder
	student.ID = input.ID

	res, err := s.OrderRepository.ResultOrderRepository(&student)
	return res, err
}

func (s *orderServiceImpl) FindAllOrderService() ([]*entities.EntityOrder, schemas.SchemaDatabaseError) {

	res, err := s.OrderRepository.FindAllOrderRepository()
	//fmt.Printf("serviceResults %+v\n",res)
	return res, err
}

func (s *orderServiceImpl) UpdateOrderService(input *schemas.SchemaOrder) (*entities.EntityOrder, schemas.SchemaDatabaseError) {

	var student schemas.SchemaOrder
	student.ID = input.ID
	student.ProductId = input.ProductId
	student.UserId = input.UserId
	student.Amount = input.Amount

	res, err := s.OrderRepository.UpdateOrderRepository(&student)
	return res, err
}


