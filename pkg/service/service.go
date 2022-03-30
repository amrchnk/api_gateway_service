package service

import (
	"context"
	"github.com/amrchnk/api-gateway/pkg/models"
)

type Implementation struct {
	IAuthService
}

func NewApiGWService(as AuthService) *Implementation {
	return &Implementation{
		IAuthService:     as,
	}
}

type IAuthService interface {
	SignUp(ctx context.Context,user models.User)(int,error)
	SignIn(ctx context.Context, login,password string) (string, error)
}