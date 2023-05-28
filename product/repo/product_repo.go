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

// GetProductByIdRepo implements domain.ProductRepo
func (p *productRepo) GetProductByIdRepo(productUUID string) (model.Product, error) {
	var product model.Product
	err := p.Db.Model(&model.Product{}).Select("id", "image_url", "name", "price", "stock", "created_at").Where("id = ?", productUUID).Find(&product).Error
	if err != nil {
		return product, errors.New("errGetAllProductRepo: " + err.Error())
	}

	return product, nil
}

// UpdateStockProductRepo implements domain.ProductRepo
func (p *productRepo) UpdateStockProductRepo(productUUID string, stock int64, isSuccess int) (*gorm.DB, error) {
	tx := p.Db.Begin()

	var stockTotal int64
	var product model.Product
	if err := p.Db.Model(&model.Product{}).Select("name,stock").Where("id = ?", productUUID).Find(&product).Error; err != nil {
		return tx, errors.New("errGetProduct: " + err.Error())
	}

	if product.Name == "" {
		return tx, errors.New("product not found")
	}

	if isSuccess == 1 {
		stockTotal = product.Stock - stock
	} else if isSuccess == 0 {
		stockTotal = product.Stock + stock
	} else {
		return tx, errors.New("errIsSuccessStock: Must be 0/1")
	}

	if stockTotal < 0 {
		return tx, errors.New("stock cannot be less than 0")
	}

	if err := tx.Model(&model.Product{}).Where("id = ?", productUUID).Updates(model.Product{Stock: stockTotal}).Error; err != nil {
		return tx, errors.New("errUpdatedStockProduct: " + err.Error())
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
