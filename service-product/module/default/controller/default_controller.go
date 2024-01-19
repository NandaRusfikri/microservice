package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"service-product/constant"
	"service-product/dto"
	"service-product/utils"
	"time"
)

type DefaultController struct {
}

func InitDefaultController(g *gin.Engine) {
	handler := &DefaultController{}

	g.GET("/", handler.MainPage)
}

func (c *DefaultController) MainPage(ctx *gin.Context) {
	jsonData := map[string]interface{}{
		"service_name": dto.CfgApp.ServiceName,
		"author":       constant.AUTHOR,
		"version":      constant.SERVICE_VERSION,
		"time_now":     time.Now(),
		"rest_api":     fmt.Sprintf(":%v", dto.CfgApp.RestPort),
	}
	utils.APIResponse(ctx, dto.APIResponse{
		Code:    200,
		Message: "success",
		Data:    jsonData,
	})
}
