package services

import (
	"fmt"
	"sort"
	"strings"

	"github.com/PetJs/blog-backend/internal/models"
	"github.com/PetJs/blog-backend/internal/repository"
	"github.com/PetJs/blog-backend/pkg/utils"
)

type PostService struct {
	Repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{Repo: repo}
}

func (s *PostService) CreatePost() (*models.Post, error) {
	slug := s.uniqueSlug("untitled")
	post := &models.Post{
		Title:  "Untitled",
		Slug:   slug,
		Status: "draft",
	}
	return s.Repo.CreatePost(post)
}

func (s *PostService) GetPublishedPosts() ([]models.Post, error) {
	return s.Repo.GetAllPublishedPosts()
}

func (s *PostService) GetPostBySlug(slug string) (*models.Post, error) {
	return s.Repo.GetPostBySlug(slug)
}

func (s *PostService) UpdatePost(id string, title, excerpt, coverImage string) (*models.Post, error) {
	updates := map[string]interface{}{}

	if title != "" {
		updates["title"] = title
		updates["slug"] = s.uniqueSlug(title)
	}
	if excerpt != "" {
		updates["excerpt"] = excerpt
	}
	if coverImage != "" {
		updates["cover_image"] = coverImage
	}

	return s.Repo.UpdatePost(id, updates)
}

func (s *PostService) PublishPost(id string) (*models.Post, error) {
	post, err := s.Repo.GetPostByID(id)
	if err != nil {
		return nil, err
	}

	// Collect readable text from text and audio blocks ordered by position
	sort.Slice(post.Blocks, func(i, j int) bool {
		return post.Blocks[i].Position < post.Blocks[j].Position
	})

	var fullText strings.Builder
	for _, block := range post.Blocks {
		if block.Type == "text" || block.Type == "audio" {
			// Strip HTML tags from text blocks before sending to ElevenLabs
			text := utils.StripHTML(block.Content)
			if text != "" {
				fullText.WriteString(text + " ")
			}
		}
	}

	audioURL, err := utils.GenerateElevenLabsAudio(strings.TrimSpace(fullText.String()))
	if err != nil {
		return nil, fmt.Errorf("failed to generate audio: %w", err)
	}

	updated, err := s.Repo.UpdatePost(id, map[string]interface{}{
		"status":               "published",
		"elevenlabs_audio_url": audioURL,
	})
	if err != nil {
		return nil, err
	}

	updated.Blocks = post.Blocks
	return updated, nil
}

func (s *PostService) DeletePost(id string) error {
	return s.Repo.DeletePost(id)
}

// uniqueSlug generates a slug from title and appends a counter if it already exists.
func (s *PostService) uniqueSlug(title string) string {
	base := utils.GenerateSlug(title)
	slug := base
	counter := 1
	for s.Repo.SlugExists(slug) {
		slug = fmt.Sprintf("%s-%d", base, counter)
		counter++
	}
	return slug
}
