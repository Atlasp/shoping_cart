package repository

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"revel_systems_shopping/model"
)

type Repository struct {
	DB        *gorm.DB
	CartRules model.CartRules
}

func NewRepository(cr model.CartRules) Repository {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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

	return Repository{DB: db, CartRules: cr}
}

func (r Repository) GetItem(id int64) (model.Item, error) {
	var item model.Item
	result := r.DB.First(&item, id)
	if result.Error != nil {
		return model.Item{}, result.Error
	}
	return item, nil
}

func (r Repository) DeleteItem(id int64) error {
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

func (r Repository) PlaceOrder(customerId int64) error {
	var ct model.CartTable
	result := r.DB.First(&ct, "customer_id = ?", customerId)
	if result.Error != nil {
		return result.Error
	}
	c, err := r.GetCart(customerId)
	if err != nil {
		return err
	}
	c.GetCartTotal(r.CartRules)
	dt := time.Now()

	if c.CartT.FinalPrice <= r.CartRules.MaxBasketSize {

		order := model.Order{
			CustomerId:    c.CustomerId,
			OrderedItems:  ct.Items,
			OrderTotal:    c.CartT.TotalPrice,
			OrderDiscount: c.CartT.Discount,
			OrderFinal:    c.CartT.FinalPrice,
			OrderDate:     dt,
		}

		ct.Items = nil

		r.DB.Create(&order)
		r.DB.Save(&ct)

		return nil
	}

	return fmt.Errorf("cannot complete order, cart exceeds maximum limit of $%d", r.CartRules.MaxBasketSize)
}

func (r Repository) AddItemToCart(customerId int64, itemId int64) error {
	var cartTable model.CartTable
	result := r.DB.First(&cartTable, "customer_id = ?", customerId)
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
func (r Repository) GetCart(customerId int64) (model.Cart, error) {
	var cartEntry model.CartTable
	result := r.DB.First(&cartEntry, "customer_id = ?", customerId)
	if result.Error != nil {
		return model.Cart{}, result.Error
	}
	cart := cartEntry.ParseCartTable()

	for _, v := range cartEntry.Items {
		item, _ := r.GetItem(v)
		cart.AddItem(item)
	}

	return cart, nil
}
