package repo

import (
	"context"
	"errors"
	"fmt"
	"go-cqrs-saga-edd/order-query/domain"
	"go-cqrs-saga-edd/order-query/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type orderQueryRepo struct {
	table *mongo.Collection
}

// GetOrderProductAll implements domain.OrderQueryRepo
func (repo *orderQueryRepo) GetOrderProductAll(ctx context.Context) ([]model.OrderProduct, error) {
	// Retrieve the cursor
	cursor, err := repo.table.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error retrieving order products: %w", err)
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and collect order products
	var orderProducts []model.OrderProduct
	for cursor.Next(ctx) {
		var orderProduct model.OrderProduct
		if err := cursor.Decode(&orderProduct); err != nil {
			return nil, fmt.Errorf("error decoding order product: %w", err)
		}
		orderProducts = append(orderProducts, orderProduct)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error iterating cursor: %w", err)
	}

	return orderProducts, nil
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
