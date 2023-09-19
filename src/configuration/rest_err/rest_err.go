package rest_err

import "net/http"

type RestErr struct {
	Message string   `json: message`
	Err     string   `json: error`
	Code    int64    `json: code`
	Causes  []Causes `json: causes`
}

type Causes struct {
	Field   string `json: field`
	Message string `json: message`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message, err string, code int64, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Not_Found",
		Code:    http.StatusNotFound,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Internal_Server_Error",
		Code:    http.StatusInternalServerError,
	}
}
