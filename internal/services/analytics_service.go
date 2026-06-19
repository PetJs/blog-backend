package services

import (
	"github.com/PetJs/blog-backend/internal/repository"
)

type AnalyticsService struct {
	Repo        *repository.AnalyticsRepository
	PostRepo    *repository.PostRepository
}

func NewAnalyticsService(repo *repository.AnalyticsRepository, postRepo *repository.PostRepository) *AnalyticsService {
	return &AnalyticsService{Repo: repo, PostRepo: postRepo}
}

type DashboardStats struct {
	Posts    repository.PostCounts       `json:"posts"`
	Blocks   []repository.BlockTypeStat  `json:"blocks_by_type"`
	Views    repository.ViewCounts       `json:"views"`
	TopPosts []repository.TopPost        `json:"top_posts"`
}

func (s *AnalyticsService) GetDashboardStats() (*DashboardStats, error) {
	posts, err := s.Repo.GetPostCounts()
	if err != nil {
		return nil, err
	}

	blocks, err := s.Repo.GetBlockTypeStats()
	if err != nil {
		return nil, err
	}

	views, err := s.Repo.GetViewCounts()
	if err != nil {
		return nil, err
	}

	topPosts, err := s.Repo.GetTopPosts(5)
	if err != nil {
		return nil, err
	}

	return &DashboardStats{
		Posts:    posts,
		Blocks:   blocks,
		Views:    views,
		TopPosts: topPosts,
	}, nil
}

func (s *AnalyticsService) RecordView(slug string) error {
	post, err := s.PostRepo.GetPostBySlug(slug)
	if err != nil {
		return err
	}
	return s.Repo.RecordView(post.ID)
}
