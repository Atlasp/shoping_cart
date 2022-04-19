package model

type Cart struct {
	CartID     int64              `json:"cart_id" gorm`
	CustomerId int64              `json:"customer_id"`
	CartItems  map[int64]CartItem `json:"basket_items"`
}

type CartItem struct {
	UnitPrice int64 `json:"unit_price"`
	Quantity  int64 `json:"quantity"`
}

func (c *Cart) AddItem(i Item) {
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
			discount += cr.Discount
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
