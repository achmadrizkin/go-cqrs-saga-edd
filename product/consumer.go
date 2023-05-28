package main

import (
	"go-cqrs-saga-edd/order-command/config"
	"go-cqrs-saga-edd/product/db"
	"go-cqrs-saga-edd/product/model"
	"go-cqrs-saga-edd/product/rabbitmq"
	"go-cqrs-saga-edd/product/repo"
	"go-cqrs-saga-edd/product/usecase"
	"log"
)

func main() {
	// jika kode mengalami crash, nomor line akan ditampilkan
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Product Service Consumer Started")

	database := db.ConnectToMysql()
	if err := database.AutoMigrate(&model.Product{}); err != nil {
		log.Panic(err.Error())
	}

	rabbitMq := rabbitmq.ConnectionToRabbitMq()
	ch := rabbitmq.ConnectionToChannelRabbitMq(rabbitMq)

	productPublisherRepo := repo.NewProductPublisherRepo(ch)
	productErrPublisher := repo.NewProductErrPublisherRepo(ch)
	productConsumerRepo := repo.NewProductConsumerRepo(ch)
	productAESRepo := repo.NewProductAESRepo()
	productRepo := repo.NewProductRepo(database)
	productConsumerUseCase := usecase.NewProductConsumerUseCase(productConsumerRepo, productAESRepo, productRepo, productErrPublisher, productPublisherRepo)

	productConsumerUseCase.ConsumerProductFromOrderUseCase(config.Config("NAME_EVENT_SUCCESS_ORDER_TO_PRODUCT_CONSUMER"), config.Config("NAME_EVENT_FAILED_PRODUCT_TO_ORDER_PUBLISHER"), config.Config("NAME_EVENT_SUCCESS_PRODUCT_TO_ORDER_QUERY_PUBLISHER"))
}
