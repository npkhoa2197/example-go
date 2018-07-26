package lendbook

import (
	"context"

	"github.com/npkhoa2197/example-go/domain"
)

// Service interface for lend book
type Service interface {
	Create(ctx context.Context, p *domain.LendBookRecord) error
	Update(ctx context.Context, p *domain.LendBookRecord) (*domain.LendBookRecord, error)
	Find(ctx context.Context, p *domain.LendBookRecord) (*domain.LendBookRecord, error)
	FindAll(ctx context.Context) ([]domain.LendBookRecord, error)
	Delete(ctx context.Context, p *domain.LendBookRecord) error
}
