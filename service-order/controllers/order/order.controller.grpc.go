package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	pb_order "service-order/proto/order"
	"service-order/schemas"
	services "service-order/services/order"
	"strconv"
)

type OrderControllerGRPC struct {
	OrderService services.OrderService
}

func NewHandlerRPCOrder(serviceCreate services.OrderService) *OrderControllerGRPC {
	return &OrderControllerGRPC{OrderService: serviceCreate}
}

func (h *OrderControllerGRPC) CreateOrderRPC(ctx context.Context, param *pb_order.CreateOrderRequest,res *pb_order.ModelProtoOrder) error {

	input:= schemas.SchemaOrder{
		UserId: uint64(param.UserId),
		ProductId: uint64(param.ProductId),
	}
	Create, err := h.OrderService.CreateOrderService(&input)
	fmt.Println("hadeh")
	fmt.Printf("Create %+v\n",Create)
	fmt.Printf("err %+v\n",err)

	switch err.Code {
	case 500:
		return errors.New("Internal Server Error")
	case 400:
		return errors.New("Name Order already exist")
	case 404:
		return errors.New(err.Message)
	case 402:
		return errors.New("Create new Order account failed")
	default:
		res.Id = strconv.Itoa(int(Create.ID))
		res.Amount = Create.Amount
		res.UserId = Create.UserId
		res.ProductId = Create.ProductId
	}
	return nil

}

func (h *OrderControllerGRPC) ListProductRPC(ctx context.Context,empty *empty.Empty, res *pb_order.ResponseEntityOrderList) error {

	//ListProduct, err := h.serviceResults.ResultsProductService()
	//
	//var ListProto  []*pb.EntityProtoProduct
	//
	//switch err.Code {
	//case 500:
	//	return errors.New("Internal Server Error")
	//default:
	//	for _, product := range ListProduct {
	//		data := pb.EntityProtoProduct{
	//			Id: strconv.FormatInt(product.ID, 10),
	//			Name: product.Name,
	//			Quantity: product.Quantity,
	//			Price: product.Price,
	//			IsActive: product.IsActive,
	//		}
	//		ListProto = append(ListProto, &data)
	//	}
	//
	//}
	//res.List = ListProto

	return nil
}




