package server

import (
	"context"
	"go-cqrs-saga-edd/order-query/domain"
	pb "go-cqrs-saga-edd/order-query/proto"
	"go-cqrs-saga-edd/order-query/utils"
)

type OrderCommandServer struct {
	pb.UnimplementedOrderCommandServiceServer
	OrderCommandUseCase domain.OrderCommandUseCase
}

func (o *OrderCommandServer) PostOrderProduct(ctx context.Context, req *pb.PostOrderProductRequest) (*pb.PostOrderProductResponse, error) {
	orderProduct, err := utils.ConverterProtoOrderProductToModel(req)
	if err != nil {
		return &pb.PostOrderProductResponse{
			StatusCode: 500,
			Message:    err.Error(),
		}, nil
	}

	// insert into mongodb
	if errCreateOrderProduct := o.OrderCommandUseCase.CreateOrderProduct(ctx, orderProduct); errCreateOrderProduct != nil {
		return &pb.PostOrderProductResponse{
			StatusCode: 500,
			Message:    errCreateOrderProduct.Error(),
		}, nil
	}

	return &pb.PostOrderProductResponse{
		StatusCode: 200,
		Message:    "Insert Data Success",
	}, nil
}
