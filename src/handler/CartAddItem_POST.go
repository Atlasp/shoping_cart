package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"revel_systems_shopping/src/repostitory"
)

func AddItemToCart(r repostitory.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId := c.Param("customer_id")
		itemId := c.Param("item_id")
		r.AddItemToCart(customerId, itemId)
		c.Status(http.StatusOK)
	}
}
