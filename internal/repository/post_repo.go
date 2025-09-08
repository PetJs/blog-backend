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

func (r *PostRepository) DeletePost(postID string, userID uint) error {
	var post models.Post

	//find thepost by id
	if err := r.DB.First(&post, postID).Error; err != nil {
		return err
	}

	//Check if the post belongs to the user
	if post.UserID != userID {
		return gorm.ErrRecordNotFound
	}

	return r.DB.Delete(&post).Error
}

func (r *PostRepository) GetPostByUser(userID uint) ([]models.Post, error){
	var posts []models.Post

	if err := r.DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}