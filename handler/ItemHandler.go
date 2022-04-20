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
// @Param   item_id     path    string     true        "ID of the item to retrieve"
// @Success      200  {object}  model.Item
// @Failure      400 {string} string "Item doesn't exist"
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

// ShowAccount godoc
// @Summary      Create an item
// @Description  Create an item
// @Tags         items
// @Accept json
// @Produce json
// @Success      200  {object}  model.Item
// @Failure      400 {string} string "Item already exist"
// @Router       /items/{item_id} [post]
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

// ShowAccount godoc
// @Summary      Removes item
// @Description  Removes item from catalog
// @Tags         items
// @Produce json
// @Param   item_id     path    string     true        "ID of the item to delete"
// @Success      200  {object}  model.Item
// @Failure      400 {string} string "Item doesn't exist"
// @Router       /items/{item_id} [delete]
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
