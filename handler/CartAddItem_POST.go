package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Repository) ItemToCart(c *gin.Context) {
	customerId := c.Param("customer_id")
	itemId := c.Param("item_id")
	cid, err := strconv.ParseInt(customerId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("%d: %s", http.StatusBadRequest, err.Error()))
		return
	}
	iid, err := strconv.ParseInt(itemId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("%d: %s", http.StatusBadRequest, err.Error()))
		return
	}
	err = r.AddItemToCart(cid, iid)
	if err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("%d: %s", http.StatusNotFound, err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, fmt.Sprintf("%s added to cart", itemId))
	}
}
