package dto

type SchemaError struct {
	StatusCode int
	Error      error
}

type SchemaErrorResponse struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Error      interface{} `json:"error"`
}
