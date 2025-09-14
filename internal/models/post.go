package models 

import "time"

type Post struct {
	ID 			uint 		`gorm:"primaryKey" json:"id"`
	Title 		string 		`gorm:"size:255;not null" json:"title"`
	ImageURL 	string		`gorm:"size:255" json:"image_url"`
	Content 	string 		`gorm:"type:text;not null" json:"content"`
	User 		User		`gorm:"foreignKey:UserID" json:"user"`
	UserID      uint        `gorm:"not null" json:"user_id"`
	CreatedAt 	*time.Time	`json:"created_at"`	
	UpdatedAt 	*time.Time	`json:"updated_at"`	
}