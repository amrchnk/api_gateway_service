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

	router.MaxMultipartMemory = 8 << 20 //обозначим максимальный размер файлов в МБ, которые пользователь может отправлять в запросе
	router.Use(CORSMiddleware())

	docs.SwaggerInfo.BasePath = "/api/v1" //инициализируем сваггер
	v1 := router.Group("/api/v1")         //группируем роуты
	{
		api := v1.Group("/auth")
		{
			api.POST("/sign-up", h.signUp)
			api.POST("/sign-in", h.signIn)
			api.POST("/logout", h.userIdentity, h.logOut)
			api.POST("/refresh",h.refreshAccessToken)
		}

		user := v1.Group("/users", h.userIdentity)
		{
			user.GET("/:id", h.getUserById)
			user.GET("/", h.adminIdentity, h.getAllUsers)
			user.DELETE("/:id", h.adminIdentity, h.deleteUserById)
			user.PUT("/", h.updateUser)
		}

		post := v1.Group("/posts")
		{
			post.POST("/",h.userIdentity, h.createPost)
			post.DELETE(":id",h.userIdentity, h.deletePostById)
			post.GET(":id", h.getPostById)
			post.PUT(":id",h.userIdentity, h.updatePostById)
			post.GET("/users/:id", h.getAllUserPosts)
			post.GET("/users/", h.getAllUsersPosts)
		}

		account := v1.Group("/account", h.userIdentity)
		{
			account.GET(":id", h.getAccountByUserId)
			account.POST(":id", h.createAccountByUserId, h.adminIdentity)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-control-expose-headers", "SetInCache-Cookie")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-AccessToken, Authorization, accept, origin, Cache-Control, X-Requested-With,SetInCache-Cookie")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
