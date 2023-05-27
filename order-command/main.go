package main

import (
	"fmt"
	"go-cqrs-saga-edd/order-command/config"
	"go-cqrs-saga-edd/order-command/db"
	"go-cqrs-saga-edd/order-command/model"
	pb "go-cqrs-saga-edd/order-command/proto"
	"go-cqrs-saga-edd/order-command/rabbitmq"
	"go-cqrs-saga-edd/order-command/repo"
	"go-cqrs-saga-edd/order-command/server"
	"go-cqrs-saga-edd/order-command/usecase"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Order Command Service Started")

	db := db.ConnectToMysql()
	if err := db.AutoMigrate(&model.Order{}); err != nil {
		log.Panic(err.Error())
	}

	rabbitMq := rabbitmq.ConnectionToRabbitMq()
	chRabbitMQ := rabbitmq.ConnectionToChannelRabbitMq(rabbitMq)

	startGRPCServer(db, chRabbitMQ)
}

func startGRPCServer(db *gorm.DB, chRabbitMQ *amqp.Channel) {
	s := grpc.NewServer()

	orderRepo := repo.NewOrderRepo(db)
	orderUseCase := usecase.NewOrderUseCase(orderRepo)
	orderAESRepo := repo.NewOrderAESRepo()
	orderPublisherRepo := repo.NewOrderPublisherRepo(chRabbitMQ)
	orderPublisherUseCase := usecase.NewOrderPublisherUseCase(orderPublisherRepo, orderAESRepo)

	pb.RegisterOrderServiceServer(s, &server.OrderServer{
		UnimplementedOrderServiceServer: pb.UnimplementedOrderServiceServer{},
		OrderUseCase:                    orderUseCase,
		OrderPublisherUseCase:           orderPublisherUseCase,
	})
	reflection.Register(s)

	// gRPC server
	lis, err := net.Listen("tcp", config.Config("GRPC_PORT"))
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	go func() {
		fmt.Println("Starting server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Menunggu hingga dihentikan dengan Ctrl + C
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Lakukan block hingga sinyal sudah didapatkan
	<-ch
	fmt.Println("Stopping the server..")
	s.Stop()
	fmt.Println("Stopping listener...")
	lis.Close()
	fmt.Println("End of Program")
}
