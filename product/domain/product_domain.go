package domain

import (
	"go-cqrs-saga-edd/product/model"
)

type ProductRepo interface {
	CreateProductRepo(model.Product) error

	GetAllProductRepo([]model.Product) ([]model.Product, error)
}

type ProductUseCase interface {
	CreateProductUseCase(model.Product) error

	GetAllProductUseCase([]model.Product) ([]model.Product, error)
}
