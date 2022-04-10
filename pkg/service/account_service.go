package service

import (
	"context"
	"github.com/amrchnk/api-gateway/internal/app/clients"
)

type AccountService struct {
	account *clients.AccountClient
}

func NewAccountService(account *clients.AccountClient) AccountService {
	return AccountService{account: account}
}

func (a AccountService) CreateAccount(ctx context.Context, userId int64) (int64, error) {
	return a.account.CreateAccountFunc(ctx,userId)
}

func (a AccountService) DeleteAccount(ctx context.Context, userId int64) (string, error) {
	return a.account.DeleteAccountFunc(ctx,userId)
}
