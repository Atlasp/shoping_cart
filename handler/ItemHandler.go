package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"revel_systems_shopping/model"
)

// ShowAccount godoc
// @Summary      Show an item
// @Description  get item by ID
// @Tags         items
// @Produce      json
// @Success      200  {object}  model.Item
// @Router       /items/{item_id} [get]
func (h *handler) GetItem(c *gin.Context) {
	id := c.Param("item_id")
	iid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("%d: %s", http.StatusBadRequest, err.Error()))
		return
	}
	item, err := h.repo.GetItem(iid)
	if err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("%d: %s", http.StatusNotFound, err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, item)
	}
}

func (h *handler) CreateItem(c *gin.Context) {
	item := model.Item{}
	err := c.Bind(&item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("%d: %s", http.StatusInternalServerError, err.Error()))
		return
	}
	err = h.repo.AddItem(item)
	if err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			c.JSON(http.StatusBadRequest, fmt.Sprintf("%d: Item id %d already exists", http.StatusBadRequest, item.Id))
			return
		} else {
			c.JSON(http.StatusInternalServerError, fmt.Sprintf("%d: %s", http.StatusInternalServerError, err.Error()))
			return
		}
	} else {
		createdItem, _ := h.repo.GetItem(item.Id)
		c.JSON(http.StatusCreated, createdItem)
	}
}

func (h *handler) RemoveItem(c *gin.Context) {
	id := c.Param("item_id")
	iid, err := strconv.ParseInt(id, 10, 64)
	_, err = h.repo.GetItem(iid)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("%d: item %s doesn't exist", http.StatusBadRequest, id))
		return
	}
	err = h.repo.DeleteItem(iid)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("%d: item %s doesn't exist", http.StatusBadRequest, id))
		return
	} else {
		c.JSON(http.StatusOK, fmt.Sprintf("%d: item %s deleted", http.StatusOK, id))
	}
}