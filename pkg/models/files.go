package models

import "mime/multipart"

type File struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
	FileName string `json:"file_name,omitempty" validate:"required"`
}

type Url struct {
	Url string `json:"url,omitempty" validate:"required"`
}