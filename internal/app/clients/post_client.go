package clients

import (
	"context"
	"github.com/amrchnk/api-gateway/pkg/models"
	"github.com/amrchnk/api-gateway/proto/account"
	"log"
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
			Categories:  post.Categories,
			AccountId:   post.AccountId,
		},
	})
	if err != nil {
		return 0, err
	}
	return res.Id, err
}

func (ac *AccountClient) DeletePostByIdFunc(ctx context.Context, postId int64) (string, error) {
	resp, err := ac.DeletePostById(ctx, &account.DeletePostByIdRequest{
		Id: postId,
	})
	if err != nil {
		log.Printf("[ERROR]: %v", err)
		return "", err
	}
	return resp.Message, err
}

func (ac *AccountClient) UpdatePostFunc(ctx context.Context, post models.Post) (string, error) {
	images := make([]*account.Image, 0, len(post.Images))
	for _, image := range post.Images {
		images = append(images, &account.Image{
			Link: image.Link,
		})
	}
	resp, err := ac.UpdatePostById(ctx, &account.UpdatePostByIdRequest{
		Post: &account.Post{
			Id:          post.Id,
			Title:       post.Title,
			Categories:  post.Categories,
			Description: post.Description,
			Images:      images,
		},
	})
	if err != nil {
		log.Printf("[ERROR]: %v", err)
		return "", err
	}

	return resp.Message, err
}

func (ac *AccountClient) GetPostByIdFunc(ctx context.Context, postId int64) (models.Post, error) {
	req, err := ac.GetPostById(ctx, &account.GetPostByIdRequest{
		Id: postId,
	})

	if err != nil {
		log.Printf("[ERROR]: %v", err)
		return models.Post{}, err
	}

	resTime, _ := time.Parse("2006-01-02 15:04:05", req.Post.CreatedAt)

	images := make([]models.Image, 0, len(req.Post.Images))
	for _, image := range req.Post.Images {
		imageResp := models.Image{
			Id:     image.Id,
			Link:   image.Link,
			PostId: image.PostId,
		}
		images = append(images, imageResp)
	}

	post := models.Post{
		Id:          req.Post.Id,
		Title:       req.Post.Title,
		Description: req.Post.Description,
		CreatedAt:   resTime,
		AccountId:   req.Post.AccountId,
		Images:      images,
		Categories:  req.Post.Categories,
	}

	return post, err
}

func (ac *AccountClient) GetPostsByUserIdFunc(ctx context.Context, userId int64) ([]models.Post, error) {
	req, err := ac.GetPostsByUserId(ctx, &account.GetUserPostsRequest{
		UserId: userId,
	})

	if err != nil {
		return []models.Post{}, err
	}

	posts := make([]models.Post, 0, len(req.Posts))
	for _, post := range req.Posts {
		images := make([]models.Image, 0, len(post.Images))
		for _, image := range post.Images {
			images = append(images, models.Image{
				Id:     image.Id,
				Link:   image.Link,
				PostId: image.PostId,
			})
		}

		resTime, _ := time.Parse("2006-01-02 15:04:05", post.CreatedAt)
		postResp := models.Post{
			Id:          post.Id,
			Title:       post.Title,
			Description: post.Description,
			CreatedAt:   resTime,
			Images:      images,
			AccountId:   post.AccountId,
		}
		posts = append(posts, postResp)
	}

	return posts, err
}

func (ac *AccountClient) GetAllUsersPostsFunc(ctx context.Context, request models.GetAllUsersPostsRequest) ([]models.GetAllUsersPosts, error) {
	resp, err := ac.GetAllUsersPosts(ctx, &account.GetAllUsersPostsRequest{
		Limit:   request.Limit,
		Offset:  request.Offset,
		Sorting: request.Sorting,
	})
	if err != nil {
		log.Printf("[ERROR]: %v", err)
		return nil, err
	}

	posts := make([]models.GetAllUsersPosts, 0, len(resp.Posts))
	for _, post := range resp.Posts {
		resTime, _ := time.Parse("2006-01-02 15:04:05", post.CreatedAt)
		posts = append(posts, models.GetAllUsersPosts{
			Id:          post.Id,
			UserId:      post.UserId,
			Images:      post.Images,
			Categories:  post.Categories,
			CreatedAt:   resTime,
			Title:       post.Title,
			Description: post.Description,
		})
	}

	return posts, nil
}
