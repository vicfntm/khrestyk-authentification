package handlers

import (
	"example.com/hello/src/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/jwt-policy")
	{
		auth.POST("/sign-up", h.IsAdmin, h.signUp)
		auth.POST("/sign-in", h.CORSMiddleware, h.signIn)

	}

	api := router.Group("/api", h.UserIdentity)
	{
		api.GET("/about", h.about)
	}

	return router
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{
		services: services,
	}
}
