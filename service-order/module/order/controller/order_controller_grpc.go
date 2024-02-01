package handlers

import (
	"context"
	"fmt"
	"service-order/dto"
	orders "service-order/module/order"
	pb_order "service-order/proto/order"
	"strconv"
)

type OrderControllerGRPC struct {
	OrderService orders.ServiceInterface
	pb_order.UnimplementedServiceOrderRPCServer
}

func NewOrderControllerGRPC(serviceCreate orders.ServiceInterface) *OrderControllerGRPC {
	return &OrderControllerGRPC{OrderService: serviceCreate}
}

func (h *OrderControllerGRPC) Create(ctx context.Context, param *pb_order.CreateRequest) (*pb_order.Order, error) {

	var res pb_order.Order
	input := dto.CreateOrderRequest{
		UserId:    uint64(param.UserId),
		ProductId: uint64(param.ProductId),
		Quantity:  uint64(param.Quantity),
	}
	fmt.Println("input ", input)
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
