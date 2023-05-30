package repo

import (
	"errors"
	"go-cqrs-saga-edd/order-query/domain"

	"github.com/streadway/amqp"
)

type orderErrPublisherRepo struct {
	ch *amqp.Channel
}

// CreateErrOrderQueryPublisherToProductRepo implements domain.OrderErrPublisherRepo
func (o *orderErrPublisherRepo) CreateErrOrderQueryPublisherToProductRepo(encryptedOrder []byte, nameQueue string) error {
	q, errDeclareQueue := declareQueue(o.ch, nameQueue)
	if errDeclareQueue != nil {
		return errors.New("err declare queue: " + errDeclareQueue.Error())
	}

	if errPublisher := publishQueue(o.ch, q, encryptedOrder); errPublisher != nil {
		return errors.New("err publish order to product: " + errPublisher.Error())
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

func NewOrderErrPublisherRepo(ch *amqp.Channel) domain.OrderErrPublisherRepo {
	return &orderErrPublisherRepo{
		ch: ch,
	}
}
