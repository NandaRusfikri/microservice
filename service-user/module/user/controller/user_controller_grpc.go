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

func (h *ServiceUserRPCServer) CutBalanceRPC(ctx context.Context, req *pb_user.CutBalanceRequest) (*pb_user.CutBalanceResponse, error) {
	var input dto.CutBalanceRequest
	var res pb_user.CutBalanceResponse
	input.UserId, _ = strconv.ParseUint(req.UserId, 10, 64)
	input.Balance = req.Amount

	user, err := h.serviceUser.CutBalance(input)

	if err.Error != nil {
		return nil, err.Error
	}

	res.Id = strconv.FormatUint(user.ID, 10)
	res.Balance = user.Balance
	return &res, nil

}

func (h *ServiceUserRPCServer) GetUserByIDRPC(ctx context.Context, req *pb_user.GetByIDRequest) (*pb_user.User, error) {

	var res pb_user.User
	userID, _ := strconv.ParseUint(req.Id, 10, 64)
	User, err := h.serviceUser.GetById(userID)
	if err.Error != nil {
		return nil, err.Error
	} else {
		res.Id = strconv.FormatUint(userID, 10)
		res.Balance = User.Balance
		res.Fullname = User.Fullname
	}
	return &res, nil

}
