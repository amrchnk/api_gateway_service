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
	Categories  []int64  `json:"categories,omitempty"`
}

type UpdatePostRequest struct {
	Id          int64     `json:"id" db:"id"`
	Title       string   `json:"title,omitempty" db:"title"`
	Description string   `json:"description,omitempty" db:"description"`
	Images      []string `json:"images,omitempty"`
	Categories  []int64  `json:"categories,omitempty"`
}

type GetAllUsersPostsRequest struct {
	Offset  int64  `json:"offset"`
	Limit   int64  `json:"limit"`
	Sorting string `json:"sorting"`
}
