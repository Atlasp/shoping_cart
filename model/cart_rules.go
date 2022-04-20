package model

import (
	"os"
	"strconv"
)

type CartRules struct {
	MaxBasketSize     int64
	DiscountThreshold int64
	Discount          int64
	FreeItemThreshold int64
}

func NewCartRules() CartRules {
	mbs := os.Getenv("MAX_BASKET_SIZE")
	dt := os.Getenv("DISCOUNT_THRESHOLD")
	d := os.Getenv("DISCOUNT")
	fit := os.Getenv("FREE_ITEM_THRESHOLD")

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
