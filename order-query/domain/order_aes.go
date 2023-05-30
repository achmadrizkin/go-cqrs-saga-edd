package domain

import "go-cqrs-saga-edd/order-query/model"

type OrderAESRepo interface {
	DecryptOrderProductAES(message []byte) (model.OrderProduct, error)
}
