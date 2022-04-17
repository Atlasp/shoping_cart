package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"revel_systems_shopping/src/model"
	"revel_systems_shopping/src/repostitory"
)

func PlaceOrder(r repostitory.Repository, d model.Discount) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId := c.Param("customer_id")
		r.PlaceOrder(customerId, d)
		c.Status(http.StatusOK)
	}
}
