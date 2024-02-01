package controller

import (
	"encoding/json"
	"service-user/constant"
	"service-user/dto"
	"service-user/module/user/usecase"
)

func NewUserControllerKafka(orderService usecase.ServicesUser) {
	go ListenTopicNewOrder(orderService)
}

func ListenTopicNewOrder(orderService usecase.ServicesUser) {

	for {
		select {
		case data := <-constant.TopicNewOrder:

			var msg map[string]interface{}
			json.Unmarshal([]byte(data), &msg)

			user_id := uint64(msg["user_id"].(float64))
			order_id := uint64(msg["order_id"].(float64))
			amount := uint64(msg["amount"].(float64))

			orderService.CutBalance(dto.CutBalanceRequest{
				UserId:  user_id,
				Balance: int64(amount),
				OrderId: order_id,
			})
		}
	}

}
