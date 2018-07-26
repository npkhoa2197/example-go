package category

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound        = errNotFound{}
	ErrUnknown         = errUnknown{}
	ErrNameIsRequired  = errNameIsRequired{}
	ErrRecordNotFound  = errRecordNotFound{}
	ErrCategoryExisted = errCategoryExisted{}
	ErrNameLength      = errNameLength{}
)

const uniqueError = "23505"

type errNotFound struct{}

func (errNotFound) Error() string {
	return "record not found"
}
func (errNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errUnknown struct{}

func (errUnknown) Error() string {
	return "unknown error"
}
func (errUnknown) StatusCode() int {
	return http.StatusBadRequest
}

type errRecordNotFound struct{}

func (errRecordNotFound) Error() string {
	return "client record not found"
}
func (errRecordNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errNameIsRequired struct{}

func (errNameIsRequired) Error() string {
	return "category name is required"
}

func (errNameIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errCategoryExisted struct{}

func (errCategoryExisted) Error() string {
	return "category name existed"
}

func (errCategoryExisted) StatusCode() int {
	return http.StatusBadRequest
}

type errNameLength struct{}

func (errNameLength) Error() string {
	return "category name must be longer than 5 characters"
}

func (errNameLength) StatusCode() int {
	return http.StatusBadRequest
}
