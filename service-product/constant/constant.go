package constant

const (
	SERVICE_VERSION   = "1.0.3"
	SERVICE_NAME      = "ServiceProduct"
	AUTHOR            = "NandaRusfikri"
	TABLE_PRODUCT     = "product"
	TABLE_TRANSACTION = "transaction"

	TOPIC_PRODUCT_STOCK_UPDATE = "product-stock-update"
)

var TopicProductStockUpdate = make(chan string)
