package helpers

import (
	"encoding/json"
	"service-order/schemas"

	"github.com/sirupsen/logrus"

)

func Strigify(payload interface{}) []byte {
	response, _ := json.Marshal(payload)
	return response
}

func Parse(payload []byte) schemas.SchemaResponses {
	var jsonResponse schemas.SchemaResponses
	err := json.Unmarshal(payload, &jsonResponse)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return jsonResponse
}
