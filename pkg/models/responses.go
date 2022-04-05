package models

type UpdateUserResponse struct {
	Id        int64     `json:"id" db:"id" binding:"required"`
	Login     string    `json:"login" db:"login"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password" db:"password_hash"`
	RoleId    int64     `json:"role_id" db:"role_id"`
}

type GetAllUsersResponse struct {
	Data []User `json:"data"`
}