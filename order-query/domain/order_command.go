package domain

import (
	"context"
	"go-cqrs-saga-edd/order-query/model"
)

type OrderCommandRepo interface {
	CreateOrderProduct(ctx context.Context, orderProduct model.OrderProduct) error
}

type OrderCommandUseCase interface {
	CreateOrderProduct(ctx context.Context, orderProduct model.OrderProduct) error
}
