package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"revel_systems_shopping/src/repostitory"
)

func DeleteItem(r repostitory.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("item_id")
		err := r.DeleteItem(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.Status(http.StatusOK)
	}
}
