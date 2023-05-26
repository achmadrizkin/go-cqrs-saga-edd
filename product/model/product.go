package model

type Product struct {
	Id        string `gorm:"type:varchar(255);primary_key"` // uuid
	Image_url string `gorm:"type:varchar(255)"`
	Name      string `gorm:"type:varchar(50)"`
	Price     int64  `gorm:"type:int"`
	Stock     int64  `gorm:"type:int"`
}
