package book

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound           = errNotFound{}
	ErrUnknown            = errUnknown{}
	ErrNameIsRequired     = errNameIsRequired{}
	ErrRecordNotFound     = errRecordNotFound{}
	ErrCategoryNotExisted = errCategoryNotExisted{}
)

const foreignKeyError = "23503"

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
	return "book name is required"
}

func (errNameIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errCategoryNotExisted struct{}

func (errCategoryNotExisted) Error() string {
	return "category not existed"
}

func (errCategoryNotExisted) StatusCode() int {
	return http.StatusBadRequest
}
