package model

import "time"

type Order struct {
	Id         string    `json:"Id"`
	ProductId  string    `json:"ProductId"`
	Quantity   int32     `json:"Quantity"`
	ShipMethod string    `json:"ShipMethod"`
	Address    string    `json:"Address"`
	Date       time.Time `json:"Date"`
}
