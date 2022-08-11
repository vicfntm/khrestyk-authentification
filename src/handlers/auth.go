package handlers

import (
	"net/http"

	"example.com/hello/src/exceptionhandlers"
	"example.com/hello/src/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {

	var input models.User

	if err := c.BindJSON(&input); err != nil {
		exceptionhandlers.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.services.Authorization.CreateUser(input)

	if err != nil {
		exceptionhandlers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) signIn(c *gin.Context) {

	var input models.LoginUserStruct

	if err := c.BindJSON(&input); err != nil {
		exceptionhandlers.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	token, expAt, err := h.services.Authorization.LoginUser(input)

	if err != nil {
		exceptionhandlers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token":   token,
		"expired": expAt,
	})
}
