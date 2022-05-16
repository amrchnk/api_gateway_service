package service

import (
	"context"
	"fmt"
	"github.com/amrchnk/api-gateway/pkg/models"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/go-playground/validator/v10"
	"log"
	"os"
	"strings"
	"time"
)

var (
	validate = validator.New()
)

type CloudService struct {
	*cloudinary.Cloudinary
}

func NewCloud() (*cloudinary.Cloudinary, error) {
	//create Cloud instance
	cld, err := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_CLOUD_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"))
	if err != nil {
		return nil, err
	}
	return cld, nil
}

func NewCloudService() CloudService {
	cloud, _ := NewCloud()
	return CloudService{
		cloud,
	}
}

func (m CloudService) UploadOneFile(path string, file models.File) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uploadParam, err := m.Upload.Upload(ctx, file.File, uploader.UploadParams{
		UseFilename: true,
		Folder:      path,
	})
	if err != nil {
		return "", err
	}
	link := uploadParam.SecureURL
	return link, nil
}

func (m CloudService) DeleteFiles(links []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, link := range links {
		arr := strings.Split(link, "/")
		fileName := strings.Split(arr[len(arr)-1], ".")[0]
		fmt.Println(arr[len(arr)-2] + fileName)
		result, err := m.Upload.Destroy(ctx, uploader.DestroyParams{
			PublicID: arr[len(arr)-2] +"/"+ fileName,
		})
		if err != nil {
			return fmt.Errorf("error while deleting file from remote server: %v", err)
		}
		fmt.Println(result.Result)
	}
	return nil
}

func (m CloudService) DeleteFile(publicID string) error {
	ctx := context.Background()

	result, err := m.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})
	if err != nil {
		log.Printf("[ERROR]: error while deleting file from remote server -  %v", err)
		return fmt.Errorf("error while deleting file from remote server: %v", err)
	}
	fmt.Println(result)
	return nil
}

func (m CloudService) FilesUpload(path string, files []models.File) ([]string, error) {
	links := make([]string, 0, len(files)) // пустой массив для ссылок

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //время, в течение которого мы должны получить ответ от сервера
	defer cancel()

	//проходимся по массиву входящих файлов и загружаем каждый файл из массива в хранилище
	for _, file := range files {
		uploadParam, err := m.Upload.Upload(ctx, file.File, uploader.UploadParams{
			Folder:   path,
		})

		if err != nil {
			return links, err
		}
		links = append(links, uploadParam.SecureURL) // добавляем ссылку на файл в массив
	}

	return links, nil
}
