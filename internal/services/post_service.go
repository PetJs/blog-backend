package services

import (
	"github.com/PetJs/blog-backend/internal/models"
	"github.com/PetJs/blog-backend/internal/repository"
)

type PostService struct {
	Repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{Repo: repo}
}

func (s *PostService) GetPosts() ([]models.Post, error) {
	return s.Repo.GetAllPosts()
}

func (s *PostService) GetPost(id string) (*models.Post, error) {
	return s.Repo.GetPostByID(id)
}

func (s *PostService) CreatePost(post models.Post) (*models.Post, error) {
	return s.Repo.CreatePost(&post)
}

func (s *PostService) UpdatePost(id string, updates map[string]interface{}) (*models.Post, error) {
	return s.Repo.UpdatePost(id, updates)
}

func (s *PostService) DeletePost(id string) error {
	return s.Repo.DeletePost(id)
}
