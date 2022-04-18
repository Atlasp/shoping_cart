package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"revel_systems_shopping/src/repostitory"
)

func ReturnCart(r repostitory.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		cartId := c.Param("customer_id")
		cart, err := r.GetCart(cartId)
		if err != nil {
			c.JSON(http.StatusNotFound, fmt.Sprintf("%d: %s", http.StatusNotFound, err.Error()))
		} else {
			c.JSON(http.StatusOK, cart.CartItems)
		}
	}
}
