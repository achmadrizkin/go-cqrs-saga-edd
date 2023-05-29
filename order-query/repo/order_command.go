package repo

import (
	"errors"
	"go-cqrs-saga-edd/order-query/domain"
	"go-cqrs-saga-edd/order-query/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type orderCommandRepo struct {
	table *mongo.Collection
}

// CreateOrderProduct implements domain.OrderQueryRepo
func (o *orderCommandRepo) CreateOrderProduct(sc mongo.SessionContext, orderProduct model.OrderProduct) error {
	_, errInsert := o.table.InsertOne(sc, orderProduct)
	if errInsert != nil {
		return errors.New("errInsertOne: " + errInsert.Error())
	}

	return nil
}

func NewOrderCommandRepo(
	table *mongo.Collection,
) domain.OrderCommandRepo {
	return &orderCommandRepo{
		table,
	}
}
