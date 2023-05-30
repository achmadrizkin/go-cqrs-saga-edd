package server

import (
	"context"
	"go-cqrs-saga-edd/order-query/domain"
	"go-cqrs-saga-edd/order-query/model"
	pb "go-cqrs-saga-edd/order-query/proto"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type OrderQueryServer struct {
	pb.UnimplementedOrderQueryServiceServer
	OrderQueryUseCase domain.OrderQueryUseCase
}

func (o *OrderQueryServer) GetOrderProductByOrderId(ctx context.Context, req *pb.GetOrderProductByOrderIdRequest) (*pb.GetOrderProductByOrderIdResponse, error) {
	orderProduct, err := o.OrderQueryUseCase.GetOrderById(ctx, req.GetId())
	if err != nil {
		return &pb.GetOrderProductByOrderIdResponse{
			StatusCode: 200,
			Message:    "errGetProductByOrderById: " + err.Error(),
		}, nil
	}

	return &pb.GetOrderProductByOrderIdResponse{
		StatusCode: 200,
		Message:    "Get Order By Id Success",
		Data:       convertToGetOrderProductResponse(orderProduct),
	}, nil
}

func (o *OrderQueryServer) GetOrderProductAll(ctx context.Context, req *pb.GetOrderProductRequest) (*pb.GetAllOrderProductResponse, error) {
	orderProducts, err := o.OrderQueryUseCase.GetOrderProductAll(ctx)
	if err != nil {
		return &pb.GetAllOrderProductResponse{
			StatusCode: 200,
			Message:    "errGetProductAll: " + err.Error(),
		}, nil
	}

	var orderProductResponses []*pb.GetOrderProductResponse
	for _, op := range orderProducts {
		orderProductResponses = append(orderProductResponses, convertToGetOrderProductResponse(op))
	}

	return &pb.GetAllOrderProductResponse{
		StatusCode: 200,
		Message:    "Get All Product Response Success",
		Data:       orderProductResponses,
	}, nil
}

func convertToGetOrderProductResponse(op model.OrderProduct) *pb.GetOrderProductResponse {
	// Convert the fields from the model.OrderProduct to the corresponding fields in GetOrderProductResponse
	return &pb.GetOrderProductResponse{
		Id:         op.Id,
		ProductId:  op.ProductId,
		Quantity:   op.Quantity,
		ShipMethod: op.ShipMethod,
		Address:    op.Address,
		Date:       convertToTimestamp(op.Date),
		Product:    convertToProduct(op.ProductData),
	}
}

func convertToProduct(p model.Product) *pb.ProductQuery {
	// Convert the fields from the model.Product to the corresponding fields in Product
	return &pb.ProductQuery{
		Id:        p.Id,
		ImageUrl:  p.Image_url,
		Price:     p.Price,
		Stock:     p.Stock,
		CreatedAt: convertToTimestamp(p.CreatedAt),
	}
}

func convertToTimestamp(t time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(t)
	return ts
}
