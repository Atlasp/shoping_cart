package model

type Item struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"unique"`
	Category string `json:"category"`
	Stock    int64  `json:"stock"`
	Price    int64  `json:"price"`
}
