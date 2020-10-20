package errors

import "net/http"

var (
	ErrorService errorServiceInterface = &errorRestService{}
)

type errorRestService struct{}

type errorServiceInterface interface {
	BadRequestErr(string) *RestErr
	InternalServerError(string) *RestErr
	NotFound(string) *RestErr
}

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func (e *errorRestService) BadRequestErr(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func (e *errorRestService) InternalServerError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

func (e *errorRestService) NotFound(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}
