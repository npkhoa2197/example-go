package category

import (
	"context"

	"github.com/lib/pq"

	"github.com/jinzhu/gorm"

	"github.com/npkhoa2197/example-go/domain"
)

// pgService implementer for Category service in Postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Validate unique constraint error return from pg. True if there is an error on unique, false otherwise
func notUnique(err error) bool {
	pqError, ok := err.(*pq.Error)

	return ok && pqError.Code == uniqueError
}

// Create implement Create for Category service
func (s *pgService) Create(_ context.Context, p *domain.Category) error {
	err := s.db.Create(p).Error

	// validate if the cateogry name existed
	if notUnique(err) {
		return ErrCategoryExisted
	}

	return err
}

// Update implement Update for Category service
func (s *pgService) Update(_ context.Context, p *domain.Category) (*domain.Category, error) {
	old := domain.Category{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	old.Name = p.Name

	err := s.db.Save(&old).Error

	// validate if the cateogry name existed
	if notUnique(err) {
		return nil, ErrCategoryExisted
	}

	return &old, err
}

// Find implement Find for Category service
func (s *pgService) Find(_ context.Context, p *domain.Category) (*domain.Category, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for Category service
func (s *pgService) FindAll(_ context.Context) ([]domain.Category, error) {
	res := []domain.Category{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for Category service
func (s *pgService) Delete(_ context.Context, p *domain.Category) error {
	old := domain.Category{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}

	if err := s.db.Where("category_id = ?", p.ID).Delete(&domain.Book{}).Error; err != nil {
		return err
	}

	return s.db.Delete(old).Error
}
