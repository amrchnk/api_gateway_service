package service

import (
	"context"
	"errors"
	"github.com/amrchnk/api-gateway/internal/app/clients"
	"github.com/amrchnk/api-gateway/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
	RoleId int64 `json:"role_id"`
}

type AuthService struct {
	auth *clients.AuthClient
}

func NewAuthService(auth *clients.AuthClient) AuthService {
	return AuthService{auth: auth}
}

func (a AuthService) SignUp(ctx context.Context, user models.User) (int64, error) {
	return a.auth.SignUpFunc(ctx, user)
}

func (a AuthService) SignIn(ctx context.Context, login, password string) (string, error) {
	resp, err := a.auth.SignInFunc(ctx, login, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		resp.Id,
		resp.RoleId,
	})

	return token.SignedString([]byte(signingKey))
}

func (a AuthService) ParseToken(accessToken string) (*tokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return nil, errors.New("token claims are not of type *tokenClaims")
	}

	return claims, nil
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
