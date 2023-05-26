package repo

import (
	"errors"
	"go-cqrs-saga-edd/product/domain"
	"go-cqrs-saga-edd/product/model"

	"gorm.io/gorm"
)

type productRepo struct {
	Db *gorm.DB
}

// CreateProductRepo implements domain.ProductRepo
func (p *productRepo) CreateProductRepo(product model.Product) error {
	if err := p.Db.Create(product).Error; err != nil {
		return errors.New("errCreateProduct: " + err.Error())
	}

	return nil
}

func NewProductRepo(db *gorm.DB) domain.ProductRepo {
	return &productRepo{
		Db: db,
	}
}
