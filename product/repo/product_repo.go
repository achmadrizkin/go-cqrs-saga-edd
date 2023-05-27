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

// GetAllProductRepo implements domain.ProductRepo
func (p *productRepo) GetAllProductRepo(productAll []model.Product) ([]model.Product, error) {
	err := p.Db.Model(&model.Product{}).Select("id", "image_url", "name", "price", "stock", "created_at").Find(&productAll).Error
	if err != nil {
		return productAll, errors.New("errGetAllProductRepo: " + err.Error())
	}

	return productAll, nil
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
