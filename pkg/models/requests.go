package models

import "mime/multipart"

type SignUpRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" db:"username" binding:"required"`
}

type SignInRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignOutRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type CreatePostRequest struct {
	Files    []*multipart.FileHeader `form:"Files" binding:"required" swaggerignore:"true"`
	PostInfo string                  `form:"PostInfo" binding:"required"`
}

type UpdateUserRequest struct {
	File []*multipart.FileHeader `swaggerignore:"true"`
	Json string
}

type CreatePostTextData struct {
	Title       string  `json:"title" db:"title" binding:"required"`
	Description string  `json:"description" db:"description"`
	Categories  []int64 `json:"categories,omitempty"`
}

type UpdatePostRequest struct {
	Files []*multipart.FileHeader
	Json  string
}

type UpdatePostRequestTextData struct {
	Id          int64    `json:"-"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Categories  []int64  `json:"categories,omitempty"`
	Images      []string `json:"images"`
}

type UpdateUserRequestTextData struct {
	Id           int64  `json:"-"`
	Login        string `json:"login" db:"login"`
	Username     string `json:"username" db:"username"`
	ProfileImage string `json:"profile_image" db:"profile_image"`
	Password     string `json:"password" db:"password_hash"`
	RoleId       int64  `json:"role_id" db:"role_id"`
}

type GetAllUsersPostsRequest struct {
	Offset  int64  `json:"offset"`
	Limit   int64  `json:"limit"`
	Sorting string `json:"sorting"`
}
