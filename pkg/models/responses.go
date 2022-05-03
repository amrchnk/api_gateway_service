package models

import "time"

type SignUpResponse struct {
	UserId    int64 `json:"user_id"`
	AccountId int64 `json:"account_id"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

type UpdateUserResponse struct {
	Id       int64  `json:"id" db:"id" binding:"required"`
	Login    string `json:"login" db:"login"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password_hash"`
	RoleId   int64  `json:"role_id" db:"role_id"`
}

type GetAllUsersResponse struct {
	Data []User `json:"data"`
}

type GetPostByIdResponse struct {
	Id          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	Categories  []int64   `json:"categories,omitempty"`
	Images      []string  `json:"images"`
}

type GetAllUserPostsResponse struct {
	Posts []GetPostByIdResponse `json:"user_posts"`
}

type GetAllUsersPosts struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at"`
	Images      []string  `json:"images" db:"image"`
	Categories  []string   `json:"categories,omitempty" db:"category"`
	UserId     int64     `json:"user_id" db:"user_id" binding:"required"`
	Username  string    `json:"username" db:"username" binding:"required"`
}

type GetAllUsersPostsResponse struct {
	Posts []GetAllUsersPosts `json:"posts"`
}