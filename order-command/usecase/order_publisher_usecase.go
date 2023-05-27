package usecase

import (
	"go-cqrs-saga-edd/order-command/domain"
	"go-cqrs-saga-edd/order-command/model"
)

type orderPublisherUseCase struct {
	orderPublisherRepo domain.OrderPublisherRepo
	orderEncryptRepo   domain.OrderAESRepo
}

// CreateOrderUseCasePublisherToProduct implements domain.OrderPublisherUseCase
func (o *orderPublisherUseCase) CreateOrderUseCasePublisherToProduct(order model.Order, nameQueue string) error {
	// encrypt the data first before sending
	encryptedOrderData, errEncryptedAES := o.orderEncryptRepo.EncryptOrderAES(order)
	if errEncryptedAES != nil {
		return errEncryptedAES
	}

	if errPublisher := o.orderPublisherRepo.CreateOrderRepoPublisherToProduct(encryptedOrderData, nameQueue); errPublisher != nil {
		return errPublisher
	}

	return nil
}

func NewOrderPublisherUseCase(orderPublisherRepo domain.OrderPublisherRepo, orderEncryptRepo domain.OrderAESRepo) domain.OrderPublisherUseCase {
	return &orderPublisherUseCase{
		orderPublisherRepo: orderPublisherRepo,
		orderEncryptRepo:   orderEncryptRepo,
	}
}
