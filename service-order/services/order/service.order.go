package order

import (
	"service-order/entities"
	"service-order/schemas"
)

type OrderService interface {
	CreateOrderService(input *schemas.SchemaOrder) (*entities.EntityOrder, schemas.SchemaDatabaseError)
 DeleteOrderService(input *schemas.SchemaOrder) (*entities.EntityOrder, schemas.SchemaDatabaseError)

 ResultOrderService(input *schemas.SchemaOrder) (*entities.EntityOrder, schemas.SchemaDatabaseError)

FindAllOrderService() ([]*entities.EntityOrder, schemas.SchemaDatabaseError)

 UpdateOrderService(input *schemas.SchemaOrder) (*entities.EntityOrder, schemas.SchemaDatabaseError)



}
