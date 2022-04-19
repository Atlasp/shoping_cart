package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "revel_systems_shopping/docs"
	"revel_systems_shopping/handler"
)

// TODO: Add decimals to price calculations
// TODO: Write unit and integration tests
// TODO: Package API in docker container
// TODO: Swagger documentation
// TODO: Create new cart if there is no item ?

// @title        Revel Shopping API
// @version      1.0
// @description  Shopping API for Revel.

// @contact.name   Simas Paulikas
// @contact.email  simaspaulikas@yahoo.com

// @host      localhost:4141
// @BasePath  /api/v1
func main() {

	cartRules := "cart_rules.json"
	dbConnection := "postgres://revel:postgres@localhost:5432/revel"
	repo := handler.NewRepository(dbConnection, cartRules)

	router := gin.Default()

	v1 := router.Group("/api/v1")

	{
		items := v1.Group("/items")
		{
			items.GET(":item_id", repo.ReturnItem)
			items.POST("", repo.CreateItem)
			items.DELETE(":item_id", repo.ItemDelete)
		}
		carts := v1.Group("/carts")
		{
			carts.GET(":customer_id", repo.ReturnCart)
			carts.GET(":customer_id/totals", repo.ReturnBasketTotal)
			carts.POST(":customer_id/items/:item_id", repo.ItemToCart)
			carts.POST(":customer_id/orders", repo.Order)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Run server
	log.Fatal(router.Run(":4141"))

}
