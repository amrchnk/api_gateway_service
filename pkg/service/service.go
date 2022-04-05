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
		IAuthService: as,
	}
}

type IAuthService interface {
	SignUp(ctx context.Context, user models.User) (int64, error)
	SignIn(ctx context.Context, login, password string) (string, error)
	ParseToken(accessToken string) (*tokenClaims, error)
	GetUserById(ctx context.Context, id int64) (models.User, error)
	DeleteUserById(ctx context.Context, id int64) (string, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	UpdateUser(ctx context.Context, user models.UpdateUserResponse) (string, error)
}
