package repository

import (
	"database/sql"

	"github.com/PetJs/blog-backend/internal/models"
)

type PostRepository struct {
	DB *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository{
	return &PostRepository{DB: db}
}

func (r *PostRepository) GetAllPosts() ([]models.Post, error){
	rows, err := r.DB.Query("SELECT id, title, content, author, created_at, updated_at FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *PostRepository) CreatePost(post models.Post) error {
	_, err := r.DB.Exec("INSERT INTO posts (title, content, author, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", post.Title, post.Content, post.Author, post.CreatedAt, post.UpdatedAt)
	return err
}