package main

import (
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

	productConsumerRepo := repo.NewProductConsumerRepo(ch)
	productAESRepo := repo.NewProductAESRepo()
	productConsumerUseCase := usecase.NewProductConsumerUseCase(productConsumerRepo, productAESRepo)

	productConsumerUseCase.ConsumerProductFromOrderUseCase("eOrderToProduct")
}
