package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"revel_systems_shopping/src/handler"
	"revel_systems_shopping/src/model"
	"revel_systems_shopping/src/repostitory"
)

func main() {
	discount := model.Discount{
		MaxBasketSize:     100,
		DiscountThreshold: 20,
		FreeItemThreshold: 5,
	}

	repo := repostitory.NewRepository()

	router := gin.Default()

	//Item endpoints
	router.GET("/item/:item_id", handler.ReturnItem(repo))
	router.POST("/item", handler.AddItem(repo))
	router.DELETE("/item/:item_id", handler.DeleteItem(repo))
	// Cart endpoints
	router.GET("/cart/:customer_id", handler.ReturnCart(repo))
	router.GET("/cart/:customer_id/totals", handler.ReturnBasketTotal(repo, discount))
	router.POST("/cart/:customer_id/items/:item_id", handler.AddItemToCart(repo))
	router.POST("/cart/:customer_id/orders", handler.PlaceOrder(repo, discount))

	// Run server
	log.Fatal(router.Run(":4141"))

}
