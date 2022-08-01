package handlers

import (
	"net/http"
	"os"

	"example.com/hello/src/exceptionhandlers"
	"github.com/gin-gonic/gin"
)

const MESSAGE = "NOT ALLOWED"

func (h *Handler) IsAdmin(c *gin.Context) {

	token := c.Query("adminToken")
	if token != os.Getenv("ADMIN_TOKEN") {
		exceptionhandlers.NewErrorResponse(c, http.StatusForbidden, MESSAGE)
	}

}
