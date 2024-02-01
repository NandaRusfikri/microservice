package dto

type ResponseError struct {
	StatusCode int
	Error      error
}

type APIResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Count   int64       `json:"count,omitempty"`
	Data    interface{} `json:"data"`
}
