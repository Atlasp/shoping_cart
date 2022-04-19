package model

import (
	"encoding/json"
	"io/ioutil"
)

type CartRules struct {
	MaxBasketSize     int64 `json:"max_basket_size"`
	DiscountThreshold int64 `json:"discount_threshold"`
	Discount          int64 `json:"discount"`
	FreeItemThreshold int64 `json:"free_item_threshold"`
}

func ReadCartRules(path string) CartRules {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	cr := CartRules{}
	err = json.Unmarshal(file, &cr)
	if err != nil {
		panic(err)
	}
	return cr
}
