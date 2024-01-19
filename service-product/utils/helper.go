package utils

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	_ "image/jpeg"
	_ "image/png"
	"net"
	"reflect"
	"service-product/dto"

	"strings"
)

func SplitOrderQuery(order string) (string, string) {
	orderA := strings.Split(order, "|")
	if len(orderA) > 1 {
		if orderA[0] != "" && orderA[1] != "" {
			return orderA[0], orderA[1]
		}
	}

	return orderA[0], "ASC"
}

func MapStruct(source interface{}, dest interface{}) {
	valueOfSource := reflect.ValueOf(source)
	valueOfDest := reflect.ValueOf(dest)

	if valueOfSource.Kind() != reflect.Struct || valueOfDest.Kind() != reflect.Ptr {
		panic("Tipe data tidak valid")
	}

	typeOfSource := valueOfSource.Type()
	typeOfDest := valueOfDest.Elem().Type()

	for i := 0; i < typeOfSource.NumField(); i++ {
		fieldName := typeOfSource.Field(i).Name
		destField, found := typeOfDest.FieldByName(fieldName)
		if found {
			destValueField := valueOfDest.Elem().FieldByName(destField.Name)
			destValueField.Set(valueOfSource.FieldByName(fieldName))
		}
	}
}

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
