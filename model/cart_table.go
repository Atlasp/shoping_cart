package model

import "github.com/lib/pq"

type CartTable struct {
	CartID     int64         `json:"cart_id" gorm:"primaryKey"`
	CustomerId int64         `json:"customer_id" gorm:"unique"`
	Items      pq.Int64Array `json:"items" gorm:"type:integer[]"`
}

func NewCart(customerId int64) CartTable {
	return CartTable{
		CustomerId: customerId,
		Items:      nil,
	}
}

func (ct CartTable) ParseCartTable() Cart {
	cartItems := make(map[int64]CartItem)

	return Cart{
		CartID:     ct.CartID,
		CustomerId: ct.CustomerId,
		CartItems:  cartItems,
	}
}
