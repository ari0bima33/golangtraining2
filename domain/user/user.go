package user

import (
	"time"
)

type User struct {
	ID        int64      `json:"id" db:"id"`
	Username  string     `json:"username" db:"username"`
	Email     string     `json:"email" db:"email"`
	Password  string     `json:"password" db:"password"`
	Age       int        `json:"age" db:"age"`
	CreatedAt time.Time  `json:"-" db:"created_at"`
	UpdatedAt *time.Time `json:"-" db:"updated_at"`
}

type CreateUser struct {
	Username string
	Email    string
	Password string
	Age      int
}

type CreateUserRequest struct {
	Age      int    `json:"age" validate:"gte=8"`
	Email    string `json:"email" validate:"empty=false"`
	Password string `json:"password" validate:"empty=false"`
	Username string `json:"username" validate:"empty=false"`
}

type CreateUserResponse struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"empty=false"`
	Password string `json:"password" validate:"empty=false"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UpdateUserRequest struct {
	Email    string `json:"email" validate:"empty=false"`
	Username string `json:"username" validate:"empty=false"`
}

type UpdateUserResponse struct {
	ID        int64      `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	Age       int        `json:"age"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}
