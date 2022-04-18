package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"revel_systems_shopping/src/repostitory"
)

func DeleteItem(r repostitory.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("item_id")
		_, err := r.GetItem(id)
		if err != nil {
			c.JSON(http.StatusNotFound, fmt.Sprintf("%d: item %s doesn't exist", http.StatusNotFound, id))
			return
		}
		err = r.DeleteItem(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, fmt.Sprintf("%s", err.Error()))
			return
		} else {
			c.JSON(http.StatusOK, fmt.Sprintf("item %s deleted", id))
		}
	}
}
