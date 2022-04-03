package models

import "time"

type User struct {
	Id        int64     `json:"id" db:"id"`
	Login     string    `json:"login" db:"login" binding:"required"`
	Username  string    `json:"username" db:"username" binding:"required"`
	Password  string    `json:"password" db:"password_hash" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	RoleId    int64     `json:"role_id" db:"role_id"`
}
