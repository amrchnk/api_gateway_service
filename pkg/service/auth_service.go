package service

import (
	"context"
	"github.com/amrchnk/api-gateway/internal/app/clients"
	"github.com/amrchnk/api-gateway/pkg/models"
)

type AuthService struct {
	auth *clients.AuthClient
}

func NewAuthService(auth *clients.AuthClient) AuthService {
	return AuthService{auth: auth}
}

func (a AuthService) SignUp(ctx context.Context, user models.User) (int64, error) {
	return a.auth.SignUpFunc(ctx, user)
}

func (a AuthService) SignIn(ctx context.Context, login, password string) (models.User, error) {
	return a.auth.SignInFunc(ctx, login, password)
}

func (a AuthService) GetUserById(ctx context.Context, id int64) (models.User, error) {
	return a.auth.GetUserByIdFunc(ctx, id)
}

func (a AuthService) DeleteUserById(ctx context.Context, id int64) (string, error) {
	return a.auth.DeleteUserByIdFunc(ctx, id)
}

func (a AuthService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return a.auth.GetAllUsersFunc(ctx)
}

func (a AuthService) UpdateUser(ctx context.Context, user models.UpdateUserResponse) (string, error) {
	return a.auth.UpdateUserFunc(ctx, user)
}
