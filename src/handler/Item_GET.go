package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"revel_systems_shopping/src/repostitory"
)

func ReturnItem(r repostitory.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("item_id")
		item, err := r.GetItem(id)
		if err != nil {
			c.JSON(http.StatusNotFound, fmt.Sprintf("%d: %s", http.StatusNotFound, err.Error()))
		} else {
			c.JSON(http.StatusOK, item)
		}
	}
}
