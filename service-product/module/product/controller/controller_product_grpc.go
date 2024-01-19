package controller

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"net/http"
	"service-product/dto"
	"service-product/module/product/usecase"
	pb "service-product/proto/product"
	"strconv"
)

type ProductControllerGRPC struct {
	service usecase.ServiceProductInterface
	pb.UnimplementedServiceProductRPCServer
}

func NewControllerProductRPC(service usecase.ServiceProductInterface) *ProductControllerGRPC {
	return &ProductControllerGRPC{
		service: service,
	}
}

func (controller *ProductControllerGRPC) Check(ctx context.Context, empty *empty.Empty) (*empty.Empty, error) {
	return empty, nil
}

func (controller *ProductControllerGRPC) GetById(ctx context.Context, param *pb.GetByIdRequest) (res *pb.Product, err error) {

	id, ErrParse := strconv.ParseUint(param.Id, 10, 64)
	if ErrParse != nil {
		return res, errors.New("Id must be number")
	}

	Product, errProduct := controller.service.GetByID(id)

	switch errProduct.StatusCode {
	case 500:
		return res, errors.New("Internal Server Error")
	case http.StatusNotFound:
		return res, nil
	default:
		res.Id = strconv.Itoa(int(Product.ID))
		res.Name = Product.Name
		res.Quantity = Product.Quantity
		res.IsActive = Product.IsActive
		res.Price = Product.Price
	}
	return res, nil

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

	switch err.StatusCode {
	case 500:
		return nil, errors.New("Internal Server Error")
	case 409:
		return nil, errors.New("Name Product already exist")
	case 402:
		return nil, errors.New("Create new Product account failed")
	default:
		resp.Id = strconv.Itoa(int(Create.ID))
		resp.Name = Create.Name
		resp.Quantity = Create.Quantity
		resp.IsActive = Create.IsActive
		resp.Price = Create.Price
	}
	return &resp, nil

}

func (controller *ProductControllerGRPC) GetList(ctx context.Context, empty *empty.Empty) (*pb.GetListResponse, error) {

	var res pb.GetListResponse
	ListProduct, err := controller.service.GetList()

	var ListProto []*pb.Product

	switch err.StatusCode {
	case 500:
		return nil, errors.New("Internal Server Error")
	default:
		for _, product := range ListProduct {
			data := pb.Product{
				Id:       strconv.FormatUint(product.ID, 10),
				Name:     product.Name,
				Quantity: product.Quantity,
				Price:    product.Price,
				IsActive: product.IsActive,
			}
			ListProto = append(ListProto, &data)
		}

	}
	res.List = ListProto

	return &res, nil
}

func (controller *ProductControllerGRPC) Update(ctx context.Context, req *pb.Product) (*pb.Product, error) {

	var input dto.SchemaProduct
	var resp pb.Product
	input.ID, _ = strconv.ParseUint(req.Id, 10, 64)
	input.Name = req.Name
	input.Price = req.Price
	input.Quantity = req.Quantity
	input.IsActive = req.IsActive

	Update, err := controller.service.Update(&input)

	switch err.StatusCode {
	case 500:
		return nil, errors.New("Internal Server Error")
	case 404:
		return nil, errors.New("Product data is not exist or deleted")
	case 409:
		return nil, errors.New("Update Product data failed")
	default:
		resp.Id = strconv.FormatUint(Update.ID, 10)
		resp.Name = Update.Name
		resp.Price = Update.Price
		resp.Quantity = Update.Quantity
		resp.IsActive = Update.IsActive
	}
	return &resp, nil

}
