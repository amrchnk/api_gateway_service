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
}

func NewApiGWService(auth AuthService, account AccountService, media CloudService, redis RedisService) *Implementation {
	return &Implementation{
		IAuthService:       auth,
		IAccountService:    account,
		ICloudinaryService: media,
		IRedisService:      redis,
	}
}

type IAuthService interface {
	SignUp(ctx context.Context, user models.User) (int64, error)
	SignIn(ctx context.Context, login, password string) (models.UserTokens, error)
	ParseToken(accessToken string) (*models.TokenClaims, error)
	GetUserById(ctx context.Context, id int64) (models.User, error)
	DeleteUserById(ctx context.Context, id int64) (string, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	UpdateUser(ctx context.Context, user models.UpdateUserResponse) (string, error)
}

type IAccountService interface {
	CreateAccount(ctx context.Context, userId int64) (int64, error)
	DeleteAccount(ctx context.Context, userId int64) (string, error)
	GetAccountByUserId(ctx context.Context, userId int64) (models.Account, error)

	CreatePost(ctx context.Context, post models.Post) (int64, error)
	DeletePostById(ctx context.Context, postId int64) (string, error)
	GetPostById(ctx context.Context, postId int64) (models.PostV2, error)
	UpdatePost(ctx context.Context, post models.UpdatePostRequest) (string, error)
	GetPostsByUserId(ctx context.Context, userId int64) ([]models.Post, error)
	GetAllUsersPosts(ctx context.Context, request models.GetAllUsersPostsRequest) ([]models.GetPostByIdResponse, error)

	GetImagesFromPost(ctx context.Context, postId int64) ([]models.Image, error)
}

type ICloudinaryService interface {
	UploadOneFile(path string, file models.File) (string, error)
	FilesUpload(path string, files []models.File) ([]string, error)
	DeleteFiles(links []string) error
	DeleteFile(publicID string) error
}

type IRedisService interface {
	GetFromCache(ctx context.Context, key string) ([]byte, error)
	SetInCache(ctx context.Context, key string, value interface{}, expiration time.Duration) error
}
