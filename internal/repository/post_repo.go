package repository

import (
	"gorm.io/gorm"

	"github.com/PetJs/blog-backend/internal/models"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository{
	return &PostRepository{DB: db}
}

func (r *PostRepository) GetAllPosts() ([]models.Post, error){

	var posts []models.Post
	if err := r.DB.Order("created_at DESC").Find(&posts).Error; err != nil {
        return nil, err
    }
	return posts, nil
}

func (r *PostRepository) CreatePost(post *models.Post) error {
	return r.DB.Create(post).Error
}