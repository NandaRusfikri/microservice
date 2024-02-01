package repository

import (
	"gorm.io/gorm"
	"net/http"
	"service-product/dto"
	"service-product/module/product/entity"
)

type ProductRepositoryInterface interface {
	Create(input *dto.SchemaProduct) (*entity.Product, dto.ResponseError)
	GetById(id uint64) (entity.Product, dto.ResponseError)
	GetList() ([]*entity.Product, dto.ResponseError)
	UpdateStock(input dto.UpdateStockRequest) (*entity.Product, dto.ResponseError)
}

type productRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) ProductRepositoryInterface {
	return &productRepository{db: db}
}

func (r *productRepository) Create(input *dto.SchemaProduct) (*entity.Product, dto.ResponseError) {

	var product entity.Product
	db := r.db.Model(&product)
	errorCode := make(chan dto.ResponseError, 1)

	checkProductExist := db.Debug().First(&product, "name = ?", input.Name)

	if checkProductExist.RowsAffected > 0 {
		errorCode <- dto.ResponseError{
			StatusCode: http.StatusConflict,
		}
		return &product, <-errorCode
	}

	product.Name = input.Name
	product.Stock = input.Quantity
	product.Price = input.Price
	product.IsActive = input.IsActive

	addNewProduct := db.Debug().Create(&product)

	if addNewProduct.RowsAffected < 1 {
		errorCode <- dto.ResponseError{
			StatusCode: http.StatusForbidden,
		}
		return &product, <-errorCode
	}
	errorCode <- dto.ResponseError{}

	return &product, <-errorCode
}

func (r *productRepository) GetById(id uint64) (entity.Product, dto.ResponseError) {

	var students entity.Product
	errorCode := make(chan dto.ResponseError, 1)
	db := r.db.Model(&students)
	students.ID = id
	resultProduct := db.Debug().First(&students, id)
	if resultProduct.RowsAffected < 1 {
		errorCode <- dto.ResponseError{
			StatusCode: http.StatusNotFound,
		}
		return students, <-errorCode
	}

	errorCode <- dto.ResponseError{}

	return students, <-errorCode
}

func (r *productRepository) GetList() ([]*entity.Product, dto.ResponseError) {

	var students []*entity.Product
	db := r.db.Model(&students)
	errorCode := make(chan dto.ResponseError, 1)

	resultsStudents := db.Debug().Find(&students)

	if resultsStudents.RowsAffected < 1 {
		errorCode <- dto.ResponseError{
			StatusCode: http.StatusNotFound,
		}
		return students, <-errorCode
	}

	errorCode <- dto.ResponseError{}
	return students, <-errorCode
}

func (r *productRepository) UpdateStock(input dto.UpdateStockRequest) (*entity.Product, dto.ResponseError) {

	var product entity.Product

	tx := r.db.Begin()

	find := tx.Model(&product).Where("id = ?", input.ProductId).First(&product)
	if find.Error != nil {
		tx.Rollback()
		return nil, dto.ResponseError{
			Error:      find.Error,
			StatusCode: 500,
		}
	}

	product.ID = input.ProductId
	product.Stock = product.Stock - input.Quantity
	update := tx.Debug().Updates(&product)

	if update.Error != nil {
		tx.Rollback()
		return nil, dto.ResponseError{
			StatusCode: 500,
			Error:      update.Error,
		}
	}

	create := tx.Debug().Create(&entity.Transaction{
		ProductId: product.ID,
		Price:     product.Price,
		Quantity:  input.Quantity,
		OrderId:   input.OrderId,
	})

	if create.Error != nil {
		tx.Rollback()
		return nil, dto.ResponseError{
			StatusCode: 500,
			Error:      create.Error,
		}
	}
	tx.Commit()
	return &product, dto.ResponseError{}
}
