package main

import (
	"fmt"
	"go-cqrs-saga-edd/product/config"
	"go-cqrs-saga-edd/product/db"
	"go-cqrs-saga-edd/product/model"
	pb "go-cqrs-saga-edd/product/proto"
	"go-cqrs-saga-edd/product/repo"
	"go-cqrs-saga-edd/product/server"
	"go-cqrs-saga-edd/product/usecase"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

func main() {
	// jika kode mengalami crash, nomor line akan ditampilkan
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Product Service Started")

	db := db.ConnectToMysql()
	if err := db.AutoMigrate(&model.Product{}); err != nil {
		log.Panic(err.Error())
	}

	startGRPCServer(db)
}

func startGRPCServer(db *gorm.DB) {
	s := grpc.NewServer()

	productRepo := repo.NewProductRepo(db)
	productUseCase := usecase.NewProductUseCase(productRepo)

	pb.RegisterProductServiceServer(s, &server.Server{
		UnimplementedProductServiceServer: pb.UnimplementedProductServiceServer{},
		ProductUseCase:                    productUseCase,
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
