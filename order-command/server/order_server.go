package server

import (
	"context"
	"go-cqrs-saga-edd/order-command/domain"
	"go-cqrs-saga-edd/order-command/model"
	pb "go-cqrs-saga-edd/order-command/proto"
	"time"

	"github.com/google/uuid"
)

type OrderServer struct {
	pb.UnimplementedOrderServiceServer
	OrderUseCase          domain.OrderUseCase
	OrderPublisherUseCase domain.OrderPublisherUseCase
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
	err, tx := o.OrderUseCase.CreateOrderUseCase(orderData)
	if err != nil {
		tx.Rollback() // Rollback the transaction in case of an error
		return &pb.PostOrderResponse{
			StatusCode: 500,
			Message:    err.Error(),
		}, nil
	}

	if errPublisher := o.OrderPublisherUseCase.CreateOrderUseCasePublisherToProduct(orderData, "eOrderToProduct"); errPublisher != nil {
		tx.Rollback() // Rollback the transaction in case of an error
		return &pb.PostOrderResponse{
			StatusCode: 500,
			Message:    errPublisher.Error(),
		}, nil
	}

	// Commit the transaction if everything is successful
	tx.Commit()

	return &pb.PostOrderResponse{
		StatusCode: 200,
		Message:    "Insert Data Order Success",
	}, nil
}
