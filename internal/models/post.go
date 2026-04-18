package models

import "time"

type Post struct {
	ID                 uint       `gorm:"primaryKey" json:"id"`
	Title              string     `gorm:"size:255;not null" json:"title"`
	Slug               string     `gorm:"size:255;uniqueIndex;not null" json:"slug"`
	Excerpt            string     `gorm:"size:500" json:"excerpt"`
	CoverImage         string     `gorm:"size:500" json:"cover_image"`
	Status             string     `gorm:"size:50;default:'draft'" json:"status"` // draft | published
	ElevenLabsAudioURL string     `gorm:"size:500" json:"elevenlabs_audio_url"`
	Blocks             []Block    `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE" json:"blocks"`
	CreatedAt          *time.Time `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
}

type Block struct {
	ID               uint       `gorm:"primaryKey" json:"id"`
	PostID           uint       `gorm:"not null;index" json:"post_id"`
	Type             string     `gorm:"size:50;not null" json:"type"` // text | image | gif | audio
	Content          string     `gorm:"type:text" json:"content"`
	OriginalAudioURL string     `gorm:"size:500" json:"original_audio_url"`
	Position         int        `gorm:"not null;default:0" json:"position"`
	CreatedAt        *time.Time `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
}
