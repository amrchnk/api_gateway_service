package service

import (
	"context"
	"github.com/amrchnk/api-gateway/pkg/models"
)

type Implementation struct {
	IAuthService
	IAccountService
	IFileService
}

func NewApiGWService(auth AuthService, account AccountService, media FileService) *Implementation {
	return &Implementation{
		IAuthService:    auth,
		IAccountService: account,
		IFileService:    media,
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

type IAccountService interface {
	CreateAccount(ctx context.Context, userId int64) (int64, error)
	DeleteAccount(ctx context.Context, userId int64) (string, error)
	GetAccountByUserId(ctx context.Context, userId int64) (models.Account, error)
	UpdateAccountByUserId(ctx context.Context, updateInfo models.UpdateAccountRequest) (string, error)

	CreatePost(ctx context.Context, post models.Post) (int64, error)
	DeletePostById(ctx context.Context, postId int64) (string, error)
	GetPostById(ctx context.Context, postId int64) (models.Post, error)
	UpdatePost(ctx context.Context, post models.Post) (string, error)
	GetPostsByUserId(ctx context.Context, userId int64) ([]models.Post, error)
	GetAllUsersPosts(ctx context.Context, request models.GetAllUsersPostsRequest) ([]models.GetAllUsersPosts, error)

	GetImagesFromPost(ctx context.Context, postId int64) ([]models.Image, error)
}

type IFileService interface {
	UploadOneFile(path string, file models.File) (string, error)
	FilesUpload(path string, files []models.File) ([]string, error)
	DeleteFiles(links []string) error
	DeleteFile(publicID string) error
}
