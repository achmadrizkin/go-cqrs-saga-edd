package usecase

import (
	"go-cqrs-saga-edd/order-command/domain"
	"go-cqrs-saga-edd/order-command/model"
)

type orderUseCase struct {
	orderRepo domain.OrderRepo
}

// CreateOrderUseCase implements domain.OrderUseCase
func (o *orderUseCase) CreateOrderUseCase(order model.Order) error {
	return o.orderRepo.CreateOrderRepo(order)
}

func NewOrderUseCase(orderRepo domain.OrderRepo) domain.OrderUseCase {
	return &orderUseCase{
		orderRepo: orderRepo,
	}
}
