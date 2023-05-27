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

// UpdateStockProductRepo implements domain.ProductRepo
func (p *productRepo) UpdateStockProductRepo(productUUID string, stock int64, isSuccess int) (*gorm.DB, error) {
	tx := p.Db.Begin()

	var stockProduct int64
	var stockTotal int64
	if err := p.Db.Model(&model.Product{}).Select("stock").Where("id = ?", productUUID).Find(&stockProduct).Error; err != nil {
		return nil, errors.New("errGetProduct: " + err.Error())
	}

	if isSuccess == 1 {
		stockTotal = stockProduct - stock
	} else if isSuccess == 0 {
		stockTotal = stockProduct + stock
	} else {
		return nil, errors.New("errIsSuccessStock: Must be 0/1")
	}

	if stockTotal < 0 {
		return nil, errors.New("stock cannot be less than 0")
	}

	if err := tx.Model(&model.Product{}).Where("id = ?", productUUID).Updates(model.Product{Stock: stockTotal}).Error; err != nil {
		return nil, errors.New("errUpdatedStockProduct: " + err.Error())
	}

	return tx, nil
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
