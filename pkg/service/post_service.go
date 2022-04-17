package service

import (
	"context"
	"github.com/amrchnk/api-gateway/pkg/models"
)

func (a AccountService) CreatePost(ctx context.Context, post models.Post) (int64, error) {
	return a.account.CreatePostFunc(ctx,post)
}
