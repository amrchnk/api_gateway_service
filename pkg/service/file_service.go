package service

import (
	"context"
	"fmt"
	"github.com/amrchnk/api-gateway/pkg/models"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/go-playground/validator/v10"
	"os"
	"time"
)

var (
	validate = validator.New()
)

func NewCloud() (*cloudinary.Cloudinary, error) {
	//create Cloud instance
	cld, err := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_CLOUD_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"))
	if err != nil {
		return nil, err
	}
	return cld, nil
}

type FileService struct {
	Cloud *cloudinary.Cloudinary
}

func NewMediaUpload() FileService {
	cloud, _ := NewCloud()
	return FileService{
		Cloud: cloud,
	}
}

func (m FileService) FilesUpload(userId int64, files []models.File) ([]string, error) {
	links := make([]string, 0, len(files))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, file := range files {
		uploadParam, err := m.Cloud.Upload.Upload(ctx, file.File, uploader.UploadParams{
			Folder: fmt.Sprintf("user%d",userId),
		})
		if err != nil {
			return links, err
		}
		links = append(links, uploadParam.SecureURL)
	}

	return links, nil
}
