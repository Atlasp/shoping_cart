package repository

import (
	"revel_systems_shopping/model"
)

func (r *Repository) PopulateDB() {
	items := []model.Item{
		{
			Id:       1,
			Name:     "Coca Cola",
			Category: "Soft Drinks",
			Stock:    56,
			Price:    3,
		},
		{
			Id:       2,
			Name:     "Tik Tak",
			Category: "Candies",
			Stock:    73,
			Price:    1,
		},
		{
			Id:       3,
			Name:     "Learn GO",
			Category: "Books",
			Stock:    3,
			Price:    35,
		},
		{
			Id:       4,
			Name:     "Marlboro Light",
			Category: "Tobacco",
			Stock:    21,
			Price:    8,
		},
		{
			Id:       5,
			Name:     "Super Coffee",
			Category: "Coffee",
			Stock:    32,
			Price:    9,
		},
		{
			Id:       6,
			Name:     "Super Nice Watch",
			Category: "Luxury goods",
			Stock:    1,
			Price:    1200,
		},
		{
			Id:       7,
			Name:     "Bread",
			Category: "Bread",
			Stock:    120,
			Price:    3,
		},
		{
			Id:       8,
			Name:     "Greek Yogurt",
			Category: "Dairy products",
			Stock:    19,
			Price:    3,
		},
		{
			Id:       9,
			Name:     "Jack Daniels",
			Category: "Alcohol",
			Stock:    23,
			Price:    18,
		},
		{
			Id:       10,
			Name:     "Sugar",
			Category: "Sugar",
			Stock:    31,
			Price:    4,
		},
	}

	userCarts := []model.CartTable{
		{
			CartID:     1,
			CustomerId: 1,
			Items:      nil,
		},
		{
			CartID:     2,
			CustomerId: 2,
			Items:      []int64{1, 2, 3, 4, 5},
		},
		{
			CartID:     3,
			CustomerId: 3,
			Items:      []int64{1, 1, 1, 1, 1, 7, 8, 3},
		},
		{
			CartID:     4,
			CustomerId: 4,
			Items:      []int64{6, 4, 5, 5, 8},
		},
		{
			CartID:     5,
			CustomerId: 5,
			Items:      []int64{10, 9, 3, 4, 4},
		},
	}

	r.DB.Create(&items)
	r.DB.Create(&userCarts)
}
