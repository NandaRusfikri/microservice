package controller

import (
	"encoding/json"
	"fmt"
	"service-product/constant"
	"service-product/dto"
	"service-product/module/product/usecase"
)

func NewProductControllerKafka(orderService usecase.ServiceProductInterface) {
	go ListenTopicProductQuantityUpdate(orderService)
}

func ListenTopicProductQuantityUpdate(orderService usecase.ServiceProductInterface) {

	for {
		select {
		case data := <-constant.TopicProductStockUpdate:

			var msg map[string]interface{}
			json.Unmarshal([]byte(data), &msg)

			fmt.Println("msg", msg)

			orderService.UpdateStock(dto.UpdateStockRequest{
				//Quantity: msg[""],
			})
		}
	}

}
