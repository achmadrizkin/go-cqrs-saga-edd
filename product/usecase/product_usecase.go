package usecase

import (
	"go-cqrs-saga-edd/product/domain"
	"go-cqrs-saga-edd/product/model"
)

type productUseCase struct {
	productRepo domain.ProductRepo
}

// CreateProductUseCase implements domain.ProductUseCase
func (p *productUseCase) CreateProductUseCase(product model.Product) error {
	return p.productRepo.CreateProductRepo(product)
}

func NewProductUseCase(productRepo domain.ProductRepo) domain.ProductUseCase {
	return &productUseCase{productRepo: productRepo}
}
