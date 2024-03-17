package models

import (
	"time"
)

func CreatePitch(
	id string,
	userId string,
	videoUrl string,
	title string,
	description string,
) *Pitch {
	return &Pitch{Id: id, Title: title, VideoUrl: videoUrl, Description: description, Status: "draft", UserId: userId}
}

type Pitch struct {
	Id          string    `gorm:"type:uuid;primary_key" json:"id"`
	Title       string    `gorm:"type:varchar(100);not null" json:"title"`
	Description string    `json:"description"`
	VideoUrl    string    `gorm:"not null" json:"video_url"`
	Status      string    `gorm:"type:varchar(100);not null;default:'draft'" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Upvote      int       `gorm:"default:0" json:"upvote"`
	Downvote    int       `gorm:"default:0" json:"downvote"`
	UserId      string    `gorm:"type:uuid;not null" json:"user_id"`
	User        User      `gorm:"foreignkey:UserId" json:"-"`
}
