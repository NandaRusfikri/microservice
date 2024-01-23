package repositorys

import (
	"gorm.io/gorm"
	"net/http"
	"service-order/entities"
	"service-order/schemas"
)

type orderRepositoryImplSQL struct {
	db *gorm.DB
}

func NewOrderRepositorySQL(db *gorm.DB) OrderRepository {
	return &orderRepositoryImplSQL{db: db}
}

func (r *orderRepositoryImplSQL) CreateOrderRepository(input *schemas.SchemaOrder) (*entities.EntityOrder, schemas.SchemaDatabaseError) {

	var product entities.EntityOrder
	db := r.db.Model(&product)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	product.Amount = input.Amount
	product.UserId = input.UserId
	product.ProductId = input.ProductId

	addNewOrder := db.Debug().Create(&product)

	if addNewOrder.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return &product, <-errorCode
	}
	errorCode <- schemas.SchemaDatabaseError{}




	return &product, <-errorCode
}

func (r *orderRepositoryImplSQL) DeleteOrderRepository(input *schemas.SchemaOrder) (*entities.EntityOrder, schemas.SchemaDatabaseError) {

	var students entities.EntityOrder
	db := r.db.Model(&students)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	students.ID = input.ID
	checkStudentId := db.Debug().First(&students)

	if checkStudentId.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return &students, <-errorCode
	}

	deleteStudentId := db.Debug().Delete(&students)

	if deleteStudentId.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return &students, <-errorCode
	}
	errorCode <- schemas.SchemaDatabaseError{}

	return &students, <-errorCode
}

func (r *orderRepositoryImplSQL) ResultOrderRepository(input *schemas.SchemaOrder) (*entities.EntityOrder, schemas.SchemaDatabaseError) {

	var students entities.EntityOrder
	db := r.db.Model(&students)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	resultStudents := db.Debug().First(&students)

	if resultStudents.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return &students, <-errorCode
	}

	return &students, <-errorCode
}

func (r *orderRepositoryImplSQL) FindAllOrderRepository() ([]*entities.EntityOrder, schemas.SchemaDatabaseError) {

	var students []*entities.EntityOrder
	db := r.db.Model(&students)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)


	resultsStudents := db.Debug().Find(&students)

	if resultsStudents.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return students, <-errorCode
	}


	errorCode <- schemas.SchemaDatabaseError{}
	return students, <-errorCode
}

func (r *orderRepositoryImplSQL) UpdateOrderRepository(input *schemas.SchemaOrder) (*entities.EntityOrder, schemas.SchemaDatabaseError) {

	var students entities.EntityOrder
	db := r.db.Model(&students)
	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	students.ID = input.ID

	checkStudentId := db.Debug().First(&students)

	if checkStudentId.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return &students, <-errorCode
	}

	students.UserId = input.UserId
	students.ProductId = input.ProductId
	students.Amount = input.Amount


	updateStudent := db.Debug().Where("id = ?", input.ID).Updates(students)

	if updateStudent.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return &students, <-errorCode
	}
	errorCode <- schemas.SchemaDatabaseError{}
	return &students, <-errorCode
}


