package comment

import "time"

type Comment struct {
	ID           int64      `json:"id" db:"id"`
	PhotoID      int64      `json:"photo_id" db:"photo_id"`
	Message      string     `json:"message" db:"message"`
	UserID       int64      `json:"user_id" db:"user_id"`
	Username     string     `json:"username" db:"username"`
	Email        string     `json:"email" db:"email"`
	PhotoTitle   string     `json:"photo_title" db:"photo_title"`
	PhotoCaption string     `json:"photo_caption" db:"photo_caption"`
	PhotoURL     string     `json:"photo_url" db:"photo_url"`
	PhotoUserID  int64      `json:"photo_user_id" db:"photo_user_id"`
	CreatedAt    time.Time  `json:"-" db:"created_at"`
	UpdatedAt    *time.Time `json:"-" db:"updated_at"`
}

type CommentList struct {
	ID        int64      `json:"id"`
	Message   string     `json:"message" `
	PhotoID   int64      `json:"photo_id" `
	UserID    int64      `json:"user_id"`
	CreatedAt time.Time  `json:"created_at" `
	UpdatedAt *time.Time `json:"updated_at" `
	User      User       `json:"user"`
	Photo     Photo      `json:"photo"`
}

type User struct {
	ID       int64  `json:"id" `
	Username string `json:"username" `
	Email    string `json:"email" `
}

type Photo struct {
	ID       int64  `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	Caption  string `json:"caption" db:"caption"`
	PhotoURL string `json:"photo_url" db:"photo_url"`
	UserID   int64  `json:"user_id" db:"user_id"`
}

type CreateCommentRequest struct {
	Message string `json:"message" validate:"empty=false"`
	PhotoID int64  `json:"photo_id"`
}

type CreateCommentResponse struct {
	ID        int64     `json:"id"`
	Message   string    `json:"message" `
	PhotoID   int64     `json:"photo_id" `
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateCommentRequest struct {
	Message string `json:"message" validate:"empty=false"`
}

type UpdateCommentResponse struct {
	ID        int64      `json:"id"`
	Message   string     `json:"message" `
	PhotoID   int64      `json:"photo_id" `
	UserID    int64      `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type DeleteCommentResponse struct {
	Message string `json:"message"`
}
