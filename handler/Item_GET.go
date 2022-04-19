package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
// @Summary      Show an item
// @Description  get item by ID
// @Tags         items
// @Produce      json
// @Success      200  {object}  model.Item
// @Router       /items/{item_id} [get]
func (r *Repository) ReturnItem(c *gin.Context) {
	id := c.Param("item_id")
	iid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("%d: %s", http.StatusBadRequest, err.Error()))
		return
	}
	item, err := r.GetItem(iid)
	if err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("%d: %s", http.StatusNotFound, err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, item)
	}
}
