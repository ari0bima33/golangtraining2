package photo

import "time"

type Photo struct {
	ID        int64      `json:"id" db:"id"`
	Title     string     `json:"title" db:"title"`
	Caption   string     `json:"caption" db:"caption"`
	PhotoURL  string     `json:"photo_url" db:"photo_url"`
	UserID    int64      `json:"user_id" db:"user_id"`
	Username  string     `json:"username" db:"username"`
	Email     string     `json:"email" db:"email"`
	CreatedAt time.Time  `json:"-" db:"created_at"`
	UpdatedAt *time.Time `json:"-" db:"updated_at"`
}

type PhotoList struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title" `
	Caption   string     `json:"caption" `
	PhotoURL  string     `json:"photo_url"`
	UserID    int64      `json:"user_id"`
	CreatedAt time.Time  `json:"created_at" `
	UpdatedAt *time.Time `json:"updated_at" `
	User      User       `json:"user"`
}

type User struct {
	Username string `json:"username" `
	Email    string `json:"email" `
}

type CreatePhotoRequest struct {
	Title    string `json:"title" validate:"empty=false"`
	Caption  string `json:"caption" validate:"empty=false"`
	PhotoURL string `json:"photo_url" validate:"empty=false"`
}

type CreatePhotoResponse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption" `
	PhotoURL  string    `json:"photo_url"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdatePhotoRequest struct {
	Title    string `json:"title" validate:"empty=false"`
	Caption  string `json:"caption" validate:"empty=false"`
	PhotoURL string `json:"photo_url" validate:"empty=false"`
}

type UpdatePhotoResponse struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption" `
	PhotoURL  string     `json:"photo_url"`
	UserID    int64      `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}
