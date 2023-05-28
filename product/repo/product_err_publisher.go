package repo

import (
	"errors"
	"go-cqrs-saga-edd/product/domain"

	"github.com/streadway/amqp"
)

type productErrPublisherRepo struct {
	ch *amqp.Channel
}

// ProductErrPublisherFromProductToOrder implements domain.ProductErrPubsliher
func (p *productErrPublisherRepo) ProductErrPublisherFromProductToOrder(encryptedOrder []byte, nameQueue string) error {
	q, errDeclareQueue := declareQueue(p.ch, nameQueue)
	if errDeclareQueue != nil {
		return errors.New("err declare queue: " + errDeclareQueue.Error())
	}

	if errPublisher := publishQueue(p.ch, q, encryptedOrder); errPublisher != nil {
		return errors.New("error errPublish product to order: " + errPublisher.Error())
	}

	return nil
}

func publishQueue(ch *amqp.Channel, q amqp.Queue, dataJSON []byte) error {
	err := ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        dataJSON,
		},
	)

	return err
}

func declareQueue(ch *amqp.Channel, nameQueue string) (amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		nameQueue,
		false,
		false,
		false,
		false,
		nil,
	)

	return q, err
}

func NewProductErrPublisherRepo(ch *amqp.Channel) domain.ProductErrPubsliher {
	return &productErrPublisherRepo{
		ch: ch,
	}
}
