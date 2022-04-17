package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"revel_systems_shopping/src/model"
	"revel_systems_shopping/src/repostitory"
)

func PlaceOrder(r repostitory.Repository, d model.Discount) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId := c.Param("customer_id")
		err := r.PlaceOrder(customerId, d)
		if err != nil {
			c.Status(http.StatusInternalServerError)
		}
		c.JSON(http.StatusOK, fmt.Sprintf("Order has been placed"))
	}
}
