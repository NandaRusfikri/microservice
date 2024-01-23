package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service-order/helpers"
	"service-order/schemas"
	services "service-order/services/order"
	"strconv"
)

type OrderControllerRestAPI struct {
	OrderService services.OrderService
}

func NewOrderControllerRestAPI(service services.OrderService) *OrderControllerRestAPI {
	return &OrderControllerRestAPI{OrderService: service}
}

func (controller *OrderControllerRestAPI) InitOrderRoutes(route *gin.Engine) {

	/**
	@description All Handler Student
	*/
	//orderRepository := repositorys.NewOrderRepository(db)
	//orderService := services.NewOrderService(orderRepository)
	//orderController := NewOrderController(orderService)


	/**
	@description All Student Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/order", controller.CreateOrderHandler)
	groupRoute.GET("/order", controller.ResultOrderHandler)
	groupRoute.GET("/order/:id", controller.FindAllOrderHandler)
	groupRoute.DELETE("/order/:id", controller.DeleteOrderHandler)
	groupRoute.PUT("/order/:id", controller.UpdateOrderHandler)


}

func (h *OrderControllerRestAPI) CreateOrderHandler(ctx *gin.Context) {

	var input schemas.SchemaOrder
	ctx.ShouldBindJSON(&input)


	_, err := h.OrderService.CreateOrderService(&input)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Npm Order already exist", err.Code, http.MethodPost, nil)
		return
	case "error_02":
		helpers.APIResponse(ctx, "Create new Order account failed", err.Code, http.MethodPost, nil)
		return
	default:
		helpers.APIResponse(ctx, "Create new Order account successfully", http.StatusCreated, http.MethodPost, nil)
	}

}

func (h *OrderControllerRestAPI) DeleteOrderHandler(ctx *gin.Context) {

	var input schemas.SchemaOrder
	input.ID,_ = strconv.ParseInt(ctx.Param("id"), 10, 64)



	_, err := h.OrderService.DeleteOrderService(&input)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Order data is not exist or deleted", err.Code, http.MethodDelete, nil)
		return
	case "error_02":
		helpers.APIResponse(ctx, "Delete Order data failed", err.Code, http.MethodDelete, nil)
		return
	default:
		helpers.APIResponse(ctx, "Delete Order data successfully", http.StatusOK, http.MethodDelete, nil)
	}
}

func (h *OrderControllerRestAPI) ResultOrderHandler(ctx *gin.Context) {

	var input schemas.SchemaOrder
	input.ID,_ = strconv.ParseInt(ctx.Param("id"), 10, 64)



	res, err := h.OrderService.ResultOrderService(&input)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Order data is not exist or deleted", err.Code, http.MethodGet, nil)
		return
	default:
		helpers.APIResponse(ctx, "Result Order data successfully", http.StatusOK, http.MethodGet, res)
	}
}

func (h *OrderControllerRestAPI) FindAllOrderHandler(ctx *gin.Context) {

	res, err := h.OrderService.FindAllOrderService()

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Orders data is not exists", err.Code, http.MethodPost, nil)
	default:
		helpers.APIResponse(ctx, "Results Orders data successfully", http.StatusOK, http.MethodPost, res)
	}
}

func (h *OrderControllerRestAPI) UpdateOrderHandler(ctx *gin.Context) {

	var input schemas.SchemaOrder
	input.ID,_ = strconv.ParseInt(ctx.Param("id"), 10, 64)
	ctx.ShouldBindJSON(&input)



	_, err := h.OrderService.UpdateOrderService(&input)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Order data is not exist or deleted", http.StatusNotFound, http.MethodPost, nil)
	case "error_02":
		helpers.APIResponse(ctx, "Update Order data failed", http.StatusForbidden, http.MethodPost, nil)
	default:
		helpers.APIResponse(ctx, "Update Order data sucessfully", http.StatusOK, http.MethodPost, nil)
	}
}