package exceptionhandlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type bException struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, bException{message})
}
