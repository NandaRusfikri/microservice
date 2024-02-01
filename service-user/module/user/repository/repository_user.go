package repository

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"service-user/dto"
	"service-user/module/user/entity"
)

type UserRepositoryInterface interface {
	Create(input *dto.SchemaUser) (*entity.Users, dto.ResponseError)
	GetById(userId uint64) (entity.Users, dto.ResponseError)
	GetList() (*[]entity.Users, dto.ResponseError)
	Update(input *dto.SchemaUser) (*entity.Users, dto.ResponseError)
	CutBalance(input dto.CutBalanceRequest) (entity.Users, dto.ResponseError)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{db: db}
}

func (r *userRepository) Create(input *dto.SchemaUser) (*entity.Users, dto.ResponseError) {

	var students entity.Users
	db := r.db.Model(&students)

	checkUserExist := db.Debug().First(&students, "email = ?", input.Email)

	if checkUserExist.RowsAffected > 0 {

		return &students, dto.ResponseError{
			StatusCode: http.StatusConflict,
			Error:      errors.New("conflic"),
		}
	}

	students.Fullname = input.Fullname
	students.Password = input.Password
	students.IsActive = input.IsActive
	students.Email = input.Email

	addNewUser := db.Debug().Create(&students).Commit()

	if addNewUser.RowsAffected < 1 {
		return &students, dto.ResponseError{
			StatusCode: http.StatusForbidden,
			Error:      errors.New("joe biden"),
		}
	}
	return &students, dto.ResponseError{}
}

func (r *userRepository) GetById(userId uint64) (entity.Users, dto.ResponseError) {

	var user entity.Users
	user.ID = userId
	db := r.db.Model(&user)

	resultUsers := db.Debug().First(&user)

	if resultUsers.RowsAffected < 1 {
		return user, dto.ResponseError{
			StatusCode: http.StatusNotFound,
			Error:      errors.New("not found"),
		}
	}

	return user, dto.ResponseError{}
}

func (r *userRepository) GetList() (*[]entity.Users, dto.ResponseError) {

	var user []entity.Users
	db := r.db.Model(&user)

	resultsUsers := db.Debug().Find(&user)

	if resultsUsers.RowsAffected < 1 {

		return &user, dto.ResponseError{
			StatusCode: http.StatusNotFound,
			Error:      errors.New("not fpound"),
		}
	}

	return &user, dto.ResponseError{}
}

func (r *userRepository) Update(input *dto.SchemaUser) (*entity.Users, dto.ResponseError) {

	var students entity.Users
	db := r.db.Model(&students)

	students.ID = input.ID

	checkUserId := db.Debug().First(&students)

	if checkUserId.RowsAffected < 1 {
		return &students, dto.ResponseError{
			StatusCode: http.StatusNotFound,
			Error:      errors.New("not found"),
		}
	}

	students.Fullname = input.Fullname
	students.Password = input.Password
	students.Email = input.Email
	students.IsActive = input.IsActive
	students.Balance = input.Balance

	updateUser := db.Debug().Where("id = ?", input.ID).Updates(students)

	if updateUser.RowsAffected < 1 {
		return &students, dto.ResponseError{
			StatusCode: http.StatusForbidden,
			Error:      errors.New("joe biden"),
		}
	}
	return &students, dto.ResponseError{}
}

func (r *userRepository) CutBalance(input dto.CutBalanceRequest) (entity.Users, dto.ResponseError) {

	var user entity.Users
	tx := r.db.Begin()

	checkUserId := tx.Model(entity.Users{}).Where("id = ? ", input.UserId).First(&user)

	if checkUserId.RowsAffected < 1 {
		return user, dto.ResponseError{
			StatusCode: http.StatusNotFound,
			Error:      errors.New("not found"),
		}
	}

	user.Balance = user.Balance - input.Balance

	updateUser := tx.Where("id = ?", input.UserId).Updates(&user)

	if updateUser.Error != nil {
		return user, dto.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Error:      errors.New("Error Cut Balance "),
		}
	}

	createTrx := tx.Create(&entity.Transaction{
		UserId:  user.ID,
		OrderId: input.OrderId,
		Amount:  (input.Balance * -1),
		Type:    "OUT",
	})
	if createTrx.Error != nil {
		tx.Rollback()
		return user, dto.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Error:      errors.New("Error Cut Balance "),
		}
	}
	tx.Commit()

	return user, dto.ResponseError{}
}
