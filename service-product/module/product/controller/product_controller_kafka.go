package controller

import (
	"encoding/json"
	"service-product/constant"
	"service-product/dto"
	"service-product/module/product/usecase"
)

func NewProductControllerKafka(orderService usecase.UsecaseInterface) {
	go ListenTopicProductQuantityUpdate(orderService)
}

func ListenTopicProductQuantityUpdate(orderService usecase.UsecaseInterface) {

	for {
		select {
		case data := <-constant.TopicProductStockUpdate:

			var msg map[string]interface{}
			json.Unmarshal([]byte(data), &msg)

			quantity := uint64(msg["quantity"].(float64))
			order_id := uint64(msg["order_id"].(float64))
			product_id := uint64(msg["product_id"].(float64))

			orderService.UpdateStock(dto.UpdateStockRequest{
				Quantity:  quantity,
				ProductId: product_id,
				OrderId:   order_id,
			})

		}
	}

}
