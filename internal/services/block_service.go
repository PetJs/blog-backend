package services

import (
	"github.com/PetJs/blog-backend/internal/models"
	"github.com/PetJs/blog-backend/internal/repository"
)

type BlockService struct {
	Repo *repository.BlockRepository
}

func NewBlockService(repo *repository.BlockRepository) *BlockService {
	return &BlockService{Repo: repo}
}

func (s *BlockService) AddBlock(postID uint, blockType, content, originalAudioURL string, position int) (*models.Block, error) {
	block := &models.Block{
		PostID:           postID,
		Type:             blockType,
		Content:          content,
		OriginalAudioURL: originalAudioURL,
		Position:         position,
	}
	return s.Repo.CreateBlock(block)
}

func (s *BlockService) UpdateBlock(id string, updates map[string]interface{}) (*models.Block, error) {
	return s.Repo.UpdateBlock(id, updates)
}

func (s *BlockService) DeleteBlock(id string) error {
	return s.Repo.DeleteBlock(id)
}
