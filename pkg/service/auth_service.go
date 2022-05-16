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
	signingKey      = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL        = 12 * time.Hour
	refreshTokenTTL = 30 * 24 * time.Hour
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

func (a AuthService) SignIn(ctx context.Context, login, password string) (models.UserTokens, error) {
	var userTokens models.UserTokens
	resp, err := a.auth.SignInFunc(ctx, login, password)
	if err != nil {
		return models.UserTokens{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: resp.Id,
		RoleId: resp.RoleId,
	})
	userTokens.Token, err = token.SignedString([]byte(signingKey))
	if err != nil {
		return models.UserTokens{}, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.RefreshTokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})
	userTokens.RefreshToken, err = refreshToken.SignedString([]byte(signingKey))
	if err != nil {
		return models.UserTokens{}, err
	}

	return userTokens, nil
}

func (a AuthService) ParseToken(accessToken string) (*models.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	/*errr:=err.(*jwt.ValidationError).Errors
	fmt.Println(errr)*/
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.TokenClaims)
	if !ok {
		return nil, errors.New("token claims are not of type *tokenClaims")
	}

	return claims, nil
}

/*func (a AuthService) UpdateToken(accessToken string)(*models.TokenClaims,error){

}*/

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
