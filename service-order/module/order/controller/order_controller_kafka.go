package handlers

import (
	"context"
	"service-order/dto"
	orders "service-order/module/order"
	pb_order "service-order/proto/order"
	"strconv"
)

type OrderControllerKafka struct {
	OrderService orders.OrderServiceInterface
	pb_order.UnimplementedServiceOrderRPCServer
}

func NewOrderControllerKafka(serviceCreate orders.OrderServiceInterface) *OrderControllerGRPC {
	return &OrderControllerGRPC{OrderService: serviceCreate}
}

func (h *OrderControllerKafka) Create(ctx context.Context, param *pb_order.CreateRequest) (*pb_order.Order, error) {

	var res pb_order.Order
	input := dto.SchemaOrder{
		UserId:    uint64(param.UserId),
		ProductId: uint64(param.ProductId),
	}
	Create, err := h.OrderService.Create(&input)

	if err.Error != nil {

	} else {
		res.Id = strconv.Itoa(int(Create.ID))
		res.Amount = Create.Amount
		res.UserId = Create.UserId
		res.ProductId = Create.ProductId
	}

	return &res, nil

}
