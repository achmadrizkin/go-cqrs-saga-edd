package main

import (
	"go-cqrs-saga-edd/product/config"
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

	productErrConsumerRepo := repo.NewProductErrConsumerRepo(ch)
	productAESRepo := repo.NewProductAESRepo()
	productRepo := repo.NewProductRepo(database)
	productErrPublisher := repo.NewProductErrPublisherRepo(ch)
	productErrConsumerUseCase := usecase.NewProductErrConsumerUseCase(productErrConsumerRepo, productAESRepo, productRepo, productErrPublisher)

	//
	productErrConsumerUseCase.ErrConsumerProductFromOrderUseCase(config.Config("E_PUBLISHER_ORDER_QUERY"), config.Config("NAME_EVENT_FAILED_PRODUCT_TO_ORDER_PUBLISHER"))
}
