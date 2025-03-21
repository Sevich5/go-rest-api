package application

import (
	"app/internal/domain"
	"errors"
	"net/http"
)

type Err struct {
	domain.Err
}

func NewAppError(message string) *Err {
	return NewAppErrorWithStatus(message, http.StatusInternalServerError)
}

func NewAppErrorWithStatus(message string, status int) *Err {
	err := &Err{Err: *domain.NewError(message)}
	err.StatusCode = status
	return err
}

func NewAppErrorFromErr(err error) *Err {
	var appErr *Err
	if errors.As(err, &appErr) {
		outErr := NewAppError(err.Error())
		outErr.StatusCode = appErr.StatusCode
		outErr.Parameters = appErr.Parameters
		return outErr
	}
	var domainErr *domain.Err
	if errors.As(err, &domainErr) {
		outErr := NewAppError(err.Error())
		outErr.StatusCode = domainErr.StatusCode
		outErr.Parameters = domainErr.Parameters
		return outErr
	}
	return NewAppError(err.Error())
}
