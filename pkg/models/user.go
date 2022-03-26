package models

import "time"

type User struct {
	Id        int       `json:"-" db:"id"`
	Login     string    `json:"login" db:"login"`
	Password  string    `json:"password_hash" db:"password_hash"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
