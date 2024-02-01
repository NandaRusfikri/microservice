package constant

const (
	SERVICE_VERSION   = "1.0.3"
	SERVICE_NAME      = "ServiceUser"
	AUTHOR            = "NandaRusfikri"
	TABLE_USERS       = "users"
	TABLE_TRANSACTION = "transaction"
	TOPIC_NEW_ORDER   = "new-order"
	TOPIC_ORDER_REPLY = "order-reply"
)

var TopicNewOrder = make(chan string)
