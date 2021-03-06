package book

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"

	"github.com/npkhoa2197/example-go/domain"
)

// pgService implementer for Book service in Postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Validate foreign key constraint error return from pg. True if there is an error on FK, false otherwise
func keyNotExisted(err error) bool {
	pqError, ok := err.(*pq.Error)

	return ok && pqError.Code == foreignKeyError
}

// Create implement Create for Book service
func (s *pgService) Create(_ context.Context, p *domain.Book) error {
	err := s.db.Create(p).Error

	if keyNotExisted(err) {
		return ErrCategoryNotExisted
	}

	return err
}

// Update implement Update for Book service
func (s *pgService) Update(_ context.Context, p *domain.Book) (*domain.Book, error) {
	old := domain.Book{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	old.Name = p.Name
	old.CategoryID = p.CategoryID
	old.Author = p.Author
	old.Description = p.Description

	err := s.db.Save(&old).Error

	if keyNotExisted(err) {
		return &old, ErrCategoryNotExisted
	}

	return &old, err
}

// Find implement Find for Book service
func (s *pgService) Find(_ context.Context, p *domain.Book) (*domain.Book, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for Book service
func (s *pgService) FindAll(_ context.Context) ([]domain.Book, error) {
	res := []domain.Book{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for Book service
func (s *pgService) Delete(_ context.Context, p *domain.Book) error {
	old := domain.Book{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}
	return s.db.Delete(old).Error
}
