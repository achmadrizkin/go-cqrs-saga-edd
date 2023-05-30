package main

import (
	"context"
	"go-cqrs-saga-edd/order-query/config"
	"go-cqrs-saga-edd/order-query/mongodb"
	"go-cqrs-saga-edd/order-query/rabbitmq"
	"go-cqrs-saga-edd/order-query/repo"
	"go-cqrs-saga-edd/order-query/usecase"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Order Query Service  Consumer Started")
	ctx := context.TODO()

	client, err := mongodb.MongoDbConn(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Connected to MongoDB Success")

	var table = mongodb.MongoCollection("orderproduct", client)

	rabbitMq := rabbitmq.ConnectionToRabbitMq()
	ch := rabbitmq.ConnectionToChannelRabbitMq(rabbitMq)

	orderAESRepo := repo.NewOrderAESRepo()
	orderConsumerRepo := repo.NewOrderConsumerRepo(ch)
	orderCommandRepo := repo.NewOrderCommandRepo(table)
	orderConsumerUseCase := usecase.NewOrderQueryConsumerUseCase(orderConsumerRepo, orderAESRepo, orderCommandRepo, client)

	orderConsumerUseCase.ConsumerOrderQueryConsumerRepo(ctx, config.Config("NAME_EVENT_SUCCESS_PRODUCT_TO_ORDER_QUERY_PUBLISHER"))
}
