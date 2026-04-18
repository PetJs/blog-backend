package repository

import (
	"gorm.io/gorm"

	"github.com/PetJs/blog-backend/internal/models"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{DB: db}
}

func (r *AdminRepository) CreateAdmin(admin *models.Admin) error {
	return r.DB.Create(admin).Error
}

func (r *AdminRepository) GetAdminByEmail(email string) (*models.Admin, error) {
	var admin models.Admin
	if err := r.DB.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
