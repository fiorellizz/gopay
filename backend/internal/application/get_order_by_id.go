package application

import (
	"context"

	"github.com/fiorellizz/gopay/internal/domain"
)

type GetOrderByIDUseCase struct {
	repo domain.OrderRepository
}

func NewGetOrderByIDUseCase(repo domain.OrderRepository) *GetOrderByIDUseCase {
	return &GetOrderByIDUseCase{repo: repo}
}

func (uc *GetOrderByIDUseCase) Execute(ctx context.Context, id string) (*domain.Order, error) {
	return uc.repo.FindByID(ctx, id)
}