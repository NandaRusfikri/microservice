package repositorys

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"service-order/entities"
	pb_product "service-order/proto/product"
	"service-order/schemas"
	"service-order/util"
	"strconv"
)

type orderRepositoryImplGRPC struct {
	//db *gorm.DB
}
func NewOrderRepositoryGRPC() OrderRepositoryGRPC {
	return &orderRepositoryImplGRPC{}
}

func (r *orderRepositoryImplGRPC) FindProductByIdRepository(ctx context.Context,product_id uint64) (entities.EntityProduct, schemas.SchemaDatabaseError) {
	ServiceProduct,_ := util.CallConsulFindService("service-product")
	var Product entities.EntityProduct
	address := fmt.Sprintf("%v:%v", ServiceProduct.Address, ServiceProduct.ServicePort)
	log.Info("GRPC DIAl : ",address)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", address, err)
	}

	product_client := pb_product.NewServiceProductHandlerClient(conn)
	ResultProduct,err := product_client.GetProductByIdRPC(context.Background(), &pb_product.ProductByIdRequest{
		Id: strconv.FormatUint(product_id,10),
	})
	if err != nil {
		log.Errorln("error Dial Product", err)
		return Product, schemas.SchemaDatabaseError{Code: 500, Message: "Internal Server Error"}
	}
	ProductID,_ := strconv.ParseUint(ResultProduct.Id,10,64)
	fmt.Println("ProductID",ProductID)

	if ProductID > 0 {
		return entities.EntityProduct{
			ID:       ProductID,
			Name:     ResultProduct.Name,
			Quantity: ResultProduct.Quantity,
			Price:    ResultProduct.Price,
			IsActive: ResultProduct.IsActive,
		}, schemas.SchemaDatabaseError{}
	}else {
		return Product, schemas.SchemaDatabaseError{Code: 404, Message: "Product Not Found"}
	}

}

