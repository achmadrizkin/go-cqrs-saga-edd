package main

import (
	"go-cqrs-saga-edd/order-command/config"
	"go-cqrs-saga-edd/order-command/db"
	"go-cqrs-saga-edd/order-command/model"
	"go-cqrs-saga-edd/order-command/rabbitmq"
	"go-cqrs-saga-edd/order-command/repo"
	"go-cqrs-saga-edd/order-command/usecase"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Order Command Service Error Consumer Started")

	database := db.ConnectToMysql()
	if err := database.AutoMigrate(&model.Order{}); err != nil {
		log.Panic(err.Error())
	}

	rabbitMq := rabbitmq.ConnectionToRabbitMq()
	ch := rabbitmq.ConnectionToChannelRabbitMq(rabbitMq)

	orderErrConsumer := repo.NewOrderErrConsumerRepo(ch)
	orderRepo := repo.NewOrderRepo(database)
	orderAESRepo := repo.NewOrderAESRepo()
	orderErrConsumerUseCase := usecase.NewOrderErrConsumerUseCase(orderErrConsumer, orderRepo, orderAESRepo)

	//
	orderErrConsumerUseCase.ConsumerErrFromOrderToProduct(config.Config("NAME_EVENT_FAILED_PRODUCT_TO_ORDER_CONSUMER"))
}
