package constant

const (
	SERVICE_VERSION   = "1.0.3"
	SERVICE_NAME      = "ServiceProduct"
	AUTHOR            = "NandaRusfikri"
	TABLE_PRODUCT     = "product"
	TABLE_TRANSACTION = "transaction"

	TOPIC_NEW_ORDER = "new-order"
)

var TopicProductStockUpdate = make(chan string)
