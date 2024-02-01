package repositorys

import (
	"fmt"
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

func (r *orderRepository) CreateOrderReply(input entities.TopicOrderReply) (bool, dto.ResponseError) {

	create := r.db.Create(&input)

	if create.Error != nil {
		return false, dto.ResponseError{
			Error:      create.Error,
			StatusCode: http.StatusInternalServerError,
		}
	}

	var dataOrder []entities.TopicOrderReply
	result := r.db.Model(&entities.TopicOrderReply{}).Where("order_id = ?", input.OrderId).Find(&dataOrder)

	if result.Error != nil {
		return false, dto.ResponseError{
			Error:      result.Error,
			StatusCode: http.StatusInternalServerError,
		}
	}

	fmt.Println("dataOrder", len(dataOrder))
	if len(dataOrder) == 2 {
		return true, dto.ResponseError{}
	}

	return false, dto.ResponseError{}
}

func (r *orderRepository) UpdateState(orderId uint64, state string) (*entities.Order, dto.ResponseError) {

	order := entities.Order{
		ID:    int64(orderId),
		State: state,
	}
	update := r.db.Debug().Updates(&order)
	if update.Error != nil {
		return nil, dto.ResponseError{
			Error:      update.Error,
			StatusCode: http.StatusInternalServerError,
		}
	}
	return &order, dto.ResponseError{}
}
