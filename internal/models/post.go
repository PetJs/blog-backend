package models 

import "time"

type Post struct {
	ID 			uint 		`gorm:"primaryKey"`
	Title 		string 		`gorm:"size:255;not null"`
	Content 	string 		`gorm:"type:text;not null"`
	Author 		string 		`gorm:"size:100;not null "`
	CreatedAt 	time.Time 	`gorm:"autoCreateTime"`
	UpdatedAt 	time.Time 	`gorm:"autoUpdateTime"`
}