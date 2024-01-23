package utils

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net"
	"service-order/dto"
)

func APIResponse(ctx *gin.Context, options ...dto.APIResponse) {

	var opt dto.APIResponse
	if len(options) > 0 {
		opt = options[0]
	}

	jsonResponse := dto.APIResponse{
		Message: opt.Message,
		Code:    opt.Code,
		Count:   opt.Count,
		Data:    opt.Data,
	}

	if jsonResponse.Code >= 400 {
		ctx.AbortWithStatusJSON(jsonResponse.Code, jsonResponse)
	} else {
		ctx.JSON(jsonResponse.Code, jsonResponse)
	}
}

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}
