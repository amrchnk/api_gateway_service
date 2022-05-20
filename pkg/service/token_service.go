package service

import (
	"errors"
	"fmt"
	"github.com/amrchnk/api-gateway/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

const (
	tokenTTL        = 12 * time.Hour
	refreshTokenTTL = 30 * 24 * time.Hour
)

type TokenService struct {
	accessSigningKey  string
	refreshSigningKey string
}

func NewTokenService(accessSigningKey, refreshSigningKey string) (*TokenService, error) {
	if accessSigningKey == "" || refreshSigningKey == "" {
		return nil, errors.New("signing key shouldn't be empty")
	}
	return &TokenService{
		accessSigningKey:  accessSigningKey,
		refreshSigningKey: refreshSigningKey,
	}, nil
}

func (t *TokenService) CreateTokens(userId, RoleId int64) (models.UserTokens, error) {
	var userTokens models.UserTokens

	userTokens.AtExpires = time.Now().Add(tokenTTL).Unix()
	userTokens.AccessUuid = uuid.New().String()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: userTokens.AtExpires,
			Id:        userTokens.AccessUuid,
		},
		UserId: userId,
		RoleId: RoleId,
	})

	var err error
	userTokens.AccessToken, err = token.SignedString([]byte(t.accessSigningKey))
	if err != nil {
		return models.UserTokens{}, err
	}

	userTokens.RtExpires = time.Now().Add(refreshTokenTTL).Unix()
	userTokens.RefreshUuid = uuid.New().String()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshTokenTTL).Unix(),
			Id:        userTokens.RefreshUuid,
		},
		UserId: userId,
	})
	userTokens.RefreshToken, err = refreshToken.SignedString([]byte(t.refreshSigningKey))
	if err != nil {
		return models.UserTokens{}, err
	}

	return userTokens, nil
}

func (t *TokenService) ParseToken(accessToken string) (*models.TokenDetails, error) {
	token, err := jwt.ParseWithClaims(accessToken, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(t.accessSigningKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.TokenClaims)
	fmt.Println(claims)
	if !ok {
		return nil, errors.New("token claims are not of type *tokenClaims")
	}

	return &models.TokenDetails{
		AccessUuid: claims.Id,
		UserId:     claims.UserId,
		RoleId:     claims.RoleId,
	}, nil
}

func (t *TokenService) ParseRefreshToken(refreshToken string) (*models.RefreshDetails, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(t.refreshSigningKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.TokenClaims)
	fmt.Println(claims)
	if !ok {
		return nil, errors.New("token claims are not of type *tokenClaims")
	}

	return &models.RefreshDetails{
		RefreshUuid: claims.StandardClaims.Id,
		UserId: claims.UserId,
	}, nil
}
