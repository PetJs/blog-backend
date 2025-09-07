package models 

import "time"

type Post struct {
	ID 			uint 		`gorm:"primaryKey" json:"id"`
	Title 		string 		`gorm:"size:255;not null" json:"title"`
	Content 	string 		`gorm:"type:text;not null" json:"content"`
	Author 		string 		`gorm:"size:100;not null" json:"author"`
	CreatedAt 	*time.Time	`json:"created_at"`	
	UpdatedAt 	*time.Time	`json:"updated_at"`	
}