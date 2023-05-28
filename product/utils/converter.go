package utils

import "go-cqrs-saga-edd/product/model"

func ConverterOrderAndProductToOrderProduct(order model.Order, product model.Product) model.OrderProduct {
	data := model.OrderProduct{
		Id:          order.Id,
		ProductId:   order.ProductId,
		Quantity:    order.Quantity,
		ShipMethod:  order.ShipMethod,
		Address:     order.Address,
		TotalPrice:  int64(order.Quantity) * product.Price,
		Date:        order.Date,
		ProductData: product,
	}

	return data
}
