package clients

import (
	"context"
	"fmt"
	account "github.com/amrchnk/api-gateway/proto/account"
	"github.com/spf13/viper"
)

type AccountClient struct {
	account.AccountServiceClient
}

var accountClientConn = &AccountClient{}

func AccountClientExecutor() *AccountClient {
	return accountClientConn
}

func InitAccountClient(ctx context.Context) {
	conn := GRPCClientConnection(ctx, fmt.Sprintf("localhost:%s", viper.GetString("api.accountGrpcPort")))
	accountClientConn.AccountServiceClient = account.NewAccountServiceClient(conn)
}

func (ac *AccountClient) CreateAccountFunc(ctx context.Context, userId int64) (int64, error) {
	res, err := ac.CreateAccountByUserId(ctx, &account.CreateAccountRequest{UserId: userId})
	if err != nil {
		return 0, err
	}
	return res.AccountId, err
}

func (ac *AccountClient) DeleteAccountFunc(ctx context.Context, userId int64) (string, error) {
	res, err := ac.DeleteAccountByUserId(ctx, &account.DeleteAccountByUserIdRequest{UserId: userId})
	if err != nil {
		return "", err
	}
	return res.Message, err
}