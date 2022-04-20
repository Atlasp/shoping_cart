package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
// @Summary      Removes item
// @Description  Removes item from catalog
// @Tags         carts
// @Produce json
// @Success      200  {object} map[int64]model.CartItem
// @Router       /carts/{customer_id} [get]
func (h *handler) ReturnCart(c *gin.Context) {
	cartId := c.Param("customer_id")
	cid, err := strconv.ParseInt(cartId, 10, 64)
	cart, err := h.repo.GetCart(cid)
	if err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("%d: %s", http.StatusNotFound, err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, cart.Items)
	}
}

// ShowAccount godoc
// @Summary      Removes item
// @Description  Removes item from catalog
// @Tags         carts
// @Produce json
// @Success      200  {object}  map[int64]model.CartItem
// @Router       /carts/{customer_id}/items/{item_id} [post]
func (h *handler) ItemToCart(c *gin.Context) {
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
	err = h.repo.AddItemToCart(cid, iid)
	if err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("%d: %s", http.StatusNotFound, err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, fmt.Sprintf("%d: %s added to cart", http.StatusOK, itemId))
	}
}

// ShowAccount godoc
// @Summary      Removes item
// @Description  Removes item from catalog
// @Tags         carts
// @Success      200  {object} map[int64]model.CartItem
// @Router       /carts/{customer_id}/totals [post]
func (h *handler) ReturnCartTotal(c *gin.Context) {
	customerId := c.Param("customer_id")
	cid, err := strconv.ParseInt(customerId, 10, 64)
	cart, err := h.repo.GetCart(cid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("%d: %s", http.StatusBadRequest, err.Error()))
	} else {
		cart.GetCartTotal(h.repo.CartRules)
		c.JSON(http.StatusOK, cart.CartT)
	}
}

// ShowAccount godoc
// @Summary      Removes item
// @Description  Removes item from catalog
// @Tags         carts
// @Success      200  {object} map[int64]model.CartItem
// @Router       /carts/{customer_id}/orders [post]
func (h *handler) PlaceOrder(c *gin.Context) {
	customerId := c.Param("customer_id")
	cid, err := strconv.ParseInt(customerId, 10, 64)
	err = h.repo.PlaceOrder(cid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("%s", err.Error()))
	} else {
		c.JSON(http.StatusOK, fmt.Sprintf("Order has been placed"))
	}
}
