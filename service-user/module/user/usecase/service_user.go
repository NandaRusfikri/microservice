package usecase

import (
	"service-user/dto"
	"service-user/module/user/entity"
	repositorys "service-user/module/user/repository"
)

type ServicesUser interface {
	CreateUserService(input *dto.SchemaUser) (*entity.Users, dto.SchemaError)

	UpdateUserService(input *dto.SchemaUser) (*entity.Users, dto.SchemaError)

	DeleteUserService(input *dto.SchemaUser) (*entity.Users, dto.SchemaError)

	ResultUserService(input *dto.SchemaUser) (*entity.Users, dto.SchemaError)

	ResultsUserService() (*[]entity.Users, dto.SchemaError)

	CutBalanceService(input *dto.SchemaCutBalanceRequest) (*entity.Users, dto.SchemaError)
}

type serviceUser struct {
	repository repositorys.UserRepository
}

func NewUserService(repository repositorys.UserRepository) ServicesUser {
	return &serviceUser{repository: repository}
}

func (s *serviceUser) CreateUserService(input *dto.SchemaUser) (*entity.Users, dto.SchemaError) {

	var student dto.SchemaUser
	student.Fullname = input.Fullname
	student.IsActive = input.IsActive
	student.Password = input.Password
	student.Email = input.Email

	res, err := s.repository.CreateUserRepository(&student)
	return res, err
}

func (s *serviceUser) UpdateUserService(input *dto.SchemaUser) (*entity.Users, dto.SchemaError) {

	var student dto.SchemaUser
	student.ID = input.ID
	student.Fullname = input.Fullname
	student.Password = input.Password
	student.Email = input.Email
	student.IsActive = input.IsActive

	res, err := s.repository.UpdateUserRepository(&student)
	return res, err
}

func (s *serviceUser) DeleteUserService(input *dto.SchemaUser) (*entity.Users, dto.SchemaError) {

	var student dto.SchemaUser
	student.ID = input.ID

	res, err := s.repository.DeleteUserRepository(&student)
	return res, err
}

func (s *serviceUser) ResultUserService(input *dto.SchemaUser) (*entity.Users, dto.SchemaError) {

	var student dto.SchemaUser
	student.ID = input.ID

	res, err := s.repository.ResultUserRepository(&student)
	return res, err
}

func (s *serviceUser) ResultsUserService() (*[]entity.Users, dto.SchemaError) {

	res, err := s.repository.ResultsUserRepository()
	return res, err
}

func (s *serviceUser) CutBalanceService(input *dto.SchemaCutBalanceRequest) (*entity.Users, dto.SchemaError) {
	res, err := s.repository.CutBalanceRepository(input)
	return res, err
}
