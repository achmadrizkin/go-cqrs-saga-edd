package domain

import "go-cqrs-saga-edd/product/model"

type ProductAESRepo interface {
	DecryptProductAES(message []byte) (model.Order, error)
	EncryptOrderProductAES(orderProduct model.OrderProduct) ([]byte, error)
}
