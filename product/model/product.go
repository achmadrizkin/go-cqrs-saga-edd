package model

import "time"

type Product struct {
	Id        string    `gorm:"type:varchar(255);primary_key" json:"Id"` // uuid
	Image_url string    `gorm:"type:varchar(255)" json:"Image_url"`
	Name      string    `gorm:"type:varchar(50)" json:"Name"`
	Price     int64     `gorm:"type:int" json:"Price"`
	Stock     int64     `gorm:"type:int" json:"Stock"`
	CreatedAt time.Time `gorm:"type:datetime" json:"CreatedAt"`
}
