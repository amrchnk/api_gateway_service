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

func (ac *AccountClient) UpdatePostFunc(ctx context.Context, post models.UpdatePostRequestTextData) (string, error) {

	resp, err := ac.UpdatePostById(ctx, &account.UpdatePostByIdRequest{
		PostId:      post.Id,
		Title:       post.Title,
		Categories:  post.Categories,
		Description: post.Description,
		Images:      post.Images,
	})
	if err != nil {
		log.Printf("[ERROR]: %v", err)
		return "", err
	}

	return resp.Message, err
}

func (ac *AccountClient) GetPostByIdFunc(ctx context.Context, postId int64) (models.PostV2, error) {
	req, err := ac.GetPostById(ctx, &account.GetPostByIdRequest{
		Id: postId,
	})

	if err != nil {
		log.Printf("[ERROR]: %v", err)
		return models.PostV2{}, err
	}

	createResTime, _ := time.Parse("2006-01-02 15:04:05", req.Post.CreatedAt)
	updateResTime, _ := time.Parse("2006-01-02 15:04:05", req.Post.UpdatedAt)

	post := models.PostV2{
		Id:          req.Post.Id,
		Title:       req.Post.Title,
		Description: req.Post.Description,
		CreatedAt:   createResTime,
		UpdatedAt:   updateResTime,
		UserId:      req.Post.UserId,
		Images:      req.Post.Images,
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

func (ac *AccountClient) GetAllUsersPostsFunc(ctx context.Context, request models.GetAllUsersPostsRequest) ([]models.GetPostByIdResponse, error) {
	resp, err := ac.GetAllUsersPosts(ctx, &account.GetAllUsersPostsRequest{
		Limit:   request.Limit,
		Offset:  request.Offset,
		Sorting: request.Sorting,
	})
	if err != nil {
		log.Printf("[ERROR]: %v", err)
		return nil, err
	}

	posts := make([]models.GetPostByIdResponse, 0, len(resp.Posts))
	for _, post := range resp.Posts {
		createResTime, _ := time.Parse("2006-01-02 15:04:05", post.CreatedAt)
		updateResTime, _ := time.Parse("2006-01-02 15:04:05", post.UpdatedAt)
		posts = append(posts, models.GetPostByIdResponse{
			Id:          post.Id,
			UserId:      post.UserId,
			Images:      post.Images,
			Categories:  post.Categories,
			CreatedAt:   createResTime,
			UpdatedAt:   updateResTime,
			Title:       post.Title,
			Description: post.Description,
		})
	}

	return posts, nil
}
