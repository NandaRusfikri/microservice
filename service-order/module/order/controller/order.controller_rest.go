package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service-order/dto"
	orders "service-order/module/order"
	"service-order/utils"
	"strconv"
)

type OrderControllerRestAPI struct {
	OrderService orders.OrderServiceInterface
}

func NewOrderControllerRestAPI(service orders.OrderServiceInterface, route *gin.Engine) {
	ctrl := OrderControllerRestAPI{OrderService: service}

	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/order", ctrl.Create)
	groupRoute.GET("/order", ctrl.GetById)
	groupRoute.GET("/order/:id", ctrl.GetList)
}

func (h *OrderControllerRestAPI) Create(ctx *gin.Context) {

	var input dto.SchemaOrder
	ctx.ShouldBindJSON(&input)

	_, err := h.OrderService.Create(&input)

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

func (h *OrderControllerRestAPI) GetById(ctx *gin.Context) {

	orderId, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	res, err := h.OrderService.GetById(orderId)

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

func (h *OrderControllerRestAPI) GetList(ctx *gin.Context) {

	res, err := h.OrderService.GetList()
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
