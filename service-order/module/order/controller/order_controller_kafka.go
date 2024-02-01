package handlers

import (
	"encoding/json"
	"fmt"
	"service-order/constant"
	"service-order/dto"
	orders "service-order/module/order"
)

func NewOrderControllerKafka(orderService orders.ServiceInterface) {
	go ListenTopicOrderReply(orderService)
}

func ListenTopicOrderReply(orderService orders.ServiceInterface) {

	for {
		select {
		case data := <-constant.ChanTopicOrderReply:

			var msg map[string]interface{}
			json.Unmarshal([]byte(data), &msg)

			orderId := uint64(msg["order_id"].(float64))
			serviceName := msg["service_name"].(string)
			status := msg["status"].(bool)

			if status {
				fmt.Println("orderId", orderId)
				fmt.Println("serviceName", serviceName)
				orderService.CreateOrderReply(dto.CreateOrderReplyRequest{
					OrderId:     orderId,
					ServiceName: serviceName,
				})
			} else {
				fmt.Println("ada Error nih dari id  service ", orderId, serviceName)
			}

		}
	}

}
