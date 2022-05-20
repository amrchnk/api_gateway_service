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

type TokenDetails struct {
	AccessUuid string
	UserId     int64 `json:"user_id"`
	RoleId     int64 `json:"role_id"`
}

type RefreshDetails struct {
	RefreshUuid string
	UserId      int64
}

type UserTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessUuid   string `json:"-"`
	RefreshUuid  string `json:"-"`
	AtExpires    int64  `json:"-"`
	RtExpires    int64  `json:"-"`
}
