package service

import (
	"context"
	"github.com/amrchnk/api-gateway/pkg/models"
)

func (a AccountService) GetImagesFromPost(ctx context.Context, postId int64) ([]models.Image, error) {
	return a.account.GetImagesFromPostFunc(ctx, postId)
}
