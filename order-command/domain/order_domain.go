package domain

import (
	"go-cqrs-saga-edd/order-command/model"

	"gorm.io/gorm"
)

type OrderRepo interface {
	CreateOrderRepo(order model.Order) (*gorm.DB, error)
	DeleteOrderRepo(order model.Order) (*gorm.DB, error)
}

type OrderUseCase interface {
	CreateOrderUseCase(order model.Order, is_success int) (*gorm.DB, error)
}
