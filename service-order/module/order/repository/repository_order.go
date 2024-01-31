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

func (r *orderRepository) Create(input entities.Order) (*entities.Order, dto.ResponseError) {

	create := r.db.Debug().Create(&input)

	if create.Error != nil {
		return nil, dto.ResponseError{
			Error:      create.Error,
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &input, dto.ResponseError{}
}

func (r *orderRepository) GetById(orderId uint64) (*entities.Order, dto.ResponseError) {

	var dataOrder entities.Order
	db := r.db.Model(&dataOrder)

	result := db.Debug().Model(&entities.Order{}).Where("id = ?", orderId).First(&dataOrder)

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
