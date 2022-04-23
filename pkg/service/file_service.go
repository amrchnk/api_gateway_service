package service

import (
	"context"
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

type Media struct {
	Cloud *cloudinary.Cloudinary
}

func NewMediaUpload() Media {
	cloud, _ := NewCloud()
	return Media{
		Cloud: cloud,
	}
}

func (m Media) FilesUpload(files []models.File) ([]string, error) {
	links := make([]string, 0, len(files))
	//validate
	/*err := validate.Struct(files)
	if err != nil {
		return links, err
	}*/

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, file := range files {
		uploadParam, err := m.Cloud.Upload.Upload(ctx, file.File, uploader.UploadParams{
			Folder: "test",
		})
		if err != nil {
			return links, err
		}
		links = append(links, uploadParam.SecureURL)
	}

	return links, nil
}

/*func (Media) RemoteUpload(url models.Url) (string, error) {
	//validate
	err := validate.Struct(url)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, errUrl := helper.ImageUploadHelper(url.Url)
	if errUrl != nil {
		return "", err
	}
	return uploadUrl, nil
}*/
