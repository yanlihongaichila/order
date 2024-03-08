package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(100)"`
	Description string  `gorm:"type:text"`
	Type        int64   `gorm:"type:int(5)"`
	Price       float64 `gorm:"type:float(8)"`
	Num         int64   `gorm:"type:int(6)"`
	Img         string  `gorm:"type:text"`
	Status      int64   `gorm:"type:tinyint(1)"`
}

func NewOrder() *Order {
	return new(Order)
}
