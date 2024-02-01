package order

import (
	"context"
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
) repoorder.ServiceInterface {
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

	data := map[string]interface{}{
		"product_id": product.ID,
		"user_id":    input.UserId,
		"quantity":   input.Quantity,
		"order_id":   res.ID,
		"amount":     order.Amount,
	}

	ayam := s.Kafka.SendMessage(constant.TOPIC_NEW_ORDER, data, 1)
	if ayam != nil {
		return nil, dto.ResponseError{
			Error: ayam,
		}
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

func (s *orderService) CreateOrderReply(input dto.CreateOrderReplyRequest) (bool, dto.ResponseError) {

	res, err := s.OrderRepository.CreateOrderReply(entities.TopicOrderReply{
		OrderId:     input.OrderId,
		ServiceName: input.ServiceName,
	})

	if res {
		s.OrderRepository.UpdateState(input.OrderId, constant.ORDER_STATE_SUCCESS)
	}
	return res, err
}
