package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Repository) ReturnCart(c *gin.Context) {
	cartId := c.Param("customer_id")
	cid, err := strconv.ParseInt(cartId, 10, 64)
	cart, err := r.GetCart(cid)
	if err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("%d: %s", http.StatusNotFound, err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, cart.CartItems)
	}
}
