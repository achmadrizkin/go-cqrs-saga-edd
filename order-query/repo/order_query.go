package repo

import (
	"context"
	"errors"
	"go-cqrs-saga-edd/order-query/domain"
	"go-cqrs-saga-edd/order-query/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type orderQueryRepo struct {
	table *mongo.Collection
}

// GetOrderById implements domain.OrderQueryRepo
func (o *orderQueryRepo) GetOrderById(ctx context.Context, id string) (model.OrderProduct, error) {
	var orderProduct model.OrderProduct

	// Filter order by ID
	filter := bson.M{"Id": id}

	// Find order matching the filter
	cursor, err := o.table.Find(ctx, filter)
	if err != nil {
		return orderProduct, errors.New("errorsFilter: " + err.Error())
	}

	// Check if any order was found
	if !cursor.Next(ctx) {
		return orderProduct, errors.New("orderProduct not found")
	}

	// Decode the matching order
	if err := cursor.Decode(&orderProduct); err != nil {
		return orderProduct, errors.New("error decoding orderProduct: " + err.Error())
	}

	return orderProduct, nil
}

func NewOrderQueryRepo(
	table *mongo.Collection,
) domain.OrderQueryRepo {
	return &orderQueryRepo{
		table,
	}
}
