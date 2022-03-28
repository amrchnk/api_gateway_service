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
	router.Use(CORSMiddleware())
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	store.Options(sessions.Options{MaxAge:   60 * 60 * 48})
	router.Use(sessions.Sessions("userSession", store))

	api:=router.Group("/api/v1/auth")
	{
		api.POST("/sign-up", h.signUp)
		api.POST("/sign-in", h.signIn)
	}

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}