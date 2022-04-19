package service

import (
	"context"
	"github.com/amrchnk/api-gateway/internal/app/clients"
	"github.com/amrchnk/api-gateway/pkg/models"
)

type AccountService struct {
	account *clients.AccountClient
}

func NewAccountService(account *clients.AccountClient) AccountService {
	return AccountService{account: account}
}

//ACCOUNT

func (a AccountService) CreateAccount(ctx context.Context, userId int64) (int64, error) {
	return a.account.CreateAccountFunc(ctx, userId)
}

func (a AccountService) DeleteAccount(ctx context.Context, userId int64) (string, error) {
	return a.account.DeleteAccountByUserIdFunc(ctx, userId)
}

func (a AccountService) GetAccountByUserId(ctx context.Context, userId int64) (models.Account, error) {
	return a.account.GetAccountByAccountIdFunc(ctx, userId)
}