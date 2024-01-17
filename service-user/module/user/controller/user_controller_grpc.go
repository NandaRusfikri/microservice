package controller

import (
	"context"
	"errors"
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

func (h *ServiceUserRPCServer) CutBalanceRPC(ctx context.Context, req *pb_user.CutBalanceRequest) (*pb_user.ModelProtoUser, error) {
	var input dto.SchemaCutBalanceRequest
	var res pb_user.ModelProtoUser
	input.UserId, _ = strconv.ParseInt(req.UserId, 10, 64)
	input.Balance = req.Amount

	CutBalance, err := h.serviceUser.CutBalanceService(&input)

	switch err.StatusCode {
	case 500:
		errors.New("Internal Server Error")
	case 400:
		errors.New("Less Balance maybe")
	default:
		res.Id = strconv.FormatInt(CutBalance.ID, 10)
		res.Balance = CutBalance.Balance
		res.Fullname = CutBalance.Fullname
	}
	return &res, nil
}

func (h *ServiceUserRPCServer) GetUserByIDRPC(ctx context.Context, req *pb_user.UserByIDRequest) (*pb_user.ModelProtoUser, error) {
	var input dto.SchemaUser
	var res pb_user.ModelProtoUser
	UserId, _ := strconv.ParseInt(req.Id, 10, 64)

	input.ID = UserId
	User, err := h.serviceUser.ResultUserService(&input)

	switch err.StatusCode {
	case 404:
		errors.New("User Id not found")
	default:
		res.Id = strconv.FormatInt(User.ID, 10)
		res.Balance = User.Balance
		res.Fullname = User.Fullname
	}
	return &res, nil

}
