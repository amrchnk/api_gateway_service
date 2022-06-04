package service

import (
	"context"
	"github.com/amrchnk/api-gateway/pkg/models"
	"time"
)

type Implementation struct {
	IAuthService
	IAccountService
	ICloudinaryService
	IRedisService
	ITokenService
}

func NewApiGWService(auth AuthService, account AccountService, media CloudService, redis *RedisService, token *TokenService) *Implementation {
	return &Implementation{
		IAuthService:       auth,
		IAccountService:    account,
		ICloudinaryService: media,
		IRedisService:      redis,
		ITokenService:      token,
	}
}

type ITokenService interface {
	ParseToken(accessToken string) (*models.TokenDetails, error)
	CreateTokens(userId, RoleId int64) (models.UserTokens, error)
	ParseRefreshToken(refreshToken string) (*models.RefreshDetails, error)
}

type IAuthService interface {
	SignUp(ctx context.Context, user models.User) (int64, error)
	SignIn(ctx context.Context, login, password string) (models.User, error)
	GetUserById(ctx context.Context, id int64) (models.User, error)
	DeleteUserById(ctx context.Context, id int64) (string, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	UpdateUser(ctx context.Context, user models.UpdateUserRequestTextData) (string, error)
}

type IAccountService interface {
	CreateAccount(ctx context.Context, userId int64) (int64, error)
	DeleteAccount(ctx context.Context, userId int64) (string, error)
	GetAccountByUserId(ctx context.Context, userId int64) (models.Account, error)

	CreatePost(ctx context.Context, post models.Post) (int64, error)
	DeletePostById(ctx context.Context, postId int64) (string, error)
	GetPostById(ctx context.Context, postId int64) (models.PostV2, error)
	UpdatePost(ctx context.Context, post models.UpdatePostRequestTextData) (string, error)
	GetPostsByUserId(ctx context.Context, userId int64) ([]models.Post, error)
	GetAllUsersPosts(ctx context.Context, request models.GetAllUsersPostsRequest) ([]models.GetPostByIdResponse, error)

	GetImagesFromPost(ctx context.Context, postId int64) ([]models.Image, error)
}

type ICloudinaryService interface {
	UploadOneFile(path string, file models.File) (string, error)
	FilesUpload(path string, files []models.File) ([]string, error)
	DeleteFiles(links []string) error
	DeleteFile(link string) error
}

type IRedisService interface {
	GetFromCache(ctx context.Context, key string) ([]byte, error)
	SetInCache(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	DeleteFromCache(ctx context.Context, key string) (int64, error)
}
