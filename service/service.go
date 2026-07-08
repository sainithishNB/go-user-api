package service

import (
	"user-api/models"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetUsers() ([]models.User, error)
	GetUserByID(id int) (models.User, error)
	UpdateUser(id int, user models.User) (models.User, error)
	DeleteUser(id int) error
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}
func (s *UserService) CreateUser(user models.User) (models.User, error) {
	return s.repo.CreateUser(user)

}
func (s *UserService) GetUsers() ([]models.User, error) {
	return s.repo.GetUsers()

}
func (s *UserService) GetUserByID(id int) (models.User, error) {
	return s.repo.GetUserByID(id)

}

func (s *UserService) UpdateUser(id int, user models.User) (models.User, error) {
	return s.repo.UpdateUser(id, user)

}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)

}
