package server

import (
	"context"
	"go-cqrs-saga-edd/product/domain"
	"go-cqrs-saga-edd/product/model"
	pb "go-cqrs-saga-edd/product/proto"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		createdAt := timestamppb.New(p.CreatedAt)
		product := &pb.Product{
			Id:        p.Id,
			ImageUrl:  p.Image_url,
			Name:      p.Name,
			Price:     p.Price,
			Stock:     p.Stock,
			CreatedAt: createdAt,
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

	var productData model.Product = model.Product{
		Id:        uuid.New().String(),
		Image_url: req.GetImageUrl(),
		Name:      req.GetName(),
		Price:     req.GetPrice(),
		Stock:     req.GetStock(),
		CreatedAt: time.Now(),
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
