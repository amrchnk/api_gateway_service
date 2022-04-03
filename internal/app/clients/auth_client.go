package clients

import (
	"context"
	"fmt"
	"github.com/amrchnk/api-gateway/pkg/models"
	auth "github.com/amrchnk/api-gateway/proto/auth"
	"github.com/spf13/viper"
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

func (ac *AuthClient) SignUpFunc(ctx context.Context, login, password string) (int64, error) {
	user := auth.User{
		Login:    login,
		Password: password,
	}
	res, err := ac.SignUp(ctx, &auth.SignUpRequest{User: &user})
	if err != nil {
		return 0, err
	}

	regResp := res.Slug

	return regResp, nil
}

func (ac *AuthClient) SignInFunc(ctx context.Context, login, password string) (models.User, error) {
	var user models.User
	authReq := auth.SignInRequest{Login: login, Password: password}
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
			Username: resp.User[i].Username,
			RoleId:   resp.User[i].UserRoleId,
		})
	}

	return users, err
}
