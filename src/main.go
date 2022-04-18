package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"revel_systems_shopping/src/handler"
	"revel_systems_shopping/src/repostitory"
)

// TODO: Add decimals to price calculations
// TODO: Write unit and integration tests
// TODO: Package API in docker container
// TODO: Create new cart if there is no item ?

func main() {

	cartRules := "cart_rules.json"
	dbConnection := "postgres://revel:postgres@localhost:5432/revel"
	repo := repostitory.NewRepository(dbConnection, cartRules)

	router := gin.Default()

	v1 := router.Group("/api/v1")

	{
		items := v1.Group("/items")
		{
			items.GET(":item_id", handler.ReturnItem(repo))
			items.POST("", handler.AddItem(repo))
			items.DELETE(":item_id", handler.DeleteItem(repo))
		}
		carts := v1.Group("/carts")
		{
			carts.GET(":customer_id", handler.ReturnCart(repo))
			carts.GET(":customer_id/totals", handler.ReturnBasketTotal(repo))
			carts.POST(":customer_id/items/:item_id", handler.AddItemToCart(repo))
			carts.POST(":customer_id/orders", handler.PlaceOrder(repo))
		}
	}

	// Run server
	log.Fatal(router.Run(":4141"))

}
