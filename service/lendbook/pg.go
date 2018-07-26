package lendbook

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"

	"github.com/npkhoa2197/example-go/domain"
)

// pgService implmenter for LendBookRecord serivce in Postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for LendBookRecord service
func (s *pgService) Create(_ context.Context, p *domain.LendBookRecord) error {
	if err := s.db.Where("book_id = ?", p.BookID).Find(&domain.LendBookRecord{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			createError := s.db.Create(p).Error
			pqError, ok := createError.(*pq.Error)

			if ok {
				if pqError.Code == foreignKeyError {
					return ErrNotExisted
				}
			}

			return createError
		}
		return err
	}
	return ErrBookNotAvailable
}

// Update implement Update for LendBookRecord service
func (s *pgService) Update(_ context.Context, p *domain.LendBookRecord) (*domain.LendBookRecord, error) {
	old := domain.LendBookRecord{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	if err := s.db.Where("book_id = ?", p.BookID).Find(&domain.LendBookRecord{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {

			old.UserID = p.UserID
			old.BookID = p.BookID
			old.From = p.From
			old.To = p.To

			updateError := s.db.Save(&old).Error
			pqError, ok := updateError.(*pq.Error)

			if ok {
				if pqError.Code == foreignKeyError {
					return nil, ErrNotExisted
				}
			}
			return nil, updateError
		}
		return nil, err
	}
	return nil, ErrBookNotAvailable
}

// Find implement Find for LendBookRecord service
func (s *pgService) Find(_ context.Context, p *domain.LendBookRecord) (*domain.LendBookRecord, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for LendBookRecord service
func (s *pgService) FindAll(_ context.Context) ([]domain.LendBookRecord, error) {
	res := []domain.LendBookRecord{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for LendBookRecord service
func (s *pgService) Delete(_ context.Context, p *domain.LendBookRecord) error {
	old := domain.LendBookRecord{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}

	return s.db.Delete(old).Error
}
