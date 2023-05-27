package domain

import "github.com/streadway/amqp"

type ProductConsumerRepo interface {
	ConsumerProductFromOrderRepo(nameQueue string) (<-chan amqp.Delivery, error)
}

type ProductConsumerUseCase interface {
	ConsumerProductFromOrderUseCase(nameQueue string) error
}
