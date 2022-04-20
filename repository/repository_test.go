package repository

import (
	"os"
	"reflect"
	"testing"
	"time"

	"revel_systems_shopping/model"
)

func CreateTestingRepository() Repository {
	cr := model.CartRules{
		MaxBasketSize:     15,
		DiscountThreshold: 10,
		Discount:          1,
		FreeItemThreshold: 5,
	}
	dbConn := os.Getenv("TEST_DB_CONN")
	repository := NewRepository(dbConn, cr)

	return repository
}

func PopulateTestTable(r Repository) {
	testCarts := []model.CartTable{
		{
			CartID:     1,
			CustomerId: 1,
			Items:      []int64{1, 2, 3},
		},
		{
			CartID:     2,
			CustomerId: 2,
			Items:      []int64{1, 2, 2, 2, 2, 2, 3, 3, 3},
		},
		{
			CartID:     3,
			CustomerId: 3,
			Items:      nil,
		},
	}

	testItems := []model.Item{
		{
			Id:       1,
			Name:     "test_1",
			Category: "test_1",
			Stock:    10,
			Price:    1,
		},
		{
			Id:       2,
			Name:     "test_2",
			Category: "test_2",
			Stock:    10,
			Price:    2,
		},
		{
			Id:       3,
			Name:     "test_3",
			Category: "test_3",
			Stock:    0,
			Price:    3,
		},
	}

	r.DB.Create(&testCarts)
	r.DB.Create(&testItems)
}

func CleanupRepo(r Repository) {
	truncate := "TRUNCATE TABLE cart_tables;TRUNCATE TABLE items;TRUNCATE TABLE orders;"
	r.DB.Exec(truncate)
}

var repo = CreateTestingRepository()

func Test_GetItem(t *testing.T) {
	PopulateTestTable(repo)
	defer CleanupRepo(repo)
	assertCorrect := func(t testing.TB, got, want interface{}) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("test getting item from database", func(t *testing.T) {
		got, _ := repo.GetItem(1)
		want := model.Item{
			Id:       1,
			Name:     "test_1",
			Category: "test_1",
			Stock:    10,
			Price:    1,
		}
		assertCorrect(t, got, want)
	})
	t.Run("test getting nonexistent item", func(t *testing.T) {
		_, got := repo.GetItem(4)
		if got == nil {
			t.Errorf("did not throw error")
		}
	})
}

func TestRepository_DeleteItem(t *testing.T) {
	PopulateTestTable(repo)
	defer CleanupRepo(repo)
	t.Run("delete existing item", func(t *testing.T) {
		err := repo.DeleteItem(1)
		if err != nil {
			t.Fail()
		}
	})
}

func TestRepository_AddItem(t *testing.T) {
	PopulateTestTable(repo)
	defer CleanupRepo(repo)

	assertCorrect := func(t testing.TB, got, want interface{}) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("add item to database", func(t *testing.T) {
		want := model.Item{
			Id:       10,
			Name:     "item_1",
			Category: "category_1",
			Stock:    1,
			Price:    1,
		}
		err := repo.AddItem(want)
		if err != nil {
			t.Fail()
		}
		got, err := repo.GetItem(want.Id)
		if err != nil {
			t.Fail()
		}
		assertCorrect(t, got, want)
	})
	t.Run("fail add existing item", func(t *testing.T) {
		item := model.Item{
			Id:       10,
			Name:     "item_1",
			Category: "category_1",
			Stock:    1,
			Price:    1,
		}
		err := repo.AddItem(item)
		if err == nil {
			t.Errorf("did not throw error")
		}
	})
}

func TestRepository_PlaceOrder(t *testing.T) {
	PopulateTestTable(repo)
	defer CleanupRepo(repo)

	assertCorrect := func(t testing.TB, got, want interface{}) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("place order", func(t *testing.T) {
		var got model.Order
		var resultItems []int64
		err := repo.PlaceOrder(1)
		if err != nil {
			t.Fail()
		}
		want := model.Order{
			CustomerId:    1,
			OrderedItems:  []int64{1, 2, 3},
			OrderTotal:    6,
			OrderDiscount: 0,
			OrderFinal:    6,
			OrderDate:     time.Time{},
		}
		repo.DB.First(&got, "customer_id = ?", 1)

		for _, i := range got.OrderedItems {
			resultItems = append(resultItems, i)
		}

		result := reflect.DeepEqual(want.OrderedItems, got.OrderedItems)
		if result != true {
			t.Errorf("got %q want %q", got.OrderedItems, want.OrderedItems)
		}
		assertCorrect(t, want.OrderTotal, got.OrderTotal)
		assertCorrect(t, want.OrderDiscount, got.OrderDiscount)
		assertCorrect(t, want.OrderFinal, got.OrderFinal)
	})
	t.Run("fail when placing order over cart limit", func(t *testing.T) {
		err := repo.PlaceOrder(2)
		if err == nil {
			t.Errorf("did not throw error")
		}
	})
}

func TestRepository_AddItemToCart(t *testing.T) {
	PopulateTestTable(repo)
	defer CleanupRepo(repo)
	t.Run("add item to cart", func(t *testing.T) {
		var got model.CartTable
		err := repo.AddItemToCart(3, 1)
		if err != nil {
			t.Fail()
		}
		if err != nil {
			t.Fail()
		}
		repo.DB.First(&got, "customer_id = ?", 3)
		want := model.CartTable{
			Items: []int64{1},
		}

		result := reflect.DeepEqual(got.Items, want.Items)
		if result != true {
			t.Errorf("got %q want %q", got.Items, want.Items)
		}
	})
	t.Run("fail when adding out of stock item", func(t *testing.T) {
		err := repo.AddItemToCart(3, 3)
		if err == nil {
			t.Errorf("did not throw error")
		}
	})
}

func TestRepository_GetCart(t *testing.T) {
	PopulateTestTable(repo)
	defer CleanupRepo(repo)
	t.Run("return cart", func(t *testing.T) {
		got, err := repo.GetCart(1)
		if err != nil {
			t.Fail()
		}
		want := model.Cart{
			CustomerId: 1,
			ItemId:     []int64{1, 2, 3},
			Items: map[int64]model.CartItem{
				1: model.CartItem{
					UnitPrice: 1,
					Quantity:  1,
				},
				2: model.CartItem{
					UnitPrice: 2,
					Quantity:  1,
				},
				3: model.CartItem{
					UnitPrice: 3,
					Quantity:  1,
				},
			},
			CartT: model.CartTotal{
				TotalPrice: 6,
				Discount:   0,
				FinalPrice: 6,
			},
		}
		result := reflect.DeepEqual(got.ItemId, want.ItemId)
		if result != true {
			t.Errorf("cart items ids are not equal")
		}
		result = reflect.DeepEqual(got.Items, want.Items)
		if result != true {
			t.Errorf("cart items are not equal")
		}
	})
}
