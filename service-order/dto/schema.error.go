package dto

type ResponseError struct {
	Error      error
	StatusCode int
}
