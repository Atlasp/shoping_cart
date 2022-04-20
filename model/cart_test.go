package model

import (
	"fmt"
	"testing"
)

func TestCart_AddItem(t *testing.T) {
	assertCorrect := func(t testing.TB, got, want interface{}) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("add one added", func(t *testing.T) {
		cart := Cart{
			CartID:     1,
			CustomerId: 1,
			Items:      map[int64]CartItem{},
		}
		item := Item{
			Id:       1,
			Name:     "Test",
			Category: "Test",
			Stock:    1,
			Price:    12,
		}
		cart.AddItem(item)
		got := cart.Items[item.Id]
		fmt.Println(got)
		want := CartItem{
			UnitPrice: item.Price,
			Quantity:  1,
		}
		fmt.Println(want)
		assertCorrect(t, got, want)
	})
	t.Run("add item multiple times", func(t *testing.T) {
		cart := Cart{
			CartID:     1,
			CustomerId: 1,
			Items:      map[int64]CartItem{},
		}
		item := Item{
			Id:       1,
			Name:     "Test",
			Category: "Test",
			Stock:    1,
			Price:    12,
		}
		cart.AddItem(item)
		cart.AddItem(item)
		got := cart.Items[item.Id]
		fmt.Println(got)
		want := CartItem{
			UnitPrice: item.Price,
			Quantity:  2,
		}
		fmt.Println(want)
		assertCorrect(t, got, want)
	})
	t.Run("add item already in cart", func(t *testing.T) {
		cart := Cart{
			CartID:     1,
			CustomerId: 1,
			Items: map[int64]CartItem{1: {
				UnitPrice: 12,
				Quantity:  1,
			}},
		}
		item := Item{
			Id:       1,
			Name:     "Test",
			Category: "Test",
			Stock:    1,
			Price:    12,
		}
		cart.AddItem(item)
		got := cart.Items[item.Id]
		fmt.Println(got)
		want := CartItem{
			UnitPrice: item.Price,
			Quantity:  2,
		}
		fmt.Println(want)
		assertCorrect(t, got, want)
	})
}

func TestCart_GetCartTotal(t *testing.T) {
	assertCorrect := func(t testing.TB, got, want interface{}) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("get cart total no discount", func(t *testing.T) {
		cr := CartRules{
			MaxBasketSize:     100,
			DiscountThreshold: 20,
			Discount:          1,
			FreeItemThreshold: 5,
		}
		cart := Cart{
			CartID:     1,
			CustomerId: 1,
			Items: map[int64]CartItem{
				1: {
					UnitPrice: 1,
					Quantity:  1,
				},
				2: {
					UnitPrice: 2,
					Quantity:  2,
				},
				3: {
					UnitPrice: 3,
					Quantity:  3,
				},
			},
		}
		cart.GetCartTotal(cr)
		got := cart.CartT
		want := CartTotal{
			TotalPrice: 14,
			Discount:   0,
			FinalPrice: 14,
		}
		assertCorrect(t, got, want)
	})
	t.Run("get cart total with discount threshold", func(t *testing.T) {
		cr := CartRules{
			MaxBasketSize:     100,
			DiscountThreshold: 20,
			Discount:          1,
			FreeItemThreshold: 5,
		}
		cart := Cart{
			CartID:     1,
			CustomerId: 1,
			Items: map[int64]CartItem{
				1: {
					UnitPrice: 10,
					Quantity:  1,
				},
				2: {
					UnitPrice: 2,
					Quantity:  2,
				},
				3: {
					UnitPrice: 3,
					Quantity:  3,
				},
			},
		}
		cart.GetCartTotal(cr)
		got := cart.CartT
		want := CartTotal{
			TotalPrice: 23,
			Discount:   1,
			FinalPrice: 22,
		}
		assertCorrect(t, got, want)
	})
	t.Run("get cart total with item threshold", func(t *testing.T) {
		cr := CartRules{
			MaxBasketSize:     100,
			DiscountThreshold: 20,
			Discount:          1,
			FreeItemThreshold: 5,
		}
		cart := Cart{
			CartID:     1,
			CustomerId: 1,
			Items: map[int64]CartItem{
				1: {
					UnitPrice: 1,
					Quantity:  1,
				},
				2: {
					UnitPrice: 2,
					Quantity:  5,
				},
				3: {
					UnitPrice: 3,
					Quantity:  1,
				},
			},
		}
		cart.GetCartTotal(cr)
		got := cart.CartT
		want := CartTotal{
			TotalPrice: 14,
			Discount:   2,
			FinalPrice: 12,
		}
		assertCorrect(t, got, want)
	})
}
