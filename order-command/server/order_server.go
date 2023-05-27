package server

import (
	"context"
	"encoding/json"
	"go-cqrs-saga-edd/order-command/domain"
	"go-cqrs-saga-edd/order-command/model"
	pb "go-cqrs-saga-edd/order-command/proto"
	"go-cqrs-saga-edd/order-command/utils"
	"log"
	"time"

	"github.com/google/uuid"
)

type OrderServer struct {
	pb.UnimplementedOrderServiceServer
	OrderUseCase domain.OrderUseCase
}

// PostOrder implements __.OrderServiceServer
func (o *OrderServer) PostOrder(ctx context.Context, req *pb.PostOrderRequest) (*pb.PostOrderResponse, error) {
	var orderData model.Order = model.Order{
		Id:         uuid.New().String(),
		ProductId:  req.GetProductId(),
		Quantity:   req.GetQuantity(),
		ShipMethod: req.GetShipMethod(),
		Address:    req.GetAddress(),
		Date:       time.Now(),
	}

	// insert into db
	if err := o.OrderUseCase.CreateOrderUseCase(orderData); err != nil {
		return &pb.PostOrderResponse{
			StatusCode: 500,
			Message:    err.Error(),
		}, nil
	}

	// The encryption key (must be 16, 24, or 32 bytes)
	key := []byte("1423456789012345")

	orderJSON, errMarshal := json.Marshal(orderData)
	if errMarshal != nil {
		return &pb.PostOrderResponse{
			StatusCode: 500,
			Message:    "errMarshal: " + errMarshal.Error(),
		}, nil
	}

	orderDataString := string(orderJSON)
	log.Println("orderDataString:" + orderDataString)

	// Encrypt
	encrypted, err := utils.EncryptAES(key, orderDataString)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Encrypted:", encrypted)

	// Decrypt
	decrypted, err := utils.DecryptAES(key, encrypted)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Decrypted:", decrypted)

	return &pb.PostOrderResponse{
		StatusCode: 200,
		Message:    "Insert Data Order Success",
	}, nil
}
