package usecase

import (
	"errors"
	"fmt"
	"go-micro.dev/v4/util/log"
	"net/http"
	"service-user/dto"
	"service-user/module/user/entity"
	repositorys "service-user/module/user/repository"
)

type ServicesUser interface {
	Create(input *dto.SchemaUser) (*entity.Users, dto.SchemaError)
	Update(input *dto.SchemaUser) (*entity.Users, dto.SchemaError)
	GetById(userId uint64) (entity.Users, dto.SchemaError)
	GetList() (*[]entity.Users, dto.SchemaError)
	CutBalance(input dto.CutBalanceRequest) (entity.Users, dto.SchemaError)
}

type userUsecase struct {
	userRepository repositorys.UserRepositoryInterface
}

func NewUserUsecase(repository repositorys.UserRepositoryInterface) ServicesUser {
	return &userUsecase{userRepository: repository}
}

func (s *userUsecase) Create(input *dto.SchemaUser) (*entity.Users, dto.SchemaError) {

	var student dto.SchemaUser
	student.Fullname = input.Fullname
	student.IsActive = input.IsActive
	student.Password = input.Password
	student.Email = input.Email

	res, err := s.userRepository.Create(&student)
	return res, err
}

func (s *userUsecase) Update(input *dto.SchemaUser) (*entity.Users, dto.SchemaError) {

	var user dto.SchemaUser
	user.ID = input.ID
	user.Fullname = input.Fullname
	user.Password = input.Password
	user.Email = input.Email
	user.IsActive = input.IsActive

	res, err := s.userRepository.Update(&user)
	return res, err
}

func (s *userUsecase) GetById(userId uint64) (entity.Users, dto.SchemaError) {
	res, err := s.userRepository.GetById(userId)
	return res, err
}

func (s *userUsecase) GetList() (*[]entity.Users, dto.SchemaError) {
	res, err := s.userRepository.GetList()
	return res, err
}

func (s *userUsecase) CutBalance(input dto.CutBalanceRequest) (entity.Users, dto.SchemaError) {

	user, err := s.GetById(input.UserId)
	if err.Error != nil {
		log.Error("aaa", err.Error)
		return user, err
	}
	fmt.Printf("%+v\n ", input)
	if (user.Balance - input.Balance) < 0 {
		if err.Error != nil {
			return user, dto.SchemaError{
				StatusCode: http.StatusForbidden,
				Error:      errors.New("balance less"),
			}
		}
	}
	return s.userRepository.CutBalance(input)
	//return dto.SchemaError{}
}
