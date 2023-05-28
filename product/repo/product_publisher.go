package repo

import (
	"errors"
	"go-cqrs-saga-edd/product/domain"

	"github.com/streadway/amqp"
)

type productPublisherRepo struct {
	ch *amqp.Channel
}

// ProductErrPublisherFromProductToOrderQuery implements domain.ProductPublisherRepo
func (p *productPublisherRepo) ProductPublisherFromProductToOrderQuery(encryptedOrder []byte, nameQueue string) error {
	q, errDeclareQueue := declareQueue(p.ch, nameQueue)
	if errDeclareQueue != nil {
		return errors.New("err declare queue: " + errDeclareQueue.Error())
	}

	if errPublisher := publishQueue(p.ch, q, encryptedOrder); errPublisher != nil {
		return errors.New("error errPublish product to order: " + errPublisher.Error())
	}

	return nil
}

func NewProductPublisherRepo(ch *amqp.Channel) domain.ProductPublisherRepo {
	return &productPublisherRepo{ch}
}
