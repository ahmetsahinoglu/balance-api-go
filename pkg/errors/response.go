package errors

import (
	"net/http"
)

type ErrorResponse struct {
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Details   interface{} `json:"details,omitempty"`
	nestedErr error       `json:"-"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}

func (e ErrorResponse) StatusCode() int {
	return e.Status
}

func (e ErrorResponse) Err(err error) ErrorResponse {
	e.nestedErr = WithStack(err)
	return e
}

func InternalServerError(msg string) ErrorResponse {
	if msg == "" {
		msg = "We encountered an error while processing your request."
	}
	return ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: msg,
	}
}

func BadRequest(msg string) ErrorResponse {
	if msg == "" {
		msg = "Your request is in a bad format."
	}
	return ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: msg,
	}
}
