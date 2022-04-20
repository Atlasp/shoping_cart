package model

type Cart struct {
	CartID     int64              `json:"cart_id" gorm`
	CustomerId int64              `json:"customer_id"`
	ItemId     []int64            `json:"item_id"`
	Items      map[int64]CartItem `json:"basket_items"`
	CartT      CartTotal          `json:"cart_t"`
}

type CartItem struct {
	UnitPrice int64 `json:"unit_price"`
	Quantity  int64 `json:"quantity"`
}

func (c *Cart) AddItem(i Item) {
	if _, ok := c.Items[i.Id]; ok {
		item := c.Items[i.Id]
		item.Quantity += 1
		c.Items[i.Id] = item
	} else {
		c.Items[i.Id] = CartItem{
			UnitPrice: i.Price,
			Quantity:  1,
		}
	}
}

func (c *Cart) GetCartTotal(cr CartRules) {
	if cr != (CartRules{}) {
		for _, v := range c.Items {
			c.CartT.TotalPrice += v.Quantity * v.UnitPrice
			c.CartT.Discount += (v.Quantity / cr.FreeItemThreshold) * v.UnitPrice
		}

		if c.CartT.TotalPrice > cr.DiscountThreshold {
			c.CartT.Discount += cr.Discount
		}
	} else {
		for _, v := range c.Items {
			c.CartT.TotalPrice += v.Quantity * v.UnitPrice
		}
	}

	c.CartT.FinalPrice = c.CartT.TotalPrice - c.CartT.Discount
}
