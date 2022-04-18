package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"revel_systems_shopping/src/model"
	"revel_systems_shopping/src/repostitory"
)

func AddItem(r repostitory.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		item := model.Item{}
		err := c.Bind(&item)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		err = r.AddItem(item)
		if err != nil {
			if strings.Contains(err.Error(), "SQLSTATE 23505") {
				c.JSON(http.StatusBadRequest, fmt.Sprintf("Item id %d already exists", item.Id))
			} else {
				c.JSON(http.StatusInternalServerError, err.Error())
			}
		} else {
			createdItem, _ := r.GetItem(item.Id)
			c.JSON(http.StatusCreated, createdItem)
		}
	}
}
