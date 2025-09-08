package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"user_id"`
	Username  string    `gorm:"size:100;unique;not null" json:"username"`
	Email     string    `gorm:"size:255;unique;not null" json:"email"`
	Password  string    `gorm:"size:255;not null"`
	CreatedAt *time.Time `json:"created_at"`
}
