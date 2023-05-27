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

// DeleteOrderRepo implements domain.OrderRepo
func (o *orderRepo) DeleteOrderRepo(order model.Order) (*gorm.DB, error) {
	tx := o.Db.Begin()

	if err := tx.Where("id = ?", order.Id).Delete(&model.Order{}).Error; err != nil {
		return nil, errors.New("errDeleteOrderRepo: " + err.Error())
	}

	return tx, nil
}

// CreateOrderRepo implements domain.OrderRepo
func (o *orderRepo) CreateOrderRepo(order model.Order) (*gorm.DB, error) {
	tx := o.Db.Begin()

	if err := tx.Create(&order).Error; err != nil {
		return nil, errors.New("errCreatedOrderRepo: " + err.Error())
	}

	return tx, nil
}

func NewOrderRepo(db *gorm.DB) domain.OrderRepo {
	return &orderRepo{
		Db: db,
	}
}
