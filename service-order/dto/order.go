package dto

type CreateOrderRequest struct {
	UserId    uint64 `json:"user_id" `
	ProductId uint64 `json:"product_id" `
	Quantity  uint64 `json:"quantity"`
}

type CreateOrderReplyRequest struct {
	OrderId     uint64
	ServiceName string
}
