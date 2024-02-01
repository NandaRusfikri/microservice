package constant

const (
	SERVICE_VERSION   = "1.0.3"
	SERVICE_NAME      = "ServiceOrder"
	AUTHOR            = "NandaRusfikri"
	TABLE_ORDER       = "order"
	TABLE_ORDER_REPLY = "order_reply"

	HTTP_GET                      = "GET"
	CONTENT_TYPE_APPLICATION_JSON = "application/json"

	TOPIC_ORDER_REPLY    = "order-reply"
	TOPIC_NEW_ORDER      = "new-order"
	SERVICE_PRODUCT_NAME = "ServiceProduct"

	ORDER_STATE_PENDING = "PENDING"
	ORDER_STATE_SUCCESS = "SUCCESS"
)

var ChanTopicNewOrder = make(chan string)
var ChanTopicOrderReply = make(chan string)
