package erorrs

import "net/http"

type RestErr struct {
	Message string
	Status  int
	Error   string
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: "Invalid Json Body",
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}

}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: "Invalid Json Body",
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}

}
