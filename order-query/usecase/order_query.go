package usecase

import (
	"context"
	"go-cqrs-saga-edd/order-query/domain"
	"go-cqrs-saga-edd/order-query/model"
)

type orderQueryUseCase struct {
	orderQueryRepo domain.OrderQueryRepo
}

// GetOrderById implements domain.OrderQueryUseCase
func (o *orderQueryUseCase) GetOrderById(ctx context.Context, id string) (model.OrderProduct, error) {
	return o.orderQueryRepo.GetOrderById(ctx, id)
}

// GetOrderProductAll implements domain.OrderQueryUseCase
func (o *orderQueryUseCase) GetOrderProductAll(ctx context.Context) ([]model.OrderProduct, error) {
	return o.orderQueryRepo.GetOrderProductAll(ctx)
}

func NewOrderQueryUseCase(orderQueryRepo domain.OrderQueryRepo) domain.OrderQueryUseCase {
	return &orderQueryUseCase{orderQueryRepo: orderQueryRepo}
}
