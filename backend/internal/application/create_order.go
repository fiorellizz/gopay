package application

import (
	"context"
	"errors"
	"time"

	"github.com/fiorellizz/gopayflow/internal/domain"
)

type CreateOrderInput struct {
	Amount float64
}

type CreateOrderOutput struct {
	ID string
}

type CreateOrderUseCase struct {
	repo domain.OrderRepository
}

func NewCreateOrderUseCase(repo domain.OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		repo: repo,
	}
}

func (uc *CreateOrderUseCase) Execute(ctx context.Context, input CreateOrderInput) (*CreateOrderOutput, error) {

	if input.Amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	order := &domain.Order{
		Amount:    input.Amount,
		Status:    domain.StatusPending,
		CreatedAt: time.Now(),
	}

	err := uc.repo.Create(ctx, order)
	if err != nil {
		return nil, err
	}

	return &CreateOrderOutput{
		ID: order.ID,
	}, nil
}