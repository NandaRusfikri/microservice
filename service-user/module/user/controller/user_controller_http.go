package controller

import (
	"net/http"
	"service-user/dto"
	services "service-user/module/user/usecase"
	"service-user/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type controllerUser struct {
	service services.ServicesUser
}

func NewUserControllerHTTP(route *gin.Engine, service services.ServicesUser) {
	controller := &controllerUser{service: service}

	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/user", controller.UserCreate)
	groupRoute.GET("/users", controller.UserList)
	groupRoute.GET("/user/:id", controller.UserDetail)
	groupRoute.DELETE("/user/:id", controller.UserDelete)
	groupRoute.PUT("/user/:id", controller.UserUpdate)

}

func (h *controllerUser) UserCreate(ctx *gin.Context) {

	var input dto.SchemaUser
	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.APIResponse(ctx, dto.APIResponse{
			Message: "request invalid " + err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	_, err := h.service.CreateUserService(&input)

	if err.Error != nil {
		util.APIResponse(ctx, dto.APIResponse{
			Code:    err.StatusCode,
			Message: err.Error.Error(),
		})
	} else {
		util.APIResponse(ctx, dto.APIResponse{
			Message: "Success",
			Code:    http.StatusOK,
		})
	}

}

func (h *controllerUser) CutBalanceHandler(ctx *gin.Context) {

	var input dto.SchemaCutBalanceRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.APIResponse(ctx, dto.APIResponse{
			Message: "request invalid " + err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	_, err := h.service.CutBalanceService(&input)
	if err.Error != nil {
		util.APIResponse(ctx, dto.APIResponse{
			Code:    err.StatusCode,
			Message: err.Error.Error(),
		})
	} else {
		util.APIResponse(ctx, dto.APIResponse{
			Message: "Success",
			Code:    http.StatusOK,
		})
	}
}

func (h *controllerUser) UserDelete(ctx *gin.Context) {

	var input dto.SchemaUser
	input.ID, _ = strconv.ParseInt(ctx.Param("id"), 10, 64)

	_, err := h.service.DeleteUserService(&input)

	if err.Error != nil {
		util.APIResponse(ctx, dto.APIResponse{
			Code:    err.StatusCode,
			Message: err.Error.Error(),
		})
	} else {
		util.APIResponse(ctx, dto.APIResponse{
			Message: "Success",
			Code:    http.StatusOK,
		})
	}
}

func (h *controllerUser) UserList(ctx *gin.Context) {

	res, err := h.service.ResultsUserService()

	if err.Error != nil {
		util.APIResponse(ctx, dto.APIResponse{
			Code:    err.StatusCode,
			Message: err.Error.Error(),
		})
	} else {
		util.APIResponse(ctx, dto.APIResponse{
			Message: "Success",
			Code:    http.StatusOK,
			Data:    res,
		})
	}
}

func (h *controllerUser) UserUpdate(ctx *gin.Context) {

	var input dto.SchemaUser
	input.ID, _ = strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.APIResponse(ctx, dto.APIResponse{
			Message: "request invalid " + err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	_, err := h.service.UpdateUserService(&input)

	if err.Error != nil {
		util.APIResponse(ctx, dto.APIResponse{
			Code:    err.StatusCode,
			Message: err.Error.Error(),
		})
	} else {
		util.APIResponse(ctx, dto.APIResponse{
			Message: "Success",
			Code:    http.StatusOK,
		})
	}
}

func (h *controllerUser) UserDetail(ctx *gin.Context) {

	var input dto.SchemaUser
	input.ID, _ = strconv.ParseInt(ctx.Param("id"), 10, 64)

	res, err := h.service.ResultUserService(&input)

	if err.Error != nil {
		util.APIResponse(ctx, dto.APIResponse{
			Code:    err.StatusCode,
			Message: err.Error.Error(),
		})
	} else {
		util.APIResponse(ctx, dto.APIResponse{
			Message: "Success",
			Code:    http.StatusOK,
			Data:    res,
		})
	}
}
