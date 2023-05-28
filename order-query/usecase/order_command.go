package usecase

import (
	"context"
	"go-cqrs-saga-edd/order-query/domain"
	"go-cqrs-saga-edd/order-query/model"
)

type orderCommandUseCase struct {
	orderCommandRepo domain.OrderCommandRepo
}

// CreateOrderProduct implements domain.OrderCommandUseCase
func (o *orderCommandUseCase) CreateOrderProduct(ctx context.Context, orderProduct model.OrderProduct) error {
	return o.orderCommandRepo.CreateOrderProduct(ctx, orderProduct)
}

func NewOrderCommandUseCase(orderCommandRepo domain.OrderCommandRepo) domain.OrderCommandUseCase {
	return &orderCommandUseCase{orderCommandRepo}
}
