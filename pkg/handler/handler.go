package handler

import (
	"github.com/amrchnk/api-gateway/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService *service.AuthService
}

func NewHandler(authService *service.AuthService)*Handler{
	return &Handler{authService: authService}
}

func (h *Handler)InitRoutes()*gin.Engine{
	router:=gin.New()
	api:=router.Group("/api/test/auth")
	{
		api.POST("/sign-up", h.signUp)
		api.POST("/sign-in", h.signIn)
	}

	return router
}