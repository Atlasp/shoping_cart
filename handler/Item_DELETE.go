package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Repository) ItemDelete(c *gin.Context) {
	id := c.Param("item_id")
	iid, err := strconv.ParseInt(id, 10, 64)
	_, err = r.GetItem(iid)
	if err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("%d: item %s doesn't exist", http.StatusNotFound, id))
		return
	}
	err = r.DeleteItem(iid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("%s", err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, fmt.Sprintf("item %s deleted", id))
	}
}
