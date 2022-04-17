package clients

import (
	"context"
	"github.com/amrchnk/api-gateway/pkg/models"
	"github.com/amrchnk/api-gateway/proto/account"
	"time"
)

func (ac *AccountClient) CreatePostFunc(ctx context.Context, post models.Post) (int64, error) {
	images := make([]*account.Image, 0, len(post.Images))
	for _, image := range post.Images {
		images = append(images, &account.Image{
			Link: image.Link,
		})
	}
	res, err := ac.CreatePost(ctx, &account.CreatePostRequest{
		Post: &account.Post{
			Title:       post.Title,
			Description: post.Description,
			Images:      images,
			AccountId:   post.AccountId,
		},
	})
	if err != nil {
		return 0, err
	}
	return res.Id, err
}

func (ac *AccountClient) DeletePostByIdFunc(ctx context.Context, postId int64) (string, error) {
	req, err := ac.DeletePostById(ctx, &account.DeletePostByIdRequest{
		Id: postId,
	})
	if err != nil {
		return "", err
	}
	return req.Message, err
}

func (ac *AccountClient) GetPostByIdFunc(ctx context.Context, postId int64) (models.Post, error) {
	req, err := ac.GetPostById(ctx, &account.GetPostByIdRequest{
		Id: postId,
	})

	if err != nil {
		return models.Post{}, err
	}

	resTime, _ := time.Parse("2006-01-02 15:04:05", req.Post.CreatedAt)

	images := make([]models.Image, 0, len(req.Post.Images))
	for i := range req.Post.Images {
		image := models.Image{
			Link: req.Post.Images[i].Link,
		}
		images = append(images, image)
	}

	post := models.Post{
		Id:          req.Post.Id,
		Title:       req.Post.Title,
		Description: req.Post.Description,
		CreatedAt:   resTime,
		AccountId:   req.Post.AccountId,
		Images:      images,
	}

	return post, err
}
