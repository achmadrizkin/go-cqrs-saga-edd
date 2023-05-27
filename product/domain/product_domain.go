package domain

import (
	"go-cqrs-saga-edd/product/model"

	"gorm.io/gorm"
)

type ProductRepo interface {
	CreateProductRepo(model.Product) error

	GetAllProductRepo([]model.Product) ([]model.Product, error)
	UpdateStockProductRepo(productUUID string, stock int64, isSuccess int) (*gorm.DB, error)
}

type ProductUseCase interface {
	CreateProductUseCase(model.Product) error

	GetAllProductUseCase([]model.Product) ([]model.Product, error)
}
