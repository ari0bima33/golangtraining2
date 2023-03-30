package socialmedia

import "time"

type Socialmedia struct {
	ID              int64      `json:"id" db:"id"`
	Name            string     `json:"name" db:"name"`
	SocialMediaUrl  string     `json:"social_media_url" db:"social_media_url"`
	UserID          int64      `json:"user_id" db:"user_id"`
	Username        string     `json:"username" db:"username"`
	Email           string     `json:"email" db:"email"`
	ProfileImageUrl string     `json:"profile_image_url" db:"profile_image_url"`
	CreatedAt       time.Time  `json:"-" db:"created_at"`
	UpdatedAt       *time.Time `json:"-" db:"updated_at"`
}

type SocialmediaList struct {
	ID             int64      `json:"id"`
	Name           string     `json:"name" `
	SocialMediaUrl string     `json:"social_media_url" `
	UserID         int64      `json:"user_id"`
	CreatedAt      time.Time  `json:"created_at" `
	UpdatedAt      *time.Time `json:"updated_at" `
	User           User       `json:"user"`
}

type User struct {
	ID              int64  `json:"id" `
	Username        string `json:"username" `
	ProfileImageUrl string `json:"profile_image_url"`
}

type CreateSocialmediaRequest struct {
	Name           string `json:"name" validate:"empty=false"`
	SocialMediaUrl string `json:"social_media_url" validate:"empty=false"`
}

type CreateSocialmediaResponse struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         int64     `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type UpdateSocialmediaRequest struct {
	Name           string `json:"name" validate:"empty=false"`
	SocialMediaUrl string `json:"social_media_url" validate:"empty=false"`
}

type UpdateSocialmediaResponse struct {
	ID             int64      `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserID         int64      `json:"user_id"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

type DeleteSocialmediaResponse struct {
	Message string `json:"message"`
}
