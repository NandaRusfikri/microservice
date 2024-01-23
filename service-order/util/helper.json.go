package util

import (
	"encoding/json"
)

func Strigify(payload interface{}) []byte {
	response, _ := json.Marshal(payload)
	return response
}

