package repostitory

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"revel_systems_shopping/src/model"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository() Repository {
	dbURL := "postgres://revel:postgres@localhost:5432/revel"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic("cannot connect to database")
	}

	err = db.AutoMigrate(
		&model.Item{},
		&model.Order{},
		&model.CartTable{},
	)
	if err != nil {
		panic("cannot migrate")
	}

	return Repository{DB: db}
}

func (r Repository) GetItem(id interface{}) (model.Item, error) {
	var item model.Item
	result := r.DB.First(&item, id)
	if result.Error != nil {
		return model.Item{}, result.Error
	}
	return item, nil
}

func (r Repository) DeleteItem(id string) error {
	result := r.DB.Delete(&model.Item{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r Repository) AddItem(item model.Item) error {
	result := r.DB.Create(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r Repository) PlaceOrder(id interface{}, d model.Discount) error {
	c, _ := r.GetCart(id)
	cartTotal := c.GetCartTotal(d)
	orderedItems := make([]int64, len(c.CartItems))
	for k, _ := range c.CartItems {
		orderedItems = append(orderedItems, k)
	}
	dt := time.Now()

	order := model.Order{
		CustomerId:    c.CustomerId,
		OrderedItems:  orderedItems,
		OrderTotal:    cartTotal.TotalPrice,
		OrderDiscount: cartTotal.Discount,
		OrderFinal:    cartTotal.FinalPrice,
		OrderDate:     dt,
	}

	cart := model.CartTable{
		CartID:     c.CartID,
		CustomerId: c.CustomerId,
		Items:      nil,
	}

	r.DB.Create(&order)
	r.DB.Save(&cart)

	return nil
}

func (r Repository) GetCart(id interface{}) (model.Cart, error) {
	var cartEntry model.CartTable
	result := r.DB.First(&cartEntry, id)
	if result.Error != nil {
		return model.NewCart(), nil
	}
	cart := cartEntry.ParseCartTable()

	for _, v := range cartEntry.Items {
		item, _ := r.GetItem(v)
		cart.AddItemToCart(item)
	}

	return cart, nil
}

// add item out of stock option
func (r Repository) AddItemToCart(cartId interface{}, itemId interface{}) {
	var cartTable model.CartTable
	r.DB.First(&cartTable, cartId)
	item, _ := r.GetItem(itemId)
	cartTable.Items = append(cartTable.Items, item.Id)
	item.Stock -= 1
	r.DB.Save(&item)
	r.DB.Save(&cartTable)
}
