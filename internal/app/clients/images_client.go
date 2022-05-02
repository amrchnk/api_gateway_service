package clients

import (
	"context"
	"github.com/amrchnk/api-gateway/pkg/models"
	"github.com/amrchnk/api-gateway/proto/account"
	"log"
)

func (ac *AccountClient) GetImagesFromPostFunc(ctx context.Context, postId int64) ([]models.Image, error) {
	var images []models.Image

	resp, err := ac.GetImagesFromPost(ctx, &account.GetImagesFromPostRequest{
		PostId: postId,
	})
	if err != nil {
		log.Printf("[ERROR]: %v", err)
		return images, err
	}
	if len(resp.Images) > 0 {
		images = make([]models.Image, 0, len(resp.Images))
		for i := range resp.Images {
			images = append(images, models.Image{
				Link: resp.Images[i].Link,
			})
		}
	}
	return images, err
}
