package constant

const (
	SERVICE_VERSION = "1.0.3"
	SERVICE_NAME    = "ServiceOrder"
	AUTHOR          = "NandaRusfikri"
	TABLE_ORDER     = "order"

	HTTP_GET                      = "GET"
	CONTENT_TYPE_APPLICATION_JSON = "application/json"

	TOPIC_ORDER_REPLY          = "orderReply"
	TOPIC_PRODUCT_STOCK_UPDATE = "product-stock-update"
	SERVICE_PRODUCT_NAME       = "ServiceProduct"

	ORDER_STATE_PENDING = "PENDING"
	ORDER_STATE_SUCCESS = "SUCCESS"
)

var ChanTopicProduct = make(chan string)
