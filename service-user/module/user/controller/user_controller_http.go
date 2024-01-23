package controller

import (
	"net/http"
	"service-user/dto"
	services "service-user/module/user/usecase"
	"service-user/utils"
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
	groupRoute.PUT("/user/:id", controller.UserUpdate)

}

func (h *controllerUser) UserCreate(ctx *gin.Context) {

	var input dto.SchemaUser
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(ctx, dto.APIResponse{
			Message: "request invalid " + err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	_, err := h.service.Create(&input)

	if err.Error != nil {
		utils.APIResponse(ctx, dto.APIResponse{
			Code:    err.StatusCode,
			Message: err.Error.Error(),
		})
	} else {
		utils.APIResponse(ctx, dto.APIResponse{
			Message: "Success",
			Code:    http.StatusOK,
		})
	}

}

func (h *controllerUser) CutBalanceHandler(ctx *gin.Context) {

	var input dto.CutBalanceRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(ctx, dto.APIResponse{
			Message: "request invalid " + err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	user, err := h.service.CutBalance(input)
	if err.Error != nil {
		utils.APIResponse(ctx, dto.APIResponse{
			Code:    err.StatusCode,
			Message: err.Error.Error(),
		})
	} else {
		utils.APIResponse(ctx, dto.APIResponse{
			Message: "Success",
			Code:    http.StatusOK,
			Data:    user,
		})
	}
}

func (h *controllerUser) UserList(ctx *gin.Context) {

	res, err := h.service.GetList()

	if err.Error != nil {
		utils.APIResponse(ctx, dto.APIResponse{
			Code:    err.StatusCode,
			Message: err.Error.Error(),
		})
	} else {
		utils.APIResponse(ctx, dto.APIResponse{
			Message: "Success",
			Code:    http.StatusOK,
			Data:    res,
		})
	}
}

func (h *controllerUser) UserUpdate(ctx *gin.Context) {

	var input dto.SchemaUser
	input.ID, _ = strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(ctx, dto.APIResponse{
			Message: "request invalid " + err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	_, err := h.service.Update(&input)

	if err.Error != nil {
		utils.APIResponse(ctx, dto.APIResponse{
			Code:    err.StatusCode,
			Message: err.Error.Error(),
		})
	} else {
		utils.APIResponse(ctx, dto.APIResponse{
			Message: "Success",
			Code:    http.StatusOK,
		})
	}
}

func (h *controllerUser) UserDetail(ctx *gin.Context) {

	userID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	res, err := h.service.GetById(userID)

	if err.Error != nil {
		utils.APIResponse(ctx, dto.APIResponse{
			Code:    err.StatusCode,
			Message: err.Error.Error(),
		})
	} else {
		utils.APIResponse(ctx, dto.APIResponse{
			Message: "Success",
			Code:    http.StatusOK,
			Data:    res,
		})
	}
}
