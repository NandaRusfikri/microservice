package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"service-order/constant"
	dto2 "service-order/dto"
	"service-order/utils"
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
		"service_name": dto2.CfgApp.ServiceName,
		"author":       constant.AUTHOR,
		"version":      constant.SERVICE_VERSION,
		"time_now":     time.Now(),
		"rest_api":     fmt.Sprintf(":%v", dto2.CfgApp.RestPort),
	}
	utils.APIResponse(ctx, dto2.APIResponse{
		Code:    200,
		Message: "success",
		Data:    jsonData,
	})
}
