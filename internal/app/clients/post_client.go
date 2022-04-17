package clients

import (
	"context"
	"github.com/amrchnk/api-gateway/pkg/models"
	"github.com/amrchnk/api-gateway/proto/account"
)

func (ac *AccountClient) CreatePostFunc(ctx context.Context, post models.Post) (int64, error) {
	images:=make([]*account.Image,0,len(post.Images))
	for _,image:=range post.Images{
		images=append(images,&account.Image{
			Link: image.Link,
		})
	}
	res, err := ac.CreatePost(ctx, &account.CreatePostRequest{
		Post:&account.Post{
			Title: post.Title,
			Description: post.Description,
			Images: images,
			AccountId:post.AccountId,
		},
	})
	if err != nil {
		return 0, err
	}
	return res.Id, err
}

func (ac *AccountClient) DeletePostByIdFunc(ctx context.Context,postId int64)(string,error){
	req,err:=ac.DeletePostById(ctx,&account.DeletePostByIdRequest{
		Id: postId,
	})
	if err != nil {
		return "", err
	}
	return req.Message,err
}
