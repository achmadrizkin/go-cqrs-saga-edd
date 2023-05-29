package server

import (
	"context"
	"errors"
	"go-cqrs-saga-edd/order-query/domain"
	pb "go-cqrs-saga-edd/order-query/proto"
	"go-cqrs-saga-edd/order-query/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type OrderCommandServer struct {
	pb.UnimplementedOrderCommandServiceServer
	OrderCommandUseCase domain.OrderCommandUseCase
	Client              *mongo.Client
}

func (o *OrderCommandServer) PostOrderProduct(ctx context.Context, req *pb.PostOrderProductRequest) (*pb.PostOrderProductResponse, error) {
	orderProduct, err := utils.ConverterProtoOrderProductToModel(req)
	if err != nil {
		return &pb.PostOrderProductResponse{
			StatusCode: 500,
			Message:    err.Error(),
		}, nil
	}

	// Set the read and write concerns for the transaction
	txOptions := options.Transaction().
		SetReadConcern(readconcern.Majority()).
		SetWriteConcern(writeconcern.New())

	// Start a session
	session, err := o.Client.StartSession()
	if err != nil {
		return &pb.PostOrderProductResponse{
			StatusCode: 500,
			Message:    err.Error(),
		}, nil
	}
	defer session.EndSession(ctx)

	// Start a transaction
	err = session.StartTransaction(txOptions)
	if err != nil {
		return &pb.PostOrderProductResponse{
			StatusCode: 500,
			Message:    err.Error(),
		}, nil
	}

	// Start Session
	if errSession := mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		errCreatedProduct := o.OrderCommandUseCase.CreateOrderProduct(sc, orderProduct)
		if err != nil {
			return errors.New("errCreatedProduct: " + errCreatedProduct.Error())
		}
		if errCommitTransaction := session.CommitTransaction(sc); err != nil {
			return errors.New("errCommitTransaction: " + errCommitTransaction.Error())
		}
		return nil
	}); errSession != nil {
		return &pb.PostOrderProductResponse{
			StatusCode: 500,
			Message:    errSession.Error(),
		}, nil
	}

	return &pb.PostOrderProductResponse{
		StatusCode: 200,
		Message:    "Insert Data Success",
	}, nil
}
