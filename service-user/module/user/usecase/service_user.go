package usecase

import (
	"errors"
	"go-micro.dev/v4/util/log"
	"net/http"
	"service-user/constant"
	"service-user/dto"
	"service-user/module/user/entity"
	repositorys "service-user/module/user/repository"
	"service-user/pkg/kafka"
)

type ServicesUser interface {
	Create(input *dto.SchemaUser) (*entity.Users, dto.ResponseError)
	Update(input *dto.SchemaUser) (*entity.Users, dto.ResponseError)
	GetById(userId uint64) (entity.Users, dto.ResponseError)
	GetList() (*[]entity.Users, dto.ResponseError)
	CutBalance(input dto.CutBalanceRequest) (entity.Users, dto.ResponseError)
}

type userUsecase struct {
	userRepository repositorys.UserRepositoryInterface
	Kafka          *kafka.Producer
}

func NewUserUsecase(repository repositorys.UserRepositoryInterface, kafka *kafka.Producer) ServicesUser {
	return &userUsecase{userRepository: repository, Kafka: kafka}
}

func (s *userUsecase) Create(input *dto.SchemaUser) (*entity.Users, dto.ResponseError) {

	var dataUser dto.SchemaUser
	dataUser.Fullname = input.Fullname
	dataUser.IsActive = input.IsActive
	dataUser.Password = input.Password
	dataUser.Email = input.Email

	res, err := s.userRepository.Create(&dataUser)
	return res, err
}

func (s *userUsecase) Update(input *dto.SchemaUser) (*entity.Users, dto.ResponseError) {

	var user dto.SchemaUser
	user.ID = input.ID
	user.Fullname = input.Fullname
	user.Password = input.Password
	user.Email = input.Email
	user.IsActive = input.IsActive

	res, err := s.userRepository.Update(&user)
	return res, err
}

func (s *userUsecase) GetById(userId uint64) (entity.Users, dto.ResponseError) {
	res, err := s.userRepository.GetById(userId)
	return res, err
}

func (s *userUsecase) GetList() (*[]entity.Users, dto.ResponseError) {
	res, err := s.userRepository.GetList()
	return res, err
}

func (s *userUsecase) CutBalance(input dto.CutBalanceRequest) (entity.Users, dto.ResponseError) {

	user, err := s.GetById(input.UserId)
	if err.Error != nil {
		log.Error(err.Error)
		return user, err
	}
	if (user.Balance - input.Balance) < 1 {
		return user, dto.ResponseError{
			StatusCode: http.StatusForbidden,
			Error:      errors.New("balance more be  less"),
		}

	}
	res, resErr := s.userRepository.CutBalance(input)

	if resErr.Error != nil {
		return res, resErr
	} else {
		data := map[string]interface{}{
			"status":       true,
			"service_name": constant.SERVICE_NAME,
			"order_id":     input.OrderId,
		}

		err := s.Kafka.SendMessage(constant.TOPIC_ORDER_REPLY, data, 1)
		if err != nil {
			return res, dto.ResponseError{
				Error: err,
			}
		}
	}
	return res, dto.ResponseError{}

}
