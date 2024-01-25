package repositorys

import (
	"gorm.io/gorm"
	"net/http"
	"service-order/dto"
	"service-order/module/order"
	"service-order/module/order/entity"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepositorySQL(db *gorm.DB) repositorys.RepositoryInterface {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(input *dto.SchemaOrder) (*entities.Order, dto.ResponseError) {

	var product entities.Order
	db := r.db.Model(&product)

	product.Amount = input.Amount
	product.UserId = input.UserId
	product.ProductId = input.ProductId

	create := db.Debug().Create(&product)

	if create.Error != nil {
		return &product, dto.ResponseError{
			Error:      create.Error,
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &product, dto.ResponseError{}
}

func (r *orderRepository) GetById(input *dto.SchemaOrder) (*entities.Order, dto.ResponseError) {

	var dataOrder entities.Order
	db := r.db.Model(&dataOrder)

	result := db.Debug().First(&dataOrder)

	if result.Error != nil {
		return nil, dto.ResponseError{
			Error:      result.Error,
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &dataOrder, dto.ResponseError{}
}

func (r *orderRepository) GetList() ([]*entities.Order, dto.ResponseError) {

	var dataOrder []*entities.Order
	db := r.db.Model(&dataOrder)

	result := db.Debug().Find(&dataOrder)

	if result.Error != nil {
		return nil, dto.ResponseError{
			Error:      result.Error,
			StatusCode: http.StatusInternalServerError,
		}
	}

	return dataOrder, dto.ResponseError{}
}
