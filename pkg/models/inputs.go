package models

type SignUpRequest struct {
	Login string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Username  string    `json:"username" db:"username" binding:"required"`
}

type SignInRequest struct {
	Login string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}