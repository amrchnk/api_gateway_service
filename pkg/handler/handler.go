package handler

import (
	"github.com/amrchnk/api-gateway/pkg/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Imp *service.Implementation
}

func NewHandler(Imp *service.Implementation)*Handler{
	return &Handler{Imp: Imp}
}

func (h *Handler)InitRoutes()*gin.Engine{
	router := gin.Default()
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	router.Use(sessions.Sessions("userSession", store))
	//r.Use(sessions.Session{})
	api:=router.Group("/api/test/auth")
	{
		api.POST("/sign-up", h.signUp)
		api.POST("/sign-in", h.signIn)
	}

	return router
}