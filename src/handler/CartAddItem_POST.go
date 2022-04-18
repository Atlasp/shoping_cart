package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"revel_systems_shopping/src/repostitory"
)

func AddItemToCart(r repostitory.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId := c.Param("customer_id")
		itemId := c.Param("item_id")
		err := r.AddItemToCart(customerId, itemId)
		if err != nil {
			c.JSON(http.StatusNotFound, fmt.Sprintf("%s", err.Error()))
		} else {
			c.JSON(http.StatusOK, fmt.Sprintf("%s added to cart", itemId))
		}
	}
}
