package repositorys

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"service-order/dto"
	"service-order/module/order/entity"
	pb_product "service-order/proto/product"
	"service-order/utils"
	"strconv"
)

type RepositoryGRPCInterface interface {
	FindProductByIdRepository(ctx context.Context, ProductId uint64) (entities.EntityProduct, dto.ResponseError)
}

type orderRepositoryImplGRPC struct {
	//db *gorm.DB
}

func NewOrderRepositoryGRPC() RepositoryGRPCInterface {
	return &orderRepositoryImplGRPC{}
}

func (r *orderRepositoryImplGRPC) FindProductByIdRepository(ctx context.Context, product_id uint64) (entities.EntityProduct, dto.ResponseError) {
	ServiceProduct, _ := utils.CallConsulFindService("ServiceProduct")
	var Product entities.EntityProduct
	address := fmt.Sprintf("%v:%v", ServiceProduct.Address, ServiceProduct.ServicePort)
	log.Info("GRPC DIAl  : ", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", address, err)
	}

	product_client := pb_product.NewServiceProductRPCClient(conn)
	ResultProduct, err := product_client.GetById(context.Background(), &pb_product.GetByIdRequest{
		Id: strconv.FormatUint(product_id, 10),
	})
	if err != nil {
		log.Errorln("error Dial Product", err)
		return Product, dto.ResponseError{StatusCode: 500, Error: err}
	}
	ProductID, _ := strconv.ParseUint(ResultProduct.Id, 10, 64)
	fmt.Println("ProductID", ProductID)

	if ProductID > 0 {
		return entities.EntityProduct{
			ID:       ProductID,
			Name:     ResultProduct.Name,
			Quantity: ResultProduct.Quantity,
			Price:    ResultProduct.Price,
			IsActive: ResultProduct.IsActive,
		}, dto.ResponseError{}
	} else {
		return Product, dto.ResponseError{StatusCode: 404, Error: errors.New("product not found")}
	}

}
