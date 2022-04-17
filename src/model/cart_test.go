package model

import (
	"reflect"
	"testing"
)

func TestCart_GetDiscountedItems(t *testing.T) {
	type fields struct {
		CartID      uint
		CustomerId  uint
		BasketItems []Item
	}
	tests := []struct {
		name   string
		fields fields
		want   CartTotal
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cart{
				CartID:     tt.fields.CartID,
				CustomerId: tt.fields.CustomerId,
				CartItems:  tt.fields.BasketItems,
			}
			if got := c.GetCartTotal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCartTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCart_GetItemCounts(t *testing.T) {
	type fields struct {
		CartID      uint
		CustomerId  uint
		BasketItems []Item
	}
	tests := []struct {
		name   string
		fields fields
		want   map[int]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cart{
				CartID:     tt.fields.CartID,
				CustomerId: tt.fields.CustomerId,
				CartItems:  tt.fields.BasketItems,
			}
			if got := c.GetItemCounts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItemCounts() = %v, want %v", got, tt.want)
			}
		})
	}
}
