package model

import (
	"os"
	"strconv"
)

type CartRules struct {
	MaxBasketSize     int64 `env:"max_basket_size"`
	DiscountThreshold int64 `env:"discount_threshold"`
	Discount          int64 `env:"discount"`
	FreeItemThreshold int64 `env:"free_item_threshold"`
}

func NewCartRules() CartRules {
	mbs := os.Getenv("max_basket_size")
	dt := os.Getenv("discount_threshold")
	d := os.Getenv("discount")
	fit := os.Getenv("free_item_threshold")

	mBS, _ := strconv.ParseInt(mbs, 10, 64)
	dT, _ := strconv.ParseInt(dt, 10, 64)
	D, _ := strconv.ParseInt(d, 10, 64)
	fIT, _ := strconv.ParseInt(fit, 10, 64)

	return CartRules{
		MaxBasketSize:     mBS,
		DiscountThreshold: dT,
		Discount:          D,
		FreeItemThreshold: fIT,
	}
}
