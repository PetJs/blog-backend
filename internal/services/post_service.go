package services

import (

	"github.com/PetJs/blog-backend/internal/models"
	"github.com/PetJs/blog-backend/internal/repository"
)


type PostService struct {
	Repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService{
	return &PostService{Repo: repo}
}

func (s *PostService) GetPosts() ([]models.Post, error){
	return s.Repo.GetAllPosts()
}

func (s *PostService) CreatePost(post models.Post) error {
	return s.Repo.CreatePost(post)
}