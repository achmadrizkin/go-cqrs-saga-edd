package domain

import (
	"context"

	"github.com/streadway/amqp"
)

type OrderQueryConsumerRepo interface {
	ConsumerOrderQuerConsumerRepo(nameQueue string) (<-chan amqp.Delivery, error)
}

type OrderQueryConsumerUseCase interface {
	ConsumerOrderQueryConsumerRepo(ctx context.Context, nameQueueConsumer string) error
}
