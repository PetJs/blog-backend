package repository

import (
	"gorm.io/gorm"

	"github.com/PetJs/blog-backend/internal/models"
)

type BlockRepository struct {
	DB *gorm.DB
}

func NewBlockRepository(db *gorm.DB) *BlockRepository {
	return &BlockRepository{DB: db}
}

func (r *BlockRepository) CreateBlock(block *models.Block) (*models.Block, error) {
	if err := r.DB.Create(block).Error; err != nil {
		return nil, err
	}
	return block, nil
}

func (r *BlockRepository) UpdateBlock(id string, updates map[string]interface{}) (*models.Block, error) {
	var block models.Block
	if err := r.DB.First(&block, id).Error; err != nil {
		return nil, err
	}
	if err := r.DB.Model(&block).Updates(updates).Error; err != nil {
		return nil, err
	}
	return &block, nil
}

func (r *BlockRepository) DeleteBlock(id string) error {
	var block models.Block
	if err := r.DB.First(&block, id).Error; err != nil {
		return err
	}
	return r.DB.Delete(&block).Error
}
