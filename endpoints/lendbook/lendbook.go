package lendbook

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"

	"github.com/npkhoa2197/example-go/domain"
	"github.com/npkhoa2197/example-go/service"
)

// CreateData data for CreateLendBookRecord
type CreateData struct {
	UserID domain.UUID `json:"user_id"`
	BookID domain.UUID `json:"book_id"`
	From   time.Time   `json:"from"`
	To     time.Time   `json:"to"`
}

// CreateRequest request struct for CreateLendBookRecord
type CreateRequest struct {
	LendBookRecord CreateData `json:"lend_book_record"`
}

// CreateResponse response struct for CreateLendBookRecord
type CreateResponse struct {
	LendBookRecord domain.LendBookRecord `json:"lend_book_record"`
}

// StatusCode customstatus code for success create LendBookRecord
func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeCreateEndpoint make endpoint for create a LendBookRecord
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req            = request.(CreateRequest)
			lendBookRecord = &domain.LendBookRecord{
				UserID: req.LendBookRecord.UserID,
				BookID: req.LendBookRecord.BookID,
				From:   req.LendBookRecord.From,
				To:     req.LendBookRecord.To,
			}
		)

		err := s.LendBookService.Create(ctx, lendBookRecord)
		if err != nil {
			return nil, err
		}

		return CreateResponse{LendBookRecord: *lendBookRecord}, nil
	}
}

// FindRequest request struct for Find a LendBookRecord
type FindRequest struct {
	LendBookRecordID domain.UUID
}

// FindResponse response struct for Find a LendBookRecord
type FindResponse struct {
	LendBookRecord *domain.LendBookRecord `json:"lend_book_record"`
}

// MakeFindEndPoint make endpoint for find LendBookRecord
func MakeFindEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var lendBookRecordFind domain.LendBookRecord
		req := request.(FindRequest)
		lendBookRecordFind.ID = req.LendBookRecordID

		lendBookRecord, err := s.LendBookService.Find(ctx, &lendBookRecordFind)
		if err != nil {
			return nil, err
		}
		return FindResponse{LendBookRecord: lendBookRecord}, nil
	}
}

// FindAllRequest request struct for FindAll LendBookRecord
type FindAllRequest struct{}

// FindAllResponse request struct for find all LendBookRecord
type FindAllResponse struct {
	LendBookRecords []domain.LendBookRecord `json:"lend_book_records"`
}

// MakeFindAllEndpoint make endpoint for find all LendBookRecord
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(FindAllRequest)
		lendBookRecords, err := s.LendBookService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{LendBookRecords: lendBookRecords}, nil
	}
}

// UpdateData data for Create
type UpdateData struct {
	ID     domain.UUID `json:"-"`
	UserID domain.UUID `json:"user_id"`
	BookID domain.UUID `json:"book_id"`
	From   time.Time   `json:"from"`
	To     time.Time   `json:"to"`
}

// UpdateRequest request struct for update
type UpdateRequest struct {
	LendBookRecord UpdateData `json:"lend_book_record"`
}

// UpdateResponse response struct for Create
type UpdateResponse struct {
	LendBookRecord domain.LendBookRecord `json:"lend_book_record"`
}

// MakeUpdateEndpoint make endpoint for update a LendBookRecord
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req            = request.(UpdateRequest)
			lendBookRecord = domain.LendBookRecord{
				Model:  domain.Model{ID: req.LendBookRecord.ID},
				UserID: req.LendBookRecord.UserID,
				BookID: req.LendBookRecord.BookID,
				From:   req.LendBookRecord.From,
				To:     req.LendBookRecord.To,
			}
		)

		res, err := s.LendBookService.Update(ctx, &lendBookRecord)
		if err != nil {
			return nil, err
		}

		return UpdateResponse{LendBookRecord: *res}, nil
	}
}

// DeleteRequest request struct for delete a LendBookRecord
type DeleteRequest struct {
	LendBookRecordID domain.UUID
}

// DeleteResponse response struct for Find a LendBookRecord
type DeleteResponse struct {
	Status string `json:"status"`
}

// MakeDeleteEndpoint make endpoint for update a LendBookRecord
func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			lendBookRecordFind = domain.LendBookRecord{}
			req                = request.(DeleteRequest)
		)
		lendBookRecordFind.ID = req.LendBookRecordID

		err := s.LendBookService.Delete(ctx, &lendBookRecordFind)
		if err != nil {
			return nil, err
		}

		return DeleteResponse{"success"}, nil
	}
}
