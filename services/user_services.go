package services

import (
	"dheepssupreme/go-microservis.git/models"
	"dheepssupreme/go-microservis.git/repositories"
)

// UserService interface untuk mendefinisikan layanan user
type UserService interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user *models.User) error
}

// userService struct yang mengimplementasikan UserService
type userService struct{}

// NewUserService membuat instance baru dari UserService
func NewUserService() UserService {
	return &userService{}
}

// GetAllUsers mengambil semua data user
func (s *userService) GetAllUsers() ([]models.User, error) {
	return repositories.FindAllUsers()
}

// CreateUser menambahkan user baru
func (s *userService) CreateUser(user *models.User) error {
	return repositories.CreateUser(user)
}
