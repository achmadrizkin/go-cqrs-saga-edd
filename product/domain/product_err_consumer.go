package domain

import "github.com/streadway/amqp"

type ProductErrConsumerRepo interface {
	ErrConsumerProductFromOrderRepo(nameQueue string) (<-chan amqp.Delivery, error)
}

type ProductErrConsumerUseCase interface {
	ErrConsumerProductFromOrderUseCase(nameQueueConsumer string, nameQueueErrPublisher string) error
}
