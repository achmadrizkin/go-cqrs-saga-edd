package domain

import "go-cqrs-saga-edd/order-command/model"

type OrderPublisherRepo interface {
	// encrypted order means -> already encrypted and string type
	CreateOrderRepoPublisherToProduct(encryptedOrder []byte, nameQueue string) error
}

type OrderPublisherUseCase interface {
	CreateOrderUseCasePublisherToProduct(order model.Order, nameQueue string) error
}
