package repo

import (
	"errors"
	"go-cqrs-saga-edd/order-command/domain"

	"github.com/streadway/amqp"
)

type orderPublisherRepo struct {
	ch *amqp.Channel
}

// CreateOrderRepoPublisherToProduct implements domain.OrderPublisherRepo
func (o *orderPublisherRepo) CreateOrderRepoPublisherToProduct(encryptedOrder []byte, nameQueue string) error {
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

func NewOrderPublisherRepo(ch *amqp.Channel) domain.OrderPublisherRepo {
	return &orderPublisherRepo{
		ch: ch,
	}
}
