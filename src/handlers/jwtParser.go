package handlers

import (
	"net/http"
	"strings"

	"example.com/hello/src/exceptionhandlers"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	message             = "token is invalid"
	userNotfound        = "user not found"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		exceptionhandlers.NewErrorResponse(c, http.StatusInternalServerError, message)
		return
	}

	parts := strings.Split(header, " ")
	if len(parts) != 2 {
		exceptionhandlers.NewErrorResponse(c, http.StatusInternalServerError, message)
		return
	}

	userId, error := h.services.Authorization.ParseToken(parts[1])
	if error != nil {
		exceptionhandlers.NewErrorResponse(c, http.StatusInternalServerError, userNotfound)
		return
	}
	c.Set("user", userId)

}
