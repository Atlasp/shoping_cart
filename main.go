package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "revel_systems_shopping/docs"
	"revel_systems_shopping/handler"
	"revel_systems_shopping/model"
	"revel_systems_shopping/repository"
)

// TODO: Package API in docker container

// @title        Revel Shopping API
// @version      1.0
// @description  Shopping API for Revel.

// @contact.name   Simas Paulikas
// @contact.email  simaspaulikas@yahoo.com

// @host      localhost:4141
// @BasePath  /api/v1
func main() {

	cartRules := model.NewCartRules()
	repo := repository.NewRepository(cartRules)

	h := handler.NewHandler(repo)

	router := gin.Default()

	v1 := router.Group("/api/v1")

	{
		items := v1.Group("/items")
		{
			items.GET(":item_id", h.GetItem)
			items.POST("", h.CreateItem)
			items.DELETE(":item_id", h.RemoveItem)
		}
		carts := v1.Group("/carts")
		{
			carts.GET(":customer_id", h.ReturnCart)
			carts.GET(":customer_id/totals", h.ReturnCartTotal)
			carts.POST(":customer_id/items/:item_id", h.ItemToCart)
			carts.POST(":customer_id/orders", h.PlaceOrder)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Run server
	log.Fatal(router.Run(":4141"))

}
