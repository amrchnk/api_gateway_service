package clients

import (
	"context"
	"fmt"
	"github.com/amrchnk/api-gateway/pkg/models"
	account "github.com/amrchnk/api-gateway/proto/account"
	"github.com/spf13/viper"
	"time"
)

type AccountClient struct {
	account.AccountServiceClient
}

var accountClientConn = &AccountClient{}

func AccountClientExecutor() *AccountClient {
	return accountClientConn
}

//ACCOUNT

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

func (ac *AccountClient) DeleteAccountByUserIdFunc(ctx context.Context, userId int64) (string, error) {
	res, err := ac.DeleteAccountByUserId(ctx, &account.DeleteAccountByUserIdRequest{UserId: userId})
	if err != nil {
		return "", err
	}
	return res.Message, err
}

func (ac *AccountClient) GetAccountByUserIdFunc(ctx context.Context, userId int64) (models.Account, error) {
	res, err := ac.GetAccountByUserId(ctx, &account.GetAccountByUserIdRequest{UserId: userId})
	if err != nil {
		return models.Account{}, err
	}
	resTime, _ := time.Parse("2006-01-02 15:04:05", res.Account.CreatedAt)

	return models.Account{
		UserId:    res.Account.UserId,
		Id:        res.Account.Id,
		ProfileImage: res.Account.ProfileImage,
		CreatedAt: resTime,
	}, err
}

func (ac *AccountClient) UpdateAccountByUserIdFunc(ctx context.Context, req models.UpdateAccountRequest) (string, error) {
	msg, err := ac.UpdateAccountByUserId(ctx, &account.UpdateAccountByUserIdRequest{NewInfo: &account.Account{
		UserId:       req.UserId,
		ProfileImage: req.ProfileImage,
	}})
	if err != nil {
		return "", err
	}

	return msg.Message,err
}
