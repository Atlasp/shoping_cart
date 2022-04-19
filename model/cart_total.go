package model

type CartTotal struct {
	TotalPrice int64 `json:"total_price"`
	Discount   int64 `json:"discount"`
	FinalPrice int64 `json:"final_price"`
}
