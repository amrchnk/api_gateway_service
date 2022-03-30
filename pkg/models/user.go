package models

import "time"

type User struct {
	Id        int       `json:"-" db:"id"`
	Login     string    `json:"login" db:"login" binding:"required"`
	Password  string    `json:"password" db:"password_hash" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Role      int64     `json:"role_id" db:"role_id"`
}
