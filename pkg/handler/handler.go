package handler

import (
	"github.com/amrchnk/api-gateway/docs"
	"github.com/amrchnk/api-gateway/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	Imp *service.Implementation
}

func NewHandler(Imp *service.Implementation) *Handler {
	return &Handler{Imp: Imp}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		api := v1.Group("/auth")
		{
			api.POST("/sign-up", h.signUp)
			api.POST("/sign-in", h.signIn)
			api.POST("/logout", h.logOut)
		}

		user := v1.Group("/users", h.userIdentity)
		{
			user.GET("/:id", h.AdminIdentity, h.getUserById)
			user.GET("/", h.AdminIdentity, h.getAllUsers)
			user.DELETE("/:id", h.AdminIdentity, h.deleteUserById)
			user.PUT("/", h.updateUser)
		}

		account := v1.Group("/account", h.userIdentity,h.AdminIdentity)
		{
			account.GET(":id",h.getAccountByUserId)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-control-expose-headers", "Set-Cookie")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With,Set-Cookie")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
