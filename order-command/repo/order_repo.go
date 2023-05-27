package repo

import (
	"errors"
	"go-cqrs-saga-edd/order-command/domain"
	"go-cqrs-saga-edd/order-command/model"

	"gorm.io/gorm"
)

type orderRepo struct {
	Db *gorm.DB
}

// CreateProductRepo implements domain.OrderRepo
func (o *orderRepo) CreateOrderRepo(order model.Order) error {
	if err := o.Db.Create(order).Error; err != nil {
		return errors.New("errCreatedOrderRepo: " + err.Error())
	}

	return nil
}

func NewOrderRepo(db *gorm.DB) domain.OrderRepo {
	return &orderRepo{
		Db: db,
	}
}
