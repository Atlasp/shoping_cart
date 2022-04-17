package model

type Discount struct {
	MaxBasketSize     int64 `json:"max_basket_size"`
	DiscountThreshold int64 `json:"discount_threshold"`
	FreeItemThreshold int64 `json:"free_item_threshold"`
}
