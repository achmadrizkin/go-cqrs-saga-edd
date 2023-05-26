package usecase

import (
	"go-cqrs-saga-edd/product/model"
	"go-cqrs-saga-edd/product/repo"

	"gorm.io/gorm"
)

// CreateProduct implements domain.ProductUseCase
func CreateProductUseCase(db *gorm.DB, product model.Product) error {
	return repo.CreateProductRepo(db, product)
}
