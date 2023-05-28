package domain

import "github.com/streadway/amqp"

type OrderErrConsumerRepo interface {
	ConsumerErrFromOrderToProduct(nameQueue string) (<-chan amqp.Delivery, error)
}

type OrderErrConsumerUseCase interface {
	ConsumerErrFromOrderToProduct(nameQueue string) error
}
