package order

import (
	"context"
	"fmt"
	"service-order/constant"
	"service-order/dto"
	repoorder "service-order/module/order"
	"service-order/module/order/entity"
	repoproduct "service-order/module/product/repository"
	"service-order/pkg/kafka"
)

type orderService struct {
	OrderRepository repoorder.RepositoryInterface
	OrderRepoRPC    repoproduct.RepositoryGRPCInterface
	Kafka           *kafka.Producer
}

func NewOrderService(
	repository repoorder.RepositoryInterface,
	repoProduct repoproduct.RepositoryGRPCInterface,
	kafka *kafka.Producer,
) repoorder.OrderServiceInterface {
	return &orderService{
		OrderRepository: repository,
		OrderRepoRPC:    repoProduct,
		Kafka:           kafka,
	}

}

func (s *orderService) Create(input *dto.CreateOrderRequest) (*entities.Order, dto.ResponseError) {

	product, err := s.OrderRepoRPC.FindProductByIdRepository(context.Background(), input.ProductId)
	if product.ID < 1 {
		return nil, err
	}
	var order entities.Order
	order.ProductId = product.ID
	order.UserId = input.UserId
	order.Price = product.Price
	order.Quantity = input.Quantity
	order.Amount = product.Price * input.Quantity
	order.State = constant.ORDER_STATE_PENDING

	res, err := s.OrderRepository.Create(order)

	ayam := s.Kafka.SendMessage(constant.TOPIC_PRODUCT_STOCK_UPDATE, map[string]interface{}{
		"product_id": product.ID,
		"user_id":    input.UserId,
		"quantity":   input.Quantity,
		"order_id":   res.ID,
	}, 1)
	if ayam != nil {
		fmt.Println(ayam.Error())
	}

	return res, err
}

func (s *orderService) GetById(orderId uint64) (*entities.Order, dto.ResponseError) {

	res, err := s.OrderRepository.GetById(orderId)
	return res, err
}

func (s *orderService) GetList() ([]*entities.Order, dto.ResponseError) {

	res, err := s.OrderRepository.GetList()
	return res, err
}
