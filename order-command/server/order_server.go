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

	return &pb.PostOrderResponse{
		StatusCode: 200,
		Message:    "Insert Data Order Success",
	}, nil
}
