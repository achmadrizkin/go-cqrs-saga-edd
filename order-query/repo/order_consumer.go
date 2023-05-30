package repo

import (
	"errors"
	"go-cqrs-saga-edd/order-query/domain"

	"github.com/streadway/amqp"
)

type orderConsumerRepo struct {
	ch *amqp.Channel
}

// ConsumerOrderQuerConsumerRepo implements domain.OrderQueryConsumerRepo
func (o *orderConsumerRepo) ConsumerOrderQuerConsumerRepo(nameQueue string) (<-chan amqp.Delivery, error) {
	msgs, err := declareConsumer(o.ch, nameQueue)
	if err != nil {
		return nil, errors.New("failed to create consumer: " + err.Error())
	}

	return msgs, nil
}

func NewOrderConsumerRepo(ch *amqp.Channel) domain.OrderQueryConsumerRepo {
	return &orderConsumerRepo{
		ch: ch,
	}
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
