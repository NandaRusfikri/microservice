package util

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"service-order/constant"
	"service-order/schemas"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	_ "image/jpeg"
	_ "image/png"

	"net/http"
	"net/http/httptest"
	"os"
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
func SplitBasicAuthBase64(token string) (string, string) {

	var decodedByte, _ = base64.StdEncoding.DecodeString(token)
	var decodedString = string(decodedByte)

	orderA := strings.Split(decodedString, ":")
	if len(orderA) > 1 {
		if orderA[0] != "" && orderA[1] != "" {
			return orderA[0], orderA[1]
		}
	}

	return orderA[0], ""
}

func ExcludeCreateSaveDB(columns ...string) []string {
	columnOmit := []string{"created_at", "create_user"}
	if len(columns) > 0 {
		columnOmit = append(columnOmit, columns...)
	}
	return columnOmit
}

func ExcludeCreateDeleteSaveDB(columns ...string) []string {
	columnOmit := []string{"created_at", "create_user", "deleted_at", "deleted_user"}
	if len(columns) > 0 {
		columnOmit = append(columnOmit, columns...)
	}
	return columnOmit
}

func ExcludeDeleteSaveDB(columns ...string) []string {
	columnOmit := []string{"deleted_at", "deleted_user"}
	if len(columns) > 0 {
		columnOmit = append(columnOmit, columns...)
	}
	return columnOmit
}

func HashPassword(password string) (string, error) {
	sha := ""

	h := hmac.New(sha256.New, []byte(constant.KEY_PASSWORD))
	h.Write([]byte(password))
	sha = hex.EncodeToString(h.Sum(nil))

	return sha, nil
}

func DeleteFile(data string) error {
	err := os.Remove(data)
	if err != nil {
		log.Println("err_remove_file", err)
	}
	return err
}
func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}





func HttpTestRequest(method, url string, payload []byte) (*httptest.ResponseRecorder, *http.Request, error) {

	request := make(chan *http.Request, 1)
	errors := make(chan error, 1)

	if binary.Size(payload) > 0 {
		req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
		request <- req
		errors <- err
	} else {
		req, err := http.NewRequest(method, url, nil)
		request <- req
		errors <- err
	}

	rr := httptest.NewRecorder()

	return rr, <-request, <-errors
}
//func APIResponse(ctx *gin.Context, Message string, StatusCode int, Count *int64, Data interface{}) {
//
//	jsonResponse := schemas.SchemaResponses{
//		StatusCode:    StatusCode,
//		//Method:  Method,
//		Count:Count,
//		Message: Message,
//		Data:    Data,
//		//Items:   Items,
//	}
//
//	if StatusCode >= 400 {
//		ctx.AbortWithStatusJSON(StatusCode, jsonResponse)
//	} else {
//		ctx.JSON(StatusCode, jsonResponse)
//	}
//}

func ValidatorErrorResponse(ctx *gin.Context, StatusCode int, Method string, Error interface{}) {
	errResponse := schemas.SchemaErrorResponse{
		StatusCode: StatusCode,
		Method:     Method,
		Error:      Error,
	}

	ctx.AbortWithStatusJSON(StatusCode, errResponse)
}


