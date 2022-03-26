package main

import (
	"context"
	"fmt"
	design_app "github.com/amrchnk/api-gateway"
	"github.com/amrchnk/api-gateway/internal/app/clients"
	"github.com/amrchnk/api-gateway/pkg/handler"
	"github.com/amrchnk/api-gateway/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func init() {
	err := initConfig()
	if err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	fmt.Println(viper.GetString("port"))
	err = godotenv.Load(filepath.Join("././", ".env"))
	if err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
}

func main(){
	ctx:=context.Background()
	clients.InitAuthClient(ctx)

	authService:=service.NewAuthService(clients.AuthClientExecutor())
	GwService:=service.NewApiGWService(authService)
	handlers := handler.NewHandler(GwService)

	srv := new(design_app.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Printf("App Started at the port %s",viper.GetString("port"))

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