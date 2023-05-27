package domain

import "go-cqrs-saga-edd/order-command/model"

type OrderRepo interface {
	CreateOrderRepo(model.Order) error
}

type OrderUseCase interface {
	CreateOrderUseCase(model.Order) error
}
