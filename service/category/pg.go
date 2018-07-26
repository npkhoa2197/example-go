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

// Create implement Create for Category service
func (s *pgService) Create(_ context.Context, p *domain.Category) error {
	err := s.db.Create(p).Error
	pqError, ok := err.(*pq.Error)

	if ok && pqError.Code == uniqueError {
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
	pqError, ok := err.(*pq.Error)

	if ok && pqError.Code == uniqueError {
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

	var books []domain.Book

	s.db.Where("category_id = ?", old.ID).Find(&books)

	for _, book := range books {
		s.db.Delete(book)
	}

	return s.db.Delete(old).Error
}
