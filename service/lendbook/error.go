package lendbook

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound         = errNotFound{}
	ErrUnknown          = errUnknown{}
	ErrFieldIsRequired  = errFieldIsRequired{}
	ErrRecordNotFound   = errRecordNotFound{}
	ErrNotExisted       = errNotExisted{}
	ErrBookNotAvailable = errBookNotAvailable{}
)

const foreignKeyError = "23503"
const nullError = "23502"

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

type errFieldIsRequired struct{}

func (errFieldIsRequired) Error() string {
	return "book/user is required"
}

func (errFieldIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errNotExisted struct{}

func (errNotExisted) Error() string {
	return "book/user not existed"
}

func (errNotExisted) StatusCode() int {
	return http.StatusBadRequest
}

type errBookNotAvailable struct{}

func (errBookNotAvailable) Error() string {
	return "book not available"
}

func (errBookNotAvailable) StatusCode() int {
	return http.StatusBadRequest
}
