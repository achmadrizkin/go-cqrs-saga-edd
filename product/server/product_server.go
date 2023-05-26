package server

import (
	"context"
	"go-cqrs-saga-edd/product/domain"
	"go-cqrs-saga-edd/product/model"
	pb "go-cqrs-saga-edd/product/proto"

	"github.com/google/uuid"
)

type Server struct {
	pb.UnimplementedProductServiceServer
	ProductUseCase domain.ProductUseCase
}

func (s *Server) GetProductAll(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	var productAll []model.Product

	productAllResponse, err := s.ProductUseCase.GetAllProductUseCase(productAll)
	if err != nil {
		return &pb.GetProductResponse{
			StatusCode: 500,
			Message:    err.Error(),
		}, nil
	}

	// Convert model.Product to []*pb.Product
	var products []*pb.Product
	for _, p := range productAllResponse {
		product := &pb.Product{
			ImageUrl: p.Image_url,
			Name:     p.Name,
			Price:    p.Price,
			Stock:    p.Stock,
		}
		products = append(products, product)
	}

	return &pb.GetProductResponse{
		StatusCode: 200,
		Message:    "Success Get All Product",
		Data:       products,
	}, nil
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
	err := s.ProductUseCase.CreateProductUseCase(productData)
	if err != nil {
		return &pb.PostProductResponse{
			StatusCode: 500,
			Message:    err.Error(),
		}, nil
	}

	return &pb.PostProductResponse{
		StatusCode: 200,
		Message:    "Insert Product Success",
	}, nil
}
