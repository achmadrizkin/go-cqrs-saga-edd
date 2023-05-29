package domain

import (
	"go-cqrs-saga-edd/order-query/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type OrderCommandRepo interface {
	CreateOrderProduct(sc mongo.SessionContext, orderProduct model.OrderProduct) error
}

type OrderCommandUseCase interface {
	CreateOrderProduct(sc mongo.SessionContext, orderProduct model.OrderProduct) error
}
