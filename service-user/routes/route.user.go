package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	handlers "service-user/module/user/controller"
	repositorys "service-user/module/user/repository"
	services "service-user/module/user/usecase"
)

func InitUserRoutes(db *gorm.DB, route *gin.Engine) {

	UserRepository := repositorys.NewUserRepository(db)
	UserService := services.NewUserUsecase(UserRepository)

	controller := handlers.NewUserControllerHTTP(UserService)
	/**
	@description All User Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/user", controller.CreateUserHandler)
	groupRoute.GET("/user", controller.ResultsUserHandler)
	groupRoute.GET("/user/:id", controller.ResultUserHandler)
	groupRoute.DELETE("/user/:id", controller.DeleteUserHandler)
	groupRoute.PUT("/user/:id", controller.UpdateUserHandler)
}
