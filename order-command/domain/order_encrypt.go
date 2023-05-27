package domain

import "go-cqrs-saga-edd/order-command/model"

type OrderEncryptRepo interface {
	EncryptOrderAES(model.Order) ([]byte, error)
}
