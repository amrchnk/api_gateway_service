package main

import (
	"context"
	design_app "github.com/amrchnk/api-gateway"
	"github.com/amrchnk/api-gateway/internal/app/clients"
	"github.com/amrchnk/api-gateway/pkg/handler"
	"github.com/amrchnk/api-gateway/pkg/repository/cache"
	"github.com/amrchnk/api-gateway/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var (
	cacheClient *cache.RedisClient
)

func init() {
	err := initConfig()
	if err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	err = godotenv.Load(filepath.Join("././", ".env"))
	if err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
}

// @title Designers App Swagger API
// @version 1.0
// @description Swagger API for Designers App
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email mirchenko1702@gmail.com

// @host      localhost:8000
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	ctx := context.Background()
	clients.InitAuthClient(ctx)
	clients.InitAccountClient(ctx)
	cc, err := cache.NewRedisClient(ctx, &cache.Options{
		Addr:     viper.GetString("redis.host"),
		Password: viper.GetString("redis.password"),
		DB:       0,
	})
	if err != nil {
		log.Fatalf("failed to initialize redis: %s", err.Error())
	}

	cacheClient = cc
	authService := service.NewAuthService(clients.AuthClientExecutor())
	tokenService, err := service.NewTokenService(os.Getenv("ACCESS_KEY"), os.Getenv("REFRESH_KEY"))
	if err != nil {
		log.Fatalf("failed to initialize tokens: %s", err.Error())
	}
	accountService := service.NewAccountService(clients.AccountClientExecutor())
	mediaService := service.NewCloudService()
	redisService := service.NewRedisService(cacheClient)

	GwService := service.NewApiGWService(authService, accountService, mediaService, &redisService, tokenService)
	handlers := handler.NewHandler(GwService)

	srv := new(design_app.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Printf("App Started at the port %s", viper.GetString("port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("././configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
