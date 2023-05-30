package repo

import (
	"errors"
	"go-cqrs-saga-edd/product/domain"

	"github.com/streadway/amqp"
)

type productErrConsumerRepo struct {
	ch *amqp.Channel
}

// ErrConsumerProductFromOrderRepo implements domain.ProductErrConsumerRepo
func (p *productErrConsumerRepo) ErrConsumerProductFromOrderRepo(nameQueue string) (<-chan amqp.Delivery, error) {
	msgs, err := declareConsumer(p.ch, nameQueue)
	if err != nil {
		return nil, errors.New("failed to create consumer: " + err.Error())
	}

	return msgs, nil
}

func NewProductErrConsumerRepo(ch *amqp.Channel) domain.ProductErrConsumerRepo {
	return &productErrConsumerRepo{
		ch: ch,
	}
}
