package repositorys

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductGetById(t *testing.T) {
	repo := NewOrderRepositoryGRPC()
	ctx := 	context.Background()
	Product, _ := repo.FindProductByIdRepository(ctx,5)

	assert.Equal(t, int64(1), int64(Product.ID))
	//assert.Equal(t, "OK", webResponse.Status)
	//
	//jsonData, _ := json.Marshal(webResponse.Data)
	//createProductResponse := schemas.CreateProductResponse{}
	//json.Unmarshal(jsonData, &createProductResponse)
	//assert.NotNil(t, createProductResponse.Id)
	//assert.Equal(t, createProductRequest.Name, createProductResponse.Name)
	//assert.Equal(t, createProductRequest.Price, createProductResponse.Price)
	//assert.Equal(t, createProductRequest.Quantity, createProductResponse.Quantity)
}
