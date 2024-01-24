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

	var dataUser dto.SchemaUser
	dataUser.Fullname = input.Fullname
	dataUser.IsActive = input.IsActive
	dataUser.Password = input.Password
	dataUser.Email = input.Email

	res, err := s.userRepository.Create(&dataUser)
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
		log.Error(err.Error)
		return user, err
	}
	fmt.Println("balance ", (user.Balance - input.Balance))
	if (user.Balance - input.Balance) < 1 {
		return user, dto.SchemaError{
			StatusCode: http.StatusForbidden,
			Error:      errors.New("balance more be  less"),
		}

	}
	return s.userRepository.CutBalance(input)

}
