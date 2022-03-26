package service

import (
	"context"
	"github.com/amrchnk/api-gateway/internal/app/clients"
	"github.com/amrchnk/api-gateway/pkg/models"
)

type AuthService struct {
	auth *clients.AuthClient
}

func NewAuthService(auth *clients.AuthClient)AuthService{
	return AuthService{auth: auth}
}

func (a AuthService) SignUp(ctx context.Context,user models.User)(int,error){
	return a.auth.SignUpFunc(ctx,user.Login,user.Password)
}

func (a AuthService) SignIn(ctx context.Context, user models.User) (string, error){
	return a.auth.SignInFunc(ctx,user.Login,user.Password)
}

