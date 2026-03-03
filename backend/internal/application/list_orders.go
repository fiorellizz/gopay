package application

import (
	"context"

	"github.com/fiorellizz/gopayflow/internal/domain"
)

type ListOrdersUseCase struct {
	repo domain.OrderRepository
}

func NewListOrdersUseCase(repo domain.OrderRepository) *ListOrdersUseCase {
	return &ListOrdersUseCase{repo: repo}
}

func (uc *ListOrdersUseCase) Execute(ctx context.Context) ([]*domain.Order, error) {
	return uc.repo.FindAll(ctx)
}