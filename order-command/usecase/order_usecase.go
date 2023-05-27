package usecase

import (
	"go-cqrs-saga-edd/order-command/domain"
	"go-cqrs-saga-edd/order-command/model"

	"gorm.io/gorm"
)

type orderUseCase struct {
	orderRepo domain.OrderRepo
}

// CreateOrderUseCase implements domain.OrderUseCase
func (o *orderUseCase) CreateOrderUseCase(order model.Order, is_success int) (*gorm.DB, error) {
	// if is_success = 1, publish to product
	// if is_success = 0, err from product (rollback)
	// except 1,0 will error
	if is_success == 1 {
		return o.orderRepo.CreateOrderRepo(order)
	} else if is_success == 0 {
		return o.orderRepo.DeleteOrderRepo(order)
	} else {
		return nil, nil
	}
}

func NewOrderUseCase(orderRepo domain.OrderRepo) domain.OrderUseCase {
	return &orderUseCase{
		orderRepo: orderRepo,
	}
}
