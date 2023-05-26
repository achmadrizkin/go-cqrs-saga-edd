package repo

import (
	"errors"
	"go-cqrs-saga-edd/product/model"

	"gorm.io/gorm"
)

// CreateProduct implements domain.ProductRepo
func CreateProductRepo(Db *gorm.DB, product model.Product) error {
	if err := Db.Create(product).Error; err != nil {
		return errors.New("errCreateProduct: " + err.Error())
	}

	return nil
}
