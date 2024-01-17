package repository

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"service-user/dto"
	"service-user/module/user/entity"
)

type UserRepository interface {
	CreateUserRepository(input *dto.SchemaUser) (*entity.Users, dto.SchemaError)

	DeleteUserRepository(input *dto.SchemaUser) (*entity.Users, dto.SchemaError)

	ResultUserRepository(input *dto.SchemaUser) (*entity.Users, dto.SchemaError)

	ResultsUserRepository() (*[]entity.Users, dto.SchemaError)

	UpdateUserRepository(input *dto.SchemaUser) (*entity.Users, dto.SchemaError)

	CutBalanceRepository(input *dto.SchemaCutBalanceRequest) (*entity.Users, dto.SchemaError)
}

type userRepositoryImplSQL struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImplSQL{db: db}
}

func (r *userRepositoryImplSQL) CreateUserRepository(input *dto.SchemaUser) (*entity.Users, dto.SchemaError) {

	var students entity.Users
	db := r.db.Model(&students)
	errorCode := make(chan dto.SchemaError, 1)

	checkUserExist := db.Debug().First(&students, "email = ?", input.Email)

	if checkUserExist.RowsAffected > 0 {
		errorCode <- dto.SchemaError{
			StatusCode: http.StatusConflict,
			Error:      errors.New("conflic"),
		}
		return &students, <-errorCode
	}

	students.Fullname = input.Fullname
	students.Password = input.Password
	students.IsActive = input.IsActive
	students.Email = input.Email

	addNewUser := db.Debug().Create(&students).Commit()

	if addNewUser.RowsAffected < 1 {
		errorCode <- dto.SchemaError{
			StatusCode: http.StatusForbidden,
			Error:      errors.New("joe biden"),
		}
		return &students, <-errorCode
	}
	errorCode <- dto.SchemaError{}

	return &students, <-errorCode
}

func (r *userRepositoryImplSQL) DeleteUserRepository(input *dto.SchemaUser) (*entity.Users, dto.SchemaError) {

	var students entity.Users
	db := r.db.Model(&students)
	errorCode := make(chan dto.SchemaError, 1)

	checkUserId := db.Debug().First(&students)

	if checkUserId.RowsAffected < 1 {
		errorCode <- dto.SchemaError{
			StatusCode: http.StatusNotFound,
			Error:      errors.New("not found"),
		}
		return &students, <-errorCode
	}

	deleteUserId := db.Debug().Delete(&students)

	if deleteUserId.RowsAffected < 1 {
		errorCode <- dto.SchemaError{
			StatusCode: http.StatusForbidden,
			Error:      errors.New("joe biden"),
		}
		return &students, <-errorCode
	}
	errorCode <- dto.SchemaError{}

	return &students, <-errorCode
}

func (r *userRepositoryImplSQL) ResultUserRepository(input *dto.SchemaUser) (*entity.Users, dto.SchemaError) {

	var students entity.Users
	db := r.db.Model(&students)
	errorCode := make(chan dto.SchemaError, 1)

	resultUsers := db.Debug().First(&students)

	if resultUsers.RowsAffected < 1 {
		errorCode <- dto.SchemaError{
			StatusCode: http.StatusNotFound,
			Error:      errors.New("not founf"),
		}
		return &students, <-errorCode
	}
	errorCode <- dto.SchemaError{}

	return &students, <-errorCode
}

func (r *userRepositoryImplSQL) ResultsUserRepository() (*[]entity.Users, dto.SchemaError) {

	var students []entity.Users
	db := r.db.Model(&students)
	errorCode := make(chan dto.SchemaError, 1)

	resultsUsers := db.Debug().Find(&students)

	if resultsUsers.RowsAffected < 1 {
		errorCode <- dto.SchemaError{
			StatusCode: http.StatusNotFound,
			Error:      errors.New("not fpound"),
		}
		return &students, <-errorCode
	}

	errorCode <- dto.SchemaError{}
	return &students, <-errorCode
}

func (r *userRepositoryImplSQL) UpdateUserRepository(input *dto.SchemaUser) (*entity.Users, dto.SchemaError) {

	var students entity.Users
	db := r.db.Model(&students)
	errorCode := make(chan dto.SchemaError, 1)

	students.ID = input.ID

	checkUserId := db.Debug().First(&students)

	if checkUserId.RowsAffected < 1 {
		errorCode <- dto.SchemaError{
			StatusCode: http.StatusNotFound,
			Error:      errors.New("not found"),
		}
		return &students, <-errorCode
	}

	students.Fullname = input.Fullname
	students.Password = input.Password
	students.Email = input.Email
	students.IsActive = input.IsActive
	students.Balance = input.Balance

	updateUser := db.Debug().Where("id = ?", input.ID).Updates(students)

	if updateUser.RowsAffected < 1 {
		errorCode <- dto.SchemaError{
			StatusCode: http.StatusForbidden,
			Error:      errors.New("joe biden"),
		}
		return &students, <-errorCode
	}
	errorCode <- dto.SchemaError{}

	return &students, <-errorCode
}

func (r *userRepositoryImplSQL) CutBalanceRepository(input *dto.SchemaCutBalanceRequest) (*entity.Users, dto.SchemaError) {

	var students entity.Users
	db := r.db.Model(&students)
	errorCode := make(chan dto.SchemaError, 1)

	students.ID = input.UserId

	checkUserId := db.First(&students)

	if checkUserId.RowsAffected < 1 {
		errorCode <- dto.SchemaError{
			StatusCode: http.StatusNotFound,
			Error:      errors.New("mot fpund"),
		}
		return &students, <-errorCode
	}
	if (students.Balance - input.Balance) > 0 {
		students.Balance = students.Balance - input.Balance

		updateUser := db.Where("id = ?", input.UserId).Updates(students)

		if updateUser.RowsAffected < 1 {
			errorCode <- dto.SchemaError{
				StatusCode: http.StatusForbidden,
				Error:      errors.New("joe biden"),
			}
			return &students, <-errorCode
		}
	} else {
		errorCode <- dto.SchemaError{
			StatusCode: http.StatusBadRequest,
			Error:      errors.New("bad requts"),
		}
		return &students, <-errorCode
	}

	errorCode <- dto.SchemaError{}
	return &students, <-errorCode
}
