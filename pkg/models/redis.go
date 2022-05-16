package models

import "github.com/dgrijalva/jwt-go"

type Likes struct {
	UsersLikes []UserLike `json:"users_likes" db:"users_likes"`
}

type UserLike struct {
	Username     string `json:"username" db:"username" binding:"required"`
	ProfileImage string `json:"profile_image,omitempty" db:"profile_image"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
	RoleId int64 `json:"role_id"`
}

type RefreshTokenClaims struct {
	jwt.StandardClaims
}

type UserTokens struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
