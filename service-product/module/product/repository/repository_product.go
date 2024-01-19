package repository

import (
	"gorm.io/gorm"
	"net/http"
	"service-product/dto"
	"service-product/module/product/entity"
)

type ProductRepositoryInterface interface {
	Create(input *dto.SchemaProduct) (*entity.EntityProduct, dto.SchemaError)
	GetById(id uint64) (entity.EntityProduct, dto.SchemaError)
	GetList() ([]*entity.EntityProduct, dto.SchemaError)
	Update(input *dto.SchemaProduct) (*entity.EntityProduct, dto.SchemaError)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepositorySQL(db *gorm.DB) ProductRepositoryInterface {
	return &productRepository{db: db}
}

func (r *productRepository) Create(input *dto.SchemaProduct) (*entity.EntityProduct, dto.SchemaError) {

	var product entity.EntityProduct
	db := r.db.Model(&product)
	errorCode := make(chan dto.SchemaError, 1)

	checkProductExist := db.Debug().First(&product, "name = ?", input.Name)

	if checkProductExist.RowsAffected > 0 {
		errorCode <- dto.SchemaError{
			StatusCode: http.StatusConflict,
		}
		return &product, <-errorCode
	}

	product.Name = input.Name
	product.Quantity = input.Quantity
	product.Price = input.Price
	product.IsActive = input.IsActive

	addNewProduct := db.Debug().Create(&product)

	if addNewProduct.RowsAffected < 1 {
		errorCode <- dto.SchemaError{
			StatusCode: http.StatusForbidden,
		}
		return &product, <-errorCode
	}
	errorCode <- dto.SchemaError{}

	return &product, <-errorCode
}

func (r *productRepository) GetById(id uint64) (entity.EntityProduct, dto.SchemaError) {

	var students entity.EntityProduct
	errorCode := make(chan dto.SchemaError, 1)
	db := r.db.Model(&students)
	students.ID = id
	resultProduct := db.Debug().First(&students, id)
	if resultProduct.RowsAffected < 1 {
		errorCode <- dto.SchemaError{
			StatusCode: http.StatusNotFound,
		}
		return students, <-errorCode
	}

	errorCode <- dto.SchemaError{}

	return students, <-errorCode
}

func (r *productRepository) GetList() ([]*entity.EntityProduct, dto.SchemaError) {

	var students []*entity.EntityProduct
	db := r.db.Model(&students)
	errorCode := make(chan dto.SchemaError, 1)

	resultsStudents := db.Debug().Find(&students)

	if resultsStudents.RowsAffected < 1 {
		errorCode <- dto.SchemaError{
			StatusCode: http.StatusNotFound,
		}
		return students, <-errorCode
	}

	errorCode <- dto.SchemaError{}
	return students, <-errorCode
}

func (r *productRepository) Update(input *dto.SchemaProduct) (*entity.EntityProduct, dto.SchemaError) {

	var students entity.EntityProduct
	db := r.db.Model(&students)
	errorCode := make(chan dto.SchemaError, 1)

	students.ID = input.ID
	students.Name = input.Name
	students.Quantity = input.Quantity
	students.IsActive = input.IsActive
	students.Price = input.Price

	updateStudent := db.Debug().Where("id = ?", input.ID).Updates(students)

	if updateStudent.RowsAffected < 1 {
		errorCode <- dto.SchemaError{
			StatusCode: http.StatusForbidden,
		}
		return &students, <-errorCode
	}
	errorCode <- dto.SchemaError{}
	return &students, <-errorCode
}
