package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Repository) Order(c *gin.Context) {
	customerId := c.Param("customer_id")
	cid, err := strconv.ParseInt(customerId, 10, 64)
	err = r.PlaceOrder(cid, r.CartRules)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("%s", err.Error()))
	} else {
		c.JSON(http.StatusOK, fmt.Sprintf("Order has been placed"))
	}
}
