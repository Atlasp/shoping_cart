package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
// @Summary      Returns cart.
// @Description  Returns cart items given customer ID.
// @Tags         Carts
// @Produce json
// @Param   item_id     path    string     true        "ID of the item to retrieve"
// @Success      200  {object} map[int64]model.CartItem
// @Failure      400 {string} string "Item doesn't exist"
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
// @Summary      Adds item to cart.
// @Description  Adds item given an ID to a given customer cart.
// @Tags         Carts
// @Produce json
// @Param   customer_id     path    string     true        "ID of the customer to whose cart the item is added"
// @Param   item_id     path    string     true        "ID of an item that is being added"
// @Success      200  {object}  map[int64]model.CartItem
// @Failure      400 {string} string "Item doesn't exist"
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
// @Summary      Returns cart total.
// @Description  Returns cart total for a given customer.
// @Tags         Carts
// @Produce json
// @Param   customer_id     path    string     true        "ID of the customer for whom to retrieve total"
// @Success      200  {object} map[int64]model.CartItem
// @Failure      400 {string} string "Item doesn't exist"
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
// @Summary      Places an order.
// @Description  Places an order for a given customers cart.
// @Tags         Carts
// @Produce json
// @Param   customer_id     path    string     true        "ID of the customer who is making the order"
// @Success      200  {object} map[int64]model.CartItem "Order has been place"
// @Failure      500 {string} string "Item doesn't exist"
// @Router       /carts/{customer_id}/orders [post]
func (h *handler) PlaceOrder(c *gin.Context) {
	customerId := c.Param("customer_id")
	cid, err := strconv.ParseInt(customerId, 10, 64)
	err = h.repo.PlaceOrder(cid)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("%d: %s", http.StatusBadRequest, err.Error()))
	} else {
		c.JSON(http.StatusOK, fmt.Sprintf("Order has been placed"))
	}
}
