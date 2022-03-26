package clients

import (
	"context"
	"fmt"
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
	conn := GRPCClientConnection(ctx, fmt.Sprintf("localhost:%s",viper.GetString("api.authGrpcPort")))
	authClientConn.AuthServiceClient = auth.NewAuthServiceClient(conn)
}

func (ac *AuthClient) SignUpFunc(ctx context.Context, login, password string) (int, error) {
	regReq := auth.SignUpRequest{Login: login, Password: password}
	res, err := ac.SignUp(ctx, &regReq)
	if err != nil {
		return 0, err
	}

	regResp := res.Slug

	return int(regResp), nil
}

func (ac *AuthClient) SignInFunc(ctx context.Context, login, password string) (string, error) {
	authReq := auth.SignInRequest{Login: login, Password: password}
	res, err := ac.SignIn(ctx, &authReq)
	if err != nil {
		return "", err
	}
	authResp := res.Session

	return authResp, nil
}
