package models

import "time"

type Admin struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Email     string     `gorm:"size:255;unique;not null" json:"email"`
	Password  string     `gorm:"size:255;not null" json:"-"`
	CreatedAt *time.Time `json:"created_at"`
}
