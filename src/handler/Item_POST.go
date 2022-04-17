package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"revel_systems_shopping/src/model"
	"revel_systems_shopping/src/repostitory"
)

func AddItem(r repostitory.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		item := model.Item{}
		err := c.Bind(&item)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		err = r.AddItem(item)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.JSON(http.StatusCreated, item)
		}
	}
}
