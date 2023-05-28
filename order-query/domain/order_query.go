package domain

import (
	"context"
	"go-cqrs-saga-edd/order-query/model"
)

type OrderQueryRepo interface {
	GetOrderById(ctx context.Context, id string) (model.OrderProduct, error)
}

type OrderQueryUseCase interface {
	GetOrderById(ctx context.Context, id string) (model.OrderProduct, error)
}
