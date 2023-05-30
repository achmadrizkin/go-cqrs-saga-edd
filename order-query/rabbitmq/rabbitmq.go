package rabbitmq

import (
	"fmt"
	"go-cqrs-saga-edd/order-query/config"
	"log"

	"github.com/streadway/amqp"
)

func ConnectionToChannelRabbitMq(conn *amqp.Connection) (ch *amqp.Channel) {
	ch, err := conn.Channel()
	if err != nil {
		log.Println("internal server error (channel): ", err)
	}
	log.Println("success connected to ch rabbitmq")

	return ch
}

func ConnectionToRabbitMq() *amqp.Connection {
	rabbitMQServer := fmt.Sprintf("amqp://%s:%s@%s:%s/", config.Config("RABBITMQ_USER"), config.Config("RABBITMQ_PASSWORD"), config.Config("RABBITMQ_HOST"), config.Config("RABBITMQ_PORT"))
	conn, err := amqp.Dial(rabbitMQServer)
	if err != nil {
		log.Println("internal server error (connection): ", err)
	}

	log.Println("success connected to rabbitmq")
	return conn
}
