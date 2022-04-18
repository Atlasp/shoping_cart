package model

import "github.com/lib/pq"

type Cart struct {
	CartID     int64              `json:"cart_id"`
	CustomerId int64              `json:"customer_id"`
	CartItems  map[int64]CartItem `json:"basket_items"`
}

type CartTable struct {
	CartID     int64         `json:"cart_id" gorm:"primaryKey"`
	CustomerId int64         `json:"customer_id" gorm:"unique"`
	Items      pq.Int64Array `json:"items" gorm:"type:integer[]"`
}

type CartItem struct {
	UnitPrice int64 `json:"unit_price"`
	Quantity  int64 `json:"quantity"`
}

type CartTotal struct {
	TotalPrice int64 `json:"total_price"`
	Discount   int64 `json:"discount"`
	FinalPrice int64 `json:"final_price"`
}

// TODO: change this so it makes a new ID or remove it
func NewCart(customerId int) Cart {
	cartItems := make(map[int64]CartItem)
	return Cart{
		CartID:     0,
		CustomerId: int64(customerId),
		CartItems:  cartItems,
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

func (c *Cart) AddItemToCart(i Item) {
	if _, ok := c.CartItems[i.Id]; ok {
		item := c.CartItems[i.Id]
		item.Quantity += 1
		c.CartItems[i.Id] = item
	} else {
		c.CartItems[i.Id] = CartItem{
			UnitPrice: i.Price,
			Quantity:  1,
		}
	}
}

func (c Cart) GetCartTotal(cr CartRules) CartTotal {
	var total int64
	var discount int64

	if cr != (CartRules{}) {
		for _, v := range c.CartItems {
			total += v.Quantity * v.UnitPrice
			discount += (v.Quantity / cr.FreeItemThreshold) * v.UnitPrice
		}

		if total > cr.DiscountThreshold {
			discount += 1.00
		}
	} else {
		for _, v := range c.CartItems {
			total += v.Quantity * v.UnitPrice
		}
	}

	return CartTotal{
		TotalPrice: total,
		Discount:   discount,
		FinalPrice: total - discount,
	}
}
