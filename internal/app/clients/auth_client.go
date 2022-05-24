package clients

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/amrchnk/api-gateway/pkg/models"
	"github.com/amrchnk/api-gateway/proto/auth"
	"github.com/spf13/viper"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
)

type AuthClient struct {
	auth.AuthServiceClient
}

var authClientConn = &AuthClient{}

func AuthClientExecutor() *AuthClient {
	return authClientConn
}

func InitAuthClient(ctx context.Context) {
	conn := GRPCClientConnection(ctx, fmt.Sprintf("localhost:%s", viper.GetString("api.authGrpcPort")))
	authClientConn.AuthServiceClient = auth.NewAuthServiceClient(conn)
}

func (ac *AuthClient) SignUpFunc(ctx context.Context, user models.User) (int64, error) {
	res, err := ac.SignUp(ctx, &auth.SignUpRequest{User: &auth.User{
		Login:    user.Login,
		Password: generatePasswordHash(user.Password),
		Username: user.Username,
	}})
	if err != nil {
		return 0, err
	}

	regResp := res.Slug

	return regResp, nil
}

func (ac *AuthClient) SignInFunc(ctx context.Context, login, password string) (models.User, error) {
	var user models.User
	authReq := auth.SignInRequest{Login: login, Password: generatePasswordHash(password)}
	resp, err := ac.SignIn(ctx, &authReq)
	if err != nil {
		return user, err
	}
	user = models.User{
		Id:     resp.User.Slug,
		RoleId: resp.User.UserRoleId,
	}

	return user, err
}

func (ac *AuthClient) GetUserByIdFunc(ctx context.Context, id int64) (models.User, error) {
	var user models.User
	req := auth.GetUserByIdRequest{Slug: id}
	resp, err := ac.GetUserById(ctx, &req)
	if err != nil {
		return user, err
	}

	user = models.User{
		Id:       resp.User.Slug,
		Username: resp.User.Username,
		Login:    resp.User.Login,
		Password: resp.User.Password,
		RoleId:   resp.User.UserRoleId,
		ProfileImage: resp.User.ProfileImage,
	}

	return user, err
}

func (ac *AuthClient) DeleteUserByIdFunc(ctx context.Context, id int64) (string, error) {
	req := auth.DeleteUserByIdRequest{Slug: id}
	resp, err := ac.DeleteUserById(ctx, &req)
	if err != nil {
		return "", err
	}

	msg := resp.Resp

	return msg, err
}

func (ac *AuthClient) UpdateUserFunc(ctx context.Context, user models.UpdateUserResponse) (string, error) {
	userReq := auth.User{
		Slug:       user.Id,
		Username:   user.Username,
		Login:      user.Login,
		Password:   user.Password,
		UserRoleId: user.RoleId,
		ProfileImage: user.ProfileImage,
	}
	req := auth.UpdateUserRequest{
		User: &userReq,
	}
	resp, err := ac.UpdateUser(ctx, &req)
	if err != nil {
		return "", err
	}
	return resp.Resp, err
}

func (ac *AuthClient) GetAllUsersFunc(ctx context.Context) ([]models.User, error) {
	req := auth.GetAllUsersRequest{}
	resp, err := ac.GetAllUsers(ctx, &req)
	if err != nil {
		return nil, err
	}

	users := make([]models.User, 0, len(resp.User))
	for i := range resp.User {
		users = append(users, models.User{
			Id:       resp.User[i].Slug,
			Login:    resp.User[i].Login,
			Password: resp.User[i].Password,
			ProfileImage: resp.User[i].ProfileImage,
			Username: resp.User[i].Username,
			RoleId:   resp.User[i].UserRoleId,
		})
	}

	return users, err
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
