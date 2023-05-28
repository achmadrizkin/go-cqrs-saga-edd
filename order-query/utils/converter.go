package utils

import (
	"errors"
	"go-cqrs-saga-edd/order-query/model"
	pb "go-cqrs-saga-edd/order-query/proto"

	"github.com/golang/protobuf/ptypes"
)

func ConverterProtoOrderProductToModel(req *pb.PostOrderProductRequest) (model.OrderProduct, error) {
	createdAt, err := ptypes.Timestamp(req.GetProduct().GetCreatedAt())
	if err != nil {
		return model.OrderProduct{}, errors.New("errConvertCreatedAt: " + err.Error())
	}
	var productData model.Product = model.Product{
		Id:        req.GetProduct().GetId(),
		Image_url: req.GetProduct().GetImageUrl(),
		Name:      req.GetProduct().GetName(),
		Price:     req.GetProduct().GetPrice(),
		Stock:     req.GetProduct().GetStock(),
		CreatedAt: createdAt,
	}

	var orderProductData model.OrderProduct = model.OrderProduct{
		Id:          req.GetProductId(),
		ProductId:   req.GetProductId(),
		Quantity:    req.GetQuantity(),
		ShipMethod:  req.GetShipMethod(),
		Address:     req.GetAddress(),
		TotalPrice:  int64(req.Quantity) * req.GetProduct().Price,
		ProductData: productData,
	}

	return orderProductData, nil
}
