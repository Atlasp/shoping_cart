package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"revel_systems_shopping/src/repostitory"
)

func PlaceOrder(r repostitory.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId := c.Param("customer_id")
		err := r.PlaceOrder(customerId, r.CartRules)
		if err != nil {
			c.JSON(http.StatusInternalServerError, fmt.Sprintf("%s", err.Error()))
		} else {
			c.JSON(http.StatusOK, fmt.Sprintf("Order has been placed"))
		}
	}
}
