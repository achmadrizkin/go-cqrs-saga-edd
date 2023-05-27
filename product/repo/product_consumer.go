package repo

import (
	"errors"
	"go-cqrs-saga-edd/product/domain"

	"github.com/streadway/amqp"
)

type productConsumerRepo struct {
	ch *amqp.Channel
}

// ConsumerProductFromOrderRepo implements domain.ProductConsumerRepo
func (p *productConsumerRepo) ConsumerProductFromOrderRepo(nameQueue string) (<-chan amqp.Delivery, error) {
	msgs, err := declareConsumer(p.ch, nameQueue)
	if err != nil {
		return nil, errors.New("failed to create consumer: " + err.Error())
	}

	return msgs, nil
}

func NewProductConsumerRepo(ch *amqp.Channel) domain.ProductConsumerRepo {
	return &productConsumerRepo{
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
