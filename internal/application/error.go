package application

import (
	"app/internal/domain"
	"errors"
	"net/http"
)

type Err struct {
	domain.Err
}

func NewError(message string) *Err {
	err := &Err{Err: *domain.NewError(message)}
	err.StatusCode = http.StatusInternalServerError
	return err
}

func NewErrorFromErr(err error) *Err {
	var appErr *Err
	if errors.As(err, &appErr) {
		outErr := NewError(err.Error())
		outErr.StatusCode = appErr.StatusCode
		outErr.Parameters = appErr.Parameters
		return outErr
	}
	var domainErr *domain.Err
	if errors.As(err, &domainErr) {
		outErr := NewError(err.Error())
		outErr.StatusCode = domainErr.StatusCode
		outErr.Parameters = domainErr.Parameters
		return outErr
	}
	return NewError(err.Error())
}
