package order

import (
	"fmt"
	repositorys "service-order/repositorys/order"
	"service-order/schemas"
	"service-order/util"
	"testing"
)

func TestCreateOrderService(t *testing.T) {

	db := util.SetupDatabase()

	orderRepository := repositorys.NewOrderRepositorySQL(db)
	orderRepositoryRPC := repositorys.NewOrderRepositoryGRPC()
	service := NewOrderService(orderRepository,orderRepositoryRPC)

	DataOrder := schemas.SchemaOrder{
		UserId: 1,
		ProductId: 12,
	}
	CreateOrder, _ := service.CreateOrderService(&DataOrder)
	fmt.Printf("CreateOrder %+v\n",CreateOrder)

	//assert.Equal(t, int64(1), int64(Product.ID))
	//assert.Equal(t, "OK", webResponse.Status)
	//
	//jsonData, _ := json.Marshal(webResponse.Data)
	//createProductResponse := schemas.CreateProductResponse{}
	//json.Unmarshal(jsonData, &createProductResponse)
	//assert.NotNil(t, createProductResponse.Id)
	//assert.Equal(t, createProductRequest.Name, createProductResponse.Name)
	//assert.Equal(t, createProductRequest.Price, createProductResponse.Price)
	//assert.Equal(t, createProductRequest.Quantity, createProductResponse.Quantity)
}
