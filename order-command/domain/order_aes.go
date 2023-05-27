package domain

import "go-cqrs-saga-edd/order-command/model"

type OrderAESRepo interface {
	EncryptOrderAES(model.Order) ([]byte, error)
}
