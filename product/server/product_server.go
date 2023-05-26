package server

import (
	"context"
	"go-cqrs-saga-edd/product/model"
	pb "go-cqrs-saga-edd/product/proto"
	"go-cqrs-saga-edd/product/usecase"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Server struct {
	pb.UnimplementedProductServiceServer
	Db *gorm.DB
}

// mustEmbedUnimplementedProductServiceServer implements __.ProductServiceServer
func (*Server) mustEmbedUnimplementedProductServiceServer() {
}

func (s *Server) GetProductAll(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	// Implementation of GetProductAll method
	return nil, nil
}

func (s *Server) PostProduct(ctx context.Context, req *pb.PostProductRequest) (*pb.PostProductResponse, error) {

	var productRequest *pb.Product = req.GetProduct()

	var productData model.Product = model.Product{
		Id:        uuid.New().String(),
		Image_url: productRequest.ImageUrl,
		Name:      productRequest.Name,
		Price:     productRequest.Price,
		Stock:     productRequest.Stock,
	}

	// insert into db
	err := usecase.CreateProductUseCase(s.Db, productData)
	if err != nil {
		return &pb.PostProductResponse{
			StatusCode: "500",
			Message:    err.Error(),
		}, nil
	}

	return &pb.PostProductResponse{
		StatusCode: "200",
		Message:    "Insert Product Success",
	}, nil
}
