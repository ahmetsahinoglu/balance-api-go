package errors

import (
	"fmt"

	"github.com/ansel1/merry/v2"
	errs "github.com/pkg/errors"
)

func buildErrorResponse(err error) ErrorResponse {
	var errorResponse ErrorResponse
	if errs.As(err, &errorResponse) {
		return errorResponse
	}

	return InternalServerError("")
}

func New(msg string) error {
	return WithStack(fmt.Errorf("%s", msg))
}

func Wrap(err error, msg string) error {
	return merry.Wrap(err, merry.AppendMessage(msg))
}

func WithStack(err error) error {
	return merry.Wrap(err)
}

func getError(err error, errorResponse ErrorResponse) error {
	if errorResponse.nestedErr != nil {
		return errorResponse.nestedErr
	}

	return err
}
