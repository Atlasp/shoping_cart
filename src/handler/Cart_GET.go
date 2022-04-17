package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"revel_systems_shopping/src/model"
	"revel_systems_shopping/src/repostitory"
)

func ReturnCart(r repostitory.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		paramId := c.Param("customer_id")
		id, err := strconv.Atoi(paramId)
		cartId := int64(id)
		if err != nil {
			c.Error(err)
			c.Status(http.StatusNotFound)
		} else {
			cart, err := r.GetCart(cartId)
			if err != nil {
				c.JSON(http.StatusOK, model.NewCart().CartItems)
			}
			c.JSON(http.StatusOK, cart.CartItems)
		}
	}
}
