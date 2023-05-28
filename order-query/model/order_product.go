package model

import (
	"time"
)

type Product struct {
	Id        string    `bson:"Id" json:"Id"`
	Image_url string    `bson:"Image_url" json:"Image_url"`
	Name      string    `bson:"Name" json:"Name"`
	Price     int64     `bson:"Price" json:"Price"`
	Stock     int64     `bson:"Stock" json:"Stock"`
	CreatedAt time.Time `bson:"CreatedAt" json:"CreatedAt"`
}

type OrderProduct struct {
	Id          string    `bson:"Id" json:"Id"`
	ProductId   string    `bson:"ProductId" json:"ProductId"`
	Quantity    int32     `bson:"Quantity" json:"Quantity"`
	ShipMethod  string    `bson:"ShipMethod" json:"ShipMethod"`
	Address     string    `bson:"Address" json:"Address"`
	TotalPrice  int64     `bson:"TotalPrice" json:"TotalPrice"`
	Date        time.Time `bson:"Date" json:"Date"`
	ProductData Product   `bson:"ProductData" json:"ProductData"`
}
