package service

import (
	"context"
	"github.com/amrchnk/api-gateway/pkg/models"
)

func (a AccountService) CreatePost(ctx context.Context, post models.Post) (int64, error) {
	return a.account.CreatePostFunc(ctx, post)
}

func (a AccountService) DeletePostById(ctx context.Context, postId int64) (string, error) {
	return a.account.DeletePostByIdFunc(ctx, postId)
}

func (a AccountService) GetPostById(ctx context.Context, postId int64) (models.PostV2, error) {
	return a.account.GetPostByIdFunc(ctx, postId)
}

func (a AccountService) UpdatePost(ctx context.Context, post models.UpdatePostRequestTextData) (string, error) {
	return a.account.UpdatePostFunc(ctx, post)
}

func (a AccountService) GetPostsByUserId(ctx context.Context, userId int64) ([]models.Post, error) {
	return a.account.GetPostsByUserIdFunc(ctx, userId)
}

func (a AccountService) GetAllUsersPosts(ctx context.Context, request models.GetAllUsersPostsRequest) ([]models.GetPostByIdResponse, error) {
	return a.account.GetAllUsersPostsFunc(ctx, request)
}
