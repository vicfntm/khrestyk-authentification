package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) about(c *gin.Context) {
	id, _ := c.Get("user")

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
