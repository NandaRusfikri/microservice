package repositorys

import (
	"context"
	"service-order/entities"
	"service-order/schemas"
)

type OrderRepository interface {
	CreateOrderRepository(input *schemas.SchemaOrder) (*entities.EntityOrder, schemas.SchemaDatabaseError)
	DeleteOrderRepository(input *schemas.SchemaOrder) (*entities.EntityOrder, schemas.SchemaDatabaseError)
	ResultOrderRepository(input *schemas.SchemaOrder) (*entities.EntityOrder, schemas.SchemaDatabaseError)
	FindAllOrderRepository() ([]*entities.EntityOrder, schemas.SchemaDatabaseError)
	UpdateOrderRepository(input *schemas.SchemaOrder) (*entities.EntityOrder, schemas.SchemaDatabaseError)

}

type OrderRepositoryGRPC interface {
	FindProductByIdRepository(ctx context.Context,ProductId uint64) (entities.EntityProduct, schemas.SchemaDatabaseError)
}
