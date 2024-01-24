package controller

import (
	"context"
	"service-user/dto"
	services "service-user/module/user/usecase"
	pb_user "service-user/proto/user"
	"strconv"
)

type ServiceUserRPCServer struct {
	serviceUser services.ServicesUser
	pb_user.UnimplementedServiceUserRPCServer
}

func NewHandlerRPCUser(serviceUser services.ServicesUser) *ServiceUserRPCServer {
	return &ServiceUserRPCServer{
		serviceUser: serviceUser,
	}
}

func (h *ServiceUserRPCServer) CutBalance(ctx context.Context, req *pb_user.CutBalanceRequest) (*pb_user.CutBalanceResponse, error) {
	var input dto.CutBalanceRequest
	var res pb_user.CutBalanceResponse
	UserId, err := strconv.ParseUint(req.UserId, 10, 64)
	if err != nil {
		return nil, err
	}

	input.UserId = UserId
	input.Balance = req.Amount

	user, respErr := h.serviceUser.CutBalance(input)

	if respErr.Error != nil {
		return nil, respErr.Error
	}

	res.Id = strconv.FormatUint(user.ID, 10)
	res.Balance = user.Balance
	return &res, nil

}

func (h *ServiceUserRPCServer) GetById(ctx context.Context, req *pb_user.GetByIDRequest) (*pb_user.User, error) {

	var res pb_user.User
	UserId, err := strconv.ParseUint(req.Id, 10, 64)
	if err != nil {
		return nil, err
	}

	User, respErr := h.serviceUser.GetById(UserId)
	if respErr.Error != nil {
		return nil, respErr.Error
	} else {
		res.Id = strconv.FormatUint(UserId, 10)
		res.Balance = User.Balance
		res.Fullname = User.Fullname
	}
	return &res, nil

}
