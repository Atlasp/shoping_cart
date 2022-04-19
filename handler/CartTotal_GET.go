package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Repository) ReturnBasketTotal(c *gin.Context) {
	customerId := c.Param("customer_id")
	cid, err := strconv.ParseInt(customerId, 10, 64)
	cart, err := r.GetCart(cid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("%d: %s", http.StatusBadRequest, err.Error()))
	} else {
		c.JSON(http.StatusOK, cart.GetCartTotal(r.CartRules))
	}
}
