package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"service-product/module/product/controller"
	"service-product/module/product/repository"
	"service-product/module/product/usecase"
)

func InitProductRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Student
	*/
	productRepository := repository.NewProductRepositorySQL(db)
	productService := usecase.NewServiceProduct(productRepository)
	productHandler := controller.NewControllerProductHTTP(productService)

	/**
	@description All Student Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/product", productHandler.Create)
	groupRoute.GET("/product", productHandler.ResultsProductHandler)
	groupRoute.GET("/product/:id", productHandler.ResultProductHandler)
	groupRoute.DELETE("/product/:id", productHandler.DeleteProductHandler)
	groupRoute.PUT("/product/:id", productHandler.UpdateProductHandler)

}
