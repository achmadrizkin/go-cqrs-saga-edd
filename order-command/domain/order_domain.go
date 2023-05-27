package domain

import (
	"go-cqrs-saga-edd/order-command/model"

	"gorm.io/gorm"
)

type OrderRepo interface {
	CreateOrderRepo(order model.Order) (error, *gorm.DB)
}

type OrderUseCase interface {
	CreateOrderUseCase(order model.Order) (error, *gorm.DB)
}
