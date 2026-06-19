package models

import "time"

type PostView struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	PostID   uint      `gorm:"not null;index" json:"post_id"`
	ViewedAt time.Time `gorm:"autoCreateTime;index" json:"viewed_at"`
}
