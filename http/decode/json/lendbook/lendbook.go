package lendbook

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/npkhoa2197/example-go/domain"
	lendbookEndpoint "github.com/npkhoa2197/example-go/endpoints/lendbook"
)

// FindRequest .
func FindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	lendbookRecordID, err := domain.UUIDFromString(chi.URLParam(r, "lendbook_id"))
	if err != nil {
		return nil, err
	}
	return lendbookEndpoint.FindRequest{LendBookRecordID: lendbookRecordID}, nil
}

// FindAllRequest .
func FindAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return lendbookEndpoint.FindAllRequest{}, nil
}

// CreateRequest .
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req lendbookEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// UpdateRequest .
func UpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	lendbookRecordID, err := domain.UUIDFromString(chi.URLParam(r, "lendbook_id"))
	if err != nil {
		return nil, err
	}

	var req lendbookEndpoint.UpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.LendBookRecord.ID = lendbookRecordID

	return req, nil
}

// DeleteRequest .
func DeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	lendbookRecordID, err := domain.UUIDFromString(chi.URLParam(r, "lendbook_id"))
	if err != nil {
		return nil, err
	}
	return lendbookEndpoint.DeleteRequest{LendBookRecordID: lendbookRecordID}, nil
}
