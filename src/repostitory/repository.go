package repostitory

import (
	"fmt"
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

func (r Repository) DeleteItem(id interface{}) error {
	result := r.DB.Delete(&model.Item{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Add item validation
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
	for k := range c.CartItems {
		orderedItems = append(orderedItems, k)
	}
	dt := time.Now()

	if cartTotal.FinalPrice < d.MaxBasketSize {

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

	return fmt.Errorf("cart exceeds maximum limit of %d", d.MaxBasketSize)
}

func (r Repository) GetCart(id interface{}) (model.Cart, error) {
	var cartEntry model.CartTable
	result := r.DB.First(&cartEntry, id)
	if result.Error != nil {
		return model.NewCart(), result.Error
	}
	cart := cartEntry.ParseCartTable()

	for _, v := range cartEntry.Items {
		item, _ := r.GetItem(v)
		cart.AddItemToCart(item)
	}

	return cart, nil
}

func (r Repository) AddItemToCart(cartId interface{}, itemId interface{}) error {
	var cartTable model.CartTable
	result := r.DB.First(&cartTable, cartId)
	if result.Error != nil {
		return result.Error
	}
	item, err := r.GetItem(itemId)
	if err != nil {
		return err
	}
	if item.Stock > 0 {
		cartTable.Items = append(cartTable.Items, item.Id)
		item.Stock -= 1
		r.DB.Save(&item)
		r.DB.Save(&cartTable)
		return nil
	}

	return fmt.Errorf("%s out of stock", item.Name)
}
