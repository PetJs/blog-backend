package repository

import (
	"gorm.io/gorm"

	"github.com/PetJs/blog-backend/internal/models"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (r *PostRepository) CreatePost(post *models.Post) (*models.Post, error) {
	if err := r.DB.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (r *PostRepository) GetAllPublishedPosts() ([]models.Post, error) {
	var posts []models.Post
	if err := r.DB.Where("status = ?", "published").Order("created_at DESC").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) GetPostBySlug(slug string) (*models.Post, error) {
	var post models.Post
	if err := r.DB.Preload("Blocks", func(db *gorm.DB) *gorm.DB {
		return db.Order("position ASC")
	}).Where("slug = ?", slug).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) GetPostByID(id string) (*models.Post, error) {
	var post models.Post
	if err := r.DB.Preload("Blocks", func(db *gorm.DB) *gorm.DB {
		return db.Order("position ASC")
	}).First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) UpdatePost(id string, updates map[string]interface{}) (*models.Post, error) {
	var post models.Post
	if err := r.DB.First(&post, id).Error; err != nil {
		return nil, err
	}
	if err := r.DB.Model(&post).Updates(updates).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) SlugExists(slug string) bool {
	var count int64
	r.DB.Model(&models.Post{}).Where("slug = ?", slug).Count(&count)
	return count > 0
}

func (r *PostRepository) DeletePost(id string) error {
	var post models.Post
	if err := r.DB.First(&post, id).Error; err != nil {
		return err
	}
	return r.DB.Delete(&post).Error
}
