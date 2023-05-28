package repo

import (
	"errors"
	"go-cqrs-saga-edd/order-command/domain"

	"github.com/streadway/amqp"
)

type orderErrConsumerRepo struct {
	ch *amqp.Channel
}

// ConsumerErrFromOrderToProduct implements domain.OrderErrConsumerRepo
func (o *orderErrConsumerRepo) ConsumerErrFromOrderToProduct(nameQueue string) (<-chan amqp.Delivery, error) {
	msgs, err := declareConsumer(o.ch, nameQueue)
	if err != nil {
		return nil, errors.New("failed to create consumer: " + err.Error())
	}

	return msgs, nil
}

func declareConsumer(ch *amqp.Channel, nameConsumer string) (<-chan amqp.Delivery, error) {
	msgs, err := ch.Consume(
		nameConsumer,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	return msgs, err
}

func NewOrderErrConsumerRepo(ch *amqp.Channel) domain.OrderErrConsumerRepo {
	return &orderErrConsumerRepo{ch: ch}
}
