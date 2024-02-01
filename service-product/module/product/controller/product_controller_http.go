package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service-product/dto"
	"service-product/module/product/usecase"
	"service-product/utils"
	"strconv"
)

type ProductControllerHTTP struct {
	service usecase.UsecaseInterface
}

func NewControllerProductHTTP(route *gin.Engine, service usecase.UsecaseInterface) {
	controller := &ProductControllerHTTP{service: service}

	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/product", controller.Create)
	groupRoute.GET("/products", controller.GetList)
	groupRoute.GET("/product/:id", controller.GetById)
	groupRoute.PUT("/product/:id", controller.Update)
}
func (controller *ProductControllerHTTP) Create(ctx *gin.Context) {

	var input dto.SchemaProduct
	ctx.ShouldBindJSON(&input)

	product, err := controller.service.Create(&input)

	if err.Error != nil {
		utils.APIResponse(ctx, dto.APIResponse{
			Code:    err.StatusCode,
			Message: err.Error.Error(),
		})
	} else {
		utils.APIResponse(ctx, dto.APIResponse{
			Message: "Success",
			Code:    http.StatusOK,
			Data:    product,
		})
	}

}

func (controller *ProductControllerHTTP) GetById(ctx *gin.Context) {

	ID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	res, err := controller.service.GetByID(ID)

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

func (controller *ProductControllerHTTP) GetList(ctx *gin.Context) {

	res, err := controller.service.GetList()

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

func (controller *ProductControllerHTTP) Update(ctx *gin.Context) {

	var input dto.UpdateStockRequest
	input.ProductId, _ = strconv.ParseUint(ctx.Param("id"), 10, 64)
	ctx.ShouldBindJSON(&input)

	_, err := controller.service.UpdateStock(input)

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
