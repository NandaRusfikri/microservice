package repositorys

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"service-order/constant"
	"service-order/dto"
	pb_product "service-order/proto/product"
	"service-order/utils"
	"strconv"
)

type RepositoryGRPCInterface interface {
	FindProductByIdRepository(ctx context.Context, ProductId uint64) (dto.ProductExternalResponse, dto.ResponseError)
}

type orderRepositoryGRPC struct {
	//db *gorm.DB
}

func NewOrderRepositoryGRPC() RepositoryGRPCInterface {
	return &orderRepositoryGRPC{}
}

func (r *orderRepositoryGRPC) FindProductByIdRepository(ctx context.Context, productId uint64) (dto.ProductExternalResponse, dto.ResponseError) {
	ServiceProduct, _ := utils.CallConsulFindService(constant.SERVICE_PRODUCT_NAME)
	var Product dto.ProductExternalResponse
	address := fmt.Sprintf("%v:%v", ServiceProduct.Address, ServiceProduct.ServicePort)
	log.Info("GRPC DIAl  : ", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", address, err)
	}

	product_client := pb_product.NewServiceProductRPCClient(conn)
	ResultProduct, err := product_client.GetById(context.Background(), &pb_product.GetByIdRequest{
		Id: strconv.FormatUint(productId, 10),
	})
	if err != nil {
		log.Errorln("error Dial Product", err)
		return Product, dto.ResponseError{StatusCode: 500, Error: err}
	}
	ProductID, _ := strconv.ParseUint(ResultProduct.Id, 10, 64)
	fmt.Println("ProductID", ProductID)

	if ProductID > 0 {
		return dto.ProductExternalResponse{
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
