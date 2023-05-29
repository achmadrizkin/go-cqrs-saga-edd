package usecase

import (
	"go-cqrs-saga-edd/order-query/domain"
	"go-cqrs-saga-edd/order-query/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type orderCommandUseCase struct {
	orderCommandRepo domain.OrderCommandRepo
}

// CreateOrderProduct implements domain.OrderCommandUseCase
func (o *orderCommandUseCase) CreateOrderProduct(sc mongo.SessionContext, orderProduct model.OrderProduct) error {
	return o.orderCommandRepo.CreateOrderProduct(sc, orderProduct)
}

func NewOrderCommandUseCase(orderCommandRepo domain.OrderCommandRepo) domain.OrderCommandUseCase {
	return &orderCommandUseCase{orderCommandRepo}
}
