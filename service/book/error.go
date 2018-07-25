package book

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound            = errNotFound{}
	ErrUnknown             = errUnknown{}
	ErrNameIsRequired      = errNameIsRequired{}
	ErrDescriptionRequired = errDescriptionRequired{}
	ErrRecordNotFound      = errRecordNotFound{}
	ErrCategoryNotExisted  = errCategoryNotExisted{}
	ErrNameLength          = errNameLength{}
	ErrDescriptionLength   = errDescriptionLength{}
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

type errNameLength struct{}

func (errNameLength) Error() string {
	return "name must be longer than 5 characters"
}

func (errNameLength) StatusCode() int {
	return http.StatusBadRequest
}

type errDescriptionLength struct{}

func (errDescriptionLength) Error() string {
	return "description must be longer than 5 characters"
}

func (errDescriptionLength) StatusCode() int {
	return http.StatusBadRequest
}

type errDescriptionRequired struct{}

func (errDescriptionRequired) Error() string {
	return "description is required "
}

func (errDescriptionRequired) StatusCode() int {
	return http.StatusBadRequest
}
