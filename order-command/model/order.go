package model

import "time"

type Order struct {
	Id         string    `gorm:"type:varchar(255);primary_key"` // uuid
	ProductId  string    `gorm:"type:varchar(255)"`             // uuid product id
	Quantity   int32     `gorm:"type:int"`
	ShipMethod string    `gorm:"type:varchar(10)"`
	Address    string    `gorm:"type:varchar(255)"`
	Date       time.Time `gorm:"type:datetime"`
}
