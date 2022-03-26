package service

/*import (
	"context"
	"github.com/amrchnk/api-gateway/internal/app/clients"
	"github.com/amrchnk/api-gateway/pkg/models"
)

type Service struct {
	IAuthorization
}

type IAuthorization interface {
	SignUp(ctx context.Context,user models.User)(int,error)
	SignIn(ctx context.Context, user models.User) (string, error)
}

func NewService(auth *clients.AuthClient)*Service{
	return &Service{
		IAuthorization:NewAuthService(auth),
	}
}*/