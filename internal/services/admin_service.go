package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/PetJs/blog-backend/internal/models"
	"github.com/PetJs/blog-backend/internal/repository"
)

type AdminService struct {
	Repo *repository.AdminRepository
}

func NewAdminService(repo *repository.AdminRepository) *AdminService {
	return &AdminService{Repo: repo}
}

func (s *AdminService) LoginAdmin(email, password string) (*models.Admin, error) {
	admin, err := s.Repo.GetAdminByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return admin, nil
}

// SeedAdmin creates the admin account if it doesn't exist yet.
func (s *AdminService) SeedAdmin(email, password string) error {
	_, err := s.Repo.GetAdminByEmail(email)
	if err == nil {
		// Admin already exists
		return nil
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin := &models.Admin{
		Email:    email,
		Password: string(hashed),
	}
	return s.Repo.CreateAdmin(admin)
}
