package models

type SignUpRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" db:"username" binding:"required"`
}

type SignInRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreatePostRequest struct {
	Title       string   `json:"title" db:"title" binding:"required"`
	Description string   `json:"description" db:"description"`
	Images      []string `json:"images" binding:"required"`
}

type UpdateAccountRequest struct {
	UserId    int64     `json:"user_id" db:"user_id"`
	ProfileImage string    `json:"profile_image" db:"profile_image"`
}
