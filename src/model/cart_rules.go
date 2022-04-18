package model

import (
	"encoding/json"
	"io/ioutil"
)

type CartRules struct {
	MaxBasketSize     int64 `json:"max_basket_size"`
	DiscountThreshold int64 `json:"discount_threshold"`
	FreeItemThreshold int64 `json:"free_item_threshold"`
}

func ReadCartRules(path string) CartRules {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return CartRules{}
	}
	cr := CartRules{}
	err = json.Unmarshal([]byte(file), &cr)
	if err != nil {
		return CartRules{}
	}
	return cr
}
