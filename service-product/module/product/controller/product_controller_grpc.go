package controller

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"service-product/dto"
	"service-product/module/product/usecase"
	pb "service-product/proto/product"
	"strconv"
)

type ProductControllerGRPC struct {
	service usecase.ServiceProductInterface
	pb.UnimplementedServiceProductRPCServer
}

func NewControllerProductGRPC(service usecase.ServiceProductInterface) *ProductControllerGRPC {
	return &ProductControllerGRPC{
		service: service,
	}
}

func (controller *ProductControllerGRPC) Check(ctx context.Context, empty *empty.Empty) (*empty.Empty, error) {
	return empty, nil
}

func (controller *ProductControllerGRPC) GetById(ctx context.Context, param *pb.GetByIdRequest) (*pb.Product, error) {

	var res pb.Product
	id, ErrParse := strconv.ParseUint(param.Id, 10, 64)
	if ErrParse != nil {
		return nil, errors.New("Id must be number")
	}

	Product, errProduct := controller.service.GetByID(id)

	if errProduct.Error != nil {
		return nil, errProduct.Error
	} else {
		res.Id = strconv.Itoa(int(Product.ID))
		res.Name = Product.Name
		res.Quantity = Product.Stock
		res.IsActive = Product.IsActive
		res.Price = Product.Price
	}

	return &res, nil

}

func (controller *ProductControllerGRPC) Create(ctx context.Context, param *pb.CreateRequest) (*pb.Product, error) {

	var resp pb.Product
	input := dto.SchemaProduct{
		IsActive: param.IsActive,
		Name:     param.Name,
		Quantity: param.Quantity,
		Price:    param.Price,
	}
	Create, err := controller.service.Create(&input)

	if err.Error != nil {
		return nil, err.Error
	} else {
		resp.Id = strconv.Itoa(int(Create.ID))
		resp.Name = Create.Name
		resp.Quantity = Create.Stock
		resp.IsActive = Create.IsActive
		resp.Price = Create.Price
	}

	return &resp, nil

}

func (controller *ProductControllerGRPC) GetList(ctx context.Context, empty *empty.Empty) (*pb.GetListResponse, error) {

	var res pb.GetListResponse
	ListProduct, err := controller.service.GetList()

	if err.Error != nil {
		return nil, err.Error
	} else {
		var ListProto []*pb.Product
		for _, product := range ListProduct {
			data := pb.Product{
				Id:       strconv.FormatUint(product.ID, 10),
				Name:     product.Name,
				Quantity: product.Stock,
				Price:    product.Price,
				IsActive: product.IsActive,
			}
			ListProto = append(ListProto, &data)
		}
		res.List = ListProto
	}
	return &res, nil
}
