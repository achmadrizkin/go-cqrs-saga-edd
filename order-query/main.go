package main

import (
	"context"
	"fmt"
	"go-cqrs-saga-edd/order-query/config"
	"go-cqrs-saga-edd/order-query/mongodb"
	pb "go-cqrs-saga-edd/order-query/proto"
	"go-cqrs-saga-edd/order-query/repo"
	"go-cqrs-saga-edd/order-query/server"
	"go-cqrs-saga-edd/order-query/usecase"
	"log"
	"net"
	"os"
	"os/signal"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Order Query Service Started")
	ctx := context.TODO()

	client, err := mongodb.MongoDbConn(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Connected to MongoDB Success")

	var table = mongodb.MongoCollection("orderproduct", client)

	startGrpcServer(table, client)
}

func startGrpcServer(table *mongo.Collection, client *mongo.Client) {
	s := grpc.NewServer()

	orderCommandRepo := repo.NewOrderCommandRepo(table)
	orderCommandUseCase := usecase.NewOrderCommandUseCase(orderCommandRepo)

	pb.RegisterOrderCommandServiceServer(s, &server.OrderCommandServer{
		UnimplementedOrderCommandServiceServer: pb.UnimplementedOrderCommandServiceServer{},
		OrderCommandUseCase:                    orderCommandUseCase,
		Client:                                 client,
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
