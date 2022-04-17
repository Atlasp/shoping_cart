package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"revel_systems_shopping/src/model"
	"revel_systems_shopping/src/repostitory"
)

func ReturnBasketTotal(r repostitory.Repository, d model.Discount) gin.HandlerFunc {
	return func(c *gin.Context) {
		paramId := c.Param("customer_id")
		id, err := strconv.Atoi(paramId)
		cartID := uint(id)
		if err != nil {
			c.Error(err)
			c.Status(http.StatusNotFound)
		} else {
			cart, _ := r.GetCart(cartID)
			c.JSON(http.StatusOK, cart.GetCartTotal(d))
		}
	}
}
